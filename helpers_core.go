// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"strconv"
)

func SetCreateVnicOptions(rawCreateVnicDetails interface{}) (vnicOpts *baremetal.CreateVnicOptions, err error) {
	vnic := rawCreateVnicDetails.(map[string]interface{})

	vnicOpts = &baremetal.CreateVnicOptions{}
	vnicOpts.SubnetID = vnic["subnet_id"].(string)

	displayName := vnic["display_name"]
	if displayName != nil {
		vnicOpts.DisplayName = displayName.(string)
	}

	hostnameLabel := vnic["hostname_label"]
	if hostnameLabel != nil {
		vnicOpts.HostnameLabel = hostnameLabel.(string)
	}

	privateIp := vnic["private_ip"]
	if privateIp != nil {
		vnicOpts.PrivateIp = privateIp.(string)
	}

	// Work around for tf bug https://github.com/hashicorp/terraform/issues/13512.
	// For bool values that are nested in maps, if the value is set to true/false then
	// it will appear here as "1"/"0". However, if the value is set to "true"/"false"
	// then it will appear here as "true"/"false". ParseBool() handles both of these cases.
	assignPublicIp := vnic["assign_public_ip"]
	if assignPublicIp != nil {
		vnicOpts.AssignPublicIp = new(bool)
		*vnicOpts.AssignPublicIp, err = strconv.ParseBool(assignPublicIp.(string))
	}

	skipSourceDestCheck := vnic["skip_source_dest_check"]
	if skipSourceDestCheck != nil {
		vnicOpts.SkipSourceDestCheck = new(bool)
		*vnicOpts.SkipSourceDestCheck, err = strconv.ParseBool(skipSourceDestCheck.(string))
	}

	return
}
