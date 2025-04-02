package service

import "blog_backend/global"

type jwtService struct {
}

func (jwtService *jwtService) IsInBlackList(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

func LoadAll() {

}
