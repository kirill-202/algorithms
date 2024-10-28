package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
	"sort"
	"reflect"
	"strings"
	"os/signal"
	"syscall"
)

const logFilePath string = "./applogs.log"
const checkInterval = 10 * time.Second



type dirEntrySorter []fs.DirEntry

func (d dirEntrySorter) Len() int {
    return len(d)
}

func (d dirEntrySorter) Less(i, j int) bool {
    return d[i].Name() < d[j].Name()
}

func (d dirEntrySorter) Swap(i, j int) {
    d[i], d[j] = d[j], d[i]
}



func GetEntriesNames(entries []fs.DirEntry) string {
	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	nameString := "[ " + strings.Join(names, " | ") + " ]"
	return nameString 
}

func MonitorDicrectory(dir string, dirEntr map[string][]fs.DirEntry, logger *log.Logger) {
	fmt.Printf("Monitoring for %s is running\n", dir)

	for {
		
		entries, _ := os.ReadDir(dir)
		oldEntries := dirEntr[dir]

		sort.Sort(dirEntrySorter(entries))
		sort.Sort(dirEntrySorter(oldEntries))

		if !reflect.DeepEqual(entries, oldEntries) {
			logger.Printf("dir content is updted from %v to %v\n", 
			GetEntriesNames(oldEntries),
			GetEntriesNames(entries),
		 	)
			dirEntr[dir] = entries
		}

		time.Sleep(checkInterval)
	}
}


func GetEntries(dirs []string, readLogger *log.Logger) map[string][]fs.DirEntry {

	dirToEntries := make(map[string][]fs.DirEntry)
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir); if err != nil {
				readLogger.Printf("can't read %v entries, skip\n", dir)
				
		} else {
			dirToEntries[dir] = entries
		}
	}
	return dirToEntries
}

func main() {

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("Failed to open file: %v\n", err)
        return
    }
    defer logFile.Close()

	dirReadLogger := log.New(logFile, "[ERROR READING DIR]", log.LstdFlags)
	dirWriteLogger := log.New(logFile, "[WRITE EVENT]", log.LstdFlags)
	


	if len(os.Args) < 2 {
		log.Fatalln("Please provide at least one directory")
	}
	dirs := os.Args[1:]
	dirWithEntries := GetEntries(dirs, dirReadLogger)

	for i:=0; i < len(dirs); i++ {
		go MonitorDicrectory(dirs[i], dirWithEntries, dirWriteLogger)
	}


	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("The program has finished  working")
}