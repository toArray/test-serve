package test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    testReq "github.com/flipped-aurora/gin-vue-admin/server/model/test/request"
    "gorm.io/gorm"
)

type TestService struct {
}

// CreateTest 创建Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestNameService *TestService) CreateTest(TestName test.Test) (err error) {
	err = global.GVA_DB.Create(&TestName).Error
	return err
}

// DeleteTest 删除Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestNameService *TestService)DeleteTest(TestName test.Test) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&test.Test{}).Where("id = ?", TestName.ID).Update("deleted_by", TestName.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&TestName).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteTestByIds 批量删除Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestNameService *TestService)DeleteTestByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&test.Test{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&test.Test{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateTest 更新Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestNameService *TestService)UpdateTest(TestName test.Test) (err error) {
	err = global.GVA_DB.Save(&TestName).Error
	return err
}

// GetTest 根据id获取Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestNameService *TestService)GetTest(id uint) (TestName test.Test, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&TestName).Error
	return
}

// GetTestInfoList 分页获取Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestNameService *TestService)GetTestInfoList(info testReq.TestSearch) (list []test.Test, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&test.Test{})
    var TestNames []test.Test
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Aa != "" {
        db = db.Where("aa LIKE ?","%"+ info.Aa+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&TestNames).Error
	return  TestNames, total, err
}
