// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package visual_builder

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_visual_builder "github.com/oracle/oci-go-sdk/v56/visualbuilder"
)

func VisualBuilderVbInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVisualBuilderVbInstances,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vb_instance_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(VisualBuilderVbInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readVisualBuilderVbInstances(d *schema.ResourceData, m interface{}) error {
	sync := &VisualBuilderVbInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbInstanceClient()

	return tfresource.ReadResource(sync)
}

type VisualBuilderVbInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_visual_builder.VbInstanceClient
	Res    *oci_visual_builder.ListVbInstancesResponse
}

func (s *VisualBuilderVbInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VisualBuilderVbInstancesDataSourceCrud) Get() error {
	request := oci_visual_builder.ListVbInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_visual_builder.ListVbInstancesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "visual_builder")

	response, err := s.Client.ListVbInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVbInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VisualBuilderVbInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("VisualBuilderVbInstancesDataSource-", VisualBuilderVbInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	vbInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, VbInstanceSummaryToMap(item))
	}
	vbInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, VisualBuilderVbInstancesDataSource().Schema["vb_instance_summary_collection"].Elem.(*schema.Resource).Schema)
		vbInstance["items"] = items
	}

	resources = append(resources, vbInstance)
	if err := s.D.Set("vb_instance_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
