//访问用户登录日志信息
package base

import (
	"pms/models/base"
	"pms/utils"
	"strconv"
)

//列表视图列数-1，第一列为checkbox

const (
	recordListCellLength = 8
)

type RecordController struct {
	BaseController
}

func (this *RecordController) Get() {
	action := this.GetString(":action")
	switch action {
	case "list":
		this.List()
	default:
		this.List()

	}
}
func (this *RecordController) List() {

	this.Data["listName"] = "登录日志"
	this.Layout = "base/base.html"
	this.TplName = "user/record_list.html"
	condArr := make(map[string]interface{})
	page := this.Input().Get("page")
	offset := this.Input().Get("offset")
	var (
		err         error
		pageInt64   int64
		offsetInt64 int64
	)
	if pageInt, ok := strconv.Atoi(page); ok == nil {
		pageInt64 = int64(pageInt)
	}
	if offsetInt, ok := strconv.Atoi(offset); ok == nil {
		offsetInt64 = int64(offsetInt)
	}
	var records []base.Record
	paginator, err, records := base.ListRecord(condArr, this.User.Id, pageInt64, offsetInt64)
	paginator.Url = "/record"
	this.Data["Paginator"] = paginator
	tableInfo := new(utils.TableInfo)
	tableInfo.Url = "/record"
	tableTitle := make(map[string]interface{})
	tableTitle["titleName"] = [recordListCellLength]string{"邮箱", "手机", "用户名", "中文用户名", "开始时间", "结束时间", "登录IP", "操作"}
	tableInfo.Title = tableTitle
	tableBody := make(map[string]interface{})
	bodyLines := make([]interface{}, 0, 20)
	if err == nil {
		for _, record := range records {
			oneLine := make([]interface{}, recordListCellLength, recordListCellLength)
			lineInfo := make(map[string]interface{})
			action := map[string]map[string]string{}
			id := int(record.Id)

			lineInfo["id"] = id
			oneLine[0] = record.User.Email
			oneLine[1] = record.User.Mobile
			oneLine[2] = record.User.Name
			oneLine[3] = record.User.NameZh

			oneLine[4] = record.CreateDate.Format("2006-01-02 15:04:05")
			oneLine[5] = record.Logout.Format("2006-01-02 15:04:05")
			oneLine[6] = record.Ip

			oneLine[7] = action
			lineData := make(map[string]interface{})
			lineData["oneLine"] = oneLine
			lineData["lineInfo"] = lineInfo
			bodyLines = append(bodyLines, lineData)
		}
		tableBody["bodyLines"] = bodyLines
		tableInfo.Body = tableBody
		tableInfo.TitleLen = recordListCellLength
		tableInfo.TitleIndexLen = recordListCellLength - 1
		tableInfo.BodyLen = paginator.CurrentPageSize
		this.Data["tableInfo"] = tableInfo
	}
}
