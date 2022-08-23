package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	trainNo       uint64
	departure     uint64
	arrival       uint64
	cost          float64
	departureTime uint64
	arrivalTime   uint64
}

type Nodes struct {
	Arr         []*Node        // file
	stations    []uint64       // unique stations
	stationsMap map[uint64]int // used to convert station number to iterator (e.g. station 1909 = 1, 1929 = 2)
	paths       []string       // all possible paths after findTime() or findCost()
	pathCosts   []float64      // costs of all paths
}

//var stations = [6]int{1, 2, 3, 4, 5, 6}

func (n *Nodes) makeTimeMatrix() [][]uint64 {
	numOfStations := len(n.stations)
	timeMatrix := make([][]uint64, numOfStations)
	for i := range timeMatrix {
		timeMatrix[i] = make([]uint64, numOfStations)
	}
	for i, e := range n.Arr {
		iTime := n.getTime(i)
		if timeMatrix[n.stationsMap[e.departure]-1][n.stationsMap[e.arrival]-1] == 0 ||
			timeMatrix[n.stationsMap[e.departure]-1][n.stationsMap[e.arrival]-1] > iTime {

			timeMatrix[n.stationsMap[e.departure]-1][n.stationsMap[e.arrival]-1] = iTime
		}

	}
	return timeMatrix
}

func (n *Nodes) makeCostMatrix() [][]float64 {
	numOfStations := len(n.stations)
	costMatrix := make([][]float64, numOfStations)
	for i := range costMatrix {
		costMatrix[i] = make([]float64, numOfStations)
	}
	for _, e := range n.Arr {
		iCost := e.cost
		if costMatrix[n.stationsMap[e.departure]-1][n.stationsMap[e.arrival]-1] == 0 ||
			costMatrix[n.stationsMap[e.departure]-1][n.stationsMap[e.arrival]-1] > iCost {

			costMatrix[n.stationsMap[e.departure]-1][n.stationsMap[e.arrival]-1] = iCost
		}

	}
	return costMatrix
}

func (n *Nodes) uniqueStations() {
	k := 1
	for _, e := range n.Arr {
		if n.stationsMap[e.departure] == 0 {
			n.stationsMap[e.departure] = k
			n.stations = append(n.stations, e.departure)
			k++
		}
		if n.stationsMap[e.arrival] == 0 {
			n.stationsMap[e.arrival] = k
			n.stations = append(n.stations, e.arrival)
			k++
		}
	}
}

func getTimeStr(strings []string) float64 {
	//size will always be 3 (h,m,s)
	dur, _ := time.ParseDuration(strings[0] + "h" + strings[1] + "m" + strings[2] + "s")
	return dur.Seconds()

}

func ParseCSV(name string) *Nodes {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var nodes Nodes
	i := 0
	reader := csv.NewReader(file)
	reader.Comma = ';'
	for {
		record, e := reader.Read()
		if e != nil {
			break
		}
		TrainNo, _ := strconv.Atoi(record[0])
		Departure, _ := strconv.Atoi(record[1])
		Arrival, _ := strconv.Atoi(record[2])
		Cost, _ := strconv.ParseFloat(record[3], 64)
		DepT := getTimeStr(strings.Split(record[4], ":"))
		ArrT := getTimeStr(strings.Split(record[5], ":"))
		temp := &Node{
			trainNo:       uint64(TrainNo),
			departure:     uint64(Departure),
			arrival:       uint64(Arrival),
			cost:          Cost,
			departureTime: uint64(DepT),
			arrivalTime:   uint64(ArrT),
		}
		nodes.Arr = append(nodes.Arr, temp)
		i++
	}
	return &nodes
}

func (n *Nodes) getTime(i int) uint64 {
	h24 := uint64(86400) // 24h in seconds
	if n.Arr[i].departureTime >= n.Arr[i].arrivalTime {
		return n.Arr[i].arrivalTime + (h24 - n.Arr[i].departureTime)
	}
	return n.Arr[i].arrivalTime - n.Arr[i].departureTime
}

func (n *Nodes) start() {
	n.stationsMap = make(map[uint64]int)
	n.uniqueStations()
	visited := make([]int, len(n.stations))
	costs := n.makeCostMatrix()
	n.findCost("", visited, costs, 0, 0, -1)
	fmt.Println("--------COST---------")
	n.ShowMin()
	timeCosts := n.makeTimeMatrix()
	n.findTime("", visited, timeCosts, 0, 0, -1)
	fmt.Println("--------TIME (in seconds)---------")
	n.ShowMin()

}

func (n *Nodes) ShowMin() bool {
	var min float64

	min = n.pathCosts[0]

	// same size
	for i := range n.paths {
		if min > n.pathCosts[i] {
			min = n.pathCosts[i]

		}
		//uncomment this line to see ALL possible paths and their costs
		//fmt.Println(n.paths[i], n.pathCosts[i])
	}
	for i := range n.paths {
		if min == n.pathCosts[i] {
			fmt.Println("path [", n.paths[i], "] cost", n.pathCosts[i])
		}
	}
	n.paths = nil
	n.pathCosts = nil
	return true
}

func (n *Nodes) findCost(path string, visited []int, costs [][]float64, depth int, cost float64, last int) {
	if depth == len(visited) {
		n.paths = append(n.paths, path)
		n.pathCosts = append(n.pathCosts, cost)

		return
	}
	for i := range visited {
		if visited[i] == 0 {

			visited[i] = depth + 1
			if last != -1 {
				if costs[last][i] == 0 {
					visited[i] = 0
					continue
				}
				n.findCost(path+"->"+strconv.Itoa(int(n.stations[i])), visited, costs, depth+1, cost+costs[last][i], i)
			} else {
				n.findCost(path+strconv.Itoa(int(n.stations[i])), visited, costs, depth+1, cost, i)
			}
			visited[i] = 0
		}

	}

}

func (n *Nodes) findTime(path string, visited []int, costs [][]uint64, depth int, cost uint64, last int) {
	//last node, so we need to check
	if depth == len(visited) {
		n.paths = append(n.paths, path)
		n.pathCosts = append(n.pathCosts, float64(cost))
		return
	}
	for i := range visited {
		if visited[i] == 0 {

			visited[i] = depth + 1
			if last != -1 {
				if costs[last][i] == 0 {
					visited[i] = 0
					continue
				}
				n.findTime(path+"->"+strconv.Itoa(int(n.stations[i])), visited, costs, depth+1, cost+costs[last][i], i)
			} else {
				n.findTime(path+strconv.Itoa(int(n.stations[i])), visited, costs, depth+1, cost, i)
			}
			visited[i] = 0
		}

	}

}
