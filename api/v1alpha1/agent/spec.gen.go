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