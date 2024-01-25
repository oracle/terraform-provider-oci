// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabasePreferredCredentialDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabasePreferredCredential,
		Schema: map[string]*schema.Schema{
			"credential_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"is_accessible": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"named_credential_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
	}
}

func readSingularDatabaseManagementManagedDatabasePreferredCredential(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabasePreferredCredentialDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabasePreferredCredentialDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetPreferredCredentialResponse
}

func (s *DatabaseManagementManagedDatabasePreferredCredentialDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabasePreferredCredentialDataSourceCrud) Get() error {
	request := oci_database_management.GetPreferredCredentialRequest{}

	if credentialName, ok := s.D.GetOkExists("credential_name"); ok {
		tmp := credentialName.(string)
		request.CredentialName = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetPreferredCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabasePreferredCredentialDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabasePreferredCredentialDataSource-", DatabaseManagementManagedDatabasePreferredCredentialDataSource(), s.D))

	if s.Res.GetIsAccessible() != nil {
		s.D.Set("is_accessible", *s.Res.GetIsAccessible())
	}

	switch v := (s.Res.PreferredCredential).(type) {
	case oci_database_management.BasicPreferredCredential:
		s.D.Set("type", "BASIC")

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		s.D.Set("role", v.Role)

		if v.UserName != nil {
			s.D.Set("user_name", *v.UserName)
		}

		if v.CredentialName != nil {
			s.D.Set("credential_name", *v.CredentialName)
		}

		if v.IsAccessible != nil {
			s.D.Set("is_accessible", *v.IsAccessible)
		}

		s.D.Set("status", v.Status)
	case oci_database_management.NamedPreferredCredential:
		s.D.Set("type", "NAMED_CREDENTIAL")

		if v.NamedCredentialId != nil {
			s.D.Set("named_credential_id", *v.NamedCredentialId)
		}

		if v.CredentialName != nil {
			s.D.Set("credential_name", *v.CredentialName)
		}

		if v.IsAccessible != nil {
			s.D.Set("is_accessible", *v.IsAccessible)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}

	s.D.Set("status", s.Res.GetStatus())

	return nil
}
