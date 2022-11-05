CREATE TABLE Usuario (
  	id MEDIUMINT NOT NULL AUTO_INCREMENT,
  	nome varchar(255) NOT NULL,
    senha varchar(255) NOT NULL,
  	email varchar(255) NOT NULL UNIQUE,
  
  	PRIMARY KEY (id)
);

