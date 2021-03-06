# pdf2jpg_gin
PDFをアップロードして、スライドとして閲覧するアプリケーションのバックエンドです。

## 使用技術
<img src="https://img.shields.io/badge/Go-v1.17-578f91.svg?logo=go&style=flat">
<img src="https://img.shields.io/badge/gin-v1.7.7-578f91.svg?logo=go&style=flat">

※ pdfの変換に`popper`の`pdftoppm`を使用しています。その為、インストールが必要です。

## *popper*のインストール
### Ubuntu 20.04
```shell
sudo apt-get install poppler-utils
```
※ その他OSは各自インストールを行ってください。

## 開発の始め方
0. `go mod tidy`
1. `go run main.go`  
-> [localhost:8000](http://localhost:8000)で立ち上がります

※環境変数"PORT"にポート番号が指定されている場合は、そちらで上書き可能です。

## OpenAPIについて
本プロジェクトでは、[swaggo/swag](https://github.com/swaggo/swag#mime-types)を採用しています。
[OpenAPI comment format](https://github.com/swaggo/swag#declarative-comments-format)に従って適宜API仕様を追加してください。

APIコメントの変更後、サーバー立ち上げ前に以下のコマンドを実行することでAPI仕様書を更新できます。
```shell
swag init
```

なお、OpenAPIは[localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)で確認できます。

![Screenshot 2021-12-15 at 23-50-27 Swagger UI](https://user-images.githubusercontent.com/80180411/146210515-e96e3984-7425-4e4f-8bd9-559132f23c9d.png)


## 使用イメージ(Front: React)


https://user-images.githubusercontent.com/80180411/146194767-320d804c-2495-402c-8be4-6097bb5d4026.mp4

