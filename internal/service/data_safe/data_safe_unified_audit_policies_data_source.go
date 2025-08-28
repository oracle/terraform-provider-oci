// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeUnifiedAuditPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUnifiedAuditPolicies,
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
			"security_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unified_audit_policy_definition_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unified_audit_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unified_audit_policy_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeUnifiedAuditPolicyResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeUnifiedAuditPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUnifiedAuditPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUnifiedAuditPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListUnifiedAuditPoliciesResponse
}

func (s *DataSafeUnifiedAuditPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUnifiedAuditPoliciesDataSourceCrud) Get() error {
	request := oci_data_safe.ListUnifiedAuditPoliciesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListUnifiedAuditPoliciesAccessLevelEnum(accessLevel.(string))
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

	if securityPolicyId, ok := s.D.GetOkExists("security_policy_id"); ok {
		tmp := securityPolicyId.(string)
		request.SecurityPolicyId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListUnifiedAuditPoliciesLifecycleStateEnum(state.(string))
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

	if unifiedAuditPolicyDefinitionId, ok := s.D.GetOkExists("unified_audit_policy_definition_id"); ok {
		tmp := unifiedAuditPolicyDefinitionId.(string)
		request.UnifiedAuditPolicyDefinitionId = &tmp
	}

	if unifiedAuditPolicyId, ok := s.D.GetOkExists("id"); ok {
		tmp := unifiedAuditPolicyId.(string)
		request.UnifiedAuditPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListUnifiedAuditPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUnifiedAuditPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeUnifiedAuditPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUnifiedAuditPoliciesDataSource-", DataSafeUnifiedAuditPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	unifiedAuditPolicy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UnifiedAuditPolicySummaryToMap(item))
	}
	unifiedAuditPolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeUnifiedAuditPoliciesDataSource().Schema["unified_audit_policy_collection"].Elem.(*schema.Resource).Schema)
		unifiedAuditPolicy["items"] = items
	}

	resources = append(resources, unifiedAuditPolicy)
	if err := s.D.Set("unified_audit_policy_collection", resources); err != nil {
		return err
	}

	return nil
}
