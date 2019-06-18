package main

import channelsfunn "channelsfunn/lib"

func main() {
	processor := channelsfunn.GetProcessor("electricity-consumption-by-sectors.csv", "elo.db")
	err := processor.Process()
	if err != nil {
		panic(err)
	}
}
