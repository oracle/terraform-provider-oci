// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globally_distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GloballyDistributedDatabasePrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GloballyDistributedDatabasePrivateEndpointResource(), fieldMap, readSingularGloballyDistributedDatabasePrivateEndpoint)
}

func readSingularGloballyDistributedDatabasePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabasePrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	return tfresource.ReadResource(sync)
}

type GloballyDistributedDatabasePrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_globally_distributed_database.ShardedDatabaseServiceClient
	Res    *oci_globally_distributed_database.GetPrivateEndpointResponse
}

func (s *GloballyDistributedDatabasePrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GloballyDistributedDatabasePrivateEndpointDataSourceCrud) Get() error {
	request := oci_globally_distributed_database.GetPrivateEndpointRequest{}

	if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "globally_distributed_database")

	response, err := s.Client.GetPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GloballyDistributedDatabasePrivateEndpointDataSourceCrud) SetData() error {
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

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.PrivateIp != nil {
		s.D.Set("private_ip", *s.Res.PrivateIp)
	}

	s.D.Set("sharded_databases", s.Res.ShardedDatabases)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
