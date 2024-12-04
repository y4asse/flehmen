# 必要な環境
- Docker
- Make
- Google Cloud CLI
- Atlas
  - https://atlasgo.io/getting-started/#installation

# セットアップ
以下のコマンドで全てのDockerコンテナが立ち上がります。
```
$ make up
```

以下のコマンドでデータベースのマイグレーションを実行します。
```
$ make db_migrate
```
