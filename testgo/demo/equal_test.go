package demo

import (
	"testing"
)

//可以通过-run=RegExp来指定运行的测试用例
//还可以通过/来指定要运行的子测试用例
//例如：go test -v -run=Split/simple只会运行simple对应的子测试用例
func TestEqual(t *testing.T) {
	/*
		got := Equal(3, 4)
		want := false

	*/
	type test struct {
		a   int
		b   int
		res bool
	}
	//使用map存储，会能直接反映出哪个有问题
	/*
	   	tests := []test{
	   		{a: 1, b: 1, res: true},
	   		{a: 2, b: 2, res: true},
	     }
	*/
	tests := map[string]test{
		"first":  {a: 1, b: 1, res: true},
		"second": {a: 2, b: 2, res: true},
	}

	//用t.run的时候把每个输入放在匿名函数里，go test的时候-v
	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			got := Equal(v.a, v.b)
			if got != v.res {
				t.Errorf("name is %s, should be %v, but is:%v", name, v.res, got)
			}
		})
	}

	/*
		  if got != want {
				t.Errorf("should be %v, but is:%v", want, got)
		  }
	*/
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(3, 4)
	}
}
