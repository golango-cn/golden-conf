package test

import (
	"fmt"
	"github.com/golango.cn/golden-conf/conf-database/model"
	"testing"
)

func TestFinds(t *testing.T) {

	env := new(model.Environment)
	find, err := env.Finds()

	fmt.Println(find, err)

}
