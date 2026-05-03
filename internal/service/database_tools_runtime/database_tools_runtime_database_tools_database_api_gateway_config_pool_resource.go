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

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext,
		ReadContext:   readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext,
		UpdateContext: updateDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext,
		DeleteContext: deleteDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"database_tools_connection_id": {
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
			"pool_route_value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"advanced_properties": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"database_actions_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"initial_pool_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"jwt_profile_audience": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jwt_profile_issuer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jwt_profile_jwk_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jwt_profile_role_claim_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_pool_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_pool_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rest_enabled_sql_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func createDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res                    *oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPool
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud) ID() string {
	var databaseToolsDatabaseApiGatewayConfigId string
	if v, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		databaseToolsDatabaseApiGatewayConfigId = v.(string)
	}

	poolKey := ""
	if s.Res != nil {
		if key := (*s.Res).GetKey(); key != nil {
			poolKey = *key
		}
	}

	if databaseToolsDatabaseApiGatewayConfigId != "" && poolKey != "" {
		return GetDatabaseToolsDatabaseApiGatewayConfigPoolCompositeId(databaseToolsDatabaseApiGatewayConfigId, poolKey)
	}

	return s.D.Id()
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest{}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.CreateDatabaseToolsDatabaseApiGatewayConfigPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPool
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolRequest{}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if poolKey, ok := s.D.GetOkExists("key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}

	databaseToolsDatabaseApiGatewayConfigId, poolKey, err := parseDatabaseToolsDatabaseApiGatewayConfigPoolCompositeId(s.D.Id())
	if err == nil {
		request.DatabaseToolsDatabaseApiGatewayConfigId = &databaseToolsDatabaseApiGatewayConfigId
		request.PoolKey = &poolKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPool
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolRequest{}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}
	if poolKey, ok := s.D.GetOkExists("key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigPoolRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.UpdateDatabaseToolsDatabaseApiGatewayConfigPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigPool
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolRequest{}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if poolKey, ok := s.D.GetOkExists("key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	_, err := s.Client.DeleteDatabaseToolsDatabaseApiGatewayConfigPool(ctx, request)
	return err
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud) SetData() error {

	databaseToolsDatabaseApiGatewayConfigId, poolKey, err := parseDatabaseToolsDatabaseApiGatewayConfigPoolCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId)
		s.D.Set("key", poolKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolDefault:
		s.D.Set("type", "DEFAULT")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		s.D.Set("database_actions_status", v.DatabaseActionsStatus)

		if v.DatabaseToolsConnectionId != nil {
			s.D.Set("database_tools_connection_id", *v.DatabaseToolsConnectionId)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.InitialPoolSize != nil {
			s.D.Set("initial_pool_size", *v.InitialPoolSize)
		}

		if v.JwtProfileAudience != nil {
			s.D.Set("jwt_profile_audience", *v.JwtProfileAudience)
		}

		if v.JwtProfileIssuer != nil {
			s.D.Set("jwt_profile_issuer", *v.JwtProfileIssuer)
		}

		if v.JwtProfileJwkUrl != nil {
			s.D.Set("jwt_profile_jwk_url", *v.JwtProfileJwkUrl)
		}

		if v.JwtProfileRoleClaimName != nil {
			s.D.Set("jwt_profile_role_claim_name", *v.JwtProfileRoleClaimName)
		}

		if v.MaxPoolSize != nil {
			s.D.Set("max_pool_size", *v.MaxPoolSize)
		}

		if v.MinPoolSize != nil {
			s.D.Set("min_pool_size", *v.MinPoolSize)
		}

		if v.PoolRouteValue != nil {
			s.D.Set("pool_route_value", *v.PoolRouteValue)
		}

		s.D.Set("rest_enabled_sql_status", v.RestEnabledSqlStatus)
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetDatabaseToolsDatabaseApiGatewayConfigPoolCompositeId(databaseToolsDatabaseApiGatewayConfigId string, poolKey string) string {
	databaseToolsDatabaseApiGatewayConfigId = url.PathEscape(databaseToolsDatabaseApiGatewayConfigId)
	poolKey = url.PathEscape(poolKey)
	compositeId := "databaseToolsDatabaseApiGatewayConfigs/" + databaseToolsDatabaseApiGatewayConfigId + "/pools/" + poolKey
	return compositeId
}

func parseDatabaseToolsDatabaseApiGatewayConfigPoolCompositeId(compositeId string) (databaseToolsDatabaseApiGatewayConfigId string, poolKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsDatabaseApiGatewayConfigs/.*/pools/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	databaseToolsDatabaseApiGatewayConfigId, _ = url.PathUnescape(parts[1])
	poolKey, _ = url.PathUnescape(parts[3])

	return
}

func DatabaseToolsDatabaseApiGatewayConfigPoolSummaryToMap(obj oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary:
		result["type"] = "DEFAULT"
		if v.Key != nil {
			result["key"] = *v.Key
		}
		if v.DisplayName != nil {
			result["display_name"] = *v.DisplayName
		}
		if v.DatabaseToolsConnectionId != nil {
			result["database_tools_connection_id"] = *v.DatabaseToolsConnectionId
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest(request *oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) error {
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
		details := oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if databaseActionsStatus, ok := s.D.GetOkExists("database_actions_status"); ok {
			details.DatabaseActionsStatus = oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum(databaseActionsStatus.(string))
		}
		if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
			tmp := databaseToolsConnectionId.(string)
			details.DatabaseToolsConnectionId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if initialPoolSize, ok := s.D.GetOkExists("initial_pool_size"); ok {
			tmp := initialPoolSize.(int)
			details.InitialPoolSize = &tmp
		}
		if jwtProfileAudience, ok := s.D.GetOkExists("jwt_profile_audience"); ok {
			tmp := jwtProfileAudience.(string)
			details.JwtProfileAudience = &tmp
		}
		if jwtProfileIssuer, ok := s.D.GetOkExists("jwt_profile_issuer"); ok {
			tmp := jwtProfileIssuer.(string)
			details.JwtProfileIssuer = &tmp
		}
		if jwtProfileJwkUrl, ok := s.D.GetOkExists("jwt_profile_jwk_url"); ok {
			tmp := jwtProfileJwkUrl.(string)
			details.JwtProfileJwkUrl = &tmp
		}
		if jwtProfileRoleClaimName, ok := s.D.GetOkExists("jwt_profile_role_claim_name"); ok {
			tmp := jwtProfileRoleClaimName.(string)
			details.JwtProfileRoleClaimName = &tmp
		}
		if maxPoolSize, ok := s.D.GetOkExists("max_pool_size"); ok {
			tmp := maxPoolSize.(int)
			details.MaxPoolSize = &tmp
		}
		if minPoolSize, ok := s.D.GetOkExists("min_pool_size"); ok {
			tmp := minPoolSize.(int)
			details.MinPoolSize = &tmp
		}
		if poolRouteValue, ok := s.D.GetOkExists("pool_route_value"); ok {
			tmp := poolRouteValue.(string)
			details.PoolRouteValue = &tmp
		}
		if restEnabledSqlStatus, ok := s.D.GetOkExists("rest_enabled_sql_status"); ok {
			details.RestEnabledSqlStatus = oci_database_tools_runtime.CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum(restEnabledSqlStatus.(string))
		}
		request.CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigPoolRequest(request *oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) error {
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
		details := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDefaultDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if databaseActionsStatus, ok := s.D.GetOkExists("database_actions_status"); ok {
			details.DatabaseActionsStatus = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsDatabaseActionsStatusEnum(databaseActionsStatus.(string))
		}
		if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
			tmp := databaseToolsConnectionId.(string)
			details.DatabaseToolsConnectionId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if initialPoolSize, ok := s.D.GetOkExists("initial_pool_size"); ok {
			tmp := initialPoolSize.(int)
			details.InitialPoolSize = &tmp
		}
		if jwtProfileAudience, ok := s.D.GetOkExists("jwt_profile_audience"); ok {
			tmp := jwtProfileAudience.(string)
			details.JwtProfileAudience = &tmp
		}
		if jwtProfileIssuer, ok := s.D.GetOkExists("jwt_profile_issuer"); ok {
			tmp := jwtProfileIssuer.(string)
			details.JwtProfileIssuer = &tmp
		}
		if jwtProfileJwkUrl, ok := s.D.GetOkExists("jwt_profile_jwk_url"); ok {
			tmp := jwtProfileJwkUrl.(string)
			details.JwtProfileJwkUrl = &tmp
		}
		if jwtProfileRoleClaimName, ok := s.D.GetOkExists("jwt_profile_role_claim_name"); ok {
			tmp := jwtProfileRoleClaimName.(string)
			details.JwtProfileRoleClaimName = &tmp
		}
		if maxPoolSize, ok := s.D.GetOkExists("max_pool_size"); ok {
			tmp := maxPoolSize.(int)
			details.MaxPoolSize = &tmp
		}
		if minPoolSize, ok := s.D.GetOkExists("min_pool_size"); ok {
			tmp := minPoolSize.(int)
			details.MinPoolSize = &tmp
		}
		if poolRouteValue, ok := s.D.GetOkExists("pool_route_value"); ok {
			tmp := poolRouteValue.(string)
			details.PoolRouteValue = &tmp
		}
		if restEnabledSqlStatus, ok := s.D.GetOkExists("rest_enabled_sql_status"); ok {
			details.RestEnabledSqlStatus = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetailsRestEnabledSqlStatusEnum(restEnabledSqlStatus.(string))
		}
		request.UpdateDatabaseToolsDatabaseApiGatewayConfigPoolDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
