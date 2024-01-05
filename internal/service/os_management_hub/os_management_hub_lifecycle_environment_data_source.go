// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubLifecycleEnvironmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["lifecycle_environment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubLifecycleEnvironmentResource(), fieldMap, readSingularOsManagementHubLifecycleEnvironment)
}

func readSingularOsManagementHubLifecycleEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleEnvironmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubLifecycleEnvironmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.LifecycleEnvironmentClient
	Res    *oci_os_management_hub.GetLifecycleEnvironmentResponse
}

func (s *OsManagementHubLifecycleEnvironmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubLifecycleEnvironmentDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetLifecycleEnvironmentRequest{}

	if lifecycleEnvironmentId, ok := s.D.GetOkExists("lifecycle_environment_id"); ok {
		tmp := lifecycleEnvironmentId.(string)
		request.LifecycleEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetLifecycleEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubLifecycleEnvironmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("arch_type", s.Res.ArchType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	managedInstanceIds := []interface{}{}
	for _, item := range s.Res.ManagedInstanceIds {
		managedInstanceIds = append(managedInstanceIds, ManagedInstanceDetailsToMap(item))
	}
	s.D.Set("managed_instance_ids", managedInstanceIds)

	s.D.Set("os_family", s.Res.OsFamily)

	stages := []interface{}{}
	for _, item := range s.Res.Stages {
		stages = append(stages, LifecycleStageToMap(item))
	}
	s.D.Set("stages", stages)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeModified != nil {
		s.D.Set("time_modified", s.Res.TimeModified.String())
	}

	s.D.Set("vendor_name", s.Res.VendorName)

	return nil
}
