package template

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
)

type QueryTemplate struct {
	ct              *ContextTemplate
	context         string
	name            string
	applicationCode string
}

func NewQueryTemplate(context, query string) *QueryTemplate {
	qt := &QueryTemplate{
		context: context,
		name:    strcase.ToLowerCamel(query),
		ct:      NewContextTemplate(context),
	}

	cQuery := strcase.ToCamel(query)
	lcQuery := strcase.ToLowerCamel(query)

	qt.applicationCode = fmt.Sprintf(queryTempalte, cQuery, lcQuery, cQuery, lcQuery, cQuery, cQuery, cQuery, cQuery)

	return qt
}

func (q *QueryTemplate) Generate() error {
	q.ct.GenerateIfNotExists()

	f, err := os.Create("./" + q.context + "/Application/" + q.name + "Query.go")

	if err != nil {
		return err
	}

	if _, err := f.WriteString(q.applicationCode); err != nil {
		return err
	}

	return nil
}
