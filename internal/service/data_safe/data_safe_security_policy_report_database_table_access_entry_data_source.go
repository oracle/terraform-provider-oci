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

func DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeSecurityPolicyReportDatabaseTableAccessEntry,
		Schema: map[string]*schema.Schema{
			"database_table_access_entry_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_policy_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"access_through_object": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"access_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"are_all_tables_accessible": {
				Type:     schema.TypeBool,
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
			"is_access_constrained_by_label_security": {
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
			"is_access_constrained_by_view": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_access_constrained_by_virtual_private_database": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_sensitive": {
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
		},
	}
}

func readSingularDataSafeSecurityPolicyReportDatabaseTableAccessEntry(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetDatabaseTableAccessEntryResponse
}

func (s *DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSourceCrud) Get() error {
	request := oci_data_safe.GetDatabaseTableAccessEntryRequest{}

	if databaseTableAccessEntryKey, ok := s.D.GetOkExists("database_table_access_entry_key"); ok {
		tmp := databaseTableAccessEntryKey.(string)
		request.DatabaseTableAccessEntryKey = &tmp
	}

	if securityPolicyReportId, ok := s.D.GetOkExists("security_policy_report_id"); ok {
		tmp := securityPolicyReportId.(string)
		request.SecurityPolicyReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetDatabaseTableAccessEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSource-", DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSource(), s.D))

	if s.Res.AccessThroughObject != nil {
		s.D.Set("access_through_object", *s.Res.AccessThroughObject)
	}

	s.D.Set("access_type", s.Res.AccessType)

	if s.Res.AreAllTablesAccessible != nil {
		s.D.Set("are_all_tables_accessible", *s.Res.AreAllTablesAccessible)
	}

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

	if s.Res.IsAccessConstrainedByLabelSecurity != nil {
		s.D.Set("is_access_constrained_by_label_security", *s.Res.IsAccessConstrainedByLabelSecurity)
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

	if s.Res.IsAccessConstrainedByView != nil {
		s.D.Set("is_access_constrained_by_view", *s.Res.IsAccessConstrainedByView)
	}

	if s.Res.IsAccessConstrainedByVirtualPrivateDatabase != nil {
		s.D.Set("is_access_constrained_by_virtual_private_database", *s.Res.IsAccessConstrainedByVirtualPrivateDatabase)
	}

	if s.Res.IsSensitive != nil {
		s.D.Set("is_sensitive", *s.Res.IsSensitive)
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

	return nil
}
