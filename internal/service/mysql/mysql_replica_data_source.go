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

func MysqlReplicaDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["replica_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MysqlReplicaResource(), fieldMap, readSingularMysqlReplica)
}

func readSingularMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()

	return tfresource.ReadResource(sync)
}

type MysqlReplicaDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.ReplicasClient
	Res    *oci_mysql.GetReplicaResponse
}

func (s *MysqlReplicaDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlReplicaDataSourceCrud) Get() error {
	request := oci_mysql.GetReplicaRequest{}

	if replicaId, ok := s.D.GetOkExists("replica_id"); ok {
		tmp := replicaId.(string)
		request.ReplicaId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.GetReplica(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlReplicaDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationId != nil {
		s.D.Set("configuration_id", *s.Res.ConfigurationId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
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

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.IsDeleteProtected != nil {
		s.D.Set("is_delete_protected", *s.Res.IsDeleteProtected)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.PortX != nil {
		s.D.Set("port_x", *s.Res.PortX)
	}

	if s.Res.ReplicaOverrides != nil {
		s.D.Set("replica_overrides", []interface{}{ReplicaOverridesToMap(s.Res.ReplicaOverrides)})
	} else {
		s.D.Set("replica_overrides", nil)
	}

	if s.Res.SecureConnections != nil {
		s.D.Set("secure_connections", []interface{}{SecureConnectionDetailsToMap(s.Res.SecureConnections)})
	} else {
		s.D.Set("secure_connections", nil)
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

	return nil
}
