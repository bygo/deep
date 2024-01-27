package deep_test

import (
	"encoding/json"
	"github.com/bygo/deep"
	"testing"
)

func TestInt(t *testing.T) {
	src := []int{1, 2, 3}
	want := []int{1, 2, 3}
	dst := deep.Copy(src)
	if len(dst) != len(want) {
		t.Errorf("\nreal: length == %d\nwant: length == %d", len(dst), len(want))
	}
	for idx := range dst {
		if dst[idx] != want[idx] {
			t.Errorf("\nreal: %v\nwant: %v", dst[idx], want[idx])
		}
	}
}

func TestAny(t *testing.T) {
	src := []any{1, "b", 1.2, complex(1, 2), any(nil), User{}}
	dst := deep.Copy(src)
	if len(dst) != len(src) {
		t.Errorf("\nreal: length == %d\nwant: length == %d", len(dst), len(src))
	}
	for idx := range dst {
		if dst[idx] != src[idx] {
			t.Errorf("\nreal: %v\nwant: %v", dst[idx], src[idx])
		}
	}
}

type User struct {
	ID       int
	password string
}

type Complex struct {
	Hash1  map[string]string
	Hash2  map[int]User
	Users1 []User
	Users2 []*User
	Users3 *[]User
	Users4 *[]*User
}

func TestBasics(t *testing.T) {
	src := Complex{
		Hash1:  map[string]string{"a": "b"},
		Hash2:  map[int]User{1: {ID: 1}},
		Users1: []User{{ID: 1, password: "123456"}},
		Users2: []*User{{ID: 2}},
		Users3: &[]User{{ID: 3}},
		Users4: &[]*User{{ID: 4}},
	}
	dst := deep.Copy(src)
	dstJson, err := json.Marshal(dst)
	if err != nil {
		panic(err)
	}
	srcJson, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	if len(dstJson) != len(srcJson) {
		t.Errorf("\nreal: length == %d\nwant: length == %d", len(dstJson), len(srcJson))
	}
	if string(dstJson) != string(srcJson) {
		t.Errorf("\nreal: %v\nwant: %v", string(dstJson), string(srcJson))
	}
	if &src.Users1[0] == &dst.Users1[0] {
		t.Errorf("same ptr address %v == %v \n", src.Users1[0], dst.Users1[0])
	}
	if src.Users2[0] == dst.Users2[0] {
		t.Errorf("same ptr address\n")
	}
	if &(*src.Users3)[0] == &(*dst.Users3)[0] {
		t.Errorf("same ptr address\n")
	}
	if (*src.Users4)[0] == (*dst.Users4)[0] {
		t.Errorf("same ptr address\n")
	}
}
