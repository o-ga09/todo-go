# TODO アプリ with Go

<img width="1000px" height="300px" src="https://github.com/o-ga09/tic-tac-toe-go/assets/54522966/9fc52696-3879-4f01-a818-98b86ddf774f">

## Features

ユーザーとユーザーに紐づくタスクの管理する API

## How to Use

## ユーザー処理

- ユーザーの一覧

```bash
curl http://localhost:8080/api/v1/users
```

- 各ユーザーの取得

```bash
curl http://localhost:8080/api/v1/user/1
```

- ユーザーの作成

```bash
curl -X POST -H 'Content-type: application/json' -d '{"name":"test","password":"dmwmdodk"}' http://localhost:8080/api/v1/user
```

- ユーザーの更新

```bash
curl -X PATCH -H 'Content-type: application/json' -d '{"name":"hogehoge","password":""}' http://localhost:8080/api/v1/user/5
```

- ユーザーの削除

```bash
curl -X DELETE http://localhost:8080/api/v1/users/5
```

## タスクの取得

- タスク一覧

```bash
curl http://localhost:8080/api/v1/tasks
```

- 各タスク

```bash
curl http://localhost:8080/api/v1/task/1
```

- タスクの作成

```bash
curl -X POST -H 'Content-type: application/json' -d '{"name":"sigoto"}' http://localhost:8080/api/v1/task
```

- タスクの更新

```bash
url -X PATCH -H 'Content-type: application/json' -d '{"name":"working"}' http://localhost:8080/api/v1/task/5
```

- タスクの削除

```bash
curl -X DELETE http://localhost:8080/api/v1/task/5
```

## How to Build

docker イメージにビルドします

```bash
make build
```

## How to Run

ローカルで go アプリケーションを起動します

```bash
make up
make run
```

## Post Condition

ローカルでの開発環境を片付けます

```bash
make down
```

docker コンテナにマウントしたデータも削除する場合

```bash
make rm
```
