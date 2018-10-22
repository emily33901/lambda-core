package engine

import (
	"github.com/galaco/Gource-Engine/engine/config"
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/lib/gameinfo"
	"runtime"
	"time"
)

// Game engine
// Only 1 can be initialised
type Engine struct {
	EventManager    event.Manager
	Managers        []core.IManager
	running         bool
	simulationSpeed float64
}

// Initialise the engine, and attached managers
func (engine *Engine) Initialise() {
	// Load engine configuration
	engine.loadConfig()

	// Derive and register game resource paths
	filesystem.RegisterGameResourcePaths(config.Get().GameDirectory, gameinfo.Get())
}

// Run the engine
func (engine *Engine) Run() {
	engine.running = true

	// Begin the event manager thread in the background
	event.GetEventManager().RunConcurrent()
	// Anything that runs concurrently can start now
	for _, manager := range engine.Managers {
		manager.RunConcurrent()
	}

	dt := 0.0
	startingTime := time.Now().UTC()

	for engine.running == true {
		for _, manager := range engine.Managers {
			manager.Update(dt)
		}

		for _, manager := range engine.Managers {
			manager.PostUpdate()
		}

		dt = (float64(time.Now().UTC().Sub(startingTime).Nanoseconds()/1000000) / 1000) * engine.simulationSpeed
		startingTime = time.Now().UTC()
	}

	for _, manager := range engine.Managers {
		manager.Unregister()
	}
}

// Add a new manager to the engine
func (engine *Engine) AddManager(manager core.IManager) {
	engine.Managers = append(engine.Managers, manager)
	manager.Register()
}

func (engine *Engine) Close() {
	engine.running = false
}

func (engine *Engine) SetSimulationSpeed(multiplier float64) {
	if multiplier < 0 {
		return
	}
	engine.simulationSpeed = multiplier
}

func (engine *Engine) loadConfig() {
	cfg, err := config.Load()
	if err != nil {
		debug.Log(err)
	}
	gameinfo.LoadConfig(cfg.GameDirectory)
}

func NewEngine() *Engine {
	runtime.LockOSThread()
	return &Engine{
		simulationSpeed: 1.0,
	}
}
