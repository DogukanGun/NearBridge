package db

/*
DatabaseHandler object for general database
*/
type DatabaseHandler struct {
	TargetDB    string
	TargetTable string
	DbUri       string
	DbPassword  string
	DbClient    interface{}
	DbSession   interface{}
}
