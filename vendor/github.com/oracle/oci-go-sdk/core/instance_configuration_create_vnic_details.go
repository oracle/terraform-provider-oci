// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// InstanceConfigurationCreateVnicDetails Please see CreateVnicDetails
type InstanceConfigurationCreateVnicDetails struct {
	AssignPublicIp *bool `mandatory:"false" json:"assignPublicIp"`

	// A user-friendly name for the VNIC. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	PrivateIp *string `mandatory:"false" json:"privateIp"`

	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`

	SubnetId *string `mandatory:"false" json:"subnetId"`
}

func (m InstanceConfigurationCreateVnicDetails) String() string {
	return common.PointerString(m)
}
