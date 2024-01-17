// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeSecurityPolicyReportDatabaseViewAccessEntry,
		Schema: map[string]*schema.Schema{
			"database_view_access_entry_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_policy_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"access_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"column_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"grant_from_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"grantee": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"grantor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_access_constrained_by_database_vault": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_access_constrained_by_real_application_security": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_access_constrained_by_redaction": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_access_constrained_by_sql_firewall": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_access_constrained_by_virtual_private_database": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privilege": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privilege_grantable": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privilege_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"table_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"table_schema": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"view_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"view_schema": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"view_text": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDataSafeSecurityPolicyReportDatabaseViewAccessEntry(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetDatabaseViewAccessEntryResponse
}

func (s *DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSourceCrud) Get() error {
	request := oci_data_safe.GetDatabaseViewAccessEntryRequest{}

	if databaseViewAccessEntryKey, ok := s.D.GetOkExists("database_view_access_entry_key"); ok {
		tmp := databaseViewAccessEntryKey.(string)
		request.DatabaseViewAccessEntryKey = &tmp
	}

	if securityPolicyReportId, ok := s.D.GetOkExists("security_policy_report_id"); ok {
		tmp := securityPolicyReportId.(string)
		request.SecurityPolicyReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetDatabaseViewAccessEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSource-", DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSource(), s.D))

	s.D.Set("access_type", s.Res.AccessType)

	if s.Res.ColumnName != nil {
		s.D.Set("column_name", *s.Res.ColumnName)
	}

	if s.Res.GrantFromRole != nil {
		s.D.Set("grant_from_role", *s.Res.GrantFromRole)
	}

	if s.Res.Grantee != nil {
		s.D.Set("grantee", *s.Res.Grantee)
	}

	if s.Res.Grantor != nil {
		s.D.Set("grantor", *s.Res.Grantor)
	}

	if s.Res.IsAccessConstrainedByDatabaseVault != nil {
		s.D.Set("is_access_constrained_by_database_vault", *s.Res.IsAccessConstrainedByDatabaseVault)
	}

	if s.Res.IsAccessConstrainedByRealApplicationSecurity != nil {
		s.D.Set("is_access_constrained_by_real_application_security", *s.Res.IsAccessConstrainedByRealApplicationSecurity)
	}

	if s.Res.IsAccessConstrainedByRedaction != nil {
		s.D.Set("is_access_constrained_by_redaction", *s.Res.IsAccessConstrainedByRedaction)
	}

	if s.Res.IsAccessConstrainedBySqlFirewall != nil {
		s.D.Set("is_access_constrained_by_sql_firewall", *s.Res.IsAccessConstrainedBySqlFirewall)
	}

	if s.Res.IsAccessConstrainedByVirtualPrivateDatabase != nil {
		s.D.Set("is_access_constrained_by_virtual_private_database", *s.Res.IsAccessConstrainedByVirtualPrivateDatabase)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("privilege", s.Res.Privilege)

	s.D.Set("privilege_grantable", s.Res.PrivilegeGrantable)

	if s.Res.PrivilegeType != nil {
		s.D.Set("privilege_type", *s.Res.PrivilegeType)
	}

	if s.Res.TableName != nil {
		s.D.Set("table_name", *s.Res.TableName)
	}

	if s.Res.TableSchema != nil {
		s.D.Set("table_schema", *s.Res.TableSchema)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.ViewName != nil {
		s.D.Set("view_name", *s.Res.ViewName)
	}

	if s.Res.ViewSchema != nil {
		s.D.Set("view_schema", *s.Res.ViewSchema)
	}

	if s.Res.ViewText != nil {
		s.D.Set("view_text", *s.Res.ViewText)
	}

	return nil
}
