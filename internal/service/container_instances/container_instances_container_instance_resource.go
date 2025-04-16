// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package container_instances

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_container_instances "github.com/oracle/oci-go-sdk/v65/containerinstances"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerInstancesContainerInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createContainerInstancesContainerInstance,
		Read:     readContainerInstancesContainerInstance,
		Update:   updateContainerInstancesContainerInstance,
		Delete:   deleteContainerInstancesContainerInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"containers": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"image_url": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"arguments": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"command": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
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
						"environment_variables": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem:     schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"health_checks": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"health_check_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"HTTP",
											"TCP",
										}, true),
									},
									// Optional
									"failure_action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"failure_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"headers": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"initial_delay_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"interval_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"status": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"status_details": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"success_threshold": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"timeout_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"is_resource_principal_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"resource_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"memory_limit_in_gbs": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vcpus_limit": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"security_context": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"capabilities": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"add_capabilities": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"drop_capabilities": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												// Computed
											},
										},
									},
									"is_non_root_user_check_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_root_file_system_readonly": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"run_as_group": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"run_as_user": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"security_context_type": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"LINUX",
										}, true),
									},

									// Computed
								},
							},
						},
						"volume_mounts": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"mount_path": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"volume_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"is_read_only": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"partition": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"sub_path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"working_directory": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"container_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"container_instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"exit_code": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
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
						"time_terminated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape_config": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"ocpus": {
							Type:     schema.TypeFloat,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"memory_in_gbs": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"networking_bandwidth_in_gbps": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"processor_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"vnics": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
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
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: func(k, o, l string, d *schema.ResourceData) bool { return true }, // this is in place due to a known issue, to be removed after fix
							Elem:             schema.TypeString,
						},
						"hostname_label": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_public_ip_assigned": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"skip_source_dest_check": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"vnic_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Optional
			"container_restart_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"dns_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"nameservers": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"searches": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"fault_domain": {
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
			"graceful_shutdown_timeout_in_seconds": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"image_pull_secrets": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"registry_endpoint": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"secret_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BASIC",
								"VAULT",
							}, true),
						},

						// Optional
						"password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"username": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_container_instances.ContainerInstanceLifecycleStateInactive),
					string(oci_container_instances.ContainerInstanceLifecycleStateActive),
				}, true),
			},
			"volumes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"volume_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CONFIGFILE",
								"EMPTYDIR",
							}, true),
						},

						// Optional
						"backing_store": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"configs": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"data": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"file_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
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
			"container_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
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
			"volume_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createContainerInstancesContainerInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerInstancesContainerInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerInstanceClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_container_instances.ContainerInstanceLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_container_instances.ContainerInstanceLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopContainerInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_container_instances.ContainerInstanceLifecycleStateInactive)
	}
	return nil

}

func readContainerInstancesContainerInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerInstancesContainerInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerInstanceClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateContainerInstancesContainerInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerInstancesContainerInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerInstanceClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_container_instances.ContainerInstanceLifecycleStateActive == oci_container_instances.ContainerInstanceLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_container_instances.ContainerInstanceLifecycleStateInactive == oci_container_instances.ContainerInstanceLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartContainerInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_container_instances.ContainerInstanceLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopContainerInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_container_instances.ContainerInstanceLifecycleStateInactive)
	}

	return nil
}

func deleteContainerInstancesContainerInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerInstancesContainerInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerInstanceClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ContainerInstancesContainerInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_container_instances.ContainerInstanceClient
	VirtualNetworkClient   *oci_core.VirtualNetworkClient
	Res                    *oci_container_instances.ContainerInstance
	DisableNotFoundRetries bool
}

func (s *ContainerInstancesContainerInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ContainerInstancesContainerInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_container_instances.ContainerInstanceLifecycleStateCreating),
	}
}

func (s *ContainerInstancesContainerInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_container_instances.ContainerInstanceLifecycleStateActive),
	}
}

func (s *ContainerInstancesContainerInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_container_instances.ContainerInstanceLifecycleStateDeleting),
	}
}

func (s *ContainerInstancesContainerInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_container_instances.ContainerInstanceLifecycleStateDeleted),
	}
}

