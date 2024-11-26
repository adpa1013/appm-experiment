-- name: GetAppointmentsBySeries :many
SELECT ap.kEdAppointmentId, ap.strName, ap.dtmBegin, ap.dtmEnd
FROM tblEdAppointment ap
WHERE ap.kEdAppointmentId IN (
    SELECT aseries.kEdAppointmentId
    FROM tblEdAppointmentSeries aseries
    WHERE aseries.kEdAppointmentId = $1
);
