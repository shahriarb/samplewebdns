package main

import (
	"io"
	"net"
	"net/http"
	//	"os/exec"
)

func hello(w http.ResponseWriter, r *http.Request) {
	/*	var (
			cmdOut []byte
			err    error
		)
	*/

	hostname := r.URL.Query().Get("hostname")
	if len(hostname) != 0 {
		/*		cmdName := "dig"
				cmdArgs := []string{"@localhost", hostname}

				if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
					io.WriteString(w, err.Error())
				} else {
					io.WriteString(w, string(cmdOut))
				}*/

		if ip, err := net.ResolveIPAddr("ip4", hostname); err != nil {
			io.WriteString(w, err.Error())
		} else {
			io.WriteString(w, ip.String())
		}

	} else {
		io.WriteString(w, "There is no hostname !")
	}

}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
