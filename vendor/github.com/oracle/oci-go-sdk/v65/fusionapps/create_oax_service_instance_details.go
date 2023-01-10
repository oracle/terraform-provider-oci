// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOaxServiceInstanceDetails The information about new Analytics Warehouse instance being provisioned.
type CreateOaxServiceInstanceDetails struct {

	// The service instance type being provisioned
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A unique Name for Analytics Warehouse.
	Name *string `mandatory:"true" json:"name"`

	// Comparment where the instance is to be created
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// This is the description for Analytics Warehouse Service.
	Description *string `mandatory:"false" json:"description"`

	FawAdminInfo *FawAdminInfoDetails `mandatory:"false" json:"FawAdminInfo"`
}

//GetDisplayName returns DisplayName
func (m CreateOaxServiceInstanceDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateOaxServiceInstanceDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m CreateOaxServiceInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOaxServiceInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOaxServiceInstanceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOaxServiceInstanceDetails CreateOaxServiceInstanceDetails
	s := struct {
		DiscriminatorParam string `json:"serviceInstanceType"`
		MarshalTypeCreateOaxServiceInstanceDetails
	}{
		"ANALYTICS_WAREHOUSE",
		(MarshalTypeCreateOaxServiceInstanceDetails)(m),
	}

	return json.Marshal(&s)
}
