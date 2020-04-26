# !/bin/bash

until mysqladmin ping -h db --silent; do
  echo 'waiting for connection mysql...'
  sleep 2
done

echo 'Server is staring!'
exec go run main.go
