package handlers

import (
	"fmt"
	"mgtu/digital-trace/main-backend-service/internal/password_hasher"
	"testing"
)

func TestXxx(t *testing.T) {
	pass, _ := password_hasher.HashPassword("asd")
	fmt.Println(pass)
}
