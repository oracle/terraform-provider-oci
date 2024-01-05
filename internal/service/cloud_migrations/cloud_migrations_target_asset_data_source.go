// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudMigrationsTargetAssetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["target_asset_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudMigrationsTargetAssetResource(), fieldMap, readSingularCloudMigrationsTargetAsset)
}

func readSingularCloudMigrationsTargetAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsTargetAssetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

type CloudMigrationsTargetAssetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_migrations.MigrationClient
	Res    *oci_cloud_migrations.GetTargetAssetResponse
}

func (s *CloudMigrationsTargetAssetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudMigrationsTargetAssetDataSourceCrud) Get() error {
	request := oci_cloud_migrations.GetTargetAssetRequest{}

	if targetAssetId, ok := s.D.GetOkExists("target_asset_id"); ok {
		tmp := targetAssetId.(string)
		request.TargetAssetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_migrations")

	response, err := s.Client.GetTargetAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudMigrationsTargetAssetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.TargetAsset.GetId())
	switch v := (s.Res.TargetAsset).(type) {
	case oci_cloud_migrations.VmTargetAsset:
		s.D.Set("type", "INSTANCE")

		if v.BlockVolumesPerformance != nil {
			s.D.Set("block_volumes_performance", *v.BlockVolumesPerformance)
		}

		if v.MsLicense != nil {
			s.D.Set("ms_license", *v.MsLicense)
		}

		s.D.Set("preferred_shape_type", v.PreferredShapeType)

		if v.RecommendedSpec != nil {
			s.D.Set("recommended_spec", []interface{}{LaunchInstanceDetailsToMap(v.RecommendedSpec, true)})
		} else {
			s.D.Set("recommended_spec", nil)
		}

		if v.TestSpec != nil {
			s.D.Set("test_spec", []interface{}{LaunchInstanceDetailsToMap(v.TestSpec, true)})
		} else {
			s.D.Set("test_spec", nil)
		}

		if v.UserSpec != nil {
			s.D.Set("user_spec", []interface{}{LaunchInstanceDetailsToMap(v.UserSpec, true)})
		} else {
			s.D.Set("user_spec", nil)
		}

		compatibilityMessages := []interface{}{}
		for _, item := range v.CompatibilityMessages {
			compatibilityMessages = append(compatibilityMessages, CompatibilityMessageToMap(item))
		}
		s.D.Set("compatibility_messages", compatibilityMessages)

		if v.CreatedResourceId != nil {
			s.D.Set("created_resource_id", *v.CreatedResourceId)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.EstimatedCost != nil {
			s.D.Set("estimated_cost", []interface{}{CostEstimationToMap(v.EstimatedCost)})
		} else {
			s.D.Set("estimated_cost", nil)
		}

		if v.IsExcludedFromExecution != nil {
			s.D.Set("is_excluded_from_execution", *v.IsExcludedFromExecution)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.MigrationAsset != nil {
			s.D.Set("migration_asset", []interface{}{MigrationAssetToMap(v.MigrationAsset)})
		} else {
			s.D.Set("migration_asset", nil)
		}

		if v.MigrationPlanId != nil {
			s.D.Set("migration_plan_id", *v.MigrationPlanId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeAssessed != nil {
			s.D.Set("time_assessed", v.TimeAssessed.String())
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.TargetAsset)
		return nil
	}

	return nil
}
