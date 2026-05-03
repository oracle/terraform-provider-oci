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

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext,
		ReadContext:   readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext,
		UpdateContext: updateDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext,
		DeleteContext: deleteDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"database_tools_database_api_gateway_config_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"global_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"certificate_bundle": {
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
								"FILENAME",
								"SELF_SIGNED",
							}, true),
						},

						// Optional
						"certificate_private_key": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"format": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"certificate_public": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"format": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"database_api_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"document_root": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"https_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pool_route": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_routing_header": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_source": {
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

func createDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res                    *oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigGlobal
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) ID() string {
	var databaseToolsDatabaseApiGatewayConfigId string
	if v, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		databaseToolsDatabaseApiGatewayConfigId = v.(string)
	}

	globalKey := ""
	if s.Res != nil {
		if key := (*s.Res).GetKey(); key != nil {
			globalKey = *key
		}
	}

	if databaseToolsDatabaseApiGatewayConfigId != "" && globalKey != "" {
		return GetDatabaseToolsDatabaseApiGatewayConfigGlobalCompositeId(databaseToolsDatabaseApiGatewayConfigId, globalKey)
	}

	return s.D.Id()
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest{}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}
	if globalKey, ok := s.D.GetOkExists("global_key"); ok {
		request.GlobalKey = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum(globalKey.(string))
	}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigGlobal
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigGlobalRequest{}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if globalKey, ok := s.D.GetOkExists("global_key"); ok {
		request.GlobalKey = oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum(globalKey.(string))
	}

	databaseToolsDatabaseApiGatewayConfigId, globalKey, err := parseDatabaseToolsDatabaseApiGatewayConfigGlobalCompositeId(s.D.Id())
	if err == nil {
		request.DatabaseToolsDatabaseApiGatewayConfigId = &databaseToolsDatabaseApiGatewayConfigId
		request.GlobalKey = oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum(globalKey)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigGlobal(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigGlobal
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest{}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}
	if globalKey, ok := s.D.GetOkExists("global_key"); ok {
		request.GlobalKey = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum(globalKey.(string))
	}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools_runtime")

	response, err := s.Client.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfigGlobal
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) SetData() error {

	databaseToolsDatabaseApiGatewayConfigId, globalKey, err := parseDatabaseToolsDatabaseApiGatewayConfigGlobalCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId)
		s.D.Set("global_key", globalKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if key := (*s.Res).GetKey(); key != nil {
		s.D.Set("key", *key)
	}

	s.D.Set("metadata_source", (*s.Res).GetMetadataSource())

	if timeCreated := (*s.Res).GetTimeCreated(); timeCreated != nil {
		s.D.Set("time_created", timeCreated.String())
	}

	if timeUpdated := (*s.Res).GetTimeUpdated(); timeUpdated != nil {
		s.D.Set("time_updated", timeUpdated.String())
	}

	switch v := (*s.Res).(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigGlobalDefault:
		s.D.Set("type", "DEFAULT")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		if v.CertificateBundle != nil {
			certificateBundleArray := []interface{}{}
			if certificateBundleMap := DatabaseApiGatewayConfigCertificateBundleToMap(&v.CertificateBundle); certificateBundleMap != nil {
				certificateBundleArray = append(certificateBundleArray, certificateBundleMap)
			}
			s.D.Set("certificate_bundle", certificateBundleArray)
		} else {
			s.D.Set("certificate_bundle", nil)
		}

		s.D.Set("database_api_status", v.DatabaseApiStatus)

		if v.DocumentRoot != nil {
			s.D.Set("document_root", *v.DocumentRoot)
		}

		if v.HttpPort != nil {
			s.D.Set("http_port", *v.HttpPort)
		}

		if v.HttpsPort != nil {
			s.D.Set("https_port", *v.HttpsPort)
		}

		s.D.Set("pool_route", v.PoolRoute)

		if v.PoolRoutingHeader != nil {
			s.D.Set("pool_routing_header", *v.PoolRoutingHeader)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetDatabaseToolsDatabaseApiGatewayConfigGlobalCompositeId(databaseToolsDatabaseApiGatewayConfigId string, globalKey string) string {
	databaseToolsDatabaseApiGatewayConfigId = url.PathEscape(databaseToolsDatabaseApiGatewayConfigId)
	globalKey = url.PathEscape(globalKey)
	compositeId := "databaseToolsDatabaseApiGatewayConfigs/" + databaseToolsDatabaseApiGatewayConfigId + "/globals/" + globalKey
	return compositeId
}

func parseDatabaseToolsDatabaseApiGatewayConfigGlobalCompositeId(compositeId string) (databaseToolsDatabaseApiGatewayConfigId string, globalKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsDatabaseApiGatewayConfigs/.*/globals/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	databaseToolsDatabaseApiGatewayConfigId, _ = url.PathUnescape(parts[1])
	globalKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) mapToDatabaseApiGatewayConfigCertificateBundle(fieldKeyFormat string) (oci_database_tools_runtime.DatabaseApiGatewayConfigCertificateBundle, error) {
	var baseObject oci_database_tools_runtime.DatabaseApiGatewayConfigCertificateBundle
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("FILENAME"):
		details := oci_database_tools_runtime.DatabaseApiGatewayConfigCertificateBundleFileName{}
		if certificatePrivateKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_private_key")); ok {
			if tmpList := certificatePrivateKey.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "certificate_private_key"), 0)
				tmp, err := s.mapToDatabaseApiGatewayConfigCertificatePrivateKeyFileName(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert certificate_private_key, encountered error: %v", err)
				}
				details.CertificatePrivateKey = &tmp
			}
		}
		if certificatePublic, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_public")); ok {
			if tmpList := certificatePublic.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "certificate_public"), 0)
				tmp, err := s.mapToDatabaseApiGatewayConfigCertificatePublicFileName(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert certificate_public, encountered error: %v", err)
				}
				details.CertificatePublic = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("SELF_SIGNED"):
		details := oci_database_tools_runtime.DatabaseApiGatewayConfigCertificateBundleSelfSigned{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func DatabaseApiGatewayConfigCertificateBundleToMap(obj *oci_database_tools_runtime.DatabaseApiGatewayConfigCertificateBundle) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools_runtime.DatabaseApiGatewayConfigCertificateBundleFileName:
		result["type"] = "FILENAME"

		if v.CertificatePrivateKey != nil {
			result["certificate_private_key"] = []interface{}{DatabaseApiGatewayConfigCertificatePrivateKeyFileNameToMap(v.CertificatePrivateKey)}
		}

		if v.CertificatePublic != nil {
			result["certificate_public"] = []interface{}{DatabaseApiGatewayConfigCertificatePublicFileNameToMap(v.CertificatePublic)}
		}
	case oci_database_tools_runtime.DatabaseApiGatewayConfigCertificateBundleSelfSigned:
		result["type"] = "SELF_SIGNED"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) mapToDatabaseApiGatewayConfigCertificatePrivateKeyFileName(fieldKeyFormat string) (oci_database_tools_runtime.DatabaseApiGatewayConfigCertificatePrivateKeyFileName, error) {
	result := oci_database_tools_runtime.DatabaseApiGatewayConfigCertificatePrivateKeyFileName{}

	if format, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format")); ok {
		result.Format = oci_database_tools_runtime.DatabaseApiGatewayConfigCertificatePrivateKeyFileNameFormatEnum(format.(string))
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	return result, nil
}

