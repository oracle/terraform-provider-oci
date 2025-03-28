// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDataSafeAlertPolicyRuleHints.GetIdFn = getDataSafeAlertPolicyRuleId
	exportDataSafeMaskingPoliciesMaskingColumnHints.GetIdFn = getDataSafeMaskingPoliciesMaskingColumnId
	exportDataSafeSensitiveDataModelsSensitiveColumnHints.GetIdFn = getDataSafeSensitiveDataModelsSensitiveColumnId

	exportDataSafeTargetDatabasePeerTargetDatabaseHints.GetIdFn = getDataSafeTargetDatabasePeerTargetDatabaseId
	exportDataSafeSensitiveDataModelReferentialRelationHints.GetIdFn = getDataSafeSensitiveDataModelReferentialRelationId
	exportDataSafeSensitiveTypeGroupGroupedSensitiveTypeHints.GetIdFn = getDataSafeSensitiveTypeGroupGroupedSensitiveTypeId
	exportDataSafeDiscoveryJobsResultHints.GetIdFn = getDataSafeDiscoveryJobsResultId
	exportDataSafeAlertPolicyHints.FindResourcesOverrideFn = findAlertPolicies
	tf_export.RegisterCompartmentGraphs("data_safe", dataSafeResourceGraph)
}

func getDataSafeTargetDatabasePeerTargetDatabaseId(resource *tf_export.OCIResource) (string, error) {

	peerTargetDatabaseId, ok := resource.SourceAttributes["peer_target_database_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find peerTargetDatabaseId for DataSafe TargetDatabasePeerTargetDatabase")
	}
	targetDatabaseId := resource.Parent.Id
	return GetTargetDatabasePeerTargetDatabaseCompositeId(peerTargetDatabaseId, targetDatabaseId), nil
}

// Custom overrides for generating composite IDs within the resource discovery framework
func getDataSafeDiscoveryJobsResultId(resource *tf_export.OCIResource) (string, error) {

	discoveryJobId := resource.Parent.Id
	resultKey, ok := resource.SourceAttributes["result_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find resultKey for DataSafe DiscoveryJobsResult")
	}
	return GetDiscoveryJobsResultCompositeId(discoveryJobId, resultKey), nil
}

func getDataSafeAlertPolicyRuleId(resource *tf_export.OCIResource) (string, error) {

	alertPolicyId := resource.Parent.Id
	ruleKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find ruleKey for DataSafe AlertPolicyRule")
	}
	return GetAlertPolicyRuleCompositeId(alertPolicyId, ruleKey), nil
}

func getDataSafeMaskingPoliciesMaskingColumnId(resource *tf_export.OCIResource) (string, error) {

	maskingColumnKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find maskingColumnKey for DataSafe MaskingPoliciesMaskingColumn")
	}
	maskingPolicyId := resource.Parent.Id
	return GetMaskingPoliciesMaskingColumnCompositeId(maskingColumnKey, maskingPolicyId), nil
}

func getDataSafeSensitiveDataModelsSensitiveColumnId(resource *tf_export.OCIResource) (string, error) {

	sensitiveColumnKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find sensitiveColumnKey for DataSafe SensitiveDataModelsSensitiveColumn")
	}
	sensitiveDataModelId := resource.Parent.Id
	return GetSensitiveDataModelsSensitiveColumnCompositeId(sensitiveColumnKey, sensitiveDataModelId), nil
}

func getDataSafeSensitiveDataModelReferentialRelationId(resource *tf_export.OCIResource) (string, error) {

	referentialRelationKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find referentialRelationKey for DataSafe SensitiveDataModelReferentialRelation")
	}
	sensitiveDataModelId := resource.Parent.Id
	return GetSensitiveDataModelReferentialRelationCompositeId(referentialRelationKey, sensitiveDataModelId), nil
}

