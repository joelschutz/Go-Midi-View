package internal

type ImageConfig struct {
	// Window dimensions
	// defaults to 240x480
	Height int
	Width  int

	Debug bool
}

func NewImageConfig() *ImageConfig {
	return &ImageConfig{
		Height: 240,
		Width:  480,

		Debug: false,
	}
}
