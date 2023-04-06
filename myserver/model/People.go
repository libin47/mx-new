package model

// 取保候审人员
import (
	"github.com/qiniu/qmgo/field"
)

type People struct {
	field.DefaultField `bson:",inline"`
	Name               string `bson:"name"`
	Phone              string `bson:"phone"`
	IDNumber           string `bson:"id_number"`
	IDEvent            string `bson:"id_event"`
	Gender             string `bson:"gender"`
	Age                string `bson:"age"`
	DateBorn           string `bson:"date_born"`
	DateStart          string `bson:"data_start"`
	DateEnd            string `bson:"data_end"`
	Stauts             string `bson:"stauts"`
	Danger             string `bson:"danger"`
	Organ              string `bson:"organ"`
}

// 姓名 手机号 身份证号 事件编号
// 性别 年龄
// 检测开始时间 结束时间 状态 危险级别
// 负责机关 管理人员 管理人员电话
// 认证文件
