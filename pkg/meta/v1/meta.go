package v1

import "time"

// Object 用于存储数据的字段对象
type Object interface {
	GetID() uint64
	SetID(id uint64)
	GetName() string
	SetName(name string)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
}

// ListInterface 用于获取分页查询数据
type ListInterface interface {
	GetTotal() int64
	SetTotal(count int64)
}

// Type 用于定义操作REST资源的类型
type Type interface {
	GetKind() string
	SetKind(kind string)
	GetApiVersion() string
	SetApiVersion(version string)
}

func (meta *ObjectMeta) GetID() uint64            { return meta.ID }
func (meta *ObjectMeta) SetID(id uint64)          { meta.ID = id }
func (meta *ObjectMeta) GetName() string          { return meta.Name }
func (meta *ObjectMeta) SetName(name string)      { meta.Name = name }
func (meta *ObjectMeta) GetCreatedAt() time.Time  { return meta.CreatedAt }
func (meta *ObjectMeta) SetCreatedAt(t time.Time) { meta.CreatedAt = t }
func (meta *ObjectMeta) GetUpdatedAt() time.Time  { return meta.UpdatedAt }
func (meta *ObjectMeta) SetUpdatedAt(t time.Time) { meta.UpdatedAt = t }

func (list *ListMeta) GetTotal() int64      { return list.Total }
func (List *ListMeta) SetTotal(total int64) { List.Total = total }

func (meta *TypeMeta) GetApiVersion() string        { return meta.ApiVersion }
func (meta *TypeMeta) SetApiVersion(version string) { meta.ApiVersion = version }
func (meta *TypeMeta) GetKind() string              { return meta.Kind }
func (meta *TypeMeta) SetKind(kind string)          { meta.Kind = kind }
