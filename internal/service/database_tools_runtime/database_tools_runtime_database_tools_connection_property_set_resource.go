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

func DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext,
		ReadContext:   readDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext,
		UpdateContext: updateDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext,
		DeleteContext: deleteDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"database_tools_connection_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"APEX",
					"APEX_DOCUMENT_GENERATOR",
					"APEX_FA_INTEGRATION",
					"ORACLE_DATABASE_EXTERNAL_AUTHENTICATION",
				}, true),
			},
			"property_set_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"authentication_substitutions": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"autonomous_database_resource_principal_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"credential_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"function_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identity_provider": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AZURE_AD",
								"NONE",
								"OCI_IAM",
							}, true),
						},

						// Optional
						"configs": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},

						// Computed
					},
				},
			},
			"instance_dbms_credential_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invoke_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_storage_bucket_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_storage_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_storage_namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"print_server_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"is_mutable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"prerequisites_check": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"user_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res                    *oci_database_tools_runtime.PropertySet
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) ID() string {
	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		if propertySetKey, ok := s.D.GetOkExists("property_set_key"); ok {
			return GetDatabaseToolsConnectionPropertySetCompositeId(databaseToolsConnectionId.(string), propertySetKey.(string))
		}
	}
	return s.D.Id()
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.UpdatePropertySetRequest{}
	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}
	if propertySetKey, ok := s.D.GetOkExists("property_set_key"); ok {
		request.PropertySetKey = oci_database_tools_runtime.UpdatePropertySetPropertySetKeyEnum(propertySetKey.(string))
	}
	err := s.populateTopLevelPolymorphicUpdatePropertySetRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.UpdatePropertySet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.PropertySet
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetPropertySetRequest{}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if propertySetKey, ok := s.D.GetOkExists("key"); ok {
		request.PropertySetKey = oci_database_tools_runtime.GetPropertySetPropertySetKeyEnum(propertySetKey.(string))
	}

	databaseToolsConnectionId, propertySetKey, err := parseDatabaseToolsConnectionPropertySetCompositeId(s.D.Id())
	if err == nil {
		request.DatabaseToolsConnectionId = &databaseToolsConnectionId
		request.PropertySetKey = oci_database_tools_runtime.GetPropertySetPropertySetKeyEnum(propertySetKey)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.GetPropertySet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.PropertySet
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.UpdatePropertySetRequest{}
	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}
	if propertySetKey, ok := s.D.GetOkExists("property_set_key"); ok {
		request.PropertySetKey = oci_database_tools_runtime.UpdatePropertySetPropertySetKeyEnum(propertySetKey.(string))
	}
	err := s.populateTopLevelPolymorphicUpdatePropertySetRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.UpdatePropertySet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.PropertySet
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) SetData() error {

	databaseToolsConnectionId, propertySetKey, err := parseDatabaseToolsConnectionPropertySetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("database_tools_connection_id", databaseToolsConnectionId)
		s.D.Set("property_set_key", propertySetKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_database_tools_runtime.PropertySetApex:
		s.D.Set("key", "APEX")
		s.clearNonApexPropertySetFields()

		if v.UserKey != nil {
			s.D.Set("user_key", *v.UserKey)
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}

		if v.IsMutable != nil {
			s.D.Set("is_mutable", *v.IsMutable)
		}
	case oci_database_tools_runtime.PropertySetApexDocumentGenerator:
		s.D.Set("key", "APEX_DOCUMENT_GENERATOR")
		s.clearNonApexDocumentGeneratorPropertySetFields()

		s.D.Set("autonomous_database_resource_principal_status", v.AutonomousDatabaseResourcePrincipalStatus)

		if v.CredentialKey != nil {
			s.D.Set("credential_key", *v.CredentialKey)
		}

		if v.FunctionId != nil {
			s.D.Set("function_id", *v.FunctionId)
		}

		if v.InvokeEndpoint != nil {
			s.D.Set("invoke_endpoint", *v.InvokeEndpoint)
		}

		if v.ObjectStorageBucketCompartmentId != nil {
			s.D.Set("object_storage_bucket_compartment_id", *v.ObjectStorageBucketCompartmentId)
		}

		if v.ObjectStorageEndpoint != nil {
			s.D.Set("object_storage_endpoint", *v.ObjectStorageEndpoint)
		}

		if v.ObjectStorageNamespace != nil {
			s.D.Set("object_storage_namespace", *v.ObjectStorageNamespace)
		}

		s.D.Set("print_server_type", v.PrintServerType)

		if v.IsMutable != nil {
			s.D.Set("is_mutable", *v.IsMutable)
		}
	case oci_database_tools_runtime.PropertySetApexFaIntegration:
		s.D.Set("key", "APEX_FA_INTEGRATION")
		s.clearNonApexFaIntegrationPropertySetFields()

		s.D.Set("authentication_substitutions", v.AuthenticationSubstitutions)

		if v.InstanceDbmsCredentialEnabled != nil {
			s.D.Set("instance_dbms_credential_enabled", *v.InstanceDbmsCredentialEnabled)
		}

		if v.PrerequisitesCheck != nil {
			s.D.Set("prerequisites_check", []interface{}{ApexFaIntegrationPrerequisitesCheckToMap(v.PrerequisitesCheck)})
		} else {
			s.D.Set("prerequisites_check", nil)
		}

		if v.IsMutable != nil {
			s.D.Set("is_mutable", *v.IsMutable)
		}
	case oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthentication:
		s.D.Set("key", "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION")
		s.clearNonOracleDatabaseExternalAuthenticationPropertySetFields()

		if v.IdentityProvider != nil {
			identityProviderArray := []interface{}{}
			if identityProviderMap := PropertySetOracleDatabaseExternalAuthenticationIdentityProviderToMap(&v.IdentityProvider); identityProviderMap != nil {
				identityProviderArray = append(identityProviderArray, identityProviderMap)
			}
			s.D.Set("identity_provider", identityProviderArray)
		} else {
			s.D.Set("identity_provider", nil)
		}

		if v.IsMutable != nil {
			s.D.Set("is_mutable", *v.IsMutable)
		}
	default:
		log.Printf("[WARN] Received 'key' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) clearNonApexPropertySetFields() {
	s.D.Set("authentication_substitutions", nil)
	s.D.Set("autonomous_database_resource_principal_status", nil)
	s.D.Set("credential_key", nil)
	s.D.Set("function_id", nil)
	s.D.Set("identity_provider", nil)
	s.D.Set("instance_dbms_credential_enabled", nil)
	s.D.Set("invoke_endpoint", nil)
	s.D.Set("object_storage_bucket_compartment_id", nil)
	s.D.Set("object_storage_endpoint", nil)
	s.D.Set("object_storage_namespace", nil)
	s.D.Set("prerequisites_check", nil)
	s.D.Set("print_server_type", nil)
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) clearNonApexDocumentGeneratorPropertySetFields() {
	s.D.Set("authentication_substitutions", nil)
	s.D.Set("identity_provider", nil)
	s.D.Set("instance_dbms_credential_enabled", nil)
	s.D.Set("prerequisites_check", nil)
	s.D.Set("user_key", nil)
	s.D.Set("version", nil)
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) clearNonApexFaIntegrationPropertySetFields() {
	s.D.Set("autonomous_database_resource_principal_status", nil)
	s.D.Set("credential_key", nil)
	s.D.Set("function_id", nil)
	s.D.Set("identity_provider", nil)
	s.D.Set("invoke_endpoint", nil)
	s.D.Set("object_storage_bucket_compartment_id", nil)
	s.D.Set("object_storage_endpoint", nil)
	s.D.Set("object_storage_namespace", nil)
	s.D.Set("print_server_type", nil)
	s.D.Set("user_key", nil)
	s.D.Set("version", nil)
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) clearNonOracleDatabaseExternalAuthenticationPropertySetFields() {
	s.D.Set("authentication_substitutions", nil)
	s.D.Set("autonomous_database_resource_principal_status", nil)
	s.D.Set("credential_key", nil)
	s.D.Set("function_id", nil)
	s.D.Set("instance_dbms_credential_enabled", nil)
	s.D.Set("invoke_endpoint", nil)
	s.D.Set("object_storage_bucket_compartment_id", nil)
	s.D.Set("object_storage_endpoint", nil)
	s.D.Set("object_storage_namespace", nil)
	s.D.Set("prerequisites_check", nil)
	s.D.Set("print_server_type", nil)
	s.D.Set("user_key", nil)
	s.D.Set("version", nil)
}

func GetDatabaseToolsConnectionPropertySetCompositeId(databaseToolsConnectionId string, propertySetKey string) string {
	databaseToolsConnectionId = url.PathEscape(databaseToolsConnectionId)
	propertySetKey = url.PathEscape(propertySetKey)
	compositeId := "databaseToolsConnections/" + databaseToolsConnectionId + "/propertySets/" + propertySetKey
	return compositeId
}

func parseDatabaseToolsConnectionPropertySetCompositeId(compositeId string) (databaseToolsConnectionId string, propertySetKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsConnections/.*/propertySets/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	databaseToolsConnectionId, _ = url.PathUnescape(parts[1])
	propertySetKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) mapToApexFaIntegrationPrerequisitesCheck(fieldKeyFormat string) (oci_database_tools_runtime.ApexFaIntegrationPrerequisitesCheck, error) {
	result := oci_database_tools_runtime.ApexFaIntegrationPrerequisitesCheck{}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_database_tools_runtime.ApexFaIntegrationPrerequisitesCheckStatusEnum(status.(string))
	}

	if statusDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status_details")); ok {
		interfaces := statusDetails.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "status_details")) {
			result.StatusDetails = tmp
		}
	}

	return result, nil
}

