CREATE DATABASE IF NOT EXISTS todo;
USE todo;

CREATE TABLE `todo` (
    `id` INT(255) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `user` (
    `id` INT(255) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);


INSERT INTO `user` (`name`,`password`,`created_at`,`updated_at`) VALUES ("テスト","nbioewjffklw",NOW(),NOW());
INSERT INTO `todo` (`name`,`created_at`,`updated_at`) VALUES ("テスト",NOW(),NOW());