package links

import (
	"log"

	database "graphql-nr/internal/pkg/db/postgres"
	"graphql-nr/internal/pkg/db/users"
)

// #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

//#2
func (link Link) Save() int64 {
	//#3
	sqlStr := `INSERT INTO "Links"("Title","Address") VALUES($1, $2) RETURNING "ID"`
	// log.Fatal("str:", sqlStr)
	stmt, err := database.Db.Prepare(sqlStr)
	if err != nil {
		log.Print(err)
	}
	//#4
	//res, err := stmt.Exec(link.Title, link.Address)
	var resID int64
	err = stmt.QueryRow(link.Title, link.Address).Scan(&resID)
	if err != nil {
		log.Print(err)
	}
	//rowsAffected, err := res.RowsAffected()
	return resID
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare(`select "ID", "Title", "Address" from "Links"`)
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Print(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
	}
	return links
}
