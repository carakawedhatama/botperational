package sql

import (
	"context"
	"fmt"

	"botperational/config"
	"botperational/internal/adapter/repository"
	"botperational/internal/domain/on_leave"

	"github.com/runsystemid/golog"
)

type OnLeaveRepository struct {
	DB  *repository.Sqlx `inject:"database"`
	Cfg *config.Config   `inject:"config"`
}

func (r *OnLeaveRepository) Startup() error { return nil }

func (r *OnLeaveRepository) Shutdown() error { return nil }

func (r *OnLeaveRepository) GetOnLeaveEmployee(ctx context.Context) ([]*on_leave.OnLeave, error) {
	query := `
	Select * From get_employee_on_leave
	`
	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		golog.Error(ctx, fmt.Sprintf("Error GetOnLeaveEmployee : %v", err.Error()), err)
		return nil, err
	}
	defer stmt.Close()

	params := map[string]interface{}{}

	rows, err := stmt.QueryxContext(ctx, params)
	if err != nil {
		golog.Error(ctx, fmt.Sprintf("Error GetOnLeaveEmployee : %v", err.Error()), err)
		return nil, err
	}
	defer rows.Close()

	var (
		_empName  string
		_deptName string
	)
	datas := make([]*on_leave.OnLeave, 0)

	for rows.Next() {
		err = rows.Scan(
			&_empName,
			&_deptName,
		)

		if err != nil {
			golog.Error(ctx, fmt.Sprintf("Error GetOnLeaveEmployee : %v", err.Error()), err)
			return nil, err
		}

		data := &on_leave.OnLeave{
			EmpName:  _empName,
			DeptName: _deptName,
		}

		datas = append(datas, data)

	}

	return datas, err

}
