package system

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterRole struct {

}

func (r *RouterRole) InitRouterRole(Router *gin.RouterGroup) (R gin.IRoutes) {
	roleGroup := Router.Group("role")
	apiRole := api.GroupApiApp.ApiSystem.ApiRole
	{
		// 获取角色列表
		roleGroup.POST("role-list", apiRole.GetRoleList)
		// 编辑角色信息
		roleGroup.PUT("role-edit", apiRole.EditRole)
		// 新增角色
		roleGroup.POST("role-add", apiRole.AddRole)
		// 删除角色
		roleGroup.DELETE("role-delete", apiRole.DeleteRole)
		// 根据ID查找角色
		roleGroup.POST("role-id", apiRole.QueryRoleById)
		// 获取角色菜单
		roleGroup.POST("role-menu", apiRole.GetRoleMenuList)
		// 编辑角色菜单
		roleGroup.PUT("role-menu-edit", apiRole.EditRoleMenu)
		// 获取角色API
		roleGroup.POST("role-api", apiRole.GetRoleApiList)
		// 编辑角色Api
		roleGroup.PUT("role-api-edit", apiRole.EditRoleApi)
		// 根据角色查找用户
		roleGroup.POST("role-user", apiRole.QueryUserByRole)
		// 从角色中移除某个用户
		roleGroup.POST("role-user-remove", apiRole.RemoveRoleUser)
		// 添加用户到某个角色
		roleGroup.POST("role-user-add", apiRole.AddRoleUser)
	}
	return Router
}
