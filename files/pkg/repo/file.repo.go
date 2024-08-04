package repo

import (
	"common/grpcs/file/pb"
	"context"
	"files-microservice/pkg/database"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type FileRepo struct {
}

func NewFileRepo() *FileRepo {
	return &FileRepo{}
}

func (f *FileRepo) SaveFile(ctx context.Context, file *pb.File) error {

	sql := `INSERT INTO "files" (path,originalName) VALUES ($1,$2) RETURNING id;`

	pool, err := database.NewPool()

	if err != nil {
		return err
	}

	var id uint64
	if err := pool.QueryRow(ctx, sql, file.Path, file.OriginalName).Scan(&id); err != nil {
		return err
	}

	file.Id = id

	return nil
}

func (f *FileRepo) GetById(ctx context.Context, id uint64) (*pb.File, error) {
	sql := `SELECT * FROM "files" WHERE id=$1;`

	pool, err := database.NewPool()
	if err != nil {
		return nil, err
	}

	var (
		file      pb.File
		createdAt time.Time
	)

	if err := pool.QueryRow(ctx, sql, id).Scan(
		&file.Id,
		&file.Path,
		&file.OriginalName,
		&createdAt,
		&file.CreatedBy,
	); err != nil {
		return nil, err
	}

	//todo: set file path

	file.CreatedAt = timestamppb.New(createdAt)

	return &file, nil

}

func (f *FileRepo) Listfiles(ctx context.Context) ([]*pb.File, error) {

	sql := `SELECT * FROM "files"`

	pool, err := database.NewPool()
	if err != nil {
		return nil, err
	}

	rows, errQ := pool.Query(ctx, sql)

	if errQ != nil {
		return nil, errQ
	}

	defer rows.Close()

	var files []*pb.File
	for rows.Next() {
		var (
			file      pb.File
			createdAt time.Time
		)
		if err := rows.Scan(
			&file.Id,
			&file.Path,
			&file.OriginalName,
			&createdAt,
			&file.CreatedBy,
		); err != nil {
			return nil, err
		}

		file.CreatedAt = timestamppb.New(createdAt)
		files = append(files, &file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil

}