func getDataSafeSensitiveTypeGroupGroupedSensitiveTypeId(resource *tf_export.OCIResource) (string, error) {

	sensitiveTypeGroupId := resource.Parent.Id
	return GetSensitiveTypeGroupGroupedSensitiveTypeCompositeId(sensitiveTypeGroupId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDataSafeDataSafePrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_data_safe_private_endpoint",
	DatasourceClass:        "oci_data_safe_data_safe_private_endpoints",
	DatasourceItemsAttr:    "data_safe_private_endpoints",
	ResourceAbbreviation:   "data_safe_private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.ListDataSafePrivateEndpointsLifecycleStateActive),
	},
}

var exportDataSafeOnPremConnectorHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_data_safe_on_prem_connector",
	DatasourceClass:      "oci_data_safe_on_prem_connectors",
	DatasourceItemsAttr:  "on_prem_connectors",
	ResourceAbbreviation: "on_prem_connector",
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.OnPremConnectorLifecycleStateInactive),
		string(oci_data_safe.OnPremConnectorLifecycleStateActive),
	},
}

var exportDataSafeTargetDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_target_database",
	DatasourceClass:        "oci_data_safe_target_databases",
	DatasourceItemsAttr:    "target_databases",
	ResourceAbbreviation:   "target_database",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateActive),
		string(oci_data_safe.TargetDatabaseLifecycleStateNeedsAttention),
	},
}

var exportDataSafeSecurityAssessmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_security_assessment",
	DatasourceClass:        "oci_data_safe_security_assessments",
	DatasourceItemsAttr:    "security_assessments",
	ResourceAbbreviation:   "security_assessment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.SecurityAssessmentLifecycleStateSucceeded),
	},
}

var exportDataSafeUserAssessmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_data_safe_user_assessment",
	DatasourceClass:      "oci_data_safe_user_assessments",
	DatasourceItemsAttr:  "user_assessments",
	ResourceAbbreviation: "user_assessment",
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.UserAssessmentLifecycleStateSucceeded),
	},
}

var exportDataSafeUnsetSecurityAssessmentBaselineHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_data_safe_unset_security_assessment_baseline",
	ResourceAbbreviation: "unset_security_assessment_baseline",
}

var exportDataSafeReportDefinitionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_report_definition",
	DatasourceClass:        "oci_data_safe_report_definitions",
	DatasourceItemsAttr:    "report_definition_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "report_definition",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.ReportDefinitionLifecycleStateActive),
	},
}

var exportDataSafeAuditTrailHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_audit_trail",
	DatasourceClass:        "oci_data_safe_audit_trails",
	DatasourceItemsAttr:    "audit_trail_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "audit_trail",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.AuditTrailLifecycleStateActive),
		string(oci_data_safe.AuditTrailLifecycleStateNeedsAttention),
	},
}

var exportDataSafeAlertHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_alert",
	DatasourceClass:        "oci_data_safe_alerts",
	DatasourceItemsAttr:    "alert_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "alert",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.AlertLifecycleStateSucceeded),
	},
}

var exportDataSafeAuditArchiveRetrievalHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_audit_archive_retrieval",
	DatasourceClass:        "oci_data_safe_audit_archive_retrievals",
	DatasourceItemsAttr:    "audit_archive_retrieval_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "audit_archive_retrieval",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.AuditArchiveRetrievalLifecycleStateActive),
		string(oci_data_safe.AuditArchiveRetrievalLifecycleStateNeedsAttention),
	},
}

var exportDataSafeAuditProfileHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_audit_profile",
	DatasourceClass:        "oci_data_safe_audit_profiles",
	DatasourceItemsAttr:    "audit_profile_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "audit_profile",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.AuditProfileLifecycleStateActive),
		string(oci_data_safe.AuditProfileLifecycleStateNeedsAttention),
	},
}

var exportDataSafeAlertPolicyRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_alert_policy_rule",
	DatasourceClass:        "oci_data_safe_alert_policy_rules",
	DatasourceItemsAttr:    "alert_policy_rule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "alert_policy_rule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.AlertPolicyRuleLifecycleStateActive),
	},
}

var exportDataSafeAlertPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_alert_policy",
	DatasourceClass:        "oci_data_safe_alert_policies",
	DatasourceItemsAttr:    "alert_policy_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "alert_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.AlertPolicyLifecycleStateActive),
	},
}

var exportDataSafeAuditPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_audit_policy",
	DatasourceClass:        "oci_data_safe_audit_policies",
	DatasourceItemsAttr:    "audit_policy_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "audit_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.AuditPolicyLifecycleStateActive),
		string(oci_data_safe.AuditPolicyLifecycleStateNeedsAttention),
	},
}

var exportDataSafeTargetAlertPolicyAssociationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_target_alert_policy_association",
	DatasourceClass:        "oci_data_safe_target_alert_policy_associations",
	DatasourceItemsAttr:    "target_alert_policy_association_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "target_alert_policy_association",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.AlertPolicyLifecycleStateActive),
	},
}

var exportDataSafeReportHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_report",
	DatasourceClass:        "oci_data_safe_reports",
	DatasourceItemsAttr:    "report_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "report",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.ReportLifecycleStateActive),
	},
}

var exportDataSafeSensitiveTypeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_sensitive_type",
	DatasourceClass:        "oci_data_safe_sensitive_types",
	DatasourceItemsAttr:    "sensitive_type_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sensitive_type",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.DiscoveryLifecycleStateActive),
	},
}

var exportDataSafeMaskingPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_masking_policy",
	DatasourceClass:        "oci_data_safe_masking_policies",
	DatasourceItemsAttr:    "masking_policy_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "masking_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.MaskingLifecycleStateActive),
		string(oci_data_safe.MaskingLifecycleStateNeedsAttention),
	},
}

var exportDataSafeMaskingPoliciesMaskingColumnHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_masking_policies_masking_column",
	DatasourceClass:        "oci_data_safe_masking_policies_masking_columns",
	DatasourceItemsAttr:    "masking_column_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "masking_policies_masking_column",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.MaskingColumnLifecycleStateActive),
		string(oci_data_safe.MaskingColumnLifecycleStateNeedsAttention),
	},
}

var exportDataSafeLibraryMaskingFormatHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_library_masking_format",
	DatasourceClass:        "oci_data_safe_library_masking_formats",
	DatasourceItemsAttr:    "library_masking_format_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "library_masking_format",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.MaskingLifecycleStateActive),
		string(oci_data_safe.MaskingLifecycleStateNeedsAttention),
	},
}

var exportDataSafeSensitiveDataModelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_sensitive_data_model",
	DatasourceClass:        "oci_data_safe_sensitive_data_models",
	DatasourceItemsAttr:    "sensitive_data_model_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sensitive_data_model",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.DiscoveryLifecycleStateActive),
	},
}

var exportDataSafeSensitiveDataModelsSensitiveColumnHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_sensitive_data_models_sensitive_column",
	DatasourceClass:        "oci_data_safe_sensitive_data_models_sensitive_columns",
	DatasourceItemsAttr:    "sensitive_column_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sensitive_data_models_sensitive_column",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.SensitiveColumnLifecycleStateActive),
	},
}

var exportDataSafeDiscoveryJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_discovery_job",
	DatasourceClass:        "oci_data_safe_discovery_jobs",
	DatasourceItemsAttr:    "discovery_job_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "discovery_job",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.DiscoveryLifecycleStateActive),
	},
}

var exportDataSafeSdmMaskingPolicyDifferenceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_sdm_masking_policy_difference",
	DatasourceClass:        "oci_data_safe_sdm_masking_policy_differences",
	DatasourceItemsAttr:    "sdm_masking_policy_difference_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sdm_masking_policy_difference",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateActive),
	},
}

var exportDataSafeDiscoveryJobsResultHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_discovery_jobs_result",
	DatasourceClass:        "oci_data_safe_discovery_jobs_results",
	DatasourceItemsAttr:    "discovery_job_result_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "discovery_jobs_result",
}

var exportDataSafeTargetDatabasePeerTargetDatabaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_target_database_peer_target_database",
	DatasourceClass:        "oci_data_safe_target_database_peer_target_databases",
	DatasourceItemsAttr:    "peer_target_database_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "target_database_peer_target_database",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateActive),
		string(oci_data_safe.TargetDatabaseLifecycleStateNeedsAttention),
	},
}

var exportDataSafeCalculateAuditVolumeAvailableHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_data_safe_calculate_audit_volume_available",
	ResourceAbbreviation: "calculate_audit_volume_available",
}

var exportDataSafeCalculateAuditVolumeCollectedHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_data_safe_calculate_audit_volume_collected",
	ResourceAbbreviation: "calculate_audit_volume_collected",
}

var exportDataSafeGenerateOnPremConnectorConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_data_safe_generate_on_prem_connector_configuration",
	ResourceAbbreviation: "generate_on_prem_connector_configuration",
}

var exportDataSafeSensitiveTypesExportHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_sensitive_types_export",
	DatasourceClass:        "oci_data_safe_sensitive_types_exports",
	DatasourceItemsAttr:    "sensitive_types_export_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sensitive_types_export",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.SensitiveTypesExportLifecycleStateActive),
	},
}

var exportDataSafeSensitiveDataModelReferentialRelationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_sensitive_data_model_referential_relation",
	DatasourceClass:        "oci_data_safe_sensitive_data_model_referential_relations",
	DatasourceItemsAttr:    "referential_relation_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sensitive_data_model_referential_relation",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.ReferentialRelationLifecycleStateActive),
	},
}

var exportDataSafeSensitiveTypeGroupGroupedSensitiveTypeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_sensitive_type_group_grouped_sensitive_type",
	DatasourceClass:        "oci_data_safe_sensitive_type_group_grouped_sensitive_types",
	DatasourceItemsAttr:    "grouped_sensitive_type_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sensitive_type_group_grouped_sensitive_type",
}

var exportDataSafeSensitiveTypeGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_safe_sensitive_type_group",
	DatasourceClass:        "oci_data_safe_sensitive_type_groups",
	DatasourceItemsAttr:    "sensitive_type_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sensitive_type_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_safe.SensitiveTypeGroupLifecycleStateActive),
	},
}

var dataSafeResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataSafeDataSafePrivateEndpointHints},
		{TerraformResourceHints: exportDataSafeOnPremConnectorHints},
		{TerraformResourceHints: exportDataSafeTargetDatabaseHints},
		{TerraformResourceHints: exportDataSafeSecurityAssessmentHints},
		{TerraformResourceHints: exportDataSafeUserAssessmentHints},
		{TerraformResourceHints: exportDataSafeReportDefinitionHints},
		{TerraformResourceHints: exportDataSafeAuditTrailHints},
		{TerraformResourceHints: exportDataSafeAlertHints},
		{TerraformResourceHints: exportDataSafeAuditArchiveRetrievalHints},
		{TerraformResourceHints: exportDataSafeAuditProfileHints},
		{TerraformResourceHints: exportDataSafeAlertPolicyHints},
		{TerraformResourceHints: exportDataSafeAuditPolicyHints},
		{TerraformResourceHints: exportDataSafeTargetAlertPolicyAssociationHints},
		{TerraformResourceHints: exportDataSafeReportHints},
		{TerraformResourceHints: exportDataSafeSensitiveTypeHints},
		{TerraformResourceHints: exportDataSafeMaskingPolicyHints},
		{TerraformResourceHints: exportDataSafeLibraryMaskingFormatHints},
		{TerraformResourceHints: exportDataSafeSensitiveDataModelHints},
		{TerraformResourceHints: exportDataSafeDiscoveryJobHints},
		{TerraformResourceHints: exportDataSafeSdmMaskingPolicyDifferenceHints},
		{TerraformResourceHints: exportDataSafeSensitiveTypesExportHints},
		{TerraformResourceHints: exportDataSafeSensitiveTypeGroupHints},
	},
	"oci_data_safe_alert_policy": {
		{
			TerraformResourceHints: exportDataSafeAlertPolicyRuleHints,
			DatasourceQueryParams: map[string]string{
				"alert_policy_id": "id",
			},
		},
	},
	"oci_data_safe_target_database": {
		{
			TerraformResourceHints: exportDataSafeTargetDatabasePeerTargetDatabaseHints,
			DatasourceQueryParams: map[string]string{
				"target_database_id": "id",
			},
		},
	},
	"oci_data_safe_masking_policy": {
		{
			TerraformResourceHints: exportDataSafeMaskingPoliciesMaskingColumnHints,
			DatasourceQueryParams: map[string]string{
				"masking_policy_id": "id",
			},
		},
	},
	"oci_data_safe_sensitive_data_model": {
		{
			TerraformResourceHints: exportDataSafeSensitiveDataModelReferentialRelationHints,
			DatasourceQueryParams: map[string]string{
				"sensitive_data_model_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDataSafeSensitiveDataModelsSensitiveColumnHints,
			DatasourceQueryParams: map[string]string{
				"sensitive_data_model_id": "id",
			},
		},
	},
	"oci_data_safe_sensitive_type_group": {
		{
			TerraformResourceHints: exportDataSafeSensitiveTypeGroupGroupedSensitiveTypeHints,
			DatasourceQueryParams: map[string]string{
				"sensitive_type_group_id": "id",
			},
		},
	},
	"oci_data_safe_discovery_job": {
		{
			TerraformResourceHints: exportDataSafeDiscoveryJobsResultHints,
			DatasourceQueryParams: map[string]string{
				"discovery_job_id": "id",
			},
		},
	},
}

func findAlertPolicies(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) (resources []*tf_export.OCIResource, err error) {
	results := []*tf_export.OCIResource{}
	request := oci_data_safe.ListAlertPoliciesRequest{}
	tmp := true
	request.IsUserDefined = &tmp
	request.CompartmentId = &parent.CompartmentId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := ctx.Clients.DataSafeClient().ListAlertPolicies(context.Background(), request)
	if err != nil {
		return nil, err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := ctx.Clients.DataSafeClient().ListAlertPolicies(context.Background(), request)
		if err != nil {
			return nil, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, alertPolicy := range response.Items {
		alertPolicyResource := tf_export.ResourcesMap[tfMeta.ResourceClass]

		d := alertPolicyResource.TestResourceData()
		d.SetId(*alertPolicy.Id)

		if err := alertPolicyResource.Read(d, ctx.Clients); err != nil {
			rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
			ctx.AddErrorToList(rdError)
			continue
		}

		state := d.Get("state")
		if state != nil && len(tfMeta.DiscoverableLifecycleStates) > 0 {
			discoverable := false
			for _, val := range tfMeta.DiscoverableLifecycleStates {
				if strings.EqualFold(state.(string), val) {
					discoverable = true
					break
				}
			}
			if !discoverable {
				continue
			}
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    parent.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(alertPolicyResource.Schema, d),
			RawResource:      alertPolicy,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		if resource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(resource.SourceAttributes, alertPolicyResource.Schema); err != nil {
			resource.TerraformName = fmt.Sprintf("%s_%s", parent.Parent.TerraformName, *alertPolicy.DisplayName)
			resource.TerraformName = tf_export.CheckDuplicateResourceName(resource.TerraformName)
		}

		results = append(results, resource)
	}

	return results, nil
}
