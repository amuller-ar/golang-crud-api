# Code Challenge Alan Müller

Bienvenidos a mi code challenge, mas abajo explicaré como ejecutar y probar la api

## Ejecutar el código

### Crear imagen y ejecutar con docker

Creamos la imagen

    docker build . -t lahaus-challenge 

Corremos la imagen como contenedor

    docker run -p 8080:8080 lahaus-challenge

### Comandos para ejecutar el server y poder realizar pruebas localmente

Obtenemos las dependencias

    go get -u -d -v

Compilamos el código

    go build .

Ejecutamos

    go run main.go

### Endpoints Disponibles

Creado de propiedades

    POST /v1/properties

Actualización

    PUT v1/properties/{id}
