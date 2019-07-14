package gogen

import (
	"testing"
)

func Test_StringInSlice(t *testing.T) {
	f := func(str string, sl []string, want bool) {
		t.Helper()
		ok := StringInSlice(str, sl)
		if ok != want {
			t.Fatalf("result mismatch; get: %v , want: %v", ok, want)
		}
	}

	f("apple", []string{}, false)
	f("apple", []string{"pineapple"}, false)
	f("apple", []string{"pineapple", "apple", "onion"}, true)
}

func Test_NumInSlice(t *testing.T) {
	f := func(num int, sl []int, want bool) {
		t.Helper()
		ok := NumInSlice(num, sl)
		if ok != want {
			t.Fatalf("result mismatch; get: %v , want: %v", ok, want)
		}
	}

	f(1, []int{}, false)
	f(1, []int{2}, false)
	f(1, []int{2, 3, 1}, true)
}

func Test_InterfaceInSlice(t *testing.T) {
	f := func(inf interface{}, sl []interface{}, want bool) {
		t.Helper()
		ok := InterfaceInSlice(inf, sl)
		if ok != want {
			t.Fatalf("result mismatch; get: %v , want: %v", ok, want)
		}
	}

	f(1, []interface{}{}, false)
	f(1, []interface{}{2}, false)
	f(1, []interface{}{"2", "1"}, false)
	f(1, []interface{}{1, 2, 3}, true)
	f("1", []interface{}{"2", "1"}, true)
}

func Test_CompareStringSlices(t *testing.T) {
	f := func(sl1 []string, sl2 []string, want bool) {
		t.Helper()
		ok := CompareStringSlices(sl1, sl2)
		if ok != want {
			t.Fatalf("result mismatch; get: %v , want: %v", ok, want)
		}
	}

	f([]string{"one", "two"}, []string{"five"}, false)
	f([]string{"one", "two"}, []string{"one", "two"}, true)
	f([]string{"one", "two"}, []string{"two", "one"}, true)
	f([]string{}, []string{}, true)
	f([]string{"five"}, []string{}, false)
}

func Test_FullCompareStringSlices(t *testing.T) {
	f := func(sl1 []string, sl2 []string, want bool) {
		t.Helper()
		ok := FullCompareStringSlices(sl1, sl2)
		if ok != want {
			t.Fatalf("result mismatch; get: %v , want: %v", ok, want)
		}
	}

	f([]string{"one", "two"}, []string{"five"}, false)
	f([]string{"one", "two"}, []string{"one", "two"}, true)
	f([]string{"one", "two"}, []string{"two", "one"}, false)
	f([]string{}, []string{}, true)
	f([]string{"five"}, []string{}, false)
}

func Test_SliceInStringSlices(t *testing.T) {
	f := func(sl []string, ssl [][]string, want bool) {
		t.Helper()
		ok := SliceInStringSlices(sl, ssl)
		if ok != want {
			t.Fatalf("result mismatch; get: %v , want: %v", ok, want)
		}
	}

	f([]string{"one"},
		[][]string{
			[]string{"five"},
		},
		false,
	)
	f([]string{"one"},
		[][]string{
			[]string{"five", "nine"},
			[]string{"one"},
		},
		true,
	)
	f([]string{"one", "two"},
		[][]string{
			[]string{"five", "nine"},
			[]string{"one", "two"}},
		true,
	)
	f([]string{"one", "two", "eleven"},
		[][]string{
			[]string{"five", "nine"},
			[]string{"eleven", "two", "one"},
			[]string{"ten"},
		},
		true,
	)
	f([]string{},
		[][]string{
			[]string{},
			[]string{"eleven", "two", "one"},
		},
		true,
	)
	f([]string{},
		[][]string{},
		false,
	)
}

func Test_IsOdd(t *testing.T) {
	f := func(num int, want bool) {
		t.Helper()
		odd := IsOdd(num)
		if odd != want {
			t.Fatalf("result mismatch; get: %v , want: %v", odd, want)
		}
	}

	f(4, false)
	f(7, true)
	f(0, false)
}
