// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service. Use this API to install, configure, and manage resources via the "infrastructure-as-code" model. For more information, see Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TerraformVersionCollection The list of Terraform versions supported for use with stacks.
type TerraformVersionCollection struct {

	// Collection of supported Terraform versions.
	Items []TerraformVersionSummary `mandatory:"false" json:"items"`
}

func (m TerraformVersionCollection) String() string {
	return common.PointerString(m)
}
