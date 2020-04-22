// writer: doumeng
// date: 202004
package main

import (
	"flag"
	"log"
	"os"

	"github.com/weiqiang333/kube-admission-image/web"

	"github.com/weiqiang333/kube-admission-image/pkg/config"
)

var configs config.FlagVar

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.LUTC)
	logFile, err := os.OpenFile("logs/kube-admission-image.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.SetOutput(logFile)

	flag.StringVar(&configs.Addrss, "address", "0.0.0.0:8080", "run listening port")
	flag.BoolVar(&configs.Tls, "tls", false, "Turn on TLS (-cert, -key)")
	flag.StringVar(&configs.CertFile, "cert", "tls.crt", "Cert file for TLS")
	flag.StringVar(&configs.KeyFile, "key", "tls.crt", "Key file for TLS")

	flag.Parse()
	if configs.Tls {
		if _, err := os.Stat(configs.CertFile); os.IsNotExist(err) {
			log.Fatalf("Please check your certificate: %s", configs.CertFile)
		}
		if _, err := os.Stat(configs.KeyFile); os.IsNotExist(err) {
			log.Fatalf("Please check your key: %s", configs.KeyFile)
		}
	}
}

func main() {
	web.Web(configs)
}
