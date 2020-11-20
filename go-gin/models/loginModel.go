package models

type Login struct {
	LoginId  string `json:"userName" binding:"required`
	Password string `json:"password" binding:"required`
}

type Modules struct {
	ModuleId        string `json:"moduleId"`
	ModuleName      string `json:moduleName`
	ModuleShortName string `json:moduleShortName`
	ModuleIconClass string `json:moduleIconClass`
}

type UserInfo struct {
	Token   string
	Modules []Modules
}

type UserDetails struct {
	FirstName    string `json:"firstName" binding:"required`
	LastName     string `json: lastName`
	MobileNumber string `json: mobileNumber`
	DOB          string `json:DOB`
	EmailId      string `json:emailId`
}
