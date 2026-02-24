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
	// String com os dados de conexão do PostgreSQL
	dsn := "host=localhost user=postgres password=1746 dbname=travel_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	// Cria um logger personalizado do GORM
	newLogger := logger.New(

		// Define que os logs serão exibidos no terminal (stdout)
		log.New(os.Stdout, "\r\n", log.LstdFlags),

		// Configurações do logger
		logger.Config{
			SlowThreshold: time.Second, // Marca como "lenta" query que passar de 1s
			LogLevel:      logger.Info, // Mostra todas as queries executadas
			Colorful:      true,        // Deixa o log colorido no terminal
		},
	)

	// Abre conexão com o banco usando o driver do Postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // Usa o logger criado acima
	})

	// Verifica se deu erro ao conectar
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err) // Encerra o programa se falhar
	}

	// Mostra mensagem se conectou com sucesso
	fmt.Println(" Conectado ao PostgresSQL!")

	// Cria ou atualiza automaticamente a tabela User no banco
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Erro ao migrar tabela: %v", err) // Encerra se der erro na migração
	}
}
