package main

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "Are you bored at home from CB? 😣 Watched too much Big Bang Theory but no one to play *Rock Paper Scissors Lizard Spock* with you? 🥺 Here's the perfect bot for you!"
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle the /play command here
func (a *application) playHandler(m *tbot.Message) {
	buttons := makeButtons()
	a.client.SendMessage(m.Chat.ID, "Pick an option:", tbot.OptInlineKeyboardMarkup(buttons))
}

// Handle the /score command here
func (a *application) scoreHandler(m *tbot.Message) {
	msg := fmt.Sprintf("Scores:\nWins: %v\nDraws: %v\nLosses: %v", a.wins, a.draws, a.losses)
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle the /reset command here
func (a *application) resetHandler(m *tbot.Message) {
	a.wins, a.draws, a.losses = 0, 0, 0
	msg := "Scores have been reset to 0."
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle button presses here
func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	humanMove := cq.Data
	msg := a.draw(humanMove)
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}

func (a *application) rulesHandler(m *tbot.Message) {
	msg := "Rock, Paper, Scissors, Lizard, Spock is an expansion on the game Rock, Paper Scissors, first used by Sheldon and Raj in Big Bang Theory to settle a dispute about what to watch on TV. \n\nHere are the rules of the game: \n• Scissors cuts Paper \n• Paper covers Rock \n• Rock crushes Lizard \n• Lizard poisons Spock \n• Spock smashes Scissors \n• Scissors decapitates Lizard \n• Lizard eats Paper \n• Paper disproves Spock \n• Spock vaporizes Rock \n• (and as always) Rock crushes Scissors"
	a.client.SendMessage(m.Chat.ID, msg)
}
