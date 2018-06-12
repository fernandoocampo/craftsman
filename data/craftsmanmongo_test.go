package data_test

import (
	"os"
	"testing"
)

func setup() {
}

func shutdown() {

}

func TestInsert(t *testing.T) {
}

func TestSearchByID(t *testing.T) {
}

func TestUpdateState(t *testing.T) {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
