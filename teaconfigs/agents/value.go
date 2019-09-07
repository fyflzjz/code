package agents

import (
	"encoding/json"
	"github.com/TeaWeb/code/teaconfigs/notices"
	"github.com/TeaWeb/code/teadb/shared"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/utils/time"
	"github.com/pquerna/ffjson/ffjson"
	"time"
)

// 应用指标值定义
type Value struct {
	Id          shared.ObjectId     `bson:"_id" json:"id"`                  // 数据库存储的ID
	NodeId      string              `bson:"nodeId" json:"nodeId"`           // 节点ID
	AgentId     string              `bson:"agentId" json:"agentId"`         // Agent ID
	AppId       string              `bson:"appId" json:"appId"`             // App ID
	ItemId      string              `bson:"itemId" json:"itemId"`           // 监控项ID
	Timestamp   int64               `bson:"timestamp" json:"timestamp"`     // Agent时间戳
	CreatedAt   int64               `bson:"createdAt" json:"createdAt"`     // Master时间戳
	Value       interface{}         `bson:"value" json:"value"`             // 值，可以是个标量，或者一个组合的值
	Error       string              `bson:"error" json:"error"`             // 错误信息
	NoticeLevel notices.NoticeLevel `bson:"noticeLevel" json:"noticeLevel"` // 通知级别
	IsNotified  bool                `bson:"isNotified" json:"isNotified"`   // 是否已通知
	ThresholdId string              `bson:"thresholdId" json:"thresholdId"` // 阈值ID
	Threshold   string              `bson:"threshold" json:"threshold"`     // 阈值描述
	TimeFormat  struct {
		Year   string `bson:"year" json:"year"`
		Month  string `bson:"month" json:"month"`
		Week   string `bson:"week" json:"week"`
		Day    string `bson:"day" json:"day"`
		Hour   string `bson:"hour" json:"hour"`
		Minute string `bson:"minute" json:"minute"`
		Second string `bson:"second" json:"second"`
	} `bson:"timeFormat" json:"timeFormat"`
}

// 获取新对象
func NewValue() *Value {
	return &Value{
		Id: shared.NewObjectId(),
	}
}

// 设置时间
func (this *Value) SetTime(t time.Time) {
	this.Timestamp = t.Unix()
	this.TimeFormat.Year = timeutil.Format("Y", t)
	this.TimeFormat.Month = timeutil.Format("Ym", t)
	this.TimeFormat.Week = timeutil.Format("YW", t)
	this.TimeFormat.Day = timeutil.Format("Ymd", t)
	this.TimeFormat.Hour = timeutil.Format("YmdH", t)
	this.TimeFormat.Minute = timeutil.Format("YmdHi", t)
	this.TimeFormat.Second = timeutil.Format("YmdHis", t)
}

// 设置数据库列值
func (this *Value) SetDBColumns(v maps.Map) {
	id, err := shared.ObjectIdFromHex(v.GetString("_id"))
	if err != nil {
		logs.Error(err)
	} else {
		this.Id = id
	}
	this.NodeId = v.GetString("nodeId")
	this.AgentId = v.GetString("agentId")
	this.AppId = v.GetString("appId")
	this.ItemId = v.GetString("itemId")
	this.Timestamp = v.GetInt64("timestamp")
	this.CreatedAt = v.GetInt64("createdAt")
	this.jsonDecode(v.Get("value"), &this.Value)
	this.Error = v.GetString("error")
	this.NoticeLevel = v.GetUint8("noticeLevel")
	this.IsNotified = v.GetInt("isNotified") > 0
	this.ThresholdId = v.GetString("thresholdId")
	this.Threshold = v.GetString("threshold")
	this.TimeFormat.Year = v.GetString("timeFormat_year")
	this.TimeFormat.Month = v.GetString("timeFormat_month")
	this.TimeFormat.Week = v.GetString("timeFormat_week")
	this.TimeFormat.Day = v.GetString("timeFormat_day")
	this.TimeFormat.Hour = v.GetString("timeFormat_hour")
	this.TimeFormat.Minute = v.GetString("timeFormat_minute")
	this.TimeFormat.Second = v.GetString("timeFormat_second")
}

// 获取数据库列值
func (this *Value) DBColumns() maps.Map {
	if this.Id.IsZero() {
		this.Id = shared.NewObjectId()
	}
	valueJSON, err := json.Marshal(this.Value)
	if err != nil {
		logs.Error(err)
	}
	return maps.Map{
		"_id":               this.Id.Hex(),
		"nodeId":            this.NodeId,
		"agentId":           this.AgentId,
		"appId":             this.AppId,
		"itemId":            this.ItemId,
		"timestamp":         this.Timestamp,
		"createdAt":         this.CreatedAt,
		"value":             valueJSON,
		"error":             this.Error,
		"noticeLevel":       this.NoticeLevel,
		"isNotified":        this.IsNotified,
		"thresholdId":       this.ThresholdId,
		"threshold":         this.Threshold,
		"timeFormat_year":   this.TimeFormat.Year,
		"timeFormat_month":  this.TimeFormat.Month,
		"timeFormat_week":   this.TimeFormat.Week,
		"timeFormat_day":    this.TimeFormat.Day,
		"timeFormat_hour":   this.TimeFormat.Hour,
		"timeFormat_minute": this.TimeFormat.Minute,
		"timeFormat_second": this.TimeFormat.Second,
	}
}

func (this *Value) jsonDecode(data interface{}, vPtr interface{}) {
	if data == nil {
		return
	}
	b, ok := data.([]byte)
	if ok {
		_ = ffjson.Unmarshal(b, vPtr)
	}
	s, ok := data.(string)
	if ok {
		_ = ffjson.Unmarshal([]byte(s), vPtr)
	}
}
