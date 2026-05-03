// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools_runtime

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialWithContext,
		Schema: map[string]*schema.Schema{
			"credential_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_tools_connection_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"related_resource": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"windows_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetUserCredentialResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetUserCredentialRequest{}

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if userKey, ok := s.D.GetOkExists("user_key"); ok {
		tmp := userKey.(string)
		request.UserKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.GetUserCredential(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSource-", DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSource(), s.D))

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		s.D.Set("credential_key", credentialKey.(string))
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		s.D.Set("database_tools_connection_id", databaseToolsConnectionId.(string))
	}

	if userKey, ok := s.D.GetOkExists("user_key"); ok {
		s.D.Set("user_key", userKey.(string))
	}

	if s.Res.Enabled != nil {
		s.D.Set("enabled", *s.Res.Enabled)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("key_type", s.Res.KeyType)

	if s.Res.Owner != nil {
		s.D.Set("owner", *s.Res.Owner)
	}

	if s.Res.RelatedResource != nil {
		s.D.Set("related_resource", []interface{}{CredentialRelatedResourceToMap(s.Res.RelatedResource)})
	} else {
		s.D.Set("related_resource", nil)
	}

	if s.Res.UserName != nil {
		s.D.Set("user_name", *s.Res.UserName)
	}

	if s.Res.WindowsDomain != nil {
		s.D.Set("windows_domain", *s.Res.WindowsDomain)
	}

	return nil
}
