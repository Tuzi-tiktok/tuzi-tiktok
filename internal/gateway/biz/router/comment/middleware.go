// Code generated by hertz generator.

package comment

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

func _commentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		control.Authentication(),
	}
}

func _comment0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcommentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}
