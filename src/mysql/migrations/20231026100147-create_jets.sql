
-- +migrate Up
CREATE TABLE IF NOT EXISTS jets (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    pilot_id int(11) unsigned NOT NULL,
    age int NOT NULL,
    name varchar(100) NOT NULL,
    color varchar(100) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`pilot_id`) REFERENCES pilots(`id`)
);

INSERT INTO jets (pilot_id, age, name, color) VALUES (1, 20, "jetname1", "red");

-- +migrate Down
DROP TABLE IF EXISTS jets;
