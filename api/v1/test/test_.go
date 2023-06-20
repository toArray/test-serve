package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/test"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type TestApi struct {
}

var TestNameService = service.ServiceGroupApp.TestServiceGroup.TestService


// CreateTest 创建Test
// @Tags Test
// @Summary 创建Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test true "创建Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TestName/createTest [post]
func (TestNameApi *TestApi) CreateTest(c *gin.Context) {
	var TestName test.Test
	err := c.ShouldBindJSON(&TestName)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    TestName.CreatedBy = utils.GetUserID(c)
	if err := TestNameService.CreateTest(TestName); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTest 删除Test
// @Tags Test
// @Summary 删除Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test true "删除Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TestName/deleteTest [delete]
func (TestNameApi *TestApi) DeleteTest(c *gin.Context) {
	var TestName test.Test
	err := c.ShouldBindJSON(&TestName)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    TestName.DeletedBy = utils.GetUserID(c)
	if err := TestNameService.DeleteTest(TestName); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestByIds 批量删除Test
// @Tags Test
// @Summary 批量删除Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TestName/deleteTestByIds [delete]
func (TestNameApi *TestApi) DeleteTestByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := TestNameService.DeleteTestByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTest 更新Test
// @Tags Test
// @Summary 更新Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Test true "更新Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TestName/updateTest [put]
func (TestNameApi *TestApi) UpdateTest(c *gin.Context) {
	var TestName test.Test
	err := c.ShouldBindJSON(&TestName)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    TestName.UpdatedBy = utils.GetUserID(c)
	if err := TestNameService.UpdateTest(TestName); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTest 用id查询Test
// @Tags Test
// @Summary 用id查询Test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test.Test true "用id查询Test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TestName/findTest [get]
func (TestNameApi *TestApi) FindTest(c *gin.Context) {
	var TestName test.Test
	err := c.ShouldBindQuery(&TestName)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reTestName, err := TestNameService.GetTest(TestName.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTestName": reTestName}, c)
	}
}

// GetTestList 分页获取Test列表
// @Tags Test
// @Summary 分页获取Test列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query testReq.TestSearch true "分页获取Test列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TestName/getTestList [get]
func (TestNameApi *TestApi) GetTestList(c *gin.Context) {
	var pageInfo testReq.TestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := TestNameService.GetTestInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
