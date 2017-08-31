package GetOIDSFromDB

import (
	"database/sql"
	"log"

	_ "./packages/pq-master"
)

//OIDS struct
type OIDS struct {
	ID                    int
	Value, Name, Descript string
}

// OIDSObject for query
var OIDSObject OIDS

//OIDSValue array
var OIDSValue []string

//OIDSDescript array
var OIDSDescript []string

//OIDSdbID array
var OIDSdbID []int

func GetOIDSFromDB(query string) (OIDSdbID []int, OIDSValue []string, OIDSDescript []string) {
	//db connect
	db, err := sql.Open("postgres", "user=admin password=admin dbname=oidsdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//get rows
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//scan rows
	for rows.Next() {
		err := rows.Scan(&OIDSObject.ID, &OIDSObject.Value, &OIDSObject.Descript)
		if err != nil {
			log.Fatal(err)
		}
		OIDSdbID = append(OIDSdbID, OIDSObject.ID)
		OIDSValue = append(OIDSValue, OIDSObject.Value)
		OIDSDescript = append(OIDSDescript, OIDSObject.Descript)
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return OIDSdbID, OIDSValue, OIDSDescript
}
