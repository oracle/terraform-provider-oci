// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_dashboard

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_management_dashboard "github.com/oracle/oci-go-sdk/v56/managementdashboard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ManagementDashboardManagementDashboardsExportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularManagementDashboardManagementDashboardsExport,
		Schema: map[string]*schema.Schema{
			"export_dashboard_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"export_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularManagementDashboardManagementDashboardsExport(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementDashboardManagementDashboardsExportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DashxApisClient()

	return tfresource.ReadResource(sync)
}

type ManagementDashboardManagementDashboardsExportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_dashboard.DashxApisClient
	Res    *oci_management_dashboard.ExportDashboardResponse
}

func (s *ManagementDashboardManagementDashboardsExportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementDashboardManagementDashboardsExportDataSourceCrud) Get() error {
	request := oci_management_dashboard.ExportDashboardRequest{}

	if exportDashboardId, ok := s.D.GetOkExists("export_dashboard_id"); ok {
		tmp := exportDashboardId.(string)
		request.ExportDashboardId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_dashboard")

	response, err := s.Client.ExportDashboard(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementDashboardManagementDashboardsExportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementDashboardManagementDashboardsExportDataSource-", ManagementDashboardManagementDashboardsExportDataSource(), s.D))

	var exportDetailsBytes, err = json.Marshal(s.Res.ManagementDashboardExportDetails)
	if err != nil {
		return fmt.Errorf("unable to Marshal ManagementDashboardExportDetails, encountered error: %v", err)
	}
	exportDetails := string(exportDetailsBytes)
	s.D.Set("export_details", exportDetails)

	return nil
}
