package entity

type TagSet map[string]bool

func NewTagSet(tags ...string) *TagSet

func (ts *TagSet) Has(tag string) bool

func (ts *TagSet) Add(tag string)

func (ts *TagSet) Remove(tag string) bool