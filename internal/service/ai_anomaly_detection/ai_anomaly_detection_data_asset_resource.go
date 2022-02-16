// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v58/aianomalydetection"
)

func AiAnomalyDetectionDataAssetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiAnomalyDetectionDataAsset,
		Read:     readAiAnomalyDetectionDataAsset,
		Update:   updateAiAnomalyDetectionDataAsset,
		Delete:   deleteAiAnomalyDetectionDataAsset,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_source_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"data_source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"INFLUX",
								"ORACLE_ATP",
								"ORACLE_OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"atp_password_secret_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"atp_user_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"cwallet_file_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"database_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ewallet_file_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"key_store_file_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"measurement_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"object": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ojdbc_file_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"password_secret_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"table_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"tnsnames_file_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"truststore_file_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"version_specific_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"influx_version": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"V_1_8",
											"V_2_0",
										}, true),
									},

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"database_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"organization_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"retention_policy_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"wallet_password_secret_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},

						// Computed
					},
				},
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createAiAnomalyDetectionDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.CreateResource(d, sync)
}

func readAiAnomalyDetectionDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

func updateAiAnomalyDetectionDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiAnomalyDetectionDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiAnomalyDetectionDataAssetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res                    *oci_ai_anomaly_detection.DataAsset
	DisableNotFoundRetries bool
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.DataAssetLifecycleStateActive),
	}
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.DataAssetLifecycleStateDeleted),
	}
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) Create() error {
	request := oci_ai_anomaly_detection.CreateDataAssetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataSourceDetails, ok := s.D.GetOkExists("data_source_details"); ok {
		if tmpList := dataSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_source_details", 0)
			tmp, err := s.mapToDataSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataSourceDetails = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.CreateDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetDataAssetRequest{}

	tmp := s.D.Id()
	request.DataAssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.GetDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_anomaly_detection.UpdateDataAssetRequest{}

	tmp := s.D.Id()
	request.DataAssetId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.UpdateDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) Delete() error {
	request := oci_ai_anomaly_detection.DeleteDataAssetRequest{}

	tmp := s.D.Id()
	request.DataAssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	_, err := s.Client.DeleteDataAsset(context.Background(), request)
	return err
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataSourceDetails != nil {
		dataSourceDetailsArray := []interface{}{}
		if dataSourceDetailsMap := DataSourceDetailsToMap(&s.Res.DataSourceDetails); dataSourceDetailsMap != nil {
			dataSourceDetailsArray = append(dataSourceDetailsArray, dataSourceDetailsMap)
		}
		s.D.Set("data_source_details", dataSourceDetailsArray)
	} else {
		s.D.Set("data_source_details", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.PrivateEndpointId != nil {
		s.D.Set("private_endpoint_id", *s.Res.PrivateEndpointId)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func AiDataAssetSummaryToMap(obj oci_ai_anomaly_detection.DataAssetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	dataSourceDetailsResults := []interface{}{}
	var baseObject oci_ai_anomaly_detection.DataSourceDetails
	dsObj := oci_ai_anomaly_detection.DataSourceDetailsObjectStorage{}
	baseObject = dsObj
	dataSourceDetailsResults = append(dataSourceDetailsResults, DataSourceDetailsToMap(&baseObject))

	result["data_source_details"] = dataSourceDetailsResults

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.PrivateEndpointId != nil {
		result["private_endpoint_id"] = string(*obj.PrivateEndpointId)
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) mapToDataSourceDetails(fieldKeyFormat string) (oci_ai_anomaly_detection.DataSourceDetails, error) {
	var baseObject oci_ai_anomaly_detection.DataSourceDetails
	//discriminator
	dataSourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_source_type"))
	var dataSourceType string
	if ok {
		dataSourceType = dataSourceTypeRaw.(string)
	} else {
		dataSourceType = "" // default value
	}
	switch strings.ToLower(dataSourceType) {
	case strings.ToLower("INFLUX"):
		details := oci_ai_anomaly_detection.DataSourceDetailsInflux{}
		if measurementName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "measurement_name")); ok {
			tmp := measurementName.(string)
			details.MeasurementName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		if versionSpecificDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_specific_details")); ok {
			if tmpList := versionSpecificDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "version_specific_details"), 0)
				tmp, err := s.mapToInfluxDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert version_specific_details, encountered error: %v", err)
				}
				details.VersionSpecificDetails = tmp
			}
		}
		baseObject = details
	case strings.ToLower("ORACLE_ATP"):
		details := oci_ai_anomaly_detection.DataSourceDetailsAtp{}
		if atpPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "atp_password_secret_id")); ok {
			tmp := atpPasswordSecretId.(string)
			details.AtpPasswordSecretId = &tmp
		}
		if atpUserName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "atp_user_name")); ok {
			tmp := atpUserName.(string)
			details.AtpUserName = &tmp
		}
		if cwalletFileSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cwallet_file_secret_id")); ok {
			tmp := cwalletFileSecretId.(string)
			details.CwalletFileSecretId = &tmp
		}
		if databaseName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_name")); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if ewalletFileSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ewallet_file_secret_id")); ok {
			tmp := ewalletFileSecretId.(string)
			details.EwalletFileSecretId = &tmp
		}
		if keyStoreFileSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_file_secret_id")); ok {
			tmp := keyStoreFileSecretId.(string)
			details.KeyStoreFileSecretId = &tmp
		}
		if ojdbcFileSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ojdbc_file_secret_id")); ok {
			tmp := ojdbcFileSecretId.(string)
			details.OjdbcFileSecretId = &tmp
		}
		if tableName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "table_name")); ok {
			tmp := tableName.(string)
			details.TableName = &tmp
		}
		if tnsnamesFileSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tnsnames_file_secret_id")); ok {
			tmp := tnsnamesFileSecretId.(string)
			details.TnsnamesFileSecretId = &tmp
		}
		if truststoreFileSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "truststore_file_secret_id")); ok {
			tmp := truststoreFileSecretId.(string)
			details.TruststoreFileSecretId = &tmp
		}
		if walletPasswordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wallet_password_secret_id")); ok {
			tmp := walletPasswordSecretId.(string)
			details.WalletPasswordSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("ORACLE_OBJECT_STORAGE"):
		details := oci_ai_anomaly_detection.DataSourceDetailsObjectStorage{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown data_source_type '%v' was specified", dataSourceType)
	}
	return baseObject, nil
}

