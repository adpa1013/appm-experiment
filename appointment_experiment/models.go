// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package appointment_experiment

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type MonthEnum string

const (
	MonthEnumJanuary   MonthEnum = "January"
	MonthEnumFebruary  MonthEnum = "February"
	MonthEnumMarch     MonthEnum = "March"
	MonthEnumApril     MonthEnum = "April"
	MonthEnumMay       MonthEnum = "May"
	MonthEnumJune      MonthEnum = "June"
	MonthEnumJuly      MonthEnum = "July"
	MonthEnumAugust    MonthEnum = "August"
	MonthEnumSeptember MonthEnum = "September"
	MonthEnumOctober   MonthEnum = "October"
	MonthEnumNovember  MonthEnum = "November"
	MonthEnumDecember  MonthEnum = "December"
)

func (e *MonthEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = MonthEnum(s)
	case string:
		*e = MonthEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for MonthEnum: %T", src)
	}
	return nil
}

type NullMonthEnum struct {
	MonthEnum MonthEnum
	Valid     bool // Valid is true if MonthEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMonthEnum) Scan(value interface{}) error {
	if value == nil {
		ns.MonthEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.MonthEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMonthEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.MonthEnum), nil
}

type ScheduleTypeEnum string

const (
	ScheduleTypeEnumDaily   ScheduleTypeEnum = "Daily"
	ScheduleTypeEnumWeekly  ScheduleTypeEnum = "Weekly"
	ScheduleTypeEnumMonthly ScheduleTypeEnum = "Monthly"
	ScheduleTypeEnumYearly  ScheduleTypeEnum = "Yearly"
)

func (e *ScheduleTypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ScheduleTypeEnum(s)
	case string:
		*e = ScheduleTypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for ScheduleTypeEnum: %T", src)
	}
	return nil
}

type NullScheduleTypeEnum struct {
	ScheduleTypeEnum ScheduleTypeEnum
	Valid            bool // Valid is true if ScheduleTypeEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullScheduleTypeEnum) Scan(value interface{}) error {
	if value == nil {
		ns.ScheduleTypeEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ScheduleTypeEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullScheduleTypeEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ScheduleTypeEnum), nil
}

type Appointment struct {
	AppointmentID int32
	Name          pgtype.Text
	StartTime     pgtype.Timestamp
	EndTime       pgtype.Timestamp
}

type AppointmentSeries struct {
	SeriesID        int32
	AppointmentID   int32
	ScheduleType    ScheduleTypeEnum
	EndTime         pgtype.Timestamp
	DayWeekSchedule pgtype.Int2
	MonthSchedule   pgtype.Int4
	YearSchedule    pgtype.Int4
	MonthOfYear     NullMonthEnum
}

type ExcludedFromSeries struct {
	ExcludedFromSeriesID        int32
	AppointmentID               int32
	DecAppointmentID            int32
	SourceSeriesAppointmentTime pgtype.Timestamp
}
