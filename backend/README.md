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
マイグレーションとは
　　やせが作ったやつがあやのパソコンにも入る
    gitみたいな感じやね
```
$ make db_migrate
```

以下のコマンドでデータベースのシーディングを実行します。
シーディングとは
  ダミーデータを入れること
```
$ make db_seed
```
