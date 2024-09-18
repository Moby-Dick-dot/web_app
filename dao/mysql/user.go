package mysql

import (
	"database/sql"
	g "web_app/global"
	"web_app/model"
	"web_app/request"
	"web_app/utils"
)

func Register(user *model.User) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	err = db.Get(&count, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count > 0 {
		// 用户已存在
		return g.ErrorUserExit
	}
	// 生成user_id
	userID, err := utils.GetSnowFlakeID()
	if err != nil {
		return g.ErrorGenIDFailed
	}
	// 生成加密密码
	password := utils.EncryptPassword([]byte(user.Password))
	// 把用户插入数据库
	sqlStr = "insert into user(user_id, username, password) values (?,?,?)"
	_, err = db.Exec(sqlStr, userID, user.UserName, password)
	return
}

func Login(user *request.User) (err error) {
	originPassword := user.Password // 记录一下原始密码
	sqlStr := "select user_id, username, password from user where username = ?"
	err = db.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		return g.ErrorUserNotExit
	}
	password := utils.EncryptPassword([]byte(originPassword))
	if user.Password != password {
		return g.ErrorPasswordWrong
	}
	return
}
