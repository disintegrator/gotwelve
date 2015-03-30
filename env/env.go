package env

import (
	"fmt"
	"os"
)

// VarSpec defines a config variable. It specifies its name in the environment,
// a parser function to process the raw value, a default value and whether or
// not the variable must be set in the environment
type VarSpec struct {
	Name         string
	Parser       func(string) string
	DefaultValue string
	Optional     bool
}

// Config is a mapping that maps and config name to its value retreived from the
// environment
type Config map[string]string
type envVarNotDefined struct {
	name string
}

func (e *envVarNotDefined) Error() string {
	return fmt.Sprintf("Environment variable not defined: %s", e.name)
}

func noopParser(val string) string { return val }

// NewConfigFromEnv Constructs a config mapping by looking up the OS environment
// variables
func NewConfigFromEnv(spec map[string]VarSpec) (Config, error) {
	source := make(map[string]string, len(spec))
	for _, varspec := range spec {
		source[varspec.Name] = os.Getenv(varspec.Name)
	}
	return newConfig(spec, source)
}

func newConfig(spec map[string]VarSpec, source map[string]string) (Config, error) {
	config := make(Config)
	for key, envVar := range spec {
		parser := noopParser
		if envVar.Parser != nil {
			parser = envVar.Parser
		}

		val := source[envVar.Name]

		if val == "" {
			val = envVar.DefaultValue
		}

		if val == "" && !envVar.Optional {
			return nil, &envVarNotDefined{envVar.Name}
		}

		val = parser(val)

		config[key] = val
	}
	return config, nil
}

// AsFlag treats a config variable as a flag returning false for zero-length
// strings and true otherwise
func (e Config) AsFlag(name string) bool {
	return e[name] != ""
}
