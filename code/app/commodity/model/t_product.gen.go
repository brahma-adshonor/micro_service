// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTProduct = "t_product"

// TProduct mapped from table <t_product>
type TProduct struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 商品id
	Name            string    `gorm:"column:name;not null" json:"name"`                                       // 商品名称
	Price           int32     `gorm:"column:price;not null" json:"price"`                                     // 商品价格, 价格分
	Pic             string    `gorm:"column:pic" json:"pic"`                                                  // 商品图片
	Description     string    `gorm:"column:description;not null" json:"description"`                         // 商品描述
	Sale            int32     `gorm:"column:sale" json:"sale"`                                                // 销量
	DeleteStatus    int32     `gorm:"column:delete_status" json:"delete_status"`                              // 删除状态：0->未删除；1->已删除
	PublishStatus   int32     `gorm:"column:publish_status" json:"publish_status"`                            // 上架状态：0->下架；1->上架
	RecommendStatus int32     `gorm:"column:recommend_status" json:"recommend_status"`                        // 推荐状态；0->不推荐；5->五星推荐
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
}

// TableName TProduct's table name
func (*TProduct) TableName() string {
	return TableNameTProduct
}
