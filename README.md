
# Go Console Menu
![license](https://img.shields.io/hexpm/l/plug.svg)

This library provides a way to quickly create the menu for your GoLang console app.

##  Overview

### Classes

#### Menu
This is the type that you use to make your menus.
It has two constructors one for if the menu items can be updated (shown or hidden) and another for if they can't. 
These should be called like this: 

`GoConsoleMenu.NewMenu(title)`

Or

`GoConsoleMenu.NewUpdatableMenu(title, updateMenuItemsFunc)`

##### Methods
- `Display()` this starts this menu. This only needs to be called on the root menu in your system, as all sub-menus are handled by this library.
- `AddMenuItem(GoConsoleMenu.NewSubmenuItem(id, description, subMenu))` this adds a submenu item to the menu. 
- `AddMenuItem(GoConsoleMenu.NewActionItem(id, description, action))` this adds a action item to the menu. 
- `AddHiddenMenuItem` this is a helper method that adds a menu item, which is then hidden.
- `ShowMenuItem(id)` this can be used to show hidden menu items, most commonly in the method above. This uses the unique id given to the menu item.
- `HideMenuItem(id)` this can be used to hide menu items.

#### MenuItem
This is the class used to define items for the menus in your system. 
It has two constructors one for if the item is a sub menu and another for if its an action. 
These should be called like this: 

`GoConsoleMenu.NewSubmenuItem(id, description, subMenu)`

Or

`GoConsoleMenu.NewActionItem(id, description, action)`

##### Methods
- `Hide()` which is used on menu items, to hide them from the list.
- `Show()` which is used on hidden menu items, to show them in the list.
- `SetAsExitOption()` which is used to set menu items as the exit option for a menu, either going to the parent menu, or exiting the application.

## Example
```Go
func main() {
	mainMenu := GoConsoleMenu.NewMenu("Welcome to the main menu.", func(menu *GoConsoleMenu.Menu) {})
	mainMenu.AddMenuItem(GoConsoleMenu.NewActionItem(100, "Exit menu", func() {}).SetAsExitOption())
	mainMenu.AddMenuItem(GoConsoleMenu.NewActionItem(101, "Print Hello World", func() 
	{
		fmt.Println("Hello World!")
	}))
	mainMenu.Display()
}
```

#### Output
```text
Welcome to the main menu
0. Exit menu
1. Print Hello World
Select option: 1
Hello World!

Welcome to the main menu
0. Exit menu
1. Print Hello World
Select option: 0

Process finished with exit code 0
```

Look at <a>https://gist.github.com/lukewarlow/3bb76d5aefe688c05c1546ca6bf05dbf<a> for a better example implementation of the library.
