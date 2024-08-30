DROP SCHEMA IF EXISTS PlayerData;
CREATE SCHEMA PlayerData;
USE PlayerData;

DROP TABLE IF EXISTS PlayerData;

CREATE TABLE PlayerData (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    UUID CHAR(36) NOT NULL UNIQUE,
    InitialLoginDateTime DATETIME NOT NULL,
    LastLoginDateTime DATETIME,
    PlayTime INT DEFAULT 0,
    ExperiencePoints INT DEFAULT 0,  -- 経験値
    Currency INT DEFAULT 0,  -- 通貨
    TotalBuildBlocks INT DEFAULT 0,  -- 総建築ブロック数
    TotalDestroyBlocks INT DEFAULT 0,  -- 総破壊ブロック数
    TotalMobKills INT DEFAULT 0  -- 総敵モブキル数
);

INSERT INTO PlayerData (UUID, InitialLoginDateTime, LastLoginDateTime, PlayTime, ExperiencePoints, Currency, TotalBuildBlocks, TotalDestroyBlocks, TotalMobKills)
VALUES
('550e8400-e29b-41d4-a716-446655440000', '2024-08-25 12:00:00', '2024-08-30 16:45:00', 7200, 1500, 100, 200, 100, 50),
('e29b5500-41d4-4a71-8400-446655440001', '2024-08-26 14:30:00', '2024-08-29 10:30:00', 5400, 1200, 150, 250, 150, 75),
('d441b400-7164-550e-29b8-446655440002', '2024-08-27 09:15:00', '2024-08-28 18:20:00', 3600, 800, 200, 150, 75, 25),
('a716e400-8400-29b5-41d4-446655440003', '2024-08-28 11:00:00', '2024-08-30 14:00:00', 4800, 1100, 80, 175, 120, 40),
('29b54100-4466-55e4-8400-550ea7160004', '2024-08-29 13:45:00', '2024-08-30 15:30:00', 6000, 1400, 120, 220, 160, 65);
