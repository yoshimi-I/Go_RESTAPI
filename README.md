# Go言語チュートリアル
- 文法の解説と簡単なAPIを設計するところまでをまとめます。
# 目次
# まず最初にやること
1. docker-composeからGoのイメージの取得
    - 簡単にいうとbuildでイメージの作成を行い,up -dでコンテナなの中に入ってください
```
docker compose build 
docker cmpose up -d　
```

```
go mod init github.com/yoshimi-I/Go_RESTAPI
```
- といった感じでリポジトリのurlのhttps:以降をinitの後に続けたものをターミナルに打ち込んでください。
- そうすることでgo.modというバージョンを管理するファイルが作られます。
  - これにパッケージをインポートいていく感じですね、Reactでいうpackage.jsonみたいなやつ
- そういうルールだと思うようにしました。