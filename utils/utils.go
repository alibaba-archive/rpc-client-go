package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"hash"
	rand2 "math/rand"
	"net/url"
	"time"
)

type UUID [16]byte

const numBytes = "1234567890"

func GetUUID() (uuidHex string) {
	uuid := NewUUID()
	uuidHex = hex.EncodeToString(uuid[:])
	return
}

func GetTimeInFormatISO8601() (timeStr string) {
	gmt := time.FixedZone("GMT", 0)

	return time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numBytes[rand2.Intn(len(numBytes))]
	}
	return string(b)
}

func NewUUID() UUID {
	ns := UUID{}
	safeRandom(ns[:])
	u := newFromHash(md5.New(), ns, RandStringBytes(16))
	u[6] = (u[6] & 0x0f) | (byte(2) << 4)
	u[8] = (u[8]&(0xff>>2) | (0x02 << 6))

	return u
}

func newFromHash(h hash.Hash, ns UUID, name string) UUID {
	u := UUID{}
	h.Write(ns[:])
	h.Write([]byte(name))
	copy(u[:], h.Sum(nil))

	return u
}

func safeRandom(dest []byte) {
	if _, err := rand.Read(dest); err != nil {
		panic(err)
	}
}

func GetUrlFormedMap(source map[string]string) (urlEncoded string) {
	urlEncoder := url.Values{}
	for key, value := range source {
		urlEncoder.Add(key, value)
	}
	urlEncoded = urlEncoder.Encode()
	return
}

func (u UUID) String() string {
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}
