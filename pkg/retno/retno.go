package retno

type RetNo int64

//go:generate stringer -type RetNo -linecomment
const (
	OK             RetNo = iota // OK
	INVALID_DATA                // 请求参数不正确
	UNAUTHORIZED                // 无权限
	INTERNAL_ERROR              // 服务器内部错误
	NOT_FOUND                   // 无该资源
)
