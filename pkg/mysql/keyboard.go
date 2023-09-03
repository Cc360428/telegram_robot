/**
 * @Author: Cc
 * @Description: 键盘
 * @File: keyboard
 * @Version: 1.0.0
 * @Date: 2023/7/28 15:23
 * @Software : GoLand
 */

package mysql

import (
	"github.com/Cc360428/telegram_robot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type KeyBoardType int

const (
	UpdateVersion KeyBoardType = 2
	ShareDocument KeyBoardType = 3
	ShellCmd      KeyBoardType = 4
)

type KeyBoardRow int

const (
	KeyBoardRowOne   KeyBoardRow = 1
	KeyBoardRowTwo   KeyBoardRow = 2
	KeyBoardRowThree KeyBoardRow = 3
)

type Keyboard struct {
	gorm.Model
	Name  string
	Data  string
	Type  int32 // 版本更新
	RowId int32 // 行
}

func GetKeyboardAllByType(keyType KeyBoardType) tgbotapi.InlineKeyboardMarkup {
	row := make([][]tgbotapi.InlineKeyboardButton, 0)
	var keyboards []Keyboard

	client.Raw("select `name`,`data`,`row_id` from t_keyboard where type = ?", keyType).Scan(&keyboards)
	switch keyType {
	case UpdateVersion:
		rowIn := make([]tgbotapi.InlineKeyboardButton, 0)

		for _, keyboard := range keyboards {
			rowIn = append(rowIn, tgbotapi.NewInlineKeyboardButtonData(keyboard.Name, keyboard.Data))
		}

		row = append(row, rowIn)
	case ShareDocument:
		rowsIn := make(map[int][]tgbotapi.InlineKeyboardButton)
		for _, keyboard := range keyboards {
			rowsIn[int(keyboard.RowId)] = append(rowsIn[int(keyboard.RowId)], tgbotapi.NewInlineKeyboardButtonURL(keyboard.Name, keyboard.Data))
		}

		for _, button := range rowsIn {
			row = append(row, button)
		}

		return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: utils.SortMapInt64(rowsIn)}
	case ShellCmd:
		rowIn := make([]tgbotapi.InlineKeyboardButton, 0)

		for _, keyboard := range keyboards {
			rowIn = append(rowIn, tgbotapi.NewInlineKeyboardButtonData(keyboard.Name, keyboard.Data))
		}

		row = append(row, rowIn)
	}

	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: row}
}
