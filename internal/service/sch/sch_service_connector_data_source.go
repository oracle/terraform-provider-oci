// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package sch

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_sch "github.com/oracle/oci-go-sdk/v65/sch"
)

func SchServiceConnectorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["service_connector_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(SchServiceConnectorResource(), fieldMap, readSingularSchServiceConnector)
}

func readSingularSchServiceConnector(d *schema.ResourceData, m interface{}) error {
	sync := &SchServiceConnectorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceConnectorClient()

	return tfresource.ReadResource(sync)
}

type SchServiceConnectorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_sch.ServiceConnectorClient
	Res    *oci_sch.GetServiceConnectorResponse
}

func (s *SchServiceConnectorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SchServiceConnectorDataSourceCrud) Get() error {
	request := oci_sch.GetServiceConnectorRequest{}

	if serviceConnectorId, ok := s.D.GetOkExists("service_connector_id"); ok {
		tmp := serviceConnectorId.(string)
		request.ServiceConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "sch")

	response, err := s.Client.GetServiceConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SchServiceConnectorDataSourceCrud) SetData() error {
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

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := SourceDetailsToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		s.D.Set("source", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.Target != nil {
		targetArray := []interface{}{}
		if targetMap := TargetDetailsToMap(&s.Res.Target); targetMap != nil {
			targetArray = append(targetArray, targetMap)
		}
		s.D.Set("target", targetArray)
	} else {
		s.D.Set("target", nil)
	}

	tasks := []interface{}{}
	for _, item := range s.Res.Tasks {
		tasks = append(tasks, TaskDetailsToMap(item))
	}
	s.D.Set("tasks", tasks)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
