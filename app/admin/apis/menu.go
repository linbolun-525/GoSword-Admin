package apis

import (
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/common/api"
	"project/utils"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// InsertMenuHandler 新增菜单
// @Summary 新增菜单
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：菜单管理 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.InsertMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseInsertMenu
// @Router /api/menus [post]
func InsertMenuHandler(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.InsertMenuDto)
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("InsertMenuHandler failed", zap.String("username", user.UserName), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	//业务逻辑处理
	m := new(service.Menu)
	if err := m.InsertMenu(p, user.UserId); err != nil {
		zap.L().Error("insert menu failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	//返回响应
	app.ResponseSuccess(c, nil)
}

// SelectMenuHandler 查询菜单
// @Summary 查询菜单
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：菜单管理 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.SelectMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseInsertMenu
// @Router /api/menus [get]
func SelectMenuHandler(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.SelectMenuDto)
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	if err := c.ShouldBindQuery(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("SelectMenuHandler failed", zap.String("username", user.UserName), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	//业务逻辑处理
	m := new(service.Menu)
	var data []*bo.SelectMenuBo
	data, err = m.SelectMenu(p)
	if err != nil {
		zap.L().Error("select menu failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	//返回响应
	app.ResponseSuccess(c, data)
}

// DeleteMenuHandler 删除菜单
// @Summary 删除菜单
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：菜单管理 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.DeleteMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseDeleteMenu
// @Router /api/menus [delete]
func DeleteMenuHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	//	绑定校验参数
	var ids []int
	if err := c.ShouldBind(&ids); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("DeleteMenuHandler failed", zap.String("username", user.UserName), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	menu := new(service.Menu)
	if err := menu.DeleteMenu(ids); err != nil {
		zap.L().Error("DeleteMenu failed", zap.Error(err))
		app.ResponseError(c, app.CodeDeleteOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// UpdateMenuHandler 更新菜单
// @Summary 更新菜单
// @Description Author：Cgl 2021/01/30 获得身份令牌
// @Tags 系统：菜单管理 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.DeleteMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseUpdateMenu
// @Router /api/menus [put]
func UpdateMenuHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	userId := user.UserId
	//	绑定校验参数
	p := new(dto.UpdateMenuDto)
	if err := c.ShouldBind(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("UpdateMenuHandler failed", zap.String("username", user.UserName), zap.Error(err))
		c.Error(err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	menu := new(service.Menu)
	if err := menu.UpdateMenu(p, userId); err != nil {
		zap.L().Error("UpdateMenu failed", zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}
	app.ResponseSuccess(c, nil)
}

// SelectForeNeedMenuHandler 查询前端所需菜单
// @Summary 查询前端所需菜单
// @Description Author：Cgl 2021/02/01 获得身份令牌
// @Tags 系统：菜单管理 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.DeleteMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseSelectForeNeedMenu
// @Router /api/menus/build [get]
func SelectForeNeedMenuHandler(c *gin.Context) {
	//获取上下文中信息
	user, err := api.GetCurrentUserInfo(c)
	if err != nil {
		zap.L().Error("GetCurrentUserInfo failed", zap.Error(err))
		return
	}
	//业务逻辑处理
	//TODO
	m := new(service.Menu)
	var data []*bo.SelectForeNeedMenuBo
	data, err = m.SelectForeNeedMenu(user)
	if err != nil {
		zap.L().Error("select menu failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	//返回响应
	app.ResponseSuccess(c, data)

}

// ReturnToAllMenusHandler 返回所有菜单
// @Summary 查询出该级别下属菜单
// @Description Author：Cgl 2021/02/04 获得身份令牌
// @Tags 系统：菜单管理 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body pid false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ReturnToAllMenusBo
// @Router /api/menus/lazy [get]
func ReturnToAllMenusHandler(c *gin.Context) {
	var pidInt int
	pid := c.Query("pid")
	pidInt, err := utils.StringToInt(pid)
	if err != nil {
		zap.L().Error("string转int出错", zap.Error(err))
		app.ResponseSuccess(c, app.CodeParamTypeBindError)
		return
	}
	// 业务逻辑处理
	// TODO
	m := new(service.Menu)
	var data []*bo.ReturnToAllMenusBo
	data, err = m.ReturnToAllMenus(pidInt)
	if err != nil {
		zap.L().Error("ReturnToAllMenus failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	// 返回响应
	app.ResponseSuccess(c, data)
}

// DownMenusHandler 导出菜单数据
// @Summary 导出菜单数据
// @Description Author：Cgl 2021/02/04 获得身份令牌
// @Tags 系统：菜单管理 Menu Controller
// @Accept application/json
// @Produce application/json
// @Param object body dto.DownloadMenuDto false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponseLogin
// @Router /api/menus/download [get]
func DownMenusHandler(c *gin.Context) {

	// 1. 获取参数，检验参数
	menu := new(dto.DownloadMenuDto)
	if err := c.ShouldBind(&menu); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	orderJsonData, err := utils.OrderJson(menu.Orders)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2. 参数正确执行响应
	m := service.Menu{}
	menuData, err := m.DownloadMenuInfoBo(menu, orderJsonData)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3. 返回文件数据
	var res []interface{}
	for _, menu := range menuData {
		res = append(res, &bo.DownloadMenuInfoBo{
			Title:      menu.Title,
			Type:       menu.Type,
			Permission: menu.Permission,
			IFrame:     menu.IFrame,
			Hidden:     menu.Hidden,
			Cache:      menu.Cache,
			CreateTime: menu.CreateTime,
		})
	}
	content := utils.ToExcel([]string{`菜单标题`, `菜单类型`, `权限标识`, `外链菜单`, `菜单可见`, `是否缓存`, `创建时间`}, res)
	utils.ResponseXls(c, content, "菜单数据")

}
