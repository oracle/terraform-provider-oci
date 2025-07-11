// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementCloudDbSystemDiscoveryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementCloudDbSystemDiscovery,
		Read:     readDatabaseManagementCloudDbSystemDiscovery,
		Update:   updateDatabaseManagementCloudDbSystemDiscovery,
		Delete:   deleteDatabaseManagementCloudDbSystemDiscovery,
		Schema: map[string]*schema.Schema{
			// Required
			"agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cloud_db_system_discovery_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dbaas_parent_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deployment_type": {
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
			"patch_operations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"operation": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"MERGE",
							}, true),
						},
						"selection": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"value": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"is_selected_for_monitoring": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"connector": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"connector_type": {
													Type:     schema.TypeString,
													Required: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"agent_id": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"connection_info": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"component_type": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional
															"connection_credentials": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"credential_type": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		// Optional
																		"credential_name": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"password_secret_id": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"role": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"ssl_secret_id": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"user_name": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		// Computed
																	},
																},
															},
															"connection_string": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"host_name": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"hosts": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"port": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"protocol": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"service": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		// Computed
																	},
																},
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"discovered_components": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"adr_home_directory": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"asm_instances": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"adr_home_directory": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"associated_components": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"association_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"component_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"component_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"can_enable_all_current_pdbs": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"cluster_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cluster_instances": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"adr_home_directory": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cluster_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connector": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"agent_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"connection_failure_message": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"connection_info": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"component_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"connection_credentials": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"credential_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"credential_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"named_credential_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"password_secret_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"role": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"ssl_secret_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"user_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"connection_string": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"host_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"hosts": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"port": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"protocol": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"service": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"connection_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"connector_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_connection_status_last_updated": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"crs_base_directory": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_role": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connector": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"agent_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connection_failure_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connection_info": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"component_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"connection_credentials": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"credential_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"credential_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"named_credential_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"password_secret_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"role": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"ssl_secret_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"connection_string": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"host_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"hosts": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"port": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"protocol": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"service": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"connection_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connector_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_connection_status_last_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu_core_count": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"crs_base_directory": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_edition": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_instances": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"adr_home_directory": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"node_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_home": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"db_node_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_packs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dbaas_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"endpoints": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"host": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"protocol": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"services": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"grid_home": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"guid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"home_directory": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_auto_enable_pluggable_database": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_cluster": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_flex_cluster": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_flex_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_selected_for_monitoring": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"listener_alias": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"listener_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"log_directory": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_size_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"network_configurations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"network_number": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"network_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subnet": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"node_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ocr_file_location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"oracle_home": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pluggable_databases": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connector": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"agent_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"connection_failure_message": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"connection_info": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"component_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"connection_credentials": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"credential_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"credential_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"named_credential_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"password_secret_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"role": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"ssl_secret_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"user_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"connection_string": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"host_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"hosts": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"port": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"protocol": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"service": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"connection_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"connector_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_connection_status_last_updated": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"container_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"guid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scan_configurations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"network_number": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"scan_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"scan_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"scan_protocol": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trace_directory": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vip_configurations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_number": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"node_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"grid_home": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func createDatabaseManagementCloudDbSystemDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemDiscoveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementCloudDbSystemDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemDiscoveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementCloudDbSystemDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemDiscoveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementCloudDbSystemDiscovery(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemDiscoveryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseManagementCloudDbSystemDiscoveryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.CloudDbSystemDiscovery
	PatchResponse          *oci_database_management.CloudDbSystemDiscovery
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.CloudDbSystemDiscoveryLifecycleStateCreating),
	}
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.CloudDbSystemDiscoveryLifecycleStateActive),
	}
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.CloudDbSystemDiscoveryLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.CloudDbSystemDiscoveryLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) Create() error {
	request := oci_database_management.CreateCloudDbSystemDiscoveryRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbaasParentInfrastructureId, ok := s.D.GetOkExists("dbaas_parent_infrastructure_id"); ok {
		tmp := dbaasParentInfrastructureId.(string)
		request.DbaasParentInfrastructureId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if deploymentType, ok := s.D.GetOkExists("deployment_type"); ok {
		request.DeploymentType = oci_database_management.CloudDbSystemDeploymentTypeEnum(deploymentType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.CreateCloudDbSystemDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	err = s.getCloudDbSystemDiscoveryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management"), oci_database_management.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	if _, ok := s.D.GetOkExists("patch_operations"); ok {
		err = s.Patch()
		if err != nil {
			log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
			return err
		}
	}
	return nil
}
func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) Patch() error {
	request := oci_database_management.PatchCloudDbSystemDiscoveryRequest{}

	tmpId := s.D.Id()
	request.CloudDbSystemDiscoveryId = &tmpId

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
		interfaces := patchOperations.([]interface{})
		tmp := make([]oci_database_management.PatchInstruction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_operations", stateDataIndex)
			converted, err := s.mapToPatchInstruction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("patch_operations") {
			request.Items = tmp
		}
	}

	if len(request.Items) == 0 {
		return nil
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")
	response, err := s.Client.PatchCloudDbSystemDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	s.PatchResponse = &response.CloudDbSystemDiscovery
	return nil
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) getCloudDbSystemDiscoveryFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_management.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	cloudDbSystemDiscoveryId, err := cloudDbSystemDiscoveryWaitForWorkRequest(workId, "dbsystemdiscovery",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*cloudDbSystemDiscoveryId)

	return s.Get()
}

func cloudDbSystemDiscoveryWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func cloudDbSystemDiscoveryWaitForWorkRequest(wId *string, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_management.DbManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_management")
	retryPolicy.ShouldRetryOperation = cloudDbSystemDiscoveryWorkRequestShouldRetryFunc(timeout)

	response := oci_database_management.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_database_management.WorkRequestStatusInProgress),
			string(oci_database_management.WorkRequestStatusAccepted),
			string(oci_database_management.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_database_management.WorkRequestStatusSucceeded),
			string(oci_database_management.WorkRequestStatusFailed),
			string(oci_database_management.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_management.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_database_management.WorkRequestStatusFailed || response.Status == oci_database_management.WorkRequestStatusCanceled {
		return nil, getErrorFromDatabaseManagementCloudDbSystemDiscoveryWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseManagementCloudDbSystemDiscoveryWorkRequest(client *oci_database_management.DbManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_management.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_management.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) Get() error {
	request := oci_database_management.GetCloudDbSystemDiscoveryRequest{}

	tmp := s.D.Id()
	request.CloudDbSystemDiscoveryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetCloudDbSystemDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudDbSystemDiscovery
	return nil
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) Update() error {
	request := oci_database_management.UpdateCloudDbSystemDiscoveryRequest{}

	tmp := s.D.Id()
	request.CloudDbSystemDiscoveryId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateCloudDbSystemDiscovery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudDbSystemDiscovery
	err = s.Patch()
	if err != nil {
		log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
		return err
	}
	return nil
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) Delete() error {
	request := oci_database_management.DeleteCloudDbSystemDiscoveryRequest{}

	tmp := s.D.Id()
	request.CloudDbSystemDiscoveryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	_, err := s.Client.DeleteCloudDbSystemDiscovery(context.Background(), request)
	return err
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) SetData() error {
	if s.Res.AgentId != nil {
		s.D.Set("agent_id", *s.Res.AgentId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbaasParentInfrastructureId != nil {
		s.D.Set("dbaas_parent_infrastructure_id", *s.Res.DbaasParentInfrastructureId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Id != nil {
		s.D.Set("cloud_db_system_discovery_id", *s.Res.Id)
	}

	s.D.Set("deployment_type", s.Res.DeploymentType)

	discoveredComponents := []interface{}{}
	for _, item := range s.Res.DiscoveredComponents {
		discoveredComponents = append(discoveredComponents, DiscoveredCloudDbSystemComponentToMap(item))
	}
	s.D.Set("discovered_components", discoveredComponents)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GridHome != nil {
		s.D.Set("grid_home", *s.Res.GridHome)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
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

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToAsmConnectionString(fieldKeyFormat string) (oci_database_management.AsmConnectionString, error) {
	result := oci_database_management.AsmConnectionString{}

	if hosts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hosts")); ok {
		interfaces := hosts.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hosts")) {
			result.Hosts = tmp
		}
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_database_management.AsmConnectionStringProtocolEnum(protocol.(string))
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		tmp := service.(string)
		result.Service = &tmp
	}

	return result, nil
}

func AsmConnectionStringToMap(obj *oci_database_management.AsmConnectionString) map[string]interface{} {
	result := map[string]interface{}{}

	result["hosts"] = obj.Hosts

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.Service != nil {
		result["service"] = string(*obj.Service)
	}

	return result
}

func AssociatedCloudComponentToMap(obj oci_database_management.AssociatedCloudComponent) map[string]interface{} {
	result := map[string]interface{}{}

	result["association_type"] = string(obj.AssociationType)

	if obj.ComponentId != nil {
		result["component_id"] = string(*obj.ComponentId)
	}

	result["component_type"] = string(obj.ComponentType)

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToCloudAsmConnectionCredentials(fieldKeyFormat string) (oci_database_management.CloudAsmConnectionCredentials, error) {
	var baseObject oci_database_management.CloudAsmConnectionCredentials
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "DETAILS" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("DETAILS"):
		details := oci_database_management.CloudAsmConnectionCredentialsByDetails{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.CloudAsmConnectionCredentialsByDetailsRoleEnum(role.(string))
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	case strings.ToLower("NAME_REFERENCE"):
		details := oci_database_management.CloudAsmConnectionCredentialsByName{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func CloudAsmConnectionCredentialsToMap(obj *oci_database_management.CloudAsmConnectionCredentials) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.CloudAsmConnectionCredentialsByDetails:
		result["credential_type"] = "DETAILS"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["role"] = string(v.Role)

		if v.UserName != nil {
			result["user_name"] = string(*v.UserName)
		}
	case oci_database_management.CloudAsmConnectionCredentialsByName:
		result["credential_type"] = "NAME_REFERENCE"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}
	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToCloudClusterNetworkConfiguration(fieldKeyFormat string) (oci_database_management.CloudClusterNetworkConfiguration, error) {
	result := oci_database_management.CloudClusterNetworkConfiguration{}

	if networkNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_number")); ok {
		tmp := networkNumber.(int)
		result.NetworkNumber = &tmp
	}

	if networkType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_type")); ok {
		result.NetworkType = oci_database_management.CloudClusterNetworkConfigurationNetworkTypeEnum(networkType.(string))
	}

	if subnet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet")); ok {
		tmp := subnet.(string)
		result.Subnet = &tmp
	}

	return result, nil
}

func CloudClusterNetworkConfigurationToMap(obj oci_database_management.CloudClusterNetworkConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NetworkNumber != nil {
		result["network_number"] = int(*obj.NetworkNumber)
	}

	result["network_type"] = string(obj.NetworkType)

	if obj.Subnet != nil {
		result["subnet"] = string(*obj.Subnet)
	}

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToCloudClusterScanListenerConfiguration(fieldKeyFormat string) (oci_database_management.CloudClusterScanListenerConfiguration, error) {
	result := oci_database_management.CloudClusterScanListenerConfiguration{}

	if networkNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_number")); ok {
		tmp := networkNumber.(int)
		result.NetworkNumber = &tmp
	}

	if scanName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_name")); ok {
		tmp := scanName.(string)
		result.ScanName = &tmp
	}

	if scanPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_port")); ok {
		tmp := scanPort.(int)
		result.ScanPort = &tmp
	}

	if scanProtocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_protocol")); ok {
		result.ScanProtocol = oci_database_management.CloudClusterScanListenerConfigurationScanProtocolEnum(scanProtocol.(string))
	}

	return result, nil
}

func CloudClusterScanListenerConfigurationToMap(obj oci_database_management.CloudClusterScanListenerConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NetworkNumber != nil {
		result["network_number"] = int(*obj.NetworkNumber)
	}

	if obj.ScanName != nil {
		result["scan_name"] = string(*obj.ScanName)
	}

	if obj.ScanPort != nil {
		result["scan_port"] = int(*obj.ScanPort)
	}

	result["scan_protocol"] = string(obj.ScanProtocol)

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToCloudClusterVipConfiguration(fieldKeyFormat string) (oci_database_management.CloudClusterVipConfiguration, error) {
	result := oci_database_management.CloudClusterVipConfiguration{}

	if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
		tmp := address.(string)
		result.Address = &tmp
	}

	if networkNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_number")); ok {
		tmp := networkNumber.(int)
		result.NetworkNumber = &tmp
	}

	if nodeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_name")); ok {
		tmp := nodeName.(string)
		result.NodeName = &tmp
	}

	return result, nil
}

func CloudClusterVipConfigurationToMap(obj oci_database_management.CloudClusterVipConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address != nil {
		result["address"] = string(*obj.Address)
	}

	if obj.NetworkNumber != nil {
		result["network_number"] = int(*obj.NetworkNumber)
	}

	if obj.NodeName != nil {
		result["node_name"] = string(*obj.NodeName)
	}

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToCloudDbSystemConnectionInfo(fieldKeyFormat string) (oci_database_management.CloudDbSystemConnectionInfo, error) {
	var baseObject oci_database_management.CloudDbSystemConnectionInfo
	//discriminator
	componentTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "component_type"))
	var componentType string
	if ok {
		componentType = componentTypeRaw.(string)
	} else {
		componentType = "" // default value
	}
	switch strings.ToLower(componentType) {
	case strings.ToLower("ASM"):
		details := oci_database_management.CloudAsmConnectionInfo{}
		if connectionCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_credentials")); ok {
			if tmpList := connectionCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_credentials"), 0)
				tmp, err := s.mapToCloudAsmConnectionCredentials(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_credentials, encountered error: %v", err)
				}
				details.ConnectionCredentials = tmp
			}
		}
		if connectionString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_string")); ok {
			if tmpList := connectionString.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_string"), 0)
				tmp, err := s.mapToAsmConnectionString(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_string, encountered error: %v", err)
				}
				details.ConnectionString = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("DATABASE"):
		details := oci_database_management.CloudDatabaseConnectionInfo{}
		if connectionCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_credentials")); ok {
			if tmpList := connectionCredentials.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_credentials"), 0)
				tmp, err := s.mapToDatabaseConnectionCredentials(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_credentials, encountered error: %v", err)
				}
				details.ConnectionCredentials = tmp
			}
		}
		if connectionString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_string")); ok {
			if tmpList := connectionString.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_string"), 0)
				tmp, err := s.mapToDatabaseConnectionString(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_string, encountered error: %v", err)
				}
				details.ConnectionString = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown component_type '%v' was specified", componentType)
	}
	return baseObject, nil
}

