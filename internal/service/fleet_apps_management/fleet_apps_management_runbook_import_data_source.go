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

func FleetAppsManagementRunbookImportDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularFleetAppsManagementRunbookImportWithContext,
		Schema: map[string]*schema.Schema{
			"import_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"runbook_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"runbook_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"runbook_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tracking_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularFleetAppsManagementRunbookImportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetAppsManagementRunbookImportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FleetAppsManagementRunbookImportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementRunbooksClient
	Res    *oci_fleet_apps_management.GetRunbookImportResponse
}

func (s *FleetAppsManagementRunbookImportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementRunbookImportDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_apps_management.GetRunbookImportRequest{}

	if importId, ok := s.D.GetOkExists("import_id"); ok {
		tmp := importId.(string)
		request.ImportId = &tmp
	}

	if runbookId, ok := s.D.GetOkExists("runbook_id"); ok {
		tmp := runbookId.(string)
		request.RunbookId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetRunbookImport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementRunbookImportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementRunbookImportDataSource-", FleetAppsManagementRunbookImportDataSource(), s.D))

	s.D.Set("details", s.Res.Details)

	if s.Res.RunbookName != nil {
		s.D.Set("runbook_name", *s.Res.RunbookName)
	}

	if s.Res.RunbookVersion != nil {
		s.D.Set("runbook_version", *s.Res.RunbookVersion)
	}

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	if s.Res.TrackingId != nil {
		s.D.Set("tracking_id", *s.Res.TrackingId)
	}

	return nil
}
