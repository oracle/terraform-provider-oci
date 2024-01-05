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

func HealthChecksPingProbeResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readHealthChecksPingProbeResults,
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
			"ping_probe_results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"connection": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"dns": {
							Type:     schema.TypeList,
							Computed: true,
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
						"error_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"icmp_code": {
							Type:     schema.TypeInt,
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
						"latency_in_ms": {
							Type:     schema.TypeFloat,
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
						"start_time": {
							Type:     schema.TypeFloat,
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

func readHealthChecksPingProbeResults(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingProbeResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.ReadResource(sync)
}

type HealthChecksPingProbeResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_health_checks.HealthChecksClient
	Res    *oci_health_checks.ListPingProbeResultsResponse
}

func (s *HealthChecksPingProbeResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *HealthChecksPingProbeResultsDataSourceCrud) Get() error {
	request := oci_health_checks.ListPingProbeResultsRequest{}

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

	response, err := s.Client.ListPingProbeResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPingProbeResults(context.Background(), request)
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

func (s *HealthChecksPingProbeResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("HealthChecksPingProbeResultsDataSource-", HealthChecksPingProbeResultsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		pingProbeResult := map[string]interface{}{
			"probe_configuration_id": *r.ProbeConfigurationId,
		}

		if r.Connection != nil {
			pingProbeResult["connection"] = []interface{}{ConnectionToMap(r.Connection)}
		} else {
			pingProbeResult["connection"] = nil
		}

		if r.Dns != nil {
			pingProbeResult["dns"] = []interface{}{DNSToMap(r.Dns)}
		} else {
			pingProbeResult["dns"] = nil
		}

		if r.DomainLookupEnd != nil {
			pingProbeResult["domain_lookup_end"] = *r.DomainLookupEnd
		}

		if r.DomainLookupStart != nil {
			pingProbeResult["domain_lookup_start"] = *r.DomainLookupStart
		}

		pingProbeResult["error_category"] = r.ErrorCategory

		if r.ErrorMessage != nil {
			pingProbeResult["error_message"] = *r.ErrorMessage
		}

		if r.IcmpCode != nil {
			pingProbeResult["icmp_code"] = *r.IcmpCode
		}

		if r.IsHealthy != nil {
			pingProbeResult["is_healthy"] = *r.IsHealthy
		}

		if r.IsTimedOut != nil {
			pingProbeResult["is_timed_out"] = *r.IsTimedOut
		}

		if r.Key != nil {
			pingProbeResult["key"] = *r.Key
		}

		if r.LatencyInMs != nil {
			pingProbeResult["latency_in_ms"] = *r.LatencyInMs
		}

		pingProbeResult["protocol"] = r.Protocol

		if r.StartTime != nil {
			pingProbeResult["start_time"] = *r.StartTime
		}

		if r.Target != nil {
			pingProbeResult["target"] = *r.Target
		}

		if r.VantagePointName != nil {
			pingProbeResult["vantage_point_name"] = *r.VantagePointName
		}

		resources = append(resources, pingProbeResult)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, HealthChecksPingProbeResultsDataSource().Schema["ping_probe_results"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("ping_probe_results", resources); err != nil {
		return err
	}

	return nil
}

func ConnectionToMap(obj *oci_health_checks.Connection) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address != nil {
		result["address"] = string(*obj.Address)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	return result
}
