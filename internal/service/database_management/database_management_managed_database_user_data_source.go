// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v56/databasemanagement"
)

func DatabaseManagementManagedDatabaseUserDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseUser,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"all_shared": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"common": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"consumer_group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_collation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_tablespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"editions_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_shared": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"implicit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inherited": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_temp_tablespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oracle_maintained": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password_versions": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_connect": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"temp_tablespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_expiring": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_login": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_locked": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_password_changed": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseUser(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseUserDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseUserDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetUserResponse
}

func (s *DatabaseManagementManagedDatabaseUserDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseUserDataSourceCrud) Get() error {
	request := oci_database_management.GetUserRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if userName, ok := s.D.GetOkExists("user_name"); ok {
		tmp := userName.(string)
		request.UserName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseUserDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseUserDataSource-", DatabaseManagementManagedDatabaseUserDataSource(), s.D))

	s.D.Set("all_shared", s.Res.AllShared)

	s.D.Set("authentication", s.Res.Authentication)

	s.D.Set("common", s.Res.Common)

	if s.Res.ConsumerGroup != nil {
		s.D.Set("consumer_group", *s.Res.ConsumerGroup)
	}

	if s.Res.DefaultCollation != nil {
		s.D.Set("default_collation", *s.Res.DefaultCollation)
	}

	if s.Res.DefaultTablespace != nil {
		s.D.Set("default_tablespace", *s.Res.DefaultTablespace)
	}

	s.D.Set("editions_enabled", s.Res.EditionsEnabled)

	if s.Res.ExternalName != nil {
		s.D.Set("external_name", *s.Res.ExternalName)
	}

	s.D.Set("external_shared", s.Res.ExternalShared)

	s.D.Set("implicit", s.Res.Implicit)

	s.D.Set("inherited", s.Res.Inherited)

	if s.Res.LocalTempTablespace != nil {
		s.D.Set("local_temp_tablespace", *s.Res.LocalTempTablespace)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("oracle_maintained", s.Res.OracleMaintained)

	if s.Res.PasswordVersions != nil {
		s.D.Set("password_versions", *s.Res.PasswordVersions)
	}

	if s.Res.Profile != nil {
		s.D.Set("profile", *s.Res.Profile)
	}

	s.D.Set("proxy_connect", s.Res.ProxyConnect)

	s.D.Set("status", s.Res.Status)

	if s.Res.TempTablespace != nil {
		s.D.Set("temp_tablespace", *s.Res.TempTablespace)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpiring != nil {
		s.D.Set("time_expiring", s.Res.TimeExpiring.String())
	}

	if s.Res.TimeLastLogin != nil {
		s.D.Set("time_last_login", s.Res.TimeLastLogin.String())
	}

	if s.Res.TimeLocked != nil {
		s.D.Set("time_locked", s.Res.TimeLocked.String())
	}

	if s.Res.TimePasswordChanged != nil {
		s.D.Set("time_password_changed", s.Res.TimePasswordChanged.String())
	}

	return nil
}

func UserSummaryToMap(obj oci_database_management.UserSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultTablespace != nil {
		result["default_tablespace"] = string(*obj.DefaultTablespace)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Profile != nil {
		result["profile"] = string(*obj.Profile)
	}

	result["status"] = string(obj.Status)

	if obj.TempTablespace != nil {
		result["temp_tablespace"] = string(*obj.TempTablespace)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeExpiring != nil {
		result["time_expiring"] = obj.TimeExpiring.String()
	}

	return result
}
