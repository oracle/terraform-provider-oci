// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LicenseManagerConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LicenseManagerConfigurationResource(), fieldMap, readSingularLicenseManagerConfiguration)
}

func readSingularLicenseManagerConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

type LicenseManagerConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_license_manager.LicenseManagerClient
	Res    *oci_license_manager.GetConfigurationResponse
}

func (s *LicenseManagerConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LicenseManagerConfigurationDataSourceCrud) Get() error {
	request := oci_license_manager.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LicenseManagerConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LicenseManagerConfigurationDataSource-", LicenseManagerConfigurationDataSource(), s.D))

	s.D.Set("email_ids", s.Res.EmailIds)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
