// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// PostInstallationActionsEnum Enum with underlying type: string
type PostInstallationActionsEnum string

// Set of constants representing the allowable values for PostInstallationActionsEnum
const (
	PostInstallationActionsChangeMinimumKeyLengthForEc            PostInstallationActionsEnum = "CHANGE_MINIMUM_KEY_LENGTH_FOR_EC"
	PostInstallationActionsChangeMinimumKeyLengthForDsa           PostInstallationActionsEnum = "CHANGE_MINIMUM_KEY_LENGTH_FOR_DSA"
	PostInstallationActionsChangeMinimumKeyLengthForDiffieHellman PostInstallationActionsEnum = "CHANGE_MINIMUM_KEY_LENGTH_FOR_DIFFIE_HELLMAN"
	PostInstallationActionsChangeMinimumKeyLengthForRsaSignedJars PostInstallationActionsEnum = "CHANGE_MINIMUM_KEY_LENGTH_FOR_RSA_SIGNED_JARS"
	PostInstallationActionsDisableTls                             PostInstallationActionsEnum = "DISABLE_TLS"
	PostInstallationActionsUseOsCacerts                           PostInstallationActionsEnum = "USE_OS_CACERTS"
	PostInstallationActionsUseSystemProxies                       PostInstallationActionsEnum = "USE_SYSTEM_PROXIES"
	PostInstallationActionsSetupHttpProxy                         PostInstallationActionsEnum = "SETUP_HTTP_PROXY"
	PostInstallationActionsSetupHttpsProxy                        PostInstallationActionsEnum = "SETUP_HTTPS_PROXY"
	PostInstallationActionsSetupFtpProxy                          PostInstallationActionsEnum = "SETUP_FTP_PROXY"
	PostInstallationActionsSetupSocksProxy                        PostInstallationActionsEnum = "SETUP_SOCKS_PROXY"
	PostInstallationActionsAddFileHandler                         PostInstallationActionsEnum = "ADD_FILE_HANDLER"
	PostInstallationActionsLoggingLevel                           PostInstallationActionsEnum = "LOGGING_LEVEL"
)

var mappingPostInstallationActionsEnum = map[string]PostInstallationActionsEnum{
	"CHANGE_MINIMUM_KEY_LENGTH_FOR_EC":              PostInstallationActionsChangeMinimumKeyLengthForEc,
	"CHANGE_MINIMUM_KEY_LENGTH_FOR_DSA":             PostInstallationActionsChangeMinimumKeyLengthForDsa,
	"CHANGE_MINIMUM_KEY_LENGTH_FOR_DIFFIE_HELLMAN":  PostInstallationActionsChangeMinimumKeyLengthForDiffieHellman,
	"CHANGE_MINIMUM_KEY_LENGTH_FOR_RSA_SIGNED_JARS": PostInstallationActionsChangeMinimumKeyLengthForRsaSignedJars,
	"DISABLE_TLS":        PostInstallationActionsDisableTls,
	"USE_OS_CACERTS":     PostInstallationActionsUseOsCacerts,
	"USE_SYSTEM_PROXIES": PostInstallationActionsUseSystemProxies,
	"SETUP_HTTP_PROXY":   PostInstallationActionsSetupHttpProxy,
	"SETUP_HTTPS_PROXY":  PostInstallationActionsSetupHttpsProxy,
	"SETUP_FTP_PROXY":    PostInstallationActionsSetupFtpProxy,
	"SETUP_SOCKS_PROXY":  PostInstallationActionsSetupSocksProxy,
	"ADD_FILE_HANDLER":   PostInstallationActionsAddFileHandler,
	"LOGGING_LEVEL":      PostInstallationActionsLoggingLevel,
}

var mappingPostInstallationActionsEnumLowerCase = map[string]PostInstallationActionsEnum{
	"change_minimum_key_length_for_ec":              PostInstallationActionsChangeMinimumKeyLengthForEc,
	"change_minimum_key_length_for_dsa":             PostInstallationActionsChangeMinimumKeyLengthForDsa,
	"change_minimum_key_length_for_diffie_hellman":  PostInstallationActionsChangeMinimumKeyLengthForDiffieHellman,
	"change_minimum_key_length_for_rsa_signed_jars": PostInstallationActionsChangeMinimumKeyLengthForRsaSignedJars,
	"disable_tls":        PostInstallationActionsDisableTls,
	"use_os_cacerts":     PostInstallationActionsUseOsCacerts,
	"use_system_proxies": PostInstallationActionsUseSystemProxies,
	"setup_http_proxy":   PostInstallationActionsSetupHttpProxy,
	"setup_https_proxy":  PostInstallationActionsSetupHttpsProxy,
	"setup_ftp_proxy":    PostInstallationActionsSetupFtpProxy,
	"setup_socks_proxy":  PostInstallationActionsSetupSocksProxy,
	"add_file_handler":   PostInstallationActionsAddFileHandler,
	"logging_level":      PostInstallationActionsLoggingLevel,
}

// GetPostInstallationActionsEnumValues Enumerates the set of values for PostInstallationActionsEnum
func GetPostInstallationActionsEnumValues() []PostInstallationActionsEnum {
	values := make([]PostInstallationActionsEnum, 0)
	for _, v := range mappingPostInstallationActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetPostInstallationActionsEnumStringValues Enumerates the set of values in String for PostInstallationActionsEnum
func GetPostInstallationActionsEnumStringValues() []string {
	return []string{
		"CHANGE_MINIMUM_KEY_LENGTH_FOR_EC",
		"CHANGE_MINIMUM_KEY_LENGTH_FOR_DSA",
		"CHANGE_MINIMUM_KEY_LENGTH_FOR_DIFFIE_HELLMAN",
		"CHANGE_MINIMUM_KEY_LENGTH_FOR_RSA_SIGNED_JARS",
		"DISABLE_TLS",
		"USE_OS_CACERTS",
		"USE_SYSTEM_PROXIES",
		"SETUP_HTTP_PROXY",
		"SETUP_HTTPS_PROXY",
		"SETUP_FTP_PROXY",
		"SETUP_SOCKS_PROXY",
		"ADD_FILE_HANDLER",
		"LOGGING_LEVEL",
	}
}

// GetMappingPostInstallationActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPostInstallationActionsEnum(val string) (PostInstallationActionsEnum, bool) {
	enum, ok := mappingPostInstallationActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