func (s *ContainerInstancesContainerInstanceResourceCrud) Create() error {
	request := oci_container_instances.CreateContainerInstanceRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if containerRestartPolicy, ok := s.D.GetOkExists("container_restart_policy"); ok {
		request.ContainerRestartPolicy = oci_container_instances.ContainerInstanceContainerRestartPolicyEnum(containerRestartPolicy.(string))
	}

	if containers, ok := s.D.GetOkExists("containers"); ok {
		interfaces := containers.([]interface{})
		tmp := make([]oci_container_instances.CreateContainerDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "containers", stateDataIndex)
			converted, err := s.mapToCreateContainerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("containers") {
			request.Containers = tmp
		}
	}

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

	if dnsConfig, ok := s.D.GetOkExists("dns_config"); ok {
		if tmpList := dnsConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dns_config", 0)
			tmp, err := s.mapToCreateContainerDnsConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DnsConfig = &tmp
		}
	}

	if faultDomain, ok := s.D.GetOkExists("fault_domain"); ok {
		tmp := faultDomain.(string)
		request.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gracefulShutdownTimeoutInSeconds, ok := s.D.GetOkExists("graceful_shutdown_timeout_in_seconds"); ok {
		tmp := gracefulShutdownTimeoutInSeconds.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert gracefulShutdownTimeoutInSeconds string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.GracefulShutdownTimeoutInSeconds = &tmpInt64
	}

	if imagePullSecrets, ok := s.D.GetOkExists("image_pull_secrets"); ok {
		interfaces := imagePullSecrets.([]interface{})
		tmp := make([]oci_container_instances.CreateImagePullSecretDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_pull_secrets", stateDataIndex)
			converted, err := s.mapToCreateImagePullSecretDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("image_pull_secrets") {
			request.ImagePullSecrets = tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if shapeConfig, ok := s.D.GetOkExists("shape_config"); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shape_config", 0)
			tmp, err := s.mapToCreateContainerInstanceShapeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ShapeConfig = &tmp
		}
	}

	if vnics, ok := s.D.GetOkExists("vnics"); ok {
		interfaces := vnics.([]interface{})
		tmp := make([]oci_container_instances.CreateContainerVnicDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vnics", stateDataIndex)
			converted, err := s.mapToCreateContainerVnicDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("vnics") {
			request.Vnics = tmp
		}
	}

	if volumes, ok := s.D.GetOkExists("volumes"); ok {
		interfaces := volumes.([]interface{})
		tmp := make([]oci_container_instances.CreateContainerVolumeDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "volumes", stateDataIndex)
			converted, err := s.mapToCreateContainerVolumeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("volumes") {
			request.Volumes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")
	response, err := s.Client.CreateContainerInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	return s.getContainerInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance"), oci_container_instances.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ContainerInstancesContainerInstanceResourceCrud) getContainerInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_container_instances.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	containerInstanceId, err := containerInstanceWaitForWorkRequest(workId, "containerinstance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*containerInstanceId)

	return s.Get()
}

func (s *ContainerInstancesContainerInstanceResourceCrud) getContainerFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_container_instances.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	_, err := containerInstanceWaitForWorkRequest(workId, "container",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	return nil
}

func containerInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "containerinstance", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_container_instances.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func containerInstanceWaitForWorkRequest(wId *string, entityType string, action oci_container_instances.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_container_instances.ContainerInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerinstance")
	retryPolicy.ShouldRetryOperation = containerInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_container_instances.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_container_instances.OperationStatusInProgress),
			string(oci_container_instances.OperationStatusAccepted),
			string(oci_container_instances.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_container_instances.OperationStatusSucceeded),
			string(oci_container_instances.OperationStatusFailed),
			string(oci_container_instances.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_container_instances.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_container_instances.OperationStatusFailed || response.Status == oci_container_instances.OperationStatusCanceled {
		return nil, getErrorFromContainerInstancesContainerInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromContainerInstancesContainerInstanceWorkRequest(client *oci_container_instances.ContainerInstanceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_container_instances.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_container_instances.ListWorkRequestErrorsRequest{
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

func (s *ContainerInstancesContainerInstanceResourceCrud) Get() error {
	request := oci_container_instances.GetContainerInstanceRequest{}

	tmp := s.D.Id()
	request.ContainerInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

	response, err := s.Client.GetContainerInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerInstance
	return nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_container_instances.UpdateContainerInstanceRequest{}

	tmp := s.D.Id()
	request.ContainerInstanceId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

	if vnics, ok := s.D.GetOkExists("vnics"); ok && s.D.HasChange("vnics") {

		interfaces := vnics.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vnics", stateDataIndex)

			updateVnicDetails, err := s.mapToUpdateVnicDetailsInstance(fieldKeyFormat)
			if err != nil {
				return err
			}

			vnic_id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_id"))
			str_vnic_id := vnic_id.(string)
			if ok {
				request := oci_core.UpdateVnicRequest{
					VnicId:            &str_vnic_id,
					UpdateVnicDetails: updateVnicDetails,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
					},
				}

				response, err := s.VirtualNetworkClient.UpdateVnic(context.Background(), request)

				if err != nil {
					log.Printf("[ERROR] Primary VNIC could not be updated during container instance Update: %q (ContainerInstance ID: \"%v\", State: %q)", err, *s.Res.Id, s.Res.LifecycleState)
					return err
				}

				log.Printf("[INFO] Container Instances Vnic Update Work Request Executed. (OpcRequestId: \"%v\", VnicId: \"%v\")", *response.OpcRequestId, *request.VnicId)
			}
		}
	}

	if containers, ok := s.D.GetOkExists("containers"); ok && s.D.HasChange("containers") {
		interfaces := containers.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "containers", stateDataIndex)

			if _, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_id")); ok {

				request, err := s.mapToUpdateContainer(fieldKeyFormat)
				if err != nil {
					return err
				}

				request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

				response, err := s.Client.UpdateContainer(context.Background(), request)
				if err != nil {
					return err
				}

				workId := response.OpcWorkRequestId

				err = s.getContainerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance"), oci_container_instances.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
				if err != nil {
					return err
				}

				log.Printf("[INFO] Container Instances Container Update Work Request Executed. (OpcWorkRequestId: \"%v\", ContainerId: \"%v\")", *workId, *request.ContainerId)
			}
		}
	}

	response, err := s.Client.UpdateContainerInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	return s.getContainerInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance"), oci_container_instances.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToUpdateVnicDetailsInstance(fieldKeyFormat string) (oci_core.UpdateVnicDetails, error) {
	result := oci_core.UpdateVnicDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok && hostnameLabel != "" {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	result.NsgIds = []string{}
	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.NsgIds = tmp
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	return result, nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) Delete() error {
	request := oci_container_instances.DeleteContainerInstanceRequest{}

	tmp := s.D.Id()
	request.ContainerInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

	response, err := s.Client.DeleteContainerInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := containerInstanceWaitForWorkRequest(workId, "containerinstance",
		oci_container_instances.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ContainerInstancesContainerInstanceResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContainerCount != nil {
		s.D.Set("container_count", *s.Res.ContainerCount)
	}

	s.D.Set("container_restart_policy", s.Res.ContainerRestartPolicy)

	containers := []interface{}{}
	for _, item := range s.Res.Containers {
		result := map[string]interface{}{}

		if item.ContainerId != nil {
			result["container_id"] = string(*item.ContainerId)

			request := oci_container_instances.GetContainerRequest{}
			request.ContainerId = item.ContainerId

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

			response, err := s.Client.GetContainer(context.Background(), request)
			if err != nil {
				return err
			}

			container := response.Container
			result = ContainerToMap(container)
		}
		containers = append(containers, result)
	}
	s.D.Set("containers", containers)

	vnics := []interface{}{}
	for _, item := range s.Res.Vnics {

		result := map[string]interface{}{}

		if item.VnicId != nil {
			result["vnic_id"] = string(*item.VnicId)

			request := oci_core.GetVnicRequest{}
			request.VnicId = item.VnicId

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

			response, err := s.VirtualNetworkClient.GetVnic(context.Background(), request)
			if err != nil {
				return err
			}

			vnic := response.Vnic
			result = VnicDetailsToMap(vnic, false)
		}
		vnics = append(vnics, result)
	}
	s.D.Set("vnics", vnics)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsConfig != nil {
		s.D.Set("dns_config", []interface{}{ContainerDnsConfigToMap(s.Res.DnsConfig)})
	} else {
		s.D.Set("dns_config", nil)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GracefulShutdownTimeoutInSeconds != nil {
		s.D.Set("graceful_shutdown_timeout_in_seconds", strconv.FormatInt(*s.Res.GracefulShutdownTimeoutInSeconds, 10))
	}

	imagePullSecrets := []interface{}{}
	for index, item := range s.Res.ImagePullSecrets {
		imagePullSecrets = append(imagePullSecrets, ImagePullSecretToMap(item, s.D, index))
	}
	s.D.Set("image_pull_secrets", imagePullSecrets)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	if s.Res.ShapeConfig != nil {
		s.D.Set("shape_config", []interface{}{ContainerInstanceShapeConfigToMap(s.Res.ShapeConfig)})
	} else {
		s.D.Set("shape_config", nil)
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

	if s.Res.VolumeCount != nil {
		s.D.Set("volume_count", *s.Res.VolumeCount)
	}

	volumes := []interface{}{}
	for index, item := range s.Res.Volumes {
		volumes = append(volumes, ContainerVolumeToMap(item, s.D, index))
	}
	s.D.Set("volumes", volumes)

	return nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) StartContainerInstance() error {
	request := oci_container_instances.StartContainerInstanceRequest{}

	idTmp := s.D.Id()
	request.ContainerInstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

	_, err := s.Client.StartContainerInstance(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res.LifecycleState == oci_container_instances.ContainerInstanceLifecycleStateActive
	}
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerInstancesContainerInstanceResourceCrud) StopContainerInstance() error {
	request := oci_container_instances.StopContainerInstanceRequest{}

	idTmp := s.D.Id()
	request.ContainerInstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

	_, err := s.Client.StopContainerInstance(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res.LifecycleState == oci_container_instances.ContainerInstanceLifecycleStateInactive
	}
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToContainerCapabilities(fieldKeyFormat string) (oci_container_instances.ContainerCapabilities, error) {
	result := oci_container_instances.ContainerCapabilities{}

	if addCapabilities, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "add_capabilities")); ok {
		interfaces := addCapabilities.([]interface{})
		tmp := make([]oci_container_instances.ContainerCapabilityTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i], _ = oci_container_instances.GetMappingContainerCapabilityTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "add_capabilities")) {
			result.AddCapabilities = tmp
		}
	}

	if dropCapabilities, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "drop_capabilities")); ok {
		interfaces := dropCapabilities.([]interface{})
		tmp := make([]oci_container_instances.ContainerCapabilityTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i], _ = oci_container_instances.GetMappingContainerCapabilityTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "drop_capabilities")) {
			result.DropCapabilities = tmp
		}
	}

	return result, nil
}

