// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v60/common"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v60/dataconnectivity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DataConnectivityRegistryDataAssetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataConnectivityRegistryDataAsset,
		Read:     readDataConnectivityRegistryDataAsset,
		Update:   updateDataConnectivityRegistryDataAsset,
		Delete:   deleteDataConnectivityRegistryDataAsset,
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
			"asset_properties": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"default_connection": {
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
							Computed: true,
							Optional: true,
						},
						"name": {
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
							Required: true,
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
						"properties": {
							Type:     schema.TypeMap,
							Required: true,
							Elem:     schema.TypeString,
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
						"type": {
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
			"external_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				Required: true,
			},
			"model_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"native_type_system": {
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
						"name": {
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
						"type_mapping_from": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"type_mapping_to": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"types": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"model_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"config_definition": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"config_parameter_definitions": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"is_contained": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
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
												"name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"object_status": {
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

												// Computed
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dt_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"model_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"object_status": {
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
									"type_system_name": {
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

func createDataConnectivityRegistryDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDataConnectivityRegistryDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDataConnectivityRegistryDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataConnectivityRegistryDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryDataAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataConnectivityRegistryDataAssetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_connectivity.DataConnectivityManagementClient
	Res                    *oci_data_connectivity.DataAsset
	DisableNotFoundRetries bool
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) ID() string {
	return GetRegistryDataAssetCompositeId(*(*s.Res).Key, s.D.Get("registry_id").(string))
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) Create() error {
	request := oci_data_connectivity.CreateDataAssetRequest{}

	if assetProperties, ok := s.D.GetOkExists("asset_properties"); ok {
		request.AssetProperties = tfresource.ObjectMapToStringMap(assetProperties.(map[string]interface{}))
	}

	if defaultConnection, ok := s.D.GetOkExists("default_connection"); ok {
		if tmpList := defaultConnection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "default_connection", 0)
			tmp, err := s.mapToConnection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DefaultConnection = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if externalKey, ok := s.D.GetOkExists("external_key"); ok {
		tmp := externalKey.(string)
		request.ExternalKey = &tmp
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
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

	if nativeTypeSystem, ok := s.D.GetOkExists("native_type_system"); ok {
		if tmpList := nativeTypeSystem.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "native_type_system", 0)
			tmp, err := s.mapToTypeSystem(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NativeTypeSystem = &tmp
		}
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
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

	response, err := s.Client.CreateDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) Get() error {
	request := oci_data_connectivity.GetDataAssetRequest{}

	if dataAssetKey, ok := s.D.GetOkExists("key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	dataAssetKey, registryId, err := parseRegistryDataAssetCompositeId(s.D.Id())
	if err == nil {
		request.DataAssetKey = &dataAssetKey
		request.RegistryId = &registryId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.GetDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) Update() error {
	request := oci_data_connectivity.UpdateDataAssetRequest{}

	if assetProperties, ok := s.D.GetOkExists("asset_properties"); ok {
		request.AssetProperties = tfresource.ObjectMapToStringMap(assetProperties.(map[string]interface{}))
	}

	if dataAssetKey, ok := s.D.GetOkExists("key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	if defaultConnection, ok := s.D.GetOkExists("default_connection"); ok {
		if tmpList := defaultConnection.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "default_connection", 0)
			tmp, err := s.mapToConnection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DefaultConnection = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if externalKey, ok := s.D.GetOkExists("external_key"); ok {
		tmp := externalKey.(string)
		request.ExternalKey = &tmp
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
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

	if nativeTypeSystem, ok := s.D.GetOkExists("native_type_system"); ok {
		if tmpList := nativeTypeSystem.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "native_type_system", 0)
			tmp, err := s.mapToTypeSystem(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NativeTypeSystem = &tmp
		}
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
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

	response, err := s.Client.UpdateDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataAsset
	return nil
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) Delete() error {
	request := oci_data_connectivity.DeleteDataAssetRequest{}

	if dataAssetKey, ok := s.D.GetOkExists("key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	_, err := s.Client.DeleteDataAsset(context.Background(), request)
	return err
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) SetData() error {

	dataAssetKey, registryId, err := parseRegistryDataAssetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("data_asset_key", &dataAssetKey)
		s.D.Set("registry_id", &registryId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("asset_properties", s.Res.AssetProperties)

	if s.Res.DefaultConnection != nil {
		s.D.Set("default_connection", []interface{}{DataConnectivityConnectionToMap(s.Res.DefaultConnection)})
	} else {
		s.D.Set("default_connection", nil)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.ExternalKey != nil {
		s.D.Set("external_key", *s.Res.ExternalKey)
	}

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{DataConnectivityObjectMetadataToMap(s.Res.Metadata)})
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

	if s.Res.NativeTypeSystem != nil {
		s.D.Set("native_type_system", []interface{}{TypeSystemToMap(s.Res.NativeTypeSystem)})
	} else {
		s.D.Set("native_type_system", nil)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	s.D.Set("properties", tfresource.GenericMapToJsonMap(s.Res.Properties))

	if s.Res.RegistryMetadata != nil {
		s.D.Set("registry_metadata", []interface{}{DataConnectivityRegistryMetadataToMap(s.Res.RegistryMetadata)})
	} else {
		s.D.Set("registry_metadata", nil)
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}

func GetRegistryDataAssetCompositeId(dataAssetKey string, registryId string) string {
	dataAssetKey = url.PathEscape(dataAssetKey)
	registryId = url.PathEscape(registryId)
	compositeId := "registries/" + registryId + "/dataAssets/" + dataAssetKey
	return compositeId
}

func parseRegistryDataAssetCompositeId(compositeId string) (dataAssetKey string, registryId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("registries/.*/dataAssets/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	registryId, _ = url.PathUnescape(parts[1])
	dataAssetKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToAggregatorSummary(fieldKeyFormat string) (oci_data_connectivity.AggregatorSummary, error) {
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

func AggregatorSummaryToMap(obj *oci_data_connectivity.AggregatorSummary) map[string]interface{} {
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

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToConfigDefinition(fieldKeyFormat string) (oci_data_connectivity.ConfigDefinition, error) {
	result := oci_data_connectivity.ConfigDefinition{}

	if configParameterDefinitions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_parameter_definitions")); ok {
		resultMap, err := s.objectMapToConfigParameterDefinitionsMap(configParameterDefinitions)
		if err != nil {
			return result, err
		}
		if len(resultMap) > 0 {
			result.ConfigParameterDefinitions = resultMap
		}
	}

	if isContained, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_contained")); ok {
		tmp := isContained.(bool)
		result.IsContained = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
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

	return result, nil
}

func ConfigDefinitionToMap(obj *oci_data_connectivity.ConfigDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	result["config_parameter_definitions"] = obj.ConfigParameterDefinitions

	if obj.IsContained != nil {
		result["is_contained"] = bool(*obj.IsContained)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
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

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{ParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func DataConnectivityConnectionToMap(obj *oci_data_connectivity.Connection) map[string]interface{} {
	result := map[string]interface{}{}

	connectionProperties := []interface{}{}
	for _, item := range obj.ConnectionProperties {
		connectionProperties = append(connectionProperties, ConnectionPropertyToMap(item))
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
		result["metadata"] = []interface{}{DataConnectivityObjectMetadataToMap(obj.Metadata)}
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
		result["primary_schema"] = []interface{}{SchemaToMap(obj.PrimarySchema)}
	}

	result["properties"] = tfresource.GenericMapToJsonMap(obj.Properties)

	if obj.RegistryMetadata != nil {
		result["registry_metadata"] = []interface{}{DataConnectivityRegistryMetadataToMap(obj.RegistryMetadata)}
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func ConnectionPropertyToMap(obj oci_data_connectivity.ConnectionProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func DataConnectivityDataAssetSummaryToMap(obj oci_data_connectivity.DataAssetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["asset_properties"] = obj.AssetProperties

	if obj.DefaultConnection != nil {
		result["default_connection"] = []interface{}{DataConnectivityConnectionToMap(obj.DefaultConnection)}
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

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{DataConnectivityObjectMetadataToMap(obj.Metadata)}
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

	if obj.NativeTypeSystem != nil {
		result["native_type_system"] = []interface{}{TypeSystemToMap(obj.NativeTypeSystem)}
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	result["properties"] = tfresource.GenericMapToJsonMap(obj.Properties)

	if obj.RegistryMetadata != nil {
		result["registry_metadata"] = []interface{}{DataConnectivityRegistryMetadataToMap(obj.RegistryMetadata)}
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToDataType(fieldKeyFormat string) (oci_data_connectivity.DataType, error) {
	result := oci_data_connectivity.DataType{}

	if configDefinition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_definition")); ok {
		if tmpList := configDefinition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "config_definition"), 0)
			tmp, err := s.mapToConfigDefinition(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert config_definition, encountered error: %v", err)
			}
			result.ConfigDefinition = &tmp
		}
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if dtType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dt_type")); ok {
		result.DtType = oci_data_connectivity.DataTypeDtTypeEnum(dtType.(string))
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
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

	if typeSystemName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type_system_name")); ok {
		tmp := typeSystemName.(string)
		result.TypeSystemName = &tmp
	}

	return result, nil
}

func DataTypeToMap(obj oci_data_connectivity.DataType) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigDefinition != nil {
		result["config_definition"] = []interface{}{ConfigDefinitionToMap(obj.ConfigDefinition)}
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["dt_type"] = string(obj.DtType)

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
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

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{ParentReferenceToMap(obj.ParentRef)}
	}

	if obj.TypeSystemName != nil {
		result["type_system_name"] = string(*obj.TypeSystemName)
	}

	return result
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToObjectMetadata(fieldKeyFormat string) (oci_data_connectivity.ObjectMetadata, error) {
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
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return result, err
		}
		result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
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

func DataConnectivityObjectMetadataToMap(obj *oci_data_connectivity.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Aggregator != nil {
		result["aggregator"] = []interface{}{AggregatorSummaryToMap(obj.Aggregator)}
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

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToParentReference(fieldKeyFormat string) (oci_data_connectivity.ParentReference, error) {
	result := oci_data_connectivity.ParentReference{}

	if parent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent")); ok {
		tmp := parent.(string)
		result.Parent = &tmp
	}

	return result, nil
}

func ParentReferenceToMap(obj *oci_data_connectivity.ParentReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Parent != nil {
		result["parent"] = string(*obj.Parent)
	}

	return result
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToRegistryMetadata(fieldKeyFormat string) (oci_data_connectivity.RegistryMetadata, error) {
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
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return result, err
		}
		result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
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

func DataConnectivityRegistryMetadataToMap(obj *oci_data_connectivity.RegistryMetadata) map[string]interface{} {
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

func SchemaToMap(obj *oci_data_connectivity.Schema) map[string]interface{} {
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
		result["metadata"] = []interface{}{DataConnectivityObjectMetadataToMap(obj.Metadata)}
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
		result["parent_ref"] = []interface{}{ParentReferenceToMap(obj.ParentRef)}
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	return result
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToTypeSystem(fieldKeyFormat string) (oci_data_connectivity.TypeSystem, error) {
	result := oci_data_connectivity.TypeSystem{}

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

	if typeMappingFrom, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type_mapping_from")); ok {
		result.TypeMappingFrom = tfresource.ObjectMapToStringMap(typeMappingFrom.(map[string]interface{}))
	}

	if typeMappingTo, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type_mapping_to")); ok {
		result.TypeMappingTo = tfresource.ObjectMapToStringMap(typeMappingTo.(map[string]interface{}))
	}

	if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
		interfaces := types.([]interface{})
		tmp := make([]oci_data_connectivity.DataType, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "types"), stateDataIndex)
			converted, err := s.mapToDataType(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "types")) {
			result.Types = tmp
		}
	}

	return result, nil
}

func TypeSystemToMap(obj *oci_data_connectivity.TypeSystem) map[string]interface{} {
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
		result["parent_ref"] = []interface{}{ParentReferenceToMap(obj.ParentRef)}
	}

	result["type_mapping_from"] = obj.TypeMappingFrom

	result["type_mapping_to"] = obj.TypeMappingTo

	types := []interface{}{}
	for _, item := range obj.Types {
		types = append(types, DataTypeToMap(item))
	}
	result["types"] = types

	return result
}

func mapToPropertyMap(rm map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for k, v := range rm {
		var val interface{}
		//Use the string value that was passed if it is not a valid JSON string
		if err := json.Unmarshal([]byte(v.(string)), &val); err == nil {
			result[k] = val
		} else {
			result[k] = v.(string)
		}
	}
	return result, nil
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) objectMapToConfigParameterDefinitionsMap(configParameterDefinitions interface{}) (map[string]oci_data_connectivity.ConfigParameterDefinition, error) {
	resultMap := map[string]oci_data_connectivity.ConfigParameterDefinition{}
	set := configParameterDefinitions.(*schema.Set)
	tmpList := set.List()
	for _, ifc := range tmpList {
		stateDataIndex := originGroupsHashCodeForSets(ifc)
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "origin_groups", stateDataIndex)
		converted, err := s.mapToConfigParameterDefinition(fieldKeyFormat)
		if err != nil {
			return nil, err
		}

		if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
			tmp := label.(string)
			resultMap[tmp] = converted
		}
	}

	return resultMap, nil
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToConfigParameterDefinition(fieldKeyFormat string) (oci_data_connectivity.ConfigParameterDefinition, error) {
	result := oci_data_connectivity.ConfigParameterDefinition{}

	if parameterType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameter_type")); ok {
		tmp := parameterType.(string)
		type_, err := DataConnectivityJsonToBaseType(tmp)
		if err == nil {
			result.ParameterType = type_
		}
	}

	if parameterName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameter_name")); ok {
		tmp := parameterName.(string)
		result.ParameterName = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if classFieldName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "class_field_name")); ok {
		tmp := classFieldName.(string)
		result.ClassFieldName = &tmp
	}

	if isStatic, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_static")); ok {
		tmp := isStatic.(bool)
		result.IsStatic = &tmp
	}

	if isClassFieldValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_class_field_value")); ok {
		tmp := isClassFieldValue.(bool)
		result.IsClassFieldValue = &tmp
	}

	return result, nil
}

func DataConnectivityJsonToBaseType(data string) (oci_data_connectivity.BaseType, error) {
	var val DataConnectivityBaseType
	if err := json.Unmarshal([]byte(data), &val); err == nil {
		if schemaData, err := UnmarshalPolymorphicDataConnectivityBaseTypeJSON(val.ModelType, data); err == nil {
			return schemaData, nil
		} else {
			return nil, err
		}
	}
	return nil, nil
}

type DataConnectivityBaseType struct {
	JsonData     []byte
	Key          *string                                `mandatory:"false" json:"key"`
	ModelVersion *string                                `mandatory:"false" json:"modelVersion"`
	ParentRef    *oci_data_connectivity.ParentReference `mandatory:"false" json:"parentRef"`
	Name         *string                                `mandatory:"false" json:"name"`
	ObjectStatus *int                                   `mandatory:"false" json:"objectStatus"`
	Description  *string                                `mandatory:"false" json:"description"`
	ModelType    string                                 `json:"modelType"`
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func UnmarshalPolymorphicDataConnectivityBaseTypeJSON(modelType string, data string) (oci_data_connectivity.BaseType, error) {

	var err error
	switch modelType {
	case "CONFIGURED_TYPE":
		mm := oci_data_connectivity.ConfiguredType{}
		err = json.Unmarshal([]byte(data), &mm)
		return mm, err
	case "DERIVED_TYPE":
		mm := oci_data_connectivity.DerivedType{}
		err = json.Unmarshal([]byte(data), &mm)
		return mm, err
	case "DATA_TYPE":
		mm := oci_data_connectivity.DataType{}
		err = json.Unmarshal([]byte(data), &mm)
		return mm, err
	case "COMPOSITE_TYPE":
		mm := oci_data_connectivity.CompositeType{}
		err = json.Unmarshal([]byte(data), &mm)
		return mm, err
	default:
		return nil, nil
	}
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToConnection(fieldKeyFormat string) (oci_data_connectivity.Connection, error) {
	result := oci_data_connectivity.Connection{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_version")); ok {
		tmp := objectVersion.(int)
		result.ObjectVersion = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if primarySchema, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_schema")); ok {
		if tmpList := primarySchema.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "primary_schema", 0)
			tmp, err := s.mapToSchema(fieldKeyFormat)
			if err != nil {
				return result, err
			}
			result.PrimarySchema = &tmp
		}
	}

	if rawProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
		properties, err := mapToPropertyMap(rawProperties.(map[string]interface{}))
		if err == nil {
			result.Properties = properties
		}
	}

	if connectionProperties, ok := s.D.GetOkExists("connection_properties"); ok {
		interfaces := connectionProperties.([]interface{})
		tmp := make([]oci_data_connectivity.ConnectionProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_properties", stateDataIndex)
			converted, err := s.mapToConnectionProperty(fieldKeyFormat)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("connection_properties") {
			result.ConnectionProperties = tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		result.IsDefault = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapToObjectMetadata(fieldKeyFormat)
			if err != nil {
				return result, err
			}
			result.Metadata = &tmp
		}
	}

	if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
		if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
			tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
			if err != nil {
				return result, err
			}
			result.RegistryMetadata = &tmp
		}
	}

	return result, nil
}

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToConnectionProperty(fieldKeyFormat string) (oci_data_connectivity.ConnectionProperty, error) {
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

func (s *DataConnectivityRegistryDataAssetResourceCrud) mapToSchema(fieldKeyFormat string) (oci_data_connectivity.Schema, error) {
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

func originGroupsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if label, ok := m["label"]; ok && label != "" {
		buf.WriteString(fmt.Sprintf("%v-", label))
	}

	if originGroup, ok := m["origin_group"]; ok {
		if tmpList := originGroup.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("origin_group-")
			for _, originGroupRaw := range tmpList {
				tmpMap := originGroupRaw.(map[string]interface{})
				if name, ok := tmpMap["origin"]; ok {
					buf.WriteString(fmt.Sprintf("%v-", name))
				}
				if value, ok := tmpMap["weight"]; ok {
					buf.WriteString(fmt.Sprintf("%v-", value))
				}
			}
		}
	}
	return utils.GetStringHashcode(buf.String())
}
