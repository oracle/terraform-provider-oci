// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelPackage Details of `Package` object.
type ModelPackage struct {

	// Unique immutable identifier that was assigned when the Package was registered.
	Id *string `mandatory:"true" json:"id"`

	// ID of the publisher providing the package.
	PublisherId *string `mandatory:"true" json:"publisherId"`

	// Name of package.
	Name *string `mandatory:"true" json:"name"`

	// Display name for the package (displayed in UI and user-facing applications).
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Version of the package.
	Version *string `mandatory:"true" json:"version"`

	// When the package was uploaded. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUploaded *common.SDKTime `mandatory:"true" json:"timeUploaded"`

	// When the package was last published. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimePublished *common.SDKTime `mandatory:"true" json:"timePublished"`

	// Description of the package.
	Description *string `mandatory:"true" json:"description"`

	// A list of resource types describing the content of the package.
	ResourceTypes []string `mandatory:"true" json:"resourceTypes"`

	// A map of resource type to metadata key/value map that further describes the content for the resource types in this package.. Keys are resource type names, values are a map of name/value pairs per resource type.
	ResourceTypesMetadata []ResourceTypeMetadata `mandatory:"true" json:"resourceTypesMetadata"`

	// A map of metadata key/value pairs that further describes the publisher and the platform in which the package might be used.
	PublisherMetadata []MetadataProperty `mandatory:"true" json:"publisherMetadata"`

	ImportContract *ImportContract `mandatory:"true" json:"importContract"`

	DefaultParameterValues *DefaultParameterValues `mandatory:"true" json:"defaultParameterValues"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ModelPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
