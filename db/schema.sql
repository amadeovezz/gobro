CREATE DATABASE gobro;
USE gobro;


CREATE TABLE IF NOT EXISTS `conn` (
    `time` DOUBLE UNSIGNED NOT NULL,
    `conn_uid` VARCHAR(255) NOT NULL ,
	`orig_ip` VARCHAR(40) NOT NULL,
	`orig_port` VARCHAR(5) NOT NULL,
	`resp_ip` VARCHAR(40) NOT NULL,
    `resp_port` VARCHAR(5) NOT NULL,
    `proto` VARCHAR(20) DEFAULT '-',
    `service` VARCHAR(80) DEFAULT '-',
    `duration` DECIMAL DEFAULT 0,
    `in_bytes` BIGINT UNSIGNED DEFAULT 0,
    `out_bytes` BIGINT UNSIGNED DEFAULT 0,
    `in_packets` BIGINT UNSIGNED DEFAULT 0,
    `out_packets` BIGINT UNSIGNED DEFAULT 0
);

CREATE TABLE IF NOT EXISTS `dns` (
	`time` DOUBLE UNSIGNED NOT NULL,
    `conn_uid` VARCHAR(255) NOT NULL,
	`orig_ip` VARCHAR(40) NOT NULL,
	`orig_port` VARCHAR(5) NOT NULL,
	`resp_ip` VARCHAR(40) NOT NULL,
    `resp_port` VARCHAR(5) NOT NULL,
    `proto` VARCHAR(20) DEFAULT '-',
    `trans_is` SMALLINT UNSIGNED NOT NULL,
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
	`time` DOUBLE UNSIGNED NOT NULL,
    `conn_uid` VARCHAR(255) NOT NULL NOT NULL,
    `orig_ip` VARCHAR(40) NOT NULL,
	`orig_port` VARCHAR(5) NOT NULL,
	`resp_ip` VARCHAR(40) NOT NULL,
    `resp_port` VARCHAR(5) NOT NULL,
	`status_code` VARCHAR(40) NOT NULL,
    `direction` VARCHAR(40) NOT NULL,
    `lan_client` VARCHAR(255) NOT NULL,
    `server_server` VARCHAR(255) NOT NULL
);
