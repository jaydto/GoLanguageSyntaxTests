package main

import (
	"package1/channelsgo"
	"package1/genericsgo"
	"package1/goroutines"
	"package1/pointersgo"
	"package1/runesgo"
	"package1/stringsgo"
	"package1/structsgo"
)

// "package1/arraysgo"
// "package1/arraysgo"
// "package1/mapsgo"
// "package1/speedtestgo"
// "package1/variablesgo"

// import "package1/variablesgo"

func main() {

	// arraysgo.ArrayTests()
	// mapsgo.MyMaps()
	// arraysgo.Looping2()
	// speedtestgo.SpeedTest()
	runesgo.RunesGo()
	stringsgo.StringManipulations()

	structsgo.Mystructs()

	pointersgo.MyPointers()

	goroutines.MyGoroutines()
	channelsgo.MyChannels()
	genericsgo.MyGenerics()

	// variablesgo.VariablesTest()

}
