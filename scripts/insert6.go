package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jaswdr/faker"
)

func main() {
	// Instancia de faker para generar nombres
	f := faker.New()

	dsn := "admin:admin123@tcp(127.0.0.1:3306)/universidad"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Verificar conexión
	if err = db.Ping(); err != nil {
		log.Fatal("Error al hacer ping a la base de datos:", err)
	}
	fmt.Println("Conexión exitosa a MySQL")

	// Arreglo de carreras disponibles
	carreras := []string{"Ingeniería", "Medicina", "Derecho", "Arquitectura", "Psicología"}

	// Construir la consulta de inserción múltiple
	var sb strings.Builder
	sb.WriteString("INSERT INTO estudiantes (nombre, edad, carrera, promedio) VALUES ")
	valores := make([]interface{}, 0, 1000*4)

	for i := 0; i < 1000; i++ {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString("(?, ?, ?, ?)")

		// Generar datos aleatorios
		nombre := f.Person().FirstName()              // Primer Nombre aleatorio
		edad := rand.Intn(30-18+1) + 18               // Edad entre 18 y 30
		carrera := carreras[rand.Intn(len(carreras))] // Carrera aleatoria
		promedio := float64(rand.Intn(1000)) / 100.0  // Promedio entre 0.00 y 9.99

		valores = append(valores, nombre, edad, carrera, promedio)
	}

	// Ejecutar la consulta de inserción
	query := sb.String()
	result, err := db.Exec(query, valores...)
	if err != nil {
		log.Fatal("Error al insertar registros:", err)
	}
	rows, _ := result.RowsAffected()
	fmt.Printf("Se han insertado %d registros.\n", rows)
}
