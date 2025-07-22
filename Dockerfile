# Usa una imagen base de Go directamente
FROM golang:1.23.2-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum para descargar las dependencias
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copia todo el código fuente de la aplicación
COPY cmd ./cmd
COPY internal ./internal
COPY database ./database
COPY models ./models
# COPY services ./services
# COPY cert/ ./cert/

# Construye la aplicación Go
# NOTA: Sin CGO_ENABLED=0, la compilación usará el valor por defecto
# Si tu aplicación o sus dependencias requieren bibliotecas C,
# y estas no están presentes en la imagen alpine, el binario podría fallar al ejecutarse.
RUN GOOS=linux go build -o products-api ./cmd/api

# # Instala CA certificates para HTTPS
# RUN apk --no-cache add ca-certificates

# Expone el puerto en el que la aplicación Go escuchará
# Creo que no es necesario, con 9090 en el compose sube.
EXPOSE 9090  

# Comando para iniciar la aplicación cuando el contenedor se inicie
CMD ["./products-api"]