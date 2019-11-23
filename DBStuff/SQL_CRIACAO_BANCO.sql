-- MySQL Script generated by MySQL Workbench
-- Sat Nov 23 02:36:54 2019
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mydb` DEFAULT CHARACTER SET utf8 ;
USE `mydb` ;

-- -----------------------------------------------------
-- Table `mydb`.`gerente`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`gerente` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(45) NOT NULL,
  `email` VARCHAR(100) NOT NULL,
  `senha` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`restaurante`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`restaurante` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(45) NOT NULL,
  `endereco` VARCHAR(45) NOT NULL,
  `telefone` VARCHAR(45) NOT NULL,
  `horario_atendimento` VARCHAR(45) NOT NULL,
  `gerente_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_restaurante_gerente_idx` (`gerente_id` ASC),
  CONSTRAINT `fk_restaurante_gerente`
    FOREIGN KEY (`gerente_id`)
    REFERENCES `mydb`.`gerente` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`cardapio`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`cardapio` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cor` VARCHAR(30) NOT NULL,
  `logo` BLOB NOT NULL,
  `restaurante_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_cardapio_restaurante1_idx` (`restaurante_id` ASC) ,
  CONSTRAINT `fk_cardapio_restaurante1`
    FOREIGN KEY (`restaurante_id`)
    REFERENCES `mydb`.`restaurante` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`item` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(45) NOT NULL,
  `descricao` VARCHAR(150) NULL,
  `preco` DOUBLE NOT NULL,
  `cardapio_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_item_cardapio1_idx` (`cardapio_id` ASC) ,
  CONSTRAINT `fk_item_cardapio1`
    FOREIGN KEY (`cardapio_id`)
    REFERENCES `mydb`.`cardapio` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;