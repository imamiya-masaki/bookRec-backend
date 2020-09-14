# 起動
```
$ docker-compose build
$ docker-compose up -d
```

# ログの確認
```
$ docker-compose logs golang # go
$ docker-compose logs mysql # mysql
```

# 編集の反映
```
$ docker-compose down
$ docker-compose up -d
```

# ディレクトリ構成
- app: アプリケーションフォルダ
    - controller:
    - database: 全体的なデータベース周りの処理
    - middleware: 
    - models: モデル定義など
    - main.go: アプリケーション起動とルーティング
- docker: docker関連
    -  api: api関連
    - db: データベース関連
        - db_data: データベースのデータを保持
        - db_init: データベース初期化用のSQL