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

func DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"is_enabled_for_processing": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subsetting_rule_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subsetting_rule_processing_chain_objects_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: subsettingRuleProcessingChainObjectSummarySchema(),
							},
						},
					},
				},
			},
		},
	}
}

// subsettingRuleProcessingChainObjectSummarySchema is the schema for items
// returned by the processing-chain-objects data source. Keep this separate
// from the resource schema: the resource has required input attributes that
// must not be part of a computed data-source summary object.
func subsettingRuleProcessingChainObjectSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"approximate_row_count_before_subsetting": {Type: schema.TypeString, Computed: true},
		"child_columns": {
			Type: schema.TypeList, Computed: true,
			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"child_object_name":                    {Type: schema.TypeString, Computed: true},
		"child_schema_name":                    {Type: schema.TypeString, Computed: true},
		"estimated_row_count_after_subsetting": {Type: schema.TypeString, Computed: true},
		"is_enabled_for_processing":            {Type: schema.TypeBool, Computed: true},
		"key":                                  {Type: schema.TypeString, Computed: true},
		"parent_columns": {
			Type: schema.TypeList, Computed: true,
			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"parent_object_name":             {Type: schema.TypeString, Computed: true},
		"parent_schema_name":             {Type: schema.TypeString, Computed: true},
		"propagation_impact":             {Type: schema.TypeString, Computed: true},
		"subsetting_schema_relation_key": {Type: schema.TypeString, Computed: true},
	}
}

func readDataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingRuleProcessingChainObjectsResponse
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingRuleProcessingChainObjectsRequest{}

	if isEnabledForProcessing, ok := s.D.GetOkExists("is_enabled_for_processing"); ok {
		tmp := isEnabledForProcessing.(bool)
		request.IsEnabledForProcessing = &tmp
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingRuleKey, ok := s.D.GetOkExists("subsetting_rule_key"); ok {
		tmp := subsettingRuleKey.(string)
		request.SubsettingRuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettingRuleProcessingChainObjects(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingRuleProcessingChainObjects(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSource-", DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingPolicySubsettingRuleProcessingChainObject := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingRuleProcessingChainObjectSummaryToMap(item))
	}
	subsettingPolicySubsettingRuleProcessingChainObject["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingPolicySubsettingRuleProcessingChainObjectsDataSource().Schema["subsetting_rule_processing_chain_objects_collection"].Elem.(*schema.Resource).Schema)
		subsettingPolicySubsettingRuleProcessingChainObject["items"] = items
	}

	resources = append(resources, subsettingPolicySubsettingRuleProcessingChainObject)
	if err := s.D.Set("subsetting_rule_processing_chain_objects_collection", resources); err != nil {
		return err
	}

	return nil
}
