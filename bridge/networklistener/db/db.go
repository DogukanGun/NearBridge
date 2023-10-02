package db

/*
DatabaseHandler object for general database
*/
type DatabaseHandler struct {
	targetDB    string
	targetTable string
	dbUri       string
	dbPassword  string
	dbClient    interface{}
	dbSession   interface{}
}
