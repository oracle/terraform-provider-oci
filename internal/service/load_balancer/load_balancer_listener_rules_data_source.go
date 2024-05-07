// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LoadBalancerListenerRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerListenerRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"listener_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"listener_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"action": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"allowed_methods": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"are_invalid_characters_allowed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"conditions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"attribute_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"operator": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"default_max_connections": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"header": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"http_large_header_size_in_kb": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ip_max_connections": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"ip_addresses": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"max_connections": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"prefix": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"redirect_uri": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"host": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"port": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"protocol": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"query": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"response_code": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"status_code": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"suffix": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readLoadBalancerListenerRules(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerListenerRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerListenerRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListListenerRulesResponse
}

func (s *LoadBalancerListenerRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerListenerRulesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListListenerRulesRequest{}

	if listenerName, ok := s.D.GetOkExists("listener_name"); ok {
		tmp := listenerName.(string)
		request.ListenerName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListListenerRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerListenerRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerListenerRulesDataSource-", LoadBalancerListenerRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		listenerRule := map[string]interface{}{}

		if r.RuleSetName != nil {
			listenerRule["name"] = *r.RuleSetName
		}

		if r.Rule != nil {
			ruleArray := []interface{}{}
			if ruleMap := RuleToMap(r.Rule, true); ruleMap != nil {
				ruleArray = append(ruleArray, ruleMap)
			}
			listenerRule["rule"] = ruleArray
		} else {
			listenerRule["rule"] = nil
		}

		resources = append(resources, listenerRule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoadBalancerListenerRulesDataSource().Schema["listener_rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("listener_rules", resources); err != nil {
		return err
	}

	return nil
}
