// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package zpr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_zpr "github.com/oracle/oci-go-sdk/v65/zpr"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ZprZprPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readZprZprPolicies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
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
			"zpr_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ZprZprPolicyResource()),
						},
					},
				},
			},
		},
	}
}

func readZprZprPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &ZprZprPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ZprClient()

	return tfresource.ReadResource(sync)
}

type ZprZprPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_zpr.ZprClient
	Res    *oci_zpr.ListZprPoliciesResponse
}

func (s *ZprZprPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ZprZprPoliciesDataSourceCrud) Get() error {
	request := oci_zpr.ListZprPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_zpr.ZprPolicyLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "zpr")

	response, err := s.Client.ListZprPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListZprPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ZprZprPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ZprZprPoliciesDataSource-", ZprZprPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	zprPolicy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ZprPolicySummaryToMap(item))
	}
	zprPolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ZprZprPoliciesDataSource().Schema["zpr_policies"].Elem.(*schema.Resource).Schema)
		zprPolicy["items"] = items
	}

	resources = append(resources, zprPolicy)
	if err := s.D.Set("zpr_policies", resources); err != nil {
		return err
	}

	return nil
}
