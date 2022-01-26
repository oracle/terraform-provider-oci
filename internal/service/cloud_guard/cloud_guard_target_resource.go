// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v56/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func CloudGuardTargetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardTarget,
		Read:     readCloudGuardTarget,
		Update:   updateCloudGuardTarget,
		Delete:   deleteCloudGuardTarget,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_resource_type": {
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
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_detector_recipes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"detector_recipe_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"detector_rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"details": {
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"condition_groups": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"compartment_id": {
																Type:     schema.TypeString,
																Required: true,
															},
															"condition": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: utils.JsonStringDiffSuppressFunction,
															},

															// Optional

															// Computed
														},
													},
												},

												// Computed
												"configurations": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"config_key": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"values": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"list_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"managed_list_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"is_configuration_allowed": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"labels": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"risk_level": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"detector_rule_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"detector": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_list_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"recommendation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
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
							},
						},

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"detector": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"effective_detector_rules": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"condition_groups": {
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
															"condition": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"configurations": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"config_key": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"values": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"list_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"managed_list_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"is_configuration_allowed": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"labels": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"risk_level": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"detector": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"detector_rule_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_list_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"recommendation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
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
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
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
				},
			},
			"target_responder_recipes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"responder_recipe_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"responder_rules": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"details": {
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"condition": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													DiffSuppressFunc: utils.JsonStringDiffSuppressFunction,
												},
												"configurations": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"config_key": {
																Type:     schema.TypeString,
																Required: true,
															},
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"value": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},
												"mode": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
												"is_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"responder_rule_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"policies": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"supported_modes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"effective_responder_rules": {
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
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"condition": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"configurations": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"config_key": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"is_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"policies": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"responder_rule_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"supported_modes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner": {
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
				},
			},

			// Computed
			"inherited_by_compartments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"lifecyle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recipe_count": {
				Type:     schema.TypeInt,
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

func createCloudGuardTarget(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardTarget(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardTarget(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardTarget(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardTargetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.Target
	DisableNotFoundRetries bool
}

func (s *CloudGuardTargetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardTargetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateCreating),
	}
}

func (s *CloudGuardTargetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateActive),
	}
}

func (s *CloudGuardTargetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleting),
	}
}

func (s *CloudGuardTargetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleted),
	}
}

