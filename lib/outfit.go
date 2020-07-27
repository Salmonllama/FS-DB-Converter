package lib

import "fmt"

type Outfit struct {
	Id           string `gorethink:"id"`
	Link         string `gorethink:"link"`
	Submitter    string `gorethink:"submitter"`
	Tag          string `gorethink:"tag"`
	Meta         string
	Created      int64
	Updated      int64
	Deleted      bool
	Featured     bool
	DisplayCount int
	DeleteHash   string
}

func (outfit *Outfit) String() string {
	return fmt.Sprintf("Outfit: ["+
		"id: %v, "+
		"link %v, "+
		"submitter: %v, "+
		"tag: %v, "+
		"meta: %v, "+
		"created: %v"+
		"updated: %v, "+
		"deleted: %v, "+
		"featured: %v, "+
		"display count: %v, "+
		"delete hash: %v]",
		outfit.Id,
		outfit.Link,
		outfit.Submitter,
		outfit.Tag,
		outfit.Meta,
		outfit.Created,
		outfit.Updated,
		outfit.Deleted,
		outfit.Featured,
		outfit.DisplayCount,
		outfit.DeleteHash)
}
