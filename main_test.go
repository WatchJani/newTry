package main

import (
	"log"
	"root/stack"
	"testing"
)

func BenchmarkBufferSpeed(b *testing.B) {
	b.StopTimer()

	stacK := stack.NewStack(8)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data, _ := stacK.Pop()
		stacK.Push(data)
	}
}

func Benchmark(b *testing.B) {
	b.StopTimer()

	stack := stack.NewStack(8)
	memTable, err := stack.Pop()
	if err != nil {
		log.Println(err)
	}

	memTableBuilder := NewBuilder(memTable)

	mTable := NewMemTable()

	data := LoadDataStack()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if i > 999_999 {
			continue
		}

		index := i * int(DataBlock)
		mTable.Inset(string(data[index:index+int(DataBlock)]), memTableBuilder)
	}
}

func BenchmarkGeneratorKey(b *testing.B) {
	b.StopTimer()

	buffer := make([]byte, DataBlock)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		KeyGenerator(DataBlock, buffer)
	}
}
