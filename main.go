package main

func main() {
	c := loadConfig()
	for _, l := range c.LOG_PATHS {
		readLogs(l)
	}
}
