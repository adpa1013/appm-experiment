package main

import (
	"appointment-experiment/appointment_experiment"
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
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
func collectAppointments(count int) []appointment_experiment.BatchInsertAppointmentParams {
	var appointments []appointment_experiment.BatchInsertAppointmentParams
	for i := 0; i < count; i++ {
		appointments = append(appointments, appointment_experiment.BatchInsertAppointmentParams{
			Name:      pgtype.Text{String: "Richtig toller Termin " + strconv.Itoa(i), Valid: true},
			StartTime: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
			EndTime:   pgtype.Timestamp{Time: time.Now().Add(1 * time.Hour).UTC(), Valid: true},
		})
	}

	return appointments
}

func runBatched(appm []appointment_experiment.BatchInsertAppointmentParams) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://paul:Strong@Passw0rd@localhost:5432/appointment_db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries := appointment_experiment.New(conn)
	_, err = queries.BatchInsertAppointment(ctx, appm)

	if err != nil {
		log.Fatal(err)
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
	var wg sync.WaitGroup // Synchronize goroutines

	start := time.Now()
	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			run(5000)
		}()
		fmt.Println("Started run", i)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Concurrent insert took:", duration.Seconds())

	start = time.Now()

	batchRuns := collectAppointments(1000000)

	var wgBatch sync.WaitGroup

	chunkSize := 50000
	for i := 0; i < len(batchRuns); i += chunkSize {
		end := i + chunkSize
		if end > len(batchRuns) {
			end = len(batchRuns)
		}

		wgBatch.Add(1)
		go func(batch []appointment_experiment.BatchInsertAppointmentParams) {
			defer wgBatch.Done()
			runBatched(batch)
		}(batchRuns[i:end])
	}

	wgBatch.Wait()

	duration = time.Since(start)
	res := fmt.Sprintf("Cuncurrently batch insert %d appointments took: %f Seconds!", len(batchRuns), duration.Seconds())
	fmt.Println(res)

	start = time.Now()

	runBatched(batchRuns)

	duration = time.Since(start)
	res = fmt.Sprintf("Batch sequential insert %d appointments took: %f Seconds!", len(batchRuns), duration.Seconds())
	fmt.Println(res)
	// if err := cleanUp(); err != nil {
	// 	log.Fatal(err)
	// }
}