func (s *CloudGuardTargetResourceCrud) Create() error {
	request := oci_cloud_guard.CreateTargetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_guard.LifecycleStateEnum(state.(string))
	}

	if targetDetectorRecipes, ok := s.D.GetOkExists("target_detector_recipes"); ok {
		interfaces := targetDetectorRecipes.([]interface{})
		tmp := make([]oci_cloud_guard.CreateTargetDetectorRecipeDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_detector_recipes", stateDataIndex)
			converted, err := s.mapToCreateTargetDetectorRecipeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("target_detector_recipes") {
			request.TargetDetectorRecipes = tmp
		}
	}

	if targetResourceId, ok := s.D.GetOkExists("target_resource_id"); ok {
		tmp := targetResourceId.(string)
		request.TargetResourceId = &tmp
	}

	if targetResourceType, ok := s.D.GetOkExists("target_resource_type"); ok {
		request.TargetResourceType = oci_cloud_guard.TargetResourceTypeEnum(targetResourceType.(string))
	}

	if targetResponderRecipes, ok := s.D.GetOkExists("target_responder_recipes"); ok {
		interfaces := targetResponderRecipes.([]interface{})
		tmp := make([]oci_cloud_guard.CreateTargetResponderRecipeDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_responder_recipes", stateDataIndex)
			converted, err := s.mapToCreateTargetResponderRecipeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("target_responder_recipes") {
			request.TargetResponderRecipes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Target
	return nil
}

func (s *CloudGuardTargetResourceCrud) Get() error {
	request := oci_cloud_guard.GetTargetRequest{}

	tmp := s.D.Id()
	request.TargetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Target
	return nil
}

func (s *CloudGuardTargetResourceCrud) Update() error {
	request := oci_cloud_guard.UpdateTargetRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_guard.LifecycleStateEnum(state.(string))
	}

	if targetDetectorRecipes, ok := s.D.GetOkExists("target_detector_recipes"); ok {
		interfaces := targetDetectorRecipes.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateTargetDetectorRecipe, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_detector_recipes", stateDataIndex)
			converted, err := s.mapToUpdateTargetDetectorRecipe(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("target_detector_recipes") {
			request.TargetDetectorRecipes = tmp
		}
	}

	tmp := s.D.Id()
	request.TargetId = &tmp

	if targetResponderRecipes, ok := s.D.GetOkExists("target_responder_recipes"); ok {
		interfaces := targetResponderRecipes.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateTargetResponderRecipe, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_responder_recipes", stateDataIndex)
			converted, err := s.mapToUpdateTargetResponderRecipe(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("target_responder_recipes") {
			request.TargetResponderRecipes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Target
	return nil
}

func (s *CloudGuardTargetResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteTargetRequest{}

	tmp := s.D.Id()
	request.TargetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.DeleteTarget(context.Background(), request)
	return err
}

func (s *CloudGuardTargetResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	s.D.Set("inherited_by_compartments", s.Res.InheritedByCompartments)

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	if s.Res.RecipeCount != nil {
		s.D.Set("recipe_count", *s.Res.RecipeCount)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	targetDetectorRecipes := []interface{}{}
	for _, item := range s.Res.TargetDetectorRecipes {
		targetDetectorRecipes = append(targetDetectorRecipes, TargetDetectorRecipeToMap(item))
	}
	s.D.Set("target_detector_recipes", targetDetectorRecipes)

	if s.Res.TargetResourceId != nil {
		s.D.Set("target_resource_id", *s.Res.TargetResourceId)
	}

	s.D.Set("target_resource_type", s.Res.TargetResourceType)

	targetResponderRecipes := []interface{}{}
	for _, item := range s.Res.TargetResponderRecipes {
		targetResponderRecipes = append(targetResponderRecipes, TargetResponderRecipeToMap(item))
	}
	s.D.Set("target_responder_recipes", targetResponderRecipes)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ConditionGroupToMap(obj oci_cloud_guard.ConditionGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Condition != nil {
		condition, err := tfresource.ConvertObjectToJsonString(obj.Condition)
		if err == nil {
			result["condition"] = condition
		}
	}

	return result
}

func (s *CloudGuardTargetResourceCrud) mapToCreateTargetDetectorRecipeDetails(fieldKeyFormat string) (oci_cloud_guard.CreateTargetDetectorRecipeDetails, error) {
	result := oci_cloud_guard.CreateTargetDetectorRecipeDetails{}

	if detectorRecipeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "detector_recipe_id")); ok {
		tmp := detectorRecipeId.(string)
		result.DetectorRecipeId = &tmp
	}

	if detectorRules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "detector_rules")); ok {
		interfaces := detectorRules.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateTargetRecipeDetectorRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "detector_rules"), stateDataIndex)
			converted, err := s.mapToUpdateTargetRecipeDetectorRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "detector_rules")) {
			result.DetectorRules = tmp
		}
	}

	return result, nil
}

func TargetDetectorRecipeToMap(obj oci_cloud_guard.TargetDetectorRecipe) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["detector"] = string(obj.Detector)

	if obj.DetectorRecipeId != nil {
		result["detector_recipe_id"] = string(*obj.DetectorRecipeId)
	}

	detectorRules := []interface{}{}
	for _, item := range obj.DetectorRules {
		detectorRules = append(detectorRules, TargetDetectorRecipeDetectorRuleToMap(item))
	}
	result["detector_rules"] = detectorRules

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	effectiveDetectorRules := []interface{}{}
	for _, item := range obj.EffectiveDetectorRules {
		effectiveDetectorRules = append(effectiveDetectorRules, TargetDetectorRecipeDetectorRuleToMap(item))
	}
	result["effective_detector_rules"] = effectiveDetectorRules

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["owner"] = string(obj.Owner)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *CloudGuardTargetResourceCrud) mapToCreateTargetResponderRecipeDetails(fieldKeyFormat string) (oci_cloud_guard.CreateTargetResponderRecipeDetails, error) {
	result := oci_cloud_guard.CreateTargetResponderRecipeDetails{}

	if responderRecipeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "responder_recipe_id")); ok {
		tmp := responderRecipeId.(string)
		result.ResponderRecipeId = &tmp
	}

	if responderRules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "responder_rules")); ok {
		interfaces := responderRules.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateTargetRecipeResponderRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "responder_rules"), stateDataIndex)
			converted, err := s.mapToUpdateTargetRecipeResponderRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "responder_rules")) {
			result.ResponderRules = tmp
		}
	}

	return result, nil
}

