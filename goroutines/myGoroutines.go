package goroutines

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

var dbData = []string{"hi", "hello", "greetings", "hola"}
var slice = []string{}
var gw = sync.WaitGroup{}

// var m=sync.Mutex{}
var m = sync.RWMutex{}

func MyGoroutines() {

	t1 := time.Now()
	// for i := 0; i < len(dbData); i++ {
	for i := 0; i < 10000; i++ {
		gw.Add(1)
		// go dbCall(i)
		count()
		//  dbCall(i)
	}
	gw.Wait()

	fmt.Printf("\n Total execution time :%v\n", time.Since(t1))
	fmt.Println("The results for this operation are", slice)

}
func dbCall(i int) {
	// var delay float32 = rand.Float32() * 2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	// m.Lock()
	// // fmt.Println("The result from database is: ", dbData[i])
	// m.Unlock()

	// slice = append(slice, dbData[i])
	gw.Done()
}

func save(res string) {
	m.Lock()
	slice = append(slice, res)
	m.Unlock()

}

func log() {
	m.RLock()
	fmt.Printf("\n Results %v \n", slice)
	m.RUnlock()

}

func count() {
	var res int

	for i := 0; i < 100000; i++ {
		res = res + i

	}

	gw.Done()
}
