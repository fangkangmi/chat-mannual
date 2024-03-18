CREATE TABLE `account` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `owner` VARCHAR(255) NOT NULL,
  `balance` BIGINT NOT NULL,
  `currency` VARCHAR(255) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `entries` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `account_id` INT NOT NULL,
  `amount` BIGINT NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`account_id`) REFERENCES `account` (`id`)
);

CREATE TABLE `transfers` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `from_account_id` INT NOT NULL,
  `to_account_id` INT NOT NULL,
  `amount` BIGINT NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`from_account_id`) REFERENCES `account` (`id`),
  FOREIGN KEY (`to_account_id`) REFERENCES `account` (`id`)
);

CREATE INDEX `account_index_0` ON `account` (`owner`);

CREATE INDEX `entries_index_1` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_4` ON transfers (from_account_id, to_account_id);