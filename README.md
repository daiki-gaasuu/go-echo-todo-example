# Go ＋ Vue3 TODO サンプル

## バックエンド

- Go1.20.1
- Gorm
- Echo
- OpenAPI

### ディレクトリ構成

番号は Clean Architecture に準拠した場合の外側からの順番

- server
  - cmd
    - main.go
  - infrastructure→1
    - router.go
    - sql_handler.go→ 外部パッケージ(gorm)使用のため一番外に定義
  - interfaces→2
    - controllers
    - database→ 実際にデータをやり取りする処理
      - sql_handler.go→SqlHandler の振る舞いを定義。database 層から実際に呼び出すのはこの振る舞い。
      - todo_repository.go
  - usecase→3
    - todo_interactor.go→interfaces/controllers への Gateway の役割
    - todo_repository.go→interfaces/database からの Input Port の役割
  - domain→4
    - todo.go
- compose.yml

## フロントエンド

- Vite4.1.0
- Vuetify3
- Vue3
- Node19.6.1
- yarn4.0.0-rc.39

## OpenAPI

- バックエンド
  - https://github.com/deepmap/oapi-codegen
  - プロジェクトルートにて`make gen-backend`
- フロントエンド
  - プロジェクトルートにて`make gen-frontend`
  - https://github.com/OpenAPITools/openapi-generator

## 起動方法

1. 上記の OpenAPI の生成
2. `cd server`
3. `go run main.go`
4. client ルートにて`yarn`
5. yarn run dev
