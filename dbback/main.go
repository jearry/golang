package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/jearry/golang/util"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	file, err := os.OpenFile("dbbak.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	}
}

func main() {

	current_date := util.GetCurrentDate()
	sqlfilename := current_date + ".sql"
	zipfilename := current_date + ".7z"

	stdin_name := "-si" + sqlfilename

	cmd_dump := exec.Command("mysqldump.exe", "-hxx.abc.com", "-uuser", "-ppasswd", "dbname")
	cmd_7z := exec.Command("7z.exe", "-ppasswd", "-mhe", "-r", stdin_name, "a", zipfilename, sqlfilename)

	pipe, err := cmd_7z.StdinPipe()

	if err != nil {
		log.Fatal("cmd pipe error, ", err)
	}
	cmd_dump.Stdout = pipe

	if err = cmd_dump.Start(); err != nil {
		log.Fatal("cmd dump error, ", err)
	}

	if err = cmd_7z.Start(); err != nil {
		log.Fatal("cmd zip error, ", err)
	}

	cmd_dump.Wait()

	log.Println("dump sucess")

	os.Remove(sqlfilename)
}