func ContainerCapabilitiesToMap(obj *oci_container_instances.ContainerCapabilities) map[string]interface{} {
	result := map[string]interface{}{}

	result["add_capabilities"] = obj.AddCapabilities

	result["drop_capabilities"] = obj.DropCapabilities

	return result
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToContainerConfigFile(fieldKeyFormat string) (oci_container_instances.ContainerConfigFile, error) {
	result := oci_container_instances.ContainerConfigFile{}

	if data, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data")); ok {
		tmp, err := base64.StdEncoding.DecodeString(data.(string))
		if err != nil {
			return result, err
		}
		result.Data = []byte(tmp)
	}

	if fileName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_name")); ok {
		tmp := fileName.(string)
		result.FileName = &tmp
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	return result, nil
}

func ContainerConfigFileToMap(obj oci_container_instances.ContainerConfigFile, d *schema.ResourceData, volumeIndex int, configIndex int) map[string]interface{} {
	result := map[string]interface{}{}

	if configs, ok := d.GetOkExists("volumes"); ok {
		if tmpList := configs.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%s.%d.%%s", "volumes", volumeIndex, "configs", configIndex)
			if data, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data")); ok {
				tmp := data.(string)
				result["data"] = &tmp
			}
		}
	}

	if obj.FileName != nil {
		result["file_name"] = string(*obj.FileName)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	return result
}

