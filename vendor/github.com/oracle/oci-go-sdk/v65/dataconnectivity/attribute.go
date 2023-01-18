// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Attribute The attribute object fo the registry that can be used to get the connector details.
type Attribute struct {

	// The name of of the attribute.
	Name *string `mandatory:"true" json:"name"`

	// True if attribute is sensitive.
	IsSensitive *bool `mandatory:"false" json:"isSensitive"`

	// True if attribute is mandatory.
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`

	// True if attribute is generated.
	IsGenerated *bool `mandatory:"false" json:"isGenerated"`

	// True if attribute is encoded.
	IsBase64Encoded *bool `mandatory:"false" json:"isBase64Encoded"`

	// The list of valid keys.
	ValidKeyList []string `mandatory:"false" json:"validKeyList"`

	// The attribute type details.
	AttributeType *string `mandatory:"false" json:"attributeType"`
}

func (m Attribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Attribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
