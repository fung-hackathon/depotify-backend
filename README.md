# depotify-backend

Hackathon for FUN 2022 バックエンド用の repository です。

## サーバの起動方法

### 環境変数

- ARRIVAL_REDIRECT_URL<br>
  /arrive に GET リクエストした際にリダイレクトする先を指定<br>
  例) `https://example.com/arrive`
- FRONT_ORIGIN<br>
  フロントエンドの Origin を設定。CORS設定用。<br>
  例) `https://example.com`
- GOOGLE_APPLICATION_CREDENTIALS<br>
  Firebase の認証用ファイルへのPath<br>
  例) `path/to/credentials.json`
- MODE<br>
  ログのモード設定。`production`にするとログが抑制される。<br>
  例) `production`
- YOLP_APPID<br>
  [Yahoo! Open Local Platform](https://developer.yahoo.co.jp/webapi/map/) にアクセスする際に必要な APP ID を入れる。<br>
  例) なし
### ローカル

```
go run ./cmd/main.go
```

### Render
#### Build Command
```
cd cmd && go build -tags netgo -ldflags '-s -w' -o app
```

####  Start Command
```
./cmd/app
```


## For Developers

see [this](fordev.md)

## Links

[funhackathon2022 repository](https://github.com/fung-hackathon/funhackathon2022)
