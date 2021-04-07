package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"k8s.io/klog/v2"
)

func mainRun(c *cli.Context) error {
	klog.V(0).Info("Hello")
	klog.V(1).Info("Hello1")
	klog.V(2).Info("Hello2")
	klog.V(3).Info("Hello3")
	klog.V(4).Info("Hello4")

	return cli.Exit("done", 0)
}

func main() {
	flags := []cli.Flag{
		altsrc.NewIntFlag(&cli.IntFlag{Name: "loglevel", Aliases: []string{"l"}, Value: 3}),
		&cli.StringFlag{Name: "config", Value: "eg_config.yaml", HasBeenSet: true},
		&cli.StringFlag{Name: "global-config", Value: "eg_global-config.yaml", HasBeenSet: true},
	}

	before := func(c *cli.Context) error {
		// Command line flags always overwrite configuration files
		first := altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))
		err := first(c)
		if err != nil {
			klog.Error(err)
		}
		// The second config map will not overwrite the first
		second := altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("global-config"))
		err = second(c)
		if err != nil {
			klog.Error(err)
		}

		fs := flag.NewFlagSet("", flag.PanicOnError)
		klog.InitFlags(fs)
		return fs.Set("v", strconv.Itoa(c.Int("loglevel")))
	}

	app := &cli.App{
		Name:  "klog-cli",
		Usage: "An example of setting up klog with a cli application",

		Flags:  flags,
		Before: before,
		Action: mainRun,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
