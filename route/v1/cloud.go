package v1

import (
	"github.com/labstack/echo/v4"
)

func GetStorage(ctx echo.Context) error {
	// idStr := ctx.QueryParam("id")
	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	return ctx.JSON(common_err.SUCCESS, model.Result{Success: common_err.CLIENT_ERROR, Message: common_err.GetMsg(common_err.CLIENT_ERROR), Data: err.Error()})
	// 	return
	// }
	// storage, err := service.MyService.Storage().GetStorageById(uint(id))
	// if err != nil {
	// 	return ctx.JSON(common_err.SUCCESS, model.Result{Success: common_err.SERVICE_ERROR, Message: common_err.GetMsg(common_err.SERVICE_ERROR), Data: err.Error()})
	// 	return
	// }
	// return ctx.JSON(common_err.SUCCESS, model.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS), Data: storage})
	return nil
}
