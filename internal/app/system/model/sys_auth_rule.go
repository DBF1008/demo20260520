package model

type SysAuthRuleInfoRes struct {
	Id        uint   `orm:"id,primary"  json:"id"`
	Pid       uint   `orm:"pid"         json:"pid"`
	Name      string `orm:"name,unique" json:"name"`
	Title     string `orm:"title"       json:"title"`
	Icon      string `orm:"icon"        json:"icon"`
	Condition string `orm:"condition"   json:"condition"`
	Remark    string `orm:"remark"      json:"remark"`
	MenuType  uint   `orm:"menu_type"   json:"menuType"`
	Weigh     int    `orm:"weigh"       json:"weigh"`
	IsHide    uint   `orm:"is_hide" json:"isHide"`
	IsCached  uint   `orm:"is_cached"  json:"isCached"`
	IsAffix   uint   `orm:"is_affix" json:"isAffix"`
	Path      string `orm:"path"        json:"path"`
	Redirect  string `orm:"redirect"   json:"redirect"`
	Component string `orm:"component"   json:"component"`
	IsIframe  uint   `orm:"is_iframe"    json:"isIframe"`
	IsLink    uint   `orm:"is_link" json:"isLink"`
	LinkUrl   string `orm:"link_url" json:"linkUrl"`
}


type SysAuthRuleTreeRes struct {
	*SysAuthRuleInfoRes
	Children []*SysAuthRuleTreeRes `json:"children"`
}

type UserMenu struct {
	Id        uint   `json:"id"`
	Pid       uint   `json:"pid"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Path      string `json:"path"`
	*MenuMeta `json:"meta"`
}

type UserMenus struct {
	*UserMenu `json:""`
	Children  []*UserMenus `json:"children"`
}

type MenuMeta struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	IsLink      string `json:"isLink"`
	IsHide      bool   `json:"isHide"`
	IsKeepAlive bool   `json:"isKeepAlive"`
	IsAffix     bool   `json:"isAffix"`
	IsIframe    bool   `json:"isIframe"`
}
