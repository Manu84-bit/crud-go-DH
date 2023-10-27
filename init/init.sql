create DATABASE IF not EXISTS go_db;
use go_db;
CREATE TABLE if not EXISTS dentists(
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    last_name VARCHAR(255),
    license VARCHAR(255) UNIQUE KEY
);

INSERT INTO dentists(`name`, last_name, license) VALUES
('Lionel', 'Messi', 'LM134'),
('James', 'Rodríguez', 'JM234'),
('Falcao', 'García', 'FG334')
ON DUPLICATE KEY UPDATE `license` = `license`;  

CREATE TABLE if not EXISTS patients(
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    last_name VARCHAR(255),
    address VARCHAR(255),
    dni int UNIQUE KEY,
    discharge_date VARCHAR(255)
);

INSERT INTO patients (`name`, last_name, `address`, dni, discharge_date) VALUES
('Miguel', 'Casale', 'Av. Siempre viva 34', 45678, '2023-03-30'),
('Manuel', 'Casas', 'Av. Siempre viva 54', 45675, '2023-03-31')
ON DUPLICATE KEY UPDATE `dni` = `dni`;


CREATE TABLE if not EXISTS `appointments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `patient_id` int NOT NULL,
  `dentist_id` int NOT NULL,
  `date` VARCHAR(255) UNIQUE KEY,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`patient_id`) REFERENCES `patients` (`id`),
  FOREIGN KEY (`dentist_id`) REFERENCES `dentists` (`id`)
);

INSERT INTO appointments (patient_id, dentist_id, `date`) VALUES
(1,2, '2022-08-15 14:00:45'),
(2,1, '2021-08-15 14:30:45')
ON DUPLICATE KEY UPDATE `date`= `date`
; 
