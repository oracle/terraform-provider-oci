// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedDatabasePrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["distributed_database_private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DistributedDatabaseDistributedDatabasePrivateEndpointResource(), fieldMap, readSingularDistributedDatabaseDistributedDatabasePrivateEndpointWithContext)
}

func readSingularDistributedDatabaseDistributedDatabasePrivateEndpointWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbPrivateEndpointServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedDbPrivateEndpointServiceClient
	Res    *oci_distributed_database.GetDistributedDatabasePrivateEndpointResponse
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedDatabasePrivateEndpointRequest{}

	if distributedDatabasePrivateEndpointId, ok := s.D.GetOkExists("distributed_database_private_endpoint_id"); ok {
		tmp := distributedDatabasePrivateEndpointId.(string)
		request.DistributedDatabasePrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.GetDistributedDatabasePrivateEndpoint(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointDataSourceCrud) SetData() error {
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

	globallyDistributedAutonomousDatabases := []interface{}{}
	for _, item := range s.Res.GloballyDistributedAutonomousDatabases {
		globallyDistributedAutonomousDatabases = append(globallyDistributedAutonomousDatabases, DistributedAutonomousDatabaseAssociatedWithPrivateEndpointToMap(item))
	}
	s.D.Set("globally_distributed_autonomous_databases", globallyDistributedAutonomousDatabases)

	globallyDistributedDatabases := []interface{}{}
	for _, item := range s.Res.GloballyDistributedDatabases {
		globallyDistributedDatabases = append(globallyDistributedDatabases, DistributedDatabaseAssociatedWithPrivateEndpointToMap(item))
	}
	s.D.Set("globally_distributed_databases", globallyDistributedDatabases)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.PrivateIp != nil {
		s.D.Set("private_ip", *s.Res.PrivateIp)
	}

	if s.Res.ProxyComputeInstanceId != nil {
		s.D.Set("proxy_compute_instance_id", *s.Res.ProxyComputeInstanceId)
	}

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
