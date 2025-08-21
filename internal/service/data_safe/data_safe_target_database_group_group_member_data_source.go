// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeTargetDatabaseGroupGroupMemberDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeTargetDatabaseGroupGroupMember,
		Schema: map[string]*schema.Schema{
			"target_database_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"target_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularDataSafeTargetDatabaseGroupGroupMember(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseGroupGroupMemberDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabaseGroupGroupMemberDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetGroupMembersResponse
}

func (s *DataSafeTargetDatabaseGroupGroupMemberDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabaseGroupGroupMemberDataSourceCrud) Get() error {
	request := oci_data_safe.GetGroupMembersRequest{}

	if targetDatabaseGroupId, ok := s.D.GetOkExists("target_database_group_id"); ok {
		tmp := targetDatabaseGroupId.(string)
		request.TargetDatabaseGroupId = &tmp
	}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetGroupMembers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeTargetDatabaseGroupGroupMemberDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetDatabaseGroupGroupMemberDataSource-", DataSafeTargetDatabaseGroupGroupMemberDataSource(), s.D))

	s.D.Set("target_databases", s.Res.TargetDatabases)

	return nil
}
