# api-tickets
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
