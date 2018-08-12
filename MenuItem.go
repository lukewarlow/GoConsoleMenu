package GoConsoleMenu

type MenuItem struct {
	Id int
	Description string
	Menu *Menu
	Action func()
	IsVisible bool
	IsExitOption bool
}

func NewItemInternal(id int) MenuItem {
	m := MenuItem{Id: id}
	return m
}

func NewActionItem(id int, description string, action func()) MenuItem {
	m := MenuItem{Id: id, Description: description, Action: action, IsVisible: true}
	return m
}

func NewSubmenuItem(id int, description string, menu *Menu) MenuItem {
	m := MenuItem{Id: id, Description: description, Menu: menu, IsVisible: true}
	return m
}

func (m MenuItem) Show() MenuItem {
	m.IsVisible = true
	return m
}

func (m MenuItem) Hide() MenuItem {
	m.IsVisible = false
	return m
}

func (m MenuItem) SetAsExitOption() MenuItem {
	m.IsExitOption = true
	return m
}


func (m MenuItem) Run() bool {
	if m.Action != nil {
		m.Action()
	} else {
		m.Menu.Display()
	}
	return !m.IsExitOption
}

