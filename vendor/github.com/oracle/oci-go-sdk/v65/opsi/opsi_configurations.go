// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OpsiConfigurations An OPSI configuration resource is a container for storing custom values for customizable configuration items exposed by Operations Insights.
// Operations Insights exposes different sets of customizable configuration items through different OPSI configuration types.
// UX_CONFIGURATION: OPSI configuration resource of this type can be created only once in each compartment. It is a compartment level singleton resource.
// When configuration values, for an OPSI configuration type that supports compartment level singleton (e.g: UX_CONFIGURATION) resource, are queried for a compartment,
// following will be the order of preference.
// 1. If the specified compartment has an OPSI configuration resource, first preference will be given to the custom values inside that.
// 2. If the root compartment has an OPSI configuration resource, it will be considered as applicable to all compartments of that tenency,
// hence second preference will be given to the custom values inside that.
// 3. Default configuration will be considered as a final fallback option.
type OpsiConfigurations struct {

	// OPSI Configuration Object.
	OpsiConfigurations *interface{} `mandatory:"false" json:"opsiConfigurations"`
}

func (m OpsiConfigurations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpsiConfigurations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
