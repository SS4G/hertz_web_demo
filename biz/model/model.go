package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

const (
	ALPHA_BET = "abcdefghijklmnopqrstuvwxyz"
)

var IMAGE_URLS = [10]string{
	"https://pic1.zhimg.com/80/v2-ac6781b718e51ed9487e6dac55acf28a_1440w.webp?source=1940ef5c",
	"https://pic1.zhimg.com/80/v2-37058e9aa8202ac7d6b435748f9d8b0d_1440w.webp?source=1940ef5c",
	"https://picx.zhimg.com/80/v2-4cf9aa91cb2d0e2292df1f8561a3819c_1440w.webp?source=1940ef5c",
	"https://pic1.zhimg.com/80/v2-d873e1aa21a53c3af45e417b68606c91_1440w.webp?source=1940ef5c",
	"https://picx.zhimg.com/80/v2-be171a1375b91a07fc5d9aa9844d5486_1440w.webp?source=1940ef5c",
	"https://picx.zhimg.com/80/v2-d1b4faf473ebfdf8673ea27cb9ebd1e8_1440w.webp?source=1940ef5c",
	"https://picx.zhimg.com/80/v2-11c7e649f10debdd081a15d926a39d69_1440w.webp?source=1940ef5c",
	"https://picx.zhimg.com/80/v2-9b6c16b7e0db7458b37f208f53038ade_1440w.webp?source=1940ef5c",
	"https://picx.zhimg.com/80/v2-34567f8b4eee551acff9357cdd44e72c_1440w.webp?source=1940ef5c",
	"https://picx.zhimg.com/80/v2-6b97f8854def20af3ccff10e470326c3_1440w.webp?source=1940ef5c",
}

var randGen *rand.Rand = nil

// init randgen
func init() {
	randGen = rand.New(rand.NewSource(99))
	log.Infof("model init called")
}

// 相当于所有的model类都继承自 grom.Model
// 通过内嵌字段实现了继承
type User struct {
	gorm.Model
	//UserId   uint32 `json:"user_id" column:"user_id"`
	UserName string `json:"user_name" column:"user_name"`
	Password string `json:"password" column:"password"`
}

type Doc struct {
	gorm.Model
	//DocId    uint32 `json:"doc_id" column:"doc_id"`
	ImageUrl   string `json:"image_url" column:"image_url"`
	Title      string `json:"title" column:"title"`
	Text       string `json:"text" column:"text"`
	AuthorId   uint   `json:"author_id" column:"author_id"`
	AuthorName string `json:"author_name" column:"author_name"`
}

//Only for test use
func NewRandomUser() *User {
	userNames := [5]string{"Nick", "Dean", "Jeff", "Andy", "Lucy"}
	randUserName := fmt.Sprintf("%s_%d", userNames[randGen.Intn(5)], rand.Intn(1000000))
	return &User{
		UserName: randUserName,
		Password: "12345",
	}
}

func NewRandomDoc() *Doc {
	titles := [5]string{
		"我昨天买的JK服",
		"自拍了一张",
		"女主播需要大哥点点关注",
		"大哥给我刷了个宝马",
		"刷个城堡可以加微信",
	}

	texts := [5]string{
		"看这是我的新JK",
		"昨天在公园的自拍",
		"家人们赶紧点点关注，对面马上要反超了",
		"王哥 我好感动哦",
		"加了微信有福利啊哥，加一个看看呗",
	}

	randImageUrl := IMAGE_URLS[randGen.Intn(len(IMAGE_URLS))]
	randIdx := randGen.Intn(5)
	randTitle := titles[randIdx] + fmt.Sprintf(" 主题发布@%s", time.Now().Format(time.RFC822))
	randText := texts[randIdx] + fmt.Sprintf(" 文章发布@%s", time.Now().Format(time.RFC822))
	return &Doc{
		ImageUrl:   randImageUrl,
		Title:      randTitle,
		Text:       randText,
		AuthorId:   1,
		AuthorName: "wangyan.666",
	}
}
