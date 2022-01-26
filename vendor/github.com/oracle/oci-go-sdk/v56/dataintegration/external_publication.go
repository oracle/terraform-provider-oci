// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ExternalPublication The external published object contains the audit summary information and the definition of the task.
type ExternalPublication struct {

	// The unique OCID of the identifier that is returned after creating the Oracle Cloud Infrastructure Data Flow application.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The OCID of the compartment where the application is created in the Oracle Cloud Infrastructure Data Flow Service.
	ApplicationCompartmentId *string `mandatory:"false" json:"applicationCompartmentId"`

	// The name of the application.
	DisplayName *string `mandatory:"false" json:"displayName"`

	ResourceConfiguration *ResourceConfiguration `mandatory:"false" json:"resourceConfiguration"`

	ConfigurationDetails *ConfigurationDetails `mandatory:"false" json:"configurationDetails"`

	// The status of the publishing action to Oracle Cloud Infrastructure Data Flow.
	Status ExternalPublicationStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The error of the published object in the application.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object type.
	ModelType *string `mandatory:"false" json:"modelType"`

	// This is a version number that is used by the service to upgrade objects if needed through releases of the service.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects. Other values are reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m ExternalPublication) String() string {
	return common.PointerString(m)
}

// ExternalPublicationStatusEnum Enum with underlying type: string
type ExternalPublicationStatusEnum string

// Set of constants representing the allowable values for ExternalPublicationStatusEnum
const (
	ExternalPublicationStatusSuccessful ExternalPublicationStatusEnum = "SUCCESSFUL"
	ExternalPublicationStatusFailed     ExternalPublicationStatusEnum = "FAILED"
	ExternalPublicationStatusPublishing ExternalPublicationStatusEnum = "PUBLISHING"
)

var mappingExternalPublicationStatus = map[string]ExternalPublicationStatusEnum{
	"SUCCESSFUL": ExternalPublicationStatusSuccessful,
	"FAILED":     ExternalPublicationStatusFailed,
	"PUBLISHING": ExternalPublicationStatusPublishing,
}

// GetExternalPublicationStatusEnumValues Enumerates the set of values for ExternalPublicationStatusEnum
func GetExternalPublicationStatusEnumValues() []ExternalPublicationStatusEnum {
	values := make([]ExternalPublicationStatusEnum, 0)
	for _, v := range mappingExternalPublicationStatus {
		values = append(values, v)
	}
	return values
}
