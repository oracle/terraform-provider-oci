// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v65/dataconnectivity"

	"terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/tfresource"
)

func DataConnectivityRegistryConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataConnectivityRegistryConnection,
		Read:     readDataConnectivityRegistryConnection,
		Update:   updateDataConnectivityRegistryConnection,
		Delete:   deleteDataConnectivityRegistryConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"properties": {
				Type:     schema.TypeMap,
				Required: true,
				Elem:     schema.TypeString,
			},
			"registry_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"connection_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"aggregator": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"identifier": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"aggregator_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"created_by": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"created_by_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"identifier_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"info_fields": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"time_updated": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"updated_by": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"updated_by_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"model_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"object_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"primary_schema": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"identifier": {
							Type:     schema.TypeString,
							Required: true,
						},
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"model_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"default_connection": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"external_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_has_containers": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"aggregator": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"description": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"identifier": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"aggregator_key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"created_by": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"created_by_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"identifier_path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"info_fields": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"is_favorite": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"labels": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"registry_version": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"time_created": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"time_updated": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"updated_by": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"updated_by_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"model_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_status": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"object_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"parent": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"resource_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"registry_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"aggregator_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"created_by_user_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"created_by_user_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"time_updated": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"updated_by_user_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"updated_by_user_name": {
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
	}
}

func createDataConnectivityRegistryConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDataConnectivityRegistryConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDataConnectivityRegistryConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataConnectivityRegistryConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataConnectivityRegistryConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_connectivity.DataConnectivityManagementClient
	Res                    *oci_data_connectivity.Connection
	DisableNotFoundRetries bool
}

func (s *DataConnectivityRegistryConnectionResourceCrud) ID() string {
	return GetRegistryConnectionCompositeId(*(*s.Res).Key, s.D.Get("registry_id").(string))
}