func ApexFaIntegrationPrerequisitesCheckToMap(obj *oci_database_tools_runtime.ApexFaIntegrationPrerequisitesCheck) map[string]interface{} {
	result := map[string]interface{}{}

	result["status"] = string(obj.Status)

	result["status_details"] = obj.StatusDetails

	return result
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) mapToPropertySetOracleDatabaseExternalAuthenticationIdentityProvider(fieldKeyFormat string) (oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProvider, error) {
	var baseObject oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProvider
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("AZURE_AD"):
		details := oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd{}
		if configs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configs")); ok {
			details.Configs = tfresource.ObjectMapToStringMap(configs.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProviderNone{}
		baseObject = details
	case strings.ToLower("OCI_IAM"):
		details := oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProviderOciIam{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func PropertySetOracleDatabaseExternalAuthenticationIdentityProviderToMap(obj *oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProvider) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAd:
		result["type"] = "AZURE_AD"
		result["configs"] = normalizePropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAdConfigs(v.Configs)
	case oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProviderNone:
		result["type"] = "NONE"
	case oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthenticationIdentityProviderOciIam:
		result["type"] = "OCI_IAM"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func normalizePropertySetOracleDatabaseExternalAuthenticationIdentityProviderAzureAdConfigs(configs map[string]string) map[string]string {
	if configs == nil {
		return nil
	}

	normalized := map[string]string{}
	for k, v := range configs {
		normalized[k] = v
	}

	if appID, ok := normalized["app_id"]; ok {
		delete(normalized, "app_id")
		if _, exists := normalized["application_id"]; !exists {
			normalized["application_id"] = appID
		}
	}

	return normalized
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceCrud) populateTopLevelPolymorphicUpdatePropertySetRequest(request *oci_database_tools_runtime.UpdatePropertySetRequest) error {
	//discriminator
	keyRaw, ok := s.D.GetOkExists("key")
	var key string
	if ok {
		key = keyRaw.(string)
	} else {
		key = "" // default value
	}
	switch strings.ToLower(key) {
	case strings.ToLower("APEX_DOCUMENT_GENERATOR"):
		details := oci_database_tools_runtime.UpdatePropertySetApexDocumentGeneratorDetails{}
		if autonomousDatabaseResourcePrincipalStatus, ok := s.D.GetOkExists("autonomous_database_resource_principal_status"); ok {
			details.AutonomousDatabaseResourcePrincipalStatus = oci_database_tools_runtime.PropertySetApexDocumentGeneratorAutonomousDatabaseResourcePrincipalStatusEnum(autonomousDatabaseResourcePrincipalStatus.(string))
		}
		if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
			tmp := credentialKey.(string)
			details.CredentialKey = &tmp
		}
		if functionId, ok := s.D.GetOkExists("function_id"); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
		}
		if invokeEndpoint, ok := s.D.GetOkExists("invoke_endpoint"); ok {
			tmp := invokeEndpoint.(string)
			details.InvokeEndpoint = &tmp
		}
		if objectStorageBucketCompartmentId, ok := s.D.GetOkExists("object_storage_bucket_compartment_id"); ok {
			tmp := objectStorageBucketCompartmentId.(string)
			details.ObjectStorageBucketCompartmentId = &tmp
		}
		if objectStorageEndpoint, ok := s.D.GetOkExists("object_storage_endpoint"); ok {
			tmp := objectStorageEndpoint.(string)
			details.ObjectStorageEndpoint = &tmp
		}
		if objectStorageNamespace, ok := s.D.GetOkExists("object_storage_namespace"); ok {
			tmp := objectStorageNamespace.(string)
			details.ObjectStorageNamespace = &tmp
		}
		if printServerType, ok := s.D.GetOkExists("print_server_type"); ok {
			details.PrintServerType = oci_database_tools_runtime.PropertySetApexDocumentGeneratorPrintServerTypeEnum(printServerType.(string))
		}
		request.UpdatePropertySetDetails = details
	case strings.ToLower("APEX_FA_INTEGRATION"):
		details := oci_database_tools_runtime.UpdatePropertySetApexFaIntegrationDetails{}
		if authenticationSubstitutions, ok := s.D.GetOkExists("authentication_substitutions"); ok {
			details.AuthenticationSubstitutions = tfresource.ObjectMapToStringMap(authenticationSubstitutions.(map[string]interface{}))
		}
		if instanceDbmsCredentialEnabled, ok := s.D.GetOkExists("instance_dbms_credential_enabled"); ok {
			tmp := instanceDbmsCredentialEnabled.(string)
			details.InstanceDbmsCredentialEnabled = &tmp
		}
		request.UpdatePropertySetDetails = details
	case strings.ToLower("ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"):
		details := oci_database_tools_runtime.UpdatePropertySetOracleDatabaseExternalAuthenticationDetails{}
		if identityProvider, ok := s.D.GetOkExists("identity_provider"); ok {
			if tmpList := identityProvider.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "identity_provider", 0)
				tmp, err := s.mapToPropertySetOracleDatabaseExternalAuthenticationIdentityProvider(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.IdentityProvider = tmp
			}
		}
		request.UpdatePropertySetDetails = details
	default:
		return fmt.Errorf("unknown key '%v' was specified", key)
	}
	return nil
}
