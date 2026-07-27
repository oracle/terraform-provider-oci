// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbClusterSummary A summary of a shared-storage DB cluster.
type DbClusterSummary struct {

	// The OCID of the shared-storage DB cluster.
	Id *string `mandatory:"true" json:"id"`

	// Name of the shared-storage DB cluster. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current lifecycle state of the shared-storage DB cluster.
	LifecycleState DbClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Intended MySQL version of the shared-storage DB cluster.
	// You can specify or change the intended MySQL version using the create or update shared-storage DB cluster API calls.
	// All DB nodes run the same version, except temporarily during rolling maintenance.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The date and time the shared-storage DB cluster was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the shared-storage DB cluster was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the configuration used by the shared-storage DB cluster.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The shape of the shared-storage DB cluster. The shape determines the resources, CPU cores and memory, allocated for each DB node in the shared-storage DB cluster.
	// To get a list of shapes, use the ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The OCID of the compartment the shared-storage DB cluster belongs in.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The list of email addresses that receive information from Oracle about the specified OCI shared-storage DB cluster resource.
	// Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates,
	// information about system hardware, and other information needed by administrators.
	// Up to 10 email addresses can be added to the contacts for a shared-storage DB cluster.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	DataStorage *DbClusterDataStorage `mandatory:"false" json:"dataStorage"`

	// A list of OCIDs of all DB nodes attached to the shared-storage DB cluster.
	DbNodeIds []string `mandatory:"false" json:"dbNodeIds"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Description of the shared-storage DB cluster.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Additional information about the current lifecycle state of the shared-storage DB cluster.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Actual MySQL version used for all DB nodes in a shared-storage DB cluster.
	// This version is controlled by the service and could be different from the intended MySQL version (mysqlVersion)
	// for the shared-storage DB cluster as a side effect of service maintenance events.
	CurrentMysqlVersion *string `mandatory:"false" json:"currentMysqlVersion"`

	// The port the MySQL instance listens on.
	Port *int `mandatory:"false" json:"port"`

	// The port on which X Plugin listens for TCP/IP connections for all DB nodes.
	PortX *int `mandatory:"false" json:"portX"`

	ReadEndpoint *DbClusterReadEndpoint `mandatory:"false" json:"readEndpoint"`

	// The OCID of the subnet the shared-storage DB cluster is associated with.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	Endpoint *DbClusterEndpoint `mandatory:"false" json:"endpoint"`
}

func (m DbClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbClusterLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
