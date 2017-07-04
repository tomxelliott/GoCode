package shopping

import(
	"shopping/db"	
	// "github.com/mattn/go-sqlite3" 
	// This is the package location in $GOPATH/
	// would import a folder that we had imported using:
	// "go get github.com/mattn/go-sqlite3"
)

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
