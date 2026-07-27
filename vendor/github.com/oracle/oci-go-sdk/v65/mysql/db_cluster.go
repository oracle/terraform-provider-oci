// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbCluster A shared-storage DB cluster is a logical unit that consists of one or more DB nodes, which are compute instances that run MySQL Server and share the same storage, with one or more secondary DB nodes accessible for reads.
type DbCluster struct {

	// The OCID of the shared-storage DB cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the shared-storage DB cluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The shape of the shared-storage DB cluster. The shape determines the resources, CPU cores and memory, for each DB node in the shared-storage DB cluster.
	// To get a list of shapes, use the ListShapes operation.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	DataStorage *DbClusterDataStorage `mandatory:"true" json:"dataStorage"`

	// Name of the shared-storage DB cluster. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the shared-storage DB cluster.
	LifecycleState DbClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Maintenance *MaintenanceDetails `mandatory:"true" json:"maintenance"`

	// Intended MySQL version of the shared-storage DB cluster.
	// You can specify or change the intended MySQL version using the create or update shared-storage DB cluster API calls.
	// All DB nodes run the same version, except temporarily during rolling maintenance.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The OCID of the subnet the shared-storage DB cluster is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the shared-storage DB cluster was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the shared-storage DB cluster was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the configuration used by the shared-storage DB cluster.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The list of email addresses that receive information from Oracle about the specified OCI shared-storage DB cluster resource.
	// Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates,
	// information about system hardware, and other information needed by administrators.
	// Up to 10 email addresses can be added to the contacts for a shared-storage DB cluster.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// A list with a summary of all the DB nodes attached to the shared-storage DB cluster.
	DbNodes []DbNode `mandatory:"false" json:"dbNodes"`

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

	// OCIDs of network security groups used for the VNIC attachment.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The port the MySQL instance listens on.
	Port *int `mandatory:"false" json:"port"`

	// The port on which X Plugin listens for TCP/IP connections for all DB nodes.
	PortX *int `mandatory:"false" json:"portX"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see ZPR Artifacts (https://docs.oracle.com/en-us/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	Source DbClusterSource `mandatory:"false" json:"source"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	Endpoint *DbClusterEndpoint `mandatory:"false" json:"endpoint"`

	ReadEndpoint *DbClusterReadEndpoint `mandatory:"false" json:"readEndpoint"`
}

