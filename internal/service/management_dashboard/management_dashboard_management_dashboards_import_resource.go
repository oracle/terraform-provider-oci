// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_dashboard

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_management_dashboard "github.com/oracle/oci-go-sdk/v56/managementdashboard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ManagementDashboardManagementDashboardsImportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createManagementDashboardManagementDashboardsImport,
		Read:     readManagementDashboardManagementDashboardsImport,
		Delete:   deleteManagementDashboardManagementDashboardsImport,
		Schema: map[string]*schema.Schema{
			// Optional
			"import_details_file": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"import_details"},
			},
			"import_details": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"import_details_file"},
			},
		},
	}
}

func createManagementDashboardManagementDashboardsImport(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementDashboardManagementDashboardsImportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DashxApisClient()

	return tfresource.CreateResource(d, sync)
}

func readManagementDashboardManagementDashboardsImport(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementDashboardManagementDashboardsImport(d *schema.ResourceData, m interface{}) error {
	return nil
}

type ManagementDashboardManagementDashboardsImportResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_management_dashboard.DashxApisClient
	DisableNotFoundRetries bool
}

func (s *ManagementDashboardManagementDashboardsImportResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("ManagementDashboardManagementDashboardsImportResource-", ManagementDashboardManagementDashboardsImportResource(), s.D)
}

func (s *ManagementDashboardManagementDashboardsImportResourceCrud) Create() error {
	request := oci_management_dashboard.ImportDashboardRequest{}

	var flag = false
	if importDetails, ok := s.D.GetOkExists("import_details"); ok {
		flag = true
		var importDetailsObj oci_management_dashboard.ManagementDashboardImportDetails
		err := json.Unmarshal([]byte(importDetails.(string)), &importDetailsObj)
		if err != nil {
			return err
		}
		request.ManagementDashboardImportDetails = importDetailsObj
	}

	if importDetailsFilePath, ok := s.D.GetOkExists("import_details_file"); ok {
		flag = true
		importDetailsFileData, err := ioutil.ReadFile(importDetailsFilePath.(string))
		if err != nil {
			return fmt.Errorf("unable to read import_details_file: %s", err)
		}

		var importDetailsObj oci_management_dashboard.ManagementDashboardImportDetails
		err = json.Unmarshal(importDetailsFileData, &importDetailsObj)
		if err != nil {
			return err
		}
		request.ManagementDashboardImportDetails = importDetailsObj
	}

	if !flag {
		return fmt.Errorf("Either import_details or import_details_file must be provided")
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_dashboard")

	_, err := s.Client.ImportDashboard(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *ManagementDashboardManagementDashboardsImportResourceCrud) SetData() error {
	return nil
}
