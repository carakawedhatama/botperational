package sql

import (
	"context"
	"fmt"

	"botperational/config"
	"botperational/internal/adapter/repository"
	"botperational/internal/domain/on_birthday"
	"github.com/runsystemid/golog"
)

type OnBirthdayRepository struct {
	DB  *repository.Sqlx `inject:"database"`
	Cfg *config.Config   `inject:"config"`
}

func (r *OnBirthdayRepository) Startup() error { return nil }

func (r *OnBirthdayRepository) Shutdown() error { return nil }

func (r *OnBirthdayRepository) GetOnBirthdayEmployee(ctx context.Context) ([]*on_birthday.OnBirthday, error) {
	query := `
	Select * From get_employee_on_birthday
	`
	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		golog.Error(ctx, fmt.Sprintf("Error GetOnBirthdayEmployee : %v", err.Error()), err)
		return nil, err
	}
	defer stmt.Close()

	params := map[string]interface{}{}

	rows, err := stmt.QueryxContext(ctx, params)
	if err != nil {
		golog.Error(ctx, fmt.Sprintf("Error GetOnBirthdayEmployee : %v", err.Error()), err)
		return nil, err
	}
	defer rows.Close()

	var (
		_empName           string
		_deptName          string
		_posName           string
		_birthDayPosterUrl string
		_age               int
		_gender            string
	)
	datas := make([]*on_birthday.OnBirthday, 0)

	for rows.Next() {
		err = rows.Scan(
			&_empName,
			&_deptName,
			&_posName,
			&_birthDayPosterUrl,
			&_age,
			&_gender,
		)

		if err != nil {
			golog.Error(ctx, fmt.Sprintf("Error GetOnBirthdayEmployee : %v", err.Error()), err)
			return nil, err
		}

		data := &on_birthday.OnBirthday{
			EmpName:           _empName,
			DeptName:          _deptName,
			PosName:           _posName,
			BirthdayPosterUrl: _birthDayPosterUrl,
			Age:               _age,
			Gender:            _gender,
		}

		datas = append(datas, data)

	}

	return datas, err

}
