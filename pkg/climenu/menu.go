package climenu

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type Menu struct {
	userMessage string
	position    int
	menuItems   []MenuItem
}

type MenuItem string

func NewMenu(userMessage string, position int) *Menu {
	return &Menu{
		userMessage: userMessage,
		position:    position,
		menuItems:   make([]MenuItem, 0),
	}
}

func (m *Menu) AddItem(tag, label string) {
	m.menuItems = append(m.menuItems, MenuItem(label))
}

func (m *Menu) moveUp() {
	m.position--
	m.position = (m.position + len(m.menuItems)) % len(m.menuItems)
}

func (m *Menu) moveDown() {
	m.position++
	m.position = (m.position + len(m.menuItems)) % len(m.menuItems)
}

func (m *Menu) drawMenu(redraw bool) {
	if redraw {
		fmt.Printf("\033[%dA", len(m.menuItems)-1)
	}
	for i, item := range m.menuItems {
		if i == m.position {
			fmt.Printf("-> %s", item)
		} else {
			fmt.Printf("   %s", item)
		}
		if i < len(m.menuItems)-1 {
			fmt.Print("\n")
		}
	}

}

func (m *Menu) RunMenu() (chosenIndex int) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	fmt.Printf("%s\n", m.userMessage)
	fmt.Printf("Use the arrow keys to navigate and press Enter to select\n")
	fmt.Printf("Press ESC to exit\n")
	m.drawMenu(false)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch key {
		case keyboard.KeyArrowUp:
			m.moveUp()
		case keyboard.KeyArrowDown:
			m.moveDown()
		case keyboard.KeyEnter:
			return m.position
		case keyboard.KeyEsc:
			return -1
		}
		m.drawMenu(true)
	}
}
