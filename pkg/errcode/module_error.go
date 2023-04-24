package errcode

var (
	ErrorGetTokenFail   = NewError(20010001, "获取Token失败")
	ErrorGetTagListFail = NewError(20010002, "获取标签列表失败")
)