func ContainerInstanceShapeConfigToMap(obj *oci_container_instances.ContainerInstanceShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.NetworkingBandwidthInGbps != nil {
		result["networking_bandwidth_in_gbps"] = float32(*obj.NetworkingBandwidthInGbps)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	if obj.ProcessorDescription != nil {
		result["processor_description"] = string(*obj.ProcessorDescription)
	}

	return result
}

func ContainerInstanceSummaryToMap(obj oci_container_instances.ContainerInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ContainerCount != nil {
		result["container_count"] = int(*obj.ContainerCount)
	}

	result["container_restart_policy"] = string(obj.ContainerRestartPolicy)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.GracefulShutdownTimeoutInSeconds != nil {
		result["graceful_shutdown_timeout_in_seconds"] = strconv.FormatInt(*obj.GracefulShutdownTimeoutInSeconds, 10)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.ShapeConfig != nil {
		result["shape_config"] = []interface{}{ContainerInstanceShapeConfigToMap(obj.ShapeConfig)}
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

	if obj.VolumeCount != nil {
		result["volume_count"] = int(*obj.VolumeCount)
	}

	return result
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateContainerDetails(fieldKeyFormat string) (oci_container_instances.CreateContainerDetails, error) {
	result := oci_container_instances.CreateContainerDetails{}

	if arguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "arguments")); ok {
		interfaces := arguments.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "arguments")) {
			result.Arguments = tmp
		}
	}

	if command, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command")); ok {
		interfaces := command.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "command")) {
			result.Command = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
		result.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if healthChecks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "health_checks")); ok {
		interfaces := healthChecks.([]interface{})
		tmp := make([]oci_container_instances.CreateContainerHealthCheckDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "health_checks"), stateDataIndex)
			converted, err := s.mapToCreateContainerHealthCheckDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "health_checks")) {
			result.HealthChecks = tmp
		}
	}

	if imageUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_url")); ok {
		tmp := imageUrl.(string)
		result.ImageUrl = &tmp
	}

	if isResourcePrincipalDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_resource_principal_disabled")); ok {
		tmp := isResourcePrincipalDisabled.(bool)
		result.IsResourcePrincipalDisabled = &tmp
	}

	if resourceConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_config")); ok {
		if tmpList := resourceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "resource_config"), 0)
			tmp, err := s.mapToCreateContainerResourceConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert resource_config, encountered error: %v", err)
			}
			result.ResourceConfig = &tmp
		}
	}

	if securityContext, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "security_context")); ok {
		if tmpList := securityContext.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "security_context"), 0)
			tmp, err := s.mapToCreateSecurityContextDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert security_context, encountered error: %v", err)
			}
			result.SecurityContext = tmp
		}
	}

	if volumeMounts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_mounts")); ok {
		interfaces := volumeMounts.([]interface{})
		tmp := make([]oci_container_instances.CreateVolumeMountDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "volume_mounts"), stateDataIndex)
			converted, err := s.mapToCreateVolumeMountDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "volume_mounts")) {
			result.VolumeMounts = tmp
		}
	}

	if workingDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "working_directory")); ok {
		tmp := workingDirectory.(string)
		result.WorkingDirectory = &tmp
	}

	return result, nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateContainerDnsConfigDetails(fieldKeyFormat string) (oci_container_instances.CreateContainerDnsConfigDetails, error) {
	result := oci_container_instances.CreateContainerDnsConfigDetails{}

	if nameservers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nameservers")); ok {
		interfaces := nameservers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nameservers")) {
			result.Nameservers = tmp
		}
	}

	if options, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "options")); ok {
		interfaces := options.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "options")) {
			result.Options = tmp
		}
	}

	if searches, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "searches")); ok {
		interfaces := searches.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "searches")) {
			result.Searches = tmp
		}
	}

	return result, nil
}

