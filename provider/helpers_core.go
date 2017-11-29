// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"
)

var createVnicDetailsSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"assign_public_ip": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
			ForceNew: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"hostname_label": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"private_ip": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"skip_source_dest_check": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"subnet_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	},
}

// vnicDetailsList is assumed to be non-nil and non-empty.
func SetCreateVnicOptions(vnicDetailsList []interface{}) (vnicOpts *baremetal.CreateVnicOptions) {
	vnic := vnicDetailsList[0].(map[string]interface{})

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

	assignPublicIp := vnic["assign_public_ip"]
	if assignPublicIp != nil {
		vnicOpts.AssignPublicIp = new(bool)
		*vnicOpts.AssignPublicIp = assignPublicIp.(bool)
	}

	skipSourceDestCheck := vnic["skip_source_dest_check"]
	if skipSourceDestCheck != nil {
		vnicOpts.SkipSourceDestCheck = new(bool)
		*vnicOpts.SkipSourceDestCheck = skipSourceDestCheck.(bool)
	}

	return
}

// vnicDetailsList is assumed to be non-nil and non-empty.
func SetUpdateVnicOptions(vnicDetailsList []interface{}) (vnicOpts *baremetal.UpdateVnicOptions) {
	vnic := vnicDetailsList[0].(map[string]interface{})
	vnicOpts = &baremetal.UpdateVnicOptions{}

	displayName := vnic["display_name"]
	if displayName != nil {
		vnicOpts.DisplayName = displayName.(string)
	}

	hostnameLabel := vnic["hostname_label"]
	if hostnameLabel != nil {
		vnicOpts.HostnameLabel = hostnameLabel.(string)
	}

	skipSourceDestCheck := vnic["skip_source_dest_check"]
	if skipSourceDestCheck != nil {
		vnicOpts.SkipSourceDestCheck = new(bool)
		*vnicOpts.SkipSourceDestCheck = skipSourceDestCheck.(bool)
	}

	return
}

func RefreshCreateVnicDetails(resourceData *schema.ResourceData, vnic *baremetal.Vnic) {
	vnicDetails := make(map[string]interface{})
	vnicDetails["subnet_id"] = vnic.SubnetID
	vnicDetails["assign_public_ip"] = len(vnic.PublicIPAddress) > 0
	vnicDetails["display_name"] = vnic.DisplayName
	vnicDetails["hostname_label"] = vnic.HostnameLabel
	vnicDetails["private_ip"] = vnic.PrivateIPAddress
	vnicDetails["skip_source_dest_check"] = vnic.SkipSourceDestCheck
	resourceData.Set("create_vnic_details", []interface{}{vnicDetails})
}
