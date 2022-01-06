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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_data_connectivity "github.com/oracle/oci-go-sdk/v60/dataconnectivity"
)

func DataConnectivityRegistryFolderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataConnectivityRegistryFolder,
		Read:     readDataConnectivityRegistryFolder,
		Update:   updateDataConnectivityRegistryFolder,
		Delete:   deleteDataConnectivityRegistryFolder,
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
			"registry_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"data_assets": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
						"name": {
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
										Required: true,
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
									"properties": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
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
						"properties": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
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

			// Computed
		},
	}
}

func createDataConnectivityRegistryFolder(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryFolderResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDataConnectivityRegistryFolder(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryFolderResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDataConnectivityRegistryFolder(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryFolderResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataConnectivityRegistryFolder(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryFolderResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataConnectivityRegistryFolderResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_connectivity.DataConnectivityManagementClient
	Res                    *oci_data_connectivity.Folder
	DisableNotFoundRetries bool
}

func (s *DataConnectivityRegistryFolderResourceCrud) ID() string {
	return GetRegistryFolderCompositeId(*(*s.Res).Key, s.D.Get("registry_id").(string))
}

func (s *DataConnectivityRegistryFolderResourceCrud) Create() error {
	request := oci_data_connectivity.CreateFolderRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
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

	if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
			tmp, err := s.mapToParentReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ParentRef = &tmp
		}
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.CreateFolder(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Folder
	return nil
}

func (s *DataConnectivityRegistryFolderResourceCrud) Get() error {
	request := oci_data_connectivity.GetFolderRequest{}

	if folderKey, ok := s.D.GetOkExists("key"); ok {
		tmp := folderKey.(string)
		request.FolderKey = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	folderKey, registryId, err := parseRegistryFolderCompositeId(s.D.Id())
	if err == nil {
		request.FolderKey = &folderKey
		request.RegistryId = &registryId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.GetFolder(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Folder
	return nil
}

func (s *DataConnectivityRegistryFolderResourceCrud) Update() error {
	request := oci_data_connectivity.UpdateFolderRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if folderKey, ok := s.D.GetOkExists("key"); ok {
		tmp := folderKey.(string)
		request.FolderKey = &tmp
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
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

	if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
			tmp, err := s.mapToParentReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ParentRef = &tmp
		}
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	response, err := s.Client.UpdateFolder(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Folder
	return nil
}

func (s *DataConnectivityRegistryFolderResourceCrud) Delete() error {
	request := oci_data_connectivity.DeleteFolderRequest{}

	if folderKey, ok := s.D.GetOkExists("key"); ok {
		tmp := folderKey.(string)
		request.FolderKey = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_connectivity")

	_, err := s.Client.DeleteFolder(context.Background(), request)
	return err
}

func (s *DataConnectivityRegistryFolderResourceCrud) SetData() error {

	folderKey, registryId, err := parseRegistryFolderCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("folder_key", &folderKey)
		s.D.Set("registry_id", &registryId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
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

	if s.Res.ParentRef != nil {
		s.D.Set("parent_ref", []interface{}{DataConnectivityParentReferenceToMap(s.Res.ParentRef)})
	} else {
		s.D.Set("parent_ref", nil)
	}

	return nil
}

func GetRegistryFolderCompositeId(folderKey string, registryId string) string {
	folderKey = url.PathEscape(folderKey)
	registryId = url.PathEscape(registryId)
	compositeId := "registries/" + registryId + "/folders/" + folderKey
	return compositeId
}

func parseRegistryFolderCompositeId(compositeId string) (folderKey string, registryId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("registries/.*/folders/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	registryId, _ = url.PathUnescape(parts[1])
	folderKey, _ = url.PathUnescape(parts[3])

	return
}

func FolderSummaryToMap(obj oci_data_connectivity.FolderSummary) map[string]interface{} {
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
		result["parent_ref"] = []interface{}{DataConnectivityParentReferenceToMap(obj.ParentRef)}
	}

	return result
}

func (s *DataConnectivityRegistryFolderResourceCrud) mapToParentReference(fieldKeyFormat string) (oci_data_connectivity.ParentReference, error) {
	result := oci_data_connectivity.ParentReference{}

	if parent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent")); ok {
		tmp := parent.(string)
		result.Parent = &tmp
	}

	return result, nil
}

func DataConnectivityParentReferenceToMap(obj *oci_data_connectivity.ParentReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Parent != nil {
		result["parent"] = string(*obj.Parent)
	}

	return result
}
