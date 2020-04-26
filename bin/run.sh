# !/bin/bash

until mysqladmin ping -h db --silent; do
  echo 'waiting for connection mysql...'
  sleep 2
done

echo 'Database is launched!!'
exec go run main.go
