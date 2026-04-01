// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmIscsiDetails The host iSCSI details.
type OlvmIscsiDetails struct {

	// Address of iSCSI
	Address *string `mandatory:"false" json:"address"`

	// Disk ID of iSCSI
	DiskId *string `mandatory:"false" json:"diskId"`

	// Initiator of iSCSI
	Initiator *string `mandatory:"false" json:"initiator"`

	// LUN Mapping of iSCSI
	LunMapping *int `mandatory:"false" json:"lunMapping"`

	// Number of paths of iSCSI
	Paths *int `mandatory:"false" json:"paths"`

	// Port number of iSCSI
	Port *int `mandatory:"false" json:"port"`

	// Portal of iSCSI
	Portal *string `mandatory:"false" json:"portal"`

	// Product ID of iSCSI
	ProductId *string `mandatory:"false" json:"productId"`

	// Serial of iSCSI
	Serial *string `mandatory:"false" json:"serial"`

	// size of iSCSI
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// Status of iSCSI
	Status *string `mandatory:"false" json:"status"`

	// Storage Domain ID of iSCSI
	StorageDomainId *string `mandatory:"false" json:"storageDomainId"`

	// target of iSCSI
	Target *string `mandatory:"false" json:"target"`

	// Username of iSCSI
	Username *string `mandatory:"false" json:"username"`

	// Vendor ID of iSCSI
	VendorId *string `mandatory:"false" json:"vendorId"`

	// Volume Group ID of iSCSI
	VolumeGroupId *string `mandatory:"false" json:"volumeGroupId"`
}

func (m OlvmIscsiDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmIscsiDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
