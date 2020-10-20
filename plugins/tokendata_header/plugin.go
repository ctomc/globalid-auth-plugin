package main

import (
	"github.com/solo-io/ext-auth-plugins/api"
	impl "gitlab.com/globalid/experiments/globalid-auth-plugin/plugins/tokendata_header/pkg"
)

func main() {}

// Compile-time assertion
var _ api.ExtAuthPlugin = new(impl.TokendataPlugin)

// This is the exported symbol that Gloo will look for.
//noinspection GoUnusedGlobalVariable
var Plugin impl.TokendataPlugin
