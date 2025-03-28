package domain

// 将原有的128位requestId的结构体
// 分两个uint64进行存储
type RequestHead struct {
	part1 uint64
	part2 uint64
}

func NewRequestHead(req RequestId) *RequestHead {
	var part2 uint64
	// 将序列号存储在高16位
	part2 |= uint64(req.Sequence) << 16
	// 中间的32位存储ip
	part2 |= uint64(req.NextIp) << 48
	// 低16位存储端口
	part2 |= uint64(req.NextPort)

	return &RequestHead{
		part1: req.TimeId,
		part2: part2,
	}
}

type Message struct {
	ReqHead       *RequestHead
	ReqInfo       string
	ReqStatusCode int
}

// 接口只有两个 查询数据和新增数据
type MesssageService interface {
	// 可以根据多种查询条件进行查询
	SearchMessage(args ...any)
	// 增加全量数据
	InsertMessage(reqMessage *Message)
}

type MessageRepository interface {
	InsertMessage(reqMessage *Message) bool
}
