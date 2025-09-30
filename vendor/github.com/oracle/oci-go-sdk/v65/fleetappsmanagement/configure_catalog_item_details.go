// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigureCatalogItemDetails Contains the details required to configure a catalog item.
type ConfigureCatalogItemDetails struct {

	// The Oracle Cloud Object Storage namespace where the artifact or variables are stored.
	StorageNamespace *string `mandatory:"true" json:"storageNamespace"`

	// The name of the Object Storage bucket that contains the catalog item configuration object.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// OCID of the Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Array of key value pairs specifying variables or parameters to be used when configuring the catalog item.
	InputVariables []KeyValueProperty `mandatory:"true" json:"inputVariables"`

	// The name of the object file in the specified bucket containing catalog item configuration details.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// The unique identifier of the catalog listing to which configuration applies.
	CatalogListingId *string `mandatory:"false" json:"catalogListingId"`

	// The unique identifier for the specific version of the catalog listing.
	CatalogListingVersionId *string `mandatory:"false" json:"catalogListingVersionId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ConfigureCatalogItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigureCatalogItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
