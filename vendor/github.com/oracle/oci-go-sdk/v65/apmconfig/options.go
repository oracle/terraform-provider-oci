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

// Options An object that represents configuration options.
type Options struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID is generated
	// when the item is created.
	Id *string `mandatory:"false" json:"id"`

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `mandatory:"false" json:"etag"`

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
func (m Options) GetId() *string {
	return m.Id
}

// GetTimeCreated returns TimeCreated
func (m Options) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m Options) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetCreatedBy returns CreatedBy
func (m Options) GetCreatedBy() *string {
	return m.CreatedBy
}

// GetUpdatedBy returns UpdatedBy
func (m Options) GetUpdatedBy() *string {
	return m.UpdatedBy
}

// GetEtag returns Etag
func (m Options) GetEtag() *string {
	return m.Etag
}

// GetFreeformTags returns FreeformTags
func (m Options) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m Options) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m Options) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Options) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Options) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOptions Options
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeOptions
	}{
		"OPTIONS",
		(MarshalTypeOptions)(m),
	}

	return json.Marshal(&s)
}
