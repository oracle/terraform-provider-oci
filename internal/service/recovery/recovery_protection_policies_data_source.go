// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RecoveryProtectionPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readRecoveryProtectionPoliciesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protection_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protection_policy_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(RecoveryProtectionPolicyResource()),
						},
					},
				},
			},
		},
	}
}

func readRecoveryProtectionPoliciesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &RecoveryProtectionPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type RecoveryProtectionPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_recovery.DatabaseRecoveryClient
	Res    *oci_recovery.ListProtectionPoliciesResponse
}

func (s *RecoveryProtectionPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RecoveryProtectionPoliciesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_recovery.ListProtectionPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if owner, ok := s.D.GetOkExists("owner"); ok {
		request.Owner = oci_recovery.ListProtectionPoliciesOwnerEnum(owner.(string))
	}

	if protectionPolicyId, ok := s.D.GetOkExists("id"); ok {
		tmp := protectionPolicyId.(string)
		request.ProtectionPolicyId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_recovery.ListProtectionPoliciesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "recovery")

	response, err := s.Client.ListProtectionPolicies(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProtectionPolicies(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RecoveryProtectionPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RecoveryProtectionPoliciesDataSource-", RecoveryProtectionPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	protectionPolicy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProtectionPolicySummaryToMap(item))
	}
	protectionPolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, RecoveryProtectionPoliciesDataSource().Schema["protection_policy_collection"].Elem.(*schema.Resource).Schema)
		protectionPolicy["items"] = items
	}

	resources = append(resources, protectionPolicy)
	if err := s.D.Set("protection_policy_collection", resources); err != nil {
		return err
	}

	return nil
}
