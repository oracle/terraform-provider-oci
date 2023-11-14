// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"strings"
)

// ServiceEnum Enum with underlying type: string
type ServiceEnum string

// Set of constants representing the allowable values for ServiceEnum
const (
	ServiceZookeeper ServiceEnum = "ZOOKEEPER"
	ServiceAms       ServiceEnum = "AMS"
	ServiceHdfs      ServiceEnum = "HDFS"
	ServiceYarn      ServiceEnum = "YARN"
	ServiceMapreduce ServiceEnum = "MAPREDUCE"
	ServiceOozie     ServiceEnum = "OOZIE"
	ServiceHbase     ServiceEnum = "HBASE"
	ServiceSpark     ServiceEnum = "SPARK"
	ServiceHive      ServiceEnum = "HIVE"
	ServiceKafka     ServiceEnum = "KAFKA"
	ServiceFlink     ServiceEnum = "FLINK"
	ServiceRegistry  ServiceEnum = "REGISTRY"
)

var mappingServiceEnum = map[string]ServiceEnum{
	"ZOOKEEPER": ServiceZookeeper,
	"AMS":       ServiceAms,
	"HDFS":      ServiceHdfs,
	"YARN":      ServiceYarn,
	"MAPREDUCE": ServiceMapreduce,
	"OOZIE":     ServiceOozie,
	"HBASE":     ServiceHbase,
	"SPARK":     ServiceSpark,
	"HIVE":      ServiceHive,
	"KAFKA":     ServiceKafka,
	"FLINK":     ServiceFlink,
	"REGISTRY":  ServiceRegistry,
}

var mappingServiceEnumLowerCase = map[string]ServiceEnum{
	"zookeeper": ServiceZookeeper,
	"ams":       ServiceAms,
	"hdfs":      ServiceHdfs,
	"yarn":      ServiceYarn,
	"mapreduce": ServiceMapreduce,
	"oozie":     ServiceOozie,
	"hbase":     ServiceHbase,
	"spark":     ServiceSpark,
	"hive":      ServiceHive,
	"kafka":     ServiceKafka,
	"flink":     ServiceFlink,
	"registry":  ServiceRegistry,
}

// GetServiceEnumValues Enumerates the set of values for ServiceEnum
func GetServiceEnumValues() []ServiceEnum {
	values := make([]ServiceEnum, 0)
	for _, v := range mappingServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceEnumStringValues Enumerates the set of values in String for ServiceEnum
func GetServiceEnumStringValues() []string {
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

// GetMappingServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceEnum(val string) (ServiceEnum, bool) {
	enum, ok := mappingServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
