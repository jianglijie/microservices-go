package localdb

import (
	"config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
	"utils"
)

var (
	once  sync.Once
	Mysql *mysqlClient
)

type mysqlClient struct {
	instance *sql.DB
}

func init() {
	Mysql = &mysqlClient{}
}

func (rc *mysqlClient) GetInstance() *sql.DB {
	once.Do(func() {
		dbUrl := config.Config().Mysql.User + ":" + config.Config().Mysql.Password + "@tcp(" +
			config.Config().Mysql.Server + ":" + config.Config().Mysql.Port + ")/" +
			config.Config().Mysql.DefaultDatabase + "?charset=utf8"

		var err error
		rc.instance, err = sql.Open("mysql", dbUrl)
		if err != nil {
			content := fmt.Sprintf("mysql open error: %s", err)
			fields := make(map[string]interface{})
			fields["type"] = "mysql"
			fields["ope"] = "open"
			utils.LogError(content, fields)
		}

		//设置参数
		rc.instance.SetMaxOpenConns(5000)
		rc.instance.SetMaxIdleConns(1000)
		rc.instance.SetConnMaxLifetime(30 * time.Minute)

		err = rc.instance.Ping()
		if err != nil {
			content := fmt.Sprintf("mysql ping error: %s", err)
			fields := make(map[string]interface{})
			fields["type"] = "mysql"
			fields["ope"] = "ping"
			utils.LogError(content, fields)
		}
	})
	return rc.instance
}

func (rc *mysqlClient) GetOne(sqlStr string, columns []string) (result map[string]string) {
	row := rc.GetInstance().QueryRow(sqlStr)
	values := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))
	result = make(map[string]string)

	for i := range values {
		scans[i] = &values[i]
	}
	if err := row.Scan(scans...); err != nil && err != sql.ErrNoRows {
		content := fmt.Sprintf("mysql row scan error: %s", err)
		fields := make(map[string]interface{})
		fields["type"] = "mysql"
		fields["ope"] = "scan"
		fields["sql"] = sqlStr
		utils.LogError(content, fields)
		return
	}
	for key, val := range values {
		result[columns[key]] = string(val)
	}
	return
}

func (rc *mysqlClient) GetAll(sqlStr string) (result []map[string]string) {
	rows, err := rc.GetInstance().Query(sqlStr)
	if err != nil {
		content := fmt.Sprintf("mysql query error: %s", err)
		fields := make(map[string]interface{})
		fields["type"] = "mysql"
		fields["ope"] = "query"
		fields["sql"] = sqlStr
		utils.LogError(content, fields)
		return
	}
	defer rows.Close()
	columns, err := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		_ = rows.Scan(scans...)
		each := make(map[string]string)
		for i, col := range values {
			each[columns[i]] = string(col)
		}
		result = append(result, each)
	}
	return
}
