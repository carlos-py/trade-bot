# trade-bot
Crypto trading bot written in Go

## Dependencies

Third-party packages: go-bittrex and go-toml
```
go get github.com/pelletier/go-toml
go get github.com/toorop/go-bittrex

go install github.com/pelletier/go-toml
go install github.com/toorop/go-bittrex
```
## Configuration

See the config.toml file for configuration options

## Run or Build from source

```
go run trade-bot.go
```

```
go build trade-bot.go && ./trade-bot
```
