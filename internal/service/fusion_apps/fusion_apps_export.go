package fusion_apps

import (
	"fmt"

	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportFusionAppsFusionEnvironmentRefreshActivityHints.GetIdFn = getFusionAppsFusionEnvironmentRefreshActivityId
	exportFusionAppsFusionEnvironmentAdminUserHints.GetIdFn = getFusionAppsFusionEnvironmentAdminUserId
	exportFusionAppsFusionEnvironmentDataMaskingActivityHints.GetIdFn = getFusionAppsFusionEnvironmentDataMaskingActivityId
	tf_export.RegisterCompartmentGraphs("fusion_apps", fusionAppsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getFusionAppsFusionEnvironmentRefreshActivityId(resource *tf_export.OCIResource) (string, error) {

	fusionEnvironmentId := resource.Parent.Id
	refreshActivityId, ok := resource.SourceAttributes["refresh_activity_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find refreshActivityId for FusionApps FusionEnvironmentRefreshActivity")
	}
	return GetFusionEnvironmentRefreshActivityCompositeId(fusionEnvironmentId, refreshActivityId), nil
}

func getFusionAppsFusionEnvironmentAdminUserId(resource *tf_export.OCIResource) (string, error) {

	adminUsername, ok := resource.SourceAttributes["admin_username"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find adminUsername for FusionApps FusionEnvironmentAdminUser")
	}
	fusionEnvironmentId := resource.Parent.Id
	return GetFusionEnvironmentAdminUserCompositeId(adminUsername, fusionEnvironmentId), nil
}

func getFusionAppsFusionEnvironmentDataMaskingActivityId(resource *tf_export.OCIResource) (string, error) {

	dataMaskingActivityId, ok := resource.SourceAttributes["data_masking_activity_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataMaskingActivityId for FusionApps FusionEnvironmentDataMaskingActivity")
	}
	fusionEnvironmentId := resource.Parent.Id
	return GetFusionEnvironmentDataMaskingActivityCompositeId(dataMaskingActivityId, fusionEnvironmentId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportFusionAppsFusionEnvironmentRefreshActivityHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fusion_apps_fusion_environment_refresh_activity",
	DatasourceClass:        "oci_fusion_apps_fusion_environment_refresh_activities",
	DatasourceItemsAttr:    "refresh_activity_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fusion_environment_refresh_activity",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fusion_apps.RefreshActivityLifecycleStateAccepted),
		string(oci_fusion_apps.RefreshActivityLifecycleStateInProgress),
		string(oci_fusion_apps.RefreshActivityLifecycleStateNeedsAttention),
		string(oci_fusion_apps.RefreshActivityLifecycleStateSucceeded),
	},
}

var exportFusionAppsFusionEnvironmentAdminUserHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fusion_apps_fusion_environment_admin_user",
	DatasourceClass:        "oci_fusion_apps_fusion_environment_admin_users",
	DatasourceItemsAttr:    "admin_user_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fusion_environment_admin_user",
}

var exportFusionAppsFusionEnvironmentFamilyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fusion_apps_fusion_environment_family",
	DatasourceClass:        "oci_fusion_apps_fusion_environment_families",
	DatasourceItemsAttr:    "fusion_environment_family_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fusion_environment_family",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateActive),
	},
}

var exportFusionAppsFusionEnvironmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fusion_apps_fusion_environment",
	DatasourceClass:        "oci_fusion_apps_fusion_environments",
	DatasourceItemsAttr:    "fusion_environment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fusion_environment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fusion_apps.FusionEnvironmentLifecycleStateActive),
	},
}

var exportFusionAppsFusionEnvironmentDataMaskingActivityHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fusion_apps_fusion_environment_data_masking_activity",
	DatasourceClass:        "oci_fusion_apps_fusion_environment_data_masking_activities",
	DatasourceItemsAttr:    "data_masking_activity_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fusion_environment_data_masking_activity",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fusion_apps.DataMaskingActivityLifecycleStateSucceeded),
	},
}

var fusionAppsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFusionAppsFusionEnvironmentFamilyHints},
		{TerraformResourceHints: exportFusionAppsFusionEnvironmentHints},
	},
	"oci_fusion_apps_fusion_environment": {
		{
			TerraformResourceHints: exportFusionAppsFusionEnvironmentAdminUserHints,
			DatasourceQueryParams: map[string]string{
				"fusion_environment_id": "id",
			},
		},
		{
			TerraformResourceHints: exportFusionAppsFusionEnvironmentDataMaskingActivityHints,
			DatasourceQueryParams: map[string]string{
				"fusion_environment_id": "id",
			},
		},
		{
			TerraformResourceHints: exportFusionAppsFusionEnvironmentRefreshActivityHints,
			DatasourceQueryParams: map[string]string{
				"fusion_environment_id": "id",
			},
		},
	},
}
