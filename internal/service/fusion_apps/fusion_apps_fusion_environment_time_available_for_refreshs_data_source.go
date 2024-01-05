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

func FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFusionAppsFusionEnvironmentTimeAvailableForRefreshs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_available_for_refresh_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
		},
	}
}

func readFusionAppsFusionEnvironmentTimeAvailableForRefreshs(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListTimeAvailableForRefreshesResponse
}

func (s *FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSourceCrud) Get() error {
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
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTimeAvailableForRefreshes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSource-", FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fusionEnvironmentTimeAvailableForRefresh := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TimeAvailableForRefreshSummaryToMap(item))
	}
	fusionEnvironmentTimeAvailableForRefresh["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSource().Schema["time_available_for_refresh_collection"].Elem.(*schema.Resource).Schema)
		fusionEnvironmentTimeAvailableForRefresh["items"] = items
	}

	resources = append(resources, fusionEnvironmentTimeAvailableForRefresh)
	if err := s.D.Set("time_available_for_refresh_collection", resources); err != nil {
		return err
	}

	return nil
}

func TimeAvailableForRefreshSummaryToMap(obj oci_fusion_apps.TimeAvailableForRefreshSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeAvailableForRefresh != nil {
		result["time_available_for_refresh"] = obj.TimeAvailableForRefresh.String()
	}

	return result
}
