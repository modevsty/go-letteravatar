package contpickerapi

// ShapePicker allows picking a shape.
type ContPickable interface {
	WithName() *ContPicker
	WithBgcolor() *ContPicker
	WithShape() *ContPicker
	WithFormat() *ContPicker
	WithSize() *ContPicker
}
