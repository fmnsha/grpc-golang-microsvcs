package repo

import (
	"common/grpcs/user/pb"
	"context"
	"fmt"
	"users-microservice/pkg/database"

	"github.com/lib/pq"
	"github.com/samber/do"
)

type CapabilityRepo interface {
	CreateCap(ctx context.Context, capability *pb.Capability) error
	ReadCap(ctx context.Context, id uint64) (*pb.Capability, error)
	UpdateCap(ctx context.Context, capability *pb.Capability) error
	DeleteCap(ctx context.Context, id uint64) error
	ListCap() ([]*pb.Capability, error)
	//
	ListCapByIds(ctx context.Context, id ...uint64) ([]*pb.Capability, error)
	AddRoleCapabilities(ctx context.Context, roleCapability *pb.RoleCapability) error
	ReadUserCapabilities(ctx context.Context, userid uint64) ([]string, error)
}
type capabilityrepo struct{}

func NewCapabilityRepo(i *do.Injector) (CapabilityRepo, error) {
	return &capabilityrepo{}, nil
}

func (c *capabilityrepo) CreateCap(ctx context.Context, capability *pb.Capability) error {
	newCapabilityStatment := `
	INSERT INTO "capabilities" (
		label, 
		name, 
		reserved
		) 
	VALUES ($1, $2, $3) RETURNING id;
	`

	pool, errGetPool := database.NewPool()
	if errGetPool != nil {
		return errGetPool
	}
	var id uint64
	err := pool.QueryRow(ctx, newCapabilityStatment,
		capability.Label,
		capability.Name,
		capability.Reserved,
	).Scan(&id)

	if err != nil {
		return err
	}

	capability.Id = id

	return nil

}
func (c *capabilityrepo) ReadCap(ctx context.Context, id uint64) (*pb.Capability, error) {
	return nil, nil
}
func (c *capabilityrepo) UpdateCap(ctx context.Context, capability *pb.Capability) error {
	return nil
}
func (c *capabilityrepo) DeleteCap(ctx context.Context, id uint64) error {
	return nil
}

func (c *capabilityrepo) ListCap() ([]*pb.Capability, error) {
	return nil, nil
}

func (c *capabilityrepo) ListCapByIds(ctx context.Context, id ...uint64) ([]*pb.Capability, error) {
	query := `SELECT  * FROM capabilities WHERE id= ANY($1)`

	pool, errGetPool := database.NewPool()
	if errGetPool != nil {
		return nil, errGetPool
	}

	rows, err := pool.Query(context.Background(), query, pq.Array(id))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	capabilities := []*pb.Capability{}
	for rows.Next() {
		cap := &pb.Capability{}
		if err := rows.Scan(&cap.Id, &cap.Name, &cap.Label, &cap.Reserved); err != nil {
			return nil, err
		}
		capabilities = append(capabilities, cap)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return capabilities, nil
}

// /
func (c *capabilityrepo) AddRoleCapabilities(ctx context.Context, roleCapability *pb.RoleCapability) error {
	sql := `
		INSERT INTO "role_capabilities" (role_id,capability_id) VALUES ($1,$2)
	`
	pool, err := database.NewPool()

	if err != nil {
		return err
	}

	result, err := pool.Exec(ctx, sql, roleCapability.Roleid, roleCapability.Capabilityid)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(result)

	return nil
}

func (c *capabilityrepo) ReadUserCapabilities(ctx context.Context, userid uint64) ([]string, error) {
	query := `SELECT capabilities.name FROM role_capabilities JOIN capabilities ON capabilities.id = role_capabilities.capability_id WHERE role_id = ANY(SELECT role_id FROM user_roles WHERE user_id = $1)  ;`

	pool, err := database.NewPool()
	if err != nil {
		return nil, err
	}

	rows, err := pool.Query(ctx, query, userid)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	capabilities := []string{}
	for rows.Next() {
		var capability string
		if err := rows.Scan(&capability); err != nil {
			return nil, err
		}

		capabilities = append(capabilities, capability)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(capabilities)

	return capabilities, nil

}
