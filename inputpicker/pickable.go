package inputpicker

// Pickable defines the interface for input pickers.
type Pickable interface {
	AskInput() (string, error)
}
