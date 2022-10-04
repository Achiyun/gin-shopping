package middlewares

import (
	"encoding/json"
	"fmt"

	"github.com/Achiyun/gin-shopping/server/intetnal/app/global"
	models "github.com/Achiyun/gin-shopping/server/intetnal/app/model/admin"

	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(c *gin.Context) {

	//进行权限判断 没有登录的用户 不能进入后台管理中心

	//1、获取Url访问的地址  /admin/captcha

	//2、获取Session里面保存的用户信息

	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行

	//4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	//  1、获取Url访问的地址   /admin/captcha?t=0.8706946438889653
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	// 2、获取Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	fmt.Println("----------------------------1我是权限判断")

	if ok {
		var userinfoStruct models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if err != nil || !(userinfoStruct != models.Manager{} && userinfoStruct.Username != "") {

			if pathname != "/admin/user/login" && pathname != "/admin/user/doLogin" && pathname != "/base/captcha" {
				c.Redirect(302, "/admin/user/login")
			}
		} else { //用户登录成功 权限判断
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			if userinfoStruct.IsSuper == 0 && !excludeAuthPath("/"+urlPath) {

				// 1、根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
				roleAccess := []models.RoleAccess{}
				global.GVA_DB.Where("role_id=?", userinfoStruct.RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}
				// 2、获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
				// pathname      /admin/manager
				access := models.Access{}
				global.GVA_DB.Where("url = ?", urlPath).Find(&access)
				//3、判断当前访问的url对应的权限id 是否在权限列表的id中
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort()
				}
			}
		}
	} else { //没有登录
		if pathname != "/admin/user/login" && pathname != "/admin/user/doLogin" && pathname != "/base/captcha" {
			c.Redirect(302, "/admin/user/login")
		}
	}

}

//排除权限判断的方法

func excludeAuthPath(urlPath string) bool {
	// config, iniErr := ini.Load("./conf/app.ini")
	// if iniErr != nil {
	// 	fmt.Printf("Fail to read file: %v", iniErr)
	// 	os.Exit(1)
	// }
	// excludeAuthPath := config.Section("").Key("excludeAuthPath").String()
	excludeAuthPath := global.GVA_CONFIG.System.ExcludeAuthPath

	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	// return true
	fmt.Printf("excludeAuthPathSlice: %v\n", excludeAuthPathSlice)
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