func CloudDbSystemConnectionInfoToMap(obj *oci_database_management.CloudDbSystemConnectionInfo) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.CloudAsmConnectionInfo:
		result["component_type"] = "ASM"

		if v.ConnectionCredentials != nil {
			connectionCredentialsArray := []interface{}{}
			if connectionCredentialsMap := CloudAsmConnectionCredentialsToMap(&v.ConnectionCredentials); connectionCredentialsMap != nil {
				connectionCredentialsArray = append(connectionCredentialsArray, connectionCredentialsMap)
			}
			result["connection_credentials"] = connectionCredentialsArray
		}

		if v.ConnectionString != nil {
			result["connection_string"] = []interface{}{AsmConnectionStringToMap(v.ConnectionString)}
		}
	case oci_database_management.CloudDatabaseConnectionInfo:
		result["component_type"] = "DATABASE"

		if v.ConnectionCredentials != nil {
			connectionCredentialsArray := []interface{}{}
			if connectionCredentialsMap := DatabaseConnectionCredentialsToMap(&v.ConnectionCredentials); connectionCredentialsMap != nil {
				connectionCredentialsArray = append(connectionCredentialsArray, connectionCredentialsMap)
			}
			result["connection_credentials"] = connectionCredentialsArray
		}

		if v.ConnectionString != nil {
			result["connection_string"] = []interface{}{DatabaseConnectionStringToMap(v.ConnectionString)}
		}
	default:
		log.Printf("[WARN] Received 'component_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToCloudDbSystemDiscoveryConnector(fieldKeyFormat string) (oci_database_management.CloudDbSystemDiscoveryConnector, error) {
	var baseObject oci_database_management.CloudDbSystemDiscoveryConnector
	//discriminator
	connectorTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector_type"))
	var connectorType string
	if ok {
		connectorType = connectorTypeRaw.(string)
	} else {
		connectorType = "" // default value
	}
	switch strings.ToLower(connectorType) {
	case strings.ToLower("MACS"):
		details := oci_database_management.CloudDbSystemDiscoveryMacsConnector{}
		if agentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_id")); ok {
			tmp := agentId.(string)
			details.AgentId = &tmp
		}
		if connectionInfo, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_info")); ok {
			if tmpList := connectionInfo.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connection_info"), 0)
				tmp, err := s.mapToCloudDbSystemConnectionInfo(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert connection_info, encountered error: %v", err)
				}
				details.ConnectionInfo = tmp
			}
		}
		if connectionFailureMessage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_failure_message")); ok {
			tmp := connectionFailureMessage.(string)
			details.ConnectionFailureMessage = &tmp
		}
		if connectionStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_status")); ok {
			tmp := connectionStatus.(string)
			details.ConnectionStatus = &tmp
		}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if timeConnectionStatusLastUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_connection_status_last_updated")); ok {
			tmp, err := time.Parse(time.RFC3339, timeConnectionStatusLastUpdated.(string))
			if err != nil {
				return details, err
			}
			details.TimeConnectionStatusLastUpdated = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown connector_type '%v' was specified", connectorType)
	}
	return baseObject, nil
}

