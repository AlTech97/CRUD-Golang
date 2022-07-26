/*CREAZIONE DEL DATABASE*/

Drop DATABASE IF EXISTS comunity;
CREATE DATABASE comunity;
USE comunity;

/*CREAZIONE TABELLE*/
CREATE TABLE persona (
	c_f CHAR(16) Not NULL,
	nome VARCHAR(20),
	cognome VARCHAR(20),
	PRIMARY KEY (c_f)
	);

CREATE TABLE telefono(
	numero VARCHAR(15) NOT NULL,
	intestatario CHAR(16) NOT NULL,
	PRIMARY KEY (numero),
	FOREIGN KEY (intestatario) REFERENCES persona(c_f)
	ON UPDATE CASCADE
	ON DELETE CASCADE
	);
	
/* POPOLAZIONE DATABASE*/
INSERT INTO persona
VALUES ('SGSMNL82R21D643D', 'Manuele', 'Saggese');

INSERT INTO persona
VALUES ('MLNDLA94D54B590W', 'Dalia', 'Milani');

INSERT INTO telefono
VALUES ('3321234530', 'SGSMNL82R21D643D');

INSERT INTO telefono
VALUES ('3351049390', 'SGSMNL82R21D643D');

INSERT INTO telefono
VALUES ('3545787510', 'MLNDLA94D54B590W');