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

// DbClusterSnapshot A snapshot of selected shared-storage DB cluster configuration details captured at the time the backup was taken.
type DbClusterSnapshot struct {

	// The OCID of the shared-storage DB cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the shared-storage DB cluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DataStorage *DbClusterDataStorage `mandatory:"true" json:"dataStorage"`

	// Name of the shared-storage DB cluster. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// MySQL version of the shared-storage DB cluster.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The OCID of the subnet the shared-storage DB cluster is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The administrator username.
	AdminUsername *string `mandatory:"false" json:"adminUsername"`

	// The OCID of the configuration of the shared-storage DB cluster.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The shape of the shared-storage DB cluster at the time the backup was taken.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// A list of DB nodes that were part of the shared-storage DB cluster at the time the backup was taken.
	DbNodes []DbNodeSnapshot `mandatory:"false" json:"dbNodes"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Description of the shared-storage DB cluster.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	Maintenance *MaintenanceDetails `mandatory:"false" json:"maintenance"`

	// OCIDs of network security groups used for the VNIC attachment.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The port the MySQL instance listens on.
	Port *int `mandatory:"false" json:"port"`

	// The port on which X Plugin listens for TCP/IP connections for all DB nodes.
	PortX *int `mandatory:"false" json:"portX"`

	Endpoint *DbClusterEndpoint `mandatory:"false" json:"endpoint"`

	ReadEndpoint *DbClusterReadEndpoint `mandatory:"false" json:"readEndpoint"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see ZPR Artifacts (https://docs.oracle.com/en-us/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DbClusterSnapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbClusterSnapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
