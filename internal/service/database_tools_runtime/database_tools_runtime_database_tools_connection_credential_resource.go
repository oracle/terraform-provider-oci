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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsRuntimeDatabaseToolsConnectionCredentialWithContext,
		ReadContext:   readDatabaseToolsRuntimeDatabaseToolsConnectionCredentialWithContext,
		UpdateContext: updateDatabaseToolsRuntimeDatabaseToolsConnectionCredentialWithContext,
		DeleteContext: deleteDatabaseToolsRuntimeDatabaseToolsConnectionCredentialWithContext,
		Schema: map[string]*schema.Schema{
			// Required
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
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BASIC",
				}, true),
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional

			// Computed
			"enabled": {
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
			"windows_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseToolsRuntimeDatabaseToolsConnectionCredentialWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsRuntimeDatabaseToolsConnectionCredentialWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsRuntimeDatabaseToolsConnectionCredentialWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsRuntimeDatabaseToolsConnectionCredentialWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res                    *oci_database_tools_runtime.Credential
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud) ID() string {
	return GetDatabaseToolsConnectionCredentialCompositeId(s.D.Get("key").(string), s.D.Get("database_tools_connection_id").(string))
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.CreateCredentialRequest{}
	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}
	err := s.populateTopLevelPolymorphicCreateCredentialRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.CreateCredential(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Credential
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetCredentialRequest{}

	if credentialKey, ok := s.D.GetOkExists("key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	credentialKey, databaseToolsConnectionId, err := parseDatabaseToolsConnectionCredentialCompositeId(s.D.Id())
	if err == nil {
		request.CredentialKey = &credentialKey
		request.DatabaseToolsConnectionId = &databaseToolsConnectionId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.GetCredential(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Credential
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.UpdateCredentialRequest{}
	if credentialKey, ok := s.D.GetOkExists("key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}
	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}
	err := s.populateTopLevelPolymorphicUpdateCredentialRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.UpdateCredential(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Credential
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.DeleteCredentialRequest{}

	if credentialKey, ok := s.D.GetOkExists("key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	_, err := s.Client.DeleteCredential(ctx, request)
	return err
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud) SetData() error {

	credentialKey, databaseToolsConnectionId, err := parseDatabaseToolsConnectionCredentialCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", credentialKey)
		s.D.Set("database_tools_connection_id", databaseToolsConnectionId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
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

func GetDatabaseToolsConnectionCredentialCompositeId(credentialKey string, databaseToolsConnectionId string) string {
	credentialKey = url.PathEscape(credentialKey)
	databaseToolsConnectionId = url.PathEscape(databaseToolsConnectionId)
	compositeId := "databaseToolsConnections/" + databaseToolsConnectionId + "/credentials/" + credentialKey
	return compositeId
}

func parseDatabaseToolsConnectionCredentialCompositeId(compositeId string) (credentialKey string, databaseToolsConnectionId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsConnections/.*/credentials/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	databaseToolsConnectionId, _ = url.PathUnescape(parts[1])
	credentialKey, _ = url.PathUnescape(parts[3])

	return
}

func CredentialRelatedResourceToMap(obj *oci_database_tools_runtime.CredentialRelatedResource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	result["type"] = string(obj.Type)

	return result
}

func CredentialSummaryToMap(obj oci_database_tools_runtime.CredentialSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Enabled != nil {
		result["enabled"] = string(*obj.Enabled)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["key_type"] = string(obj.KeyType)

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	if obj.RelatedResource != nil {
		result["related_resource"] = []interface{}{CredentialRelatedResourceToMap(obj.RelatedResource)}
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	if obj.WindowsDomain != nil {
		result["windows_domain"] = string(*obj.WindowsDomain)
	}

	return result
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud) populateTopLevelPolymorphicCreateCredentialRequest(request *oci_database_tools_runtime.CreateCredentialRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("BASIC"):
		details := oci_database_tools_runtime.CreateCredentialBasicDetails{}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if userName, ok := s.D.GetOkExists("user_name"); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		if key, ok := s.D.GetOkExists("key"); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		request.CreateCredentialDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceCrud) populateTopLevelPolymorphicUpdateCredentialRequest(request *oci_database_tools_runtime.UpdateCredentialRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("BASIC"):
		details := oci_database_tools_runtime.UpdateCredentialBasicDetails{}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if userName, ok := s.D.GetOkExists("user_name"); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		request.UpdateCredentialDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
