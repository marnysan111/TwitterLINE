package auth

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func LineAuth() (*linebot.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("LINE接続エラー:", err)
	}
	Sercret := os.Getenv("CHANNEL_SERCRET")
	Token := os.Getenv("CHANNEL_TOKEN")
	bot, err := linebot.New(
		Sercret, Token,
	)
	return bot, nil
}
