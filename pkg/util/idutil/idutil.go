package idutil

import (
	"github.com/zxm1124/component-base/pkg/util/iputil"

	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids"
)

const (
	//Alphabet62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		ip := iputil.GetLocalIP()

		return uint16([]byte(ip)[2])<<8 + uint16([]byte(ip)[3]), nil
	}

	sf = sonyflake.NewSonyflake(st)
}

// GetDistributeID 获得分布式ID
func GetDistributeID() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}

	return id
}

// GetInstanceID 获取资源ID
func GetInstanceID(id uint64, prefix string) string {
	hd := hashids.NewData()
	hd.Alphabet = Alphabet36
	hd.Salt = "!3asd@asd"
	hd.MinLength = 6

	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}

	str, err := h.Encode([]int{int(id)})
	if err != nil {
		panic(err)
	}

	return prefix + str
}
