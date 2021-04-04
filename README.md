# Using klog in a urfave/cli

`k8s.io/klog/v2` will try to use the std `flags` pkg for options.
If you want to use a different flags package you need to setup some co-ordination.

The `urfave/cli/v2` package will grab the `-v` flag for the version.
Just to be on the safe side this example will create a new flag called `loglevel`
and then maps that back to the `v` flag for klog.

It also uses the `altsrc` flags so that the `loglevel` flag can be set from a configuration file if needed.
