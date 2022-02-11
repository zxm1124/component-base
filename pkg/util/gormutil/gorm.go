package gormutil

type OffsetAndLimit struct {
	Offset int
	Limit  int
}

var DefaultLimit = 100

// Unpointer 如果未添加offset和limit，则设置默认值
func Unpointer(offset, limit *int64) *OffsetAndLimit {
	var o, l int = 0, DefaultLimit

	if offset != nil {
		o = int(*offset)
	}

	if limit != nil {
		l = int(*limit)
	}

	return &OffsetAndLimit{
		Offset: o,
		Limit:  l,
	}
}
