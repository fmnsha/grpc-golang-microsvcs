package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/metadata"
)

func NewPool() (*pgxpool.Pool, error) {
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, err
	}

	config.BeforeConnect = func(ctx context.Context, cfg *pgx.ConnConfig) error {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println("metadata not available")
			return errors.New("metadata not available")
		}
		fmt.Println(md)
		clients := md["client"]
		if len(clients) == 0 {
			fmt.Println("client not available")
			return errors.New("client not available")
		}
		client := md["client"][0]
		cfg.Database = client
		return nil
	}

	// config.BeforeAcquire = func(ctx context.Context, conn *pgx.Conn) bool {
	// 	// set the tenant id into this connection's setting
	// 	//tenantID := ctx.Value("tenant_id").(string)
	// 	md, ok := metadata.FromIncomingContext(ctx)
	// 	if !ok {
	// 		fmt.Println("metadata not available")
	// 		return false
	// 	}
	// 	fmt.Println(md)
	// 	clients := md["client"]
	// 	if len(clients) == 0 {
	// 		fmt.Println("client not available")
	// 		return false
	// 	}
	// 	client := md["client"][0]
	// 	fmt.Println("the client is", client)

	// 	_, err := conn.Exec(ctx, "select set_tenant($1)", client)
	// 	if err != nil {
	// 		fmt.Println("the err is", err)
	// 		return false // or better to log the error, and then `return false` to destroy this connection instead of leaving it open.
	// 	}
	// 	return true
	// }

	config.AfterRelease = func(conn *pgx.Conn) bool {

		fmt.Println("connection after release")
		// set the setting to be empty before this connection is released to pool
		// _, err := conn.Exec(context.Background(), "select set_tenant($1)", "")
		// if err != nil {
		// 	panic(err) // or better to log the error, and then`return false` to destroy this connection instead of leaving it open.
		// }
		return true
	}

	// config.MaxConns = int32(20)
	// config.MaxConnLifetime = time.Minute
	// config.MaxConnIdleTime = time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	return pool, err
}
