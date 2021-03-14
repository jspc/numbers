package numbers

import (
	"testing"
)

func TestItoa(t *testing.T) {
	for _, test := range []struct {
		arg    int
		expect string
	}{
		{0, "공"},
		{1, "일"},
		{10, "십"},
		{100, "백"},
		{1000, "천"},
		{10000, "만"},
		{100000, "십만"},

		// Slightly more complex numbers
		{12345611, "천 이백 삼십 사만 오천 육백 십 일"},
	} {
		t.Run("", func(t *testing.T) {
			got := Itoa(test.arg)
			if test.expect != got {
				t.Errorf("expected %q, received %q", test.expect, got)
			}
		})
	}
}
