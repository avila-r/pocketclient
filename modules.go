package pocketclient

type Module struct{}

type (
	ModuleAdmin       Module
	ModuleCollections Module
)

var (
	Admin       = &ModuleAdmin{}
	Collections = &ModuleCollections{}
)

func (c *PocketClient) Admin() *ModuleAdmin {
	return Admin
}

func (c *PocketClient) Collections() *ModuleCollections {
	return Collections
}
