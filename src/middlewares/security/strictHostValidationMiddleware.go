package middlewares

import (
	"base-go-app/src/common/utils/environment"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func StrictHostValidationMiddleware() gin.HandlerFunc {
	allowedHosts := environment.GetAsSlice("ALLOWED_HOSTS", []string{""})
	// Build the regex pattern for allowed hosts
	allowedHostPattern := buildAllowedHostPattern(allowedHosts)
	// Define the list of suspicious characters
	suspiciousCharacters := []string{
		"@", "..", "%", "\\", "`", "\"", "<", ">", "{", "}",
		"|", "^", "~", "[", "]", "(", ")", ";", "&", "$", "#", "'",
	}
	return func(c *gin.Context) {
		host := c.Request.Host
		// Validate the host against the allowed host pattern
		if !allowedHostPattern.MatchString(host) || containsSuspiciousCharacters(host, suspiciousCharacters) {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}

}

func buildAllowedHostPattern(allowedHosts []string) *regexp.Regexp {
	// Escape dots and join hosts into a regex pattern
	for i, host := range allowedHosts {
		allowedHosts[i] = regexp.QuoteMeta(host)
	}
	pattern := "^(" + strings.Join(allowedHosts, "|") + ")(\\:\\d+)?$"
	return regexp.MustCompile(pattern)
}

func containsSuspiciousCharacters(host string, suspiciousCharacters []string) bool {
	for _, character := range suspiciousCharacters {
		if strings.Contains(host, character) {
			return true
		}
	}
	return false
}
