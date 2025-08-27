// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredCloudDbHome The details of an Oracle DB home discovered in a cloud DB system discovery run.
type DiscoveredCloudDbHome struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The location of the DB home.
	HomeDirectory *string `mandatory:"true" json:"homeDirectory"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Dbaas OCI resource matching the discovered DB system component.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedCloudComponent `mandatory:"false" json:"associatedComponents"`

	// The state of the discovered DB system component.
	Status DiscoveredCloudDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredCloudDbHome) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredCloudDbHome) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredCloudDbHome) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredCloudDbHome) GetResourceId() *string {
	return m.ResourceId
}

// GetDbaasId returns DbaasId
func (m DiscoveredCloudDbHome) GetDbaasId() *string {
	return m.DbaasId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredCloudDbHome) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredCloudDbHome) GetStatus() DiscoveredCloudDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredCloudDbHome) GetAssociatedComponents() []AssociatedCloudComponent {
	return m.AssociatedComponents
}

func (m DiscoveredCloudDbHome) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredCloudDbHome) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveredCloudDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredCloudDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DiscoveredCloudDbHome) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredCloudDbHome DiscoveredCloudDbHome
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredCloudDbHome
	}{
		"DATABASE_HOME",
		(MarshalTypeDiscoveredCloudDbHome)(m),
	}

	return json.Marshal(&s)
}
