package aiutility

// Reasoner represents the root of the AI system.
// It contains all the actions and considerations.
type Reasoner struct {
	Actions        []*Action                // The actions that can be performed by the agent.
	Considerations map[string]Consideration // The considerations that can be used to calculate the utility of an action.
}

// NewReasoner returns a new reasoner.
func NewReasoner() *Reasoner {
	return &Reasoner{
		Actions:        make([]*Action, 0),
		Considerations: make(map[string]Consideration),
	}
}

// BestAction returns the action with the highest utility.
func (r *Reasoner) BestAction() *Action {
	var bestAction *Action
	var bestUtility float64
	for _, action := range r.Actions {
		utility := action.Utility(r)
		if utility > bestUtility {
			bestAction = action
			bestUtility = utility
		}
	}
	return bestAction
}

// Action represents an action that can be performed by an agent.
type Action struct {
	Name    string                  // The name of the action.
	Utility func(*Reasoner) float64 // The utility of the action.
	Execute func()                  // The function that executes the action.
}

// Consideration represents a consideration that can be used to calculate the utility of an action.
// TODO: Introduce parameters for considerations.
// - This could be entity a and b for calculating the distance between two entities, etc.
type Consideration func(params any) float64
