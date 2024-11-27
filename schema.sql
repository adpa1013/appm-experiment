CREATE TYPE Month AS ENUM (
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December'
);

CREATE TYPE ScheduleType AS ENUM (
  'Daily',
  'Weekly',
  'Monthly',
  'Yearly'
);

CREATE TABLE tblEdAppointmentSeries (
  kEdAppointmentId integer PRIMARY KEY,
  eScheduleType ScheduleType NOT NULL,
  dtmEnd timestamp,
  kDayWeekSchedule smallint,
  kMonthSchedule integer,
  kYearSchedule integer,
  eMonthOfYear Month
);

CREATE TABLE tblExcludedFromSeries (
  kDecAppointmentId integer PRIMARY KEY,
  kEdAppointmentId integer NOT NULL,
  dtmSrcSeriesApmt timestamp
);

CREATE TABLE tblEdAppointment (
  kEdAppointmentId integer PRIMARY KEY,
  strName text,
  dtmBegin timestamp,
  dtmEnd timestamp
);

ALTER TABLE tblEdAppointment
  ADD FOREIGN KEY (kEdAppointmentId)
  REFERENCES tblEdAppointmentSeries (kEdAppointmentId);

ALTER TABLE tblExcludedFromSeries
  ADD FOREIGN KEY (kEdAppointmentId)
  REFERENCES tblEdAppointmentSeries (kEdAppointmentId);

ALTER TABLE tblExcludedFromSeries
  ADD FOREIGN KEY (kDecAppointmentId)
  REFERENCES tblEdAppointment (kEdAppointmentId);
