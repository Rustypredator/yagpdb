package util

import (
	"github.com/jonas747/dcmd"
	"github.com/jonas747/yagpdb/bot"
	"github.com/jonas747/yagpdb/common"
)

func RequireBotAdmin(inner dcmd.RunFunc) dcmd.RunFunc {
	return func(data *dcmd.Data) (interface{}, error) {
		if admin, err := bot.IsBotAdmin(data.Msg.Author.ID); admin && err == nil {
			return inner(data)
		}

		return "", nil
	}
}

func RequireOwner(inner dcmd.RunFunc) dcmd.RunFunc {
	return func(data *dcmd.Data) (interface{}, error) {
		if data.Msg.Author.ID == int64(common.ConfOwner.GetInt()) {
			return inner(data)
		}

		return "", nil
	}
}
