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

// Rule Rule for DYNAMIC selection.
type Rule struct {

	// Tenancy Id (Root Compartment Id)for which the rule is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Compartment ID to dynamically search resources.
	// Provide the compartment ID to which the rule is applicable.
	ResourceCompartmentId *string `mandatory:"true" json:"resourceCompartmentId"`

	// Rule Conditions
	Conditions []Condition `mandatory:"true" json:"conditions"`

	// Based on what the rule is created.
	// It can be based on a resourceProperty or a tag.
	// If based on a tag, basis will be 'definedTagEquals'
	// If based on a resource property, basis will be 'inventoryProperties'
	Basis *string `mandatory:"false" json:"basis"`
}

func (m Rule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Rule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
