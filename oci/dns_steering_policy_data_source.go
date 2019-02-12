// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	oci_dns "github.com/oracle/oci-go-sdk/dns"
)

func DnsSteeringPolicyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDnsSteeringPolicy,
		Schema: map[string]*schema.Schema{
			"steering_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"answers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"rdata": {
							Type:     schema.TypeString,
							Required: true,
						},
						"rtype": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"is_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"pool": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"health_check_monitor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"rule_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FILTER",
								"HEALTH",
								"LIMIT",
								"PRIORITY",
								"WEIGHTED",
							}, true),
						},

						// Optional
						"cases": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"answer_data": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"answer_condition": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"should_keep": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"case_condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"default_answer_data": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"answer_condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"should_keep": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"default_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"self": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularDnsSteeringPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient

	return ReadResource(sync)
}

type DnsSteeringPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetSteeringPolicyResponse
}

func (s *DnsSteeringPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsSteeringPolicyDataSourceCrud) Get() error {
	request := oci_dns.GetSteeringPolicyRequest{}

	if steeringPolicyId, ok := s.D.GetOkExists("steering_policy_id"); ok {
		tmp := steeringPolicyId.(string)
		request.SteeringPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "dns")

	response, err := s.Client.GetSteeringPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DnsSteeringPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	answers := []interface{}{}
	for _, item := range s.Res.Answers {
		answers = append(answers, SteeringPolicyAnswerToMap(item))
	}
	s.D.Set("answers", answers)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HealthCheckMonitorId != nil {
		s.D.Set("health_check_monitor_id", *s.Res.HealthCheckMonitorId)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, SteeringPolicyRuleToMap(item))
	}
	s.D.Set("rules", rules)

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("template", s.Res.Template)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Ttl != nil {
		s.D.Set("ttl", *s.Res.Ttl)
	}

	return nil
}
