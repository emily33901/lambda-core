package loader

import (
	"github.com/emily33901/lambda-core/core/filesystem"
	"github.com/emily33901/lambda-core/core/loader/prop"
	"github.com/emily33901/lambda-core/core/logger"
	"github.com/emily33901/lambda-core/core/model"
	"github.com/emily33901/lambda-core/core/resource"
	"github.com/emily33901/lambda-core/lib/util"
	"github.com/galaco/bsp/primitives/game"
	"strings"
)

// LoadStaticProps GetFile all staticprops referenced in a
// bsp's game lump
func LoadStaticProps(propLump *game.StaticPropLump, fs filesystem.IFileSystem) []model.StaticProp {
	ResourceManager := resource.Manager()
	errorProp, err := prop.LoadProp(ResourceManager.ErrorModelName(), fs)
	// If we have no error model, expect this to be fatal issue
	if errorProp == nil && err != nil {
		logger.Panic(err)
	}

	propPaths := make([]string, 0)
	for _, propEntry := range propLump.PropLumps {
		propPaths = append(propPaths, propLump.DictLump.Name[propEntry.GetPropType()])
	}

	propPaths = util.RemoveDuplicatesFromList(propPaths)
	logger.Notice("Found %d staticprops", len(propPaths))

	numLoaded := 0
	for _, path := range propPaths {
		if !strings.HasSuffix(path, ".mdl") {
			path += ".mdl"
		}
		_, err := prop.LoadProp(path, fs)
		if err != nil {
			continue
		}
		numLoaded++
	}

	logger.Notice("Loaded %d props, failed to load %d props", numLoaded, len(propPaths)-numLoaded)

	staticPropList := make([]model.StaticProp, 0)

	for _, propEntry := range propLump.PropLumps {
		modelName := propLump.DictLump.Name[propEntry.GetPropType()]
		m := ResourceManager.Model(modelName)
		if m != nil {
			staticPropList = append(staticPropList, *model.NewStaticProp(propEntry, &propLump.LeafLump, m))
			continue
		}
		// Model missing, use error model
		m = ResourceManager.Model(ResourceManager.ErrorModelName())
		staticPropList = append(staticPropList, *model.NewStaticProp(propEntry, &propLump.LeafLump, m))
	}

	return staticPropList
}
