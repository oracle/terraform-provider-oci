// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssociationRuleOverrideParameters The parameters to be used while creating associations by applying a rule override.
type AssociationRuleOverrideParameters struct {

	// The OCID of the log group which would contain the collected logs.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`

	// The association properties. Properties of multiple source types can be specified.
	// When creating associations, relevant ones will be included for corresponding sources.
	Properties []AssociationProperty `mandatory:"false" json:"properties"`

	// The credentials to be used while creating associations by applying this rule override.
	Credentials []CollectionRuleCredential `mandatory:"false" json:"credentials"`
}

func (m AssociationRuleOverrideParameters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociationRuleOverrideParameters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
