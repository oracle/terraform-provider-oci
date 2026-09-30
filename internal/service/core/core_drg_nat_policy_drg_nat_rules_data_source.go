// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreDrgNatPolicyDrgNatRulesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readCoreDrgNatPolicyDrgNatRulesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"drg_nat_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_nat_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"drg_nat_policy_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"drg_nat_rule_priority": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"original_destination": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"original_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"translated_destination": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"translated_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreDrgNatPolicyDrgNatRulesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreDrgNatPolicyDrgNatRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type CoreDrgNatPolicyDrgNatRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDrgNatRulesResponse
}

func (s *CoreDrgNatPolicyDrgNatRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgNatPolicyDrgNatRulesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_core.ListDrgNatRulesRequest{}

	if drgNatPolicyId, ok := s.D.GetOkExists("drg_nat_policy_id"); ok {
		tmp := drgNatPolicyId.(string)
		request.DrgNatPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDrgNatRules(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgNatRules(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDrgNatPolicyDrgNatRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDrgNatPolicyDrgNatRulesDataSource-", CoreDrgNatPolicyDrgNatRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drgNatPolicyDrgNatRule := map[string]interface{}{
			"drg_nat_policy_id": *r.DrgNatPolicyId,
		}

		if r.DrgNatRulePriority != nil {
			drgNatPolicyDrgNatRule["drg_nat_rule_priority"] = strconv.FormatInt(*r.DrgNatRulePriority, 10)
		}

		if r.Id != nil {
			drgNatPolicyDrgNatRule["id"] = *r.Id
		}

		if r.OriginalDestination != nil {
			drgNatPolicyDrgNatRule["original_destination"] = *r.OriginalDestination
		}

		if r.OriginalSource != nil {
			drgNatPolicyDrgNatRule["original_source"] = *r.OriginalSource
		}

		if r.TranslatedDestination != nil {
			drgNatPolicyDrgNatRule["translated_destination"] = *r.TranslatedDestination
		}

		if r.TranslatedSource != nil {
			drgNatPolicyDrgNatRule["translated_source"] = *r.TranslatedSource
		}

		resources = append(resources, drgNatPolicyDrgNatRule)
	}

	sort.Slice(resources, func(i, j int) bool {
		iPriority, _ := strconv.Atoi(resources[i]["drg_nat_rule_priority"].(string))
		jPriority, _ := strconv.Atoi(resources[j]["drg_nat_rule_priority"].(string))
		return iPriority < jPriority
	})

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDrgNatPolicyDrgNatRulesDataSource().Schema["drg_nat_rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drg_nat_rules", resources); err != nil {
		return err
	}

	return nil
}
