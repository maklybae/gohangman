package draw

import (
	"makly/hangman/internal/domain"
)

const HangmanHeight = 7

var StringStates = map[domain.State]string{
	domain.Initial: `
   +---+
   |   |
       |
       |
       |
       |
=========`,
	domain.Head: `
   +---+
   |   |
   O   |
       |
       |
       |
=========`,
	domain.Body: `
   +---+
   |   |
   O   |
   |   |
       |
       |
=========`,
	domain.LeftArm: `
   +---+
   |   |
   O   |
  /|   |
       |
       |
=========`,
	domain.RightArm: `
   +---+
   |   |
   O   |
  /|\  |
       |
       |
=========`,
	domain.LeftLeg: `
   +---+
   |   |
   O   |
  /|\  |
  /    |
       |
=========`,
	domain.RightLeg: `
   +---+
   |   |
   O   |
  /|\  |
  / \  |
       |
=========`,
}
