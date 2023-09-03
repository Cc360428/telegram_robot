/**
 * @Author: Cc
 * @Description: 描述
 * @File: utils
 * @Version: 1.0.0
 * @Date: 2023/7/26 09:52
 * @Software : GoLand
 */

package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sort"
)

// SortMapInt64 ...
// @Description:
// @param rate
// @return playerIds
func SortMapInt64(rate map[int][]tgbotapi.InlineKeyboardButton) (teKeys [][]tgbotapi.InlineKeyboardButton) {

	keys := make([]int, 0, len(rate))
	for k := range rate {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// 打印排序结果
	for _, k := range keys {
		teKeys = append(teKeys, rate[k])
	}
	return teKeys
}
