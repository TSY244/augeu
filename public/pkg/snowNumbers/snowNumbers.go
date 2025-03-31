package snowNumbers

import (
	"augeu/public/util/convert"
	"errors"
	"sync"
	"time"
)

var (
	snow *Snowflake
)

const (
	epoch            int64 = 1577836800000 // 自定义的起始时间戳（2020-01-01 00:00:00 UTC）
	workerIDBits     uint  = 5
	datacenterIDBits       = 5
	sequenceBits           = 12

	maxWorkerID     int64 = -1 ^ (-1 << workerIDBits)
	maxDatacenterID int64 = -1 ^ (-1 << datacenterIDBits)
	sequenceMask    int64 = -1 ^ (-1 << sequenceBits)

	workerIDShift      = sequenceBits + datacenterIDBits
	datacenterIDShift  = sequenceBits
	timestampLeftShift = sequenceBits + datacenterIDBits + workerIDBits

	twepoch = int64(epoch)
)

type Snowflake struct {
	mu           sync.Mutex
	timestamp    int64
	workerID     int64
	datacenterID int64
	sequence     int64
}

func init() {
	snow, _ = NewSnowflake(1, 1)
}

// NewSnowflake returns a new snowflake node that can be used to generate snowflake IDs.
func NewSnowflake(workerID, datacenterID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID out of range")
	}
	if datacenterID < 0 || datacenterID > maxDatacenterID {
		return nil, errors.New("datacenter ID out of range")
	}
	return &Snowflake{
		workerID:     workerID,
		datacenterID: datacenterID,
		timestamp:    0,
		sequence:     0,
	}, nil
}

func (sf *Snowflake) TillNextMillis(lastTimestamp int64) int64 {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	for timestamp <= lastTimestamp {
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return timestamp
}

func (sf *Snowflake) NextID() (int64, error) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	if timestamp < sf.timestamp {
		return 0, errors.New("clock moved backwards. Refusing to generate id")
	}

	if sf.timestamp == timestamp {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			timestamp = sf.TillNextMillis(sf.timestamp)
		}
	} else {
		sf.sequence = 0
	}

	sf.timestamp = timestamp

	id := ((timestamp - twepoch) << timestampLeftShift) |
		(sf.datacenterID << datacenterIDShift) |
		(sf.workerID << workerIDShift) |
		sf.sequence

	return id, nil
}

func GetAnID() (int64, error) {
	return snow.NextID()
}

func GetAnStrID() (string, error) {
	id, err := snow.NextID()
	if err != nil {
		return "", err
	}
	return convert.Int642Str(id), nil
}
