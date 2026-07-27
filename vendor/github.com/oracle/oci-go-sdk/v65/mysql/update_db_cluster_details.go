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

// UpdateDbClusterDetails Details used to update a shared-storage DB cluster.
type UpdateDbClusterDetails struct {

	// The password for the administrative user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character, and
	// 1 special (non-alphanumeric) character.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// The OCID of the configuration to be used for the shared-storage DB cluster.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The shape of the shared-storage DB cluster. The shape determines the resources, CPU cores and memory, allocated for each DB node in the shared-storage DB cluster.
	// To get a list of shapes, use the ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The list of email addresses that receive information from Oracle about the specified OCI shared-storage DB cluster resource.
	// Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates,
	// information about system hardware, and other information needed by administrators.
	// Up to 10 email addresses can be added to the contacts for a shared-storage DB cluster.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	DataStorage *UpdateDbClusterDataStorageDetails `mandatory:"false" json:"dataStorage"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Description of the shared-storage DB cluster.
	Description *string `mandatory:"false" json:"description"`

	// Name of the shared-storage DB cluster. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	Maintenance *UpdateDbClusterMaintenanceDetails `mandatory:"false" json:"maintenance"`

	// MySQL version to use for all DB nodes in the shared-storage DB cluster.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see ZPR Artifacts (https://docs.oracle.com/en-us/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// OCIDs of network security groups used for the VNIC attachment.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	Endpoint *UpdateDbClusterEndpointDetails `mandatory:"false" json:"endpoint"`

	ReadEndpoint *UpdateDbClusterReadEndpointDetails `mandatory:"false" json:"readEndpoint"`
}

func (m UpdateDbClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDbClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
