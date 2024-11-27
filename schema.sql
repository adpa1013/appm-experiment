CREATE TYPE month_enum AS ENUM (
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

CREATE TYPE schedule_type_enum AS ENUM (
  'Daily',
  'Weekly',
  'Monthly',
  'Yearly'
);

CREATE TABLE appointment_series (
  series_id SERIAL PRIMARY KEY,
  appointment_id integer NOT NULL,
  schedule_type schedule_type_enum NOT NULL,
  end_time timestamp,
  day_week_schedule smallint,
  month_schedule integer,
  year_schedule integer,
  month_of_year month_enum
);

CREATE TABLE excluded_from_series (
  excluded_from_series_id SERIAL PRIMARY KEY,
  appointment_id integer NOT NULL,
  dec_appointment_id integer NOT NULL,
  source_series_appointment_time timestamp
);

CREATE TABLE appointment (
  appointment_id SERIAL PRIMARY KEY ,
  name text,
  start_time timestamp,
  end_time timestamp
);

ALTER TABLE appointment_series
  ADD FOREIGN KEY (appointment_id)
  REFERENCES appointment (appointment_id);

ALTER TABLE excluded_from_series
  ADD FOREIGN KEY (appointment_id)
  REFERENCES appointment (appointment_id);

ALTER TABLE excluded_from_series
  ADD FOREIGN KEY (dec_appointment_id)
  REFERENCES appointment (appointment_id);