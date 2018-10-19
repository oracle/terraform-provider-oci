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

// CreateInstancePoolDetails The data to create an instance pool.
type CreateInstancePoolDetails struct {

	// The OCID of the compartment containing the instance pool
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the instance configuration associated to the instance pool.
	InstanceConfigurationId *string `mandatory:"true" json:"instanceConfigurationId"`

	// The placement configurations for the instance pool.
	// There should be 1 placement configuration for each desired AD.
	PlacementConfigurations []CreateInstancePoolPlacementConfigurationDetails `mandatory:"true" json:"placementConfigurations"`

	// The number of instances that should be in the instance pool.
	Size *int `mandatory:"true" json:"size"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The user-friendly name.  Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateInstancePoolDetails) String() string {
	return common.PointerString(m)
}
