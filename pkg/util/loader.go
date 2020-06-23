package util

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func LoadSwagger(filePath string) (*openapi3.Swagger, error) {
	loader := openapi3.NewSwaggerLoader()
	loader.IsExternalRefsAllowed = true
	swagger, err := loader.LoadSwaggerFromFile(filePath)
	if err != nil {
		return nil, err
	}
	return swagger, nil
}
