/**
 * @Author: Cc
 * @Description: 描述
 * @File: keyboard
 * @Version: 1.0.0
 * @Date: 2023/7/25 16:48
 * @Software : GoLand
 */

package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var NumericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("GameInfo"),
		tgbotapi.NewKeyboardButton("ShareInfo"),
	),
)
