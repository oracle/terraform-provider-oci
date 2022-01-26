// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreNetworkSecurityGroupSecurityRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreNetworkSecurityGroupSecurityRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"direction": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_security_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"direction": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"icmp_options": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"code": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_valid": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stateless": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"tcp_options": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"destination_port_range": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"max": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"min": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"source_port_range": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"max": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"min": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"udp_options": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"destination_port_range": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"max": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"min": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"source_port_range": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"max": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"min": {
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
				},
			},
		},
	}
}

func readCoreNetworkSecurityGroupSecurityRules(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNetworkSecurityGroupSecurityRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreNetworkSecurityGroupSecurityRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListNetworkSecurityGroupSecurityRulesResponse
}

func (s *CoreNetworkSecurityGroupSecurityRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreNetworkSecurityGroupSecurityRulesDataSourceCrud) Get() error {
	request := oci_core.ListNetworkSecurityGroupSecurityRulesRequest{}

	if direction, ok := s.D.GetOkExists("direction"); ok {
		request.Direction = oci_core.ListNetworkSecurityGroupSecurityRulesDirectionEnum(direction.(string))
	}

	if networkSecurityGroupId, ok := s.D.GetOkExists("network_security_group_id"); ok {
		tmp := networkSecurityGroupId.(string)
		request.NetworkSecurityGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListNetworkSecurityGroupSecurityRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNetworkSecurityGroupSecurityRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreNetworkSecurityGroupSecurityRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreNetworkSecurityGroupSecurityRulesDataSource-", CoreNetworkSecurityGroupSecurityRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		networkSecurityGroupSecurityRule := map[string]interface{}{}

		if r.Description != nil {
			networkSecurityGroupSecurityRule["description"] = *r.Description
		}

		if r.Destination != nil {
			networkSecurityGroupSecurityRule["destination"] = *r.Destination
		}

		networkSecurityGroupSecurityRule["destination_type"] = r.DestinationType

		networkSecurityGroupSecurityRule["direction"] = r.Direction

		if r.IcmpOptions != nil {
			networkSecurityGroupSecurityRule["icmp_options"] = []interface{}{nsgIcmpOptionsToMap(r.IcmpOptions)}
		} else {
			networkSecurityGroupSecurityRule["icmp_options"] = nil
		}

		if r.Id != nil {
			networkSecurityGroupSecurityRule["id"] = *r.Id
		}

		if r.IsValid != nil {
			networkSecurityGroupSecurityRule["is_valid"] = *r.IsValid
		}

		if r.Protocol != nil {
			networkSecurityGroupSecurityRule["protocol"] = *r.Protocol
		}

		if r.Source != nil {
			networkSecurityGroupSecurityRule["source"] = *r.Source
		}

		networkSecurityGroupSecurityRule["source_type"] = r.SourceType

		if r.IsStateless != nil {
			networkSecurityGroupSecurityRule["stateless"] = *r.IsStateless
		}

		if r.TcpOptions != nil {
			networkSecurityGroupSecurityRule["tcp_options"] = []interface{}{nsgTcpOptionsToMap(r.TcpOptions)}
		} else {
			networkSecurityGroupSecurityRule["tcp_options"] = nil
		}

		if r.TimeCreated != nil {
			networkSecurityGroupSecurityRule["time_created"] = r.TimeCreated.String()
		}

		if r.UdpOptions != nil {
			networkSecurityGroupSecurityRule["udp_options"] = []interface{}{nsgUdpOptionsToMap(r.UdpOptions)}
		} else {
			networkSecurityGroupSecurityRule["udp_options"] = nil
		}

		resources = append(resources, networkSecurityGroupSecurityRule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreNetworkSecurityGroupSecurityRulesDataSource().Schema["security_rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("security_rules", resources); err != nil {
		return err
	}

	return nil
}
