package repo

import (
	"common/grpcs/user/pb"
	"context"
	"errors"
	"fmt"
	"time"
	"users-microservice/pkg/database"
	"users-microservice/util"

	"github.com/jackc/pgx/v5"
	"github.com/samber/do"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *pb.User) (*pb.User, error)
	ReadUser(ctx context.Context, id uint64) (*pb.User, error)
	UpdateUser(ctx context.Context, user *pb.User) error
	DeleteUser(ctx context.Context, id uint64) error
	ListUser(ctx context.Context, found func(*pb.User) error) error
	GetUserByEmail(ctx context.Context, email string) (*pb.User, error)
	ResetPassword(ctx context.Context, id uint64, password string) error
}

type userrepo struct{}

func NewUserRepo(i *do.Injector) (UserRepo, error) {
	return &userrepo{}, nil
}

func (u *userrepo) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {

	password := user.GetPassword()

	hash, errHash := bcrypt.GenerateFromPassword([]byte(password), 10)

	if errHash != nil {
		return nil, errHash
	}

	newUserStatment := `
	INSERT INTO "users" (
		firstname, 
		lastname, 
		email,
		emailConfirmation,
		password
		) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id,createdAt, updatedAt;
	`

	pool, errg := database.NewPool()
	if errg != nil {
		return nil, errg
	}

	var (
		id        uint64
		createdAt time.Time
		updatedAt time.Time
	)
	err := pool.QueryRow(ctx, newUserStatment,
		user.FirstName,
		user.LastName,
		user.Email,
		user.EmailConfirmation,
		string(hash),
	).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user.Id = id
	user.CreatedAt = timestamppb.New(createdAt)
	user.UpdatedAt = timestamppb.New(updatedAt)
	return user, nil

}

func (u *userrepo) ReadUser(ctx context.Context, id uint64) (*pb.User, error) {

	query := `
	SELECT 
	id,
	firstname,
	lastname,
	email,
	emailConfirmation,
	createdAt,
	createdBy,
	updatedAt,
	updatedBy
	FROM users WHERE id=$1;
	`

	pool, err := database.NewPool()

	if err != nil {
		return nil, err
	}

	var user pb.User
	var (
		createdAt time.Time
		updatedAt time.Time
	)
	if err := pool.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.EmailConfirmation,
		&createdAt,
		&user.CreatedBy,
		&updatedAt,
		&user.UpdatedBy,
	); err != nil {
		fmt.Println(err)
		return nil, errors.New("error get user")
	}

	user.CreatedAt = timestamppb.New(createdAt)
	user.UpdatedAt = timestamppb.New(updatedAt)

	return &user, nil

}

func (u *userrepo) UpdateUser(ctx context.Context, user *pb.User) error {

	password := user.GetPassword()

	pool, err := database.NewPool()
	if err != nil {
		return err
	}

	updatedAt := time.Now()
	var id uint64
	if password == "" {
		updateStatement := `
			UPDATE users
			SET 
			firstName = $1,
			lastName = $2,
			email = $3,
			emailConfirmation = $4,
			updatedAt = $5
			WHERE id = $6 RETURNING id;
		`

		if err := pool.QueryRow(ctx, updateStatement,
			user.FirstName,
			user.LastName,
			user.Email,
			user.EmailConfirmation,
			updatedAt,
			user.Id,
		).Scan(&id); err != nil {
			fmt.Println(err)
			return err
		}

	} else {

		hash, errHash := bcrypt.GenerateFromPassword([]byte(password), 10)

		if errHash != nil {
			return errHash
		}

		updateStatement := `
			UPDATE users
			SET 
			firstName = $1,
			lastName = $2,
			email = $3,
			password = $4,
			emailConfirmation = $5,
			updatedAt = $6
			WHERE id = $7 RETURNING id;
		`

		if err := pool.QueryRow(ctx, updateStatement,
			user.FirstName,
			user.LastName,
			user.Email,
			string(hash),
			user.EmailConfirmation,
			updatedAt,
			user.Id,
		).Scan(&id); err != nil {
			fmt.Println(err)
			return err
		}

	}

	user.UpdatedAt = timestamppb.New(updatedAt)

	return nil
}

func (u *userrepo) DeleteUser(ctx context.Context, id uint64) error {
	return nil
}

func (u *userrepo) ListUser(ctx context.Context, found func(*pb.User) error) error {
	query := `
		SELECT
		*
		FROM users
		ORDER BY id
		;
	`

	pool, err := database.NewPool()
	if err != nil {
		return err
	}

	rows, err := pool.Query(ctx, query)

	if err != nil {
		return err
	}

	defer rows.Close()

	//result := []*pb.User{}
	for rows.Next() {
		user := &pb.User{}
		var (
			createdAt time.Time
			updatedAt time.Time
		)
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.EmailConfirmation,
			&user.Password,
			&createdAt,
			&user.CreatedBy,
			&updatedAt,
			&user.UpdatedBy,
		); err != nil {
			fmt.Println(err)
			return err
		}

		user.CreatedAt = timestamppb.New(createdAt)
		user.UpdatedAt = timestamppb.New(updatedAt)
		user.Password = ""

		if err := found(user); err != nil {
			return err
		}

		//result = append(result, user)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (u *userrepo) GetUserByEmail(ctx context.Context, email string) (*pb.User, error) {
	query := `
	SELECT 
	Id,
	firstName,
	lastName,
	email,
	emailconfirmation,
	password
	FROM users WHERE email=$1;
	`
	pool, errGetPool := database.NewPool()
	if errGetPool != nil {
		return nil, errGetPool
	}
	var user pb.User
	if err := pool.QueryRow(ctx, query, email).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.EmailConfirmation,
		&user.Password,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u *userrepo) ResetPassword(ctx context.Context, id uint64, password string) error {

	pool, err := database.NewPool()
	if err != nil {
		return err
	}

	hash, errHash := bcrypt.GenerateFromPassword([]byte(password), 10)

	if errHash != nil {
		return errHash
	}

	updateStatement := `
			UPDATE users
			SET 
			password = $1,
			updatedAt = $2
			WHERE id = $3;
		`

	if _, err := pool.Exec(ctx, updateStatement,
		hash,
		time.Now(),
		id,
	); err != nil {
		return err
	}

	return nil

}
