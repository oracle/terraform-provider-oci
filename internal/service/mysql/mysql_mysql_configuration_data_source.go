// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlMysqlConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MysqlMysqlConfigurationResource(), fieldMap, readSingularMysqlMysqlConfiguration)
}

func readSingularMysqlMysqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()

	return tfresource.ReadResource(sync)
}

type MysqlMysqlConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.MysqlaasClient
	Res    *oci_mysql.GetConfigurationResponse
}

func (s *MysqlMysqlConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlConfigurationDataSourceCrud) Get() error {
	request := oci_mysql.GetConfigurationRequest{}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlMysqlConfigurationDataSourceCrud) SetData() error {
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

	if s.Res.InitVariables != nil {
		s.D.Set("init_variables", []interface{}{InitializationVariablesToMap(s.Res.InitVariables)})
	} else {
		s.D.Set("init_variables", nil)
	}

	if s.Res.ParentConfigurationId != nil {
		s.D.Set("parent_configuration_id", *s.Res.ParentConfigurationId)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.Variables != nil {
		s.D.Set("variables", []interface{}{ConfigurationVariablesToMap(s.Res.Variables)})
	} else {
		s.D.Set("variables", nil)
	}

	return nil
}
