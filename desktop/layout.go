package desktop

import "github.com/leukipp/cortile/store"

type Layout interface {
	Apply()
	AddClient(c *store.Client)
	RemoveClient(c *store.Client)
	MakeMaster(c *store.Client)
	SwapClient(c1 *store.Client, c2 *store.Client)
	NextClient() *store.Client
	PreviousClient() *store.Client
	IncreaseMaster()
	DecreaseMaster()
	IncreaseSlave()
	DecreaseSlave()
	IncreaseProportion()
	DecreaseProportion()
	UpdateProportions(c *store.Client, d *store.Directions)
	GetManager() *store.Manager
	GetName() string
}
