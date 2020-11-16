package unittest

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	got := Split("babcbde", "b")
// 	want := []string{"", "a", "c", "de"}
// 	if !reflect.DeepEqual(got, want) { //切片比较
// 		t.Errorf("want:%v but got:%v", want, got)
// 	}
// }

//测试一组，且可以指定测试
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case1": {"abcde", "b", []string{"a", "cde"}},
		"case2": {"acdae", "a", []string{"", "cd", "e"}},
	}
	//go test -v 测试执行
	//指定测试执行 go test -run=TestSplit/case1
	//go test -cover -coverprofile=c.out  把覆盖率为多少的测试代码输处到 c.out文件里面
	//html形式查看上面文件 go tool cover -html=c.out
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v but got:%#v\n", tc.want, got)
			}
		})
	}

}

//基准测试 BenchmarkSplit
//go test -bench=Split
//BenchmarkSplit-6 GOMAXROCS值    628574执行次数   194 ns/op 执行的次数的平均值
// //go test -bench=Split -benchmem 可以查看内存情况
// func BenchmarkSplit(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Split("a:b:c", ":")
// 	}
// }
