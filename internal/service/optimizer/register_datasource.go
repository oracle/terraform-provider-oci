// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_optimizer_categories", OptimizerCategoriesDataSource())
	tfresource.RegisterDatasource("oci_optimizer_category", OptimizerCategoryDataSource())
	tfresource.RegisterDatasource("oci_optimizer_enrollment_status", OptimizerEnrollmentStatusDataSource())
	tfresource.RegisterDatasource("oci_optimizer_enrollment_statuses", OptimizerEnrollmentStatusesDataSource())
	tfresource.RegisterDatasource("oci_optimizer_histories", OptimizerHistoriesDataSource())
	tfresource.RegisterDatasource("oci_optimizer_profile", OptimizerProfileDataSource())
	tfresource.RegisterDatasource("oci_optimizer_profile_level", OptimizerProfileLevelDataSource())
	tfresource.RegisterDatasource("oci_optimizer_profile_levels", OptimizerProfileLevelsDataSource())
	tfresource.RegisterDatasource("oci_optimizer_profiles", OptimizerProfilesDataSource())
	tfresource.RegisterDatasource("oci_optimizer_recommendation", OptimizerRecommendationDataSource())
	tfresource.RegisterDatasource("oci_optimizer_recommendation_strategies", OptimizerRecommendationStrategiesDataSource())
	tfresource.RegisterDatasource("oci_optimizer_recommendation_strategy", OptimizerRecommendationStrategyDataSource())
	tfresource.RegisterDatasource("oci_optimizer_recommendations", OptimizerRecommendationsDataSource())
	tfresource.RegisterDatasource("oci_optimizer_resource_action", OptimizerResourceActionDataSource())
	tfresource.RegisterDatasource("oci_optimizer_resource_actions", OptimizerResourceActionsDataSource())
}
