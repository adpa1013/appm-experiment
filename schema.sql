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
  dtmEnd datetime,
  kDayWeekSchedule tinyint,
  kMonthSchedule int,
  kYearSchedule int,
  eMonthOfYear Month
);

CREATE TABLE tblExcludedFromSeries (
  kDecAppointmentId integer PRIMARY KEY,
  kEdAppointmentId integer NOT NULL,
  dtmSrcSeriesApmt datetime
);

CREATE TABLE tblEdAppointment (
  kEdAppointmentId integer PRIMARY KEY,
  strName string,
  dtmBegin datetime,
  dtmEnd datetime
);

ALTER TABLE tblEdAppointment ADD FOREIGN KEY (kEdAppointmentId) REFERENCES tblEdAppointmentSeries (kEdAppointmentId);
ALTER TABLE tblExcludedFromSeries ADD FOREIGN KEY (kEdAppointmentId) REFERENCES tblEdAppointmentSeries (kEdAppointmentId);
ALTER TABLE tblExcludedFromSeries ADD FOREIGN KEY (kDecAppointmentId) REFERENCES tblEdAppointment (kEdAppointmentId);
