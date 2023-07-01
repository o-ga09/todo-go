DROP TABLE IF EXISTS `todo`;

CREATE TABLE `todo` (
    `id` INT(255) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `createdAt` TIMESTAMP NOT NULL,
    `updatedAt` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `user` (
    `id` INT(255) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `createdAt` TIMESTAMP NOT NULL,
    `updatedAt` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
);