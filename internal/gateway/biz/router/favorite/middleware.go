// Code generated by hertz generator.

package favorite

import (
	"github.com/cloudwego/hertz/pkg/app"
	"tuzi-tiktok/gateway/biz/router/control"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		control.Authentication(),
	}
}

func _favorvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getfavoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}
