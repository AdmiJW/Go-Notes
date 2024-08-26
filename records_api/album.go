package records_api

// Represents data about a record album
type album struct {
	ID		string 		`json:"id"`
	Title	string 		`json:"title" binding:"required,max=10"`
	Artist	string 		`json:"artist" binding:"required,max=10"`
	Price	float64 	`json:"price" binding:"required,gt=0"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}