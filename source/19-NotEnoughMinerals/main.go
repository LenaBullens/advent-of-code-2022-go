package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

//===Blueprint===

type Blueprint struct {
	id                   int
	oreBotCost           int
	clayBotCost          int
	obsidianBotOreCost   int
	obsidianBotClayCost  int
	geodeBotOreCost      int
	geodeBotObsidianCost int
	highestOreCost       int
}

func createBlueprint(input string) Blueprint {
	splitString1 := strings.Split(input, ": ")
	idSection := splitString1[0]
	costSection := splitString1[1]
	id, error := strconv.Atoi(idSection[10:])
	if error != nil {
		log.Fatal(error)
	}

	splitString2 := strings.Split(costSection, ". ")
	oreBotSection := splitString2[0]
	oreBotSplit := strings.Split(oreBotSection, " ")
	clayBotSection := splitString2[1]
	clayBotSplit := strings.Split(clayBotSection, " ")
	obsidianBotSection := splitString2[2]
	obsidianBotSplit := strings.Split(obsidianBotSection, " ")
	geodeBotSection := splitString2[3]
	geodeBotSplit := strings.Split(geodeBotSection, " ")

	oreBotCost, error := strconv.Atoi(oreBotSplit[4])
	if error != nil {
		log.Fatal(error)
	}
	clayBotCost, error := strconv.Atoi(clayBotSplit[4])
	if error != nil {
		log.Fatal(error)
	}
	obsidianBotOreCost, error := strconv.Atoi(obsidianBotSplit[4])
	if error != nil {
		log.Fatal(error)
	}
	obsidianBotClayCost, error := strconv.Atoi(obsidianBotSplit[7])
	if error != nil {
		log.Fatal(error)
	}
	geodeBotOreCost, error := strconv.Atoi(geodeBotSplit[4])
	if error != nil {
		log.Fatal(error)
	}
	geodeBotObsidianCost, error := strconv.Atoi(geodeBotSplit[7])
	if error != nil {
		log.Fatal(error)
	}

	oreCosts := []int{oreBotCost, clayBotCost, obsidianBotOreCost, geodeBotOreCost}
	sort.Ints(oreCosts)
	highestOreCost := oreCosts[len(oreCosts)-1]

	return Blueprint{id: id, oreBotCost: oreBotCost, clayBotCost: clayBotCost, obsidianBotOreCost: obsidianBotOreCost, obsidianBotClayCost: obsidianBotClayCost, geodeBotOreCost: geodeBotOreCost, geodeBotObsidianCost: geodeBotObsidianCost, highestOreCost: highestOreCost}
}

//===State===

type State struct {
	oreBots      int
	ore          int
	clayBots     int
	clay         int
	obsidianBots int
	obsidian     int
	geodeBots    int
	geodes       int
	minutes      int
}

func createState(oreBots int, ore int, clayBots int, clay int, obsidianBots int, obsidian int, geodeBots int, geodes int, minutes int) State {
	return State{oreBots: oreBots, ore: ore, clayBots: clayBots, clay: clay, obsidianBots: obsidianBots, obsidian: obsidian, geodeBots: geodeBots, geodes: geodes, minutes: minutes}
}

func createInitialState() State {
	return createState(1, 0, 0, 0, 0, 0, 0, 0, 0)
}

func advanceTime(state State) State {
	return createState(state.oreBots, state.ore+state.oreBots, state.clayBots, state.clay+state.clayBots, state.obsidianBots, state.obsidian+state.obsidianBots, state.geodeBots, state.geodes+state.geodeBots, state.minutes+1)
}

