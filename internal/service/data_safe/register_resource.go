// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_data_safe_alert", DataSafeAlertResource())
	tfresource.RegisterResource("oci_data_safe_audit_archive_retrieval", DataSafeAuditArchiveRetrievalResource())
	tfresource.RegisterResource("oci_data_safe_audit_policy", DataSafeAuditPolicyResource())
	tfresource.RegisterResource("oci_data_safe_audit_profile", DataSafeAuditProfileResource())
	tfresource.RegisterResource("oci_data_safe_audit_trail", DataSafeAuditTrailResource())
	tfresource.RegisterResource("oci_data_safe_compare_security_assessment", DataSafeCompareSecurityAssessmentResource())
	tfresource.RegisterResource("oci_data_safe_compare_user_assessment", DataSafeCompareUserAssessmentResource())
	tfresource.RegisterResource("oci_data_safe_data_safe_configuration", DataSafeDataSafeConfigurationResource())
	tfresource.RegisterResource("oci_data_safe_data_safe_private_endpoint", DataSafeDataSafePrivateEndpointResource())
	tfresource.RegisterResource("oci_data_safe_discovery_job", DataSafeDiscoveryJobResource())
	tfresource.RegisterResource("oci_data_safe_library_masking_format", DataSafeLibraryMaskingFormatResource())
	tfresource.RegisterResource("oci_data_safe_masking_policies_masking_column", DataSafeMaskingPoliciesMaskingColumnResource())
	tfresource.RegisterResource("oci_data_safe_masking_policy", DataSafeMaskingPolicyResource())
	tfresource.RegisterResource("oci_data_safe_on_prem_connector", DataSafeOnPremConnectorResource())
	tfresource.RegisterResource("oci_data_safe_report_definition", DataSafeReportDefinitionResource())
	tfresource.RegisterResource("oci_data_safe_security_assessment", DataSafeSecurityAssessmentResource())
	tfresource.RegisterResource("oci_data_safe_sensitive_data_model", DataSafeSensitiveDataModelResource())
	tfresource.RegisterResource("oci_data_safe_sensitive_data_models_sensitive_column", DataSafeSensitiveDataModelsSensitiveColumnResource())
	tfresource.RegisterResource("oci_data_safe_sensitive_type", DataSafeSensitiveTypeResource())
	tfresource.RegisterResource("oci_data_safe_set_security_assessment_baseline", DataSafeSetSecurityAssessmentBaselineResource())
	tfresource.RegisterResource("oci_data_safe_set_user_assessment_baseline", DataSafeSetUserAssessmentBaselineResource())
	tfresource.RegisterResource("oci_data_safe_target_alert_policy_association", DataSafeTargetAlertPolicyAssociationResource())
	tfresource.RegisterResource("oci_data_safe_target_database", DataSafeTargetDatabaseResource())
	tfresource.RegisterResource("oci_data_safe_unset_security_assessment_baseline", DataSafeUnsetSecurityAssessmentBaselineResource())
	tfresource.RegisterResource("oci_data_safe_unset_user_assessment_baseline", DataSafeUnsetUserAssessmentBaselineResource())
	tfresource.RegisterResource("oci_data_safe_user_assessment", DataSafeUserAssessmentResource())
	tfresource.RegisterResource("oci_data_safe_mask_data", DataSafeMaskDataResource())
	tfresource.RegisterResource("oci_data_safe_add_sdm_columns", DataSafeAddColumnsFromSdmResource())
	tfresource.RegisterResource("oci_data_safe_sensitive_data_models_apply_discovery_job_results", DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResource())
	tfresource.RegisterResource("oci_data_safe_discovery_jobs_result", DataSafeDiscoveryJobsResultResource())
}
