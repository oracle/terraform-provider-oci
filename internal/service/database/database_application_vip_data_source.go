// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseApplicationVipDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["application_vip_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseApplicationVipResource(), fieldMap, readSingularDatabaseApplicationVip)
}

func readSingularDatabaseApplicationVip(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseApplicationVipDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseApplicationVipDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetApplicationVipResponse
}

func (s *DatabaseApplicationVipDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseApplicationVipDataSourceCrud) Get() error {
	request := oci_database.GetApplicationVipRequest{}

	if applicationVipId, ok := s.D.GetOkExists("application_vip_id"); ok {
		tmp := applicationVipId.(string)
		request.ApplicationVipId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetApplicationVip(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseApplicationVipDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CloudVmClusterId != nil {
		s.D.Set("cloud_vm_cluster_id", *s.Res.CloudVmClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostnameLabel != nil {
		s.D.Set("hostname_label", *s.Res.HostnameLabel)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeAssigned != nil {
		s.D.Set("time_assigned", s.Res.TimeAssigned.String())
	}

	return nil
}
