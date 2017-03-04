package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"xd5d.com/util"
)

//FileModifyDef
type FileModifyDef struct {
	FileName string
	Op       string
}

//ModifyDef
type ModifyDef struct {
	TotalCount int
	Files      []FileModifyDef
}

var (
	StaticFileType = [...]string{".php", ".htm", ".html", ".js", ".css"}
	ChangeFiles    []FileModifyDef

	MmonitorDirNum int
	WatcherObj     *fsnotify.Watcher
)

func init() {
	var err error
	WatcherObj, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("create watcher err, ", err)
	}
	runtime.SetFinalizer(WatcherObj, (*fsnotify.Watcher).Close)
}

func checkAddFileName(name string, op string, checkDir bool) {
	ext := filepath.Ext(name)
	found := false
	for _, t := range StaticFileType {
		if strings.EqualFold(t, ext) {
			found = true
			break
		}
	}

	if checkDir {
		if util.IsDir(name) {
			found = true
		}
	}

	if found {
		in_list := false
		for index, _ := range ChangeFiles {
			if strings.EqualFold(ChangeFiles[index].FileName, name) && strings.EqualFold(ChangeFiles[index].Op, op) {
				in_list = true
				break
			}
		}
		if !in_list {
			log.Println("deteck file change, ", name, op)

			ChangeFiles = append(ChangeFiles, FileModifyDef{name, op})
		}
	}
}

func inExcludeDir(dir string) bool {
	in := false

	for _, v := range MainCfg.ExcludeDir {
		if strings.HasPrefix(dir, v) {
			in = true
			break
		}
	}

	return in
}

func walkDir(dirPth string) (files []string, err error) {
	files = make([]string, 0, 30)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Println("walk err", err)
		} else {
			if fi.IsDir() {
				if !inExcludeDir(filename) {
					files = append(files, filename)
				} else {
					log.Println("file in exclude dir, ", filename)
				}
			}
		}
		return nil
	})
	return files, err
}

func addWatcher(dir string) error {

	filenames, err := walkDir(dir)
	if err != nil {
		log.Println("walk dir err", err)
		return err
	} else {
		for _, name := range filenames {
			//log.Println("walk dir: ", name)
			err = WatcherObj.Add(name)
			if err != nil {
				log.Println("watch err, ", err)
				return err
			} else {
				MmonitorDirNum++
			}
		}
	}

	return nil
}

func fileMonitorLoop() {
	timerMail := time.NewTicker(time.Duration(MainCfg.IntervalTime) * time.Second)

	for {
		select {
		case event := <-WatcherObj.Events:
			if event.Op&fsnotify.Create == fsnotify.Create {
				if util.IsDir(event.Name) {
					log.Println("detect new dir created: ", event.Name)
					checkAddFileName(event.Name, "Create", true)

					if !inExcludeDir(event.Name) {
						log.Println("add new dir to monitor list, ", event.Name)

						err := WatcherObj.Add(event.Name)
						if err != nil {
							log.Println("watch new dir err, ", err)
						} else {
							MmonitorDirNum++
						}
					}
				} else {
					checkAddFileName(event.Name, "Create", false)

					//log.Println("create file:", event.Name)
				}
			} else if event.Op&fsnotify.Rename == fsnotify.Rename {
				checkAddFileName(event.Name, "Rename", false)

				//log.Println("rename file:", event.Name)
			} else if event.Op&fsnotify.Write == fsnotify.Write {
				checkAddFileName(event.Name, "Modify", false)

				//log.Println("modified file:", event.Name)
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				checkAddFileName(event.Name, "Del", false)

				//log.Println("del file:", event.Name)
			}
		case err := <-WatcherObj.Errors:
			log.Println("event error:", err)
		case <-timerMail.C:
			if len(ChangeFiles) > 0 {

				sendAlarmMail(ChangeFiles)

				ChangeFiles = []FileModifyDef{}
			}
		}
	}
}

func addFileMonitor() error {
	var err error
	for _, dir := range MainCfg.InlcudeDir {
		log.Println("include dir, ", dir)
		err = addWatcher(dir)

		if err != nil {
			log.Fatal("add monitor err, ", err)
			break
		}
	}
	return err
}