func (s *DataConnectivityRegistryConnectionResourceCrud) Create() error {
	request := oci_data_connectivity.CreateConnectionRequest{}

	if connectionProperties, ok := s.D.GetOkExists("connection_properties"); ok {
		interfaces := connectionProperties.([]interface{})
		tmp := make([]oci_data_connectivity.ConnectionProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_properties", stateDataIndex)
			converted, err := s.mapToConnectionProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("connection_properties") {
			request.ConnectionProperties = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapToObjectMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Metadata = &tmp
		}
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		tmp := modelType.(string)
		request.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
	}

	if primarySchema, ok := s.D.GetOkExists("primary_schema"); ok {
		if tmpList := primarySchema.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "primary_schema", 0)
			tmp, err := s.mapToSchema(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PrimarySchema = &tmp
		}
	}

	if rawProperties, ok := s.D.GetOkExists("properties"); ok {
		properties, err := mapToPropertyMap(rawProperties.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.Properties = properties
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
		if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
			tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegistryMetadata = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.CreateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *DataConnectivityRegistryConnectionResourceCrud) Get() error {
	request := oci_data_connectivity.GetConnectionRequest{}

	if connectionKey, ok := s.D.GetOkExists("key"); ok {
		tmp := connectionKey.(string)
		request.ConnectionKey = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	connectionKey, registryId, err := parseRegistryConnectionCompositeId(s.D.Id())
	if err == nil {
		request.ConnectionKey = &connectionKey
		request.RegistryId = &registryId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *DataConnectivityRegistryConnectionResourceCrud) Update() error {
	request := oci_data_connectivity.UpdateConnectionRequest{}

	if connectionKey, ok := s.D.GetOkExists("key"); ok {
		tmp := connectionKey.(string)
		request.ConnectionKey = &tmp
	}

	if connectionProperties, ok := s.D.GetOkExists("connection_properties"); ok {
		interfaces := connectionProperties.([]interface{})
		tmp := make([]oci_data_connectivity.ConnectionProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_properties", stateDataIndex)
			converted, err := s.mapToConnectionProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("connection_properties") {
			request.ConnectionProperties = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapToObjectMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Metadata = &tmp
		}
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		tmp := modelType.(string)
		request.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
	}

	if primarySchema, ok := s.D.GetOkExists("primary_schema"); ok {
		if tmpList := primarySchema.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "primary_schema", 0)
			tmp, err := s.mapToSchema(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PrimarySchema = &tmp
		}
	}

	if rawProperties, ok := s.D.GetOkExists("properties"); ok {
		properties, err := mapToPropertyMap(rawProperties.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.Properties = properties
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
		if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
			tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegistryMetadata = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.UpdateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *DataConnectivityRegistryConnectionResourceCrud) Delete() error {
	request := oci_data_connectivity.DeleteConnectionRequest{}

	if connectionKey, ok := s.D.GetOkExists("key"); ok {
		tmp := connectionKey.(string)
		request.ConnectionKey = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	_, err := s.Client.DeleteConnection(context.Background(), request)
	return err
}

func (s *DataConnectivityRegistryConnectionResourceCrud) SetData() error {

	connectionKey, registryId, err := parseRegistryConnectionCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("connection_key", &connectionKey)
		s.D.Set("registry_id", &registryId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	connectionProperties := []interface{}{}
	for _, item := range s.Res.ConnectionProperties {
		connectionProperties = append(connectionProperties, DataConnectivityConnectionPropertyToMap(item))
	}
	s.D.Set("connection_properties", connectionProperties)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{DataConnectivityDataConnectivityObjectMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.ModelType != nil {
		s.D.Set("model_type", *s.Res.ModelType)
	}

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	if s.Res.PrimarySchema != nil {
		s.D.Set("primary_schema", []interface{}{DataConnectivitySchemaToMap(s.Res.PrimarySchema)})
	} else {
		s.D.Set("primary_schema", nil)
	}

	s.D.Set("properties", tfresource.GenericMapToJsonMap(s.Res.Properties))

	if s.Res.RegistryMetadata != nil {
		s.D.Set("registry_metadata", []interface{}{DataConnectivityDataConnectivityRegistryMetadataToMap(s.Res.RegistryMetadata)})
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}

func GetRegistryConnectionCompositeId(connectionKey string, registryId string) string {
	connectionKey = url.PathEscape(connectionKey)
	registryId = url.PathEscape(registryId)
	compositeId := "registries/" + registryId + "/connections/" + connectionKey
	return compositeId
}

func parseRegistryConnectionCompositeId(compositeId string) (connectionKey string, registryId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("registries/.*/connections/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	registryId, _ = url.PathUnescape(parts[1])
	connectionKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DataConnectivityRegistryConnectionResourceCrud) mapToAggregatorSummary(fieldKeyFormat string) (oci_data_connectivity.AggregatorSummary, error) {
	result := oci_data_connectivity.AggregatorSummary{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func DataConnectivityAggregatorSummaryToMap(obj *oci_data_connectivity.AggregatorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func DataConnectivityConnectionPropertyToMap(obj oci_data_connectivity.ConnectionProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func DataConnectivityConnectionSummaryToMap(obj oci_data_connectivity.ConnectionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	connectionProperties := []interface{}{}
	for _, item := range obj.ConnectionProperties {
		connectionProperties = append(connectionProperties, DataConnectivityConnectionPropertyToMap(item))
	}
	result["connection_properties"] = connectionProperties

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DataConnectivityDataConnectivityObjectMetadataToMap(obj.Metadata)}
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	if obj.PrimarySchema != nil {
		result["primary_schema"] = []interface{}{DataConnectivitySchemaToMap(obj.PrimarySchema)}
	}

	result["properties"] = tfresource.GenericMapToJsonMap(obj.Properties)

	if obj.RegistryMetadata != nil {
		result["registry_metadata"] = []interface{}{DataConnectivityDataConnectivityRegistryMetadataToMap(obj.RegistryMetadata)}
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DataConnectivityRegistryConnectionResourceCrud) mapToObjectMetadata(fieldKeyFormat string) (oci_data_connectivity.ObjectMetadata, error) {
	result := oci_data_connectivity.ObjectMetadata{}

	if aggregator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator")); ok {
		if tmpList := aggregator.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "aggregator"), 0)
			tmp, err := s.mapToAggregatorSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert aggregator, encountered error: %v", err)
			}
			result.Aggregator = &tmp
		}
	}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if createdBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by")); ok {
		tmp := createdBy.(string)
		result.CreatedBy = &tmp
	}

	if createdByName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by_name")); ok {
		tmp := createdByName.(string)
		result.CreatedByName = &tmp
	}

	if identifierPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier_path")); ok {
		tmp := identifierPath.(string)
		result.IdentifierPath = &tmp
	}

	if infoFields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "info_fields")); ok {
		result.InfoFields = tfresource.ObjectMapToStringMap(infoFields.(map[string]interface{}))
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if registryVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_version")); ok {
		tmp := registryVersion.(int)
		result.RegistryVersion = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		if timeCreated != "" {
			tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
			if err != nil {
				return result, err
			}
			result.TimeCreated = &oci_common.SDKTime{Time: tmp}
		}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		if timeUpdated != "" {
			tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
			if err != nil {
				return result, err
			}
			result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
		}
	}

	if updatedBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by")); ok {
		tmp := updatedBy.(string)
		result.UpdatedBy = &tmp
	}

	if updatedByName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by_name")); ok {
		tmp := updatedByName.(string)
		result.UpdatedByName = &tmp
	}

	return result, nil
}

func DataConnectivityDataConnectivityObjectMetadataToMap(obj *oci_data_connectivity.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Aggregator != nil {
		result["aggregator"] = []interface{}{DataConnectivityAggregatorSummaryToMap(obj.Aggregator)}
	}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.CreatedByName != nil {
		result["created_by_name"] = string(*obj.CreatedByName)
	}

	if obj.IdentifierPath != nil {
		result["identifier_path"] = string(*obj.IdentifierPath)
	}

	result["info_fields"] = obj.InfoFields

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	if obj.UpdatedBy != nil {
		result["updated_by"] = string(*obj.UpdatedBy)
	}

	if obj.UpdatedByName != nil {
		result["updated_by_name"] = string(*obj.UpdatedByName)
	}

	return result
}

func (s *DataConnectivityRegistryConnectionResourceCrud) mapToParentReference(fieldKeyFormat string) (oci_data_connectivity.ParentReference, error) {
	result := oci_data_connectivity.ParentReference{}

	if parent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent")); ok {
		tmp := parent.(string)
		result.Parent = &tmp
	}

	return result, nil
}

func (s *DataConnectivityRegistryConnectionResourceCrud) mapToRegistryMetadata(fieldKeyFormat string) (oci_data_connectivity.RegistryMetadata, error) {
	result := oci_data_connectivity.RegistryMetadata{}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if createdByUserId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by_user_id")); ok {
		tmp := createdByUserId.(string)
		result.CreatedByUserId = &tmp
	}

	if createdByUserName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by_user_name")); ok {
		tmp := createdByUserName.(string)
		result.CreatedByUserName = &tmp
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if registryVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_version")); ok {
		tmp := registryVersion.(int)
		result.RegistryVersion = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		if timeCreated != "" {
			tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
			if err != nil {
				return result, err
			}
			result.TimeCreated = &oci_common.SDKTime{Time: tmp}
		}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		if timeUpdated != "" {
			tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
			if err != nil {
				return result, err
			}
			result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
		}
	}

	if updatedByUserId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by_user_id")); ok {
		tmp := updatedByUserId.(string)
		result.UpdatedByUserId = &tmp
	}

	if updatedByUserName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by_user_name")); ok {
		tmp := updatedByUserName.(string)
		result.UpdatedByUserName = &tmp
	}

	return result, nil
}

