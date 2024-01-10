// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// MonitoredResourceAliasCredential Monitored Resource Alias Credential Details
type MonitoredResourceAliasCredential struct {

	// The source type and source name combination,delimited with (.) separator.
	// Example: {source type}.{source name} and source type max char limit is 63.
	Source *string `mandatory:"true" json:"source"`

	// The name of the alias, within the context of the source.
	Name *string `mandatory:"true" json:"name"`

	Credential *MonitoredResourceAliasSourceCredential `mandatory:"true" json:"credential"`
}

func (m MonitoredResourceAliasCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredResourceAliasCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
