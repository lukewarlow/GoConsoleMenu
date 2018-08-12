package GoConsoleMenu

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type Menu struct {
	Title string
	MenuItems []MenuItem
	UpdateMenuItems func(menu *Menu)
}

func NewUpdatableMenu(title string, updateMenuItems func(menu *Menu)) *Menu {
	m := new(Menu)
	m.Title = title
	m.MenuItems = []MenuItem{}
	m.UpdateMenuItems = updateMenuItems
	return m
}

func NewMenu(title string) *Menu {
	m := new(Menu)
	m.Title = title
	m.MenuItems = []MenuItem{}
	m.UpdateMenuItems = nil
	return m
}

func (m *Menu) Display() {
	scanner := bufio.NewScanner(os.Stdin)
	repeat := true

	for repeat {
		if m.UpdateMenuItems != nil {
			m.UpdateMenuItems(m)
		}
		fmt.Println()
		fmt.Println(m.Title)
		i := 0
		for i < len(m.MenuItems) {
			if m.MenuItems[i].IsVisible {
				fmt.Printf("%d. " + m.MenuItems[i].Description + "\n", i)
			}
			i++
		}
		fmt.Print("Select Option: ")
		scanner.Scan()
		input := scanner.Text()
		itemIndex, _ := strconv.ParseInt(input, 10, 64)
		menuItem := m.MenuItems[itemIndex]
		if menuItem.IsVisible {
			repeat = menuItem.Run()
		} else {
			//throw an error
		}
		//Catch errors
	}
}

func (m *Menu) AddMenuItem(menuItem MenuItem) {
	if !contains(m.MenuItems, menuItem) {
		items := append(m.MenuItems, menuItem)
		m.MenuItems = items
	} else {
		//throw an error
	}
}

func (m *Menu) AddHiddenMenuItem(menuItem MenuItem) {
	m.AddMenuItem(menuItem.Hide())
}

func (m *Menu) ShowMenuItem(itemId int) {
	menuItem := NewItemInternal(itemId)
	itemIndex := indexOf(m.MenuItems, menuItem)
	if itemIndex > 0 {
		m.MenuItems[itemIndex].IsVisible = true
	} else {
		//throw error
	}
}

func (m *Menu) HideMenuItem(itemId int) {
	menuItem := NewItemInternal(itemId)
	itemIndex := indexOf(m.MenuItems, menuItem)
	if itemIndex > 0 {
		m.MenuItems[itemIndex].IsVisible = false
	} else {
		//throw error
	}
}

func contains(s []MenuItem, e MenuItem) bool {
	for _, a := range s {
		if a.Id == e.Id {
			return true
		}
	}
	return false
}

func indexOf(s []MenuItem, o MenuItem) int {
	for i, m := range s {
		if m.Id == o.Id {
			return i
		}
	}
	return -1
}