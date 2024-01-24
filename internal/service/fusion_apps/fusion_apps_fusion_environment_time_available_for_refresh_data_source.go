// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularFusionAppsFusionEnvironmentTimeAvailableForRefresh,
		Schema: map[string]*schema.Schema{
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"time_available_for_refresh": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_fusion_apps_fusion_environment_time_available_for_refresh", "oci_fusion_apps_fusion_environment_time_available_for_refreshs"),
	}
}

func readSingularFusionAppsFusionEnvironmentTimeAvailableForRefresh(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListTimeAvailableForRefreshesResponse
}

func (s *FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSourceCrud) Get() error {
	request := oci_fusion_apps.ListTimeAvailableForRefreshesRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.ListTimeAvailableForRefreshes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSource-", FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TimeAvailableForRefreshSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
