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

// DiscoveredCloudDbInstance The details of an Oracle DB instance discovered in a cloud DB system discovery run.
type DiscoveredCloudDbInstance struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The name of the DB instance.
	InstanceName *string `mandatory:"true" json:"instanceName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Dbaas OCI resource matching the discovered DB system component.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedCloudComponent `mandatory:"false" json:"associatedComponents"`

	// The name of the host on which the DB instance is running.
	HostName *string `mandatory:"false" json:"hostName"`

	// The name of the DB instance node.
	NodeName *string `mandatory:"false" json:"nodeName"`

	// The Oracle home location of the DB instance.
	OracleHome *string `mandatory:"false" json:"oracleHome"`

	// The Automatic Diagnostic Repository (ADR) home directory for the DB instance.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	// The state of the discovered DB system component.
	Status DiscoveredCloudDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredCloudDbInstance) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredCloudDbInstance) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredCloudDbInstance) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredCloudDbInstance) GetResourceId() *string {
	return m.ResourceId
}

// GetDbaasId returns DbaasId
func (m DiscoveredCloudDbInstance) GetDbaasId() *string {
	return m.DbaasId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredCloudDbInstance) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredCloudDbInstance) GetStatus() DiscoveredCloudDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredCloudDbInstance) GetAssociatedComponents() []AssociatedCloudComponent {
	return m.AssociatedComponents
}

func (m DiscoveredCloudDbInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredCloudDbInstance) ValidateEnumValue() (bool, error) {
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
func (m DiscoveredCloudDbInstance) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredCloudDbInstance DiscoveredCloudDbInstance
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredCloudDbInstance
	}{
		"DATABASE_INSTANCE",
		(MarshalTypeDiscoveredCloudDbInstance)(m),
	}

	return json.Marshal(&s)
}
