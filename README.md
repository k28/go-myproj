# GOでCLIツールを作成する際のテンプレート

CLIツールを作成する際には以下のファイル構成で作成する

```
├── LICENSE
├── Makefile
├── README.md
├── bin
├── cmd
│   └── myproj
│       ├── main.go
│       └── main_test.go
├── go.mod
├── myproj.go
├── myproj_test.go
└── util.go
```

# 知っている人は知っているけど、初心者がつまずく点

- 他のパッケージから呼ぶ関数は大文字から始める(mainパッケージから呼ぶ関数など)
- リポジトリの名前とツールの名称は同じにする(go-は付けてもOK)

