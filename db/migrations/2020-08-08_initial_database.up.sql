-- dataseries contains the time series configuration data.
CREATE TABLE dataseries
(
    id          int(10) unsigned                     NOT NULL AUTO_INCREMENT,
    created_at  timestamp                            NOT NULL DEFAULT current_timestamp(),
    updated_at  timestamp                            NOT NULL DEFAULT current_timestamp(),
    name        varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    description text COLLATE utf8_unicode_ci                  DEFAULT NULL,
    label       varchar(255) COLLATE utf8_unicode_ci          DEFAULT NULL,
    api_key     varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci;

-- dataseries_summaries is a materialized view for getting the newest value in
-- a time series.
-- The application is expected to update current_value, min_value and max_value
-- whenever a new value is added to the readings table.
CREATE TABLE dataseries_summaries
(
    id            int(10) unsigned                     NOT NULL AUTO_INCREMENT,
    created_at    timestamp                            NOT NULL DEFAULT current_timestamp(),
    updated_at    timestamp                            NOT NULL DEFAULT current_timestamp(),
    current_value varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    min_value     varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    max_value     varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    dataseries_id int(10) unsigned                     NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY dataseries_summaries_dataseries_id_unique (dataseries_id),
    CONSTRAINT dataseries_summaries_dataseries_id_foreign FOREIGN KEY (dataseries_id) REFERENCES dataseries (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci;

-- readings will contain the values received from the devices. Each value is
-- associated with a dataseries but there is no foreign key constraint in order
-- to make inserts a little bit faster.
CREATE TABLE readings
(
    id            int(10) unsigned                     NOT NULL AUTO_INCREMENT,
    created_at    timestamp                            NOT NULL DEFAULT current_timestamp(),
    updated_at    timestamp                            NOT NULL DEFAULT current_timestamp(),
    value         varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    dataseries_id int(10) unsigned                     NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci;

-- The readings are most often queried given the dataseries id and the number
-- of newest readings.
CREATE INDEX idx_readings_dataseries_created_at ON readings (dataseries_id, created_at DESC);
