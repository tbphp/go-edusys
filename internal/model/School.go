package model

import "gorm.io/gorm"

type School struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"type:varchar;comment:学校名称" binding:"required,min=2,max=50" label:"学校名称"`
	OwnerId      uint   `json:"owner_id" gorm:"index:inx_owner;comment:学校创建者"`
	Status       int    `json:"status" gorm:"type:smallint;comment:状态：1.待审核,2.已批准，3.已拒绝"`
	RejectReason string `json:"reject_reason" gorm:"type:varchar;default:'';comment:拒绝原因"`
	Times
	DeletedAt gorm.DeletedAt `json:"-"`
}
