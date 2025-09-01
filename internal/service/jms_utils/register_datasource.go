// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_utils

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_jms_utils_analyze_applications_configuration", JmsUtilsAnalyzeApplicationsConfigurationDataSource())
	tfresource.RegisterDatasource("oci_jms_utils_java_migration_analysi", JmsUtilsJavaMigrationAnalysiDataSource())
	tfresource.RegisterDatasource("oci_jms_utils_java_migration_analysis", JmsUtilsJavaMigrationAnalysisDataSource())
	tfresource.RegisterDatasource("oci_jms_utils_performance_tuning_analysi", JmsUtilsPerformanceTuningAnalysiDataSource())
	tfresource.RegisterDatasource("oci_jms_utils_performance_tuning_analysis", JmsUtilsPerformanceTuningAnalysisDataSource())
	tfresource.RegisterDatasource("oci_jms_utils_subscription_acknowledgment_configuration", JmsUtilsSubscriptionAcknowledgmentConfigurationDataSource())
}
