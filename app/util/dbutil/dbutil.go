package dbutil


import (
		"database/sql"
		"gopkg.in/gorp.v1"
		_ "github.com/mattn/go-sqlite3"
		"log"
		"github.com/petetheman79/idnumber/app/util/idnumberutil"
)

func InsertID(idnumber idnumberutil.ID) {
	// initialize the DbMap
    dbmap := initDb()
    defer dbmap.Db.Close()

    // insert rows - auto increment PKs will be set properly after the insert
    err := dbmap.Insert(&idnumber)
    checkErr(err, "Insert failed")
}

func InsertIDList(idnumberlist []idnumberutil.ID) {
	// initialize the DbMap
    dbmap := initDb()
    defer dbmap.Db.Close()

    // insert rows - auto increment PKs will be set properly after the insert
	for i := 0; i < len(idnumberlist); i++ {
		id := idnumberlist[i]
		err := dbmap.Insert(&id)
		checkErr(err, "Insert failed")
	}
}

func GetIDNumberList() []idnumberutil.ID {
	// initialize the DbMap
    dbmap := initDb()
    defer dbmap.Db.Close()
	
	var err error
	var idnumberlist []idnumberutil.ID
    _, err = dbmap.Select(&idnumberlist, "select * from ids")
    checkErr(err, "Select failed")
	
	return idnumberlist
}


func initDb() *gorp.DbMap {
    // connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish
    db, err := sql.Open("sqlite3", "/tmp/id_db.bin")
    checkErr(err, "sql.Open failed")

    // construct a gorp DbMap
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

    // add a table, setting the table name to 'posts' and
    // specifying that the Id property is an auto incrementing PK
    dbmap.AddTableWithName(idnumberutil.ID{}, "ids").SetKeys(true, "Id")

    // create the table. in a production system you'd generally
    // use a migration tool, or create the tables via scripts
    err = dbmap.CreateTablesIfNotExists()
    checkErr(err, "Create tables failed")

    return dbmap
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}