func CloudDbSystemDiscoveryConnectorToMap(obj *oci_database_management.CloudDbSystemDiscoveryConnector) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.CloudDbSystemDiscoveryMacsConnector:
		result["connector_type"] = "MACS"

		if v.AgentId != nil {
			result["agent_id"] = string(*v.AgentId)
		}

		if v.ConnectionInfo != nil {
			connectionInfoArray := []interface{}{}
			if connectionInfoMap := CloudDbSystemConnectionInfoToMap(&v.ConnectionInfo); connectionInfoMap != nil {
				connectionInfoArray = append(connectionInfoArray, connectionInfoMap)
			}
			result["connection_info"] = connectionInfoArray
		}
	default:
		log.Printf("[WARN] Received 'connector_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func CloudDbSystemDiscoverySummaryToMap(obj oci_database_management.CloudDbSystemDiscoverySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DbaasParentInfrastructureId != nil {
		result["dbaas_parent_infrastructure_id"] = string(*obj.DbaasParentInfrastructureId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["deployment_type"] = string(obj.DeploymentType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToCloudListenerEndpoint(fieldKeyFormat string) (oci_database_management.CloudListenerEndpoint, error) {
	var baseObject oci_database_management.CloudListenerEndpoint
	//discriminator
	protocolRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol"))
	var protocol string
	if ok {
		protocol = protocolRaw.(string)
	} else {
		protocol = "" // default value
	}
	switch strings.ToLower(protocol) {
	case strings.ToLower("IPC"):
		details := oci_database_management.CloudListenerIpcEndpoint{}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if services, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "services")); ok {
			interfaces := services.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "services")) {
				details.Services = tmp
			}
		}
		baseObject = details
	case strings.ToLower("TCP"):
		details := oci_database_management.CloudListenerTcpEndpoint{}
		if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if services, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "services")); ok {
			interfaces := services.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "services")) {
				details.Services = tmp
			}
		}
		baseObject = details
	case strings.ToLower("TCPS"):
		details := oci_database_management.CloudListenerTcpsEndpoint{}
		if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if services, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "services")); ok {
			interfaces := services.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "services")) {
				details.Services = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown protocol '%v' was specified", protocol)
	}
	return baseObject, nil
}

