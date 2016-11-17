package amazon

import "testing"

func TestCartClear(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	req := client.CartClear(CartClearParameters{})
	if req == nil {
		t.Error("Expected not nil but got nil")
	}
	Test{client, req.Client}.Compare(t)
}

func TestCartClearBuildQuery(t *testing.T) {
	client, _ := New("AK", "SK", "JP")
	q := client.CartClear(CartClearParameters{}).buildQuery()
	Test{0, len(q)}.Compare(t)
}