# google-apps-utils

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/yagihiro/google-apps-utils/blob/master/LICENSE

Google Apps for Work の管理者用ユーティリティコマンドです

# Installation

google-apps-utils をインストールします

```
$ export GOPATH=/path/to/your/gopath
$ export PATH=$GOPATH/bin:$PATH
$ go get -u github.com/yagihiro/google-apps-utils
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

### ユーザー一覧を表示

ドメインに所属する全ユーザーリストをターミナルに表示します。
```
$ google-apps-utils list
```

### ユーザーを追加

ドメインにユーザーを追加します。ターミナルに初期パスワードが表示されるので、primaryemail と合わせてユーザーにお伝えください。
```
$ google-apps-utils create -g givenname -f familyname -e primaryemail
```

### グループ一覧を表示

ドメインに所属する全グループリストをターミナルに表示します。
```
$ google-apps-utils grouplist
```

### グループを追加

ドメインにグループを追加します。
```
$ google-apps-utils groupcreate -e email -d description -n name
```

### グループのメンバーを表示

ドメインに所属するグループのメンバーを一覧表示します
```
$ google-apps-utils groupmemberlist -k groupKey
```

### グループメンバーを追加

ドメインに所属するグループにメンバーを追加します。
```
$ google-apps-utils groupmembercreate -k groupKey -e email -r role
```


### トークンのリセット

ローカルにキャッシュ済みのトークンをリセットします。
```
$ google-apps-utils reset
```

# Files

* $HOME/.google-apps-utils/client_secret.json
  * OAuth Client Secret です
  * 詳細は [こちら](https://developers.google.com/admin-sdk/directory/v1/quickstart/go) を参照してください
* $HOME/.google-apps-utils/token.json
  * OAuth token のキャッシュファイルです

# 開発者向け

依存しているライブラリをインストールします

```
$ export GOPATH=xxx
$ go get golang.org/x/net/context
$ go get golang.org/x/oauth2
$ go get golang.org/x/oauth2/google
$ go get google.golang.org/api/admin/directory/v1
$ go get google.golang.org/api/admin/reports/v1
$ go get github.com/urfave/cli
```

# Reference

* https://developers.google.com/admin-sdk/directory/
* https://godoc.org/google.golang.org/api/admin/directory/v1

# License

MIT License

# Author

[Hiroki Yagita](https://github.com/yagihiro)
