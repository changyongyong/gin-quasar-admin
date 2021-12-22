package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysApi = new(sysApi)

type sysApi struct{}

var sysApiData = []system.SysApi{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1001, Remark: "获取用户列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1002, Remark: "编辑用户信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1003, Remark: "新增用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1004, Remark: "删除用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1005, Remark: "根据ID查找用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1006, Remark: "获取用户的菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-menu", ApiMethod: "GET",
	},
	//{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1007, Remark: "获取用户的角色列表", CreatedAt: time.Now(), CreatedBy: "admin"},
	//	ApiGroup: "user", ApiPath: "/user/user-role", ApiMethod: "GET",
	//},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1008, Remark: "用户修改密码", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "user", ApiPath: "/user/user-change-password", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1009, Remark: "获取角色列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1010, Remark: "编辑角色信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1011, Remark: "新增角色", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1012, Remark: "删除角色", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1013, Remark: "根据ID查找角色", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1014, Remark: "获取角色菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-menu", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1015, Remark: "编辑角色菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-menu-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1016, Remark: "获取角色API", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-api", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1017, Remark: "编辑角色Api", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-api-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1018, Remark: "根据角色查找用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-user", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1019, Remark: "从角色中移除某个用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-user-remove", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1020, Remark: "添加用户到某个角色", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-user-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1021, Remark: "修改角色部门数据权限", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "role", ApiPath: "/role/role-dept-data-permission-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1022, Remark: "获取菜单列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1023, Remark: "编辑菜单信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1024, Remark: "新增菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1025, Remark: "删除菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1026, Remark: "根据ID查找菜单", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "menu", ApiPath: "/menu/menu-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1027, Remark: "获取部门列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1028, Remark: "编辑部门信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1029, Remark: "新增部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1030, Remark: "删除部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1031, Remark: "根据ID查找部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1032, Remark: "根据部门查找用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-user", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1033, Remark: "从部门中移除某个用户", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-user-remove", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1034, Remark: "添加用户到某个部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dept", ApiPath: "/dept/dept-user-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1035, Remark: "获取根字典列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1036, Remark: "编辑字典信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1037, Remark: "新增字典", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1038, Remark: "删除字典", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1039, Remark: "根据ID查找字典", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "dict", ApiPath: "/dict/dict-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1040, Remark: "获取api列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "api", ApiPath: "/api/api-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1041, Remark: "编辑api信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "api", ApiPath: "/api/api-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1042, Remark: "新增api", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "api", ApiPath: "/api/api-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1043, Remark: "删除api", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "api", ApiPath: "/api/api-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1044, Remark: "上传头像", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "upload", ApiPath: "/upload/avatar", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1045, Remark: "上传文件", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "upload", ApiPath: "/upload/file", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1046, Remark: "上传网站Logo", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "upload", ApiPath: "/upload/web-logo", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1047, Remark: "上传标签页Logo", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "upload", ApiPath: "/upload/header-logo", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1048, Remark: "获取后台配置列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-backend", ApiPath: "/config-backend/config-backend-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1049, Remark: "编辑后台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-backend", ApiPath: "/config-backend/config-backend-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1050, Remark: "新增后台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-backend", ApiPath: "/config-backend/config-backend-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1051, Remark: "删除后台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-backend", ApiPath: "/config-backend/config-backend-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1052, Remark: "获取前台配置列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-frontend", ApiPath: "/config-frontend/config-frontend-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1053, Remark: "编辑前台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-frontend", ApiPath: "/config-frontend/config-frontend-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1054, Remark: "新增前台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-frontend", ApiPath: "/config-frontend/config-frontend-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1055, Remark: "删除前台配置", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "config-frontend", ApiPath: "/config-frontend/config-frontend-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1056, Remark: "获取登录日志列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "log", ApiPath: "/log/log-login-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1057, Remark: "删除登录日志", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "log", ApiPath: "/log/log-login-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1058, Remark: "获取操作日志列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "log", ApiPath: "/log/log-operation-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1059, Remark: "删除登录日志", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "log", ApiPath: "/log/log-operation-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1060, Remark: "获取消息列表", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "notice", ApiPath: "/notice/notice-list", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1061, Remark: "编辑消息信息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "notice", ApiPath: "/notice/notice-edit", ApiMethod: "PUT",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1062, Remark: "新增消息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "notice", ApiPath: "/notice/notice-add", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1063, Remark: "删除消息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "notice", ApiPath: "/notice/notice-delete", ApiMethod: "DELETE",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1064, Remark: "根据ID查找消息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "notice", ApiPath: "/notice/notice-id", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1065, Remark: "根据ID查找消息并阅读", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "notice", ApiPath: "/notice/notice-id-read", ApiMethod: "POST",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1066, Remark: "发送消息", CreatedAt: time.Now(), CreatedBy: "admin"},
		ApiGroup: "notice", ApiPath: "/notice/notice-send", ApiMethod: "POST",
	},
}

func (s *sysApi) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysApi{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_api 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[Gin-Quasar-Admin] --> sys_api 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_api 表初始数据成功！")
		global.GqaLog.Info("[Gin-Quasar-Admin] --> sys_api 表初始数据成功！")
		return nil
	})
}
