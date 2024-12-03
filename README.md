# これは何？

```svgファイル``` を ```pngファイル``` に変換するツールです。

## インストール

```sh
go install github.com/devlights/svg2png/cmd/svg2png@latest 
```

ライブラリとして利用する場合は

```sh
go get github.com/devlights/svg2png@latest
```

## 使い方

```sh
$ svg2png.exe -in SVGファイルパス -out PNGファイルパス
```

### HELP

```sh
$ svg2png.exe -help
Usage of svg2png.exe:
  -debug
        debug mode
  -in string
        input file
  -opacity float
        opacity (default 1)
  -out string
        output file (default "out.png")
  -size int
        size in pixels (default 1000)
```

## ビルド時の前提条件

- [Go](https://go.dev/) 1.23 以降
- [Task](https://taskfile.dev/#/)

## ビルド

[Task](https://taskfile.dev/#/) を使っています。ビルドファイルは [Taskfile.yml](./Taskfile.yml) です。

```sh
$ task build
```

## 補足

本ツールは [github.com/typomedia/rasterize](https://github.com/typomedia/rasterize/tree/master) の実装を参考にしました。

素晴らしいアイデアを共有してくれてありがとうございます。
