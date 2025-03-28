// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generic Artifacts Content API
//
// API covering the Generic Artifacts Service content
// Use this API to put and get generic artifact content.
//

package genericartifactscontent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenericArtifact The metadata of the artifact.
type GenericArtifact struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the artifact.
	// Example: `ocid1.genericartifact.oc1..exampleuniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The artifact name with the format of `<artifact-path>:<artifact-version>`. The artifact name is truncated to a maximum length of 255.
	// Example: `project01/my-web-app/artifact-abc:1.0.0`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// A user-defined path to describe the location of an artifact. Slashes do not create a directory structure, but you can use slashes to organize the repository. An artifact path does not include an artifact version.
	// Example: `project01/my-web-app/artifact-abc`
	ArtifactPath *string `mandatory:"true" json:"artifactPath"`

	// A user-defined string to describe the artifact version.
	// Example: `1.1.0` or `1.2-beta-2`
	Version *string `mandatory:"true" json:"version"`

	// The SHA256 digest for the artifact. When you upload an artifact to the repository, a SHA256 digest is calculated and added to the artifact properties.
	Sha256 *string `mandatory:"true" json:"sha256"`

	// The size of the artifact in bytes.
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`

	// The current state of the artifact.
	LifecycleState GenericArtifactLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// An RFC 3339 timestamp indicating when the repository was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m GenericArtifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericArtifact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenericArtifactLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGenericArtifactLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenericArtifactLifecycleStateEnum Enum with underlying type: string
type GenericArtifactLifecycleStateEnum string

// Set of constants representing the allowable values for GenericArtifactLifecycleStateEnum
const (
	GenericArtifactLifecycleStateAvailable GenericArtifactLifecycleStateEnum = "AVAILABLE"
	GenericArtifactLifecycleStateDeleting  GenericArtifactLifecycleStateEnum = "DELETING"
	GenericArtifactLifecycleStateDeleted   GenericArtifactLifecycleStateEnum = "DELETED"
)

var mappingGenericArtifactLifecycleStateEnum = map[string]GenericArtifactLifecycleStateEnum{
	"AVAILABLE": GenericArtifactLifecycleStateAvailable,
	"DELETING":  GenericArtifactLifecycleStateDeleting,
	"DELETED":   GenericArtifactLifecycleStateDeleted,
}

var mappingGenericArtifactLifecycleStateEnumLowerCase = map[string]GenericArtifactLifecycleStateEnum{
	"available": GenericArtifactLifecycleStateAvailable,
	"deleting":  GenericArtifactLifecycleStateDeleting,
	"deleted":   GenericArtifactLifecycleStateDeleted,
}

// GetGenericArtifactLifecycleStateEnumValues Enumerates the set of values for GenericArtifactLifecycleStateEnum
func GetGenericArtifactLifecycleStateEnumValues() []GenericArtifactLifecycleStateEnum {
	values := make([]GenericArtifactLifecycleStateEnum, 0)
	for _, v := range mappingGenericArtifactLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGenericArtifactLifecycleStateEnumStringValues Enumerates the set of values in String for GenericArtifactLifecycleStateEnum
func GetGenericArtifactLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingGenericArtifactLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenericArtifactLifecycleStateEnum(val string) (GenericArtifactLifecycleStateEnum, bool) {
	enum, ok := mappingGenericArtifactLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
