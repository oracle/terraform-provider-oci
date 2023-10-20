// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryDrProtectionGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDisasterRecoveryDrProtectionGroups,
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
			"dr_protection_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_sub_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dr_protection_group_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DisasterRecoveryDrProtectionGroupResource()),
						},
					},
				},
			},
		},
	}
}

func readDisasterRecoveryDrProtectionGroups(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrProtectionGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.ReadResource(sync)
}

type DisasterRecoveryDrProtectionGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_disaster_recovery.DisasterRecoveryClient
	Res    *oci_disaster_recovery.ListDrProtectionGroupsResponse
}

func (s *DisasterRecoveryDrProtectionGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DisasterRecoveryDrProtectionGroupsDataSourceCrud) Get() error {
	request := oci_disaster_recovery.ListDrProtectionGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drProtectionGroupId, ok := s.D.GetOkExists("id"); ok {
		tmp := drProtectionGroupId.(string)
		request.DrProtectionGroupId = &tmp
	}

	if lifecycleSubState, ok := s.D.GetOkExists("lifecycle_sub_state"); ok {
		request.LifecycleSubState = oci_disaster_recovery.ListDrProtectionGroupsLifecycleSubStateEnum(lifecycleSubState.(string))
	}

	if role, ok := s.D.GetOkExists("role"); ok {
		request.Role = oci_disaster_recovery.ListDrProtectionGroupsRoleEnum(role.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_disaster_recovery.ListDrProtectionGroupsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "disaster_recovery")

	response, err := s.Client.ListDrProtectionGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrProtectionGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DisasterRecoveryDrProtectionGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DisasterRecoveryDrProtectionGroupsDataSource-", DisasterRecoveryDrProtectionGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	drProtectionGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DrProtectionGroupSummaryToMap(item))
	}
	drProtectionGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DisasterRecoveryDrProtectionGroupsDataSource().Schema["dr_protection_group_collection"].Elem.(*schema.Resource).Schema)
		drProtectionGroup["items"] = items
	}

	resources = append(resources, drProtectionGroup)
	if err := s.D.Set("dr_protection_group_collection", resources); err != nil {
		return err
	}

	return nil
}
