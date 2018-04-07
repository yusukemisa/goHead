## goHead
goによるheadコマンド実装です。

#### 以下仕様を満たす

- [x] 引数で渡されたファイルの先頭の最大N行をそのまま出力する
- [x] Nはデフォルトでは10
- [x] オプション -n でNを指定できる

#### 追加機能

- [x] オプション -c で出力するバイト数を指定可能
- [x] 複数ファイル指定可能にする。その場合`==> {ファイル名} <==`でファイル内容を区切る

#### 実行例
```
$ myhead -n=3 hoge.txt
A
B
C
```

## installation

```
$ go get github.com/yusukemisa/goHead
```

## test(可能であれば)
* 主要な処理をmainから追い出してテストコード書けるようにしたい
* headコマンドの結果と一致するテストを書く
