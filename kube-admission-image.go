// writer: doumeng
// date: 202004
package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/weiqiang333/kube-admission-image/pkg/config"
	"github.com/weiqiang333/kube-admission-image/web"
)

var configs config.FlagVar

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.LUTC)
	logFile, err := os.OpenFile("logs/kube-admission-image.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.SetOutput(logFile)

	pflag.String("address", "0.0.0.0:8080", "run listening port")
	tls := pflag.Bool("tls", false, "Turn on TLS (-cert, -key)")
	cert := pflag.String("cert", "tls.crt", "Cert file for TLS")
	key := pflag.String("key", "tls.crt", "Key file for TLS")
	pflag.String("sourceDefaultPolicy", "allow", "images source default policy."+
		"\nplease configure the default reject policy carefully.\nOptions (allow|reject)")
	pflag.StringSlice("sourceAllowPolicy", []string{}, "Policy that allows images source."+
		"\nUser: --sourceAllowPolicy=terminus.io,terminus.com")
	pflag.StringSlice("sourceRejectPolicy", []string{}, "Policy that reject images source")

	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	if *tls {
		if _, err := os.Stat(*cert); os.IsNotExist(err) {
			log.Fatalf("Please check your certificate: %s", *cert)
		}
		if _, err := os.Stat(*key); os.IsNotExist(err) {
			log.Fatalf("Please check your key: %s", *key)
		}
	}
}

func main() {
	web.Web()
}
