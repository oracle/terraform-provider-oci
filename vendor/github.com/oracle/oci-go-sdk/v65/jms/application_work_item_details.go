// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplicationWorkItemDetails The work item details with JFR related information.
type ApplicationWorkItemDetails struct {

	// The unique key of the application of the JFR.
	ApplicationKey *string `mandatory:"true" json:"applicationKey"`

	// The application name.
	ApplicationName *string `mandatory:"true" json:"applicationName"`

	// The unique key of the application installation of the JFR.
	ApplicationInstallationKey *string `mandatory:"false" json:"applicationInstallationKey"`

	// The full path on which application installation was detected.
	ApplicationInstallationPath *string `mandatory:"false" json:"applicationInstallationPath"`

	// The work item type.
	WorkItemType WorkItemTypeEnum `mandatory:"false" json:"workItemType,omitempty"`
}

// GetWorkItemType returns WorkItemType
func (m ApplicationWorkItemDetails) GetWorkItemType() WorkItemTypeEnum {
	return m.WorkItemType
}

func (m ApplicationWorkItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationWorkItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWorkItemTypeEnum(string(m.WorkItemType)); !ok && m.WorkItemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkItemType: %s. Supported values are: %s.", m.WorkItemType, strings.Join(GetWorkItemTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ApplicationWorkItemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApplicationWorkItemDetails ApplicationWorkItemDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeApplicationWorkItemDetails
	}{
		"APPLICATION",
		(MarshalTypeApplicationWorkItemDetails)(m),
	}

	return json.Marshal(&s)
}
