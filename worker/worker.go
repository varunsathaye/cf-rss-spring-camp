package worker

import "mongo-with-golang/stores"

func PerformWork(m *stores.MongoStore){
	for{
		m.OpenConnectionWithMongoDB()
		final_data, _ := (cfclient.RecentActions(5))
		m.StoreRecentActionsInTheDatabase(final_data)
		delay := 5 * time.Minute
		<-time.After(delay)
		
	}
}
