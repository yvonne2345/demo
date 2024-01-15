package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BlacklistVulnerability struct {
	ID                  int64  `gorm:"primarykey" json:"id"` // 主键ID
	DetailId            int64  `json:"detailId" form:"detailId" gorm:"column:detail_Id;comment:;size:20;"`
	RiskLevel           int8   `json:"riskLevel" form:"riskLevel" gorm:"column:risk_level;comment:风险等级，0低1中2高;"`
	Name                string `json:"name" form:"name" gorm:"column:name;comment:漏洞名称;size:255;"`
	PublishDate         string `json:"publishDate" form:"publishDate" gorm:"column:publish_date;comment:发布时间;"`
	Action              int8   `json:"action" form:"action" gorm:"column:action;comment:动作，1告警, 2阻断;"`
	Status              int8   `json:"status" form:"status" gorm:"column:status;comment:开启状态，0 关闭, 1 开启;"`
	Cve                 string `json:"cve" form:"cve" gorm:"column:cve;comment:漏洞统一编号;size:255;"`
	VulnerabilitySource string `json:"vulnerabilitySource" form:"vulnerabilitySource" gorm:"column:vulnerability_source;comment:漏洞来源;"`
	Severity            int8   `json:"severity" form:"severity" gorm:"column:severity;comment:危害级别0低 1中 2高;size:255;"`
	TriggerDevice       string `json:"triggerDevice" form:"triggerDevice" gorm:"column:trigger_device;comment:漏洞类型;size:255;"`
	RuleSource          string `json:"ruleSource" form:"ruleSource" gorm:"column:rule_source;comment:;size:255;"`
	EventHandle         int    `json:"eventHandle" form:"eventHandle" gorm:"column:event_handle;comment:规则描述;"`
	AffectedFirm        string `json:"affectedFirm" form:"affectedFirm" gorm:"column:affected_firm;comment:;"`
	AttackRequirement   string `json:"attackRequirement" form:"attackRequirement" gorm:"column:attack_requirement;comment:攻击条件;"`
	Description         string `json:"description" form:"description" gorm:"column:description;comment:;"`
	Type                string `json:"type" form:"type" gorm:"column:type;comment:;"`
	Sid                 int64  `json:"sid" form:"sid" gorm:"column:sid;comment:;size:1024;"`
	SignName            string `json:"signName" form:"signName" gorm:"column:sign_name;comment:;size:10;"`
	Priority            int    `json:"priority" form:"priority" gorm:"column:priority;comment:;size:10;"`
	Message             string `json:"-" form:"-" gorm:"column:message;comment:;size:512;"`
}

