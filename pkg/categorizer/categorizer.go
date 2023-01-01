// categorizer is the package responsible to parse
// and apply rules for categorization present inside
// the YAML Ruleuration files
//
// Below, an example of YAML:
//
//	---
//	rules:
//	- regex: STARBUCKS
//		category: BREAKFAST
//	- regex: MC(.*)DONALDS
//		category: LUNCH
//	- regex: WALMART
//		category: SUPERMARKET
package categorizer

import (
	"os"

	"github.com/thiagohdeplima/financial-control/pkg/fileutil"
	"gopkg.in/yaml.v3"
)

type Rule struct {
	Pattern  string `yaml:"pattern"`
	Category string `yaml:"category"`
}

type Rules struct {
	Rules []Rule `yaml:"rules"`
}

func ParseRules(yamlPath string) (Rules, error) {
	var yml = Rules{}

	if err := fileutil.FileExistsWithError(yamlPath); err != nil {
		return yml, err
	}

	ymlBytes, err := os.ReadFile(yamlPath)
	if err != nil {
		return yml, err
	}

	if err := yaml.Unmarshal(ymlBytes, &yml); err != nil {
		return yml, err
	}

	return yml, nil
}
