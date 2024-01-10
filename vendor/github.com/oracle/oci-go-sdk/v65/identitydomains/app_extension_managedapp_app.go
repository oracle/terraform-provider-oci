// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AppExtensionManagedappApp Managed App
type AppExtensionManagedappApp struct {

	// If true, the accounts of the application are managed through an ICF connector bundle
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Connected *bool `mandatory:"false" json:"connected"`

	// If true, the managed app can be authoritative.
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	CanBeAuthoritative *bool `mandatory:"false" json:"canBeAuthoritative"`

	// If true, sync from the managed app will be performed as authoritative sync.
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsAuthoritative *bool `mandatory:"false" json:"isAuthoritative"`

	// If true, the managed app is an On-Premise app.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsOnPremiseApp *bool `mandatory:"false" json:"isOnPremiseApp"`

	// If true, the managed app is a directory.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsDirectory *bool `mandatory:"false" json:"isDirectory"`

	// If true, the managed app supports schema discovery.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsSchemaDiscoverySupported *bool `mandatory:"false" json:"isSchemaDiscoverySupported"`

	// If true, the managed app supports schema customization.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsSchemaCustomizationSupported *bool `mandatory:"false" json:"isSchemaCustomizationSupported"`

	// If true, sync run-time operations are enabled for this App.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	EnableSync *bool `mandatory:"false" json:"enableSync"`

	// If true, send sync summary as notification upon job completion.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	EnableSyncSummaryReportNotification *bool `mandatory:"false" json:"enableSyncSummaryReportNotification"`

	// If true, send activation email to new users created from authoritative sync.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	EnableAuthSyncNewUserNotification *bool `mandatory:"false" json:"enableAuthSyncNewUserNotification"`

	// If true, admin has granted consent to perform managed app run-time operations.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AdminConsentGranted *bool `mandatory:"false" json:"adminConsentGranted"`

	// If true, the managed app requires 3-legged OAuth for authorization.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsThreeLeggedOAuthEnabled *bool `mandatory:"false" json:"isThreeLeggedOAuthEnabled"`

	// If true, indicates that Oracle Identity Cloud Service can use two-legged OAuth to connect to this ManagedApp.
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsTwoLeggedOAuthEnabled *bool `mandatory:"false" json:"isTwoLeggedOAuthEnabled"`

	// Three legged OAuth provider name in Oracle Identity Cloud Service.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	ThreeLeggedOAuthProviderName *string `mandatory:"false" json:"threeLeggedOAuthProviderName"`

	// The most recent DateTime that the configuration of this App was updated. AppServices updates this timestamp whenever AppServices updates an App's configuration with respect to synchronization.
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	SyncConfigLastModified *string `mandatory:"false" json:"syncConfigLastModified"`

	// If true, then the account form will be displayed in the Oracle Identity Cloud Service UI to interactively create or update an account for this App. If a value is not specified for this attribute, a default value of \"false\" will be assumed as the value for this attribute.
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AccountFormVisible *bool `mandatory:"false" json:"accountFormVisible"`

	// IdentityBridges associated with this App
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	IdentityBridges []AppIdentityBridges `mandatory:"false" json:"identityBridges"`

	ConnectorBundle *AppConnectorBundle `mandatory:"false" json:"connectorBundle"`

	// ConnectorBundle configuration properties
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	BundleConfigurationProperties []AppBundleConfigurationProperties `mandatory:"false" json:"bundleConfigurationProperties"`

	// Object classes
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	ObjectClasses []AppObjectClasses `mandatory:"false" json:"objectClasses"`

	BundlePoolConfiguration *AppBundlePoolConfiguration `mandatory:"false" json:"bundlePoolConfiguration"`

	FlatFileConnectorBundle *AppFlatFileConnectorBundle `mandatory:"false" json:"flatFileConnectorBundle"`

	// Flat file connector bundle configuration properties
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	FlatFileBundleConfigurationProperties []AppFlatFileBundleConfigurationProperties `mandatory:"false" json:"flatFileBundleConfigurationProperties"`

	ThreeLeggedOAuthCredential *AppThreeLeggedOAuthCredential `mandatory:"false" json:"threeLeggedOAuthCredential"`
}

func (m AppExtensionManagedappApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionManagedappApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
