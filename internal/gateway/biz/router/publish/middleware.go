// Code generated by hertz generator.

package publish

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

func _publishMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		control.Authentication(),
	}
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getpublishlistMw() []app.HandlerFunc {
	// your code...
	return nil
}
