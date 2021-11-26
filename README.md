# learn-ddd
DDDの学習用

# 立ち上げ
## docker環境構築
初回はgooseのインストールに少し時間がかかる。

```bash
docker-compose up -d
```

## マイグレーション
毎回usernameなどを指定する必要がある？ => 要調査。

```bash
docker-compose exec app ash
goose -dir migration mysql "root:root@tcp(db:3306)/ddd?parseTime=true"
```

### マイグレーションファイルを作成
```bash
docker-compose exec app ash
goose -dir migration create create_***_table sql
```