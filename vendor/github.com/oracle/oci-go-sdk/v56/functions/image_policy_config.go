// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ImagePolicyConfig Define the image signature verification policy for an application.
type ImagePolicyConfig struct {

	// Define if image signature verification policy is enabled for the application.
	IsPolicyEnabled *bool `mandatory:"true" json:"isPolicyEnabled"`

	// A list of KMS key details.
	KeyDetails []KeyDetails `mandatory:"false" json:"keyDetails"`
}

func (m ImagePolicyConfig) String() string {
	return common.PointerString(m)
}
