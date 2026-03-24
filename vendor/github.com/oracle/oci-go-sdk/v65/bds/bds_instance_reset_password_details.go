// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// BdsInstanceResetPasswordDetails The request body for resetting the password of indicated component.
type BdsInstanceResetPasswordDetails struct {

	// Target service to which this operation applies.
	Service BdsInstanceResetPasswordDetailsServiceEnum `mandatory:"true" json:"service"`

	// Base-64 encoded password for the cluster admin user.
	ClusterAdminPassword *string `mandatory:"false" json:"clusterAdminPassword"`

	// The secretId for the cluster admin user.
	SecretId *string `mandatory:"false" json:"secretId"`
}

func (m BdsInstanceResetPasswordDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsInstanceResetPasswordDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsInstanceResetPasswordDetailsServiceEnum(string(m.Service)); !ok && m.Service != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Service: %s. Supported values are: %s.", m.Service, strings.Join(GetBdsInstanceResetPasswordDetailsServiceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsInstanceResetPasswordDetailsServiceEnum Enum with underlying type: string
type BdsInstanceResetPasswordDetailsServiceEnum string

// Set of constants representing the allowable values for BdsInstanceResetPasswordDetailsServiceEnum
const (
	BdsInstanceResetPasswordDetailsServiceAmbari     BdsInstanceResetPasswordDetailsServiceEnum = "AMBARI"
	BdsInstanceResetPasswordDetailsServiceHue        BdsInstanceResetPasswordDetailsServiceEnum = "HUE"
	BdsInstanceResetPasswordDetailsServiceRanger     BdsInstanceResetPasswordDetailsServiceEnum = "RANGER"
	BdsInstanceResetPasswordDetailsServiceJupyterhub BdsInstanceResetPasswordDetailsServiceEnum = "JUPYTERHUB"
)

var mappingBdsInstanceResetPasswordDetailsServiceEnum = map[string]BdsInstanceResetPasswordDetailsServiceEnum{
	"AMBARI":     BdsInstanceResetPasswordDetailsServiceAmbari,
	"HUE":        BdsInstanceResetPasswordDetailsServiceHue,
	"RANGER":     BdsInstanceResetPasswordDetailsServiceRanger,
	"JUPYTERHUB": BdsInstanceResetPasswordDetailsServiceJupyterhub,
}

var mappingBdsInstanceResetPasswordDetailsServiceEnumLowerCase = map[string]BdsInstanceResetPasswordDetailsServiceEnum{
	"ambari":     BdsInstanceResetPasswordDetailsServiceAmbari,
	"hue":        BdsInstanceResetPasswordDetailsServiceHue,
	"ranger":     BdsInstanceResetPasswordDetailsServiceRanger,
	"jupyterhub": BdsInstanceResetPasswordDetailsServiceJupyterhub,
}

// GetBdsInstanceResetPasswordDetailsServiceEnumValues Enumerates the set of values for BdsInstanceResetPasswordDetailsServiceEnum
func GetBdsInstanceResetPasswordDetailsServiceEnumValues() []BdsInstanceResetPasswordDetailsServiceEnum {
	values := make([]BdsInstanceResetPasswordDetailsServiceEnum, 0)
	for _, v := range mappingBdsInstanceResetPasswordDetailsServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsInstanceResetPasswordDetailsServiceEnumStringValues Enumerates the set of values in String for BdsInstanceResetPasswordDetailsServiceEnum
func GetBdsInstanceResetPasswordDetailsServiceEnumStringValues() []string {
	return []string{
		"AMBARI",
		"HUE",
		"RANGER",
		"JUPYTERHUB",
	}
}

// GetMappingBdsInstanceResetPasswordDetailsServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsInstanceResetPasswordDetailsServiceEnum(val string) (BdsInstanceResetPasswordDetailsServiceEnum, bool) {
	enum, ok := mappingBdsInstanceResetPasswordDetailsServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
