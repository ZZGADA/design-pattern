package template

// RedisLabelTreeDataTemp labelTree 在redis中的数据结构
type RedisLabelTreeDataTemp struct {
	name      string
	parentIds string
}

func NewRedisMapLabelTreeDataTemp(name, parentIds string) map[string]string {
	return map[string]string{
		name:      name,
		parentIds: parentIds,
	}
}

func (lt *RedisLabelTreeDataTemp) SetName(name string) {
	lt.name = name
}

func (lt *RedisLabelTreeDataTemp) SetParentIds(parentIds string) {
	lt.parentIds = parentIds
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
