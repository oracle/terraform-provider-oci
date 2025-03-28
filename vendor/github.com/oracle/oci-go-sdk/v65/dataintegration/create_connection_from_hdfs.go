// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateConnectionFromHdfs The details to create the HDFS data asset connection.
type CreateConnectionFromHdfs struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The HDFS principal.
	HdfsPrincipal *string `mandatory:"true" json:"hdfsPrincipal"`

	// The HDFS Data Node principal.
	DataNodePrincipal *string `mandatory:"true" json:"dataNodePrincipal"`

	// The HDFS Name Node principal.
	NameNodePrincipal *string `mandatory:"true" json:"nameNodePrincipal"`

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// User-defined description for the connection.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// HDFS Realm name.
	Realm *string `mandatory:"false" json:"realm"`

	// The HDFS Key Distribution Center.
	KeyDistributionCenter *string `mandatory:"false" json:"keyDistributionCenter"`

	KeyTabContent *SensitiveAttribute `mandatory:"false" json:"keyTabContent"`
}

// GetKey returns Key
func (m CreateConnectionFromHdfs) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m CreateConnectionFromHdfs) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m CreateConnectionFromHdfs) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m CreateConnectionFromHdfs) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateConnectionFromHdfs) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m CreateConnectionFromHdfs) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m CreateConnectionFromHdfs) GetIdentifier() *string {
	return m.Identifier
}

// GetConnectionProperties returns ConnectionProperties
func (m CreateConnectionFromHdfs) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

// GetRegistryMetadata returns RegistryMetadata
func (m CreateConnectionFromHdfs) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateConnectionFromHdfs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateConnectionFromHdfs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateConnectionFromHdfs) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateConnectionFromHdfs CreateConnectionFromHdfs
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateConnectionFromHdfs
	}{
		"HDFS_CONNECTION",
		(MarshalTypeCreateConnectionFromHdfs)(m),
	}

	return json.Marshal(&s)
}
