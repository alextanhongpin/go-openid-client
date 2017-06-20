package main

func TestHello(t *testing.T) {
    expect := "go"
	got := "go"

	if got != expect {
		t.Error("Expected %s, got %s", expect, got)
	}
}