func calculateCandidateStates(startingState State, blueprint Blueprint) []State {
	var candidates []State
	//Do nothing
	doNothingState := advanceTime(startingState)
	candidates = append(candidates, doNothingState)

	//Try building geodeBot
	if startingState.ore >= blueprint.geodeBotOreCost && startingState.obsidian >= blueprint.geodeBotObsidianCost {
		geodeBotState := createState(doNothingState.oreBots, doNothingState.ore-blueprint.geodeBotOreCost, doNothingState.clayBots, doNothingState.clay, doNothingState.obsidianBots, doNothingState.obsidian-blueprint.geodeBotObsidianCost, doNothingState.geodeBots+1, doNothingState.geodes, doNothingState.minutes)
		candidates = append(candidates, geodeBotState)
	}

	//Try building obsidianBot
	if startingState.ore >= blueprint.obsidianBotOreCost && startingState.clay >= blueprint.obsidianBotClayCost {
		//Only build obsidianBot if not producing enough obsidian yet
		if startingState.obsidianBots < blueprint.geodeBotObsidianCost {
			obsidianBotState := createState(doNothingState.oreBots, doNothingState.ore-blueprint.obsidianBotOreCost, doNothingState.clayBots, doNothingState.clay-blueprint.obsidianBotClayCost, doNothingState.obsidianBots+1, doNothingState.obsidian, doNothingState.geodeBots, doNothingState.geodes, doNothingState.minutes)
			candidates = append(candidates, obsidianBotState)
		}
	}

	//Try building clayBot
	if startingState.ore >= blueprint.clayBotCost {
		//Only build clayBot if not producing enough clay yet
		if startingState.clayBots < blueprint.obsidianBotClayCost {
			clayBotState := createState(doNothingState.oreBots, doNothingState.ore-blueprint.clayBotCost, doNothingState.clayBots+1, doNothingState.clay, doNothingState.obsidianBots, doNothingState.obsidian, doNothingState.geodeBots, doNothingState.geodes, doNothingState.minutes)
			candidates = append(candidates, clayBotState)
		}
	}

	//Try building oreBot
	if startingState.ore >= blueprint.oreBotCost {
		//Only build oreBot if not producing enough ore yet
		if startingState.oreBots < blueprint.highestOreCost {
			oreBotState := createState(doNothingState.oreBots+1, doNothingState.ore-blueprint.oreBotCost, doNothingState.clayBots, doNothingState.clay, doNothingState.obsidianBots, doNothingState.obsidian, doNothingState.geodeBots, doNothingState.geodes, doNothingState.minutes)
			candidates = append(candidates, oreBotState)
		}
	}

	return candidates
}

//The maximum amount a given state could produce from x minutes onwards is if the state would make a geodeBot every
//single opportunity from now on. (ignoring resources) We check this because if a state can't match up with our current
//maximum, there's no point further going down this branch. (the state's actual potential maximum will be even lower than
//this because we're ignoring the cost for building the geodeBots)

func getTheoreticalMaximum(state State, goal int) int {
	minutesLeft := goal - state.minutes
	//Production is a guassian expansion
	return (minutesLeft*(minutesLeft+1))/2 + state.geodes + minutesLeft*state.geodeBots
}

//===Stack===

type Stack []State

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(value State) {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() (State, error) {
	if stack.IsEmpty() {
		return State{}, errors.New("stack is empty")
	} else {
		index := len(*stack) - 1
		top := (*stack)[index]
		*stack = (*stack)[:index]
		return top, nil
	}
}

func main() {
	solve1()
	solve2()
}

func solve1() {
	lines := helper.ReadLines("input-19.txt")
	var blueprints []Blueprint

	for _, line := range lines {
		blueprints = append(blueprints, createBlueprint(line))
	}

	var totalQualityLevel int

	for _, blueprint := range blueprints {
		initialState := createInitialState()
		highestGeodeCount := 0

		//Depth first search
		var stack Stack
		stack.Push(initialState)
		visited := make(map[State]bool)

		for !stack.IsEmpty() {
			current, error := stack.Pop()
			if error != nil {
				log.Fatal(error)
			}
			visited[current] = true

			if current.minutes < 24 {
				//Determine adjacent states
				candidates := calculateCandidateStates(current, blueprint)

				for _, candidate := range candidates {
					if !visited[candidate] {
						//Naive check to see if state can beat current best
						theoreticalMaximum := getTheoreticalMaximum(candidate, 24)
						if theoreticalMaximum > highestGeodeCount {
							stack.Push(candidate)
						}
					}

				}
			} else {
				if current.geodes > highestGeodeCount {
					highestGeodeCount = current.geodes
				}
			}
		}

		totalQualityLevel = totalQualityLevel + blueprint.id*highestGeodeCount
	}

	fmt.Println(totalQualityLevel)
}

func solve2() {
	lines := helper.ReadLines("input-19.txt")
	var blueprints []Blueprint

	for _, line := range lines {
		blueprints = append(blueprints, createBlueprint(line))
	}

	result := 1

	for i := 0; i < 3; i++ {
		blueprint := blueprints[i]
		initialState := createInitialState()
		highestGeodeCount := 0

		//Depth first search
		var stack Stack
		stack.Push(initialState)
		visited := make(map[State]bool)

		for !stack.IsEmpty() {
			current, error := stack.Pop()
			if error != nil {
				log.Fatal(error)
			}
			visited[current] = true

			if current.minutes < 32 {
				//Determine adjacent states
				candidates := calculateCandidateStates(current, blueprint)

				for _, candidate := range candidates {
					if !visited[candidate] {
						//Naive check to see if state can beat current best
						theoreticalMaximum := getTheoreticalMaximum(candidate, 32)
						if theoreticalMaximum > highestGeodeCount {
							stack.Push(candidate)
						}
					}

				}
			} else {
				if current.geodes > highestGeodeCount {
					highestGeodeCount = current.geodes
				}
			}
		}

		result = result * highestGeodeCount
	}

	fmt.Println(result)
}
