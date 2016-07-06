# google-apps-utils

Google Apps のユーティリティコマンドです

# Installation

google-apps-utils をインストールします

```
$ go get github.com/yagihiro/google-apps-utils
$ go install github.com/yagihiro/google-apps-utils
$ export PATH=$PATH:$GOPATH/bin
```

管理 API を有効にします

https://support.google.com/a/answer/60757

以下のドキュメントの "Step 1: Turn on the Directory API" を参考にして client_secret.json を入手します

https://developers.google.com/admin-sdk/directory/v1/quickstart/go

以下のパスに client_secret.json を置きます

```
$HOME/.google-apps-utils/client_secret.json
```

実行します
```
$ google-apps-utils --help
```


# Run

* ユーザー一覧を表示

```
$ google-apps-utils list
```

* ユーザーを追加

```
$ google-apps-utils create -g givenname -f familyname -e primaryemail
```

# 開発者向け

依存しているライブラリをインストールします

```
$ export GOPATH=xxx
$ go get golang.org/x/net/context
$ go get golang.org/x/oauth2
$ go get golang.org/x/oauth2/google
$ go get google.golang.org/api/admin/directory/v1
$ go get github.com/urfave/cli
```

# Reference

* https://developers.google.com/admin-sdk/directory/
* https://godoc.org/google.golang.org/api/admin/directory/v1

# License

MIT License
