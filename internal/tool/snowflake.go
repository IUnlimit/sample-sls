package tool

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	epoch             = int64(1577808000000)                            // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits     = uint(41)                                        // 时间戳占用位数
	datacenterIdBits  = uint(3)                                         // 数据中心id所占位数
	crawlerIdBits     = uint(7)                                         // 机器id所占位数
	sequenceBits      = uint(12)                                        // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))               // 时间戳最大值
	datacenterIdMax   = int64(-1 ^ (-1 << datacenterIdBits))            // 支持的最大数据中心id数量
	crawlerIdMax      = int64(-1 ^ (-1 << crawlerIdBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))                // 支持的最大序列id数量
	crawlerIdShift    = sequenceBits                                    // 机器id左移位数
	datacenterIdShift = sequenceBits + crawlerIdBits                    // 数据中心id左移位数
	timestampShift    = sequenceBits + crawlerIdBits + datacenterIdBits // 时间戳左移位数
)

// Snowflake 雪花算法
type Snowflake struct {
	sync.Mutex         // 锁
	timestamp    int64 // 时间戳 ，毫秒
	datacenterId int64 // 数据中心id
	crawlerId    int64 // 工作节点id
	sequence     int64 // 序列号
}

// NewSnowflake 创建一个 Snowflake 实例
func NewSnowflake(datacenterId int64, crawlerId int64) (*Snowflake, error) {
	if datacenterId < 0 || datacenterId > datacenterIdMax {
		return nil, fmt.Errorf("datacenter Id must be between 0 and %d", datacenterIdMax)
	}
	if crawlerId < 0 || crawlerId > crawlerIdMax {
		return nil, fmt.Errorf("crawler Id must be between 0 and %d", crawlerIdMax)
	}
	return &Snowflake{
		datacenterId: datacenterId,
		crawlerId:    crawlerId,
	}, nil
}

func (s *Snowflake) NextVal() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1000000 // 转毫秒
	if s.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
		s.timestamp = now
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		log.Errorf("Epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	r := (t << timestampShift) | (s.datacenterId << datacenterIdShift) | (s.crawlerId << crawlerIdShift) | (s.sequence)
	return r
}
