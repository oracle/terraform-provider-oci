// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package health_checks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_health_checks "github.com/oracle/oci-go-sdk/v56/healthchecks"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func HealthChecksHttpProbeResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readHealthChecksHttpProbeResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"probe_configuration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_time_greater_than_or_equal_to": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"start_time_less_than_or_equal_to": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_probe_results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"connect_end": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"connect_start": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"connection": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connect_duration": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"secure_connect_duration": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"dns": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"addresses": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"domain_lookup_duration": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"domain_lookup_end": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"domain_lookup_start": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"duration": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"encoded_body_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"error_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fetch_start": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"is_healthy": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_timed_out": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"probe_configuration_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"request_start": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"response_end": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"response_start": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"secure_connection_start": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"start_time": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"status_code": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"target": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vantage_point_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readHealthChecksHttpProbeResults(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksHttpProbeResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.ReadResource(sync)
}

type HealthChecksHttpProbeResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_health_checks.HealthChecksClient
	Res    *oci_health_checks.ListHttpProbeResultsResponse
}

func (s *HealthChecksHttpProbeResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *HealthChecksHttpProbeResultsDataSourceCrud) Get() error {
	request := oci_health_checks.ListHttpProbeResultsRequest{}

	if probeConfigurationId, ok := s.D.GetOkExists("probe_configuration_id"); ok {
		tmp := probeConfigurationId.(string)
		request.ProbeConfigurationId = &tmp
	}

	if startTimeGreaterThanOrEqualTo, ok := s.D.GetOkExists("start_time_greater_than_or_equal_to"); ok {
		tmp := startTimeGreaterThanOrEqualTo.(float64)
		request.StartTimeGreaterThanOrEqualTo = &tmp
	}

	if startTimeLessThanOrEqualTo, ok := s.D.GetOkExists("start_time_less_than_or_equal_to"); ok {
		tmp := startTimeLessThanOrEqualTo.(float64)
		request.StartTimeLessThanOrEqualTo = &tmp
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		tmp := target.(string)
		request.Target = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "health_checks")

	response, err := s.Client.ListHttpProbeResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListHttpProbeResults(context.Background(), request)
		if err != nil {
			return err
		}
		// health check always returns valid OpcNextPage, which causes infinite loop.
		if listResponse.Items == nil || len(listResponse.Items) == 0 {
			request.Page = nil
		} else {
			s.Res.Items = append(s.Res.Items, listResponse.Items...)
			request.Page = listResponse.OpcNextPage
		}
	}

	return nil
}

func (s *HealthChecksHttpProbeResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("HealthChecksHttpProbeResultsDataSource-", HealthChecksHttpProbeResultsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		httpProbeResult := map[string]interface{}{
			"probe_configuration_id": *r.ProbeConfigurationId,
		}

		if r.ConnectEnd != nil {
			httpProbeResult["connect_end"] = *r.ConnectEnd
		}

		if r.ConnectStart != nil {
			httpProbeResult["connect_start"] = *r.ConnectStart
		}

		if r.Connection != nil {
			httpProbeResult["connection"] = []interface{}{TcpConnectionToMap(r.Connection)}
		} else {
			httpProbeResult["connection"] = nil
		}

		if r.Dns != nil {
			httpProbeResult["dns"] = []interface{}{DNSToMap(r.Dns)}
		} else {
			httpProbeResult["dns"] = nil
		}

		if r.DomainLookupEnd != nil {
			httpProbeResult["domain_lookup_end"] = *r.DomainLookupEnd
		}

		if r.DomainLookupStart != nil {
			httpProbeResult["domain_lookup_start"] = *r.DomainLookupStart
		}

		if r.Duration != nil {
			httpProbeResult["duration"] = *r.Duration
		}

		if r.EncodedBodySize != nil {
			httpProbeResult["encoded_body_size"] = *r.EncodedBodySize
		}

		httpProbeResult["error_category"] = r.ErrorCategory

		if r.ErrorMessage != nil {
			httpProbeResult["error_message"] = *r.ErrorMessage
		}

		if r.FetchStart != nil {
			httpProbeResult["fetch_start"] = *r.FetchStart
		}

		if r.IsHealthy != nil {
			httpProbeResult["is_healthy"] = *r.IsHealthy
		}

		if r.IsTimedOut != nil {
			httpProbeResult["is_timed_out"] = *r.IsTimedOut
		}

		if r.Key != nil {
			httpProbeResult["key"] = *r.Key
		}

		httpProbeResult["protocol"] = r.Protocol

		if r.RequestStart != nil {
			httpProbeResult["request_start"] = *r.RequestStart
		}

		if r.ResponseEnd != nil {
			httpProbeResult["response_end"] = *r.ResponseEnd
		}

		if r.ResponseStart != nil {
			httpProbeResult["response_start"] = *r.ResponseStart
		}

		if r.SecureConnectionStart != nil {
			httpProbeResult["secure_connection_start"] = *r.SecureConnectionStart
		}

		if r.StartTime != nil {
			httpProbeResult["start_time"] = *r.StartTime
		}

		if r.StatusCode != nil {
			httpProbeResult["status_code"] = *r.StatusCode
		}

		if r.Target != nil {
			httpProbeResult["target"] = *r.Target
		}

		if r.VantagePointName != nil {
			httpProbeResult["vantage_point_name"] = *r.VantagePointName
		}

		resources = append(resources, httpProbeResult)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, HealthChecksHttpProbeResultsDataSource().Schema["http_probe_results"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("http_probe_results", resources); err != nil {
		return err
	}

	return nil
}

func DNSToMap(obj *oci_health_checks.Dns) map[string]interface{} {
	result := map[string]interface{}{}

	result["addresses"] = obj.Addresses

	if obj.DomainLookupDuration != nil {
		result["domain_lookup_duration"] = float64(*obj.DomainLookupDuration)
	}

	return result
}

func TcpConnectionToMap(obj *oci_health_checks.TcpConnection) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address != nil {
		result["address"] = string(*obj.Address)
	}

	if obj.ConnectDuration != nil {
		result["connect_duration"] = float64(*obj.ConnectDuration)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.SecureConnectDuration != nil {
		result["secure_connect_duration"] = float64(*obj.SecureConnectDuration)
	}

	return result
}
