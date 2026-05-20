package model

type Context struct {
	User *ContextUser
}

type ContextUser struct {
	*LoginUserRes
}
