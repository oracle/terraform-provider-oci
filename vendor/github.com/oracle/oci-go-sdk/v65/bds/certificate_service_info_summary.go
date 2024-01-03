// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateServiceInfoSummary List of TLS/SSL information of services
type CertificateServiceInfoSummary struct {

	// Name of the service
	Service CertificateServiceInfoSummaryServiceEnum `mandatory:"true" json:"service"`

	// Whether certificate is enabled or disabled
	ServiceCertificateStatus CertificateServiceInfoSummaryServiceCertificateStatusEnum `mandatory:"true" json:"serviceCertificateStatus"`

	// List of Host specific certificate details
	HostSpecificCertificateDetails []HostSpecificCertificateDetails `mandatory:"true" json:"hostSpecificCertificateDetails"`
}

func (m CertificateServiceInfoSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateServiceInfoSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCertificateServiceInfoSummaryServiceEnum(string(m.Service)); !ok && m.Service != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Service: %s. Supported values are: %s.", m.Service, strings.Join(GetCertificateServiceInfoSummaryServiceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCertificateServiceInfoSummaryServiceCertificateStatusEnum(string(m.ServiceCertificateStatus)); !ok && m.ServiceCertificateStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceCertificateStatus: %s. Supported values are: %s.", m.ServiceCertificateStatus, strings.Join(GetCertificateServiceInfoSummaryServiceCertificateStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CertificateServiceInfoSummaryServiceEnum Enum with underlying type: string
type CertificateServiceInfoSummaryServiceEnum string

// Set of constants representing the allowable values for CertificateServiceInfoSummaryServiceEnum
const (
	CertificateServiceInfoSummaryServiceZookeeper CertificateServiceInfoSummaryServiceEnum = "ZOOKEEPER"
	CertificateServiceInfoSummaryServiceAms       CertificateServiceInfoSummaryServiceEnum = "AMS"
	CertificateServiceInfoSummaryServiceHdfs      CertificateServiceInfoSummaryServiceEnum = "HDFS"
	CertificateServiceInfoSummaryServiceYarn      CertificateServiceInfoSummaryServiceEnum = "YARN"
	CertificateServiceInfoSummaryServiceMapreduce CertificateServiceInfoSummaryServiceEnum = "MAPREDUCE"
	CertificateServiceInfoSummaryServiceOozie     CertificateServiceInfoSummaryServiceEnum = "OOZIE"
	CertificateServiceInfoSummaryServiceHbase     CertificateServiceInfoSummaryServiceEnum = "HBASE"
	CertificateServiceInfoSummaryServiceSpark     CertificateServiceInfoSummaryServiceEnum = "SPARK"
	CertificateServiceInfoSummaryServiceHive      CertificateServiceInfoSummaryServiceEnum = "HIVE"
	CertificateServiceInfoSummaryServiceKafka     CertificateServiceInfoSummaryServiceEnum = "KAFKA"
	CertificateServiceInfoSummaryServiceFlink     CertificateServiceInfoSummaryServiceEnum = "FLINK"
	CertificateServiceInfoSummaryServiceRegistry  CertificateServiceInfoSummaryServiceEnum = "REGISTRY"
)

var mappingCertificateServiceInfoSummaryServiceEnum = map[string]CertificateServiceInfoSummaryServiceEnum{
	"ZOOKEEPER": CertificateServiceInfoSummaryServiceZookeeper,
	"AMS":       CertificateServiceInfoSummaryServiceAms,
	"HDFS":      CertificateServiceInfoSummaryServiceHdfs,
	"YARN":      CertificateServiceInfoSummaryServiceYarn,
	"MAPREDUCE": CertificateServiceInfoSummaryServiceMapreduce,
	"OOZIE":     CertificateServiceInfoSummaryServiceOozie,
	"HBASE":     CertificateServiceInfoSummaryServiceHbase,
	"SPARK":     CertificateServiceInfoSummaryServiceSpark,
	"HIVE":      CertificateServiceInfoSummaryServiceHive,
	"KAFKA":     CertificateServiceInfoSummaryServiceKafka,
	"FLINK":     CertificateServiceInfoSummaryServiceFlink,
	"REGISTRY":  CertificateServiceInfoSummaryServiceRegistry,
}

var mappingCertificateServiceInfoSummaryServiceEnumLowerCase = map[string]CertificateServiceInfoSummaryServiceEnum{
	"zookeeper": CertificateServiceInfoSummaryServiceZookeeper,
	"ams":       CertificateServiceInfoSummaryServiceAms,
	"hdfs":      CertificateServiceInfoSummaryServiceHdfs,
	"yarn":      CertificateServiceInfoSummaryServiceYarn,
	"mapreduce": CertificateServiceInfoSummaryServiceMapreduce,
	"oozie":     CertificateServiceInfoSummaryServiceOozie,
	"hbase":     CertificateServiceInfoSummaryServiceHbase,
	"spark":     CertificateServiceInfoSummaryServiceSpark,
	"hive":      CertificateServiceInfoSummaryServiceHive,
	"kafka":     CertificateServiceInfoSummaryServiceKafka,
	"flink":     CertificateServiceInfoSummaryServiceFlink,
	"registry":  CertificateServiceInfoSummaryServiceRegistry,
}

// GetCertificateServiceInfoSummaryServiceEnumValues Enumerates the set of values for CertificateServiceInfoSummaryServiceEnum
func GetCertificateServiceInfoSummaryServiceEnumValues() []CertificateServiceInfoSummaryServiceEnum {
	values := make([]CertificateServiceInfoSummaryServiceEnum, 0)
	for _, v := range mappingCertificateServiceInfoSummaryServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateServiceInfoSummaryServiceEnumStringValues Enumerates the set of values in String for CertificateServiceInfoSummaryServiceEnum
func GetCertificateServiceInfoSummaryServiceEnumStringValues() []string {
	return []string{
		"ZOOKEEPER",
		"AMS",
		"HDFS",
		"YARN",
		"MAPREDUCE",
		"OOZIE",
		"HBASE",
		"SPARK",
		"HIVE",
		"KAFKA",
		"FLINK",
		"REGISTRY",
	}
}

// GetMappingCertificateServiceInfoSummaryServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateServiceInfoSummaryServiceEnum(val string) (CertificateServiceInfoSummaryServiceEnum, bool) {
	enum, ok := mappingCertificateServiceInfoSummaryServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CertificateServiceInfoSummaryServiceCertificateStatusEnum Enum with underlying type: string
type CertificateServiceInfoSummaryServiceCertificateStatusEnum string

// Set of constants representing the allowable values for CertificateServiceInfoSummaryServiceCertificateStatusEnum
const (
	CertificateServiceInfoSummaryServiceCertificateStatusEnabled  CertificateServiceInfoSummaryServiceCertificateStatusEnum = "ENABLED"
	CertificateServiceInfoSummaryServiceCertificateStatusDisabled CertificateServiceInfoSummaryServiceCertificateStatusEnum = "DISABLED"
)

var mappingCertificateServiceInfoSummaryServiceCertificateStatusEnum = map[string]CertificateServiceInfoSummaryServiceCertificateStatusEnum{
	"ENABLED":  CertificateServiceInfoSummaryServiceCertificateStatusEnabled,
	"DISABLED": CertificateServiceInfoSummaryServiceCertificateStatusDisabled,
}

var mappingCertificateServiceInfoSummaryServiceCertificateStatusEnumLowerCase = map[string]CertificateServiceInfoSummaryServiceCertificateStatusEnum{
	"enabled":  CertificateServiceInfoSummaryServiceCertificateStatusEnabled,
	"disabled": CertificateServiceInfoSummaryServiceCertificateStatusDisabled,
}

// GetCertificateServiceInfoSummaryServiceCertificateStatusEnumValues Enumerates the set of values for CertificateServiceInfoSummaryServiceCertificateStatusEnum
func GetCertificateServiceInfoSummaryServiceCertificateStatusEnumValues() []CertificateServiceInfoSummaryServiceCertificateStatusEnum {
	values := make([]CertificateServiceInfoSummaryServiceCertificateStatusEnum, 0)
	for _, v := range mappingCertificateServiceInfoSummaryServiceCertificateStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateServiceInfoSummaryServiceCertificateStatusEnumStringValues Enumerates the set of values in String for CertificateServiceInfoSummaryServiceCertificateStatusEnum
func GetCertificateServiceInfoSummaryServiceCertificateStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCertificateServiceInfoSummaryServiceCertificateStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateServiceInfoSummaryServiceCertificateStatusEnum(val string) (CertificateServiceInfoSummaryServiceCertificateStatusEnum, bool) {
	enum, ok := mappingCertificateServiceInfoSummaryServiceCertificateStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