func (m DbCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbClusterLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DbCluster) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConfigurationId     *string                           `json:"configurationId"`
		CustomerContacts    []CustomerContact                 `json:"customerContacts"`
		DbNodes             []DbNode                          `json:"dbNodes"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		Description         *string                           `json:"description"`
		FreeformTags        map[string]string                 `json:"freeformTags"`
		LifecycleDetails    *string                           `json:"lifecycleDetails"`
		CurrentMysqlVersion *string                           `json:"currentMysqlVersion"`
		NsgIds              []string                          `json:"nsgIds"`
		Port                *int                              `json:"port"`
		PortX               *int                              `json:"portX"`
		SecurityAttributes  map[string]map[string]interface{} `json:"securityAttributes"`
		Source              dbclustersource                   `json:"source"`
		SystemTags          map[string]map[string]interface{} `json:"systemTags"`
		Endpoint            *DbClusterEndpoint                `json:"endpoint"`
		ReadEndpoint        *DbClusterReadEndpoint            `json:"readEndpoint"`
		Id                  *string                           `json:"id"`
		CompartmentId       *string                           `json:"compartmentId"`
		ShapeName           *string                           `json:"shapeName"`
		DataStorage         *DbClusterDataStorage             `json:"dataStorage"`
		DisplayName         *string                           `json:"displayName"`
		LifecycleState      DbClusterLifecycleStateEnum       `json:"lifecycleState"`
		Maintenance         *MaintenanceDetails               `json:"maintenance"`
		MysqlVersion        *string                           `json:"mysqlVersion"`
		SubnetId            *string                           `json:"subnetId"`
		TimeCreated         *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated         *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConfigurationId = model.ConfigurationId

	m.CustomerContacts = make([]CustomerContact, len(model.CustomerContacts))
	copy(m.CustomerContacts, model.CustomerContacts)
	m.DbNodes = make([]DbNode, len(model.DbNodes))
	copy(m.DbNodes, model.DbNodes)
	m.DefinedTags = model.DefinedTags

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.LifecycleDetails = model.LifecycleDetails

	m.CurrentMysqlVersion = model.CurrentMysqlVersion

	m.NsgIds = make([]string, len(model.NsgIds))
	copy(m.NsgIds, model.NsgIds)
	m.Port = model.Port

	m.PortX = model.PortX

	m.SecurityAttributes = model.SecurityAttributes

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(DbClusterSource)
	} else {
		m.Source = nil
	}

	m.SystemTags = model.SystemTags

	m.Endpoint = model.Endpoint

	m.ReadEndpoint = model.ReadEndpoint

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ShapeName = model.ShapeName

	m.DataStorage = model.DataStorage

	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	m.Maintenance = model.Maintenance

	m.MysqlVersion = model.MysqlVersion

	m.SubnetId = model.SubnetId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// DbClusterLifecycleStateEnum Enum with underlying type: string
type DbClusterLifecycleStateEnum string

// Set of constants representing the allowable values for DbClusterLifecycleStateEnum
const (
	DbClusterLifecycleStateCreating DbClusterLifecycleStateEnum = "CREATING"
	DbClusterLifecycleStateActive   DbClusterLifecycleStateEnum = "ACTIVE"
	DbClusterLifecycleStateInactive DbClusterLifecycleStateEnum = "INACTIVE"
	DbClusterLifecycleStateUpdating DbClusterLifecycleStateEnum = "UPDATING"
	DbClusterLifecycleStateDeleting DbClusterLifecycleStateEnum = "DELETING"
	DbClusterLifecycleStateDeleted  DbClusterLifecycleStateEnum = "DELETED"
	DbClusterLifecycleStateFailed   DbClusterLifecycleStateEnum = "FAILED"
)

var mappingDbClusterLifecycleStateEnum = map[string]DbClusterLifecycleStateEnum{
	"CREATING": DbClusterLifecycleStateCreating,
	"ACTIVE":   DbClusterLifecycleStateActive,
	"INACTIVE": DbClusterLifecycleStateInactive,
	"UPDATING": DbClusterLifecycleStateUpdating,
	"DELETING": DbClusterLifecycleStateDeleting,
	"DELETED":  DbClusterLifecycleStateDeleted,
	"FAILED":   DbClusterLifecycleStateFailed,
}

var mappingDbClusterLifecycleStateEnumLowerCase = map[string]DbClusterLifecycleStateEnum{
	"creating": DbClusterLifecycleStateCreating,
	"active":   DbClusterLifecycleStateActive,
	"inactive": DbClusterLifecycleStateInactive,
	"updating": DbClusterLifecycleStateUpdating,
	"deleting": DbClusterLifecycleStateDeleting,
	"deleted":  DbClusterLifecycleStateDeleted,
	"failed":   DbClusterLifecycleStateFailed,
}

// GetDbClusterLifecycleStateEnumValues Enumerates the set of values for DbClusterLifecycleStateEnum
func GetDbClusterLifecycleStateEnumValues() []DbClusterLifecycleStateEnum {
	values := make([]DbClusterLifecycleStateEnum, 0)
	for _, v := range mappingDbClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbClusterLifecycleStateEnumStringValues Enumerates the set of values in String for DbClusterLifecycleStateEnum
func GetDbClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDbClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbClusterLifecycleStateEnum(val string) (DbClusterLifecycleStateEnum, bool) {
	enum, ok := mappingDbClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
