// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecureAccessControlRule The access control rule for SECURE_ACCESS network type selection.
type SecureAccessControlRule struct {

	// The type of IP notation.
	IpNotation IpNotationTypeEnum `mandatory:"true" json:"ipNotation"`

	// The associated value of the selected IP notation.
	Value *string `mandatory:"true" json:"value"`

	// A comma-separated IP or CIDR address for VCN OCID IP notation selection.
	VcnIps *string `mandatory:"false" json:"vcnIps"`
}

func (m SecureAccessControlRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecureAccessControlRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIpNotationTypeEnum(string(m.IpNotation)); !ok && m.IpNotation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpNotation: %s. Supported values are: %s.", m.IpNotation, strings.Join(GetIpNotationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
