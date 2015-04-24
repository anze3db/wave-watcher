CREATE TABLE locations (
     lid    serial primary key,
     name   varchar(255)
);

INSERT INTO locations (name) VALUES ('Medulin');

CREATE TABLE readings (
     rid    serial primary key,
     lid    integer references locations(lid) not null,
     last_update timestamp not null,
     next_update timestamp not null,
     sun_rise    timestamp not null,
     sun_set     timestamp not null

);

CREATE TABLE forecasts (
     fid    serial primary key
);
