package main

import "fmt"

type RailroadWideChecker interface {
    CheckRailsWidth() int
}

type RailRoad struct {
	Width int
}

func (r *RailRoad) IsCorrectSizeTrain(s RailroadWideChecker) bool {
	return s.CheckRailsWidth() == r.Width
}

type Train struct {
	TrainWidth int
}

func (t Train) CheckRailsWidth() int {
	return t.TrainWidth
}


func main() {
	railroad := RailRoad{Width: 10}

	passangerTrain := Train{TrainWidth: 10}
	cargoTrain := Train{TrainWidth: 15}

	passangerTrainCheck := railroad.IsCorrectSizeTrain(passangerTrain)
	cargoTrainCheck := railroad.IsCorrectSizeTrain(cargoTrain)

	fmt.Printf("Can passenger train pass? %b\n", passangerTrainCheck)
    fmt.Printf("Can cargo train pass? %b\n", cargoTrainCheck)
}