func ContainerDnsConfigToMap(obj *oci_container_instances.ContainerDnsConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["nameservers"] = obj.Nameservers

	result["options"] = obj.Options

	result["searches"] = obj.Searches

	return result
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateContainerHealthCheckDetails(fieldKeyFormat string) (oci_container_instances.CreateContainerHealthCheckDetails, error) {
	var baseObject oci_container_instances.CreateContainerHealthCheckDetails
	//discriminator
	healthCheckTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "health_check_type"))
	var healthCheckType string
	if ok {
		healthCheckType = healthCheckTypeRaw.(string)
	} else {
		healthCheckType = "" // default value
	}
	switch strings.ToLower(healthCheckType) {
	case strings.ToLower("HTTP"):
		details := oci_container_instances.CreateContainerHttpHealthCheckDetails{}
		if headers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "headers")); ok {
			interfaces := headers.([]interface{})
			tmp := make([]oci_container_instances.HealthCheckHttpHeader, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "headers"), stateDataIndex)
				converted, err := s.mapToHealthCheckHttpHeader(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "headers")) {
				details.Headers = tmp
			}
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if failureAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_action")); ok {
			details.FailureAction = oci_container_instances.ContainerHealthCheckFailureActionEnum(failureAction.(string))
		}
		if failureThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_threshold")); ok {
			tmp := failureThreshold.(int)
			details.FailureThreshold = &tmp
		}
		if initialDelayInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "initial_delay_in_seconds")); ok {
			tmp := initialDelayInSeconds.(int)
			details.InitialDelayInSeconds = &tmp
		}
		if intervalInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval_in_seconds")); ok {
			tmp := intervalInSeconds.(int)
			details.IntervalInSeconds = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if successThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "success_threshold")); ok {
			tmp := successThreshold.(int)
			details.SuccessThreshold = &tmp
		}
		if timeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_seconds")); ok {
			tmp := timeoutInSeconds.(int)
			details.TimeoutInSeconds = &tmp
		}
		baseObject = details
	case strings.ToLower("TCP"):
		details := oci_container_instances.CreateContainerTcpHealthCheckDetails{}
		if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if failureAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_action")); ok {
			details.FailureAction = oci_container_instances.ContainerHealthCheckFailureActionEnum(failureAction.(string))
		}
		if failureThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_threshold")); ok {
			tmp := failureThreshold.(int)
			details.FailureThreshold = &tmp
		}
		if initialDelayInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "initial_delay_in_seconds")); ok {
			tmp := initialDelayInSeconds.(int)
			details.InitialDelayInSeconds = &tmp
		}
		if intervalInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval_in_seconds")); ok {
			tmp := intervalInSeconds.(int)
			details.IntervalInSeconds = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if successThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "success_threshold")); ok {
			tmp := successThreshold.(int)
			details.SuccessThreshold = &tmp
		}
		if timeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_seconds")); ok {
			tmp := timeoutInSeconds.(int)
			details.TimeoutInSeconds = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown health_check_type '%v' was specified", healthCheckType)
	}
	return baseObject, nil
}