func (b *BlacklistVulnerability) TableName() string {
	return "security_blacklist_vulnerability"
}

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
		DSN:                       "root:Netvine123#@!@tcp(10.25.30.131:3306)/audit?charset=utf8&parseTime=True&loc=Local", // DSN data source name
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

	var pSid []int
	NETVINE_DB.Model(&IntrusionPredefined{}).Select("sid").Find(&pSid)
	//去除alter
	//for _, p := range pSid {
	//	var str string
	//	NETVINE_DB.Model(&IntrusionPredefined{}).Select("message").Where("sid = ?", p).Find(&str)
	//	if str == "" {
	//		var msg string
	//		NETVINE_DB.Model(&IntrusionPredefined{}).Select("rule_message").Where("sid = ?", p).Find(&msg)
	//		rule := strings.ReplaceAll(msg, "alert ", "")
	//		NETVINE_DB.Model(&IntrusionPredefined{}).Where("sid = ?", p).Update("message", rule)
	//	}
	//}

	var bSid []int
	NETVINE_DB.Model(&BlacklistVulnerability{}).Select("sid").Find(&bSid)

	var sameSid []int
	var sameSidMap = make(map[int]string)
	var diffSid []int
	for _, b := range bSid {
		for _, p := range pSid {
			if b == p {
				sameSid = append(sameSid, b)
				sameSidMap[b] = "ok"
			}
		}
	}
	fmt.Println(len(sameSid))

	for _, b := range bSid {
		_, ok := sameSidMap[b]
		if !ok {
			diffSid = append(diffSid, b)
		}
	}
	fmt.Println(len(diffSid))

	for _, d := range diffSid {
		var Blacklist BlacklistVulnerability
		var Intrusion IntrusionPredefined
		NETVINE_DB.Model(&BlacklistVulnerability{}).Where("sid = ?", d).Find(&Blacklist)
		Intrusion.Sid = Blacklist.Sid
		Intrusion.RuleName = Blacklist.Name
		Intrusion.RuleSource = Blacklist.RuleSource
		Intrusion.OccurDate = Blacklist.PublishDate
		Intrusion.RuleMessage = "alter " + Blacklist.Message
		Intrusion.RiskLevel = 3
		Intrusion.VulnerabilitySource = Blacklist.VulnerabilitySource
		Intrusion.Cve = Blacklist.Cve
		Intrusion.AttackRequirement = Blacklist.AttackRequirement
		Intrusion.Description = Blacklist.AffectedFirm
		Intrusion.SignName = Blacklist.SignName
		Intrusion.ClassType = Blacklist.Type
		Intrusion.Action = Blacklist.Action
		Intrusion.TriggerDevice = Blacklist.TriggerDevice
		Intrusion.EventHandle = Blacklist.EventHandle
		Intrusion.Priority = Blacklist.Priority
		Intrusion.RuleDesc = Blacklist.Description
		Intrusion.Message = Blacklist.Message
		Intrusion.RuleDiff = 1
		Intrusion.Severity = Blacklist.Severity
		//fmt.Println(Intrusion)
		NETVINE_DB.Model(&IntrusionPredefined{}).Create(&Intrusion)
	}

	//sid一样
	//for _, b := range bSid {
	//	for _, p := range pSid {
	//		if b == p {
	//			sameSid = append(sameSid, b)
	//		}
	//	}
	//}
	//fmt.Println(len(sameSid))
	//for _, v := range sameSid {
	//	var Blacklist BlacklistVulnerability
	//	var Intrusion IntrusionPredefined
	//	NETVINE_DB.Model(&BlacklistVulnerability{}).Where("sid = ?", v).Find(&Blacklist)
	//	NETVINE_DB.Model(&IntrusionPredefined{}).Where("sid =?", v).Find(&Intrusion)
	//
	//}
	//for _, s := range sameSid {
	//
	//	var Blacklist BlacklistVulnerability
	//	var Intrusion IntrusionPredefined
	//	NETVINE_DB.Model(&BlacklistVulnerability{}).Where("sid = ?", s).Find(&Blacklist)
	//	Intrusion.Sid = Blacklist.Sid
	//	Intrusion.RuleName = Blacklist.Name
	//	Intrusion.RuleSource = Blacklist.RuleSource
	//	Intrusion.OccurDate = Blacklist.PublishDate
	//	//Intrusion.RuleMessage = "alter " + Blacklist.Message
	//	Intrusion.RiskLevel = 3
	//	Intrusion.VulnerabilitySource = Blacklist.VulnerabilitySource
	//	Intrusion.Cve = Blacklist.Cve
	//	Intrusion.AttackRequirement = Blacklist.AttackRequirement
	//	Intrusion.Description = Blacklist.AffectedFirm
	//	Intrusion.SignName = Blacklist.SignName
	//
	//	Intrusion.ClassType = Blacklist.Type
	//	Intrusion.Action = Blacklist.Action
	//	Intrusion.TriggerDevice = Blacklist.TriggerDevice
	//	Intrusion.EventHandle = Blacklist.EventHandle
	//	Intrusion.Priority = Blacklist.Priority
	//	Intrusion.RuleDesc = Blacklist.Description
	//	Intrusion.Message = Blacklist.Message
	//	Intrusion.RuleDiff = 3
	//	Intrusion.Severity = Blacklist.Severity
	//
	//	fmt.Println(Intrusion)
	//	NETVINE_DB.Model(&IntrusionPredefined{}).Where("sid =?", s).Updates(&Intrusion)
	//}
	//fmt.Println(sameSid)

	fmt.Println("====")
}
