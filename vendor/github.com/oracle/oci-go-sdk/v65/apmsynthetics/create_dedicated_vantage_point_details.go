// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDedicatedVantagePointDetails Details of the request body used to create a new dedicated vantage point.
type CreateDedicatedVantagePointDetails struct {

	// Unique dedicated vantage point name that cannot be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	DvpStackDetails DvpStackDetails `mandatory:"true" json:"dvpStackDetails"`

	// Name of the region.
	Region *string `mandatory:"true" json:"region"`

	// Status of the dedicated vantage point.
	Status DedicatedVantagePointStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDedicatedVantagePointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDedicatedVantagePointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDedicatedVantagePointStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDedicatedVantagePointStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDedicatedVantagePointDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Status          DedicatedVantagePointStatusEnum   `json:"status"`
		FreeformTags    map[string]string                 `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{} `json:"definedTags"`
		DisplayName     *string                           `json:"displayName"`
		DvpStackDetails dvpstackdetails                   `json:"dvpStackDetails"`
		Region          *string                           `json:"region"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Status = model.Status

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	nn, e = model.DvpStackDetails.UnmarshalPolymorphicJSON(model.DvpStackDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DvpStackDetails = nn.(DvpStackDetails)
	} else {
		m.DvpStackDetails = nil
	}

	m.Region = model.Region

	return
}
