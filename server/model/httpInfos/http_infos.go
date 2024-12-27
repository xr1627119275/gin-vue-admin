// 自动生成模板HttpInfos
package httpInfos

import (
	"gorm.io/datatypes"
	"time"
)

// httpInfos表 结构体  HttpInfos
type HttpInfos struct {
	CreatedAt    *time.Time     `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`                               //createdAt字段
	DeletedAt    *time.Time     `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`                               //deletedAt字段
	Headers      datatypes.JSON `json:"headers" form:"headers" gorm:"column:headers;comment:;type:text;"swaggertype:"object"`        //headers字段
	Host         string         `json:"host" form:"host" gorm:"column:host;comment:域名;size:191;"`                                    //域名
	Id           *int           `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                                  //id字段
	Path         string         `json:"path" form:"path" gorm:"column:path;type:text;comment:路径;"`                                   //路径
	RemoteAddr   string         `json:"remoteAddr" form:"remoteAddr" gorm:"column:remote_addr;comment:原始IP;size:191;"`               //原始IP
	ResContent   string         `json:"resContent" form:"resContent" gorm:"column:res_content;type:longtext;comment:响应内容字符串;"`       //响应内容字符串
	ResData      []byte         `json:"resData" form:"resData" gorm:"column:res_data;comment:响应数据;"`                                 //响应数据
	ResHeader    datatypes.JSON `json:"resHeader" form:"resHeader" gorm:"column:res_header;comment:;type:text;"swaggertype:"object"` //resHeader字段
	StatusCode   *int           `json:"statusCode" form:"statusCode" gorm:"column:status_code;comment:响应码;size:19;"`                 //响应码
	Method       string         `json:"method" form:"method" gorm:"column:method;comment:请求方法;size:191;"`                            //请求方法
	Proto        string         `json:"proto" form:"proto" gorm:"column:proto;comment:协议;size:191;"`                                 //协议
	ReqContent   string         `json:"reqContent" form:"reqContent" gorm:"column:req_content;type:text;comment:请求内容字符串;"`           //请求内容字符串
	ReqData      []byte         `json:"reqData" form:"reqData" gorm:"column:req_data;comment:请求数据;"`                                 //请求数据
	UpdatedAt    *time.Time     `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`
	WeakPassword bool           `json:"weak_pass" gorm:"type:bool;comment:弱密码"`
	//updatedAt字段
}

// TableName httpInfos表 HttpInfos自定义表名 http_infos
func (HttpInfos) TableName() string {
	return "http_infos"
}
