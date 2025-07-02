// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceNodeReplaceConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["node_replace_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BdsBdsInstanceNodeReplaceConfigurationResource(), fieldMap, readSingularBdsBdsInstanceNodeReplaceConfiguration)
}

func readSingularBdsBdsInstanceNodeReplaceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeReplaceConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceNodeReplaceConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetNodeReplaceConfigurationResponse
}

func (s *BdsBdsInstanceNodeReplaceConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceNodeReplaceConfigurationDataSourceCrud) Get() error {
	request := oci_bds.GetNodeReplaceConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if nodeReplaceConfigurationId, ok := s.D.GetOkExists("node_replace_configuration_id"); ok {
		tmp := nodeReplaceConfigurationId.(string)
		request.NodeReplaceConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetNodeReplaceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceNodeReplaceConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DurationInMinutes != nil {
		s.D.Set("duration_in_minutes", *s.Res.DurationInMinutes)
	}

	if s.Res.LevelTypeDetails != nil {
		levelTypeDetailsArray := []interface{}{}
		if levelTypeDetailsMap := LevelTypeDetailsToMap(&s.Res.LevelTypeDetails); levelTypeDetailsMap != nil {
			levelTypeDetailsArray = append(levelTypeDetailsArray, levelTypeDetailsMap)
		}
		s.D.Set("level_type_details", levelTypeDetailsArray)
	} else {
		s.D.Set("level_type_details", nil)
	}

	s.D.Set("metric_type", s.Res.MetricType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
