// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v58/apmsynthetics"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsPublicVantagePointDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApmSyntheticsPublicVantagePoint,
		Schema: map[string]*schema.Schema{
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
	}
}

func readSingularApmSyntheticsPublicVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsPublicVantagePointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsPublicVantagePointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.ListPublicVantagePointsResponse
}

func (s *ApmSyntheticsPublicVantagePointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsPublicVantagePointDataSourceCrud) Get() error {
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
	return nil
}

func (s *ApmSyntheticsPublicVantagePointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmSyntheticsPublicVantagePointDataSource-", ApmSyntheticsPublicVantagePointDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PublicVantagePointSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func GeoSummaryToMap(obj *oci_apm_synthetics.GeoSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminDivCode != nil {
		result["admin_div_code"] = string(*obj.AdminDivCode)
	}

	if obj.CityName != nil {
		result["city_name"] = string(*obj.CityName)
	}

	if obj.CountryCode != nil {
		result["country_code"] = string(*obj.CountryCode)
	}

	if obj.CountryName != nil {
		result["country_name"] = string(*obj.CountryName)
	}

	if obj.Latitude != nil {
		result["latitude"] = float64(*obj.Latitude)
	}

	if obj.Longitude != nil {
		result["longitude"] = float64(*obj.Longitude)
	}

	return result
}

func PublicVantagePointSummaryToMap(obj oci_apm_synthetics.PublicVantagePointSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Geo != nil {
		result["geo"] = []interface{}{GeoSummaryToMap(obj.Geo)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
