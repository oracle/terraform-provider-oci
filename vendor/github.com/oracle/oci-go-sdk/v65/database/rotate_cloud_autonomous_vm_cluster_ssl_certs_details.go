// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RotateCloudAutonomousVmClusterSslCertsDetails The details for configuring the SSL certificates on Cloud Autonomous VM Cluster
type RotateCloudAutonomousVmClusterSslCertsDetails struct {

	// Specify SYSTEM to use Oracle-managed certificates. Specify BYOC when you want to bring your own certificate.
	CertificateGenerationType RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum `mandatory:"true" json:"certificateGenerationType"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate to use.
	CertificateId *string `mandatory:"false" json:"certificateId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate authority.
	CertificateAuthorityId *string `mandatory:"false" json:"certificateAuthorityId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate bundle.
	CaBundleId *string `mandatory:"false" json:"caBundleId"`
}

func (m RotateCloudAutonomousVmClusterSslCertsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RotateCloudAutonomousVmClusterSslCertsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum(string(m.CertificateGenerationType)); !ok && m.CertificateGenerationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateGenerationType: %s. Supported values are: %s.", m.CertificateGenerationType, strings.Join(GetRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum Enum with underlying type: string
type RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum string

// Set of constants representing the allowable values for RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum
const (
	RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeSystem RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum = "SYSTEM"
	RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeByoc   RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum = "BYOC"
)

var mappingRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum = map[string]RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum{
	"SYSTEM": RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeSystem,
	"BYOC":   RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeByoc,
}

var mappingRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumLowerCase = map[string]RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum{
	"system": RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeSystem,
	"byoc":   RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeByoc,
}

// GetRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumValues Enumerates the set of values for RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum
func GetRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumValues() []RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum {
	values := make([]RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum, 0)
	for _, v := range mappingRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumStringValues Enumerates the set of values in String for RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum
func GetRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"BYOC",
	}
}

// GetMappingRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum(val string) (RotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnum, bool) {
	enum, ok := mappingRotateCloudAutonomousVmClusterSslCertsDetailsCertificateGenerationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
