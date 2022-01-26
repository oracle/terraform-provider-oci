// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"
)

func MysqlChannelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["channel_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MysqlChannelResource(), fieldMap, readSingularMysqlChannel)
}

func readSingularMysqlChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlChannelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ChannelsClient()

	return tfresource.ReadResource(sync)
}

type MysqlChannelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.ChannelsClient
	Res    *oci_mysql.GetChannelResponse
}

func (s *MysqlChannelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlChannelDataSourceCrud) Get() error {
	request := oci_mysql.GetChannelRequest{}

	if channelId, ok := s.D.GetOkExists("channel_id"); ok {
		tmp := channelId.(string)
		request.ChannelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.GetChannel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlChannelDataSourceCrud) SetData() error {
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

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := ChannelSourceToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		s.D.Set("source", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Target != nil {
		targetArray := []interface{}{}
		if targetMap := ChannelTargetToMap(&s.Res.Target); targetMap != nil {
			targetArray = append(targetArray, targetMap)
		}
		s.D.Set("target", targetArray)
	} else {
		s.D.Set("target", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
