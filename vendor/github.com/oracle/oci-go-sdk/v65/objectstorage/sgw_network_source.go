// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SgwNetworkSource Matches a specific Service Gateway, or a set of Service Gateways.
type SgwNetworkSource struct {

	// The ID of the SGW to match, or "ALL" to match all SGWs in the specified compartment.
	SgwId *string `mandatory:"true" json:"sgwId"`

	// The SGW must exist in the specified compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m SgwNetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SgwNetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SgwNetworkSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSgwNetworkSource SgwNetworkSource
	s := struct {
		DiscriminatorParam string `json:"networkSourceType"`
		MarshalTypeSgwNetworkSource
	}{
		"SGW",
		(MarshalTypeSgwNetworkSource)(m),
	}

	return json.Marshal(&s)
}
