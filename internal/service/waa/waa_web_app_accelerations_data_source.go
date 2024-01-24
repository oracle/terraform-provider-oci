// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WaaWebAppAccelerationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWaaWebAppAccelerations,
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"web_app_acceleration_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_app_acceleration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(WaaWebAppAccelerationResource()),
						},
					},
				},
			},
		},
	}
}

func readWaaWebAppAccelerations(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()

	return tfresource.ReadResource(sync)
}

type WaaWebAppAccelerationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waa.WaaClient
	Res    *oci_waa.ListWebAppAccelerationsResponse
}

func (s *WaaWebAppAccelerationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaaWebAppAccelerationsDataSourceCrud) Get() error {
	request := oci_waa.ListWebAppAccelerationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	// allows filtering by multiple values
	if state, ok := s.D.GetOkExists("state"); ok {
		interfaces := state.([]interface{})
		tmp := make([]oci_waa.WebAppAccelerationLifecycleStateEnum, len(interfaces))

		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waa.WebAppAccelerationLifecycleStateEnum(interfaces[i].(string))
			}
		}

		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	if webAppAccelerationPolicyId, ok := s.D.GetOkExists("web_app_acceleration_policy_id"); ok {
		tmp := webAppAccelerationPolicyId.(string)
		request.WebAppAccelerationPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waa")

	response, err := s.Client.ListWebAppAccelerations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWebAppAccelerations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WaaWebAppAccelerationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WaaWebAppAccelerationsDataSource-", WaaWebAppAccelerationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	webAppAcceleration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WebAppAccelerationSummaryToMap(item))
	}
	webAppAcceleration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WaaWebAppAccelerationsDataSource().Schema["web_app_acceleration_collection"].Elem.(*schema.Resource).Schema)
		webAppAcceleration["items"] = items
	}

	resources = append(resources, webAppAcceleration)
	if err := s.D.Set("web_app_acceleration_collection", resources); err != nil {
		return err
	}

	return nil
}
