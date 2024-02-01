package main

import (
	"mongo-with-golang/stores"
	"mongo-with-golang/worker"
)

func main() {
	var m *stores.MongoStore
	//var cfCLient *cfapi.CodeForcesClient
	worker.PerformWork(m)

}
