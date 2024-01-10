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

// RotateCloudAutonomousVmClusterOrdsCertsDetails The details for configuring the ORDS certificates on Cloud Autonomous VM Cluster
type RotateCloudAutonomousVmClusterOrdsCertsDetails struct {

	// Specify SYSTEM to use Oracle-managed certificates. Specify BYOC when you want to bring your own certificate.
	CertificateGenerationType RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum `mandatory:"true" json:"certificateGenerationType"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate to use.
	CertificateId *string `mandatory:"false" json:"certificateId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate authority.
	CertificateAuthorityId *string `mandatory:"false" json:"certificateAuthorityId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the certificate bundle.
	CaBundleId *string `mandatory:"false" json:"caBundleId"`
}

func (m RotateCloudAutonomousVmClusterOrdsCertsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RotateCloudAutonomousVmClusterOrdsCertsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum(string(m.CertificateGenerationType)); !ok && m.CertificateGenerationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateGenerationType: %s. Supported values are: %s.", m.CertificateGenerationType, strings.Join(GetRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum Enum with underlying type: string
type RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum string

// Set of constants representing the allowable values for RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum
const (
	RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeSystem RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum = "SYSTEM"
	RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeByoc   RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum = "BYOC"
)

var mappingRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum = map[string]RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum{
	"SYSTEM": RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeSystem,
	"BYOC":   RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeByoc,
}

var mappingRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumLowerCase = map[string]RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum{
	"system": RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeSystem,
	"byoc":   RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeByoc,
}

// GetRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumValues Enumerates the set of values for RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum
func GetRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumValues() []RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum {
	values := make([]RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum, 0)
	for _, v := range mappingRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumStringValues Enumerates the set of values in String for RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum
func GetRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"BYOC",
	}
}

// GetMappingRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum(val string) (RotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum, bool) {
	enum, ok := mappingRotateCloudAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
