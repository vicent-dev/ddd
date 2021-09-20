package template

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
)

type CommandTemplate struct {
	ct              *ContextTemplate
	context         string
	name            string
	applicationCode string
}

func NewCommandTemplate(context, command string) *CommandTemplate {
	ct := &CommandTemplate{
		context: context,
		name:    strcase.ToLowerCamel(command),
		ct:      NewContextTemplate(context),
	}

	cCommand := strcase.ToCamel(command)
	lcCommand := strcase.ToLowerCamel(command)

	ct.applicationCode = fmt.Sprintf(commandTempalte, cCommand, lcCommand, cCommand, lcCommand, cCommand, cCommand)

	return ct
}

func (c *CommandTemplate) Generate() error {
	c.ct.GenerateIfNotExists()

	f, err := os.Create("./" + c.context + "/Application/" + c.name + "Command.go")

	if err != nil {
		return err
	}

	if _, err := f.WriteString(c.applicationCode); err != nil {
		return err
	}

	return nil
}