func CloudListenerEndpointToMap(obj oci_database_management.CloudListenerEndpoint) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_management.CloudListenerIpcEndpoint:
		result["protocol"] = "IPC"

		if v.Key != nil {
			result["key"] = string(*v.Key)
		}

		result["services"] = v.Services
	case oci_database_management.CloudListenerTcpEndpoint:
		result["protocol"] = "TCP"

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		result["services"] = v.Services
	case oci_database_management.CloudListenerTcpsEndpoint:
		result["protocol"] = "TCPS"

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		result["services"] = v.Services
	default:
		log.Printf("[WARN] Received 'protocol' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToDatabaseConnectionCredentials(fieldKeyFormat string) (oci_database_management.DatabaseConnectionCredentials, error) {
	var baseObject oci_database_management.DatabaseConnectionCredentials
	//discriminator
	credentialTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type"))
	var credentialType string
	if ok {
		credentialType = credentialTypeRaw.(string)
	} else {
		credentialType = "DETAILS" // default value
	}
	switch strings.ToLower(credentialType) {
	case strings.ToLower("DETAILS"):
		details := oci_database_management.DatabaseConnectionCredentialsByDetails{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.DatabaseConnectionCredentialsByDetailsRoleEnum(role.(string))
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	case strings.ToLower("NAMED_CREDENTIAL"):
		details := oci_database_management.DatabaseNamedCredentialConnectionDetails{}
		if namedCredentialId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "named_credential_id")); ok {
			tmp := namedCredentialId.(string)
			details.NamedCredentialId = &tmp
		}
		baseObject = details
	case strings.ToLower("NAME_REFERENCE"):
		details := oci_database_management.DatabaseConnectionCredentailsByName{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		baseObject = details
	case strings.ToLower("SSL_DETAILS"):
		details := oci_database_management.DatabaseSslConnectionCredentials{}
		if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
			tmp := credentialName.(string)
			details.CredentialName = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
			details.Role = oci_database_management.DatabaseSslConnectionCredentialsRoleEnum(role.(string))
		}
		if sslSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_secret_id")); ok {
			tmp := sslSecretId.(string)
			details.SslSecretId = &tmp
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown credential_type '%v' was specified", credentialType)
	}
	return baseObject, nil
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToDatabaseConnectionString(fieldKeyFormat string) (oci_database_management.DatabaseConnectionString, error) {
	result := oci_database_management.DatabaseConnectionString{}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_database_management.DatabaseConnectionStringProtocolEnum(protocol.(string))
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		tmp := service.(string)
		result.Service = &tmp
	}

	return result, nil
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToDiscoveredCloudAsmInstance(fieldKeyFormat string) (oci_database_management.DiscoveredCloudAsmInstance, error) {
	result := oci_database_management.DiscoveredCloudAsmInstance{}

	if adrHomeDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "adr_home_directory")); ok {
		tmp := adrHomeDirectory.(string)
		result.AdrHomeDirectory = &tmp
	}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if instanceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_name")); ok {
		tmp := instanceName.(string)
		result.InstanceName = &tmp
	}

	return result, nil
}

