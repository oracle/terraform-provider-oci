// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AnyNetworkSource Matches all network source types, including those added to this API in the future.
type AnyNetworkSource struct {

	// The network traffic must originate from the specified IP range, expressed in CIDR notation, to match.
	// Currently, only IPv4 addresses are supported.
	SourceIpAddress *string `mandatory:"false" json:"sourceIpAddress"`
}

func (m AnyNetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnyNetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AnyNetworkSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAnyNetworkSource AnyNetworkSource
	s := struct {
		DiscriminatorParam string `json:"networkSourceType"`
		MarshalTypeAnyNetworkSource
	}{
		"ANY",
		(MarshalTypeAnyNetworkSource)(m),
	}

	return json.Marshal(&s)
}
