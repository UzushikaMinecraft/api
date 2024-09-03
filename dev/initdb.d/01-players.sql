DROP SCHEMA IF EXISTS uzsk;
CREATE SCHEMA uzsk;
USE uzsk;

DROP TABLE IF EXISTS profile;

CREATE TABLE `profile` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `uuid` CHAR(36) NOT NULL UNIQUE,
    `initial_login_date` TIMESTAMP NOT NULL,
    `last_login_date` TIMESTAMP NOT NULL,
    `total_play_time` BIGINT NOT NULL,
    `experience` FLOAT DEFAULT 0.0, 
    `currency` INT DEFAULT 0, 
    `total_build_blocks` INT DEFAULT 0,
    `total_destroy_blocks` INT DEFAULT 0, 
    `total_mob_kills` INT DEFAULT 0
);

INSERT INTO `profile` (`uuid`, `initial_login_date`, `last_login_date`, `total_play_time`, `experience`, `currency`, `total_build_blocks`, `total_destroy_blocks`, `total_mob_kills`) VALUES 
('550e8400-e29b-41d4-a716-446655440000', '2024-08-25 12:00:00', '2024-08-30 16:45:00', 7200, 1500, 100, 200, 100, 50),
('e29b5500-41d4-4a71-8400-446655440001', '2024-08-26 14:30:00', '2024-08-29 10:30:00', 5400, 1200, 150, 250, 150, 75),
('d441b400-7164-550e-29b8-446655440002', '2024-08-27 09:15:00', '2024-08-28 18:20:00', 3600, 800, 200, 150, 75, 25),
('a716e400-8400-29b5-41d4-446655440003', '2024-08-28 11:00:00', '2024-08-30 14:00:00', 4800, 1100, 80, 175, 120, 40),
('29b54100-4466-55e4-8400-550ea7160004', '2024-08-29 13:45:00', '2024-08-30 15:30:00', 6000, 1400, 120, 220, 160, 65);