func DataSourceDetailsToMap(obj *oci_ai_anomaly_detection.DataSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_anomaly_detection.DataSourceDetailsInflux:
		result["data_source_type"] = "INFLUX"

		if v.MeasurementName != nil {
			result["measurement_name"] = string(*v.MeasurementName)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.Url != nil {
			result["url"] = string(*v.Url)
		}

		if v.UserName != nil {
			result["user_name"] = string(*v.UserName)
		}

		if v.VersionSpecificDetails != nil {
			versionSpecificDetailsArray := []interface{}{}
			if versionSpecificDetailsMap := InfluxDetailsToMap(&v.VersionSpecificDetails); versionSpecificDetailsMap != nil {
				versionSpecificDetailsArray = append(versionSpecificDetailsArray, versionSpecificDetailsMap)
			}
			result["version_specific_details"] = versionSpecificDetailsArray
		}
	case oci_ai_anomaly_detection.DataSourceDetailsAtp:
		result["data_source_type"] = "ORACLE_ATP"

		if v.AtpPasswordSecretId != nil {
			result["atp_password_secret_id"] = string(*v.AtpPasswordSecretId)
		}

		if v.AtpUserName != nil {
			result["atp_user_name"] = string(*v.AtpUserName)
		}

		if v.CwalletFileSecretId != nil {
			result["cwallet_file_secret_id"] = string(*v.CwalletFileSecretId)
		}

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.EwalletFileSecretId != nil {
			result["ewallet_file_secret_id"] = string(*v.EwalletFileSecretId)
		}

		if v.KeyStoreFileSecretId != nil {
			result["key_store_file_secret_id"] = string(*v.KeyStoreFileSecretId)
		}

		if v.OjdbcFileSecretId != nil {
			result["ojdbc_file_secret_id"] = string(*v.OjdbcFileSecretId)
		}

		if v.TableName != nil {
			result["table_name"] = string(*v.TableName)
		}

		if v.TnsnamesFileSecretId != nil {
			result["tnsnames_file_secret_id"] = string(*v.TnsnamesFileSecretId)
		}

		if v.TruststoreFileSecretId != nil {
			result["truststore_file_secret_id"] = string(*v.TruststoreFileSecretId)
		}

		if v.WalletPasswordSecretId != nil {
			result["wallet_password_secret_id"] = string(*v.WalletPasswordSecretId)
		}
	case oci_ai_anomaly_detection.DataSourceDetailsObjectStorage:
		result["data_source_type"] = "ORACLE_OBJECT_STORAGE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}
	default:
		log.Printf("[WARN] Received 'data_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) mapToInfluxDetails(fieldKeyFormat string) (oci_ai_anomaly_detection.InfluxDetails, error) {
	var baseObject oci_ai_anomaly_detection.InfluxDetails
	//discriminator
	influxVersionRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "influx_version"))
	var influxVersion string
	if ok {
		influxVersion = influxVersionRaw.(string)
	} else {
		influxVersion = "" // default value
	}
	switch strings.ToLower(influxVersion) {
	case strings.ToLower("V_1_8"):
		details := oci_ai_anomaly_detection.InfluxDetailsV1v8{}
		if databaseName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_name")); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if retentionPolicyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_policy_name")); ok {
			tmp := retentionPolicyName.(string)
			details.RetentionPolicyName = &tmp
		}
		baseObject = details
	case strings.ToLower("V_2_0"):
		details := oci_ai_anomaly_detection.InfluxDetailsV2v0{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if organizationName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "organization_name")); ok {
			tmp := organizationName.(string)
			details.OrganizationName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown influx_version '%v' was specified", influxVersion)
	}
	return baseObject, nil
}

func InfluxDetailsToMap(obj *oci_ai_anomaly_detection.InfluxDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_anomaly_detection.InfluxDetailsV1v8:
		result["influx_version"] = "V_1_8"

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.RetentionPolicyName != nil {
			result["retention_policy_name"] = string(*v.RetentionPolicyName)
		}
	case oci_ai_anomaly_detection.InfluxDetailsV2v0:
		result["influx_version"] = "V_2_0"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.OrganizationName != nil {
			result["organization_name"] = string(*v.OrganizationName)
		}
	default:
		log.Printf("[WARN] Received 'influx_version' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiAnomalyDetectionDataAssetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_anomaly_detection.ChangeDataAssetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DataAssetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	_, err := s.Client.ChangeDataAssetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
