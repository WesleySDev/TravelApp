package main

import (
	"fmt"
	"log"
	"os"
	"time"

	model "github.com/WesleySDev/travel-app/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// String de conexão com o PostgreSQL
	dsn := "host=localhost user=postgres password=1746 dbname=travel_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	// Logger para mostrar as queries SQL no terminal
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // avisa se query demorar
			LogLevel:      logger.Info, // mostra todas as queries
			Colorful:      true,        // log colorido
		},
	)

	// Abre conexão com o banco
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}

	fmt.Println("✅ Conectado ao PostgreSQL!")

	// Cria/atualiza tabela User automaticamente
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Erro ao migrar tabela: %v", err)
	}
}
