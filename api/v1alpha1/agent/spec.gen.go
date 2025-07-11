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

	externalRef0 "github.com/flightctl/flightctl/api/v1alpha1"
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

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../openapi.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
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
