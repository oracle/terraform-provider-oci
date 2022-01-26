// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v56/apmsynthetics"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsPublicVantagePointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmSyntheticsPublicVantagePoints,
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
			"public_vantage_point_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"geo": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"admin_div_code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"city_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"country_code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"country_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"latitude": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"longitude": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"name": {
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

func readApmSyntheticsPublicVantagePoints(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsPublicVantagePointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsPublicVantagePointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.ListPublicVantagePointsResponse
}

func (s *ApmSyntheticsPublicVantagePointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsPublicVantagePointsDataSourceCrud) Get() error {
	request := oci_apm_synthetics.ListPublicVantagePointsRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.ListPublicVantagePoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPublicVantagePoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmSyntheticsPublicVantagePointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmSyntheticsPublicVantagePointsDataSource-", ApmSyntheticsPublicVantagePointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	publicVantagePoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PublicVantagePointSummaryToMap(item))
	}
	publicVantagePoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApmSyntheticsPublicVantagePointsDataSource().Schema["public_vantage_point_collection"].Elem.(*schema.Resource).Schema)
		publicVantagePoint["items"] = items
	}

	resources = append(resources, publicVantagePoint)
	if err := s.D.Set("public_vantage_point_collection", resources); err != nil {
		return err
	}

	return nil
}
