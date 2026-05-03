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

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext,
		ReadContext:   readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext,
		UpdateContext: updateDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext,
		DeleteContext: deleteDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"database_object_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_object_type": {
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
			"alias": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operations": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"roles": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_schemes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

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

func createDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res                    *oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud) ID() string {
	var databaseToolsDatabaseApiGatewayConfigId, poolKey string
	if v, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		databaseToolsDatabaseApiGatewayConfigId = v.(string)
	}
	if v, ok := s.D.GetOkExists("pool_key"); ok {
		poolKey = v.(string)
	}

	autoApiSpecKey := ""
	if s.Res != nil {
		if key := (*s.Res).GetKey(); key != nil {
			autoApiSpecKey = *key
		}
	}

	if databaseToolsDatabaseApiGatewayConfigId != "" && poolKey != "" && autoApiSpecKey != "" {
		return GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecCompositeId(autoApiSpecKey, databaseToolsDatabaseApiGatewayConfigId, poolKey)
	}

	return s.D.Id()
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest{}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}
	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest{}

	if autoApiSpecKey, ok := s.D.GetOkExists("key"); ok {
		tmp := autoApiSpecKey.(string)
		request.AutoApiSpecKey = &tmp
	}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}

	autoApiSpecKey, databaseToolsDatabaseApiGatewayConfigId, poolKey, err := parseDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecCompositeId(s.D.Id())
	if err == nil {
		request.AutoApiSpecKey = &autoApiSpecKey
		request.DatabaseToolsDatabaseApiGatewayConfigId = &databaseToolsDatabaseApiGatewayConfigId
		request.PoolKey = &poolKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest{}
	if autoApiSpecKey, ok := s.D.GetOkExists("key"); ok {
		tmp := autoApiSpecKey.(string)
		request.AutoApiSpecKey = &tmp
	}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}
	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest{}

	if autoApiSpecKey, ok := s.D.GetOkExists("key"); ok {
		tmp := autoApiSpecKey.(string)
		request.AutoApiSpecKey = &tmp
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

	_, err := s.Client.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx, request)
	return err
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud) SetData() error {

	autoApiSpecKey, databaseToolsDatabaseApiGatewayConfigId, poolKey, err := parseDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", autoApiSpecKey)
		s.D.Set("database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId)
		s.D.Set("pool_key", poolKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	v := *s.Res
	s.D.Set("type", "DEFAULT")

	if v.GetAlias() != nil {
		s.D.Set("alias", *v.GetAlias())
	}
	if v.GetDatabaseObjectName() != nil {
		s.D.Set("database_object_name", *v.GetDatabaseObjectName())
	}
	if v.GetDescription() != nil {
		s.D.Set("description", *v.GetDescription())
	}
	if v.GetDisplayName() != nil {
		s.D.Set("display_name", *v.GetDisplayName())
	}
	if key := v.GetKey(); key != nil {
		s.D.Set("key", *key)
	}
	if v.GetScope() != nil {
		s.D.Set("scope", *v.GetScope())
	}
	if timeCreated := v.GetTimeCreated(); timeCreated != nil {
		s.D.Set("time_created", timeCreated.String())
	}
	if timeUpdated := v.GetTimeUpdated(); timeUpdated != nil {
		s.D.Set("time_updated", timeUpdated.String())
	}

	s.D.Set("database_object_type", v.GetDatabaseObjectType())
	s.D.Set("operations", autoApiSpecOperationsToStrings(v.GetOperations()))
	s.D.Set("security_schemes", autoApiSpecSecuritySchemesToStrings(v.GetSecuritySchemes()))
	s.D.Set("roles", v.GetRoles())
	return nil
}

func GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecCompositeId(autoApiSpecKey string, databaseToolsDatabaseApiGatewayConfigId string, poolKey string) string {
	autoApiSpecKey = url.PathEscape(autoApiSpecKey)
	databaseToolsDatabaseApiGatewayConfigId = url.PathEscape(databaseToolsDatabaseApiGatewayConfigId)
	poolKey = url.PathEscape(poolKey)
	compositeId := "databaseToolsDatabaseApiGatewayConfigs/" + databaseToolsDatabaseApiGatewayConfigId + "/pools/" + poolKey + "/autoApiSpecs/" + autoApiSpecKey
	return compositeId
}

func parseDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecCompositeId(compositeId string) (autoApiSpecKey string, databaseToolsDatabaseApiGatewayConfigId string, poolKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsDatabaseApiGatewayConfigs/.*/pools/.*/autoApiSpecs/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	databaseToolsDatabaseApiGatewayConfigId, _ = url.PathUnescape(parts[1])
	poolKey, _ = url.PathUnescape(parts[3])
	autoApiSpecKey, _ = url.PathUnescape(parts[5])

	return
}

func DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSummaryToMap(obj oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch obj.(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultSummary:
		result["type"] = "DEFAULT"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest(request *oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) error {
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
		details := oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails{}
		if alias, ok := s.D.GetOkExists("alias"); ok {
			tmp := alias.(string)
			details.Alias = &tmp
		}
		if databaseObjectName, ok := s.D.GetOkExists("database_object_name"); ok {
			tmp := databaseObjectName.(string)
			details.DatabaseObjectName = &tmp
		}
		if databaseObjectType, ok := s.D.GetOkExists("database_object_type"); ok {
			details.DatabaseObjectType = oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum(databaseObjectType.(string))
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if operations, ok := s.D.GetOkExists("operations"); ok {
			interfaces := operations.([]interface{})
			tmp := make([]oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange("operations") {
				details.Operations = tmp
			}
		}
		if roles, ok := s.D.GetOkExists("roles"); ok {
			interfaces := roles.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("roles") {
				details.Roles = tmp
			}
		}
		if scope, ok := s.D.GetOkExists("scope"); ok {
			tmp := scope.(string)
			details.Scope = &tmp
		}
		if securitySchemes, ok := s.D.GetOkExists("security_schemes"); ok {
			interfaces := securitySchemes.([]interface{})
			tmp := make([]oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange("security_schemes") {
				details.SecuritySchemes = tmp
			}
		}
		request.CreateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest(request *oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) error {
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
		details := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDefaultDetails{}
		if alias, ok := s.D.GetOkExists("alias"); ok {
			tmp := alias.(string)
			details.Alias = &tmp
		}
		if databaseObjectName, ok := s.D.GetOkExists("database_object_name"); ok {
			tmp := databaseObjectName.(string)
			details.DatabaseObjectName = &tmp
		}
		if databaseObjectType, ok := s.D.GetOkExists("database_object_type"); ok {
			details.DatabaseObjectType = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsDatabaseObjectTypeEnum(databaseObjectType.(string))
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if operations, ok := s.D.GetOkExists("operations"); ok {
			interfaces := operations.([]interface{})
			tmp := make([]oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsOperationsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange("operations") {
				details.Operations = tmp
			}
		}
		if roles, ok := s.D.GetOkExists("roles"); ok {
			interfaces := roles.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("roles") {
				details.Roles = tmp
			}
		}
		if scope, ok := s.D.GetOkExists("scope"); ok {
			tmp := scope.(string)
			details.Scope = &tmp
		}
		if securitySchemes, ok := s.D.GetOkExists("security_schemes"); ok {
			interfaces := securitySchemes.([]interface{})
			tmp := make([]oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetailsSecuritySchemesEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange("security_schemes") {
				details.SecuritySchemes = tmp
			}
		}
		request.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func autoApiSpecOperationsToStrings(operations []oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecOperationsEnum) []string {
	result := make([]string, 0, len(operations))
	for _, operation := range operations {
		if operation != "" {
			result = append(result, string(operation))
		}
	}
	return result
}

func autoApiSpecSecuritySchemesToStrings(securitySchemes []oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSecuritySchemesEnum) []string {
	result := make([]string, 0, len(securitySchemes))
	for _, securityScheme := range securitySchemes {
		if securityScheme != "" {
			result = append(result, string(securityScheme))
		}
	}
	return result
}
