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
	"flag"
	)

type macaddr struct {
	name string
	mac  string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://vallder:30061997@192.168.0.13/mac_addr")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal()
	}
}

func main() {
	add := flag.Bool("add", false , "Добавить mac-адрес")
	list := flag.Bool("ls", false, "Отобразить mac-адреса в базе")
	flag.Parse()
	if *add  == true {
		add_mac()
	}
	if *list == true{
		select_all()
	}
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

func select_all() {
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
	for _, mac := range macs {
		fmt.Println("\n", mac.name, strings.TrimSpace(mac.mac))
	}
}
func add_mac() {
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
			log.Println("Mac already record")
		}
		}

	}

