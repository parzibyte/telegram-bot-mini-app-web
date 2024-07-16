package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const tokenTelegram = "Tu token va aquí"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithAllowedUpdates(bot.AllowedUpdates{"message", "edited_message", "callback_query"}),
	}

	botDeTelegram, err := bot.New(tokenTelegram, opts...)
	if err != nil {
		panic(err)
	}
	botDeTelegram.Start(ctx)
}
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	if update.Message.WebAppData != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Los datos recibidos por la webapp son %s", update.Message.WebAppData.Data),
		})
	}
	if update.Message.Text == "/registrar_notificacion" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Webapp",
			ReplyMarkup: models.ReplyKeyboardMarkup{
				Keyboard: [][]models.KeyboardButton{
					{
						{
							Text: "Por favor toca el botón para abrir la webapp",
							WebApp: &models.WebAppInfo{
								URL: "https://parzibyte.github.io/ejemplos-javascript/webapp-telegram/index.html?rutas=W3sibm9tYnJlIjoiTmV3IERvbmsgQ2l0eSIsImlkIjoxfSx7Im5vbWJyZSI6IkxvbiBMb24gUmFuY2giLCJpZCI6Mn1d&version=2",
							},
						},
					},
				},
				ResizeKeyboard:  true,
				OneTimeKeyboard: true,
			},
		})
	}
}
