package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"

	BPTree "github.com/WatchJani/BPlustTree"
)

const (
	KeySize   uint = 16
	DataBlock uint = 200
)

func main() {

	fmt.Println(LoadDataStack()[199_999_999])
	// stack := stack.NewStack(8)
	// memTable, err := stack.Pop()
	// if err != nil {
	// 	log.Println(err)
	// }

	// memTableBuilder := NewBuilder(memTable)

	// mTable := NewMemTable()

	// dataBuffer := make([]byte, DataBlock)
	// for range 5 {
	// 	newDataBlock := KeyGenerator(DataBlock, dataBuffer)
	// 	mTable.Inset(newDataBlock, memTableBuilder)
	// }
	// fmt.Println(memTableBuilder.store[:1000])
}

type MemTable struct {
	bptree *BPTree.Tree[string, int]
}

func NewMemTable() MemTable {
	return MemTable{
		bptree: BPTree.New[string, int](50),
	}
}

type Builder struct {
	store   []byte
	counter int
}

func NewBuilder(store []byte) *Builder {
	return &Builder{
		store: store,
	}
}

func (b *Builder) Len(len int) (int, bool) {
	return b.counter, b.counter+len < cap(b.store)
}

func (b *Builder) Insert(data string) {
	copy(b.store[b.counter:], []byte(data))
	b.counter += len(data)
}

func (b *Builder) Clear() {
	b.counter = 0
}

func (m *MemTable) Inset(data string, memTable *Builder) {
	key := data[:KeySize]

	offset, ok := memTable.Len(int(DataBlock))
	if !ok {
		fmt.Println(m.bptree.GetRoot())
		//send for next computing
		m.bptree = BPTree.New[string, int](50)
	
		memTable.Clear()
	}

	memTable.Insert(data)
	m.bptree.Insert(key, offset)
}

func KeyGenerator(size uint, buff []byte) string {
	if size < KeySize {
		size = KeySize
	}

	for index := range KeySize {
		buff[index] = byte(rand.Intn('Z'-'A') + 'A')
	}

	for index := KeySize; index < size; index++ {
		buff[index] = byte(rand.Intn('z'-'a') + 'a')
	}

	return string(buff)
}

func FillFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	buffer := make([]byte, DataBlock)

	for range 1_000_000 {
		_, err := writer.WriteString(KeyGenerator(DataBlock, buffer))
		if err != nil {
			fmt.Println("Writing data error:", err)
			return
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Flush err:", err)
		return
	}

	fmt.Println("Done")
}

func LoadDataStack() []byte {
	load := make([]byte, 200_000_000)

	file, err := os.Open("allData.bin")
	if err != nil {
		log.Println(err)
	}

	_, err = file.Read(load)
	if err != nil {
		log.Println(err)
	}

	// fmt.Println("byte loaded | ", n)

	return load
}
