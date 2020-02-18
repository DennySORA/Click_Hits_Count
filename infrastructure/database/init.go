package database

func InitializationDatabase(stop chan int) {
	go createHitsContext(stop)
}
