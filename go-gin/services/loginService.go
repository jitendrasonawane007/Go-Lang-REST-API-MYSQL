package services

import (
	"fmt"

	models "go-gin/models"

	helper "go-gin/helpers"

	"github.com/gchaincl/dotsql"
)

func Login(postParam models.Login) (UserInfo models.UserInfo) {
	// var Token = "test"
	// var module models.Modules
	// var db = helper.DBConnection()
	// var moduleArr []models.Modules
	// Token := "test"
	isValidLogin, roleID := validateUser(postParam)
	if isValidLogin == 1 {
		var Token = helper.GenerateToken(postParam.LoginId, postParam.Password)
		moduleArr := fetchModuleInfo(roleID)
		UserInfo.Token = Token
		UserInfo.Modules = moduleArr
	} else {
		fmt.Println(`Invalid Login`)
	}
	return UserInfo
}

/****************************************************************************/
func validateUser(postParam models.Login) (loginCount int, roleID int) {
	var db = helper.DBConnection()
	loginSQL, err := dotsql.LoadFromFile("./database/loginQuery.sql")
	if err != nil {
		fmt.Printf("Error", err)
	}
	defer db.Close()
	result, err := loginSQL.Query(db, "getLoginCount", postParam.LoginId, postParam.Password)
	defer result.Close()
	if err != nil {
		fmt.Println(`Error LoginService-10`)
	}
	for result.Next() {
		err := result.Scan(&loginCount, &roleID)
		if err != nil {
			panic(err.Error())
		}
	}
	return loginCount, roleID
}

/****************************************************************************/

func fetchModuleInfo(roleID int) (moduleArr []models.Modules) {
	var db = helper.DBConnection()
	loginSQL, err := dotsql.LoadFromFile("./database/loginQuery.sql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	result, err := loginSQL.Query(db, "getModulesByRoleId", roleID)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var module models.Modules
	for result.Next() {
		err := result.Scan(&module.ModuleId, &module.ModuleName, &module.ModuleShortName, &module.ModuleIconClass)
		if err != nil {
			panic(err.Error())
		}
		moduleArr = append(moduleArr, module)
	}
	return moduleArr

}

/****************************************************************************/

func Save(postParam models.UserDetails) {

}

/****************************************************************************/
func validateUserDetails(postParam models.UserDetails) {

}
