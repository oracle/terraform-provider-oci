// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatastoreCluster An Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm) Datastore Cluster for software-defined data center.
// The Datastore Cluster combines multiple datastores into a single datastore cluster.
type DatastoreCluster struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Datastore cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Datastores that belong to the Datastore Cluster
	DatastoreIds []string `mandatory:"true" json:"datastoreIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the Datastore.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A descriptive name for the Datastore Cluster. It must be unique within a SDDC, start with a letter, and contain only letters, digits,
	// whitespaces, dashes and underscores.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the datastore cluster.
	DatastoreClusterType DatastoreClusterTypesEnum `mandatory:"true" json:"datastoreClusterType"`

	// The date and time the Datastore Cluster was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Datastore Cluster was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Datastore Cluster.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The availability domain of the Datastore Cluster.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VMware Cluster that Datastore cluster is attached to.
	ClusterId *string `mandatory:"false" json:"clusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC that Datastore cluster is associated with.
	SddcId *string `mandatory:"false" json:"sddcId"`

	// The OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ESXi hosts to attach the
	// datastore to. All ESXi hosts must belong to the same VMware cluster.
	EsxiHostIds []string `mandatory:"false" json:"esxiHostIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DatastoreCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatastoreCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatastoreClusterTypesEnum(string(m.DatastoreClusterType)); !ok && m.DatastoreClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatastoreClusterType: %s. Supported values are: %s.", m.DatastoreClusterType, strings.Join(GetDatastoreClusterTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
