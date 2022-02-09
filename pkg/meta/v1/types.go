package v1

import (
	"encoding/json"
	"time"
)

// Extend 可存储对象额外的字段数据
type Extend map[string]interface{}

func (ext Extend) String() string {
	data, _ := json.Marshal(ext)
	return string(data)
}

// Merge 将原先存储的string数据添加到Extend中
func (ext Extend) Merge(extendShow string) Extend {
	var extend Extend
	_ = json.Unmarshal([]byte(extendShow), &extend)
	for k, v := range extend {
		if _, ok := ext[k]; !ok {
			ext[k] = v
		}
	}

	return ext
}

type ObjectMeta struct {
	// ID 数据库唯一字段
	ID uint64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`

	// InstanceID 资源唯一字段
	InstanceID string `json:"instanceID,omitempty" gorm:"unique;column:instanceID;type:varchar(32);not null"`

	// Name 资源名
	Name string `json:"name,omitempty" gorm:"colum:name;type:varchar(64);not null" validate:"name"`

	// Extend 额外字段
	Extend Extend `json:"extend,omitempty" gorm:"-" validate:"omitempty"`

	// ExtendShow 额外字段json数据
	ExtendShow string `json:"-" gorm:"extendShow" validate:"omitempty"`

	// CreatedAt
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`

	// UpdatedAt
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"updatedAt"`
}

type ListMeta struct {
	Count int64
}
