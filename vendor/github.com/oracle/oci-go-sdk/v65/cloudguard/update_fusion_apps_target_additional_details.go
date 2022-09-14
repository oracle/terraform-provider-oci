// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateFusionAppsTargetAdditionalDetails Additional details about fusion apps target.
type UpdateFusionAppsTargetAdditionalDetails struct {

	// Resources that will be monitored
	MonitoringResources []MonitoringResource `mandatory:"false" json:"monitoringResources"`

	// Service account username
	Username *string `mandatory:"false" json:"username"`

	// Service account password
	Password *string `mandatory:"false" json:"password"`

	// URL for the login service.
	LoginServiceUrl *string `mandatory:"false" json:"loginServiceUrl"`

	// Login service type, for example, OAM.
	LoginServiceType LoginServiceTypeEnum `mandatory:"false" json:"loginServiceType,omitempty"`
}

func (m UpdateFusionAppsTargetAdditionalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateFusionAppsTargetAdditionalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLoginServiceTypeEnum(string(m.LoginServiceType)); !ok && m.LoginServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoginServiceType: %s. Supported values are: %s.", m.LoginServiceType, strings.Join(GetLoginServiceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateFusionAppsTargetAdditionalDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateFusionAppsTargetAdditionalDetails UpdateFusionAppsTargetAdditionalDetails
	s := struct {
		DiscriminatorParam string `json:"targetResourceType"`
		MarshalTypeUpdateFusionAppsTargetAdditionalDetails
	}{
		"FACLOUD",
		(MarshalTypeUpdateFusionAppsTargetAdditionalDetails)(m),
	}

	return json.Marshal(&s)
}