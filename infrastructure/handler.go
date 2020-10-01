package infrastructure

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type m map[string]interface{}

func JSONErrorHandler(err error, c echo.Context) {
	// handle error
	// TODO: ここで例外発生時のレンダリングや操作など行う
	// var appErr *AppErr
	// if errors.As(err, &appErr) {

	// } else {
	// 	appErr = UnknownErr
	// }
	// c.JSON(appErr.Code, )
	fmt.Printf("%+v", errors.Wrap(err, "internal error occurs"))
	// fmt.Println()
}
