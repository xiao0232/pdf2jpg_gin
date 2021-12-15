# pdf2jpg_gin
PDFをアップロードして、スライドとして閲覧するアプリケーションのバックエンドです。

## 使用技術
<img src="https://img.shields.io/badge/Go-v1.17-578f91.svg?logo=go&style=flat">
<img src="https://img.shields.io/badge/gin-v1.7.7-578f91.svg?logo=go&style=flat">

※ pdfの変換方法として`popper`の`pdftoppm`を使用しています。その為、インストールが必要です。

## *popper*のインストール
### Ubuntu 20.04
```shell
sudo apt-get install poppler-utils
```
