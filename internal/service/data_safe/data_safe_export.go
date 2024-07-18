package data_safe

import (
	"fmt"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDataSafeMaskingPoliciesMaskingColumnHints.GetIdFn = getDataSafeMaskingPoliciesMaskingColumnId
	exportDataSafeSensitiveDataModelsSensitiveColumnHints.GetIdFn = getDataSafeSensitiveDataModelsSensitiveColumnId
	exportDataSafeTargetDatabasePeerTargetDatabaseHints.GetIdFn = getDataSafeTargetDatabasePeerTargetDatabaseId
	exportDataSafeDiscoveryJobsResultHints.GetIdFn = getDataSafeDiscoveryJobsResultId
	tf_export.RegisterCompartmentGraphs("data_safe", dataSafeResourceGraph)
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

func getDataSafeTargetDatabasePeerTargetDatabaseId(resource *tf_export.OCIResource) (string, error) {

	peerTargetDatabaseId, ok := resource.SourceAttributes["peer_target_database_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find peerTargetDatabaseId for DataSafe TargetDatabasePeerTargetDatabase")
	}
	targetDatabaseId := resource.Parent.Id
	return GetTargetDatabasePeerTargetDatabaseCompositeId(peerTargetDatabaseId, targetDatabaseId), nil
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
		{TerraformResourceHints: exportDataSafeAuditPolicyHints},
		{TerraformResourceHints: exportDataSafeTargetAlertPolicyAssociationHints},
		{TerraformResourceHints: exportDataSafeReportHints},
		{TerraformResourceHints: exportDataSafeSensitiveTypeHints},
		{TerraformResourceHints: exportDataSafeMaskingPolicyHints},
		{TerraformResourceHints: exportDataSafeLibraryMaskingFormatHints},
		{TerraformResourceHints: exportDataSafeSensitiveDataModelHints},
		{TerraformResourceHints: exportDataSafeDiscoveryJobHints},
		{TerraformResourceHints: exportDataSafeSdmMaskingPolicyDifferenceHints},
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
			TerraformResourceHints: exportDataSafeSensitiveDataModelsSensitiveColumnHints,
			DatasourceQueryParams: map[string]string{
				"sensitive_data_model_id": "id",
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
