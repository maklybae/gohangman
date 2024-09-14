package climenu

import (
	"errors"
	"fmt"

	"github.com/eiannone/keyboard"
)

const MenuIntroLines = 3

type MenuItem string

type Menu struct {
	oneLineUserMessage string
	position           int
	menuItems          []MenuItem
}

func NewMenu(oneLineUserMessage string) *Menu {
	return &Menu{
		oneLineUserMessage: oneLineUserMessage,
		position:           0,
		menuItems:          make([]MenuItem, 0),
	}
}

func (m *Menu) AddItem(label string) {
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

func (m *Menu) destroyMenu() {
	for i := 0; i < len(m.menuItems)+MenuIntroLines; i++ {
		fmt.Printf("\033[1A")
		fmt.Printf("\033[K")
	}
}

func (m *Menu) clearMenu() {
	fmt.Printf("\033[%dA", len(m.menuItems))
}

func (m *Menu) drawMenu(redraw bool) {
	if redraw {
		m.clearMenu()
	}

	for i, item := range m.menuItems {
		if i == m.position {
			fmt.Printf("-> %s\n", item)
		} else {
			fmt.Printf("   %s\n", item)
		}
	}
}

func (m *Menu) RunMenu() (chosenIndex int, err error) {
	if err := keyboard.Open(); err != nil {
		return -1, fmt.Errorf("keyboard open: %w", err)
	}

	defer func() {
		if closeErr := keyboard.Close(); closeErr != nil {
			if err != nil {
				err = errors.Join(err, closeErr)
				return
			}

			err = fmt.Errorf("keyboard close: %w", closeErr)
		}
	}()

	// Hide cursor
	fmt.Printf("\033[?25l")
	defer fmt.Printf("\033[?25h")

	defer m.destroyMenu()

	fmt.Printf("%s\n", m.oneLineUserMessage)
	fmt.Printf("Use the arrow keys to navigate and press Enter to select\n")
	fmt.Printf("Press ESC to exit\n")
	m.drawMenu(false)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key { //nolint
		case keyboard.KeyArrowUp:
			m.moveUp()
		case keyboard.KeyArrowDown:
			m.moveDown()
		case keyboard.KeyEnter:
			return m.position, nil
		case keyboard.KeyEsc:
			return -1, &ExitError{}
		default:
			continue
		}

		m.drawMenu(true)
	}
}
