// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogDataMappingDetails Log data mapping details.
type LogDataMappingDetails struct {

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The mapping rule between keys.
	DataMappingRule *string `mandatory:"true" json:"dataMappingRule"`

	// List of OCIDs of log objects and log groups.
	LogSources []string `mandatory:"true" json:"logSources"`

	// Validity state of log data mapping.
	ValidityState LogDataMappingValidityStateEnum `mandatory:"false" json:"validityState,omitempty"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m LogDataMappingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogDataMappingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLogDataMappingValidityStateEnum(string(m.ValidityState)); !ok && m.ValidityState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidityState: %s. Supported values are: %s.", m.ValidityState, strings.Join(GetLogDataMappingValidityStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
