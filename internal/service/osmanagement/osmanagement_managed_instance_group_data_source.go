// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v56/osmanagement"
)

func OsmanagementManagedInstanceGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["managed_instance_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsmanagementManagedInstanceGroupResource(), fieldMap, readSingularOsmanagementManagedInstanceGroup)
}

func readSingularOsmanagementManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementManagedInstanceGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.GetManagedInstanceGroupResponse
}

func (s *OsmanagementManagedInstanceGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstanceGroupDataSourceCrud) Get() error {
	request := oci_osmanagement.GetManagedInstanceGroupRequest{}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.GetManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsmanagementManagedInstanceGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	managedInstances := []interface{}{}
	for _, item := range s.Res.ManagedInstances {
		managedInstances = append(managedInstances, IdToMap(item))
	}
	s.D.Set("managed_instances", managedInstances)

	s.D.Set("os_family", s.Res.OsFamily)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
