// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import "github.com/MustWin/baremetal-sdk-go"

func SetCreateVnicOptions(rawCreateVnicDetails interface{}) (vnicOpts *baremetal.CreateVnicOptions) {
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

	//todo: work around for tf bug https://github.com/hashicorp/terraform/issues/13512
	assignPublicIp := vnic["assign_public_ip"]
	if assignPublicIp != nil {
		vnicOpts.AssignPublicIp = new(bool)
		*vnicOpts.AssignPublicIp = assignPublicIp.(string) == "1"
	}

	return
}