func DiscoveredCloudAsmInstanceToMap(obj oci_database_management.DiscoveredCloudAsmInstance) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdrHomeDirectory != nil {
		result["adr_home_directory"] = string(*obj.AdrHomeDirectory)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.InstanceName != nil {
		result["instance_name"] = string(*obj.InstanceName)
	}

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToDiscoveredCloudClusterInstance(fieldKeyFormat string) (oci_database_management.DiscoveredCloudClusterInstance, error) {
	result := oci_database_management.DiscoveredCloudClusterInstance{}

	if adrHomeDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "adr_home_directory")); ok {
		tmp := adrHomeDirectory.(string)
		result.AdrHomeDirectory = &tmp
	}

	if clusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cluster_id")); ok {
		tmp := clusterId.(string)
		result.ClusterId = &tmp
	}

	if connector, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector")); ok {
		if tmpList := connector.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connector"), 0)
			tmp, err := s.mapToCloudDbSystemDiscoveryConnector(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert connector, encountered error: %v", err)
			}
			result.Connector = tmp
		}
	}

	if crsBaseDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "crs_base_directory")); ok {
		tmp := crsBaseDirectory.(string)
		result.CrsBaseDirectory = &tmp
	}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if nodeRole, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_role")); ok {
		result.NodeRole = oci_database_management.DiscoveredCloudClusterInstanceNodeRoleEnum(nodeRole.(string))
	}

	return result, nil
}

