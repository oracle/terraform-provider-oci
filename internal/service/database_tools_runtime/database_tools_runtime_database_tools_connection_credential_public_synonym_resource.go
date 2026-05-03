// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools_runtime

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymWithContext,
		ReadContext:   readDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymWithContext,
		DeleteContext: deleteDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"credential_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"database_tools_connection_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func deleteDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res                    *oci_database_tools_runtime.CredentialPublicSynonym
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud) ID() string {
	return GetDatabaseToolsConnectionCredentialPublicSynonymCompositeId(s.D.Get("credential_key").(string), s.D.Get("database_tools_connection_id").(string), s.D.Get("key").(string))
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.CreateCredentialPublicSynonymRequest{}

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.CreateCredentialPublicSynonym(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CredentialPublicSynonym
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetCredentialPublicSynonymRequest{}

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if publicSynonymKey, ok := s.D.GetOkExists("key"); ok {
		tmp := publicSynonymKey.(string)
		request.PublicSynonymKey = &tmp
	}

	credentialKey, databaseToolsConnectionId, publicSynonymKey, err := parseDatabaseToolsConnectionCredentialPublicSynonymCompositeId(s.D.Id())
	if err == nil {
		request.CredentialKey = &credentialKey
		request.DatabaseToolsConnectionId = &databaseToolsConnectionId
		request.PublicSynonymKey = &publicSynonymKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.GetCredentialPublicSynonym(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CredentialPublicSynonym
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.DeleteCredentialPublicSynonymRequest{}

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if publicSynonymKey, ok := s.D.GetOkExists("key"); ok {
		tmp := publicSynonymKey.(string)
		request.PublicSynonymKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	_, err := s.Client.DeleteCredentialPublicSynonym(ctx, request)
	return err
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceCrud) SetData() error {

	credentialKey, databaseToolsConnectionId, publicSynonymKey, err := parseDatabaseToolsConnectionCredentialPublicSynonymCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("credential_key", credentialKey)
		s.D.Set("database_tools_connection_id", databaseToolsConnectionId)
		s.D.Set("key", publicSynonymKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	return nil
}

func GetDatabaseToolsConnectionCredentialPublicSynonymCompositeId(credentialKey string, databaseToolsConnectionId string, publicSynonymKey string) string {
	credentialKey = url.PathEscape(credentialKey)
	databaseToolsConnectionId = url.PathEscape(databaseToolsConnectionId)
	publicSynonymKey = url.PathEscape(publicSynonymKey)
	compositeId := "databaseToolsConnections/" + databaseToolsConnectionId + "/credentials/" + credentialKey + "/publicSynonyms/" + publicSynonymKey
	return compositeId
}

func parseDatabaseToolsConnectionCredentialPublicSynonymCompositeId(compositeId string) (credentialKey string, databaseToolsConnectionId string, publicSynonymKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsConnections/.*/credentials/.*/publicSynonyms/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	databaseToolsConnectionId, _ = url.PathUnescape(parts[1])
	credentialKey, _ = url.PathUnescape(parts[3])
	publicSynonymKey, _ = url.PathUnescape(parts[5])

	return
}

func CredentialPublicSynonymSummaryToMap(obj oci_database_tools_runtime.CredentialPublicSynonymSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	return result
}
