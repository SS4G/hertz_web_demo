package model

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
)

const (
	ALPHA_BET = "abcdefghijklmnopqrstuvwxyz"
)

var IMAGE_URLS = [6]string{
	"https://picx.zhimg.com/80/v2-c1d5fdc184a19e6693294947a9ae53fb_1440w.webp?source=1940ef5c",
	"https://pic1.zhimg.com/80/v2-5dc572dcd6923b4930724633db14602d_1440w.webp?source=1940ef5c",
	"https://pic1.zhimg.com/80/v2-e0cfb82310c6b4acfc708fcb9b1e4361_1440w.webp?source=1940ef5c",
	"https://picx.zhimg.com/80/v2-a559d84363f5f8b05e74f3b9f878379e_1440w.webp?source=1940ef5c",
	"https://pic4.zhimg.com/v2-556ba7daee5c231abde250693c85df03_b.jpg",
	"https://pic4.zhimg.com/v2-c53767509dff966734c5bad04d88b2eb_b.jpg",
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
	UserName       string `json:"user_name" column:"user_name"`
	Password       string `json:"password" column:"password"`
	AvatarImageUrl string `json:"avatar_image_url" column:"avatar_image_url"`
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

var TestUsers []*User

var TestDocs []*Doc

//Only for test use
func GenTestUsers() {
	userNames := [4]string{"Putin", "Trump", "Biden", "KimJongUn"}
	avatarImages := [4]string{
		"http://www.news.cn/mil/2022-12/22/1211711449_16716728965781n.jpg",
		"http://www.xinhuanet.com/world/2019-08/10/1124858824_15653878515811n.jpg",
		"http://www.news.cn/world/2022-01/20/1211537014_16426748585311n.jpg",
		"http://www.xinhuanet.com/world/2021-01/08/1126959898_16101753558321n.jpg",
	}

	titles := [4]string{
		"普京在俄国防部会议上最新讲话",
		"特朗普再次收到金正恩信件",
		"拜登执政一年：从踌躇满志到挫败连连",
		"金正恩在朝鲜劳动党八大提出新五年计划",
	}

	texts := [4]string{
		"自俄罗斯于2月24日对乌克兰发起特别军事行动以来，已有将近10个月时间。综合俄新社、塔斯社等多家俄媒报道，当地时间12月21日，俄罗斯总统普京、防长绍伊古出席俄国防部部务委员会扩大会议，就俄军部队战备状态等话题作出最新表态。",
		"特朗普与金正恩先后在新加坡、河内和板门店举行三次会晤。今年6月30日，特朗普在板门店与金正恩举行会谈后对媒体表示，美朝双方将在未来两三周重启工作磋商。美国国务卿蓬佩奥本月7日表示，期待能在未来数周内重启与朝方的对话磋商。",
		"拜登上台之初，支持率超过半数，一度接近六成。执政百日时，他基本兑现了广泛接种疫苗、大规模重开学校等竞选承诺，美国疫情一度显著回落。但先是德尔塔毒株带来新一波疫情，继而奥密克戎毒株席卷美国，使美国依赖疫苗摆脱疫情的设想破灭。",
		"据朝中社8日报道，朝鲜劳动党委员长、国务委员会委员长、武装力量最高司令官金正恩在朝鲜劳动党第八次代表大会上提出新的五年计划发展目标，实现可持续生产增长，推动教育、卫生、文学艺术等社会主义文化各领域创新。",
	}
	var i uint
	for i = 0; i < 4; i++ {
		userName := userNames[i]
		imageUrl := avatarImages[i]
		u := &User{
			UserName:       userName,
			Password:       "12345",
			AvatarImageUrl: imageUrl,
		}
		TestUsers = append(TestUsers, u)

		d := &Doc{
			ImageUrl:   IMAGE_URLS[i],
			Title:      titles[i],
			Text:       texts[i],
			AuthorId:   i + 1,
			AuthorName: u.UserName,
		}
		TestDocs = append(TestDocs, d)
		log.Infof("u=%+v", u)
		log.Infof("d=%+v", d)
	}
}
