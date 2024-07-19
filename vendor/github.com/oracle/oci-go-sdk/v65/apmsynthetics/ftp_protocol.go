// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// FtpProtocolEnum Enum with underlying type: string
type FtpProtocolEnum string

// Set of constants representing the allowable values for FtpProtocolEnum
const (
	FtpProtocolFtp  FtpProtocolEnum = "FTP"
	FtpProtocolFtps FtpProtocolEnum = "FTPS"
	FtpProtocolSftp FtpProtocolEnum = "SFTP"
)

var mappingFtpProtocolEnum = map[string]FtpProtocolEnum{
	"FTP":  FtpProtocolFtp,
	"FTPS": FtpProtocolFtps,
	"SFTP": FtpProtocolSftp,
}

var mappingFtpProtocolEnumLowerCase = map[string]FtpProtocolEnum{
	"ftp":  FtpProtocolFtp,
	"ftps": FtpProtocolFtps,
	"sftp": FtpProtocolSftp,
}

// GetFtpProtocolEnumValues Enumerates the set of values for FtpProtocolEnum
func GetFtpProtocolEnumValues() []FtpProtocolEnum {
	values := make([]FtpProtocolEnum, 0)
	for _, v := range mappingFtpProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetFtpProtocolEnumStringValues Enumerates the set of values in String for FtpProtocolEnum
func GetFtpProtocolEnumStringValues() []string {
	return []string{
		"FTP",
		"FTPS",
		"SFTP",
	}
}

// GetMappingFtpProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFtpProtocolEnum(val string) (FtpProtocolEnum, bool) {
	enum, ok := mappingFtpProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
