// Backup database
// Avoid directly enter your user name and password in scripts

package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/jearry/golang/util"
)

var (
	dbUser     = "user"
	dbPasspwd  = "passwd"
	dbHost     = "db.host.com"
	dbDatabase = "dbname"

	zipPasswd = "passwdasdf%^&"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	file, err := os.OpenFile("dbbak.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	}
}

func main() {

	currentDate := util.GetCurrentDate()
	sqlFilename := currentDate + ".sql"
	zipFilename := currentDate + ".7z"

	stdinName := "-si" + sqlFilename
	argHost := "-h" + dbHost
	argUser := "-u" + dbUser
	argPass := "-p" + dbPasspwd

	cmdDump := exec.Command("mysqldump.exe", argHost, argUser, argPass, dbDatabase)

	argZipPwd := "-p" + zipPasswd
	cmdZip := exec.Command("7z.exe", argZipPwd, "-mhe", "-r", stdinName, "a", zipFilename)

	pipe, err := cmdZip.StdinPipe()

	if err != nil {
		log.Fatal("cmd pipe error, ", err)
	}
	cmdDump.Stdout = pipe

	if err = cmdDump.Start(); err != nil {
		log.Fatal("cmd dump error, ", err)
	}

	if err = cmdZip.Start(); err != nil {
		log.Fatal("cmd zip error, ", err)
	}

	cmdDump.Wait()

	log.Println("dump sucess")
}
