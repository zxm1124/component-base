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
	GetPage() int
	SetPage(int)
	GetPageSize() int
	SetPageSize(int)
	GetTotal() int64
	SetTotal(count int64)
}

func (meta *ObjectMeta) GetID() uint64            { return meta.ID }
func (meta *ObjectMeta) SetID(id uint64)          { meta.ID = id }
func (meta *ObjectMeta) GetName() string          { return meta.Name }
func (meta *ObjectMeta) SetName(name string)      { meta.Name = name }
func (meta *ObjectMeta) GetCreatedAt() time.Time  { return meta.CreatedAt }
func (meta *ObjectMeta) SetCreatedAt(t time.Time) { meta.CreatedAt = t }
func (meta *ObjectMeta) GetUpdatedAt() time.Time  { return meta.UpdatedAt }
func (meta *ObjectMeta) SetUpdatedAt(t time.Time) { meta.UpdatedAt = t }

func (list *ListMeta) GetPage() int         { return list.Page }
func (list *ListMeta) SetPage(page int)     { list.Page = page }
func (list *ListMeta) GetPageSize() int     { return list.PageSize }
func (list *ListMeta) SetPageSize(size int) { list.PageSize = size }
func (list *ListMeta) GetTotal() int64      { return list.Total }
func (List *ListMeta) SetTotal(total int64) { List.Total = total }
