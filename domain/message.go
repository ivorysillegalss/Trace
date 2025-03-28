package domain

import (
	"encoding/binary"
	"net"
	"sync/atomic"
	"time"
)

var (
	counter uint32 = 0
	lastNs  int64
)

// 一条日志信息由128位组成
// 0号位保留
// 63位纳秒时间戳
// 16位递增序列
// 32位IPV4
// 16位端口
// https://raw.githubusercontent.com/ivorysillegalss/pic-bed/refs/heads/main/%E6%88%AA%E5%B1%8F2025-03-16%2019.32.33.png
type Message struct {
	TimeId   int64
	Sequence uint16
	NextIp   uint32
	NextPort uint16
}

func (m *Message) SetNo() int64 {
	time := time.Now()
	nanoTime := time.UnixNano()
	m.TimeId = nanoTime
	return nanoTime
}

func (m *Message) SetSequence(nowTime int64) uint16 {
	// 判断时间是否延迟 是否需要更新16位递增序列号
	if atomic.LoadInt64(&lastNs) != nowTime {
		atomic.StoreInt64(&lastNs, nowTime)
		atomic.StoreUint32(&counter, 0)
	}
	// 获取当前信息的id
	seq := atomic.AddUint32(&counter, 1) - 1
	v := uint16(seq & 0xFFFF)
	m.Sequence = v
	return v
}

func (m *Message) SetIpPort(ip net.IP) {
	ipv4 := ip.To4()
	ipaddr := binary.BigEndian.Uint32(ipv4)
	m.NextIp = ipaddr
}

func (m *Message) SetPort(port uint16) {
	m.NextPort = port
}