func TargetResponderRecipeToMap(obj oci_cloud_guard.TargetResponderRecipe) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	effectiveResponderRules := []interface{}{}
	for _, item := range obj.EffectiveResponderRules {
		effectiveResponderRules = append(effectiveResponderRules, TargetResponderRecipeResponderRuleToMap(item))
	}
	result["effective_responder_rules"] = effectiveResponderRules

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["owner"] = string(obj.Owner)

	if obj.ResponderRecipeId != nil {
		result["responder_recipe_id"] = string(*obj.ResponderRecipeId)
	}

	responderRules := []interface{}{}
	for _, item := range obj.ResponderRules {
		responderRules = append(responderRules, TargetResponderRecipeResponderRuleToMap(item))
	}
	result["responder_rules"] = responderRules

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *CloudGuardTargetResourceCrud) mapToResponderConfiguration(fieldKeyFormat string) (oci_cloud_guard.ResponderConfiguration, error) {
	result := oci_cloud_guard.ResponderConfiguration{}

	if configKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_key")); ok {
		tmp := configKey.(string)
		result.ConfigKey = &tmp
	}

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

func TargetDetectorDetailsToMap(obj *oci_cloud_guard.TargetDetectorDetails) map[string]interface{} {
	result := map[string]interface{}{}

	conditionGroups := []interface{}{}
	for _, item := range obj.ConditionGroups {
		conditionGroups = append(conditionGroups, ConditionGroupToMap(item))
	}
	result["condition_groups"] = conditionGroups

	configurations := []interface{}{}
	for _, item := range obj.Configurations {
		configurations = append(configurations, DetectorConfigurationToMap(item))
	}
	result["configurations"] = configurations

	if obj.IsConfigurationAllowed != nil {
		result["is_configuration_allowed"] = bool(*obj.IsConfigurationAllowed)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	result["labels"] = obj.Labels

	result["risk_level"] = string(obj.RiskLevel)

	return result
}

func TargetDetectorRecipeDetectorRuleToMap(obj oci_cloud_guard.TargetDetectorRecipeDetectorRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Details != nil {
		result["details"] = []interface{}{TargetDetectorDetailsToMap(obj.Details)}
	}

	result["detector"] = string(obj.Detector)

	if obj.DetectorRuleId != nil {
		result["detector_rule_id"] = string(*obj.DetectorRuleId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["managed_list_types"] = obj.ManagedListTypes

	if obj.Recommendation != nil {
		result["recommendation"] = string(*obj.Recommendation)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.ServiceType != nil {
		result["service_type"] = string(*obj.ServiceType)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func TargetResponderRecipeResponderRuleToMap(obj oci_cloud_guard.TargetResponderRecipeResponderRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Details != nil {
		result["details"] = []interface{}{ResponderRuleDetailsToMap(obj.Details)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["policies"] = obj.Policies

	if obj.ResponderRuleId != nil {
		result["responder_rule_id"] = string(*obj.ResponderRuleId)
	}

	result["state"] = string(obj.LifecycleState)

	result["supported_modes"] = obj.SupportedModes

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func TargetSummaryToMap(obj oci_cloud_guard.TargetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecyleDetails != nil {
		result["lifecyle_details"] = string(*obj.LifecyleDetails)
	}

	if obj.RecipeCount != nil {
		result["recipe_count"] = int(*obj.RecipeCount)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetResourceId != nil {
		result["target_resource_id"] = string(*obj.TargetResourceId)
	}

	result["target_resource_type"] = string(obj.TargetResourceType)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *CloudGuardTargetResourceCrud) mapToUpdateTargetDetectorRuleDetails(fieldKeyFormat string) (oci_cloud_guard.UpdateTargetDetectorRuleDetails, error) {
	result := oci_cloud_guard.UpdateTargetDetectorRuleDetails{}

	if conditionGroups, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition_groups")); ok {
		interfaces := conditionGroups.([]interface{})
		tmp := make([]oci_cloud_guard.ConditionGroup, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "condition_groups"), stateDataIndex)
			converted, err := s.mapToConditionGroup(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "condition_groups")) {
			result.ConditionGroups = tmp
		}
	}

	return result, nil
}

func (s *CloudGuardTargetResourceCrud) mapToUpdateTargetRecipeDetectorRuleDetails(fieldKeyFormat string) (oci_cloud_guard.UpdateTargetRecipeDetectorRuleDetails, error) {
	result := oci_cloud_guard.UpdateTargetRecipeDetectorRuleDetails{}

	if details, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "details")); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "details"), 0)
			tmp, err := s.mapToUpdateTargetDetectorRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert details, encountered error: %v", err)
			}
			result.Details = &tmp
		}
	}

	if detectorRuleId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "detector_rule_id")); ok {
		tmp := detectorRuleId.(string)
		result.DetectorRuleId = &tmp
	}

	return result, nil
}

