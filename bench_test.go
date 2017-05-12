package main

import (
	"testing"

	"github.com/leprechau/uuid-test/uuidf"

	guuid "github.com/google/uuid"
	muuid "github.com/m4rw3r/uuid"
	fuuid "github.com/rogpeppe/fastuuid"
	suuid "github.com/satori/go.uuid"
)

func BenchmarkV4Google(b *testing.B) {
	for i := 0; i < b.N; i++ {
		guuid.NewRandom()
	}
}

func BenchmarkV4Satori(b *testing.B) {
	for i := 0; i < b.N; i++ {
		suuid.NewV4()
	}
}

func BenchmarkV4m4rw3r(b *testing.B) {
	for i := 0; i < b.N; i++ {
		muuid.V4()
	}
}

func BenchmarkFastUUIDRaw(b *testing.B) {
	g := fuuid.MustNewGenerator()
	for i := 0; i < b.N; i++ {
		g.Next()
	}
}

func BenchmarkFastUUIDFormatted(b *testing.B) {
	g := fuuid.MustNewGenerator()
	for i := 0; i < b.N; i++ {
		uuidf.ToString(g.Next())
	}
}
