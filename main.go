package main

import (
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "gopkg.in/inconshreveable/log15.v2"
)

var Log = log.New()

func hello(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			Log.Info("Unhandled Error!")
			err, ok := r.(error)
			if ok {
				Log.Info("FATAL: Unhandled Error! " + err.Error())
			} else {
				Log.Info("FATAL: Unhandled Error!")
			}
			os.Exit(2)
			return
		}
	}()

	/*	var (
			cmdOut []byte
			err    error
		)
	*/
	Log.Info("Hello Request: START")

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
			Log.Info("Hello Request: Error: " + err.Error())
		} else {
			io.WriteString(w, ip.String())
			Log.Info("Hello Request:" + ip.String())
		}

	} else {
		io.WriteString(w, "There is no hostname !")
		Log.Info("Hello Request: There is no hostname!")
	}

	Log.Info("Hello Request: END")

}

func main() {
	Log.SetHandler(log.StdoutHandler)

	// listen to signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		if r := recover(); r != nil {
			Log.Info("Unhandled Error!")
			err, ok := r.(error)
			if ok {
				Log.Info("FATAL: Unhandled Error! " + err.Error())
			} else {
				Log.Info("FATAL: Unhandled Error!")
			}
			os.Exit(2)
			return
		}
	}()

	go func() {
		for {
			select {
			case sig := <-signalChan:
				Log.Info("Signal! (" + sig.String() + ")")
				return
			}
		}
	}()

	Log.Info("Web started!")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)

	Log.Info("Web stopping!")
}
