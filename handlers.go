package main

import (
	//	"fmt"

	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "Are you bored at home from CB? Watched too much Big Bang Theory but no one to play Rock Paper Scissors Lizard Spock with you? \nHere's the perfect bot for you!"
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle the /play command here
func (a *application) playHandler(m *tbot.Message) {
	buttons := makeButtons()
	a.client.SendMessage(m.Chat.ID, "Pick an option:", tbot.OptInlineKeyboardMarkup(buttons))
}

// Handle the /score command here
// func (a *application) scoreHandler(m *tbot.Message) {
// 	msg := fmt.Sprintf("Scores:\nWins: %v\nDraws: %v\nLosses: %v", a.wins, a.draws, a.losses)
// 	a.client.SendMessage(m.Chat.ID, msg)
// }

// Handle the /reset command here
// func (a *application) resetHandler(m *tbot.Message) {
// 	a.wins, a.draws, a.losses = 0, 0, 0
// 	msg := "Scores have been reset to 0."
// 	a.client.SendMessage(m.Chat.ID, msg)
// }

// Handle button presses here
func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	humanMove := cq.Data
	msg := draw(humanMove)
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}
