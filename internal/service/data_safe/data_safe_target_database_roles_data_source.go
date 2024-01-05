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

func DataSafeTargetDatabaseRolesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeTargetDatabaseRoles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"authentication_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_oracle_maintained": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"role_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"role_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"authentication_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_common": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_implicit": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_inherited": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_oracle_maintained": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_password_required": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"role_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeTargetDatabaseRoles(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseRolesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabaseRolesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListRolesResponse
}

func (s *DataSafeTargetDatabaseRolesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabaseRolesDataSourceCrud) Get() error {
	request := oci_data_safe.ListRolesRequest{}

	if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
		tmp := authenticationType.(string)
		request.AuthenticationType = &tmp
	}

	if isOracleMaintained, ok := s.D.GetOkExists("is_oracle_maintained"); ok {
		tmp := isOracleMaintained.(bool)
		request.IsOracleMaintained = &tmp
	}

	if roleName, ok := s.D.GetOkExists("role_name"); ok {
		request.RoleName = roleName.([]string)
	}

	if roleNameContains, ok := s.D.GetOkExists("role_name_contains"); ok {
		tmp := roleNameContains.(string)
		request.RoleNameContains = &tmp
	}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListRoles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRoles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeTargetDatabaseRolesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetDatabaseRolesDataSource-", DataSafeTargetDatabaseRolesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		targetDatabaseRole := map[string]interface{}{}

		if r.AuthenticationType != nil {
			targetDatabaseRole["authentication_type"] = *r.AuthenticationType
		}

		if r.IsCommon != nil {
			targetDatabaseRole["is_common"] = *r.IsCommon
		}

		if r.IsImplicit != nil {
			targetDatabaseRole["is_implicit"] = *r.IsImplicit
		}

		if r.IsInherited != nil {
			targetDatabaseRole["is_inherited"] = *r.IsInherited
		}

		if r.IsOracleMaintained != nil {
			targetDatabaseRole["is_oracle_maintained"] = *r.IsOracleMaintained
		}

		if r.IsPasswordRequired != nil {
			targetDatabaseRole["is_password_required"] = *r.IsPasswordRequired
		}

		if r.RoleName != nil {
			targetDatabaseRole["role_name"] = *r.RoleName
		}

		resources = append(resources, targetDatabaseRole)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeTargetDatabaseRolesDataSource().Schema["roles"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("roles", resources); err != nil {
		return err
	}

	return nil
}
