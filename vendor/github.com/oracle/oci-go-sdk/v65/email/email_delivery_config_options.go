// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmailDeliveryConfigOptions Information about the config params for the delivery config resource.
type EmailDeliveryConfigOptions struct {

	// The maxOutboundConnections parameter provides the means to limit the rate at which connections are made to a destination domain.
	MaxOutboundConnections *int `mandatory:"false" json:"maxOutboundConnections"`

	// Value of maxOutboundConnections in backoff state
	BackoffMaxOutboundConnections *int `mandatory:"false" json:"backoffMaxOutboundConnections"`

	// The maxMsgPerConnection parameter limits the number of messages that can be transferred by a single connection
	MaxMsgPerConnection *int `mandatory:"false" json:"maxMsgPerConnection"`

	// The maxMessageRate parameter provides the means to limit the rate at which messages are delivered to a destination domain.
	MaxMessageRate *string `mandatory:"false" json:"maxMessageRate"`

	// The backoffMaxMessageRate parameter provides the means to limit the rate at which messages are delivered to a destination domain in backoff state.
	BackoffMaxMessageRate *string `mandatory:"false" json:"backoffMaxMessageRate"`

	// Retry frequencies for message delivery retries when messages aren't successfully delivered the
	// first time. This option can be specified as a series of intervals as arguments. The first interval specifies the time
	// to wait before the first retry, the second specifies the time to wait for the second retry, and so on.
	// The last value given specifies the time to wait for all subsequent retries. Deliveries are attempted for a
	// period of time specified by the retryPeriod option
	RetryFrequenciesInMinutes []int `mandatory:"false" json:"retryFrequenciesInMinutes"`

	// Retry frequencies for message that have been placed in IP backoff mode. This option can be specified as a series of intervals as arguments. The first interval specifies the time
	// to wait before the first retry, the second specifies the time to wait for the second retry, and so on.
	// The last value given specifies the time to wait for all subsequent retries. Deliveries are attempted for a
	// period of time specified by the retryPeriod option
	BackoffRetryFrequenciesInMinutes []int `mandatory:"false" json:"backoffRetryFrequenciesInMinutes"`

	// Timeout, in seconds, for the backoff entry.
	BackoffRetryFrequencyTimeoutInSeconds *int `mandatory:"false" json:"backoffRetryFrequencyTimeoutInSeconds"`

	// TLS parameter controls whether TLS negotiation is attempted, and if it fails whether fallback should be performed or the email should be bounced.
	// Possible values are:
	//   "disable"- disables the use of TLS,
	//   "optional", which allows the use of tLS if the channel is set to allows it, and
	//   "require", which requires the use of TLS
	Tls EmailDeliveryConfigOptionsTlsEnum `mandatory:"false" json:"tls,omitempty"`
}

func (m EmailDeliveryConfigOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailDeliveryConfigOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEmailDeliveryConfigOptionsTlsEnum(string(m.Tls)); !ok && m.Tls != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Tls: %s. Supported values are: %s.", m.Tls, strings.Join(GetEmailDeliveryConfigOptionsTlsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailDeliveryConfigOptionsTlsEnum Enum with underlying type: string
type EmailDeliveryConfigOptionsTlsEnum string

// Set of constants representing the allowable values for EmailDeliveryConfigOptionsTlsEnum
const (
	EmailDeliveryConfigOptionsTlsDisable  EmailDeliveryConfigOptionsTlsEnum = "DISABLE"
	EmailDeliveryConfigOptionsTlsOptional EmailDeliveryConfigOptionsTlsEnum = "OPTIONAL"
	EmailDeliveryConfigOptionsTlsRequire  EmailDeliveryConfigOptionsTlsEnum = "REQUIRE"
)

var mappingEmailDeliveryConfigOptionsTlsEnum = map[string]EmailDeliveryConfigOptionsTlsEnum{
	"DISABLE":  EmailDeliveryConfigOptionsTlsDisable,
	"OPTIONAL": EmailDeliveryConfigOptionsTlsOptional,
	"REQUIRE":  EmailDeliveryConfigOptionsTlsRequire,
}

var mappingEmailDeliveryConfigOptionsTlsEnumLowerCase = map[string]EmailDeliveryConfigOptionsTlsEnum{
	"disable":  EmailDeliveryConfigOptionsTlsDisable,
	"optional": EmailDeliveryConfigOptionsTlsOptional,
	"require":  EmailDeliveryConfigOptionsTlsRequire,
}

// GetEmailDeliveryConfigOptionsTlsEnumValues Enumerates the set of values for EmailDeliveryConfigOptionsTlsEnum
func GetEmailDeliveryConfigOptionsTlsEnumValues() []EmailDeliveryConfigOptionsTlsEnum {
	values := make([]EmailDeliveryConfigOptionsTlsEnum, 0)
	for _, v := range mappingEmailDeliveryConfigOptionsTlsEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailDeliveryConfigOptionsTlsEnumStringValues Enumerates the set of values in String for EmailDeliveryConfigOptionsTlsEnum
func GetEmailDeliveryConfigOptionsTlsEnumStringValues() []string {
	return []string{
		"DISABLE",
		"OPTIONAL",
		"REQUIRE",
	}
}

// GetMappingEmailDeliveryConfigOptionsTlsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailDeliveryConfigOptionsTlsEnum(val string) (EmailDeliveryConfigOptionsTlsEnum, bool) {
	enum, ok := mappingEmailDeliveryConfigOptionsTlsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
