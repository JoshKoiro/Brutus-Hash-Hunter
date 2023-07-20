package compare

import "testing"

func TestCompareText(t *testing.T) {
	t.Run("comparing equal values", func(t *testing.T) {
		got := CompareText("text value", "text value")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("comparing unequal values", func(t *testing.T) {
		got := CompareText("wrong text", "right text")
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestCompareSHA256(t *testing.T) {
	t.Run("compare equal values", func(t *testing.T) {
		got := CompareSHA256("equal text", "equal text")
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("compare unequal values", func(t *testing.T) {
		got := CompareSHA256("unequal text", "equal text")
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestCompareMD5(t *testing.T) {
	t.Run("compare equal values", func(t *testing.T) {
		got := CompareMD5("equal text", "equal text")
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("compare unequal values", func(t *testing.T) {
		got := CompareMD5("equal text", "unequal text")
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
