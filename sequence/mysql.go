package sequence

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

//简历Mysql连接，执行replace into语句

const sqlReplaceIntoStub = `REPLACE INTO sequence (stub) VALUES ('a')`

type MySQl struct {
	conn sqlx.SqlConn
}

func NewMySQL(dsn string) Sequence {
	return &MySQl{
		conn: sqlx.NewMysql(dsn),
	}
}

// 取下一个号
func (m *MySQl) Next() (seq uint64, err error) {
	var stmt sqlx.StmtSession
	stmt, err = m.conn.Prepare(sqlReplaceIntoStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	defer stmt.Close()

	//执行
	var rest sql.Result
	rest, err = stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec() failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	//获取刚插入的主键id
	var lid int64
	lid, err = rest.LastInsertId()
	if err != nil {
		logx.Errorw("rest.LastInsertId() failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	return uint64(lid), nil
}
