package service

import (
	"gin_demo/dao"
)

func Login() {

	// 去数据库查询用户名密码
	dao.Login() // 具体的数据库操作

	// redis操作

	// 发通知短信
	// sms.Send()
}
