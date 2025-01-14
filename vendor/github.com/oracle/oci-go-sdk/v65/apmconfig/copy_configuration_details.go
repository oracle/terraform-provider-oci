// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CopyConfigurationDetails Array of configuration items with dependencies to copy to a destination domain.
type CopyConfigurationDetails struct {

	// Simple key-value pair that has parameters related to the import process (EnableOcidSubstitution, DestinationDomainID, …) and more.
	// Example: `{"parameter-key": "parameter-value"}`
	// Supported parameters:
	// — Enable the OCIDs in instructions to be replaced, if set to "true" The Config Service replace any OCIDs it finds
	// in the instructions.
	// — Destination APM Domain ID where the configuration Item(s) will be fast imported to.
	// — List of Configuration Type or Groups to be fast imported.
	// — the compartment Id we will fast import to,
	// if the compartment Id is not provided it will be the default destination domain compartmentId.
	ConfigurationMap map[string]string `mandatory:"true" json:"configurationMap"`
}

func (m CopyConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CopyConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
