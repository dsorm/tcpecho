package common

import (
	"bytes"
	"testing"
	"time"
)

const ServerAddress = ":7000"

func TestMain(m *testing.M) {
	ServerStart(ServerAddress)
	m.Run()
}

func TestServerStart(t *testing.T) {
	expected := []byte("Test!")
	got := SendEcho(expected, ServerAddress, 2*time.Second)

	// if not equal
	if bytes.Compare(expected, got) != 0 {
		t.Fatalf("Error while sending single echo: got %v, expected %v", string(got), string(expected))
	}

	equal, got := SendEchosAndVerify([]byte("Test!"), 16, ServerAddress, 5*time.Second)

	if !equal {
		t.Fatalf("Error while sending multiple echos in single connection: got %v, expected %v", string(got), string(expected))
	}

}

func Benchmark_1x(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendEcho([]byte("Hello world!"), ServerAddress, time.Second)
	}
}

func Benchmark_64x(b *testing.B) {
	running := true
	start := make(chan bool, 1)
	addOne := make(chan bool, 16)
	var i int

	for g := 0; g < 64; g++ {
		go func() {
			<-start
			for running {
				SendEcho([]byte("Hello world!"), ServerAddress, time.Second)
				addOne <- true
			}
		}()
	}

	start <- true
	for i = 0; i < b.N; i++ {
		<-addOne
	}
	running = false
}
