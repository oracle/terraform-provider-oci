// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Condition Rule condition
type Condition struct {

	// Attribute Group. Provide a Tag namespace if the rule is based on a tag.
	// Provide resource type if the rule is based on a resource property.
	AttrGroup *string `mandatory:"true" json:"attrGroup"`

	// Attribute Key.Provide Tag key if the rule is based on a tag.
	// Provide resource property name if the rule is based on a resource property.
	AttrKey *string `mandatory:"true" json:"attrKey"`

	// Attribute Value.Provide Tag value if the rule is based on a tag.
	// Provide resource property value if the rule is based on a resource property.
	AttrValue *string `mandatory:"true" json:"attrValue"`
}

func (m Condition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Condition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
