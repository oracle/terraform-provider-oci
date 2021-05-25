// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// CreateConfigurationDetails The details required to create a new Configuration.
type CreateConfigurationDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the associated Shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// User-provided data about the Configuration.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the Configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	Variables *ConfigurationVariables `mandatory:"false" json:"variables"`

	// The OCID of the Configuration from which the new Configuration is derived. The values in CreateConfigurationDetails.variables supersede the variables of the parent Configuration.
	ParentConfigurationId *string `mandatory:"false" json:"parentConfigurationId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateConfigurationDetails) String() string {
	return common.PointerString(m)
}
