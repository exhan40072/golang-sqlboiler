
-- +migrate Up
CREATE TABLE IF NOT EXISTS pilots (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO pilots (name) VALUES ("pilotname1");

-- +migrate Down
DROP TABLE IF EXISTS pilots;