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

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext,
		ReadContext:   readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext,
		UpdateContext: updateDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext,
		DeleteContext: deleteDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_tools_database_api_gateway_config_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pool_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional

			// Computed
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res                    *oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud) ID() string {
	var databaseToolsDatabaseApiGatewayConfigId, poolKey string
	if v, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		databaseToolsDatabaseApiGatewayConfigId = v.(string)
	}
	if v, ok := s.D.GetOkExists("pool_key"); ok {
		poolKey = v.(string)
	}

	apiSpecKey := ""
	if s.Res != nil {
		if key := (*s.Res).GetKey(); key != nil {
			apiSpecKey = *key
		}
	}

	if databaseToolsDatabaseApiGatewayConfigId != "" && poolKey != "" && apiSpecKey != "" {
		return GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCompositeId(apiSpecKey, databaseToolsDatabaseApiGatewayConfigId, poolKey)
	}

	return s.D.Id()
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest{}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}
	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest{}

	if apiSpecKey, ok := s.D.GetOkExists("key"); ok {
		tmp := apiSpecKey.(string)
		request.ApiSpecKey = &tmp
	}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}

	apiSpecKey, databaseToolsDatabaseApiGatewayConfigId, poolKey, err := parseDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCompositeId(s.D.Id())
	if err == nil {
		request.ApiSpecKey = &apiSpecKey
		request.DatabaseToolsDatabaseApiGatewayConfigId = &databaseToolsDatabaseApiGatewayConfigId
		request.PoolKey = &poolKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest{}
	if apiSpecKey, ok := s.D.GetOkExists("key"); ok {
		tmp := apiSpecKey.(string)
		request.ApiSpecKey = &tmp
	}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}
	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest{}

	if apiSpecKey, ok := s.D.GetOkExists("key"); ok {
		tmp := apiSpecKey.(string)
		request.ApiSpecKey = &tmp
	}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	_, err := s.Client.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx, request)
	return err
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud) SetData() error {

	apiSpecKey, databaseToolsDatabaseApiGatewayConfigId, poolKey, err := parseDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", apiSpecKey)
		s.D.Set("database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId)
		s.D.Set("pool_key", poolKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefault:
		s.D.Set("type", "DEFAULT")

		if v.Content != nil {
			s.D.Set("content", *v.Content)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if key := v.GetKey(); key != nil {
			s.D.Set("key", *key)
		}

		if timeCreated := v.GetTimeCreated(); timeCreated != nil {
			s.D.Set("time_created", timeCreated.String())
		}

		if timeUpdated := v.GetTimeUpdated(); timeUpdated != nil {
			s.D.Set("time_updated", timeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCompositeId(apiSpecKey string, databaseToolsDatabaseApiGatewayConfigId string, poolKey string) string {
	apiSpecKey = url.PathEscape(apiSpecKey)
	databaseToolsDatabaseApiGatewayConfigId = url.PathEscape(databaseToolsDatabaseApiGatewayConfigId)
	poolKey = url.PathEscape(poolKey)
	compositeId := "databaseToolsDatabaseApiGatewayConfigs/" + databaseToolsDatabaseApiGatewayConfigId + "/pools/" + poolKey + "/apiSpecs/" + apiSpecKey
	return compositeId
}

func parseDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCompositeId(compositeId string) (apiSpecKey string, databaseToolsDatabaseApiGatewayConfigId string, poolKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsDatabaseApiGatewayConfigs/.*/pools/.*/apiSpecs/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	databaseToolsDatabaseApiGatewayConfigId, _ = url.PathUnescape(parts[1])
	poolKey, _ = url.PathUnescape(parts[3])
	apiSpecKey, _ = url.PathUnescape(parts[5])

	return
}

func DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecSummaryToMap(obj oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultSummary:
		result["type"] = "DEFAULT"
		if key := v.GetKey(); key != nil {
			result["key"] = *key
		}
		if displayName := v.GetDisplayName(); displayName != nil {
			result["display_name"] = *displayName
		}
		if timeCreated := v.GetTimeCreated(); timeCreated != nil {
			result["time_created"] = timeCreated.String()
		}
		if timeUpdated := v.GetTimeUpdated(); timeUpdated != nil {
			result["time_updated"] = timeUpdated.String()
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest(request *oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails{}
		if content, ok := s.D.GetOkExists("content"); ok {
			tmp := content.(string)
			details.Content = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		request.CreateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest(request *oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefaultDetails{}
		if content, ok := s.D.GetOkExists("content"); ok {
			tmp := content.(string)
			details.Content = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		request.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
