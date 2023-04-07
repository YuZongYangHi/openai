package main

import (
	"flag"
	"fmt"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/config"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/models"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/routers"
	"k8s.io/klog/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	configPath string
)

func init() {

	if flag.CommandLine.Lookup("log_dir") != nil {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}

	klog.InitFlags(nil)

	defer klog.Flush()

	flag.StringVar(&configPath, "config", "config.yaml", "config path")
	flag.Parse()
}

func main() {

	if err := config.NewConfig(configPath); err != nil {
		klog.Errorf(err.Error())
		os.Exit(1)
	}

	if err := models.RegisterDatabase(); err != nil {
		klog.Errorf(err.Error())
		os.Exit(1)
	}

	s := &http.Server{
		Addr: fmt.Sprintf("%s", fmt.Sprintf("%s:%d", config.AppConfig().HTTP.Host,
			config.AppConfig().HTTP.Port)),
		Handler:        routers.NewRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		klog.Errorf(err.Error())
		os.Exit(3)
	}

	sics := make(chan os.Signal, 1)
	signal.Notify(sics, syscall.SIGINT, syscall.SIGTERM)
	<-sics
}
