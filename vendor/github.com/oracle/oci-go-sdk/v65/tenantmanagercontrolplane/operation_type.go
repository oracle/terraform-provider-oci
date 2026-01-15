// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateSenderInvitation              OperationTypeEnum = "CREATE_SENDER_INVITATION"
	OperationTypeAcceptRecipientInvitation           OperationTypeEnum = "ACCEPT_RECIPIENT_INVITATION"
	OperationTypeCancelSenderInvitation              OperationTypeEnum = "CANCEL_SENDER_INVITATION"
	OperationTypeCompleteOrderActivation             OperationTypeEnum = "COMPLETE_ORDER_ACTIVATION"
	OperationTypeActivateOrderExistingTenancy        OperationTypeEnum = "ACTIVATE_ORDER_EXISTING_TENANCY"
	OperationTypeRegisterDomain                      OperationTypeEnum = "REGISTER_DOMAIN"
	OperationTypeReleaseDomain                       OperationTypeEnum = "RELEASE_DOMAIN"
	OperationTypeCreateChildTenancy                  OperationTypeEnum = "CREATE_CHILD_TENANCY"
	OperationTypeAssignDefaultSubscription           OperationTypeEnum = "ASSIGN_DEFAULT_SUBSCRIPTION"
	OperationTypeManualLinkCreation                  OperationTypeEnum = "MANUAL_LINK_CREATION"
	OperationTypeTerminateOrganizationTenancy        OperationTypeEnum = "TERMINATE_ORGANIZATION_TENANCY"
	OperationTypeUpdateSaasCapability                OperationTypeEnum = "UPDATE_SAAS_CAPABILITY"
	OperationTypeSoftTerminateTenancy                OperationTypeEnum = "SOFT_TERMINATE_TENANCY"
	OperationTypeHardTerminateTenancy                OperationTypeEnum = "HARD_TERMINATE_TENANCY"
	OperationTypeRestoreTenancy                      OperationTypeEnum = "RESTORE_TENANCY"
	OperationTypeLogTenancyTerminationRequest        OperationTypeEnum = "LOG_TENANCY_TERMINATION_REQUEST"
	OperationTypeStandaloneTenancyTerminationRequest OperationTypeEnum = "STANDALONE_TENANCY_TERMINATION_REQUEST"
	OperationTypeSelfOptIn                           OperationTypeEnum = "SELF_OPT_IN"
	OperationTypeSelfOptOut                          OperationTypeEnum = "SELF_OPT_OUT"
	OperationTypeOrderPayloadsBackfillEtl            OperationTypeEnum = "ORDER_PAYLOADS_BACKFILL_ETL"
	OperationTypePublishToTopic                      OperationTypeEnum = "PUBLISH_TO_TOPIC"
	OperationTypeTerminateLink                       OperationTypeEnum = "TERMINATE_LINK"
	OperationTypeTransferSubscription                OperationTypeEnum = "TRANSFER_SUBSCRIPTION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_SENDER_INVITATION":               OperationTypeCreateSenderInvitation,
	"ACCEPT_RECIPIENT_INVITATION":            OperationTypeAcceptRecipientInvitation,
	"CANCEL_SENDER_INVITATION":               OperationTypeCancelSenderInvitation,
	"COMPLETE_ORDER_ACTIVATION":              OperationTypeCompleteOrderActivation,
	"ACTIVATE_ORDER_EXISTING_TENANCY":        OperationTypeActivateOrderExistingTenancy,
	"REGISTER_DOMAIN":                        OperationTypeRegisterDomain,
	"RELEASE_DOMAIN":                         OperationTypeReleaseDomain,
	"CREATE_CHILD_TENANCY":                   OperationTypeCreateChildTenancy,
	"ASSIGN_DEFAULT_SUBSCRIPTION":            OperationTypeAssignDefaultSubscription,
	"MANUAL_LINK_CREATION":                   OperationTypeManualLinkCreation,
	"TERMINATE_ORGANIZATION_TENANCY":         OperationTypeTerminateOrganizationTenancy,
	"UPDATE_SAAS_CAPABILITY":                 OperationTypeUpdateSaasCapability,
	"SOFT_TERMINATE_TENANCY":                 OperationTypeSoftTerminateTenancy,
	"HARD_TERMINATE_TENANCY":                 OperationTypeHardTerminateTenancy,
	"RESTORE_TENANCY":                        OperationTypeRestoreTenancy,
	"LOG_TENANCY_TERMINATION_REQUEST":        OperationTypeLogTenancyTerminationRequest,
	"STANDALONE_TENANCY_TERMINATION_REQUEST": OperationTypeStandaloneTenancyTerminationRequest,
	"SELF_OPT_IN":                            OperationTypeSelfOptIn,
	"SELF_OPT_OUT":                           OperationTypeSelfOptOut,
	"ORDER_PAYLOADS_BACKFILL_ETL":            OperationTypeOrderPayloadsBackfillEtl,
	"PUBLISH_TO_TOPIC":                       OperationTypePublishToTopic,
	"TERMINATE_LINK":                         OperationTypeTerminateLink,
	"TRANSFER_SUBSCRIPTION":                  OperationTypeTransferSubscription,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_sender_invitation":               OperationTypeCreateSenderInvitation,
	"accept_recipient_invitation":            OperationTypeAcceptRecipientInvitation,
	"cancel_sender_invitation":               OperationTypeCancelSenderInvitation,
	"complete_order_activation":              OperationTypeCompleteOrderActivation,
	"activate_order_existing_tenancy":        OperationTypeActivateOrderExistingTenancy,
	"register_domain":                        OperationTypeRegisterDomain,
	"release_domain":                         OperationTypeReleaseDomain,
	"create_child_tenancy":                   OperationTypeCreateChildTenancy,
	"assign_default_subscription":            OperationTypeAssignDefaultSubscription,
	"manual_link_creation":                   OperationTypeManualLinkCreation,
	"terminate_organization_tenancy":         OperationTypeTerminateOrganizationTenancy,
	"update_saas_capability":                 OperationTypeUpdateSaasCapability,
	"soft_terminate_tenancy":                 OperationTypeSoftTerminateTenancy,
	"hard_terminate_tenancy":                 OperationTypeHardTerminateTenancy,
	"restore_tenancy":                        OperationTypeRestoreTenancy,
	"log_tenancy_termination_request":        OperationTypeLogTenancyTerminationRequest,
	"standalone_tenancy_termination_request": OperationTypeStandaloneTenancyTerminationRequest,
	"self_opt_in":                            OperationTypeSelfOptIn,
	"self_opt_out":                           OperationTypeSelfOptOut,
	"order_payloads_backfill_etl":            OperationTypeOrderPayloadsBackfillEtl,
	"publish_to_topic":                       OperationTypePublishToTopic,
	"terminate_link":                         OperationTypeTerminateLink,
	"transfer_subscription":                  OperationTypeTransferSubscription,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_SENDER_INVITATION",
		"ACCEPT_RECIPIENT_INVITATION",
		"CANCEL_SENDER_INVITATION",
		"COMPLETE_ORDER_ACTIVATION",
		"ACTIVATE_ORDER_EXISTING_TENANCY",
		"REGISTER_DOMAIN",
		"RELEASE_DOMAIN",
		"CREATE_CHILD_TENANCY",
		"ASSIGN_DEFAULT_SUBSCRIPTION",
		"MANUAL_LINK_CREATION",
		"TERMINATE_ORGANIZATION_TENANCY",
		"UPDATE_SAAS_CAPABILITY",
		"SOFT_TERMINATE_TENANCY",
		"HARD_TERMINATE_TENANCY",
		"RESTORE_TENANCY",
		"LOG_TENANCY_TERMINATION_REQUEST",
		"STANDALONE_TENANCY_TERMINATION_REQUEST",
		"SELF_OPT_IN",
		"SELF_OPT_OUT",
		"ORDER_PAYLOADS_BACKFILL_ETL",
		"PUBLISH_TO_TOPIC",
		"TERMINATE_LINK",
		"TRANSFER_SUBSCRIPTION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
