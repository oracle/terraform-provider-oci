// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Control Plane API
//
// Use the Application Performance Monitoring Control Plane API to perform operations such as creating, updating,
// deleting and listing APM domains and monitoring the progress of these operations using the work request APIs.
//

package apmcontrolplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// GenerateDataKeyDetails Details of the Data Key to be generated.
type GenerateDataKeyDetails struct {

	// Name of the Data Key. The name uniquely identifies a Data Key within an APM domain.
	Name *string `mandatory:"true" json:"name"`

	// Type of the Data Key.
	Type DataKeyTypesEnum `mandatory:"true" json:"type"`
}

func (m GenerateDataKeyDetails) String() string {
	return common.PointerString(m)
}
