// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateRunbookDetails The information about new Runbook.
type CreateRunbookDetails struct {

	// Type of runbook structure.
	RunbookRelevance RunbookRunbookRelevanceEnum `mandatory:"true" json:"runbookRelevance"`

	// The lifecycle operation performed by the task.
	Operation *string `mandatory:"true" json:"operation"`

	// The OS type for the runbook.
	OsType OsTypeEnum `mandatory:"true" json:"osType"`

	Associations *Associations `mandatory:"true" json:"associations"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The platform of the runbook.
	Platform *string `mandatory:"false" json:"platform"`

	// Is the runbook default?
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Estimated time to successfully complete the runbook execution
	EstimatedTime *string `mandatory:"false" json:"estimatedTime"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRunbookDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRunbookDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRunbookRunbookRelevanceEnum(string(m.RunbookRelevance)); !ok && m.RunbookRelevance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RunbookRelevance: %s. Supported values are: %s.", m.RunbookRelevance, strings.Join(GetRunbookRunbookRelevanceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsTypeEnum(string(m.OsType)); !ok && m.OsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsType: %s. Supported values are: %s.", m.OsType, strings.Join(GetOsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
