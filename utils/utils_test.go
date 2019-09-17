package utils

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUUID(t *testing.T) {
	uuid := NewUUID()
	assert.Equal(t, 16, len(uuid))
	assert.Equal(t, 36, len(uuid.String()))
	uuidString := GetUUID()
	assert.Equal(t, 32, len(uuidString))
}

func TestGetTimeInFormatISO8601(t *testing.T) {
	s := GetTimeInFormatISO8601()
	assert.Equal(t, 20, len(s))
	// 2006-01-02T15:04:05Z
	re := regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z$`)
	assert.True(t, re.MatchString(s))
}

func TestGetUrlFormedMap(t *testing.T) {
	m := make(map[string]string)
	m["key"] = "value"
	s := GetUrlFormedMap(m)
	assert.Equal(t, "key=value", s)
	m["key2"] = "http://domain/?key=value&key2=value2"
	s2 := GetUrlFormedMap(m)
	assert.Equal(t, "key=value&key2=http%3A%2F%2Fdomain%2F%3Fkey%3Dvalue%26key2%3Dvalue2", s2)
}
