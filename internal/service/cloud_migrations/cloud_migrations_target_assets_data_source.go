// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudMigrationsTargetAssetsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readCloudMigrationsTargetAssetsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"migration_plan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_asset_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_asset_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudMigrationsTargetAssetResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudMigrationsTargetAssetsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CloudMigrationsTargetAssetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type CloudMigrationsTargetAssetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_migrations.MigrationClient
	Res    *oci_cloud_migrations.ListTargetAssetsResponse
}

func (s *CloudMigrationsTargetAssetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudMigrationsTargetAssetsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_cloud_migrations.ListTargetAssetsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if migrationPlanId, ok := s.D.GetOkExists("migration_plan_id"); ok {
		tmp := migrationPlanId.(string)
		request.MigrationPlanId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_migrations.TargetAssetLifecycleStateEnum(state.(string))
	}

	if targetAssetId, ok := s.D.GetOkExists("id"); ok {
		tmp := targetAssetId.(string)
		request.TargetAssetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_migrations")

	response, err := s.Client.ListTargetAssets(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTargetAssets(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudMigrationsTargetAssetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudMigrationsTargetAssetsDataSource-", CloudMigrationsTargetAssetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	targetAsset := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TargetAssetSummaryToMap(item, true))
	}
	targetAsset["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudMigrationsTargetAssetsDataSource().Schema["target_asset_collection"].Elem.(*schema.Resource).Schema)
		targetAsset["items"] = items
	}

	resources = append(resources, targetAsset)
	if err := s.D.Set("target_asset_collection", resources); err != nil {
		return err
	}

	return nil
}
