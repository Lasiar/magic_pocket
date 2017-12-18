package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net"
	"os"
	"os/user"
	"strings"
	"magic_pocket/lib"
	"magic_pocket/models"
	"github.com/olekukonko/tablewriter"
)



func init() {
	lib.Db = models.NewDb()
}

func main() {
	add := flag.Bool("add", false, "Добавить mac-адрес")
	list := flag.Bool("ls", false, "Отобразить mac-адреса в базе")
	flag.Parse()
	if *add == *list {
		flag.PrintDefaults()
	}
	if *add == true {
		addMac()
	}
	if *list == true {
		selectAll()
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

func selectAll() {
	var (
		macsAddrs []lib.MacAddr
		macAddr lib.MacAddr
	)
	rows, err := lib.Db.Query("SELECT * FROM mac")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&macAddr.Name, &macAddr.Mac)
		if err != nil {
			log.Fatal(err)
		}
		macsAddrs = append(macsAddrs, macAddr)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, mac := range macsAddrs {
		if mac.Mac == "" {
			continue
		}
	}


	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeaderLine(false)
	var v1 []string
	for _, v := range macsAddrs {
		if v.Mac == "" {
			continue
		}
		v1 = nil
		v1 = append(v1, strings.TrimSpace(v.Mac))
		v1 = append(v1, strings.TrimSpace(v.Name))
		table.Append(v1)
	}
	table.Render() // Send output
}
func addMac() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	if user.Name == "" {
		fmt.Print("your_name-> ")
		user.Name, _ = reader.ReadString('\n')
	}
	ArrMac, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}
	SqlStatement := "INSERT INTO mac (name, mac) VALUES ($1,$2)"
	for _, MacUser := range ArrMac {
		_, err := lib.Db.Exec(SqlStatement, user.Name, MacUser)
		if err != nil {
			log.Println("Mac already record")
		}
	}

}
