package mysql

import (
	"bluebell/models"
	"bluebell/setting"
	"testing"
)

func init() {
	dbCfg := setting.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "Long5230413",
		DB:           "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          10,
		Title:       "test",
		Content:     "just a test",
		AuthorID:    123,
		CommunityID: 2,
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("create post failed, err:%v\n", err)
	}
	t.Logf("create post success")
}
