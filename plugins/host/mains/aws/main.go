// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"

	awshp "github.com/hashicorp/boundary-plugin-host-aws/plugin"
	hp "github.com/hashicorp/boundary/sdk/plugins"
)

func main() {
	if err := hp.ServePlugin(awshp.NewAwsPlugin()); err != nil {
		fmt.Println("Error serving plugin", err)
		os.Exit(1)
	}
	os.Exit(0)
}
