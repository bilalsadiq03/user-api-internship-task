package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/bilalsadiq03/user-api-internship-task/db/sqlc"
)

type UserRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: sqlc.New(db),
	}
}


// Create User
func (r *UserRepository) Create(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.CreateUserRow, error) {

	return r.db.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}


// Get User By ID
func (r *UserRepository) GetByID(
	ctx context.Context,
	id int32,
) (sqlc.GetUserByIDRow, error) {
	return r.db.GetUserByID(ctx, id)
}

// Get All Users
func (r *UserRepository) ListPaginated(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]sqlc.ListUsersPaginatedRow, error) {
	return r.db.ListUsersPaginated(ctx, sqlc.ListUsersPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
}

// Delete User By ID
func (r *UserRepository) DeleteByID(
	ctx context.Context,
	id int32,
) error	 {
	return r.db.DeleteUser(ctx, id)
}


// Update User
func (r *UserRepository) UpdateByID(
	ctx context.Context,
	id int32,
	name string,
	dob time.Time,
) (sqlc.UpdateUserRow, error){
	return r.db.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}