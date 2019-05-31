package generator

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cv21/gen/pkg"
	. "github.com/dave/jennifer/jen"
)

const (
	generatorName            = "github.com/cv21/gen-gokit-endpoints"
	generatorVersion         = "1.0.0"
	endpointsSetName         = "EndpointsSet"
	packagePathGoKitEndpoint = "github.com/go-kit/kit/endpoint"
	generatedFilePath        = "./transport/endpoints_gen.go"
)

type (
	generatorParams struct {
		// It is an interface name of service.
		// Example: StringService
		InterfaceName string `json:"interface_name"`
	}

	endpointsGenerator struct {
	}
)

func NewGenerator() pkg.Generator {
	return &endpointsGenerator{}
}

func (e *endpointsGenerator) Generate(params *pkg.GenerateParams) (*pkg.GenerateResult, error) {
	genParams := &generatorParams{}
	err := json.Unmarshal(params.Params, genParams)
	if err != nil {
		return nil, err
	}

	iface := pkg.FindInterface(params.File, genParams.InterfaceName)
	if iface == nil {
		return nil, errors.New("could not find interface")
	}

	f := NewFile("transport")
	pkg.AddDefaultPackageComment(f, generatorName, generatorVersion)

	f.Type().Id(endpointsSetName).StructFunc(func(g *Group) {
		for _, signature := range iface.Methods {
			g.Id(signature.Name+"Endpoint").Qual(packagePathGoKitEndpoint, "Endpoint")
		}
	}).Line()

	return &pkg.GenerateResult{
		Files: []pkg.GenerateResultFile{
			{
				Path:    generatedFilePath,
				Content: []byte(fmt.Sprintf("%#v", f)),
			},
		},
	}, nil
}
