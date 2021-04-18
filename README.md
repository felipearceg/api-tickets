# api-tickets

Para probar en local:

Instalar golang

![alt text](https://github.com/felipearceg/api-tickets/blob/master/descargargo.png)

![alt text](https://github.com/felipearceg/api-tickets/blob/master/descomprimirgo.png)

![alt text](https://github.com/felipearceg/api-tickets/blob/master/editaretcprofile.png)

![alt text](https://github.com/felipearceg/api-tickets/blob/master/variabledeentorno.png)

![alt text](https://github.com/felipearceg/api-tickets/blob/master/reiniciar.png)

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

Luego ejecutar la aplicación creando un ejecutable con go build:

go build .

./api-tickets