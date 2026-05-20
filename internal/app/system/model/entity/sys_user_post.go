package entity


type SysUserPost struct {
	UserId int64 `json:"userId" description:"用户ID"`
	PostId int64 `json:"postId" description:"岗位ID"`
}
