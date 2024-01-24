// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_log_analytics_log_analytics_entity", LogAnalyticsLogAnalyticsEntityResource())
	tfresource.RegisterResource("oci_log_analytics_log_analytics_import_custom_content", LogAnalyticsLogAnalyticsImportCustomContentResource())
	tfresource.RegisterResource("oci_log_analytics_log_analytics_log_group", LogAnalyticsLogAnalyticsLogGroupResource())
	tfresource.RegisterResource("oci_log_analytics_log_analytics_object_collection_rule", LogAnalyticsLogAnalyticsObjectCollectionRuleResource())
	tfresource.RegisterResource("oci_log_analytics_log_analytics_preferences_management", LogAnalyticsLogAnalyticsPreferencesManagementResource())
	tfresource.RegisterResource("oci_log_analytics_log_analytics_resource_categories_management", LogAnalyticsLogAnalyticsResourceCategoriesManagementResource())
	tfresource.RegisterResource("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResource())
	tfresource.RegisterResource("oci_log_analytics_namespace_ingest_time_rule", LogAnalyticsNamespaceIngestTimeRuleResource())
	tfresource.RegisterResource("oci_log_analytics_namespace_ingest_time_rules_management", LogAnalyticsNamespaceIngestTimeRulesManagementResource())
	tfresource.RegisterResource("oci_log_analytics_namespace_scheduled_task", LogAnalyticsNamespaceScheduledTaskResource())
	tfresource.RegisterResource("oci_log_analytics_namespace", LogAnalyticsNamespaceResource())
}
