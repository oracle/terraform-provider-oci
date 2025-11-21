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

func FusionAppsFusionEnvironmentDataMaskingActivitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFusionAppsFusionEnvironmentDataMaskingActivities,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_masking_activity_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FusionAppsFusionEnvironmentDataMaskingActivityResource()),
						},
					},
				},
			},
		},
	}
}

func readFusionAppsFusionEnvironmentDataMaskingActivities(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentDataMaskingActivitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentDataMaskingActivitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListDataMaskingActivitiesResponse
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivitiesDataSourceCrud) Get() error {
	request := oci_fusion_apps.ListDataMaskingActivitiesRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fusion_apps.DataMaskingActivityLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.ListDataMaskingActivities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataMaskingActivities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FusionAppsFusionEnvironmentDataMaskingActivitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentDataMaskingActivitiesDataSource-", FusionAppsFusionEnvironmentDataMaskingActivitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fusionEnvironmentDataMaskingActivity := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataMaskingActivitySummaryToMap(item))
	}
	fusionEnvironmentDataMaskingActivity["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FusionAppsFusionEnvironmentDataMaskingActivitiesDataSource().Schema["data_masking_activity_collection"].Elem.(*schema.Resource).Schema)
		fusionEnvironmentDataMaskingActivity["items"] = items
	}

	resources = append(resources, fusionEnvironmentDataMaskingActivity)
	if err := s.D.Set("data_masking_activity_collection", resources); err != nil {
		return err
	}

	return nil
}
