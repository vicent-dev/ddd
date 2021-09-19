package template

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
)

type ContextTemplate struct {
	name       string
	domainCode string
}

func NewContextTemplate(context string) *ContextTemplate {
	ct := &ContextTemplate{name: strcase.ToLowerCamel(context)}
	ct.domainCode = fmt.Sprintf(domainTemplate, strcase.ToCamel(context))

	return ct
}

func (c *ContextTemplate) Generate() error {

	folders := []string{"Application", "Domain", "Infrastructure"}

	for _, f := range folders {
		if err := os.MkdirAll("./"+c.name+"/"+f, 0755); err != nil {
			return err
		}
	}

	f, err := os.Create("./" + c.name + "/Domain/" + c.name + ".go")

	if err != nil {
		return err
	}

	if _, err := f.WriteString(c.domainCode); err != nil {
		return err
	}

	return nil
}
