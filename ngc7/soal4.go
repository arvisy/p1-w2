package main

func ScheduleTask(task func()) {
	go task()
}
