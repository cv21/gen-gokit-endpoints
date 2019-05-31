package generator

import (
	"testing"

	"github.com/cv21/gen/pkg"
)

func TestEndpointsGenerator_Generate(t *testing.T) {
	testCases := []pkg.TestCase{
		{
			Name: "basic",
			Params: `
{
	"interface_name": "StringService"
}
`,
			GeneratedFilePaths: []string{"./transport/endpoints_gen.go"},
		},
	}

	pkg.RunTestCases(t, testCases, NewGenerator())
}
