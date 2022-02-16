// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// StrategyParameter The metadata associated with the strategy parameter.
type StrategyParameter struct {

	// The name of the strategy parameter.
	Name *string `mandatory:"true" json:"name"`

	// The type of strategy parameter.
	Type StrategyParameterTypeEnum `mandatory:"true" json:"type"`

	// Text describing the strategy parameter.
	Description *string `mandatory:"true" json:"description"`

	// Whether this parameter is required.
	IsRequired *bool `mandatory:"true" json:"isRequired"`

	// A default value used for the strategy parameter.
	DefaultValue *interface{} `mandatory:"false" json:"defaultValue"`

	// The list of possible values used for these strategy parameters.
	PossibleValues []interface{} `mandatory:"false" json:"possibleValues"`
}

func (m StrategyParameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StrategyParameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStrategyParameterTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetStrategyParameterTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
