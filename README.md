# api-tickets

Para probar en local:

Instalar mysql-server

Ingresar a la consola de MySQL y ejecutar las siguientes instrucciones:

create database apitickets;

use apitickets;

create table tiquetes (
	id int not null auto_increment,
	usuario int not null,
	fechacreado varchar(255),
	fechaactualizado varchar(255),
	estado varchar(255),
	primary key (id)
);
