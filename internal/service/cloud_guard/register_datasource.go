// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_cloud_guard_cloud_guard_configuration", CloudGuardCloudGuardConfigurationDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_data_mask_rule", CloudGuardDataMaskRuleDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_data_mask_rules", CloudGuardDataMaskRulesDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_data_source", CloudGuardDataSourceDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_data_source_event", CloudGuardDataSourceEventDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_data_source_events", CloudGuardDataSourceEventsDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_data_sources", CloudGuardDataSourcesDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_detector_recipe", CloudGuardDetectorRecipeDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_detector_recipes", CloudGuardDetectorRecipesDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_managed_list", CloudGuardManagedListDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_managed_lists", CloudGuardManagedListsDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_problem_entities", CloudGuardProblemEntitiesDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_problem_entity", CloudGuardProblemEntityDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_responder_recipe", CloudGuardResponderRecipeDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_responder_recipes", CloudGuardResponderRecipesDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_security_policies", CloudGuardSecurityPoliciesDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_security_policy", CloudGuardSecurityPolicyDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_security_recipe", CloudGuardSecurityRecipeDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_security_recipes", CloudGuardSecurityRecipesDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_security_zone", CloudGuardSecurityZoneDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_security_zones", CloudGuardSecurityZonesDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_target", CloudGuardTargetDataSource())
	tfresource.RegisterDatasource("oci_cloud_guard_targets", CloudGuardTargetsDataSource())
}
