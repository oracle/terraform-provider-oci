// Copyright (c) 2017, 2025 Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// these are GLOBAL variables used by JMS integration tests.
	// ensure that you have set up static resources in target JMS environment (e.g. DEV, DEV2, DEV3)
	// ensure that you also define environment variables for these static resources in your tfworkflow terminal
	JmsTenancyId = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	JmsUserId    = utils.GetEnvSettingWithBlankDefault("user_ocid")

	JmsCreateCompartmentId = utils.GetEnvSettingWithBlankDefault("compartment_id_for_create")
	JmsUpdateCompartmentId = utils.GetEnvSettingWithBlankDefault("compartment_id_for_update")

	JmsCompartmentId     = utils.GetEnvSettingWithBlankDefault("compartment_id_for_static_resources")
	JmsFleetId           = utils.GetEnvSettingWithBlankDefault("fleet_ocid")
	JmsManagedInstanceId = utils.GetEnvSettingWithBlankDefault("managed_instance_ocid")

	JmsLogGroupId       = utils.GetEnvSettingWithBlankDefault("fleet_log_group_ocid")
	JmsInventoryLogId   = utils.GetEnvSettingWithBlankDefault("fleet_inventory_log_ocid")
	JmsOperationLogId   = utils.GetEnvSettingWithBlankDefault("fleet_operation_log_ocid")
	JmsCryptoEventLogId = utils.GetEnvSettingWithBlankDefault("crypto_event_log_ocid")

	JmsAnalyticBucketId        = utils.GetEnvSettingWithBlankDefault("analytic_bucket_ocid")
	JmsAnalyticBucketNamespace = utils.GetEnvSettingWithBlankDefault("analytic_bucket_namespace")
	JmsAnalyticBucketName      = utils.GetEnvSettingWithBlankDefault("analytic_bucket_name")

	JmsUtilsBucketName      = utils.GetEnvSettingWithBlankDefault("jms_utils_bucket_name")
	JmsUtilsBucketNamespace = utils.GetEnvSettingWithBlankDefault("jms_utils_bucket_namespace")

	JmsUtilsJavaMigrationReportId     = utils.GetEnvSettingWithBlankDefault("jms_utils_java_migration_report_ocid")
	JmsUtilsPerformanceTuningReportId = utils.GetEnvSettingWithBlankDefault("jms_utils_performance_tuning_report_ocid")
)
