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

// CreateDbClusterDetails Details used to create a shared-storage DB cluster.
// In MySQL HeatWave Service, a shared-storage DB cluster is a logical unit composed of one or more
// DB nodes (compute instances running MySQL). It decouples compute from a managed storage layer,
// enabling independent scaling of compute and storage.
type CreateDbClusterDetails struct {

	// The password for the administrator. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character, and
	// 1 special (non-alphanumeric) character.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// Administrator username to use for all DB nodes.
	AdminUsername *string `mandatory:"true" json:"adminUsername"`

	// The OCID of the compartment that contains the shared-storage DB cluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The shape of the shared-storage DB cluster. The shape determines the resources, CPU cores and memory, for each DB node in the shared-storage DB cluster.
	// To get a list of shapes, use the ListShapes operation.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// A list of DB nodes in the shared-storage DB cluster. Exactly one node must be the primary.
	DbNodes []CreateDbNodeDetails `mandatory:"true" json:"dbNodes"`

	// The OCID of the subnet the shared-storage DB cluster is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID of the configuration to be used for the shared-storage DB cluster.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The list of email addresses that receive information from Oracle about the specified OCI shared-storage DB cluster resource.
	// Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates,
	// information about system hardware and other information needed by administrators.
	// Up to 10 email addresses can be added to the contacts for a shared-storage DB cluster.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	DataStorage *CreateDbClusterDataStorageDetails `mandatory:"false" json:"dataStorage"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Description of the shared-storage DB cluster.
	Description *string `mandatory:"false" json:"description"`

	// Name for the shared-storage DB cluster. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	Maintenance *CreateDbClusterMaintenanceDetails `mandatory:"false" json:"maintenance"`

	// MySQL version to use for all DB nodes in the shared-storage DB cluster.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// OCIDs of network security groups used for the VNIC attachment.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The port the MySQL instance listens on.
	Port *int `mandatory:"false" json:"port"`

	// The port on which X Plugin listens for TCP/IP connections for all DB nodes.
	PortX *int `mandatory:"false" json:"portX"`

	ReadEndpoint *CreateDbClusterReadEndpointDetails `mandatory:"false" json:"readEndpoint"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see ZPR Artifacts (https://docs.oracle.com/en-us/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	Source CreateDbClusterSourceDetails `mandatory:"false" json:"source"`

	Endpoint *CreateDbClusterEndpointDetails `mandatory:"false" json:"endpoint"`
}

func (m CreateDbClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDbClusterDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConfigurationId    *string                             `json:"configurationId"`
		CustomerContacts   []CustomerContact                   `json:"customerContacts"`
		DataStorage        *CreateDbClusterDataStorageDetails  `json:"dataStorage"`
		DefinedTags        map[string]map[string]interface{}   `json:"definedTags"`
		Description        *string                             `json:"description"`
		DisplayName        *string                             `json:"displayName"`
		FreeformTags       map[string]string                   `json:"freeformTags"`
		Maintenance        *CreateDbClusterMaintenanceDetails  `json:"maintenance"`
		MysqlVersion       *string                             `json:"mysqlVersion"`
		NsgIds             []string                            `json:"nsgIds"`
		Port               *int                                `json:"port"`
		PortX              *int                                `json:"portX"`
		ReadEndpoint       *CreateDbClusterReadEndpointDetails `json:"readEndpoint"`
		SecurityAttributes map[string]map[string]interface{}   `json:"securityAttributes"`
		Source             createdbclustersourcedetails        `json:"source"`
		Endpoint           *CreateDbClusterEndpointDetails     `json:"endpoint"`
		AdminPassword      *string                             `json:"adminPassword"`
		AdminUsername      *string                             `json:"adminUsername"`
		CompartmentId      *string                             `json:"compartmentId"`
		ShapeName          *string                             `json:"shapeName"`
		DbNodes            []CreateDbNodeDetails               `json:"dbNodes"`
		SubnetId           *string                             `json:"subnetId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConfigurationId = model.ConfigurationId

	m.CustomerContacts = make([]CustomerContact, len(model.CustomerContacts))
	copy(m.CustomerContacts, model.CustomerContacts)
	m.DataStorage = model.DataStorage

	m.DefinedTags = model.DefinedTags

	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.Maintenance = model.Maintenance

	m.MysqlVersion = model.MysqlVersion

	m.NsgIds = make([]string, len(model.NsgIds))
	copy(m.NsgIds, model.NsgIds)
	m.Port = model.Port

	m.PortX = model.PortX

	m.ReadEndpoint = model.ReadEndpoint

	m.SecurityAttributes = model.SecurityAttributes

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(CreateDbClusterSourceDetails)
	} else {
		m.Source = nil
	}

	m.Endpoint = model.Endpoint

	m.AdminPassword = model.AdminPassword

	m.AdminUsername = model.AdminUsername

	m.CompartmentId = model.CompartmentId

	m.ShapeName = model.ShapeName

	m.DbNodes = make([]CreateDbNodeDetails, len(model.DbNodes))
	copy(m.DbNodes, model.DbNodes)
	m.SubnetId = model.SubnetId

	return
}
