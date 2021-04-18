# Para probar en local

## Instalar golang

![Image 1](https://github.com/felipearceg/api-tickets/blob/master/images/descargargo.png)

![Image 2](https://github.com/felipearceg/api-tickets/blob/master/images/descomprimirgo.png)

![Image 3](https://github.com/felipearceg/api-tickets/blob/master/images/editaretcprofile.png)

![Image 4](https://github.com/felipearceg/api-tickets/blob/master/images/variabledeentorno.png)

![Image 5](https://github.com/felipearceg/api-tickets/blob/master/images/reiniciar.png)

## Instalarmysql-server

![Image 6](https://github.com/felipearceg/api-tickets/blob/master/images/update.png)

![Image 7](https://github.com/felipearceg/api-tickets/blob/master/images/instalarmysql.png)

## Ingresara la consola de MySQL y ejecutar las siguientes instrucciones:

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

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';

sudo service mysql stop

sudo service mysql start

![Image 8](https://github.com/felipearceg/api-tickets/blob/master/images/ingresarmysql.png)

![Image 9](https://github.com/felipearceg/api-tickets/blob/master/images/createdatabase.png)

## Descargar el repositorio

![Image 10](https://github.com/felipearceg/api-tickets/blob/master/images/descargarzip.png)

![Image 11](https://github.com/felipearceg/api-tickets/blob/master/images/unzip.png)


## Luego ejecutar la aplicación creando un ejecutable con go build

go build .

./api-tickets