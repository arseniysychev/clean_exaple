package domain

import (
	"github.com/google/uuid"
	"time"
)

type Element struct {
	Project    Project
	Identifier string
	Position   uint8
	Hidden     bool
}
type ElementSet []Element

type ElementInput struct {
	Element
	Required    bool
	Placeholder string
	//DefaultValue
}

type ElementText struct {
	ElementInput
	MaxLength    uint8
	MinLength    uint8
	DefaultValue string
}

type ElementRichText struct {
	ElementInput
	DefaultValue string
}

const (
	TypeNumberInt = iota
	TypeNumberFloat
)

type ElementNumber struct {
	ElementInput
	MinValue     float64
	MaxValue     float64
	DefaultValue float64
	Format       int8
}

const (
	TypeDateDT = iota
	TypeDateD
	TypeDateT
)

type ElementDate struct {
	ElementInput
	MinValue     time.Time
	MaxValue     time.Time
	DefaultValue time.Time
	Format       int8
}

type ElementSelect struct {
	ElementInput
	OptionSet OptionSet
}
type ElementSelectSingle struct {
	ElementSelect
	DefaultValue Option
}
type ElementSelectMultiply struct {
	ElementSelect
	DefaultValue Option
}

type ElementAssociatedRecord struct {
	ElementInput
	RelationUid        uuid.UUID
	Associated         Project
	AssociatedElements ElementSet
}
