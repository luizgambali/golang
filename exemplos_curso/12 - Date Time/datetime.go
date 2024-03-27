package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {

	data := time.Now()
	datautc := time.Now().UTC()

	p("******* Data e hora atual *******")
	p("Data/hora atual (local): ", data)
	p("Data/hora atual (UTC): ", datautc)

	pf("Dia %d Mes %d Ano %d\n", data.Day(), data.Month(), data.Year())
	pf("Hora %d minuto %d segundo %d\n", data.Hour(), data.Minute(), data.Second())

	p("******* NOVA DATA *******")
	newdate := time.Date(2021, time.August, 22, 11, 45, 0, 0, time.Local)
	pf("%s\n", newdate.Local())

	p("******* CONVERSAO DE FORMATO *******")
	p("Time: ", data.Format("15:04:05"))
	p("Date:", data.Format("Jan 2, 2006"))
	p("Timestamp:", data.Format(time.Stamp))
	p("ANSIC:", data.Format(time.ANSIC))
	p("UnixDate:", data.Format(time.UnixDate))
	p("Kitchen:", data.Format(time.Kitchen))

	p("SQL: ", data.Format("2006-01-02"))
	p("SQL Date Time: ", data.Format("2006-01-02 3:04:05"))
	p("Pt/BR: ", data.Format("02/01/2006"))
	p("Pt/BR: ", data.Format("02/01/2006 3:4:5 PM"))
	p("SQL: ", data.Format("2006-01-02 3:4:5"))
	data = data.AddDate(0, 0, 3)
	data = data.Add(time.Hour * 3)
	p("Pt/BR: ", data.Format("02/01/2006 15:04:05"))
}
