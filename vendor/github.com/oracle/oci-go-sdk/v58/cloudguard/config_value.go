// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ConfigValue configuration item for multi list data type
type ConfigValue struct {

	// configuration list item type, either CUSTOM or MANAGED
	ListType ConfigurationListItemTypeEnum `mandatory:"true" json:"listType"`

	// type of the managed list
	ManagedListType *string `mandatory:"true" json:"managedListType"`

	// configuration value
	Value *string `mandatory:"true" json:"value"`
}

func (m ConfigValue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigValue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConfigurationListItemTypeEnum(string(m.ListType)); !ok && m.ListType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListType: %s. Supported values are: %s.", m.ListType, strings.Join(GetConfigurationListItemTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
