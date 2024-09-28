package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func SetupRoutes(g *gin.Engine) {
	g.GET("/health", HealthCheck)

	e := g.Group("/files", ValidateTokenMiddleware())
	e.GET("/metadata/*path", HandleFileMetadata)
	e.GET("/download/*path", HandleFileDownload)
	e.POST("/upload/*path", HandleFileUpload)
	e.DELETE("/delete/*path", HandleFileDelete)
	e.GET("/list/*path", HandleFileList)
	// TODO: Implement file sharing
}

func SplitPath(path string) (user string, client string, filePath string) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return "", "", ""
	}
	return parts[1], parts[2], strings.Join(parts[3:], "/")
}

func CanAccessFile(token TokenData, user string, client string, filePath string) bool {
	// TODO: This might need to be replaced when sharing is implemented
	if token.Sub != user {
		return false
	}
	if HasAccess(token, AccessType(ALL)) {
		return true
	}
	if HasAccess(token, AccessType(APP)) && token.GetIssuer() == client {
		return true
	}
	return false
}

func HandleFileMetadata(c *gin.Context) {
	t, _ := c.Get("tokenData")
	token := t.(TokenData)
	user, client, filePath := SplitPath(c.Param("path"))
	if !CanAccessFile(token, user, client, filePath) {
		c.String(403, "You do not have permission to access this file")
		return
	}
	if !HasAccess(token, AccessType(READ)) {
		c.String(403, "You do not have permission to read this file")
		return
	}
	info, err := minioClient.StatObject(context.Background(), os.Getenv("MINIO_BUCKET"), user+"/"+client+"/"+filePath, minio.StatObjectOptions{})
	if err != nil {
		c.String(500, "Error getting file metadata: %s", err)
		return
	}
	c.JSON(200, info)
}

func HandleFileDownload(c *gin.Context) {
	t, _ := c.Get("tokenData")
	token := t.(TokenData)
	user, client, filePath := SplitPath(c.Param("path"))
	if !CanAccessFile(token, user, client, filePath) {
		c.String(403, "You do not have permission to access this file")
		return
	}
	if !HasAccess(token, AccessType(READ)) {
		c.String(403, "You do not have permission to read this file")
		return
	}
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
	url, err := minioClient.PresignedGetObject(context.Background(), os.Getenv("MINIO_BUCKET"), user+"/"+client+"/"+filePath, time.Duration(15*time.Minute), reqParams)
	if err != nil {
		c.String(500, "Error generating presigned URL: %s", err)
		return
	}

	ProxyRequest(c, "GET", url)
}

func HandleFileUpload(c *gin.Context) {
	t, _ := c.Get("tokenData")
	token := t.(TokenData)
	user, client, filePath := SplitPath(c.Param("path"))
	if !CanAccessFile(token, user, client, filePath) {
		c.String(403, "You do not have permission to access this file")
		return
	}
	if !HasAccess(token, AccessType(WRITE)) {
		c.String(403, "You do not have permission to write to this file")
		return
	}
	url, err := minioClient.PresignedPutObject(context.Background(), os.Getenv("MINIO_BUCKET"), user+"/"+client+"/"+filePath, time.Duration(15*time.Minute))
	if err != nil {
		c.String(500, "Error generating presigned URL: %s", err)
		return
	}

	ProxyRequest(c, "PUT", url)
}

func HandleFileDelete(c *gin.Context) {
	t, _ := c.Get("tokenData")
	token := t.(TokenData)
	user, client, filePath := SplitPath(c.Param("path"))
	if !CanAccessFile(token, user, client, filePath) {
		c.String(403, "You do not have permission to access this file")
		return
	}
	if !HasAccess(token, AccessType(WRITE)) { // TODO: Change to DELETE
		c.String(403, "You do not have permission to write to this file")
		return
	}
	err := minioClient.RemoveObject(context.Background(), os.Getenv("MINIO_BUCKET"), user+"/"+client+"/"+filePath, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
	if err != nil {
		c.String(500, "Error deleting file: %s", err)
		return
	}
	c.String(200, "File deleted")
}

func HandleFileList(c *gin.Context) {
	t, _ := c.Get("tokenData")
	token := t.(TokenData)
	user, client, filePath := SplitPath(c.Param("path"))
	if !CanAccessFile(token, user, client, filePath) {
		c.String(403, "You do not have permission to access this file")
		return
	}
	if !HasAccess(token, AccessType(LIST)) {
		c.String(403, "You do not have permission to write to this file")
		return
	}
	if !strings.HasSuffix(filePath, "/") && filePath != "" {
		filePath += "/"
	}
	opts := minio.ListObjectsOptions{
		Recursive:    c.Query("recursive") == "true",
		Prefix:       user + "/" + client + "/" + filePath,
		WithVersions: c.Query("versions") == "true",
	}

	files := []minio.ObjectInfo{}

	for object := range minioClient.ListObjects(context.Background(), os.Getenv("MINIO_BUCKET"), opts) {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		files = append(files, object)
	}

	c.JSON(200, files)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
