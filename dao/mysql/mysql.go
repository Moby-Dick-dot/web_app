package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	g "web_app/global"
	"web_app/model"
)

// 定义一个全局对象db
var db *sqlx.DB

// Init 定义一个初始化数据库的函数
func Init() (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed,err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}

func Close() {
	_ = db.Close()
}

func GetCommunityList() (communityList []*model.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	return
}

func GetCommunityByID(idStr string) (community *model.CommunityDetail, err error) {
	community = new(model.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, idStr)
	if err == sql.ErrNoRows {
		err = g.ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = g.ErrorQueryFailed
		return
	}
	return
}
