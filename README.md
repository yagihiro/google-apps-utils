# google-apps-utils

Google Apps のユーティリティコマンドです

# Installation

```
$ go get github.com/yagihiro/google-apps-utils
$ go install github.com/yagihiro/google-apps-utils
$ export PATH=$PATH:$GOPATH/bin
$ google-apps-utils --help
```

# Pre requirements

* 管理 API を有効にします
  * https://support.google.com/a/answer/60757

* 以下のドキュメントの "Step 1: Turn on the Directory API" を参考にして client_secret.json を入手してカレントディレクトリに置きます
  * https://developers.google.com/admin-sdk/directory/v1/quickstart/go

* 以下のコマンドをたたきます

```
$ export GOPATH=xxx
$ go get golang.org/x/net/context
$ go get golang.org/x/oauth2
$ go get golang.org/x/oauth2/google
$ go get google.golang.org/api/admin/directory/v1
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

# Reference

* https://godoc.org/google.golang.org/api/admin/directory/v1

# License

MIT License
