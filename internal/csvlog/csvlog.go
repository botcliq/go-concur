package csvlog

import (
	"encoding/csv"
	"log"
	"os"
	"sync"
)

var csvmutex sync.Mutex

func LogCsv(fname string, column []string) {
	csvmutex.Lock()
	defer csvmutex.Unlock()
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	w := csv.NewWriter(f)
	w.Write(column)
	w.Flush()
}
