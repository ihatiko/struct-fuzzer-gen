package struct_fuzzer_gen

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/protobuf/proto"
	"testing"
)

const (
	defaultSet = 1000
)

type Options struct {
	Count int
}

func GetFuzzingSet[T any](f *testing.F, opt ...Options) {
	set := defaultSet
	if len(opt) > 0 && opt[0].Count > 0 {
		set = opt[0].Count
	}
	for i := 0; i < set; i++ {
		a := new(T)
		gofakeit.Seed(int64(i))
		gofakeit.Struct(a)
		data, _ := json.Marshal(a)
		f.Add(data)
	}
}

func GetStruct[T any](byte []byte) (*T, error) {
	t := new(T)
	err := json.Unmarshal(byte, t)
	return t, err
}

func GetFuzzingProtoSet[T proto.Message](f *testing.F, opt ...Options) {
	set := defaultSet
	if len(opt) > 0 && opt[0].Count > 0 {
		set = opt[0].Count
	}
	for i := 0; i < set; i++ {
		a := new(T)
		gofakeit.Seed(int64(i))
		gofakeit.Struct(a)
		data, _ := json.Marshal(a)
		f.Add(data)
	}
}
