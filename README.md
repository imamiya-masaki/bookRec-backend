## 起動
### .envファイル作成
ルートディレクトリで.envファイルを以下の内容で作成（変数の値は各自の環境に合わせてください）
```
DATABASE=book_db
USERNAME=docker
USERPASS=docker
ROOTPASS=root
```
※TwitterのAPIキーを設定する必要あり
### コンテナの生成と起動
```
$ docker-compose build
$ docker-compose up -d
```

## ログの確認
```
$ docker-compose logs golang # go
$ docker-compose logs mysql # mysql
```

## 編集を反映
編集を反映したいときはコンテナを再生成してください
```
$ docker-compose down
$ docker-compose up -d
```

## ディレクトリ構成
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