func DiscoveredCloudClusterInstanceToMap(obj oci_database_management.DiscoveredCloudClusterInstance) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdrHomeDirectory != nil {
		result["adr_home_directory"] = string(*obj.AdrHomeDirectory)
	}

	if obj.ClusterId != nil {
		result["cluster_id"] = string(*obj.ClusterId)
	}

	if obj.Connector != nil {
		connectorArray := []interface{}{}
		if connectorMap := CloudDbSystemDiscoveryConnectorToMap(&obj.Connector); connectorMap != nil {
			connectorArray = append(connectorArray, connectorMap)
		}
		result["connector"] = connectorArray
	}

	if obj.CrsBaseDirectory != nil {
		result["crs_base_directory"] = string(*obj.CrsBaseDirectory)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	result["node_role"] = string(obj.NodeRole)

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToDiscoveredCloudDbInstance(fieldKeyFormat string) (oci_database_management.DiscoveredCloudDbInstance, error) {
	result := oci_database_management.DiscoveredCloudDbInstance{}

	if adrHomeDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "adr_home_directory")); ok {
		tmp := adrHomeDirectory.(string)
		result.AdrHomeDirectory = &tmp
	}

	if hostName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host_name")); ok {
		tmp := hostName.(string)
		result.HostName = &tmp
	}

	if instanceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_name")); ok {
		tmp := instanceName.(string)
		result.InstanceName = &tmp
	}

	if nodeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_name")); ok {
		tmp := nodeName.(string)
		result.NodeName = &tmp
	}

	if oracleHome, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oracle_home")); ok {
		tmp := oracleHome.(string)
		result.OracleHome = &tmp
	}

	return result, nil
}

func DiscoveredCloudDbInstanceToMap(obj oci_database_management.DiscoveredCloudDbInstance) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdrHomeDirectory != nil {
		result["adr_home_directory"] = string(*obj.AdrHomeDirectory)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.InstanceName != nil {
		result["instance_name"] = string(*obj.InstanceName)
	}

	if obj.NodeName != nil {
		result["node_name"] = string(*obj.NodeName)
	}

	if obj.OracleHome != nil {
		result["oracle_home"] = string(*obj.OracleHome)
	}

	return result
}

