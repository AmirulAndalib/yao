package core

// ITemplate is the interface for the ITemplate
type ITemplate interface {
	Pages() ([]IPage, error)
	Page(route string) (IPage, error)

	Blocks() ([]IBlock, error)
	Block(name string) (IBlock, error)

	Components() ([]IComponent, error)
	Component(name string) (IComponent, error)

	Assets() []string
	Locales() []SelectOption
	Themes() []SelectOption
}

// IPage is the interface for the page
type IPage interface {
	Load() error

	// Render()

	// Html()
	// Script()
	// Style()
	// Data()

	// Compile()
	// Locale()
}

// IBlock is the interface for the block
type IBlock interface {
	Compile() (string, error)
	Load() error
}

// IComponent is the interface for the component
type IComponent interface {
	Compile() (string, error)
	Load() error
}