// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v47/cloudguard"
)

func init() {
	RegisterDatasource("oci_cloud_guard_cloud_guard_configuration", CloudGuardCloudGuardConfigurationDataSource())
}

func CloudGuardCloudGuardConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CloudGuardCloudGuardConfigurationResource(), fieldMap, readSingularCloudGuardCloudGuardConfiguration)
}

func readSingularCloudGuardCloudGuardConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardCloudGuardConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).cloudGuardClient()

	return ReadResource(sync)
}

type CloudGuardCloudGuardConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetConfigurationResponse
}

func (s *CloudGuardCloudGuardConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardCloudGuardConfigurationDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardCloudGuardConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CloudGuardCloudGuardConfigurationDataSource-", CloudGuardCloudGuardConfigurationDataSource(), s.D))

	if s.Res.ReportingRegion != nil {
		s.D.Set("reporting_region", *s.Res.ReportingRegion)
	}

	if s.Res.SelfManageResources != nil {
		s.D.Set("self_manage_resources", *s.Res.SelfManageResources)
	}

	s.D.Set("status", s.Res.Status)

	return nil
}
