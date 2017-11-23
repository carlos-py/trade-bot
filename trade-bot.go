package main

import "fmt"
import "github.com/pelletier/go-toml"
import "github.com/toorop/go-bittrex"
import "time"

// Load and parse config file - trade-bot.conf
var config, err = toml.LoadFile("trade-bot.conf")

var configTree = config.Get("config").(*toml.Tree)
var base_coin = configTree.Get("base_coin").(string)
var market_coin = configTree.Get("market_coin").(string)
var api_version = configTree.Get("api_version").(string)
var base_url = configTree.Get("base_url").(string)
var api_key = configTree.Get("api_key").(string)
var api_secret = configTree.Get("api_secret").(string)
var last_sell = configTree.Get("last_sell").(float64)
var target_sell = configTree.Get("target_sell").(float64)

// Init Bittrex client
var bittrex_client = bittrex.New(api_key, api_secret)

func doEvery(d time.Duration, f func(time.Time)) {
    for x := range time.Tick(d) {
        f(x)
    }
}

func get_balance() (string) {
    base_coin_balance, _ := bittrex_client.GetBalance(base_coin)
    market_coin_balance, _ := bittrex_client.GetBalance(market_coin)
    return fmt.Sprintf("[Balances: %v %v / %v %v]", base_coin_balance.Available, base_coin, market_coin_balance.Available, market_coin)
}

func last_price(t time.Time) {
    ticker, _ := bittrex_client.GetTicker(base_coin + "-" + market_coin)
    net_diff := fmt.Sprint(((ticker.Bid - last_sell) / last_sell) * 100)
    fmt.Printf("[+] Last Bid: %v || Last Ask: %v || Last market sell: %v [ My last sell: %v Net diff: %v ]%v\n", ticker.Bid, ticker.Ask, ticker.Last, last_sell, net_diff,get_balance())
}

func main() {
    fmt.Println("****** Starting Cobranix Crypto Trader ******\n")

    // show values parsed from config
    fmt.Println("[+] The base_coin is: " + base_coin)
    fmt.Println("[+] The market_coin is: " + market_coin)
    fmt.Println("[+] The api_version is: " + api_version)
    fmt.Println("[+] The base_url is: " + base_url)
    fmt.Println("[+] The api_key is: " + api_key)
    fmt.Println("[+] The api_secret is: <secret>")
    fmt.Println("[+] The last_sell is:", last_sell)
    fmt.Println("[+] The target_sell is:", target_sell)

    // check price every 2 seconds
    doEvery(2000*time.Millisecond, last_price)
}
