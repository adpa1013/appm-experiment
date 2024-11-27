-- name: GetAppointments :many

SELECT *
FROM appointment;

-- name: InsertAppointment :exec

INSERT INTO appointment (name, start_time, end_time)
VALUES ($1,
        $2,
        $3);

-- name: BatchInsertAppointment :copyfrom

INSERT INTO appointment (name, start_time, end_time)
VALUES ($1,
        $2,
        $3);

-- name: DeleteEverything :exec

DELETE FROM appointment;
DELETE FROM appointment_series;
DELETE FROM excluded_from_series;