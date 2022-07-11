package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

{{ range .PluginModel }}
func Get{{.ModelName}}List(c *gin.Context) {
	var get{{.ModelName}}List model.RequestGet{{.ModelName}}List
	if err := gqaModel.RequestShouldBindJSON(c, &get{{.ModelName}}List); err != nil {
		return
	}
	if err, records, total := privateservice.Get{{.ModelName}}List(get{{.ModelName}}List, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取{{.ModelName}}列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取{{.ModelName}}列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  records,
			Page:     get{{.ModelName}}List.Page,
			PageSize: get{{.ModelName}}List.PageSize,
			Total:    total,
		}, c)
	}
}

func Edit{{.ModelName}}(c *gin.Context) {
	var toEdit{{.ModelName}} model.GqaPlugin{{$.PluginCode}}{{.ModelName}}
	if err := gqaModel.RequestShouldBindJSON(c, &toEdit{{.ModelName}}); err != nil {
		return
	}
	toEdit{{.ModelName}}.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.Edit{{.ModelName}}(toEdit{{.ModelName}}, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑{{.ModelName}}失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑{{.ModelName}}失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑{{.ModelName}}成功！", c)
	}
}

func Add{{.ModelName}}(c *gin.Context) {
	var toAdd{{.ModelName}} model.RequestAdd{{.ModelName}}
	if err := gqaModel.RequestShouldBindJSON(c, &toAdd{{.ModelName}}); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAdd{{.ModelName}}.Status,
			Sort:      toAdd{{.ModelName}}.Sort,
			Memo:      toAdd{{.ModelName}}.Memo,
		},
	}
	add{{.ModelName}} := &model.GqaPlugin{{$.PluginCode}}{{.ModelName}}{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		{{ $ModelName := .ModelName }}
        {{ range .ColumnList }}
        {{.ColumnName}}:                             toAdd{{$ModelName}}.{{.ColumnName}},
        {{ end }}
	}
	if err := privateservice.Add{{.ModelName}}(*add{{.ModelName}}, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加{{.ModelName}}失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加{{.ModelName}}失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加{{.ModelName}}成功！", c)
	}
}

func Delete{{.ModelName}}ById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.Delete{{.ModelName}}ById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除{{.ModelName}}失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除{{.ModelName}}失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除{{.ModelName}}成功！", c)
	}
}

func Query{{.ModelName}}ById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, record := privateservice.Query{{.ModelName}}ById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找{{.ModelName}}失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找{{.ModelName}}失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": record}, "查找{{.ModelName}}成功！", c)
	}
}
{{ end }}