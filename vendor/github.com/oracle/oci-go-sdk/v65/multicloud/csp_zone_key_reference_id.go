// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CspZoneKeyReferenceId This is CSP zone key reference
type CspZoneKeyReferenceId struct {

	// Value of keyName
	//    GcpProjectName: A human-readable name for your project. The project name isn't used by any Google APIs. You can edit the project name at any time during or after project creation. Project names do not need to be unique.
	//    AzureSubscriptionId: A unique alphanumeric string that identifies your Azure subscription.
	//    AwsAccountId: a unique 12-digit number that identifies an Amazon Web Services (AWS) account
	KeyValue *string `mandatory:"true" json:"keyValue"`

	// KeyName for Azure=AzureSubscriptionId Aws=AwsAccountId GCP=GcpProjectName
	KeyName *string `mandatory:"true" json:"keyName"`
}

func (m CspZoneKeyReferenceId) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CspZoneKeyReferenceId) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
