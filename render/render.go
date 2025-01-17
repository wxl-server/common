package render

import "github.com/wxl-server/common/json"

func Render(object any) string {
	return json.MarshalWithoutError[string](object)
}
