package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

type IntrusionPredefined struct {
	ID                  int64  `gorm:"primarykey" json:"id"` // 主键ID
	Sid                 int64  `json:"sid"`
	RuleName            string `json:"ruleName"`
	RuleSource          string `json:"ruleSource"`
	OccurDate           string `json:"occurDate"`
	ClassType           string `json:"classType"`
	RiskLevel           int8   `json:"riskLevel"` //1低 2中 3高
	VulnerabilitySource string `json:"vulnerabilitySource"`
	AttackRequirement   string `json:"attackRequirement"`
	Cve                 string `json:"cve"`
	Description         string `json:"description"`
	SignName            string `json:"signName"`
	RuleMessage         string `json:"ruleMessage"`
	Type                string `json:"type"`
	Action              int8   `json:"action" form:"action" gorm:"column:action;comment:动作，1告警, 2阻断;"`
	TriggerDevice       string `json:"triggerDevice" form:"triggerDevice" gorm:"column:trigger_device;comment:漏洞类型;size:255;"`
	EventHandle         int    `json:"eventHandle" form:"eventHandle" gorm:"column:event_handle;comment:规则描述;"`
	Priority            int    `json:"priority" form:"priority" gorm:"column:priority;comment:;size:10;"`
	RuleDesc            string `json:"ruleDesc" form:"ruleDesc" gorm:"column:rule_desc;comment:;"`
	RuleDiff            int    `json:"ruleDiff" form:"ruleDiff" gorm:"column:rule_diff;comment:;"`
	Message             string `json:"-" form:"-" gorm:"column:message;comment:;size:512;"`
	Severity            int8   `json:"severity" form:"severity" gorm:"column:severity;comment:危害级别0低 1中 2高;size:255;"`
}

func (IntrusionPredefined) TableName() string {
	return "policy_intrusion_dictionary"
}

var (
	NETVINE_DB *gorm.DB
)

func GormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       "root:Netvine123#@!@tcp(10.25.30.123:3306)/audit?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         191,                                                                                     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                                                                                   // 根据版本自动配置
	}
RETRY:
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		//mysql重连
		fmt.Println("连接失败")
		goto RETRY
	} else {
		fmt.Println("连接成功")
		return db
	}
	return nil
}

func main() {
	NETVINE_DB = GormMysql() // gorm连接数据库

	open, _ := os.Open("sid.txt")
	content, _ := ioutil.ReadAll(open)
	fmt.Println(content)
	//var pSid []int
	//NETVINE_DB.Model(&IntrusionPredefined{}).Select("sid").Find(&pSid)

	fmt.Println("====")
}
