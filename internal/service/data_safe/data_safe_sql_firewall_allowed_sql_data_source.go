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

func DataSafeSqlFirewallAllowedSqlDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeSqlFirewallAllowedSql,
		Schema: map[string]*schema.Schema{
			"sql_firewall_allowed_sql_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"sql_accessed_objects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sql_firewall_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sql_level": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sql_text": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_collected": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func readSingularDataSafeSqlFirewallAllowedSql(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallAllowedSqlDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSqlFirewallAllowedSqlDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSqlFirewallAllowedSqlResponse
}

func (s *DataSafeSqlFirewallAllowedSqlDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSqlFirewallAllowedSqlDataSourceCrud) Get() error {
	request := oci_data_safe.GetSqlFirewallAllowedSqlRequest{}

	if sqlFirewallAllowedSqlId, ok := s.D.GetOkExists("sql_firewall_allowed_sql_id"); ok {
		tmp := sqlFirewallAllowedSqlId.(string)
		request.SqlFirewallAllowedSqlId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSqlFirewallAllowedSql(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSqlFirewallAllowedSqlDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentUser != nil {
		s.D.Set("current_user", *s.Res.CurrentUser)
	}

	if s.Res.DbUserName != nil {
		s.D.Set("db_user_name", *s.Res.DbUserName)
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

	s.D.Set("sql_accessed_objects", s.Res.SqlAccessedObjects)

	if s.Res.SqlFirewallPolicyId != nil {
		s.D.Set("sql_firewall_policy_id", *s.Res.SqlFirewallPolicyId)
	}

	s.D.Set("sql_level", s.Res.SqlLevel)

	if s.Res.SqlText != nil {
		s.D.Set("sql_text", *s.Res.SqlText)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCollected != nil {
		s.D.Set("time_collected", s.Res.TimeCollected.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
