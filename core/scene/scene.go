package scene

import (
	"github.com/emily33901/lambda-core/core/model"
)

// Scene
type Scene struct {
	bsp         model.Bsp
	staticProps []model.StaticProp
}

// Bsp
func (s *Scene) Bsp() *model.Bsp {
	return &s.bsp
}

// StaticProps
func (s *Scene) StaticProps() []model.StaticProp {
	return s.staticProps
}

// NewScene
func NewScene(bsp model.Bsp, staticProps []model.StaticProp) *Scene {
	return &Scene{
		bsp:         bsp,
		staticProps: staticProps,
	}
}
