package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExist(t *testing.T) {
	tokens := "top_memberfree=1;top_nad=1;top_video_rew=0;"
	find := Exist(&tokens, "top_nad", "1")
	assert.Equal(t, true, find)
}
