package repo

import (
	"common/grpcs/user/pb"
	"context"
	"fmt"
	"users-microservice/pkg/database"

	"github.com/samber/do"
)

type RoleRepo interface {
	CreateRole(ctx context.Context, role *pb.Role) (*pb.Role, error)
	ReadRole(ctx context.Context, id uint64) (*pb.Role, error)
	UpdateRole(ctx context.Context, role *pb.Role) error
	DeleteRole(ctx context.Context, id uint64) error
	ListRole(ctx context.Context, found func(*pb.Role) error) error
	//
	AddUserRole(ctx context.Context, data *pb.UserRole) error
	DeleteUserRole(ctx context.Context, id uint64) error
}

type rolerepo struct{}

func NewRoleRepo(i *do.Injector) (RoleRepo, error) {
	return &rolerepo{}, nil
}

func (r *rolerepo) CreateRole(ctx context.Context, role *pb.Role) (*pb.Role, error) {

	newRoleStatment := `
	INSERT INTO "roles" (
		label, 
		name, 
		reserved
		) 
	VALUES ($1, $2, $3) RETURNING id;
	`

	pool, errGetPool := database.NewPool()
	if errGetPool != nil {
		return nil, errGetPool
	}

	var id uint64
	err := pool.QueryRow(ctx, newRoleStatment,
		role.Label,
		role.Name,
		role.Reserved,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	role.Id = id

	return role, nil

}
func (r *rolerepo) ReadRole(ctx context.Context, id uint64) (*pb.Role, error) {
	return nil, nil
}
func (r *rolerepo) UpdateRole(ctx context.Context, role *pb.Role) error {

	updateStatement := `
		UPDATE roles
		SET 
		label = $1, 
		name = $2
		WHERE id = $3 RETURNING id;
	`
	pool, errGetPool := database.NewPool()
	if errGetPool != nil {
		return errGetPool
	}

	var id uint64
	err := pool.QueryRow(ctx, updateStatement,
		role.Label,
		role.Name,
		role.Id,
	).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}
func (r *rolerepo) DeleteRole(ctx context.Context, id uint64) error {
	return nil
}
func (r *rolerepo) ListRole(ctx context.Context, found func(*pb.Role) error) error {
	query := `
		SELECT * FROM roles ORDER BY id;
	`

	pool, errGetPool := database.NewPool()
	if errGetPool != nil {
		return errGetPool
	}
	rows, err := pool.Query(ctx, query)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		role := &pb.Role{}
		if err := rows.Scan(
			&role.Id,
			&role.Name,
			&role.Label,
			&role.Reserved,
		); err != nil {
			return err
		}

		if err := found(role); err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

//

func (r *rolerepo) AddUserRole(ctx context.Context, data *pb.UserRole) error {

	query := `
		INSERT INTO "user_roles" (user_id, role_id) VALUES (?,?) RETURNING id;
	`

	pool, err := database.NewPool()
	if err != nil {
		return err
	}

	//var id uint64
	if _, err := pool.Exec(ctx, query, data.Userid, data.Roleid); err != nil {
		fmt.Println(err)
		return err
	}

	//data.Id = id

	return nil
}
func (r *rolerepo) DeleteUserRole(ctx context.Context, id uint64) error {

	query := `
		DELETE FROM "user_roles" WHERE id = $1 RETURNING id;
	`

	pool, err := database.NewPool()
	if err != nil {
		return err
	}
	var _id uint64
	if err := pool.QueryRow(ctx, query, id).Scan(&_id); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(_id)

	return nil
}