func DataConnectivityDataConnectivityRegistryMetadataToMap(obj *oci_data_connectivity.RegistryMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.CreatedByUserId != nil {
		result["created_by_user_id"] = string(*obj.CreatedByUserId)
	}

	if obj.CreatedByUserName != nil {
		result["created_by_user_name"] = string(*obj.CreatedByUserName)
	}

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	if obj.UpdatedByUserId != nil {
		result["updated_by_user_id"] = string(*obj.UpdatedByUserId)
	}

	if obj.UpdatedByUserName != nil {
		result["updated_by_user_name"] = string(*obj.UpdatedByUserName)
	}

	return result
}

func (s *DataConnectivityRegistryConnectionResourceCrud) mapToSchema(fieldKeyFormat string) (oci_data_connectivity.Schema, error) {
	result := oci_data_connectivity.Schema{}

	if defaultConnection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_connection")); ok {
		tmp := defaultConnection.(string)
		result.DefaultConnection = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if externalKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "external_key")); ok {
		tmp := externalKey.(string)
		result.ExternalKey = &tmp
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if isHasContainers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_has_containers")); ok {
		tmp := isHasContainers.(bool)
		result.IsHasContainers = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metadata"), 0)
			tmp, err := s.mapToObjectMetadata(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert metadata, encountered error: %v", err)
			}
			result.Metadata = &tmp
		}
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_version")); ok {
		tmp := objectVersion.(int)
		result.ObjectVersion = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if resourceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_name")); ok {
		tmp := resourceName.(string)
		result.ResourceName = &tmp
	}

	return result, nil
}

func DataConnectivitySchemaToMap(obj *oci_data_connectivity.Schema) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultConnection != nil {
		result["default_connection"] = string(*obj.DefaultConnection)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.ExternalKey != nil {
		result["external_key"] = string(*obj.ExternalKey)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.IsHasContainers != nil {
		result["is_has_containers"] = bool(*obj.IsHasContainers)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DataConnectivityDataConnectivityObjectMetadataToMap(obj.Metadata)}
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{DataConnectivityParentReferenceToMap(obj.ParentRef)}
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	return result
}

func (s *DataConnectivityRegistryConnectionResourceCrud) mapToConnectionProperty(fieldKeyFormat string) (oci_data_connectivity.ConnectionProperty, error) {
	result := oci_data_connectivity.ConnectionProperty{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}