func ContainerHealthCheckToMap(obj oci_container_instances.ContainerHealthCheck) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_container_instances.ContainerHttpHealthCheck:
		result["health_check_type"] = "HTTP"

		headers := []interface{}{}
		for _, item := range v.Headers {
			headers = append(headers, HealthCheckHttpHeaderToMap(item))
		}
		result["headers"] = headers

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		result["failure_action"] = string(v.FailureAction)

		if v.FailureThreshold != nil {
			result["failure_threshold"] = int(*v.FailureThreshold)
		}

		if v.InitialDelayInSeconds != nil {
			result["initial_delay_in_seconds"] = int(*v.InitialDelayInSeconds)
		}

		if v.IntervalInSeconds != nil {
			result["interval_in_seconds"] = int(*v.IntervalInSeconds)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		result["status"] = string(v.Status)

		if v.StatusDetails != nil {
			result["status_details"] = string(*v.StatusDetails)
		}

		if v.SuccessThreshold != nil {
			result["success_threshold"] = int(*v.SuccessThreshold)
		}

		if v.TimeoutInSeconds != nil {
			result["timeout_in_seconds"] = int(*v.TimeoutInSeconds)
		}
	case oci_container_instances.ContainerTcpHealthCheck:
		result["health_check_type"] = "TCP"

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		result["failure_action"] = string(v.FailureAction)

		if v.FailureThreshold != nil {
			result["failure_threshold"] = int(*v.FailureThreshold)
		}

		if v.InitialDelayInSeconds != nil {
			result["initial_delay_in_seconds"] = int(*v.InitialDelayInSeconds)
		}

		if v.IntervalInSeconds != nil {
			result["interval_in_seconds"] = int(*v.IntervalInSeconds)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		result["status"] = string(v.Status)

		if v.StatusDetails != nil {
			result["status_details"] = string(*v.StatusDetails)
		}

		if v.SuccessThreshold != nil {
			result["success_threshold"] = int(*v.SuccessThreshold)
		}

		if v.TimeoutInSeconds != nil {
			result["timeout_in_seconds"] = int(*v.TimeoutInSeconds)
		}
	default:
		log.Printf("[WARN] Received 'health_check_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateContainerInstanceShapeConfigDetails(fieldKeyFormat string) (oci_container_instances.CreateContainerInstanceShapeConfigDetails, error) {
	result := oci_container_instances.CreateContainerInstanceShapeConfigDetails{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateContainerResourceConfigDetails(fieldKeyFormat string) (oci_container_instances.CreateContainerResourceConfigDetails, error) {
	result := oci_container_instances.CreateContainerResourceConfigDetails{}

	if memoryLimitInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_limit_in_gbs")); ok {
		tmp := float32(memoryLimitInGBs.(float64))
		result.MemoryLimitInGBs = &tmp
	}

	if vcpusLimit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcpus_limit")); ok {
		tmp := float32(vcpusLimit.(float64))
		result.VcpusLimit = &tmp
	}

	return result, nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateContainerVnicDetails(fieldKeyFormat string) (oci_container_instances.CreateContainerVnicDetails, error) {
	result := oci_container_instances.CreateContainerVnicDetails{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if isPublicIpAssigned, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_public_ip_assigned")); ok {
		tmp := isPublicIpAssigned.(bool)
		result.IsPublicIpAssigned = &tmp
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateContainerVolumeDetails(fieldKeyFormat string) (oci_container_instances.CreateContainerVolumeDetails, error) {
	var baseObject oci_container_instances.CreateContainerVolumeDetails
	//discriminator
	volumeTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_type"))
	var volumeType string
	if ok {
		volumeType = volumeTypeRaw.(string)
	} else {
		volumeType = "" // default value
	}
	switch strings.ToLower(volumeType) {
	case strings.ToLower("CONFIGFILE"):
		details := oci_container_instances.CreateContainerConfigFileVolumeDetails{}
		if configs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configs")); ok {
			interfaces := configs.([]interface{})
			tmp := make([]oci_container_instances.ContainerConfigFile, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configs"), stateDataIndex)
				converted, err := s.mapToContainerConfigFile(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "configs")) {
				details.Configs = tmp
			}
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("EMPTYDIR"):
		details := oci_container_instances.CreateContainerEmptyDirVolumeDetails{}
		if backingStore, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backing_store")); ok {
			details.BackingStore = oci_container_instances.ContainerEmptyDirVolumeBackingStoreEnum(backingStore.(string))
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown volume_type '%v' was specified", volumeType)
	}
	return baseObject, nil
}

func ContainerVolumeToMap(obj oci_container_instances.ContainerVolume, d *schema.ResourceData, index int) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_container_instances.ContainerConfigFileVolume:
		result["volume_type"] = "CONFIGFILE"

		configs := []interface{}{}
		for configIndex, item := range v.Configs {
			configs = append(configs, ContainerConfigFileToMap(item, d, index, configIndex))
		}
		result["configs"] = configs

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	case oci_container_instances.ContainerEmptyDirVolume:
		result["volume_type"] = "EMPTYDIR"

		result["backing_store"] = string(v.BackingStore)

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	default:
		log.Printf("[WARN] Received 'volume_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateImagePullSecretDetails(fieldKeyFormat string) (oci_container_instances.CreateImagePullSecretDetails, error) {
	var baseObject oci_container_instances.CreateImagePullSecretDetails
	//discriminator
	secretTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_type"))
	var secretType string
	if ok {
		secretType = secretTypeRaw.(string)
	} else {
		secretType = "" // default value
	}
	switch strings.ToLower(secretType) {
	case strings.ToLower("BASIC"):
		details := oci_container_instances.CreateBasicImagePullSecretDetails{}
		if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if registryEndpoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_endpoint")); ok {
			tmp := registryEndpoint.(string)
			details.RegistryEndpoint = &tmp
		}
		baseObject = details
	case strings.ToLower("VAULT"):
		details := oci_container_instances.CreateVaultImagePullSecretDetails{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if registryEndpoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_endpoint")); ok {
			tmp := registryEndpoint.(string)
			details.RegistryEndpoint = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown secret_type '%v' was specified", secretType)
	}
	return baseObject, nil
}

func ImagePullSecretToMap(obj oci_container_instances.ImagePullSecret, d *schema.ResourceData, index int) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_container_instances.BasicImagePullSecret:
		result["secret_type"] = "BASIC"

		if imagePullSecrets, ok := d.GetOkExists("image_pull_secrets"); ok {
			if tmpList := imagePullSecrets.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_pull_secrets", index)
				if username, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
					tmp := username.(string)
					result["username"] = &tmp
				}

				if password, ok := d.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
					tmp := password.(string)
					result["password"] = &tmp
				}
			}
		}

		if v.RegistryEndpoint != nil {
			result["registry_endpoint"] = string(*v.RegistryEndpoint)
		}
	case oci_container_instances.VaultImagePullSecret:
		result["secret_type"] = "VAULT"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}

		if v.RegistryEndpoint != nil {
			result["registry_endpoint"] = string(*v.RegistryEndpoint)
		}
	default:
		log.Printf("[WARN] Received 'secret_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateSecurityContextDetails(fieldKeyFormat string) (oci_container_instances.CreateSecurityContextDetails, error) {
	var baseObject oci_container_instances.CreateSecurityContextDetails
	//discriminator
	securityContextTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "security_context_type"))
	var securityContextType string
	if ok {
		securityContextType = securityContextTypeRaw.(string)
	} else {
		securityContextType = "LINUX" // default value
	}
	switch strings.ToLower(securityContextType) {
	case strings.ToLower("LINUX"):
		details := oci_container_instances.CreateLinuxSecurityContextDetails{}
		if capabilities, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capabilities")); ok {
			if tmpList := capabilities.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "capabilities"), 0)
				tmp, err := s.mapToContainerCapabilities(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert capabilities, encountered error: %v", err)
				}
				details.Capabilities = &tmp
			}
		}
		if isNonRootUserCheckEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_non_root_user_check_enabled")); ok {
			tmp := isNonRootUserCheckEnabled.(bool)
			details.IsNonRootUserCheckEnabled = &tmp
		}
		if isRootFileSystemReadonly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_root_file_system_readonly")); ok {
			tmp := isRootFileSystemReadonly.(bool)
			details.IsRootFileSystemReadonly = &tmp
		}
		if runAsGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_as_group")); ok {
			tmp := runAsGroup.(int)
			details.RunAsGroup = &tmp
		}
		if runAsUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_as_user")); ok {
			tmp := runAsUser.(int)
			details.RunAsUser = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown security_context_type '%v' was specified", securityContextType)
	}
	return baseObject, nil
}

