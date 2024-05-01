// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredExternalCluster The details of an external cluster discovered in an external DB system discovery run.
type DiscoveredExternalCluster struct {

	// The identifier of the discovered DB system component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the discovered DB system component.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	IsSelectedForMonitoring *bool `mandatory:"false" json:"isSelectedForMonitoring"`

	// The list of associated components.
	AssociatedComponents []AssociatedComponent `mandatory:"false" json:"associatedComponents"`

	// The directory in which Oracle Grid Infrastructure is installed.
	GridHome *string `mandatory:"false" json:"gridHome"`

	// The version of Oracle Clusterware running in the cluster.
	Version *string `mandatory:"false" json:"version"`

	// Indicates whether the cluster is an Oracle Flex Cluster or not.
	IsFlexCluster *bool `mandatory:"false" json:"isFlexCluster"`

	// The list of network address configurations of the external cluster.
	NetworkConfigurations []ExternalClusterNetworkConfiguration `mandatory:"false" json:"networkConfigurations"`

	// The list of Virtual IP (VIP) configurations of the external cluster.
	VipConfigurations []ExternalClusterVipConfiguration `mandatory:"false" json:"vipConfigurations"`

	// The list of Single Client Access Name (SCAN) configurations of the external cluster.
	ScanConfigurations []ExternalClusterScanListenerConfiguration `mandatory:"false" json:"scanConfigurations"`

	// The location of the Oracle Cluster Registry (OCR) file.
	OcrFileLocation *string `mandatory:"false" json:"ocrFileLocation"`

	ClusterInstances []DiscoveredExternalClusterInstance `mandatory:"false" json:"clusterInstances"`

	// The state of the discovered DB system component.
	Status DiscoveredExternalDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetComponentId returns ComponentId
func (m DiscoveredExternalCluster) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m DiscoveredExternalCluster) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m DiscoveredExternalCluster) GetComponentName() *string {
	return m.ComponentName
}

// GetResourceId returns ResourceId
func (m DiscoveredExternalCluster) GetResourceId() *string {
	return m.ResourceId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m DiscoveredExternalCluster) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m DiscoveredExternalCluster) GetStatus() DiscoveredExternalDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m DiscoveredExternalCluster) GetAssociatedComponents() []AssociatedComponent {
	return m.AssociatedComponents
}

func (m DiscoveredExternalCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveredExternalCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveredExternalDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredExternalDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DiscoveredExternalCluster) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDiscoveredExternalCluster DiscoveredExternalCluster
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeDiscoveredExternalCluster
	}{
		"CLUSTER",
		(MarshalTypeDiscoveredExternalCluster)(m),
	}

	return json.Marshal(&s)
}
