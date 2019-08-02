package finance

import (
	"fmt"
	"localdb"
	"utils"
)

const (
	FinanceFlowTable = "financial_data4hour"
)

func GetCoinFinance(coinKey string, startTime uint32, endTime uint32) (ret []map[string]string) {
	sql := fmt.Sprintf("SELECT `coin_key`, `small_in`, `middle_in`, `big_in`, `super_in`, `small_out`, "+
		"`middle_out`, `big_out`, `super_out`, `finance_date`, `updated` FROM `ticker`.`%s` WHERE `coin_key`='%s' "+
		"AND `finance_date` >= %d AND `finance_date` <= %d", FinanceFlowTable, coinKey, startTime, endTime)
	ret = localdb.Mysql.GetAll(sql)
	content := fmt.Sprintf("test log")
	fields := make(map[string]interface{})
	fields["type"] = "test"
	fields["ope"] = "test-p"
	utils.LogError(content, fields)
	return
}
