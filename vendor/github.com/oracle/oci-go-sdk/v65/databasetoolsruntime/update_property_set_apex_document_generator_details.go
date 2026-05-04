// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdatePropertySetApexDocumentGeneratorDetails Contains the update details of an APEX document generator property set
type UpdatePropertySetApexDocumentGeneratorDetails struct {

	// The name of the credential used by APEX to manage Object Storage Buckets and Objects as well as invoke the Document Generator function.
	CredentialKey *string `mandatory:"false" json:"credentialKey"`

	// The Object Storage Namespace containing the Object Storage Buckets managed by APEX
	ObjectStorageNamespace *string `mandatory:"false" json:"objectStorageNamespace"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Object Storage Buckets managed by APEX
	ObjectStorageBucketCompartmentId *string `mandatory:"false" json:"objectStorageBucketCompartmentId"`

	// Object Storage Endpoint
	ObjectStorageEndpoint *string `mandatory:"false" json:"objectStorageEndpoint"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Document Generator function
	FunctionId *string `mandatory:"false" json:"functionId"`

	// The base endpoint URL to use to invoke the Document Generator function
	InvokeEndpoint *string `mandatory:"false" json:"invokeEndpoint"`

	// The print server type
	PrintServerType PropertySetApexDocumentGeneratorPrintServerTypeEnum `mandatory:"false" json:"printServerType,omitempty"`

	// The status of the Autonomous Database Serverless Resource Principal (OCI$RESOURCE_PRINCIPAL)
	AutonomousDatabaseResourcePrincipalStatus PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum `mandatory:"false" json:"autonomousDatabaseResourcePrincipalStatus,omitempty"`
}

func (m UpdatePropertySetApexDocumentGeneratorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePropertySetApexDocumentGeneratorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPropertySetApexDocumentGeneratorPrintServerTypeEnum(string(m.PrintServerType)); !ok && m.PrintServerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrintServerType: %s. Supported values are: %s.", m.PrintServerType, strings.Join(GetPropertySetApexDocumentGeneratorPrintServerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum(string(m.AutonomousDatabaseResourcePrincipalStatus)); !ok && m.AutonomousDatabaseResourcePrincipalStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutonomousDatabaseResourcePrincipalStatus: %s. Supported values are: %s.", m.AutonomousDatabaseResourcePrincipalStatus, strings.Join(GetPropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatePropertySetApexDocumentGeneratorDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatePropertySetApexDocumentGeneratorDetails UpdatePropertySetApexDocumentGeneratorDetails
	s := struct {
		DiscriminatorParam string `json:"key"`
		MarshalTypeUpdatePropertySetApexDocumentGeneratorDetails
	}{
		"APEX_DOCUMENT_GENERATOR",
		(MarshalTypeUpdatePropertySetApexDocumentGeneratorDetails)(m),
	}

	return json.Marshal(&s)
}
