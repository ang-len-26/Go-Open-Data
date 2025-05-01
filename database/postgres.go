package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL no está definido")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	DB, err = pgxpool.New(ctx, databaseUrl)
	if err != nil {
		log.Fatalf("No se pudo conectar a PostgreSQL: %v", err)
	}

	err = DB.Ping(ctx)
	if err != nil {
		log.Fatalf("No se pudo hacer ping a PostgreSQL: %v", err)
	}

	fmt.Println("✅ Conexión a PostgreSQL exitosa")
}
