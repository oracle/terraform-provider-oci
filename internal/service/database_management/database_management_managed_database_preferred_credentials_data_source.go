// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabasePreferredCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabasePreferredCredentials,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"preferred_credential_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"credential_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_accessible": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"named_credential_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"password_secret_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementManagedDatabasePreferredCredentials(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabasePreferredCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabasePreferredCredentialsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListPreferredCredentialsResponse
}

func (s *DatabaseManagementManagedDatabasePreferredCredentialsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabasePreferredCredentialsDataSourceCrud) Get() error {
	request := oci_database_management.ListPreferredCredentialsRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListPreferredCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabasePreferredCredentialsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabasePreferredCredentialsDataSource-", DatabaseManagementManagedDatabasePreferredCredentialsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabasePreferredCredential := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PreferredCredentialSummaryToMap(item))
	}
	managedDatabasePreferredCredential["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabasePreferredCredentialsDataSource().Schema["preferred_credential_collection"].Elem.(*schema.Resource).Schema)
		managedDatabasePreferredCredential["items"] = items
	}

	resources = append(resources, managedDatabasePreferredCredential)
	if err := s.D.Set("preferred_credential_collection", resources); err != nil {
		return err
	}

	return nil
}

func PreferredCredentialSummaryToMap(obj oci_database_management.PreferredCredentialSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CredentialName != nil {
		result["credential_name"] = string(*obj.CredentialName)
	}

	if obj.IsAccessible != nil {
		result["is_accessible"] = bool(*obj.IsAccessible)
	}

	if obj.NamedCredentialId != nil {
		result["named_credential_id"] = string(*obj.NamedCredentialId)
	}

	if obj.PasswordSecretId != nil {
		result["password_secret_id"] = string(*obj.PasswordSecretId)
	}

	result["role"] = string(obj.Role)

	result["status"] = string(obj.Status)

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	return result
}
