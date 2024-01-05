// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opa "github.com/oracle/oci-go-sdk/v65/opa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpaOpaInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpaOpaInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"opa_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpaOpaInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readOpaOpaInstances(d *schema.ResourceData, m interface{}) error {
	sync := &OpaOpaInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpaInstanceClient()

	return tfresource.ReadResource(sync)
}

type OpaOpaInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opa.OpaInstanceClient
	Res    *oci_opa.ListOpaInstancesResponse
}

func (s *OpaOpaInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpaOpaInstancesDataSourceCrud) Get() error {
	request := oci_opa.ListOpaInstancesRequest{}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_opa.OpaInstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opa")

	response, err := s.Client.ListOpaInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOpaInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpaOpaInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpaOpaInstancesDataSource-", OpaOpaInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	opaInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OpaInstanceSummaryToMap(item))
	}
	opaInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpaOpaInstancesDataSource().Schema["opa_instance_collection"].Elem.(*schema.Resource).Schema)
		opaInstance["items"] = items
	}

	resources = append(resources, opaInstance)
	if err := s.D.Set("opa_instance_collection", resources); err != nil {
		return err
	}

	return nil
}
