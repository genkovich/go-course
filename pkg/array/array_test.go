package array_test

import (
	"course/pkg/array"
	"testing"
)

func TestHasArrayItem(t *testing.T) {
	type testCase struct {
		name       string
		item       int
		items      []int
		wantResult bool
	}

	testCases := []testCase{
		{"positive1", 1, []int{1, 2, 3, 4, 5}, true},
		{"negative", 6, []int{1, 2, 3, 4, 5}, false},
		{"positive2", 3, []int{1, 2, 3, 4, 5}, true},
		{"positive3", -1, []int{-1, 2, 3, 4, 5}, true},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			haveRes := array.HasArrayItem(tc.item, tc.items)

			if haveRes != tc.wantResult {
				t.Logf("Mismatched result: HasArrayItem() = %v, want %v", haveRes, tc.wantResult)
				t.Fail()
			}
		})
	}
}
