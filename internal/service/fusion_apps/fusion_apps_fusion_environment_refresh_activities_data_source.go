// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentRefreshActivitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFusionAppsFusionEnvironmentRefreshActivities,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_expected_finish_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_scheduled_start_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"refresh_activity_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FusionAppsFusionEnvironmentRefreshActivityResource()),
						},
					},
				},
			},
		},
	}
}

func readFusionAppsFusionEnvironmentRefreshActivities(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentRefreshActivitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentRefreshActivitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListRefreshActivitiesResponse
}

func (s *FusionAppsFusionEnvironmentRefreshActivitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentRefreshActivitiesDataSourceCrud) Get() error {
	request := oci_fusion_apps.ListRefreshActivitiesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fusion_apps.RefreshActivityLifecycleStateEnum(state.(string))
	}

	if timeExpectedFinishLessThanOrEqualTo, ok := s.D.GetOkExists("time_expected_finish_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpectedFinishLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeExpectedFinishLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeScheduledStartGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_scheduled_start_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduledStartGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeScheduledStartGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.ListRefreshActivities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRefreshActivities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FusionAppsFusionEnvironmentRefreshActivitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentRefreshActivitiesDataSource-", FusionAppsFusionEnvironmentRefreshActivitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fusionEnvironmentRefreshActivity := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RefreshActivitySummaryToMap(item))
	}
	fusionEnvironmentRefreshActivity["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FusionAppsFusionEnvironmentRefreshActivitiesDataSource().Schema["refresh_activity_collection"].Elem.(*schema.Resource).Schema)
		fusionEnvironmentRefreshActivity["items"] = items
	}

	resources = append(resources, fusionEnvironmentRefreshActivity)
	if err := s.D.Set("refresh_activity_collection", resources); err != nil {
		return err
	}

	return nil
}
