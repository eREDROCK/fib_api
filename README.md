#  概要
fib_apiは指定されたn番目のフィボナッチ数を返すREST APIが動作するWebサーバーです．  
以下のようなリクエストを送信することで指定の番号のフィボナッチ数がjson形式で送られてきます．  
https://fib-server.onrender.com/fib?n=[得たいフィボナッチ数の番号]  

#  ソースコードの構成
```
.
├── api
│   ├── Dockerfile
│   ├── Dockerfile.dev
│   ├── controllers
│   │   └── GetFib.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── router
│   │   └── router.go
│   ├── test
│   │   └── fib_test.go
│   └── tmp
│       ├── build-errors.log
│       └── main
└── docker-compose.yml
```
| ファイル           | 説明                                                             | 
| ------------------ | ---------------------------------------------------------------- | 
| api                | REST APIに関するすべてのコードを格納したフォルダ                 | 
| docker-compose.yml | 開発環境でDockerコンテナ上にGolangの環境を構成するためのファイル | 
| Dockerfile         | 本番環境用のDockerfile                                           | 
| Dockerfile.dev     | 開発環境用のDockerfile                                           | 
| main.go            | サーバーを起動するファイル                                       | 
| controllers        | リクエストが来た際に実行される関数をまとめたフォルダ             | 
| router             | ルーティングの設定を行うフォルダ                                 | 
| test               | ユニットテストを行うフォルダ                                     | 