func (s *CloudGuardTargetResourceCrud) mapToUpdateTargetRecipeResponderRuleDetails(fieldKeyFormat string) (oci_cloud_guard.UpdateTargetRecipeResponderRuleDetails, error) {
	result := oci_cloud_guard.UpdateTargetRecipeResponderRuleDetails{}

	if details, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "details")); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "details"), 0)
			tmp, err := s.mapToUpdateTargetResponderRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert details, encountered error: %v", err)
			}
			result.Details = &tmp
		}
	}

	if responderRuleId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "responder_rule_id")); ok {
		tmp := responderRuleId.(string)
		result.ResponderRuleId = &tmp
	}

	return result, nil
}

func (s *CloudGuardTargetResourceCrud) mapToUpdateTargetResponderRuleDetails(fieldKeyFormat string) (oci_cloud_guard.UpdateTargetResponderRuleDetails, error) {
	result := oci_cloud_guard.UpdateTargetResponderRuleDetails{}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		if len(tmp) > 0 {
			var err error
			result.Condition, err = jsonToCondition(tmp)
			if err != nil {
				return result, err
			}
		}
	}

	if configurations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configurations")); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_cloud_guard.ResponderConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configurations"), stateDataIndex)
			converted, err := s.mapToResponderConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "configurations")) {
			result.Configurations = tmp
		}
	}

	if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
		result.Mode = oci_cloud_guard.ResponderModeTypesEnum(mode.(string))
	}

	return result, nil
}

func (s *CloudGuardTargetResourceCrud) mapToUpdateTargetResponderRecipe(fieldKeyFormat string) (oci_cloud_guard.UpdateTargetResponderRecipe, error) {
	result := oci_cloud_guard.UpdateTargetResponderRecipe{}

	if responderRecipeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := responderRecipeId.(string)
		result.TargetResponderRecipeId = &tmp
	}

	if responderRules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "responder_rules")); ok {
		interfaces := responderRules.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateTargetRecipeResponderRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "responder_rules"), stateDataIndex)
			converted, err := s.mapToUpdateTargetRecipeResponderRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "responder_rules")) {
			result.ResponderRules = tmp
		}
	}

	return result, nil
}

func (s *CloudGuardTargetResourceCrud) mapToUpdateTargetDetectorRecipe(fieldKeyFormat string) (oci_cloud_guard.UpdateTargetDetectorRecipe, error) {
	result := oci_cloud_guard.UpdateTargetDetectorRecipe{}

	if detectorRecipeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := detectorRecipeId.(string)
		result.TargetDetectorRecipeId = &tmp
	}

	if detectorRules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "detector_rules")); ok {
		interfaces := detectorRules.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateTargetRecipeDetectorRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "detector_rules"), stateDataIndex)
			converted, err := s.mapToUpdateTargetRecipeDetectorRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "detector_rules")) {
			result.DetectorRules = tmp
		}
	}

	return result, nil
}

func (s *CloudGuardTargetResourceCrud) mapToConditionGroup(fieldKeyFormat string) (oci_cloud_guard.ConditionGroup, error) {
	result := oci_cloud_guard.ConditionGroup{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		if len(tmp) > 0 {
			var err error
			result.Condition, err = jsonToCondition(tmp)
			if err != nil {
				return result, err
			}
		}
	}

	return result, nil
}
