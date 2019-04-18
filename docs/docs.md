
```$xslt
go build -ldflags "-X main.version=development" -o iChat cmd/iChat/iChat.go
```

```$xslt
./iChat --config-file configs/iChat.yml start
```