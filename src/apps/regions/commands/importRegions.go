package main

import (
	"base-go-app/config/settings/environment"
	"base-go-app/src/apps/regions/models"
	"base-go-app/src/database"
	"log"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

// Define worker count for parallel processing
const workerCount = 10

func main() {
	// Initialize environment variables and establish database connection
	environment.InitalizeDotEnv()
	database.ConnectPostgres()

	// Get the base directory of the app
	appBaseDir, err := environment.GetBaseDir()
	if err != nil {
		log.Fatalf("Error getting base directory: %v", err)
	}

	// Define the path to the Excel file containing region data
	filePath := filepath.Join(appBaseDir, "src", "apps", "regions", "assets", "turkey_regions.xlsx")
	f, err := excelize.OpenFile(filePath) // Open the Excel file
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer f.Close() // Ensure the file is closed after processing

	// Fetch all rows from the first sheet
	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Fatalf("Error getting rows: %v", err)
	}

	// Create a JSON object for the name of the country
	if err != nil {
		log.Fatalf("Error dumping map as json: %v", err)
	}

	countryName := map[string]interface{}{
		"en": "Turkey",
		"tr": "Türkiye",
	}
	// Create or fetch the country "Turkey" if it doesn't exist
	var country models.Country
	err = database.PostgresContext.Where("code = ?", "tr").FirstOrCreate(&country, models.Country{
		Name: countryName,
		Code: "tr",
	}).Error
	if err != nil {
		log.Fatalf("Error creating country: %v", err)
	}

	// GOROUTINES BLOCK
	// -----------------------------
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create channels for rows and errors
	rowChan := make(chan []string, workerCount)
	errorChan := make(chan error, workerCount)

	// Create workers for concurrent processing
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, rowChan, errorChan, country.ID)
	}

	// Push rows to the channel for processing
	go func() {
		defer close(rowChan)
		for _, row := range rows {
			// Skip rows that have fewer than 3 values
			if len(row) < 3 {
				log.Println("Row length is less than 3, Skipping...")
				continue
			}
			select {
			case rowChan <- row: // Send row for processing
			case <-ctx.Done(): // Stop if context is cancelled
				return
			}
		}
	}()

	// Handle errors
	go func() {
		defer close(errorChan)
		for err := range errorChan {
			log.Printf("Error: %v", err) // Log errors
		}
	}()

	// Wait for all workers to finish processing
	wg.Wait()
}

// Worker function for processing rows concurrently
func worker(ctx context.Context, wg *sync.WaitGroup, rowChan <-chan []string, errorChan chan<- error, countryID uuid.UUID) {
	defer wg.Done()

	// Continuously receive rows from the channel and process them
	for {
		select {
		case <-ctx.Done():
			return
		case row, ok := <-rowChan:
			if !ok {
				return
			}
			// Process the row and handle any errors
			if err := processRow(ctx, row, countryID); err != nil {
				errorChan <- err // Send errors to the error channel
			}
		}
	}
}

// Function to process a row and perform database operations
func processRow(ctx context.Context, row []string, countryID uuid.UUID) error {
	// Wrap database operations in a transaction for consistency
	return database.PostgresContext.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Extract city, district, and neighborhood names from the row
		cityName := row[0]
		districtName := row[1]
		neighborhoodName := row[2]

		// Create or fetch the City
		var city models.City
		if err := tx.Where("name = ? and country_id = ?", cityName, countryID).FirstOrCreate(&city, models.City{
			Name:      cityName,
			CountryID: countryID,
		}).Error; err != nil {
			return err
		}

		// Create or fetch the District
		var district models.District
		if err := tx.Where("name = ? and city_id = ?", districtName, city.ID).FirstOrCreate(&district, models.District{
			Name:      districtName,
			CountryID: countryID,
			CityID:    city.ID,
		}).Error; err != nil {
			return err
		}

		// Create or fetch the Neighborhood
		if err := tx.Where("name = ? and district_id = ?", neighborhoodName, district.ID).FirstOrCreate(&models.Neighborhood{
			Name:       neighborhoodName,
			CountryID:  countryID,
			CityID:     city.ID,
			DistrictID: district.ID,
		}).Error; err != nil {
			return err
		}

		log.Printf("Row processed: %s, %s, %s", cityName, districtName, neighborhoodName) // Log processed row
		return nil
	})
}

// NOTE:
// Without goroutines, processing the data takes about 3.45 minutes.
// With goroutines, the processing time is reduced to around 1.02 minutes.
