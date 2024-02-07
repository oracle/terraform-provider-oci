// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
)

func MysqlMysqlDbSystemDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["db_system_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MysqlMysqlDbSystemResource(), fieldMap, readSingularMysqlMysqlDbSystem)
}

func readSingularMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	return tfresource.ReadResource(sync)
}

type MysqlMysqlDbSystemDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.DbSystemClient
	Res    *oci_mysql.GetDbSystemResponse
}

func (s *MysqlMysqlDbSystemDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlDbSystemDataSourceCrud) Get() error {
	request := oci_mysql.GetDbSystemRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.GetDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlMysqlDbSystemDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.BackupPolicy != nil {
		s.D.Set("backup_policy", []interface{}{BackupPolicyToMap(s.Res.BackupPolicy)})
	} else {
		s.D.Set("backup_policy", nil)
	}

	channels := []interface{}{}
	for _, item := range s.Res.Channels {
		channels = append(channels, ChannelSummaryToMap(item))
	}
	s.D.Set("channels", channels)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationId != nil {
		s.D.Set("configuration_id", *s.Res.ConfigurationId)
	}

	s.D.Set("crash_recovery", s.Res.CrashRecovery)

	if s.Res.CurrentPlacement != nil {
		s.D.Set("current_placement", []interface{}{DbSystemPlacementToMap(s.Res.CurrentPlacement)})
	} else {
		s.D.Set("current_placement", nil)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	s.D.Set("database_management", s.Res.DatabaseManagement)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeletionPolicy != nil {
		s.D.Set("deletion_policy", []interface{}{DeletionPolicyDetailsToMap(s.Res.DeletionPolicy)})
	} else {
		s.D.Set("deletion_policy", nil)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	endpoints := []interface{}{}
	for _, item := range s.Res.Endpoints {
		endpoints = append(endpoints, DbSystemEndpointToMap(item))
	}
	s.D.Set("endpoints", endpoints)

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HeatWaveCluster != nil {
		s.D.Set("heat_wave_cluster", []interface{}{HeatWaveClusterSummaryToMap(s.Res.HeatWaveCluster)})
	} else {
		s.D.Set("heat_wave_cluster", nil)
	}

	if s.Res.HostnameLabel != nil {
		s.D.Set("hostname_label", *s.Res.HostnameLabel)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.IsHeatWaveClusterAttached != nil {
		s.D.Set("is_heat_wave_cluster_attached", *s.Res.IsHeatWaveClusterAttached)
	}

	if s.Res.IsHighlyAvailable != nil {
		s.D.Set("is_highly_available", *s.Res.IsHighlyAvailable)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Maintenance != nil {
		s.D.Set("maintenance", []interface{}{MaintenanceDetailsToMap(s.Res.Maintenance)})
	} else {
		s.D.Set("maintenance", nil)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	if s.Res.PointInTimeRecoveryDetails != nil {
		s.D.Set("point_in_time_recovery_details", []interface{}{PointInTimeRecoveryDetailsToMap(s.Res.PointInTimeRecoveryDetails)})
	} else {
		s.D.Set("point_in_time_recovery_details", nil)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.PortX != nil {
		s.D.Set("port_x", *s.Res.PortX)
	}

	if s.Res.SecureConnections != nil {
		s.D.Set("secure_connections", []interface{}{SecureConnectionDetailsToMap(s.Res.SecureConnections)})
	} else {
		s.D.Set("secure_connections", nil)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := DbSystemSourceToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		result := map[string]interface{}{}
		result["source_type"] = "NONE"
		sourceArray := []interface{}{}
		sourceArray = append(sourceArray, result)
		s.D.Set("source", sourceArray)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
