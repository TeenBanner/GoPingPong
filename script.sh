#!/bin/bash

# URL de destino
url="http://localhost:8080/ping"

# Número de solicitudes a enviar
n=1000000000000000

# Bucle for para enviar las solicitudes
for ((i=1; i<=$n; i++))
do
  # Imprime el número de solicitud actual
  echo "Enviando solicitud $i a $url"

  # Utiliza el comando curl para enviar la solicitud
  # Puedes personalizar los parámetros de curl según tus necesidades
  curl -X GET $url

  # Pausa de 1 segundo entre solicitudes (opcional)
  sleep 1
done
