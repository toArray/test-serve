// 自动生成模板Test
package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Test 结构体
type Test struct {
      global.GVA_MODEL
      Aa  string `json:"aa" form:"aa" gorm:"column:aa;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Test 表名
func (Test) TableName() string {
  return "test"
}

