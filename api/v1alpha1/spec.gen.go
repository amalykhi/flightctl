// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
