// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// NamespaceDetails Namespace details.
type NamespaceDetails struct {

	// Usage type.
	Usage NamespaceUsageEnum `mandatory:"true" json:"usage"`

	// Metric namespace.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Resource group.
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`
}

func (m NamespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceUsageEnum(string(m.Usage)); !ok && m.Usage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Usage: %s. Supported values are: %s.", m.Usage, strings.Join(GetNamespaceUsageEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
