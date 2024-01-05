// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceIpAddressCondition An access control rule condition that requires a match on the specified source IP address or address range.
type SourceIpAddressCondition struct {

	// An IPv4 or IPv6 address range that the source IP address of an incoming packet must match.
	// The service accepts only classless inter-domain routing (CIDR) format (x.x.x.x/y or x:x::x/y) strings.
	// Specify 0.0.0.0/0 or ::/0 to match all incoming traffic.
	// example: "192.168.0.0/16"
	AttributeValue *string `mandatory:"true" json:"attributeValue"`
}

func (m SourceIpAddressCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SourceIpAddressCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SourceIpAddressCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceIpAddressCondition SourceIpAddressCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypeSourceIpAddressCondition
	}{
		"SOURCE_IP_ADDRESS",
		(MarshalTypeSourceIpAddressCondition)(m),
	}

	return json.Marshal(&s)
}
