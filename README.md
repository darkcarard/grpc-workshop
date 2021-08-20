# gRPC Workshop
#### By Carlos Ardila (Charlie)

---
# Descripción
Taller práctico en donde implementamos los diferentes tipos de API en un servicio hecho con gRPC.

Los ejemplos consisten en realizar ataques a un jefe de un juego, quien se defenderá atacándonos de vuelta.

Mediante estos "ataques" podremos ver las ventajas y desventajas de cada uno de los tipos de API que tiene gRPC: unary, server streaming, client streaming y bidirectional streaming.

## Tipos de APIs o Streaming en gRPC 
![gRPC APIs](https://i0.wp.com/blog.knoldus.com/wp-content/uploads/2020/03/f.png?w=861&ssl=1)

###### Tomado de https://blog.knoldus.com/unary-streaming-via-grpc/

---
# Configuración ambiente
- Instalar Protocol Buffers en mac:

`brew install protobuf`

- Instalar dependencias de Go:
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

# Ejecución
En la carpeta `scripts` se encuentran los comandos para ejedcutar en consola:
- `protoc.sh` para crear el código a partir de la definición en protobuf
- `run.sh` para ejecutar el servidor y cliente para cada tipo de API o streaming