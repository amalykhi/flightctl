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
