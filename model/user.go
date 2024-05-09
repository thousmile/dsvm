package model

import "time"

type User struct {
	CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
	UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	Updated   int64     `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created   int64     `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
}
