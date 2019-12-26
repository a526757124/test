package models

//任务配置
type Task struct {
	TaskID      string `xml:"taskid,attr" yaml:"taskid"`           //task编号，需唯一
	IsRun       bool   `xml:"isrun,attr" yaml:"isrun"`             //标识是否允许task执行,默认为false,如设为flash,则启动后不执行task
	TaskType    string `xml:"type,attr" yaml:"type"`               //Task类型,目前支持loop、cron、queue
	DueTime     int64  `xml:"duetime,attr" yaml:"duetime"`         //开始任务的延迟时间（以毫秒为单位），如果<=0则不延迟
	Interval    int64  `xml:"interval,attr" yaml:"interval"`       //loop类型下,两次Task执行之间的间隔,单位为毫秒
	Express     string `xml:"express,attr" yaml:"express"`         //cron类型下,task执行的时间表达式，具体参考readme
	QueueSize   int64  `xml:"queuesize,attr" yaml:"queuesize"`     //queue类型下,queue初始长度
	HandlerName string `xml:"handlername,attr" yaml:"handlername"` //Task对应的HandlerName,需使用RegisterHandler进行统一注册
	HandlerData string `xml:"handlerdata,attr" yaml:"handlerdata"` //Task对应的自定义数据,可在配置源中设置
}
