package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type Item struct {
	Id          int `json:"item_id"`
	Title 	string `json:"item_name"`
	Description string `json:"item_description"`
	UserId	int
	Password	string `json:"password"`
	Type	int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"addtime"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

type ItemInfo struct {
	Id          int `json:"item_id"`
	Title 	string `json:"item_domain"`
	Password string `json:"password"`
	IsArchived string `json:"is_archived"`
	DefaultPageId	string `json:"default_page_id"`
	DefaultCatId2	string `json:"default_cat_id2"`
	DefaultCatId3	string `json:"default_cat_id3"`
	UnreadCount		UnreadCount `json:"unread_count"`
	ItemType		int		`json:"item_type"`
	IsLogin bool `json:"is_login"`
	ItemPermn bool `json:"ItemPermn"`
	ItemCreator bool `json:"ItemCreator"`
	Menu 	Menu `json:"menu"`
}
type Menu struct {
	Page []*Page `json:"pages"`
	Catalogs []*Catalogs `json:"catalogs"`
}

type UnreadCount struct {

}
func (this *Item) GetItemInfo() (ItemInfo) {
	var itemInfo ItemInfo
	itemInfo.Id = this.Id
	itemInfo.ItemType = this.Type
	itemInfo.IsLogin = false
	itemInfo.ItemPermn = false
	itemInfo.ItemCreator = false
	itemInfo.Menu = itemInfo.GetMenu()
	return itemInfo
}

func (this *ItemInfo) GetMenu()(Menu) {
	var menu Menu
	menu.Page = GetPagesByItemId(this.Id)
	menu.Catalogs = GetCatalogsByItemId(this.Id)
	return menu
}


func (this *Item) Create() (int,error) {
	o := orm.NewOrm()

	id ,err := o.Insert(this)
	if err != nil {
		beego.Error(err.Error())
		err = fmt.Errorf("%s", "新增数据失败")
	}

	return int(id),err
}

func GetMyItem(uid int) ([]*Item) {
	o := orm.NewOrm()
	var item []*Item
	num,err :=o.QueryTable("item").Filter("user_id", uid).All(&item)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,uid)

	return item
}

func GetOneItem(id int) (bool,Item) {
	o := orm.NewOrm()
	item := Item{Id:id}

	err := o.Read(&item)

	if err == nil {
		return true,item
	} else {
		return false,item
	}

}
func GetMyItemInfo(id int) (ItemInfo,error) {
	o := orm.NewOrm()
	item := Item{Id:id}
	itemInfo := ItemInfo{}
	err := o.Read(&item)
	if err != nil {
		return itemInfo,err
	}

	return itemInfo,err
}

func (this *Item) Delete() (error) {
	o := orm.NewOrm()


	_,err := o.Delete(this)

	return err

}

func (this *Item) Save() {
	o := orm.NewOrm()

	//判断是新建还是更新
	if this.Id == 0 {
		_ ,err := o.Insert(this)
		if err != nil {
			beego.Error(err.Error())
			err = fmt.Errorf("%s", "新增数据失败")
		}
	} else {
		item := Item{Id: this.Id}
		if o.Read(&item) == nil {
			item = *this
			item.Id = this.Id
			if num, err := o.Update(&item); err == nil {
				fmt.Println(num)
			}
		}
	}
}