func DatabaseApiGatewayConfigCertificatePrivateKeyFileNameToMap(obj *oci_database_tools_runtime.DatabaseApiGatewayConfigCertificatePrivateKeyFileName) map[string]interface{} {
	result := map[string]interface{}{}

	result["format"] = string(obj.Format)

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	return result
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) mapToDatabaseApiGatewayConfigCertificatePublicFileName(fieldKeyFormat string) (oci_database_tools_runtime.DatabaseApiGatewayConfigCertificatePublicFileName, error) {
	result := oci_database_tools_runtime.DatabaseApiGatewayConfigCertificatePublicFileName{}

	if format, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format")); ok {
		result.Format = oci_database_tools_runtime.DatabaseApiGatewayConfigCertificatePublicFileNameFormatEnum(format.(string))
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	return result, nil
}

func DatabaseApiGatewayConfigCertificatePublicFileNameToMap(obj *oci_database_tools_runtime.DatabaseApiGatewayConfigCertificatePublicFileName) map[string]interface{} {
	result := map[string]interface{}{}

	result["format"] = string(obj.Format)

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	return result
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest(request *oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest) error {
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
		details := oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDefaultDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if certificateBundle, ok := s.D.GetOkExists("certificate_bundle"); ok {
			if tmpList := certificateBundle.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_bundle", 0)
				tmp, err := s.mapToDatabaseApiGatewayConfigCertificateBundle(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CertificateBundle = tmp
			}
		}
		if databaseApiStatus, ok := s.D.GetOkExists("database_api_status"); ok {
			details.DatabaseApiStatus = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsDatabaseApiStatusEnum(databaseApiStatus.(string))
		}
		if documentRoot, ok := s.D.GetOkExists("document_root"); ok {
			tmp := documentRoot.(string)
			details.DocumentRoot = &tmp
		}
		if httpPort, ok := s.D.GetOkExists("http_port"); ok {
			tmp := httpPort.(int)
			details.HttpPort = &tmp
		}
		if httpsPort, ok := s.D.GetOkExists("https_port"); ok {
			tmp := httpsPort.(int)
			details.HttpsPort = &tmp
		}
		if poolRoute, ok := s.D.GetOkExists("pool_route"); ok {
			details.PoolRoute = oci_database_tools_runtime.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetailsPoolRouteEnum(poolRoute.(string))
		}
		if poolRoutingHeader, ok := s.D.GetOkExists("pool_routing_header"); ok {
			tmp := poolRoutingHeader.(string)
			details.PoolRoutingHeader = &tmp
		}
		request.UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
