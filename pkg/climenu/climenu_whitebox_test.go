package climenu //nolint

import (
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMenu(t *testing.T) {
	log.SetOutput(io.Discard)

	oneLineUserMessage := "Select an option:"
	menu := NewMenu(oneLineUserMessage)

	assert.NotNil(t, menu, "Menu should not be nil")
	assert.Equal(t, oneLineUserMessage, menu.oneLineUserMessage, "oneLineUserMessage should be set correctly")
	assert.Equal(t, 0, menu.position, "Initial position should be 0")
	assert.Empty(t, menu.menuItems, "Initial menuItems should be empty")
}

func TestAddItem(t *testing.T) {
	log.SetOutput(io.Discard)

	menu := NewMenu("Select an option:")

	menu.AddItem("Item 1")
	assert.Equal(t, 1, len(menu.menuItems), "Menu should have 1 item")
	assert.Equal(t, MenuItem("Item 1"), menu.menuItems[0], "First item should be 'Item 1'")

	menu.AddItem("Item 2")
	assert.Equal(t, 2, len(menu.menuItems), "Menu should have 2 items")
	assert.Equal(t, MenuItem("Item 2"), menu.menuItems[1], "Second item should be 'Item 2'")
}

func TestMoveDown(t *testing.T) {
	log.SetOutput(io.Discard)

	menu := NewMenu("Select an option:")
	menu.AddItem("Item 1")
	menu.AddItem("Item 2")
	menu.AddItem("Item 3")

	assert.Equal(t, 0, menu.position, "init")

	menu.moveDown()
	assert.Equal(t, 1, menu.position, "first")

	menu.moveDown()
	assert.Equal(t, 2, menu.position, "second")

	menu.moveDown()
	assert.Equal(t, 0, menu.position, "init")
}

func TestMoveUp(t *testing.T) {
	log.SetOutput(io.Discard)

	menu := NewMenu("Select an option:")
	menu.AddItem("Item 1")
	menu.AddItem("Item 2")
	menu.AddItem("Item 3")

	assert.Equal(t, 0, menu.position, "init")

	menu.moveUp()
	assert.Equal(t, 2, menu.position, "last")

	menu.moveUp()
	assert.Equal(t, 1, menu.position, "second last")

	menu.moveUp()
	assert.Equal(t, 0, menu.position, "init")
}
