package benchmarks

import "testing"

func BenchmarkReadAllFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readAllFile()
	}
}

func BenchmarkReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readFile()
	}
}

func BenchmarkStringReader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringReader()
	}
}

func BenchmarkByteReader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteReader()
	}
}

func BenchmarkStringBuilderWriteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilderWriteString()
	}
}

func BenchmarkioCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioCopy()
	}
}

func BenchmarkByteWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteWriter()
	}
}
