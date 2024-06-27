package service

import "os/user"

type SaveUser interface {
	Save(*user.User) error
}
