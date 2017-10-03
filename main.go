package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net"
	"os"
	"os/user"
	"strings"
)

type macaddr struct {
	name string
	mac  string
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			fmt.Println(a)
			as = append(as, a)
		}
	}
	return as, nil
}

func main() {
	db, err := sql.Open("postgres", "postgres://vallder:30061997@192.168.0.174/mac_addr")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM mac")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	macs := make([]*macaddr, 0)
	for rows.Next() {
		mac := new(macaddr)
		err := rows.Scan(&mac.name, &mac.mac)
		if err != nil {
			log.Fatal(err)
		}
		macs = append(macs, mac)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	if user.Name == "" {
		fmt.Print("-> ")
		user.Name, _ = reader.ReadString('\n')
		// convert CRLF to LF
		user.Name = strings.Replace(user.Name, "\n", "", -1)

	}
	ArrMac, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}
	SqlStatement := "INSERT INTO mac (name, mac) VALUES ($1,$2)"
	for _, MacUser := range ArrMac {
		_, err := db.Exec(SqlStatement, user.Name, MacUser)
		if err != nil {
			log.Println(MacUser, "mac is already recorded")
		}
	}
	for _, mac := range macs {
		/*if len(strings.TrimSpace(mac.name)) == 14 {
			mac.name = "unknow"
		}*/
		fmt.Println("\n", mac.name, strings.TrimSpace(mac.mac))
		//fmt.Println(len(strings.TrimSpace(mac.name)))
	}

}
