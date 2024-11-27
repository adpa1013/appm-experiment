package main

import (
	"appointment-experiment/appointment_experiment"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func run(count int) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://paul:Strong@Passw0rd@localhost:5432/appointment_db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries := appointment_experiment.New(conn)

	for i := 0; i < count; i++ {
		err := queries.InsertAppointment(ctx, appointment_experiment.InsertAppointmentParams{
			Name:      pgtype.Text{String: "Richtig toller Termin", Valid: true},
			StartTime: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
			EndTime:   pgtype.Timestamp{Time: time.Now().Add(1 * time.Hour).UTC(), Valid: true},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cleanUp() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://paul:Strong@Passw0rd@localhost:5432/appointment_db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries := appointment_experiment.New(conn)
	err = queries.DeleteEverything(ctx)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	run(500)
	// for i := 0; i < 10; i++ {
	// 	go run(5000)
	// 	fmt.Println("Started run", i)
	// }

	// if err := cleanUp(); err != nil {
	// 	log.Fatal(err)
	// }
}
