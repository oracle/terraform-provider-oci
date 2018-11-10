// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func RuleSetDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularRuleSet,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ADD_HTTP_REQUEST_HEADER",
								"ADD_HTTP_RESPONSE_HEADER",
								"EXTEND_HTTP_REQUEST_HEADER_VALUE",
								"EXTEND_HTTP_RESPONSE_HEADER_VALUE",
								"REMOVE_HTTP_REQUEST_HEADER",
								"REMOVE_HTTP_RESPONSE_HEADER",
							}, true),
						},
						"header": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"suffix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
		},
	}
}

func readSingularRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &RuleSetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return ReadResource(sync)
}

type RuleSetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetRuleSetResponse
}

func (s *RuleSetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RuleSetDataSourceCrud) Get() error {
	request := oci_load_balancer.GetRuleSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.RuleSetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RuleSetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RuleToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
