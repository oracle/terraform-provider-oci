// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FunctionsTargetDetails The destination function for data transferred from the source.
// For configuration instructions, see
// Creating a Connector (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type FunctionsTargetDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function.
	FunctionId *string `mandatory:"true" json:"functionId"`

	// The batch rollover size in kilobytes.
	// Only one size option can be specified: `batchSizeInKbs` or `batchSizeInNum`.
	BatchSizeInKbs *int `mandatory:"false" json:"batchSizeInKbs"`

	// The batch rollover size in number of messages.
	// Only one size option can be specified: `batchSizeInKbs` or `batchSizeInNum`.
	BatchSizeInNum *int `mandatory:"false" json:"batchSizeInNum"`

	// The batch rollover time in seconds.
	BatchTimeInSec *int `mandatory:"false" json:"batchTimeInSec"`
}

func (m FunctionsTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionsTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FunctionsTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFunctionsTargetDetails FunctionsTargetDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeFunctionsTargetDetails
	}{
		"functions",
		(MarshalTypeFunctionsTargetDetails)(m),
	}

	return json.Marshal(&s)
}
