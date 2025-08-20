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

// CollectdResourceNameConfigurationDetails Resource name generation overriding configurations for collectd resource types.
type CollectdResourceNameConfigurationDetails struct {

	// String to be suffixed to the resource name.
	Suffix *string `mandatory:"false" json:"suffix"`

	// List of property names to be included.
	IncludeProperties []string `mandatory:"false" json:"includeProperties"`

	// List of property names to be excluded.
	ExcludeProperties []string `mandatory:"false" json:"excludeProperties"`
}

func (m CollectdResourceNameConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CollectdResourceNameConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
