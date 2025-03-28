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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportImportOptionsSummary An Options object represents configuration options to be exported.
type ExportImportOptionsSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID is generated
	// when the item is created.
	Id *string `mandatory:"false" json:"id"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The name by which a configuration entity is displayed to the end user.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The options are stored here as JSON.
	Options *interface{} `mandatory:"false" json:"options"`

	// A string that specifies the group that an OPTIONS item belongs to.
	Group *string `mandatory:"false" json:"group"`

	// An optional string that describes what the options are intended or used for.
	Description *string `mandatory:"false" json:"description"`
}

// GetId returns Id
func (m ExportImportOptionsSummary) GetId() *string {
	return m.Id
}

// GetFreeformTags returns FreeformTags
func (m ExportImportOptionsSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m ExportImportOptionsSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m ExportImportOptionsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportImportOptionsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExportImportOptionsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExportImportOptionsSummary ExportImportOptionsSummary
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeExportImportOptionsSummary
	}{
		"OPTIONS",
		(MarshalTypeExportImportOptionsSummary)(m),
	}

	return json.Marshal(&s)
}
