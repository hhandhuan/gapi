package utils

import "testing"

func TestResolvePathFileName(t *testing.T) {
	path := "home/app/main.go"
	if ResolvePathFileName(path) != "main.go" {
		t.Fatal("resolve error")
	}
	t.Log("is ok")
}
