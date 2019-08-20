package localdb

import (
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"sync"
	"time"
	"utils"
)

var (
	once  sync.Once
	Mysql *mysqlClient
	mysqlConf *tomlConfig
)

type tomlConfig struct {
	Title         string
	ConfigVersion int `toml:"config_version"`
	Ticker         mysqlConfig
}

type mysqlConfig struct {
	Server          string
	Port            string
	User            string
	Password        string
	DefaultDatabase string `toml:"default_database"`
}

type mysqlClient struct {
	instance *sql.DB
}

func init() {
	Mysql = &mysqlClient{}
}

func (rc *mysqlClient) setConfig()  {
	var err error
	dbUrl := mysqlConf.Ticker.User + ":" + mysqlConf.Ticker.Password + "@tcp(" + mysqlConf.Ticker.Server + ":" +
		mysqlConf.Ticker.Port + ")/" + mysqlConf.Ticker.DefaultDatabase + "?charset=utf8&timeout=5s"
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
}

func (rc *mysqlClient) GetInstance() *sql.DB {
	once.Do(func() {
		// 载入配置
		viper.SetConfigName("mysql")   // 设置配置文件名 (不带后缀)
		viper.AddConfigPath("conf")        // 第一个搜索路径，可多个
		err := viper.ReadInConfig()     // 读取配置数据
		if err != nil {
			content := fmt.Sprintf("init mysql config error: %s", err)
			fields := make(map[string]interface{})
			fields["type"] = "mysql-config"
			fields["ope"] = "init"
			utils.LogError(content, fields)
		}
		err = viper.Unmarshal(&mysqlConf)
		if err != nil {
			content := fmt.Sprintf("unmarshal mysql config error: %s", err)
			fields := make(map[string]interface{})
			fields["type"] = "mysql-config"
			fields["ope"] = "unmarshal"
			utils.LogError(content, fields)
		}
		// 监听配置文件变动
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			//fmt.Println("Config file changed:", e.Name)
			err := viper.ReadInConfig()     // 读取配置数据
			if err != nil {
				content := fmt.Sprintf("update mysql config error: %s", err)
				fields := make(map[string]interface{})
				fields["type"] = "mysql-config"
				fields["ope"] = "update"
				utils.LogError(content, fields)
			}
			err = viper.Unmarshal(&mysqlConf)
			if err != nil {
				content := fmt.Sprintf("unmarshal mysql config error: %s", err)
				fields := make(map[string]interface{})
				fields["type"] = "mysql-config"
				fields["ope"] = "unmarshal"
				utils.LogError(content, fields)
			}
			rc.setConfig()
			// 记录更改信息
			content := fmt.Sprintf("update mysql config success")
			fields := make(map[string]interface{})
			fields["type"] = "mysql-config"
			fields["ope"] = "update-finish"
			fields["info"] = mysqlConf
			utils.LogWarn(content, fields)
		})
		rc.setConfig()
	})
	return rc.instance
}

func (rc *mysqlClient) GetOne(sqlStr string, columns []string) (result map[string]string) {
	row := rc.GetInstance().QueryRow(sqlStr)
	values := make([ ][]byte, len(columns))
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
