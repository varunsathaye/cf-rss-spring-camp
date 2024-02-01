package worker

import "mongo-with-golang/stores"

func PerformWork(m *stores.MongoStore){
	for{
		m.OpenConnectionWithMongoDB()
		
	}
}