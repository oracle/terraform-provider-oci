// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineAddonDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["addon_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ContainerengineAddonResource(), fieldMap, readSingularContainerengineAddon)
}

func readSingularContainerengineAddon(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineAddonDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineAddonDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetAddonResponse
}

func (s *ContainerengineAddonDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineAddonDataSourceCrud) Get() error {
	request := oci_containerengine.GetAddonRequest{}

	if addonName, ok := s.D.GetOkExists("addon_name"); ok {
		tmp := addonName.(string)
		request.AddonName = &tmp
	}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetAddon(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineAddonDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineAddonDataSource-", ContainerengineAddonDataSource(), s.D))

	if s.Res.AddonError != nil {
		s.D.Set("addon_error", []interface{}{AddonErrorToMap(s.Res.AddonError)})
	} else {
		s.D.Set("addon_error", nil)
	}

	configurations := []interface{}{}
	for _, item := range s.Res.Configurations {
		configurations = append(configurations, AddonConfigurationToMap(item))
	}
	s.D.Set("configurations", configurations)

	if s.Res.CurrentInstalledVersion != nil {
		s.D.Set("current_installed_version", *s.Res.CurrentInstalledVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("addon_name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
