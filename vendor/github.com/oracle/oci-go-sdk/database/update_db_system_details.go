// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateDbSystemDetails Describes the parameters for updating the DB system.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateDbSystemDetails struct {

	// The new number of CPU cores to set for the DB system. Not applicable for virtual machine DB systems.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The size, in gigabytes, to scale the attached storage up to for this virtual machine DB system. This value must be greater than current storage size. Note that the resulting total storage size attached will be greater than the amount requested to allow for REDO/RECO space and software volume. Applies only to virtual machine DB systems.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The public key portion of the key pair to use for SSH access to the DB system. Multiple public keys can be provided. The length of the combined keys cannot exceed 10,000 characters.
	SshPublicKeys []string `mandatory:"false" json:"sshPublicKeys"`

	Version *PatchDetails `mandatory:"false" json:"version"`
}

func (m UpdateDbSystemDetails) String() string {
	return common.PointerString(m)
}
