// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the <a href="https://docs.oracle.com/en-us/iaas/jms/doc/java-download.html">Java Download</a> feature of Java Management Service.
//

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DownloadUrl Download Url object for the Java artifact.
type DownloadUrl struct {

	// The URL for downloading the artifact.
	DownloadUrl *string `mandatory:"true" json:"downloadUrl"`

	// The type of download URL.
	DownloadUrlType DownloadUrlTypeEnum `mandatory:"true" json:"downloadUrlType"`
}

func (m DownloadUrl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DownloadUrl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDownloadUrlTypeEnum(string(m.DownloadUrlType)); !ok && m.DownloadUrlType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DownloadUrlType: %s. Supported values are: %s.", m.DownloadUrlType, strings.Join(GetDownloadUrlTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
