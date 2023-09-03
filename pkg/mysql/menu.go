/**
 * @Author: Cc
 * @Description: 菜单
 * @File: menu
 * @Version: 1.0.0
 * @Date: 2023/7/28 15:21
 * @Software : GoLand
 */

package mysql

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name string
}
