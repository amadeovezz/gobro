CREATE DATABASE gobro;
USE gobro;


CREATE TABLE IF NOT EXISTS `conn` (
    `ts` DOUBLE UNSIGNED NOT NULL,
    `uid` VARCHAR(255) NOT NULL ,
	`id_orig_h` VARCHAR(40) NOT NULL,
	`id_orig_p` VARCHAR(5) NOT NULL,
	`id_resp_h` VARCHAR(40) NOT NULL,
    `id_resp_p` VARCHAR(5) NOT NULL,
    `proto` VARCHAR(20) DEFAULT '-',
    `service` VARCHAR(80) DEFAULT '-',
    `duration` DECIMAL DEFAULT 0,
    `orig_bytes` BIGINT UNSIGNED DEFAULT 0,
    `resp_bytes` BIGINT UNSIGNED DEFAULT 0,
    `orig_pkts` BIGINT UNSIGNED DEFAULT 0,
    `resp_pkts` BIGINT UNSIGNED DEFAULT 0
);

CREATE TABLE IF NOT EXISTS `dns` (
	`ts` DOUBLE UNSIGNED NOT NULL,
    `uid` VARCHAR(255) NOT NULL,
	`id_orig_h` VARCHAR(40) NOT NULL,
	`id_orig_p` VARCHAR(5) NOT NULL,
	`id_resp_h` VARCHAR(40) NOT NULL,
    `id_resp_p` VARCHAR(5) NOT NULL,
    `proto` VARCHAR(20) DEFAULT '-',
    `trans_id` SMALLINT UNSIGNED NOT NULL,
    `query` TEXT NOT NULL,
    `qclass` TINYINT(1) UNSIGNED NOT NULL,
    `qclass_name` VARCHAR(255) NOT NULL,
    `qtype` TINYINT(1) UNSIGNED NOT NULL,
    `qtype_name` VARCHAR(255) NOT NULL,
	`rcode` TINYINT(1) UNSIGNED NOT NULL,
    `rcode_name` VARCHAR(255) NOT NULL,
    `AA` VARCHAR(2) NOT NULL,
    `TC` VARCHAR(2) NOT NULL,
    `RD` VARCHAR(2) NOT NULL,
    `RA` VARCHAR(2) NOT NULL,
    `Z` TINYINT(1) UNSIGNED NOT NULL,
    `answers` TEXT NOT NULL,
	`TTLs` TEXT NOT NULL,
    `rejected` VARCHAR(2) NOT NULL
);


CREATE TABLE IF NOT EXISTS `ssh` (
	`ts` DOUBLE UNSIGNED NOT NULL,
    `uid` VARCHAR(255) NOT NULL NOT NULL,
    `id_orig_h` VARCHAR(40) NOT NULL,
	`id_orip_p` VARCHAR(5) NOT NULL,
	`id_resp_h` VARCHAR(40) NOT NULL,
    `id_resp_p` VARCHAR(5) NOT NULL,
	`status` VARCHAR(40) NOT NULL,
    `direction` VARCHAR(40) NOT NULL,
    `client` VARCHAR(255) NOT NULL,
    `server` VARCHAR(255) NOT NULL
);



