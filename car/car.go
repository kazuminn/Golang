package main

import (
	"fmt"
	"math"
	"strconv"
)

type Car struct {
	id    int
	brand string
	cc    float64
	now   int
}

type CarInterface interface {
	run()
}

func (c *Car) run(cc float64) {
	c.now = c.now + int(math.Sqrt(cc))
}
func main() {

	cars := []*Car{
		{1, "toyota", 100, 0},
		{2, "suzuki", 200, 0},
	}

	for i := 1; i < 10; i++ {
		for _, a := range cars {
			a.run(a.cc)
			fmt.Println("car " + a.brand + "is " + strconv.Itoa(a.now) + "m")
		}
	}
}

/* output
car toyotais 10m
car suzukiis 14m
car toyotais 20m
car suzukiis 28m
car toyotais 30m
car suzukiis 42m
car toyotais 40m
car suzukiis 56m
car toyotais 50m
car suzukiis 70m
car toyotais 60m
car suzukiis 84m
car toyotais 70m
car suzukiis 98m
car toyotais 80m
car suzukiis 112m
car toyotais 90m
car suzukiis 126m
*/
