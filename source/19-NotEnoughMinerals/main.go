package main

type State struct {
	oreBots      int
	clayBots     int
	obsidianBots int
	geodeBots    int
	ore          int
	clay         int
	obsidian     int
	geodes       int
	timeElapsed  int
}

func main() {
}

func advanceGamestate(state State, nodes map[State]any) {
	var newStates []State

	//OreBot
	oreBotState := state
	buildOreBot := false
	if oreBotState.ore >= 4 {
		oreBotState.ore = oreBotState.ore - 4
		buildOreBot = true
	}
	oreBotState = produce(oreBotState)
	if buildOreBot {
		oreBotState.oreBots = oreBotState.oreBots + 1
	}
	oreBotState.timeElapsed = oreBotState.timeElapsed + 1
	newStates = append(newStates, oreBotState)

	//ClayBot
	clayBotState := state
	buildClayBot := false
	if clayBotState.ore >= 2 {
		clayBotState.ore = clayBotState.ore - 2
		buildClayBot = true
	}
	clayBotState = produce(clayBotState)
	if buildClayBot {
		clayBotState.clayBots = clayBotState.clayBots + 1
	}
	clayBotState.timeElapsed = clayBotState.timeElapsed + 1
	newStates = append(newStates, clayBotState)

	//ObsidianBot
	obsidianBotState := state
	buildObsidianBot := false
	if obsidianBotState.ore >= 3 && obsidianBotState.clay >= 14 {
		obsidianBotState.ore = obsidianBotState.ore - 3
		obsidianBotState.obsidian = obsidianBotState.obsidian - 14
		buildObsidianBot = true
	}
	obsidianBotState = produce(obsidianBotState)
	if buildObsidianBot {
		obsidianBotState.obsidianBots = obsidianBotState.obsidianBots + 1
	}
	obsidianBotState.timeElapsed = obsidianBotState.timeElapsed + 1
	newStates = append(newStates, obsidianBotState)

	//GeodeBot
	geodeBotState := state
	buildGeodeBot := false
	if geodeBotState.ore >= 2 && geodeBotState.obsidian >= 12 {
		geodeBotState.ore = geodeBotState.ore - 2
		geodeBotState.obsidian = geodeBotState.obsidian - 12
		buildGeodeBot = true
	}
	geodeBotState = produce(geodeBotState)
	if buildGeodeBot {
		geodeBotState.geodeBots = geodeBotState.geodeBots + 1
	}
	geodeBotState.timeElapsed = geodeBotState.timeElapsed + 1
	newStates = append(newStates, geodeBotState)
}

func produce(state State) State {
	state.ore = state.ore + state.oreBots
	state.clay = state.clay + state.clayBots
	state.obsidian = state.obsidian + state.obsidianBots
	state.geodes = state.geodes + state.geodeBots
	return state
}
