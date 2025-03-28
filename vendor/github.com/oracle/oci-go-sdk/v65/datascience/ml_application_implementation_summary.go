// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// MlApplicationImplementationSummary Summary of the MlApplicationImplementation.
type MlApplicationImplementationSummary struct {

	// The OCID of the MlApplicationImplementation. Unique identifier that is immutable after creation.
	Id *string `mandatory:"true" json:"id"`

	// ML Application Implementation name which is unique for given ML Application.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the ML Application implemented by this ML Application Implementation.
	MlApplicationId *string `mandatory:"true" json:"mlApplicationId"`

	// The name of ML Application (based on mlApplicationId).
	MlApplicationName *string `mandatory:"true" json:"mlApplicationName"`

	// The OCID of the compartment where the MlApplicationImplementation is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the MlApplicationImplementation was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the ML Application Implementation.
	LifecycleState MlApplicationImplementationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Description of ML Application Implementation defined in ML Application package descriptor
	Description *string `mandatory:"false" json:"description"`

	// The version of ML Application Package (e.g. "1.2" or "2.0.4") defined in ML Application package descriptor. Value is not mandatory only for CREATING state otherwise it must be always presented.
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// Schema of configuration which needs to be provided for each ML Application Instance. It is defined in the ML Application package descriptor.
	ConfigurationSchema []ConfigurationPropertySchema `mandatory:"false" json:"configurationSchema"`

	// List of ML Application Implementation OCIDs for which migration from this implementation is allowed. Migration means that if consumers change implementation for their instances to implementation with OCID from this list, instance components will be updated in place otherwise new instance components are created based on the new implementation and old instance components are removed.
	AllowedMigrationDestinations []string `mandatory:"false" json:"allowedMigrationDestinations"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MlApplicationImplementationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlApplicationImplementationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationImplementationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMlApplicationImplementationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
