// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AkamaiManualStreamCdnConfig Configuration fields for manual Akamai configuration.
type AkamaiManualStreamCdnConfig struct {

	// The shared secret key A, two for errorless key rotation.
	OriginAuthSecretKeyA *string `mandatory:"false" json:"originAuthSecretKeyA"`

	// Nonce identifier for originAuthSecretKeyA (used to determine key used to sign).
	OriginAuthSecretKeyNonceA *string `mandatory:"false" json:"originAuthSecretKeyNonceA"`

	// The shared secret key B, two for errorless key rotation.
	OriginAuthSecretKeyB *string `mandatory:"false" json:"originAuthSecretKeyB"`

	// Nonce identifier for originAuthSecretKeyB (used to determine key used to sign).
	OriginAuthSecretKeyNonceB *string `mandatory:"false" json:"originAuthSecretKeyNonceB"`

	// The hostname of the CDN edge server to use when building CDN URLs.
	EdgeHostname *string `mandatory:"false" json:"edgeHostname"`

	// The path to prepend when building CDN URLs.
	EdgePathPrefix *string `mandatory:"false" json:"edgePathPrefix"`

	// Whether token authentication should be used at the CDN edge.
	IsEdgeTokenAuth *bool `mandatory:"false" json:"isEdgeTokenAuth"`

	// The encryption key to use for edge token authentication.
	EdgeTokenKey *string `mandatory:"false" json:"edgeTokenKey"`

	// Salt to use when encrypting authentication token.
	EdgeTokenSalt *string `mandatory:"false" json:"edgeTokenSalt"`

	// The type of data used to compute the signature.
	OriginAuthSignType AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum `mandatory:"false" json:"originAuthSignType,omitempty"`

	// The type of encryption used to compute the signature.
	OriginAuthSignEncryption AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum `mandatory:"false" json:"originAuthSignEncryption,omitempty"`
}

func (m AkamaiManualStreamCdnConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AkamaiManualStreamCdnConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAkamaiManualStreamCdnConfigOriginAuthSignTypeEnum(string(m.OriginAuthSignType)); !ok && m.OriginAuthSignType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OriginAuthSignType: %s. Supported values are: %s.", m.OriginAuthSignType, strings.Join(GetAkamaiManualStreamCdnConfigOriginAuthSignTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum(string(m.OriginAuthSignEncryption)); !ok && m.OriginAuthSignEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OriginAuthSignEncryption: %s. Supported values are: %s.", m.OriginAuthSignEncryption, strings.Join(GetAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AkamaiManualStreamCdnConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAkamaiManualStreamCdnConfig AkamaiManualStreamCdnConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAkamaiManualStreamCdnConfig
	}{
		"AKAMAI_MANUAL",
		(MarshalTypeAkamaiManualStreamCdnConfig)(m),
	}

	return json.Marshal(&s)
}

// AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum Enum with underlying type: string
type AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum string

// Set of constants representing the allowable values for AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum
const (
	AkamaiManualStreamCdnConfigOriginAuthSignTypeForwardurl AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum = "ForwardURL"
)

var mappingAkamaiManualStreamCdnConfigOriginAuthSignTypeEnum = map[string]AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum{
	"ForwardURL": AkamaiManualStreamCdnConfigOriginAuthSignTypeForwardurl,
}

var mappingAkamaiManualStreamCdnConfigOriginAuthSignTypeEnumLowerCase = map[string]AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum{
	"forwardurl": AkamaiManualStreamCdnConfigOriginAuthSignTypeForwardurl,
}

// GetAkamaiManualStreamCdnConfigOriginAuthSignTypeEnumValues Enumerates the set of values for AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum
func GetAkamaiManualStreamCdnConfigOriginAuthSignTypeEnumValues() []AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum {
	values := make([]AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum, 0)
	for _, v := range mappingAkamaiManualStreamCdnConfigOriginAuthSignTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAkamaiManualStreamCdnConfigOriginAuthSignTypeEnumStringValues Enumerates the set of values in String for AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum
func GetAkamaiManualStreamCdnConfigOriginAuthSignTypeEnumStringValues() []string {
	return []string{
		"ForwardURL",
	}
}

// GetMappingAkamaiManualStreamCdnConfigOriginAuthSignTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAkamaiManualStreamCdnConfigOriginAuthSignTypeEnum(val string) (AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum, bool) {
	enum, ok := mappingAkamaiManualStreamCdnConfigOriginAuthSignTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum Enum with underlying type: string
type AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum string

// Set of constants representing the allowable values for AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum
const (
	AkamaiManualStreamCdnConfigOriginAuthSignEncryptionSha256Hmac AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum = "SHA256-HMAC"
)

var mappingAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum = map[string]AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum{
	"SHA256-HMAC": AkamaiManualStreamCdnConfigOriginAuthSignEncryptionSha256Hmac,
}

var mappingAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnumLowerCase = map[string]AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum{
	"sha256-hmac": AkamaiManualStreamCdnConfigOriginAuthSignEncryptionSha256Hmac,
}

// GetAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnumValues Enumerates the set of values for AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum
func GetAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnumValues() []AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum {
	values := make([]AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum, 0)
	for _, v := range mappingAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnumStringValues Enumerates the set of values in String for AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum
func GetAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnumStringValues() []string {
	return []string{
		"SHA256-HMAC",
	}
}

// GetMappingAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum(val string) (AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum, bool) {
	enum, ok := mappingAkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
