// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vbs_inst

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_vbs_inst "github.com/oracle/oci-go-sdk/v65/vbsinst"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func VbsInstVbsInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVbsInstVbsInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vbs_instance_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(VbsInstVbsInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readVbsInstVbsInstances(d *schema.ResourceData, m interface{}) error {
	sync := &VbsInstVbsInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbsInstanceClient()

	return tfresource.ReadResource(sync)
}

type VbsInstVbsInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_vbs_inst.VbsInstanceClient
	Res    *oci_vbs_inst.ListVbsInstancesResponse
}

func (s *VbsInstVbsInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VbsInstVbsInstancesDataSourceCrud) Get() error {
	request := oci_vbs_inst.ListVbsInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_vbs_inst.ListVbsInstancesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "vbs_inst")

	response, err := s.Client.ListVbsInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVbsInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VbsInstVbsInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("VbsInstVbsInstancesDataSource-", VbsInstVbsInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	vbsInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, VbsInstanceSummaryToMap(item))
	}
	vbsInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, VbsInstVbsInstancesDataSource().Schema["vbs_instance_summary_collection"].Elem.(*schema.Resource).Schema)
		vbsInstance["items"] = items
	}

	resources = append(resources, vbsInstance)
	if err := s.D.Set("vbs_instance_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
