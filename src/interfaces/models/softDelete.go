package models

type ISoftDeleteModel interface {
	Delete() error
	Restore() error
	ForceDelete() error
}
