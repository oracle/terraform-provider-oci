// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package health_checks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_health_checks "github.com/oracle/oci-go-sdk/v65/healthchecks"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func HealthChecksVantagePointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readHealthChecksVantagePoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_checks_vantage_points": {
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
									"geo_key": {
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
						"provider_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"routing": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"as_label": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"asn": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"prefix": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"weight": {
										Type:     schema.TypeInt,
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

func readHealthChecksVantagePoints(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksVantagePointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.ReadResource(sync)
}

type HealthChecksVantagePointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_health_checks.HealthChecksClient
	Res    *oci_health_checks.ListHealthChecksVantagePointsResponse
}

func (s *HealthChecksVantagePointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *HealthChecksVantagePointsDataSourceCrud) Get() error {
	request := oci_health_checks.ListHealthChecksVantagePointsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "health_checks")

	response, err := s.Client.ListHealthChecksVantagePoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListHealthChecksVantagePoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *HealthChecksVantagePointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("HealthChecksVantagePointsDataSource-", HealthChecksVantagePointsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vantagePoint := map[string]interface{}{}

		if r.DisplayName != nil {
			vantagePoint["display_name"] = *r.DisplayName
		}

		if r.Geo != nil {
			vantagePoint["geo"] = []interface{}{GeolocationToMap(r.Geo)}
		} else {
			vantagePoint["geo"] = nil
		}

		if r.Name != nil {
			vantagePoint["name"] = *r.Name
		}

		if r.ProviderName != nil {
			vantagePoint["provider_name"] = *r.ProviderName
		}

		routing := []interface{}{}
		for _, item := range r.Routing {
			routing = append(routing, RoutingToMap(item))
		}
		vantagePoint["routing"] = routing

		resources = append(resources, vantagePoint)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, HealthChecksVantagePointsDataSource().Schema["health_checks_vantage_points"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("health_checks_vantage_points", resources); err != nil {
		return err
	}

	return nil
}

func GeolocationToMap(obj *oci_health_checks.Geolocation) map[string]interface{} {
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

	if obj.GeoKey != nil {
		result["geo_key"] = string(*obj.GeoKey)
	}

	if obj.Latitude != nil {
		result["latitude"] = float64(*obj.Latitude)
	}

	if obj.Longitude != nil {
		result["longitude"] = float64(*obj.Longitude)
	}

	return result
}

func RoutingToMap(obj oci_health_checks.Routing) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AsLabel != nil {
		result["as_label"] = string(*obj.AsLabel)
	}

	if obj.Asn != nil {
		result["asn"] = int(*obj.Asn)
	}

	if obj.Prefix != nil {
		result["prefix"] = string(*obj.Prefix)
	}

	if obj.Weight != nil {
		result["weight"] = int(*obj.Weight)
	}

	return result
}
