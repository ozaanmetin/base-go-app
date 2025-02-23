package models

type IActivenessModel interface {
	Activate() error
	Deactivate() error
}
