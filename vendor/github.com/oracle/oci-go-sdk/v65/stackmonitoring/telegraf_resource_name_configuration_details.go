// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TelegrafResourceNameConfigurationDetails Resource name generation overriding configurations for telegraf resource types.
type TelegrafResourceNameConfigurationDetails struct {

	// Flag to indicate if only tags will be used for resource name generation.
	IsUseTagsOnly *bool `mandatory:"false" json:"isUseTagsOnly"`

	// List of tag names to be included.
	IncludeTags []string `mandatory:"false" json:"includeTags"`

	// List of tag names to be excluded.
	ExcludeTags []string `mandatory:"false" json:"excludeTags"`
}

func (m TelegrafResourceNameConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TelegrafResourceNameConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
