// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateConfigurationDetails Parameters to update Cloud Guard configuration details for a tenancy.
type UpdateConfigurationDetails struct {

	// The reporting region
	ReportingRegion *string `mandatory:"true" json:"reportingRegion"`

	// Status of Cloud Guard tenant
	Status CloudGuardStatusEnum `mandatory:"true" json:"status"`

	// List of service configurations for tenant
	ServiceConfigurations []ServiceConfiguration `mandatory:"false" json:"serviceConfigurations"`

	// Identifies if Oracle managed resources will be created by customers.
	// If no value is specified false is the default.
	SelfManageResources *bool `mandatory:"false" json:"selfManageResources"`
}

func (m UpdateConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudGuardStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCloudGuardStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ServiceConfigurations []serviceconfiguration `json:"serviceConfigurations"`
		SelfManageResources   *bool                  `json:"selfManageResources"`
		ReportingRegion       *string                `json:"reportingRegion"`
		Status                CloudGuardStatusEnum   `json:"status"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ServiceConfigurations = make([]ServiceConfiguration, len(model.ServiceConfigurations))
	for i, n := range model.ServiceConfigurations {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ServiceConfigurations[i] = nn.(ServiceConfiguration)
		} else {
			m.ServiceConfigurations[i] = nil
		}
	}
	m.SelfManageResources = model.SelfManageResources

	m.ReportingRegion = model.ReportingRegion

	m.Status = model.Status

	return
}
