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

// InternetNetworkSource Matches public internet traffic from a range of IP addresses.
type InternetNetworkSource struct {

	// The network traffic must originate from the specified IP range, expressed in CIDR notation, to match.
	// Currently, only IPv4 addresses are supported.
	SourceIpAddress *string `mandatory:"true" json:"sourceIpAddress"`
}

func (m InternetNetworkSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternetNetworkSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InternetNetworkSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInternetNetworkSource InternetNetworkSource
	s := struct {
		DiscriminatorParam string `json:"networkSourceType"`
		MarshalTypeInternetNetworkSource
	}{
		"INTERNET",
		(MarshalTypeInternetNetworkSource)(m),
	}

	return json.Marshal(&s)
}
