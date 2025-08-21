// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeUnifiedAuditPolicyDefinitionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUnifiedAuditPolicyDefinitions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_seeded": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unified_audit_policy_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unified_audit_policy_definition_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unified_audit_policy_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unified_audit_policy_definition_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeUnifiedAuditPolicyDefinitionResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeUnifiedAuditPolicyDefinitions(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnifiedAuditPolicyDefinitionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUnifiedAuditPolicyDefinitionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListUnifiedAuditPolicyDefinitionsResponse
}

func (s *DataSafeUnifiedAuditPolicyDefinitionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUnifiedAuditPolicyDefinitionsDataSourceCrud) Get() error {
	request := oci_data_safe.ListUnifiedAuditPolicyDefinitionsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListUnifiedAuditPolicyDefinitionsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isSeeded, ok := s.D.GetOkExists("is_seeded"); ok {
		tmp := isSeeded.(bool)
		request.IsSeeded = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum(state.(string))
	}

	if unifiedAuditPolicyCategory, ok := s.D.GetOkExists("unified_audit_policy_category"); ok {
		request.UnifiedAuditPolicyCategory = oci_data_safe.UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum(unifiedAuditPolicyCategory.(string))
	}

	if unifiedAuditPolicyDefinitionId, ok := s.D.GetOkExists("id"); ok {
		tmp := unifiedAuditPolicyDefinitionId.(string)
		request.UnifiedAuditPolicyDefinitionId = &tmp
	}

	if unifiedAuditPolicyName, ok := s.D.GetOkExists("unified_audit_policy_name"); ok {
		tmp := unifiedAuditPolicyName.(string)
		request.UnifiedAuditPolicyName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListUnifiedAuditPolicyDefinitions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUnifiedAuditPolicyDefinitions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeUnifiedAuditPolicyDefinitionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUnifiedAuditPolicyDefinitionsDataSource-", DataSafeUnifiedAuditPolicyDefinitionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	unifiedAuditPolicyDefinition := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UnifiedAuditPolicyDefinitionSummaryToMap(item))
	}
	unifiedAuditPolicyDefinition["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeUnifiedAuditPolicyDefinitionsDataSource().Schema["unified_audit_policy_definition_collection"].Elem.(*schema.Resource).Schema)
		unifiedAuditPolicyDefinition["items"] = items
	}

	resources = append(resources, unifiedAuditPolicyDefinition)
	if err := s.D.Set("unified_audit_policy_definition_collection", resources); err != nil {
		return err
	}

	return nil
}
