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
	UserId	int
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
	Catalogs []Catalogs `json:"catalogs"`
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
	itemInfo.Menu = itemInfo.GetMenu()
	return itemInfo
}

func (this *ItemInfo) GetMenu()(Menu) {
	var menu Menu
	menu.Page = GetPagesByItemId(this.Id)
	menu.Catalogs = GetCatalogsByItemId(this.Id)
	return menu
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

func GetMyItem(uid int) ([]*Item) {
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