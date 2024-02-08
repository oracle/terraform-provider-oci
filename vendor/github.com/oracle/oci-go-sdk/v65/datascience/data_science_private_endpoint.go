// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSciencePrivateEndpoint Data Science private endpoint.
type DataSciencePrivateEndpoint struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create private endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A user friendly name. It doesn't have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The OCID of a private endpoint.
	Id *string `mandatory:"true" json:"id"`

	// State of the Data Science private endpoint.
	LifecycleState DataSciencePrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user that created the private endpoint.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID of a subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time that the Data Science private endpoint was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time that the Data Science private endpoint was updated expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A user friendly description. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Details of the state of Data Science private endpoint.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// An array of network security group OCIDs.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Accesing the Data Science resource using FQDN.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// Data Science resource type.
	DataScienceResourceType DataScienceResourceTypeEnum `mandatory:"false" json:"dataScienceResourceType,omitempty"`
}

func (m DataSciencePrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSciencePrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataSciencePrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDataSciencePrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDataScienceResourceTypeEnum(string(m.DataScienceResourceType)); !ok && m.DataScienceResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataScienceResourceType: %s. Supported values are: %s.", m.DataScienceResourceType, strings.Join(GetDataScienceResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
