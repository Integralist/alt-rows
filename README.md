# alt-rows

Simple program that streams data piped into stdin and colors each line with alternative colors.

## Examples

```
seq 1 5 | go run main.go
```

![](./alt-rows.png)

```
seq 1 5 | go run main.go -pattern '(?:[5-9]|10)'
```

![](./alt-rows-pattern.png)

## Install

```
go build
sudo cp ./alt-rows /usr/local/bin/
```
