// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicySubsettingRulesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingPolicySubsettingRulesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"object": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"schema_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subsetting_rule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataSafeSubsettingPolicySubsettingRuleResource(),
						},
					},
				},
			},
		},
	}
}

func readDataSafeSubsettingPolicySubsettingRulesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicySubsettingRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingRulesResponse
}

func (s *DataSafeSubsettingPolicySubsettingRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicySubsettingRulesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingRulesRequest{}

	if object, ok := s.D.GetOkExists("object"); ok {
		interfaces := object.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object") {
			request.ObjectName = tmp
		}
	}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		interfaces := schemaName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schema_name") {
			request.SchemaName = tmp
		}
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettingRules(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingRules(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingPolicySubsettingRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingPolicySubsettingRulesDataSource-", DataSafeSubsettingPolicySubsettingRulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingPolicySubsettingRule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingRuleSummaryToMap(item))
	}
	subsettingPolicySubsettingRule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingPolicySubsettingRulesDataSource().Schema["subsetting_rule_collection"].Elem.(*schema.Resource).Schema)
		subsettingPolicySubsettingRule["items"] = items
	}

	resources = append(resources, subsettingPolicySubsettingRule)
	if err := s.D.Set("subsetting_rule_collection", resources); err != nil {
		return err
	}

	return nil
}