func DiscoveredCloudDbSystemComponentToMap(obj oci_database_management.DiscoveredCloudDbSystemComponent) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_management.DiscoveredCloudAsm:
		result["component_type"] = "ASM"

		asmInstances := []interface{}{}
		for _, item := range v.AsmInstances {
			asmInstances = append(asmInstances, DiscoveredCloudAsmInstanceToMap(item))
		}
		result["asm_instances"] = asmInstances

		if v.Connector != nil {
			connectorArray := []interface{}{}
			if connectorMap := CloudDbSystemDiscoveryConnectorToMap(&v.Connector); connectorMap != nil {
				connectorArray = append(connectorArray, connectorMap)
			}
			result["connector"] = connectorArray
		}

		if v.GridHome != nil {
			result["grid_home"] = string(*v.GridHome)
		}

		if v.IsFlexEnabled != nil {
			result["is_flex_enabled"] = bool(*v.IsFlexEnabled)
		}

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudAsmInstance:
		result["component_type"] = "ASM_INSTANCE"

		if v.AdrHomeDirectory != nil {
			result["adr_home_directory"] = string(*v.AdrHomeDirectory)
		}

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		if v.InstanceName != nil {
			result["instance_name"] = string(*v.InstanceName)
		}

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudCluster:
		result["component_type"] = "CLUSTER"

		clusterInstances := []interface{}{}
		for _, item := range v.ClusterInstances {
			clusterInstances = append(clusterInstances, DiscoveredCloudClusterInstanceToMap(item))
		}
		result["cluster_instances"] = clusterInstances

		if v.GridHome != nil {
			result["grid_home"] = string(*v.GridHome)
		}

		if v.IsFlexCluster != nil {
			result["is_flex_cluster"] = bool(*v.IsFlexCluster)
		}

		networkConfigurations := []interface{}{}
		for _, item := range v.NetworkConfigurations {
			networkConfigurations = append(networkConfigurations, CloudClusterNetworkConfigurationToMap(item))
		}
		result["network_configurations"] = networkConfigurations

		if v.OcrFileLocation != nil {
			result["ocr_file_location"] = string(*v.OcrFileLocation)
		}

		scanConfigurations := []interface{}{}
		for _, item := range v.ScanConfigurations {
			scanConfigurations = append(scanConfigurations, CloudClusterScanListenerConfigurationToMap(item))
		}
		result["scan_configurations"] = scanConfigurations

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		vipConfigurations := []interface{}{}
		for _, item := range v.VipConfigurations {
			vipConfigurations = append(vipConfigurations, CloudClusterVipConfigurationToMap(item))
		}
		result["vip_configurations"] = vipConfigurations

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudClusterInstance:
		result["component_type"] = "CLUSTER_INSTANCE"

		if v.AdrHomeDirectory != nil {
			result["adr_home_directory"] = string(*v.AdrHomeDirectory)
		}

		if v.ClusterId != nil {
			result["cluster_id"] = string(*v.ClusterId)
		}

		if v.Connector != nil {
			connectorArray := []interface{}{}
			if connectorMap := CloudDbSystemDiscoveryConnectorToMap(&v.Connector); connectorMap != nil {
				connectorArray = append(connectorArray, connectorMap)
			}
			result["connector"] = connectorArray
		}

		if v.CrsBaseDirectory != nil {
			result["crs_base_directory"] = string(*v.CrsBaseDirectory)
		}

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		result["node_role"] = string(v.NodeRole)

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudDatabase:
		result["component_type"] = "DATABASE"

		if v.CanEnableAllCurrentPdbs != nil {
			result["can_enable_all_current_pdbs"] = bool(*v.CanEnableAllCurrentPdbs)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.Connector != nil {
			connectorArray := []interface{}{}
			if connectorMap := CloudDbSystemDiscoveryConnectorToMap(&v.Connector); connectorMap != nil {
				connectorArray = append(connectorArray, connectorMap)
			}
			result["connector"] = connectorArray
		}

		if v.DbEdition != nil {
			result["db_edition"] = string(*v.DbEdition)
		}

		if v.DbId != nil {
			result["db_id"] = string(*v.DbId)
		}

		dbInstances := []interface{}{}
		for _, item := range v.DbInstances {
			dbInstances = append(dbInstances, DiscoveredCloudDbInstanceToMap(item))
		}
		result["db_instances"] = dbInstances

		if v.DbPacks != nil {
			result["db_packs"] = string(*v.DbPacks)
		}

		result["db_role"] = string(v.DbRole)

		result["db_type"] = string(v.DbType)

		if v.DbUniqueName != nil {
			result["db_unique_name"] = string(*v.DbUniqueName)
		}

		if v.DbVersion != nil {
			result["db_version"] = string(*v.DbVersion)
		}

		if v.IsAutoEnablePluggableDatabase != nil {
			result["is_auto_enable_pluggable_database"] = bool(*v.IsAutoEnablePluggableDatabase)
		}

		if v.IsCluster != nil {
			result["is_cluster"] = bool(*v.IsCluster)
		}

		pluggableDatabases := []interface{}{}
		for _, item := range v.PluggableDatabases {
			pluggableDatabases = append(pluggableDatabases, DiscoveredCloudPluggableDatabaseToMap(item))
		}
		result["pluggable_databases"] = pluggableDatabases

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudDbHome:
		result["component_type"] = "DATABASE_HOME"

		if v.HomeDirectory != nil {
			result["home_directory"] = string(*v.HomeDirectory)
		}

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudDbInstance:
		result["component_type"] = "DATABASE_INSTANCE"

		if v.AdrHomeDirectory != nil {
			result["adr_home_directory"] = string(*v.AdrHomeDirectory)
		}

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		if v.InstanceName != nil {
			result["instance_name"] = string(*v.InstanceName)
		}

		if v.NodeName != nil {
			result["node_name"] = string(*v.NodeName)
		}

		if v.OracleHome != nil {
			result["oracle_home"] = string(*v.OracleHome)
		}

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudDbNode:
		result["component_type"] = "DATABASE_NODE"

		if v.Connector != nil {
			connectorArray := []interface{}{}
			if connectorMap := CloudDbSystemDiscoveryConnectorToMap(&v.Connector); connectorMap != nil {
				connectorArray = append(connectorArray, connectorMap)
			}
			result["connector"] = connectorArray
		}

		if v.CpuCoreCount != nil {
			result["cpu_core_count"] = float32(*v.CpuCoreCount)
		}

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		if v.MemorySizeInGBs != nil {
			result["memory_size_in_gbs"] = float32(*v.MemorySizeInGBs)
		}

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudListener:
		result["component_type"] = "LISTENER"

		if v.AdrHomeDirectory != nil {
			result["adr_home_directory"] = string(*v.AdrHomeDirectory)
		}

		if v.Connector != nil {
			connectorArray := []interface{}{}
			if connectorMap := CloudDbSystemDiscoveryConnectorToMap(&v.Connector); connectorMap != nil {
				connectorArray = append(connectorArray, connectorMap)
			}
			result["connector"] = connectorArray
		}

		if v.DbNodeName != nil {
			result["db_node_name"] = string(*v.DbNodeName)
		}

		endpoints := []interface{}{}
		for _, item := range v.Endpoints {
			endpoints = append(endpoints, CloudListenerEndpointToMap(item))
		}
		result["endpoints"] = endpoints

		if v.HostName != nil {
			result["host_name"] = string(*v.HostName)
		}

		if v.ListenerAlias != nil {
			result["listener_alias"] = string(*v.ListenerAlias)
		}

		result["listener_type"] = string(v.ListenerType)

		if v.LogDirectory != nil {
			result["log_directory"] = string(*v.LogDirectory)
		}

		if v.OracleHome != nil {
			result["oracle_home"] = string(*v.OracleHome)
		}

		if v.TraceDirectory != nil {
			result["trace_directory"] = string(*v.TraceDirectory)
		}

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	case oci_database_management.DiscoveredCloudPluggableDatabase:
		result["component_type"] = "PLUGGABLE_DATABASE"

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.Connector != nil {
			connectorArray := []interface{}{}
			if connectorMap := CloudDbSystemDiscoveryConnectorToMap(&v.Connector); connectorMap != nil {
				connectorArray = append(connectorArray, connectorMap)
			}
			result["connector"] = connectorArray
		}

		if v.ContainerDatabaseId != nil {
			result["container_database_id"] = string(*v.ContainerDatabaseId)
		}

		if v.Guid != nil {
			result["guid"] = string(*v.Guid)
		}

		associatedComponents := []interface{}{}
		for _, item := range v.AssociatedComponents {
			associatedComponents = append(associatedComponents, AssociatedCloudComponentToMap(item))
		}
		result["associated_components"] = associatedComponents

		if v.ComponentId != nil {
			result["component_id"] = string(*v.ComponentId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.DbaasId != nil {
			result["dbaas_id"] = string(*v.DbaasId)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.IsSelectedForMonitoring != nil {
			result["is_selected_for_monitoring"] = bool(*v.IsSelectedForMonitoring)
		}

		if v.ResourceId != nil {
			result["resource_id"] = string(*v.ResourceId)
		}

		result["status"] = string(v.Status)
	default:
		log.Printf("[WARN] Received 'component_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToDiscoveredCloudPluggableDatabase(fieldKeyFormat string) (oci_database_management.DiscoveredCloudPluggableDatabase, error) {
	result := oci_database_management.DiscoveredCloudPluggableDatabase{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if connector, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector")); ok {
		if tmpList := connector.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connector"), 0)
			tmp, err := s.mapToCloudDbSystemDiscoveryConnector(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert connector, encountered error: %v", err)
			}
			result.Connector = tmp
		}
	}

	if containerDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_database_id")); ok {
		tmp := containerDatabaseId.(string)
		result.ContainerDatabaseId = &tmp
	}

	if guid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "guid")); ok {
		tmp := guid.(string)
		result.Guid = &tmp
	}

	return result, nil
}

