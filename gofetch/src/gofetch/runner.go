package gofetch

import (
	"time"
)

func Init() {
	db.session.AutoMigrate(&Forecast{})
}

func Run() {
	defer rerunOnPanic()
	forecast := Parse(Fetch())
	db.session.FirstOrCreate(&forecast)
	next_update := forecast.NextUpdate
	duration := next_update.Sub(time.Now().Add(time.Hour * 2))
	if duration.Minutes() < 5 {
		duration = 5 * time.Minute
	}
	print("Next run in", duration, next_update)
	<-time.After(duration)
	Run()
}

func rerunOnPanic() {
	if r := recover(); r != nil {
		print("Recovering from panic, rerunning in 1 min")
		<-time.After(time.Minute)
		Run()
	}
}
