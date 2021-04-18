# api-tickets

Para probar en local:

Instalar golang

![Image 1](https://github.com/felipearceg/api-tickets/blob/master/descargargo.png)

![Image 2](https://github.com/felipearceg/api-tickets/blob/master/descomprimirgo.png)

![Image 3](https://github.com/felipearceg/api-tickets/blob/master/editaretcprofile.png)

![Image 4](https://github.com/felipearceg/api-tickets/blob/master/variabledeentorno.png)

![Image 5](https://github.com/felipearceg/api-tickets/blob/master/reiniciar.png)

Instalarmysql-server

Ingresara la consola de MySQL y ejecutar las siguientes instrucciones:

create dtabase apitickets;

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