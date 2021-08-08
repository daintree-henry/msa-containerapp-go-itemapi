package domain

type Item struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Quantity int64  `json:"quantity"`
}

var ItemList map[int64]*Item

func init() {
	ItemList = make(map[int64]*Item)
	ItemList[1] = &Item{
		Id:       1,
		Name:     "pen",
		Category: "Office",
		Quantity: 500,
	}
	ItemList[2] = &Item{
		Id:       2,
		Name:     "note",
		Category: "Office",
		Quantity: 100,
	}
	ItemList[3] = &Item{
		Id:       3,
		Name:     "cup",
		Category: "living",
		Quantity: 200,
	}
}

func (i *Item) Validate() (err string) {
	if i.Quantity < 0 {
		return "The quantity must be greater then 0."
	}
	if i.Category == "" {
		return "The category can't be blank."
	}
	if i.Name == "" {
		return "The name can't be blank."
	}
	return ""
}

func (i *Item) Create() (*Item, string) {
	if _, check := ItemList[i.Id]; check == true {
		return nil, "The item is already exists."
	}
	ItemList[i.Id] = i
	return i, ""
}

func (i *Item) GetById(id int64) (*Item, string) {
	if id <= 0 {
		return nil, "Invalid Id"
	}
	i.Id = id
	i.Name = ItemList[id].Name
	i.Category = ItemList[id].Category
	i.Quantity = ItemList[id].Quantity
	return i, ""
}

func (i *Item) List() ([]*Item, string) {
	var items []*Item
	for _, item := range ItemList {
		items = append(items, item)
	}
	return items, ""
}

func (i *Item) DeleteById(id int64) string {
	if _, check := ItemList[id]; check == false {
		return "The user not exists"
	}
	delete(ItemList, id)
	return "The user deleted"
}
