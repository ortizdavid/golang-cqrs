package repositories

type Repository[T any] interface {
	Create(entity T) error
	GetAll() ([]T, error)
	GetById(id int) (T, error)
	GetByUniqueId(id int) (T, error)
	Exists(predicate string) bool
	Update(entity T) error
	Delete(entity T)
}