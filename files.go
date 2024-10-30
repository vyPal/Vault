package main

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
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

func GetFolderSize(path string) (int64, error) {
	var size int64
	for objeft := range minioClient.ListObjects(context.Background(), os.Getenv("MINIO_BUCKET"), minio.ListObjectsOptions{Prefix: path, Recursive: true, WithVersions: true}) {
		if objeft.Err != nil {
			return 0, objeft.Err
		}
		size += objeft.Size
	}
	return size, nil
}
