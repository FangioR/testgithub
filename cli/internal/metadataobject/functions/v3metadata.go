package functions

import (
	"path/filepath"

	"github.com/hasura/graphql-engine/cli/v2/internal/metadataobject"

	"github.com/hasura/graphql-engine/cli/v2"

	"gopkg.in/yaml.v2"
)

type V3MetadataFunctionConfig struct {
	*FunctionConfig
}

func NewV3MetadataFunctionConfig(ec *cli.ExecutionContext, baseDir string) *V3MetadataFunctionConfig {
	return &V3MetadataFunctionConfig{
		&FunctionConfig{
			MetadataDir: baseDir,
			logger:      ec.Logger,
		},
	}
}
func (t *V3MetadataFunctionConfig) Export(md yaml.MapSlice) (map[string][]byte, metadataobject.ErrParsingMetadataObject) {
	metadataBytes, err := yaml.Marshal(md)
	if err != nil {
		return nil, t.error(err)
	}
	var metadata struct {
		Sources []struct {
			Name      string          `yaml:"name"`
			Functions []yaml.MapSlice `yaml:"functions"`
		} `yaml:"sources"`
	}
	var functions interface{}
	if err := yaml.Unmarshal(metadataBytes, &metadata); err != nil {
		return nil, t.error(err)
	}
	if len(metadata.Sources) > 0 {
		// use tables of first source
		functions = metadata.Sources[0].Functions
	}
	if functions == nil {
		functions = make([]interface{}, 0)
	}
	data, err := yaml.Marshal(functions)
	if err != nil {
		return nil, t.error(err)
	}
	return map[string][]byte{
		filepath.ToSlash(filepath.Join(t.MetadataDir, t.Filename())): data,
	}, nil
}
