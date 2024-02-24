﻿# ngin-link-server
## 動作確認
1. .envを埋めます
2. dockerでgoの実行環境を用意します
```bash
docker compose up -d
```
3. 実行環境に入ります
```bash
docker exec -it ngin_link_server /bin/bash
```
4. デバッグします
```bash
go run cmd/main.go
```
