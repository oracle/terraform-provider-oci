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

func FusionAppsFusionEnvironmentFamiliesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFusionAppsFusionEnvironmentFamilies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fusion_environment_family_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fusion_environment_family_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FusionAppsFusionEnvironmentFamilyResource()),
						},
					},
				},
			},
		},
	}
}

func readFusionAppsFusionEnvironmentFamilies(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentFamiliesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentFamiliesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListFusionEnvironmentFamiliesResponse
}

func (s *FusionAppsFusionEnvironmentFamiliesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentFamiliesDataSourceCrud) Get() error {
	request := oci_fusion_apps.ListFusionEnvironmentFamiliesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fusionEnvironmentFamilyId, ok := s.D.GetOkExists("id"); ok {
		tmp := fusionEnvironmentFamilyId.(string)
		request.FusionEnvironmentFamilyId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.ListFusionEnvironmentFamilies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFusionEnvironmentFamilies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FusionAppsFusionEnvironmentFamiliesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentFamiliesDataSource-", FusionAppsFusionEnvironmentFamiliesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fusionEnvironmentFamily := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FusionEnvironmentFamilySummaryToMap(item))
	}
	fusionEnvironmentFamily["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FusionAppsFusionEnvironmentFamiliesDataSource().Schema["fusion_environment_family_collection"].Elem.(*schema.Resource).Schema)
		fusionEnvironmentFamily["items"] = items
	}

	resources = append(resources, fusionEnvironmentFamily)
	if err := s.D.Set("fusion_environment_family_collection", resources); err != nil {
		return err
	}

	return nil
}
