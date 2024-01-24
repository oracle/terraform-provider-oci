// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryDrPlansDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDisasterRecoveryDrPlans,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dr_plan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dr_plan_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dr_protection_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dr_plan_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DisasterRecoveryDrPlanResource()),
						},
					},
				},
			},
		},
	}
}

func readDisasterRecoveryDrPlans(d *schema.ResourceData, m interface{}) error {
	sync := &DisasterRecoveryDrPlansDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.ReadResource(sync)
}

type DisasterRecoveryDrPlansDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_disaster_recovery.DisasterRecoveryClient
	Res    *oci_disaster_recovery.ListDrPlansResponse
}

func (s *DisasterRecoveryDrPlansDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DisasterRecoveryDrPlansDataSourceCrud) Get() error {
	request := oci_disaster_recovery.ListDrPlansRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drPlanId, ok := s.D.GetOkExists("id"); ok {
		tmp := drPlanId.(string)
		request.DrPlanId = &tmp
	}

	if drPlanType, ok := s.D.GetOkExists("dr_plan_type"); ok {
		request.DrPlanType = oci_disaster_recovery.ListDrPlansDrPlanTypeEnum(drPlanType.(string))
	}

	if drProtectionGroupId, ok := s.D.GetOkExists("dr_protection_group_id"); ok {
		tmp := drProtectionGroupId.(string)
		request.DrProtectionGroupId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_disaster_recovery.ListDrPlansLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "disaster_recovery")

	response, err := s.Client.ListDrPlans(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrPlans(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DisasterRecoveryDrPlansDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DisasterRecoveryDrPlansDataSource-", DisasterRecoveryDrPlansDataSource(), s.D))
	resources := []map[string]interface{}{}
	drPlan := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DrPlanSummaryToMap(item))
	}
	drPlan["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DisasterRecoveryDrPlansDataSource().Schema["dr_plan_collection"].Elem.(*schema.Resource).Schema)
		drPlan["items"] = items
	}

	resources = append(resources, drPlan)
	if err := s.D.Set("dr_plan_collection", resources); err != nil {
		return err
	}

	return nil
}
