/**
 * @Author: Cc
 * @Description: 描述
 * @File: game
 * @Version: 1.0.0
 * @Date: 2023/7/28 15:03
 * @Software : GoLand
 */

package mysql

import (
	"github.com/Cc360428/HelpPackage/other"
	"github.com/Cc360428/telegram_robot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	GameId     int32 `gorm:"index:idx_name,unique"`
	Name       string
	ServerName string
	Platform   int32
	RowId      int32
}

// insert into t_game(game_id,server_name,name,platform) value(20041,'TexasPoker','德州',2);

func GetGameAllByPlatform(platformId int32) tgbotapi.InlineKeyboardMarkup {

	row := make([][]tgbotapi.InlineKeyboardButton, 0)

	var games []Game
	client.Raw("select game_id,server_name,name,row_id from t_game where platform = ?", platformId).Scan(&games)

	rowsIn := make(map[int][]tgbotapi.InlineKeyboardButton)
	for _, game := range games {
		row = append(row)
		rowsIn[int(game.RowId)] = append(rowsIn[int(game.RowId)], tgbotapi.NewInlineKeyboardButtonData(game.Name, other.Int32ToStr(game.GameId)))
	}

	for _, button := range rowsIn {
		row = append(row, button)
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: utils.SortMapInt64(rowsIn)}
}
