package lex

import (
	"fmt"
)

// ItemType describes the type of a LexItem
type ItemType int

// TypeNames contains the name of reach ItemType. This is used for
// printing the values out in a human readable format
var TypeNames = make(map[ItemType]string)

const (
	// ItemEOF is emiteed upon EOF
	ItemEOF ItemType = iota
	// ItemError is emitted upon Error
	ItemError
	// ItemDefaultMax is used as marker for your own ItemType.
	// Start your types from this + 1
	ItemDefaultMax
)

func init() {
	TypeNames[ItemEOF] = "EOF"
	TypeNames[ItemError] = "Error"
	TypeNames[ItemDefaultMax] = "Special (DefaultMax)"
}

func (t ItemType) String() string {
	name, ok := TypeNames[t]
	if !ok {
		return fmt.Sprintf("Unknown Item (%d)", t)
	}
	return name
}

// LexItem defines the interface for items emitted by the Lexer
type LexItem interface {
	Type() ItemType
	Pos() int
	Line() int
	Value() string
}

// Item is the struct that gets generated upon finding *something*
type Item struct {
	typ  ItemType
	pos  int
	line int
	val  string
}

// NewItem creates a new Item
func NewItem(t ItemType, pos int, line int, v string) Item {
	return Item{t, pos, line, v}
}

// Type returns the associated ItemType
func (l Item) Type() ItemType {
	return l.typ
}

// Pos returns the associated position
func (l Item) Pos() int {
	return l.pos
}

// Line returns the line number in which this occurred
func (l Item) Line() int {
	return l.line
}

// Value returns the associated text value
func (l Item) Value() string {
	return l.val
}

// String returns the string representation of the Item
func (l Item) String() string {
	return fmt.Sprintf("%s (%q)", l.typ, l.val)
}
