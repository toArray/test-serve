package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

// InitTestRouter 初始化 Test 路由信息
func (s *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	TestNameRouter := Router.Group("TestName").Use(middleware.OperationRecord())
	TestNameRouterWithoutRecord := Router.Group("TestName")
	var TestNameApi = v1.ApiGroupApp.TestApiGroup.TestApi
	{
		TestNameRouter.POST("createTest", TestNameApi.CreateTest)   // 新建Test
		TestNameRouter.DELETE("deleteTest", TestNameApi.DeleteTest) // 删除Test
		TestNameRouter.DELETE("deleteTestByIds", TestNameApi.DeleteTestByIds) // 批量删除Test
		TestNameRouter.PUT("updateTest", TestNameApi.UpdateTest)    // 更新Test
	}
	{
		TestNameRouterWithoutRecord.GET("findTest", TestNameApi.FindTest)        // 根据ID获取Test
		TestNameRouterWithoutRecord.GET("getTestList", TestNameApi.GetTestList)  // 获取Test列表
	}
}
