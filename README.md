## Tech Stack

- Go
- SQLite
- Gorm
- JWT
- Gin
- Testify

## Consideraciones

Para este challenge decidí utilizar sqlite como base de datos en memoria, para la autenticación de usuario se utilizo
JWT, el token se podría persistir en cache pero por cuestiones de tiempo no se implementó.

El código se separo en 3 capas (**controller**, **service**, **repository**) mas una capa de infraestructura

Los tests unitarios no están completos por cuestiones de tiempo, pero se dejo un ejemplo de test realizado con **testify**
en el controller de propiedades. Se realizaría de la misma manera en las demas capas ya que todas cuentan con interfaces
para poder mockear las funciones.

## Ejecutar el código

Dejo 2 opciones, la primera y la que recomiendo es utilizar el dockerfile para crear la imagen y correrla en un
contenedor.

### Crear imagen y ejecutar con docker

Creamos la imagen

```shell
docker build -t lahaus-challenge .
```

Corremos la imagen como contenedor

```shell
docker run -p 8080:8080 lahaus-challenge
```

### Comandos para ejecutar el server y poder realizar pruebas localmente

Obtenemos las dependencias

```shell
go get -u -d -v
```

Compilamos el código

```shell
go build .
```

Ejecutamos

```shell
go run main.go
```

# Como realizar las pruebas

Ejecutar los siguientes comandos o importar en postman

## Propiedades

### Creado de propiedades

```shell
    curl --location --request POST 'http://localhost:8080/v1/properties' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "title": "Apartamento cerca a la estación de 2",
                "location": {
                    "longitude": -99.0665887,
                    "latitude": 19.6371593
                },
                "pricing": {
                    "salePrice": 3000000
                },
                "propertyType": "HOUSE",
                "bedrooms": 3,
                "bathrooms": 2,
                "parkingSpots": 1,
                "area": 60,
                "photos": [
                    "https://cdn.pixabay.com/photo/2014/08/11/21/39/wall-416060_960_720.jpg",
                    "https://cdn.pixabay.com/photo/2016/09/22/11/55/kitchen-1687121_960_720.jpg"
                ]
            }'
```

### Actualización de propiedad

```shell


        curl --location --request PUT 'http://localhost:8080/v1/properties/1' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "title": "Apartamento cerca a la estación de transmilenio",
            "description": "TESTTTTT 4",
            "location": {
                "longitude": -74.0665887,
                "latitude": 4.6371593
            },
            "pricing": {
                "salePrice": 450000000,
                "administrativeFee": 250000
            },
            "propertyType": "HOUSE",
            "bedrooms": 3,
            "bathrooms": 2,
            "parkingSpots": 1,
            "area": 60,
            "photos": [
                "https://cdn.pixabay.com/photo/2014/08/11/21/39/wall-416060_960_720.jpg",
                "https://cdn.pixabay.com/photo/2016/09/22/11/55/kitchen-1687121_960_720.jpg"
            ],
            "status": "ACTIVE"
        }'

```

### Busqueda de propiedades

Request

```shell
    curl --location --request GET 'http://localhost:8080/v1/properties?status=ALL&pageNumber=1&pageSize=20'
```

## Usuarios

### Crear usuario

Request

```shell
    curl --location --request POST 'http://localhost:8080/v1/users/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "email": "code-challenge-lahaus@test.lh"
    }'
```

### Login

Request

```json
{
  "email": "code-challenge-lahaus@test.lh",
  "password": "code-challenge-lahaus@test.lh"
}
```

El token que obtenemos en el response lo utilizaremos para authenticarnos en los siguientes requests

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjFlNDljMmE5LTkyMzktNDU2Zi04MGNmLTAwYjgxZTUyMmY4NSIsImV4cCI6MTYzNTgwOTQ4OCwidXNlcl9pZCI6ImNvZGUtY2hhbGxlbmdlLWxhaGF1c0B0ZXN0LmxoIiwidXNlcl9uYW1lIjoiY29kZS1jaGFsbGVuZ2UtbGFoYXVzQHRlc3QubGgifQ.JOH80m4KEjwY_1oA97njxdN4zSwrHWxl0EgFimLgULc"
}
```

### Set Favorites

_En el header utilizar el token generado en el Login_

#### Request

```shell
        curl --location --request POST 'http://localhost:8080/v1/users/me/favorites' \
        --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjdjNGMxZGVlLWE3MTctNDQ5ZS1iMTkxLWVhNDM4YTUyZTllNCIsImV4cCI6MTYzNTc5NTY2MiwidXNlcl9pZCI6ImNvZGUtY2hhbGxlbmdlLWxhaGF1c0B0ZXN0LmxoIiwidXNlcl9uYW1lIjoiY29kZS1jaGFsbGVuZ2UtbGFoYXVzQHRlc3QubGgifQ.q19iSFoMRQBP0_22usx6ndkTazT6AcLgIhhbICT6VBc' \
        --header 'Content-Type: application/json' \
        --data-raw '{
             "propertyId": 1
        }'
```

### Get favorites

_En el header utilizar el token generado en el Login_

### Request

```shell
    curl --location --request GET 'http://localhost:8080/v1/users/me/favorites' \
    --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImZiZGVjMWVjLTVmZjAtNDZmOS04OTZhLTRjNGNkMTA4NjM1ZiIsImV4cCI6MTYzNTc5Mzc4NiwidXNlcl9pZCI6ImNvZGUtY2hhbGxlbmdlLWxhaGF1c0B0ZXN0LmxoIiwidXNlcl9uYW1lIjoiY29kZS1jaGFsbGVuZ2UtbGFoYXVzQHRlc3QubGgifQ.ZClw3arpa6f1qHndqCfTfRZBqzBAXubK2MrLku0xHZE' \
    --data-raw ''
```