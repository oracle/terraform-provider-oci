// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package disaster_recovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DisasterRecoveryAutomaticDrConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDisasterRecoveryAutomaticDrConfigurationsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"automatic_dr_configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dr_protection_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lifecycle_state_not_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"automatic_dr_configuration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DisasterRecoveryAutomaticDrConfigurationResource()),
						},
					},
				},
			},
		},
	}
}

func readDisasterRecoveryAutomaticDrConfigurationsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DisasterRecoveryAutomaticDrConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DisasterRecoveryClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DisasterRecoveryAutomaticDrConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_disaster_recovery.DisasterRecoveryClient
	Res    *oci_disaster_recovery.ListAutomaticDrConfigurationsResponse
}

func (s *DisasterRecoveryAutomaticDrConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DisasterRecoveryAutomaticDrConfigurationsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_disaster_recovery.ListAutomaticDrConfigurationsRequest{}

	if automaticDrConfigurationId, ok := s.D.GetOkExists("id"); ok {
		tmp := automaticDrConfigurationId.(string)
		request.AutomaticDrConfigurationId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drProtectionGroupId, ok := s.D.GetOkExists("dr_protection_group_id"); ok {
		tmp := drProtectionGroupId.(string)
		request.DrProtectionGroupId = &tmp
	}

	if lifecycleStateNotEqualTo, ok := s.D.GetOkExists("lifecycle_state_not_equal_to"); ok {
		request.LifecycleStateNotEqualTo = oci_disaster_recovery.ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum(lifecycleStateNotEqualTo.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_disaster_recovery.ListAutomaticDrConfigurationsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "disaster_recovery")

	response, err := s.Client.ListAutomaticDrConfigurations(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutomaticDrConfigurations(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DisasterRecoveryAutomaticDrConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DisasterRecoveryAutomaticDrConfigurationsDataSource-", DisasterRecoveryAutomaticDrConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	automaticDrConfiguration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AutomaticDrConfigurationSummaryToMap(item))
	}
	automaticDrConfiguration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DisasterRecoveryAutomaticDrConfigurationsDataSource().Schema["automatic_dr_configuration_collection"].Elem.(*schema.Resource).Schema)
		automaticDrConfiguration["items"] = items
	}

	resources = append(resources, automaticDrConfiguration)
	if err := s.D.Set("automatic_dr_configuration_collection", resources); err != nil {
		return err
	}

	return nil
}
