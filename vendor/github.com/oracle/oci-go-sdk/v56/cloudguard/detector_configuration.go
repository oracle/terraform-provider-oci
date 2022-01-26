// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DetectorConfiguration A single configuration applied to a detector
type DetectorConfiguration struct {

	// Unique name of the configuration
	ConfigKey *string `mandatory:"true" json:"configKey"`

	// configuration name
	Name *string `mandatory:"true" json:"name"`

	// configuration value
	Value *string `mandatory:"false" json:"value"`

	// configuration data type
	DataType *string `mandatory:"false" json:"dataType"`

	// List of configuration values
	Values []ConfigValue `mandatory:"false" json:"values"`
}

func (m DetectorConfiguration) String() string {
	return common.PointerString(m)
}
