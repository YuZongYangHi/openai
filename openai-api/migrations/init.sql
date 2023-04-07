create database openai DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use openai;

CREATE TABLE `openai_channel`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(255) NOT NULL,
    `created_at` datetime(6) NOT NULL,
    `updated_at` datetime(6) NOT NULL
);

CREATE TABLE `openai_message`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `channel_id` bigint(20) NOT NULL,
    `content` longtext NOT NULL,
    `dialog_type` int(10) NOT NULL,
    `content_type` int(10) DEFAULT 0,
    `created_at` datetime(6) NOT NULL
);