func SecurityContextToMap(obj *oci_container_instances.SecurityContext) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_container_instances.LinuxSecurityContext:
		result["security_context_type"] = "LINUX"

		if v.Capabilities != nil {
			result["capabilities"] = []interface{}{ContainerCapabilitiesToMap(v.Capabilities)}
		}

		if v.IsNonRootUserCheckEnabled != nil {
			result["is_non_root_user_check_enabled"] = bool(*v.IsNonRootUserCheckEnabled)
		}

		if v.IsRootFileSystemReadonly != nil {
			result["is_root_file_system_readonly"] = bool(*v.IsRootFileSystemReadonly)
		}

		if v.RunAsGroup != nil {
			result["run_as_group"] = int(*v.RunAsGroup)
		}

		if v.RunAsUser != nil {
			result["run_as_user"] = int(*v.RunAsUser)
		}
	default:
		log.Printf("[WARN] Received 'security_context_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateVolumeMountDetails(fieldKeyFormat string) (oci_container_instances.CreateVolumeMountDetails, error) {
	result := oci_container_instances.CreateVolumeMountDetails{}

	if isReadOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_only")); ok {
		tmp := isReadOnly.(bool)
		result.IsReadOnly = &tmp
	}

	if mountPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_path")); ok {
		tmp := mountPath.(string)
		result.MountPath = &tmp
	}

	if partition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "partition")); ok {
		tmp := partition.(int)
		result.Partition = &tmp
	}

	if subPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sub_path")); ok {
		tmp := subPath.(string)
		result.SubPath = &tmp
	}

	if volumeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_name")); ok {
		tmp := volumeName.(string)
		result.VolumeName = &tmp
	}

	return result, nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToHealthCheckHttpHeader(fieldKeyFormat string) (oci_container_instances.HealthCheckHttpHeader, error) {
	result := oci_container_instances.HealthCheckHttpHeader{}

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

func HealthCheckHttpHeaderToMap(obj oci_container_instances.HealthCheckHttpHeader) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *ContainerInstancesContainerInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_container_instances.ChangeContainerInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ContainerInstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance")

	response, err := s.Client.ChangeContainerInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getContainerInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerinstance"), oci_container_instances.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToCreateVnicDetails(fieldKeyFormat string) (oci_core.CreateVnicDetails, error) {
	result := oci_core.CreateVnicDetails{}

	if assignPrivateDnsRecord, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_private_dns_record")); ok {
		tmp := assignPrivateDnsRecord.(bool)
		result.AssignPrivateDnsRecord = &tmp
	}

	if assignPublicIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")); ok {
		tmp := assignPublicIp.(string)
		boolVal, err := strconv.ParseBool(tmp)
		if err != nil {
			return result, err
		}
		result.AssignPublicIp = &boolVal
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	if vlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan_id")); ok {
		tmp := vlanId.(string)
		result.VlanId = &tmp
	}

	return result, nil
}

