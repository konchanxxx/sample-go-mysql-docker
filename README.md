# sample-go-mysql-docker

Go,MySQLでバックエンドを構築する際にDockerで完結するような開発環境を作りたいなと思い調べつつ実装してみました。

# Usage

```sh
$ docker-compose build
$ docker-compose up -d
$ docker ps
CONTAINER ID        IMAGE                        COMMAND                  CREATED             STATUS              PORTS                               NAMES
843d3bdf2feb        sample-go-mysql-docker_api   "sh ./bin/run.sh"        4 minutes ago       Up 4 minutes        0.0.0.0:8080->8080/tcp              sample-go-mysql-docker_api_1
66b6ac91f5e2        mysql:8.0                    "docker-entrypoint.s…"   26 minutes ago      Up 4 minutes        0.0.0.0:3306->3306/tcp, 33060/tcp   sample-go-mysql-docker_db_1

# productを追加
$ curl -X POST http://localhost:8080/products

# productを取得
$ curl http://localhost:8080/products/latest
{"id":1,"name":"switch","price":29980}%
```
