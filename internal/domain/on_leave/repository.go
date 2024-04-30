package on_leave

import "context"

type Repository interface {
	GetOnLeaveEmployee(ctx context.Context) ([]*OnLeave, error)
}