func DiscoveredCloudPluggableDatabaseToMap(obj oci_database_management.DiscoveredCloudPluggableDatabase) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Connector != nil {
		connectorArray := []interface{}{}
		if connectorMap := CloudDbSystemDiscoveryConnectorToMap(&obj.Connector); connectorMap != nil {
			connectorArray = append(connectorArray, connectorMap)
		}
		result["connector"] = connectorArray
	}

	if obj.ContainerDatabaseId != nil {
		result["container_database_id"] = string(*obj.ContainerDatabaseId)
	}

	if obj.Guid != nil {
		result["guid"] = string(*obj.Guid)
	}

	return result
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_database_management.PatchInstruction, error) {
	var baseObject oci_database_management.PatchInstruction
	//discriminator
	operationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation"))
	var operation string
	if ok {
		operation = operationRaw.(string)
	} else {
		operation = "" // default value
	}
	switch strings.ToLower(operation) {
	case strings.ToLower("MERGE"):
		details := oci_database_management.PatchMergeInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			if tmpList := value.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "value"), 0)
				tmp, err := s.mapToPatchMergeInstructionValue(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert value, encountered error: %v", err)
				}
				details.Value = tmp
			}
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown operation '%v' was specified", operation)
	}
	return baseObject, nil
}

func (s *DatabaseManagementCloudDbSystemDiscoveryResourceCrud) mapToPatchMergeInstructionValue(fieldKeyFormat string) (*interface{}, error) {
	var result interface{}
	value := make(map[string]interface{})

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		value["displayName"] = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		value["compartmentId"] = &tmp
	}

	if isSelectedForMonitoring, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_selected_for_monitoring")); ok {
		tmp := isSelectedForMonitoring.(bool)
		value["isSelectedForMonitoring"] = &tmp
	}

	if connector, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector")); ok {
		if tmpList := connector.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connector"), 0)
			tmp, err := s.mapToCloudDbSystemDiscoveryConnector(fieldKeyFormatNextLevel)
			if err != nil {
				return nil, fmt.Errorf("unable to convert connector, encountered error: %v", err)
			}
			value["connector"] = tmp
		}
	}

	result = value

	return &result, nil
}
