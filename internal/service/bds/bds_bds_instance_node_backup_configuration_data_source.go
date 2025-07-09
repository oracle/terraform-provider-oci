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

func BdsBdsInstanceNodeBackupConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["node_backup_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BdsBdsInstanceNodeBackupConfigurationResource(), fieldMap, readSingularBdsBdsInstanceNodeBackupConfiguration)
}

func readSingularBdsBdsInstanceNodeBackupConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceNodeBackupConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetNodeBackupConfigurationResponse
}

func (s *BdsBdsInstanceNodeBackupConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceNodeBackupConfigurationDataSourceCrud) Get() error {
	request := oci_bds.GetNodeBackupConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if nodeBackupConfigurationId, ok := s.D.GetOkExists("node_backup_configuration_id"); ok {
		tmp := nodeBackupConfigurationId.(string)
		request.NodeBackupConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetNodeBackupConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceNodeBackupConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
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

	if s.Res.NumberOfBackupsToRetain != nil {
		s.D.Set("number_of_backups_to_retain", *s.Res.NumberOfBackupsToRetain)
	}

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	return nil
}
