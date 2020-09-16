// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v25/osmanagement"
)

func init() {
	RegisterDatasource("oci_osmanagement_managed_instance_group", OsmanagementManagedInstanceGroupDataSource())
}

func OsmanagementManagedInstanceGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["managed_instance_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(OsmanagementManagedInstanceGroupResource(), fieldMap, readSingularOsmanagementManagedInstanceGroup)
}

func readSingularOsmanagementManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).osManagementClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "osmanagement")

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
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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
