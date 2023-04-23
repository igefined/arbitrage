# Arbitrage bot

## Description

This application analyzes daily tickers for major cryptocurrencies and cryptocurrency exchanges.

## Installation

git clone https://github.com/igdotog/arbitrage.git

## Usage

You must have docker installed. To run the program, change to the project directory and run the following command:
`make build`

## Requirements

Environment variables:

| Variable             | Description                                                |
|----------------------|------------------------------------------------------------|
| `APP_NAME`           | Name of app                                                |
| `PORT`               | Usage http port                                            |
| `DB_URL`             | Database url                                               |
| `NINJA_API_KEY`      | Api Key to Ninja application interface                     |
| `CLIENT_API_KEY`     | Client Api Key which can send requests to your http server |
| `TELEGRAM_API_TOKEN` | Telegram Bot ApiKey                                        |

## How it works

Once an hour, the application updates information about current spreads.
Information about cryptocurrency pairs is requested from Ninja.
After that, the application runs through all the supported exchanges and compares the prices of each other.
If the price difference is more than 2 percent, the information is stored in the database and sent to the telegram bot.
Information about saved spreads can be found at the endpoint `/bundles`

Supported exchanges:
 - Binance
 - ByBit
 - KuCoin
 - Huobi
 - Mexc
 - Okx
