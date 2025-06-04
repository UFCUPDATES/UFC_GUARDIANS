package buttons

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func StartPanel(b *gotgbot.Bot) gotgbot.InlineKeyboardMarkup {
	btn := &Button{RowWidth: 2}

	btn.Add(
		btn.Url("🔄 Update Channel", "https://t.me/Silenthrex"),
		btn.Url("💬 Update Group", "https://t.me/BESTIE_UNITE_CLUB"),
	)

	btn.Row(
		btn.Inline("❓ Help & Commands", "help"),
	)

	btn.Row(
		btn.Url(
			"➕ Add me to Your Group",
			fmt.Sprintf("https://t.me/%s?startgroup=s&admin=delete_messages+invite_users", b.User.Username),
		),
	)

	return btn.Build()
}

func NormalStartPanel(b *gotgbot.Bot) gotgbot.InlineKeyboardMarkup {
	btn := &Button{RowWidth: 2}

	btn.Add(
		btn.Url("🔄 Update Channel", "https://t.me/Silenthrex"),
		btn.Url("💬 Update Group", "https://t.me/BESTIE_UNITE_CLUB"),
	)

	btn.Row(
		btn.Url(
			"➕ Add me to Your Group",
			fmt.Sprintf("https://t.me/%s?startgroup=s&admin=delete_messages+invite_users", b.User.Username),
		),
	)

	return btn.Build()
}
