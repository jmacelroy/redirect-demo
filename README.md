# Divert demo

Deploy the demo application in Okteto Staging by pressing the following button:

[![Develop on Okteto](https://okteto.com/develop-okteto.svg)](https://staging.okteto.dev/deploy?repository=https://github.com/jmacelroy/redirect-demo)

Once the application is running, if you access the application endpoint (for example, client-api-cindy.staging.okteto.net) the content would be:

```
{
	"item": {
		"name": "LEDx",
		"type": "Medical Equipment",
		"weight": 0.23,
		"grid_size": "2x1",
		"loot_experience": 50
	},
	"flea": {}
}
```

Now, run `okteto up`:

```
 ✓  Images successfully pulled
 ✓  Files synchronized
    Name:      cindy-loot-data
    URL:       cindy-client-api-cindy.staging.okteto.net

Welcome to your development container. Happy coding!
cindy:cindy-loot-data app>
```

Compile the application:

```
cindy:cindy-loot-data app> make build
go build -o load-data cmd/loot-data/main.go
```

Run the application:

```
cindy:cindy-loot-data app> make start
./load-data
loot data server listening on :8081
```

Now you can access the original application at client-api-cindy.staging.okteto.net and the dev version of your application at  cindy-client-api-cindy.staging.okteto.net:

```
{
	"item": {
		"name": "LEDx",
		"type": "Medical Equipment",
		"weight": 0.23,
		"grid_size": "2x1",
		"loot_experience": 50
	},
	"flea": {
		"rarity": "legendary",
		"average_price_24h": 710000.25,
		"average_price_7d": 805000.8
	}
}
```

Happy coding!