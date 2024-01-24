// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_waas "github.com/oracle/oci-go-sdk/v65/waas"
)

func WaasCustomProtectionRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWaasCustomProtectionRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"states": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_protection_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(WaasCustomProtectionRuleResource()),
			},
		},
	}
}

func readWaasCustomProtectionRules(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCustomProtectionRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasCustomProtectionRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.ListCustomProtectionRulesResponse
}

func (s *WaasCustomProtectionRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasCustomProtectionRulesDataSourceCrud) Get() error {
	request := oci_waas.ListCustomProtectionRulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayNames, ok := s.D.GetOkExists("display_names"); ok {
		interfaces := displayNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_names") {
			request.DisplayName = tmp
		}
	}

	if ids, ok := s.D.GetOkExists("ids"); ok {
		interfaces := ids.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ids") {
			request.Id = tmp
		}
	}

	if states, ok := s.D.GetOkExists("states"); ok {
		interfaces := states.([]interface{})
		tmp := make([]oci_waas.LifecycleStatesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waas.LifecycleStatesEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("states") {
			request.LifecycleState = tmp
		}
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.ListCustomProtectionRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCustomProtectionRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WaasCustomProtectionRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WaasCustomProtectionRulesDataSource-", WaasCustomProtectionRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		customProtectionRule := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			customProtectionRule["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			customProtectionRule["display_name"] = *r.DisplayName
		}

		customProtectionRule["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			customProtectionRule["id"] = *r.Id
		}

		customProtectionRule["mod_security_rule_ids"] = r.ModSecurityRuleIds

		customProtectionRule["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			customProtectionRule["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, customProtectionRule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, WaasCustomProtectionRulesDataSource().Schema["custom_protection_rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("custom_protection_rules", resources); err != nil {
		return err
	}

	return nil
}
