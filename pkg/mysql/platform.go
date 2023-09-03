/**
 * @Author: Cc
 * @Description: 平台
 * @File: platform
 * @Version: 1.0.0
 * @Date: 2023/7/28 15:10
 * @Software : GoLand
 */

package mysql

import "gorm.io/gorm"

type Platform struct {
	gorm.Model
	Name string
}
