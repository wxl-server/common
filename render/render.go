package render

import "github.com/qcq1/common/json"

func Render(object any) string {
	return json.MarshalWithoutError[string](object)
}