func (s *ContainerInstancesContainerInstanceResourceCrud) mapToUpdateContainer(fieldKeyFormat string) (oci_container_instances.UpdateContainerRequest, error) {
	result := oci_container_instances.UpdateContainerRequest{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_id")); ok {
		tmp := id.(string)
		result.ContainerId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	return result, nil
}

func VnicDetailsToMap(obj oci_core.Vnic, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	// "assign_public_ip" isn't part of the VNIC's state & is only useful at creation time (and
	// subsequent force-new creations). So persist the user-defined value in the config & Update it
	// when the user changes that value.

	// We may be importing this value; so let's set it to whether the public IP is set.
	result["is_public_ip_assigned"] = obj.PublicIp != nil && *obj.PublicIp != ""

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	if obj.SkipSourceDestCheck != nil {
		result["skip_source_dest_check"] = bool(*obj.SkipSourceDestCheck)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.VlanId != nil {
		result["vlan_id"] = string(*obj.VlanId)
	}

	if obj.Id != nil {
		result["vnic_id"] = string(*obj.Id)
	}
	return result
}

func ContainerToMap(obj oci_container_instances.Container) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ContainerInstanceId != nil {
		result["container_instance_id"] = string(*obj.ContainerInstanceId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["container_id"] = string(*obj.Id)
	}

	if obj.ImageUrl != nil {
		result["image_url"] = string(*obj.ImageUrl)
	}

	if obj.IsResourcePrincipalDisabled != nil {
		result["is_resource_principal_disabled"] = bool(*obj.IsResourcePrincipalDisabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ResourceConfig != nil {
		result["resource_config"] = []interface{}{ContainerResourceConfigToMap(obj.ResourceConfig)}
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

	if obj.Arguments != nil {
		result["arguments"] = obj.Arguments
	}

	if obj.Command != nil {
		result["command"] = obj.Command
	}

	if obj.EnvironmentVariables != nil {
		result["environment_variables"] = obj.EnvironmentVariables
	}

	if obj.FreeformTags != nil {
		result["freeform_tags"] = obj.FreeformTags
	}

	if obj.VolumeMounts != nil {
		var volumeMounts []interface{}

		for _, item := range obj.VolumeMounts {
			volumeMounts = append(volumeMounts, ContainerVolumeMountsToMap(item))
		}
		result["volume_mounts"] = volumeMounts
	}

	if obj.ExitCode != nil {
		result["exit_code"] = obj.ExitCode
	}

	if obj.HealthChecks != nil {
		var healthChecks []interface{}

		for _, item := range obj.HealthChecks {
			healthChecks = append(healthChecks, ContainerHealthCheckToMap(item))
		}

		result["health_checks"] = healthChecks
	}

	if obj.WorkingDirectory != nil {
		result["working_directory"] = obj.WorkingDirectory
	}

	if obj.SecurityContext != nil {
		securityContextArray := []interface{}{}
		if securityContextMap := SecurityContextToMap(&obj.SecurityContext); securityContextMap != nil {
			securityContextArray = append(securityContextArray, securityContextMap)
		}
		result["security_context"] = securityContextArray
	}

	return result
}

func ContainerResourceConfigToMap(obj *oci_container_instances.ContainerResourceConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryLimitInGBs != nil {
		result["memory_limit_in_gbs"] = float32(*obj.MemoryLimitInGBs)
	}

	if obj.VcpusLimit != nil {
		result["vcpus_limit"] = float32(*obj.VcpusLimit)
	}

	return result
}

func ContainerVolumeMountsToMap(obj oci_container_instances.VolumeMount) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MountPath != nil {
		result["mount_path"] = obj.MountPath
	}

	if obj.VolumeName != nil {
		result["volume_name"] = obj.VolumeName
	}

	if obj.IsReadOnly != nil {
		result["is_read_only"] = obj.IsReadOnly
	}

	if obj.Partition != nil {
		result["partition"] = obj.Partition
	}

	if obj.SubPath != nil {
		result["sub_path"] = obj.SubPath
	}

	return result
}
