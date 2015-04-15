package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnvConfigTestSuite struct {
	suite.Suite
	varName     string
	varValue    string
	emptySource map[string]string
	source      map[string]string
}

func (suite *EnvConfigTestSuite) SetupTest() {
	suite.varName = "GOTWELVE_TEST_VAR"
	suite.varValue = "test"
	suite.source = make(map[string]string)
	suite.source[suite.varName] = suite.varValue
}

func (suite *EnvConfigTestSuite) TestGetEnvironmentVariable() {
	config, _ := newConfig(SpecMap{
		"test:var": {Name: suite.varName},
	}, suite.source)

	assert.Equal(suite.T(), suite.varValue, config["test:var"])
}

func (suite *EnvConfigTestSuite) TestMissingEnvironmentVariable() {
	_, err := newConfig(SpecMap{
		"test:var": {Name: suite.varName},
	}, suite.emptySource)

	assert.NotEqual(suite.T(), nil, err)
}

func (suite *EnvConfigTestSuite) TestDefaultEnvironmentVariable() {
	config, _ := newConfig(SpecMap{
		"test:var": {
			Name:         suite.varName,
			DefaultValue: "default",
		},
	}, suite.emptySource)

	assert.Equal(suite.T(), "default", config["test:var"])
}

func (suite *EnvConfigTestSuite) TestOptionalEnvironmentVariable() {
	config, _ := newConfig(SpecMap{
		"test:var": {
			Name:     suite.varName,
			Optional: true,
		},
	}, suite.emptySource)

	assert.Equal(suite.T(), "", config["test:var"])
}

func (suite *EnvConfigTestSuite) TestEnvironmentVariableParser() {
	config, _ := newConfig(SpecMap{
		"test:var": {
			Name: suite.varName,
			Parser: func(val string) string {
				return val + " world"
			},
		},
	}, suite.source)

	assert.Equal(suite.T(), "test world", config["test:var"])
}

func (suite *EnvConfigTestSuite) TestOptionalEnvironmentVariableParser() {
	config, _ := newConfig(SpecMap{
		"test:var": {
			Name:     suite.varName,
			Optional: true,
			Parser: func(val string) string {
				return val + " world"
			},
		},
	}, suite.emptySource)

	assert.Equal(suite.T(), " world", config["test:var"])
}

func (suite *EnvConfigTestSuite) TestEnvironmentVariableFlag() {
	config, _ := newConfig(SpecMap{
		"test:var": {Name: suite.varName},
	}, suite.source)

	assert.Equal(suite.T(), true, config.AsFlag("test:var"))
}

func TestEnvConfigTestSuite(t *testing.T) {
	suite.Run(t, new(EnvConfigTestSuite))
}
