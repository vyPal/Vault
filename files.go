package main

import (
	"strings"
	"time"
)

type AccessType int

const (
	READ AccessType = iota
	WRITE
	LIST
	DELETE
	APP
	ALL
)

func (a AccessType) String() string {
	return [...]string{"files:read", "files:write", "files:list", "files:delete", "files:app", "files:all"}[a]
}

func HasAccess(token TokenData, access AccessType) bool {
	if !token.Active {
		return false
	}

	if time.Now().Unix() > int64(token.Exp) {
		return false
	}

	scopes := strings.Split(token.Scope, " ")
	for _, scope := range scopes {
		if scope == access.String() {
			return true
		}
	}

	return false
}
