// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementFleetComplianceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularFleetAppsManagementFleetComplianceWithContext,
		Schema: map[string]*schema.Schema{
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compliance_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"confirmed_target_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"non_compliant_target_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularFleetAppsManagementFleetComplianceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetAppsManagementFleetComplianceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetAppsManagementFleetComplianceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.GetComplianceResponse
}

func (s *FleetAppsManagementFleetComplianceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementFleetComplianceDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_apps_management.GetComplianceRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetCompliance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementFleetComplianceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementFleetComplianceDataSource-", FleetAppsManagementFleetComplianceDataSource(), s.D))

	if s.Res.ComplianceState != nil {
		s.D.Set("compliance_state", *s.Res.ComplianceState)
	}

	if s.Res.ConfirmedTargetCount != nil {
		s.D.Set("confirmed_target_count", *s.Res.ConfirmedTargetCount)
	}

	if s.Res.NonCompliantTargetCount != nil {
		s.D.Set("non_compliant_target_count", *s.Res.NonCompliantTargetCount)
	}

	return nil
}
