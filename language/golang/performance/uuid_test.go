package performance

import (
	"os/exec"
	"testing"

	"github.com/blend/go-sdk/uuid"
)

func BenchmarkUUID(t *testing.B) {
	for range t.N {
		uuid.V4()
		//fmt.Println(id)
	}
}

func BenchmarkUUIDByCmd(t *testing.B) {
	for range t.N {
		exec.Command("uuidgen").Output()
		//fmt.Println(string(bytes.TrimSpace(id)))
	}
}
