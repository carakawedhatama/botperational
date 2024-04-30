package on_birthday

import "context"

type Repository interface {
	GetOnBirthdayEmployee(ctx context.Context) ([]*OnBirthday, error)
}
