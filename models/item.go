package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type Item struct {
	Id          int64 `json:"item_id"`
	Title 	string `json:"item_name"`
	Description string `json:"item_description"`
	UserId	int64
	Password	string `json:"-"`
	Type	int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

type ItemInfo struct {
	Id          int64 `json:"item_id"`
	Title 	string `json:"item_domain"`
	IsArchived string `json:"is_archived"`
	DefaultPageId	string `json:"default_page_id"`
	DefaultCatId2	string `json:"default_cat_id2"`
	DefaultCatId3	string `json:"default_cat_id3"`
	UnreadCount		UnreadCount `json:"unread_count"`
	ItemType		int		`json:"item_type"`
	Menu 	Menu `json:"menu"`
	IsLogin bool `json:"is_login"`
	ItemPermn bool `json:"ItemPermn"`
	ItemCreator bool `json:"ItemCreator"`
}
type Menu struct {
	Page []Page `json:"pages"`
	Catalogs []Category `json:"catalogs"`
}

type Category struct {
	CatId int `json:"cat_id"`
	CatName string `json:"cat_name"`
	ItemId int `json:"item_id"`
	SNumber int `json:"s_number"`
	Addtime string `json:"addtime"`
	ParentCatId int `json:"parent_cat_id"`
	Level int `json:"level"`
	Page []Page `json:"page"`
	Category []Category `json:"catalogs"`
}

type Page struct {
	PageId int `json:"page_id"`
	AuthorUid int `json:"author_uid"`
	CatId int `json:"cat_id"`
	PageTitle string `json:"page_title"`
	Addtime string `json:"addtime"`
}

type UnreadCount struct {

}
func (this *Item) GetItemInfo() (ItemInfo) {
	var itemInfo ItemInfo
	itemInfo.Id = this.Id
	itemInfo.ItemType = this.Type
	itemInfo.IsLogin = true
	itemInfo.ItemPermn = true
	itemInfo.ItemCreator = true
	var pages = []Page{}
	var category = []Category{}
	itemInfo.Menu.Page = pages
	itemInfo.Menu.Catalogs = category
	//itemInfo.Menu.Page = append(itemInfo.Menu.Page, Page{})



	return itemInfo
}


func (this *Item) Create() (int64,error) {
	o := orm.NewOrm()

	id ,err := o.Insert(this)
	if err != nil {
		beego.Error(err.Error())
		err = fmt.Errorf("%s", "新增数据失败")
	}

	return id,err
}

func GetMyItem(uid int64) ([]*Item) {
	o := orm.NewOrm()
	var item []*Item
	num,err :=o.QueryTable("item").Filter("user_id", uid).All(&item)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,uid)

	return item
}

func GetOneItem(id int64) (bool,Item) {
	o := orm.NewOrm()
	item := Item{Id:id}

	err := o.Read(&item)

	if err == nil {
		return true,item
	} else {
		return false,item
	}

}
func GetMyItemInfo(id int64) (ItemInfo,error) {
	o := orm.NewOrm()
	item := Item{Id:id}
	itemInfo := ItemInfo{}
	err := o.Read(&item)
	if err != nil {
		return itemInfo,err
	}

	return itemInfo,err
}