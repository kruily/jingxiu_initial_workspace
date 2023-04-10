/**
* @file: connector.go ==> data/connector
* @package: connector
* @author: jingxiu
* @since: 2023/2/23
* @desc: //TODO
 */

package connector

import (
	"gateway/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var conn *gorm.DB
var once sync.Once

func GetConnectDB() *gorm.DB {
	once.Do(func() {
		c, err := gorm.Open(mysql.Open(config.C.DB.Source))
		if err != nil {
			panic(err)
		}
		conn = c
	})
	return conn
}
