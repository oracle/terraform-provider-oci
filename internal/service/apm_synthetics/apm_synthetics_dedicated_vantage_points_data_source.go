// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsDedicatedVantagePointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmSyntheticsDedicatedVantagePoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dedicated_vantage_point_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApmSyntheticsDedicatedVantagePointResource()),
						},
					},
				},
			},
		},
	}
}

func readApmSyntheticsDedicatedVantagePoints(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsDedicatedVantagePointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsDedicatedVantagePointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.ListDedicatedVantagePointsResponse
}

func (s *ApmSyntheticsDedicatedVantagePointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsDedicatedVantagePointsDataSourceCrud) Get() error {
	request := oci_apm_synthetics.ListDedicatedVantagePointsRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.ListDedicatedVantagePoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDedicatedVantagePoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmSyntheticsDedicatedVantagePointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmSyntheticsDedicatedVantagePointsDataSource-", ApmSyntheticsDedicatedVantagePointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	dedicatedVantagePoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DedicatedVantagePointSummaryToMap(item))
	}
	dedicatedVantagePoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApmSyntheticsDedicatedVantagePointsDataSource().Schema["dedicated_vantage_point_collection"].Elem.(*schema.Resource).Schema)
		dedicatedVantagePoint["items"] = items
	}

	resources = append(resources, dedicatedVantagePoint)
	if err := s.D.Set("dedicated_vantage_point_collection", resources); err != nil {
		return err
	}

	return nil
}
