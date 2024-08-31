package repository

type Repository interface {
	Get() (any, error)
	Find() (any, error)
	Save() error
	Update() error
	Delete() (any, error)
}
