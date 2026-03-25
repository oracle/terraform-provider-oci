// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AiDataPlatform Control Plane API
//
// Use the AiDataPlatform Control Plane API to manage Data Lakes.
//

package aidataplatform

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisableAiFeatureDetails The data to disable Ai Feature for AiDataPlatform.
type DisableAiFeatureDetails struct {

	// When true, allows deletion of the AI Lakehouse. Default Value would be false.
	IsForceDeleteAiLakehouse *bool `mandatory:"false" json:"isForceDeleteAiLakehouse"`
}

func (m DisableAiFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisableAiFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
