package repo

import (
	"dao-exchange/pkg/myerr"
	"errors"
	"testing"
)

var (
	dao *Dao
	err error
)

func TestQueryCurrentHeight(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}
	height, err := dao.QueryCurrentHeightWithTaskName("polygon:80001")
	if err != nil {
		if errors.Is(err, myerr.ErrRecordNotFound) {
			t.Log("record not found")
			return
		}
		t.Log(err)
		t.FailNow()
	}
	t.Log(height)
}
