package main

import (
	"log"
	"math"
	"math/rand"

	"github.com/Flokey82/aiutility"
)

type Agent struct {
	Health float64
	Food   float64
	Hunger float64
	Rest   float64
}

func newAgent() *Agent {
	return &Agent{
		Health: 100.0,
		Food:   100.0,
		Hunger: 0.0,
		Rest:   100.0,
	}
}

func main() {
	a := newAgent()

	// Create a new reasoner and set up the consideration functions.
	res := aiutility.NewReasoner()
	res.Considerations["health"] = func(params any) float64 {
		return a.Health / 100.0
	}
	res.Considerations["food"] = func(params any) float64 {
		return a.Food / 100.0
	}
	res.Considerations["rest"] = func(params any) float64 {
		return a.Rest / 100.0
	}
	res.Considerations["hunger"] = func(params any) float64 {
		return a.Hunger / 100.0
	}

	// Heal when our health is low.
	// To determine the utility, we express our health as a number between 0 and 1,
	// invert it (so that 0 is the worst and 1 is the best), and square it (to make
	// it more urgent the lower our health is).
	res.Actions = append(res.Actions, &aiutility.Action{
		Name: "heal",
		Utility: func(r *aiutility.Reasoner) float64 {
			return math.Pow(1-r.Considerations["health"](nil), 2)
		},
		Execute: func() {
			a.Health += 10 // Heal some health.
		},
	})

	// Eat when we're hungry.
	// To determine the utility, we express our hunger as a number between 0 and 1,
	// and cube it (to make it more urgent the hungrier we are).
	// If we don't have any food, the utility is 0.
	res.Actions = append(res.Actions, &aiutility.Action{
		Name: "eat",
		Utility: func(r *aiutility.Reasoner) float64 {
			// Consider if we have any food to eat.
			if r.Considerations["food"](nil) <= 0 {
				return 0
			}

			// Consider how hungry we are.
			return math.Pow(r.Considerations["hunger"](nil), 3)
		},
		Execute: func() {
			a.Food -= 10   // Eat some food.
			a.Hunger = 0.0 // Reset hunger.
		},
	})

	// Sleep when we're tired.
	// To determine the utility, we express our rested state as a number between 0 and 1,
	// invert it (so that 0 is the worst and 1 is the best), and square it (to make
	// it more urgent the less rested we are).
	res.Actions = append(res.Actions, &aiutility.Action{
		Name: "sleep",
		Utility: func(r *aiutility.Reasoner) float64 {
			// Consider if we're tired or rested.
			return math.Pow(1-r.Considerations["rest"](nil), 2)
		},
		Execute: func() {
			a.Rest = 100.0 // Get some rest.
		},
	})

	// Find food when we're out of food.
	// If our inventory is full, the utility is 0.
	res.Actions = append(res.Actions, &aiutility.Action{
		Name: "find food",
		Utility: func(r *aiutility.Reasoner) float64 {
			// If our inventory is full, we don't need any more food.
			if r.Considerations["food"](nil) >= 1.0 {
				return 0
			}
			// Consider how empty our food inventory is.
			return math.Pow(1-r.Considerations["food"](nil), 3)
		},
		Execute: func() {
			a.Food = 100
		},
	})

	// Idle around.
	res.Actions = append(res.Actions, &aiutility.Action{
		Name: "idle",
		Utility: func(r *aiutility.Reasoner) float64 {
			return 0.01
		},
		Execute: func() {
			// Do nothing.
		},
	})

	// Run the simulation for 100 ticks.
	hungerRate := 1.0 // Hunger increases by 1.0 per tick.
	restRate := 1.0   // Rest decreases by 1.0 per tick.
	for i := 0; i < 100; i++ {
		// Calculate hunger.
		a.Hunger += hungerRate
		if a.Hunger > 100 {
			a.Hunger = 100
		}

		// Calculate rest.
		a.Rest -= restRate
		if a.Rest < 0 {
			a.Rest = 0
		}

		// Randomly injure the agent.
		if rand.Float64() < 0.1 {
			a.Health -= 10
		}

		// Log the agent's state.
		log.Printf("health: %.1f, food: %.1f, rest: %.1f, hunger: %.1f", a.Health, a.Food, a.Rest, a.Hunger)

		// Tick the reasoner.
		action := res.BestAction()
		if action != nil {
			log.Println("best action:", action.Name)
			action.Execute()
		}
	}
}
