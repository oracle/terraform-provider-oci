// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"
)

func CoreInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("45m"),
			Update: tfresource.GetTimeoutDuration("45m"),
			Delete: tfresource.GetTimeoutDuration("75m"),
		},
		Create: createCoreInstance,
		Read:   readCoreInstance,
		Update: updateCoreInstance,
		Delete: deleteCoreInstance,
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
			"shape": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"async": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"agent_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"are_all_plugins_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_management_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_monitoring_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"plugins_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"desired_state": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"availability_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_live_migration_preferred": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"recovery_action": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"capacity_reservation_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"create_vnic_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"assign_private_dns_record": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"assign_public_ip": {
							// Change type from boolean to string because TF doesn't handle default
							// values for boolean nested objects correctly.
							Type:     schema.TypeString,
							Optional: true,
							Default:  "true",
							ValidateFunc: func(v interface{}, k string) ([]string, []error) {
								// Verify that we can parse the string value as a bool value.
								var es []error
								if _, err := strconv.ParseBool(v.(string)); err != nil {
									es = append(es, fmt.Errorf("%s: cannot parse 'assign_public_ip' as bool: %v", k, err))
								}
								return nil, es
							},
							StateFunc: func(v interface{}) string {
								// ValidateFunc runs before StateFunc. Must be valid by now.
								b, _ := tfresource.NormalizeBoolString(v.(string))
								return b
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
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"hostname_label": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Set:      utils.LiteralTypeHashCodeForSets,
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
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"dedicated_vm_host_id": {
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
			"extended_metadata": {
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: utils.JsonStringDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"fault_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"hostname_label": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				Deprecated:       tfresource.FieldDeprecatedForAnother("hostname_label", "hostname_label under create_vnic_details"),
			},
			"image": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: tfresource.FieldDeprecatedAndOverridenByAnother("image", "source_details"),
			},
			"instance_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"are_legacy_imds_endpoints_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"ipxe_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_pv_encryption_in_transit_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"launch_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"boot_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"firmware": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_consistent_volume_naming_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remote_data_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"platform_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AMD_MILAN_BM",
								"AMD_ROME_BM",
								"AMD_VM",
								"INTEL_SKYLAKE_BM",
								"INTEL_VM",
							}, true),
						},

						// Optional
						"is_measured_boot_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_secure_boot_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_trusted_platform_module_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"numa_nodes_per_socket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"preemptible_instance_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"preemption_action": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"TERMINATE",
										}, true),
									},

									// Optional
									"preserve_boot_volume": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"shape_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"baseline_ocpu_utilization": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"memory_in_gbs": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"ocpus": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},

						// Computed
						"gpu_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gpus": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_disk_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"local_disks": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_disks_total_size_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"max_vnic_attachments": {
							Type:     schema.TypeInt,
							Computed: true,
						},
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
			"preserve_boot_volume": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"source_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"bootVolume",
								"image",
							}, true),
						},

						// Optional
						"boot_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     utils.ValidateInt64TypeString,
							DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"subnet_id": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: tfresource.FieldDeprecatedForAnother("subnet_id", "subnet_id under create_vnic_details"),
			},

			// Computed
			// Add this computed boot_volume_id field even though it's not part of the API specs. This will make it easier to
			// discover the attached boot volume's ID; to preserve it for reattachment.
			"boot_volume_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"launch_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_core.InstanceLifecycleStateStopped),
					string(oci_core.InstanceLifecycleStateRunning),
				}, true),
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
			"time_maintenance_reboot_due": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Legacy custom computed convenience values
			"public_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		// CustomizeDiff for Instance resource
		// Updates of 'ssh_authorized_keys' and 'user_data' in Instance 'metadata' should result in Force New
		CustomizeDiff: customdiff.All(
			customdiff.ForceNewIfChange("metadata", func(old, new, meta interface{}) bool {
				oldMetadataMap := utils.ObjectMapToStringMap(old.(map[string]interface{}))
				newMetadataMap := utils.ObjectMapToStringMap(new.(map[string]interface{}))
				return (oldMetadataMap["ssh_authorized_keys"] != newMetadataMap["ssh_authorized_keys"]) || (oldMetadataMap["user_data"] != newMetadataMap["user_data"])
			}),
		),
	}
}

func createCoreInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	sync.BlockStorageClient = m.(*client.OracleClients).BlockstorageClient()

	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_core.InstanceLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_core.InstanceLifecycleStateStopped {
			powerOff = true
		}
	}

	if e := tfresource.CreateResourceUsingHybridPolling(sync); e != nil {
		return e
	}

	return powerOffIfNeeded(sync.D, sync, powerOff)
}

func powerOffIfNeeded(d *schema.ResourceData, sync *CoreInstanceResourceCrud, powerOff bool) error {

	if powerOff {
		if err := sync.InstanceAction(oci_core.InstanceActionActionStop, oci_core.InstanceLifecycleStateStopped); err != nil {
			return err
		}
		return tfresource.ReadResource(sync)
	}
	return nil
}

func readCoreInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	sync.BlockStorageClient = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

func updateCoreInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	sync.BlockStorageClient = m.(*client.OracleClients).BlockstorageClient()

	// switch to power on
	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_core.InstanceLifecycleStateRunning == oci_core.InstanceLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_core.InstanceLifecycleStateStopped == oci_core.InstanceLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.InstanceAction(oci_core.InstanceActionActionStart, oci_core.InstanceLifecycleStateRunning); err != nil {
			return err
		}
		sync.D.Set("state", oci_core.InstanceLifecycleStateRunning)
	}
	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}
	// switch to power off
	if powerOff {
		if err := sync.InstanceAction(oci_core.InstanceActionActionStop, oci_core.InstanceLifecycleStateStopped); err != nil {
			return err
		}
		sync.D.Set("state", oci_core.InstanceLifecycleStateStopped)
	}
	return nil
}

func deleteCoreInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	sync.BlockStorageClient = m.(*client.OracleClients).BlockstorageClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	VirtualNetworkClient   *oci_core.VirtualNetworkClient
	BlockStorageClient     *oci_core.BlockstorageClient
	Res                    *oci_core.Instance
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateProvisioning),
		string(oci_core.InstanceLifecycleStateStarting),
	}
}

func (s *CoreInstanceResourceCrud) CreatedTarget() []string {
	if asyn, ok := s.D.GetOk("async"); ok {
		tmp := asyn.(bool)
		if tmp {
			return []string{
				string(oci_core.InstanceLifecycleStateRunning),
				string(oci_core.InstanceLifecycleStateProvisioning),
				string(oci_core.InstanceLifecycleStateStarting),
			}
		}
	}
	return []string{
		string(oci_core.InstanceLifecycleStateRunning),
	}
}

func (s *CoreInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InstanceLifecycleStateTerminating),
	}
}

func (s *CoreInstanceResourceCrud) DeletedTarget() []string {

	if asyn, ok := s.D.GetOk("async"); ok {
		tmp := asyn.(bool)
		if tmp {
			return []string{
				string(oci_core.InstanceLifecycleStateTerminated),
				string(oci_core.InstanceLifecycleStateTerminating),
			}
		}
	}
	return []string{
		string(oci_core.InstanceLifecycleStateTerminated),
	}
}

func (s *CoreInstanceResourceCrud) Create() error {
	request := oci_core.LaunchInstanceRequest{}

	if agentConfig, ok := s.D.GetOkExists("agent_config"); ok {
		if tmpList := agentConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "agent_config", 0)
			tmp, err := s.mapToLaunchInstanceAgentConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AgentConfig = &tmp
		}
	}

	if availabilityConfig, ok := s.D.GetOkExists("availability_config"); ok {
		if tmpList := availabilityConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "availability_config", 0)
			tmp, err := s.mapToLaunchInstanceAvailabilityConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AvailabilityConfig = &tmp
		}
	}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createVnicDetails, ok := s.D.GetOkExists("create_vnic_details"); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "create_vnic_details", 0)
			tmp, err := s.mapToCreateVnicDetailsInstance(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CreateVnicDetails = &tmp
		}
	}

	if dedicatedVmHostId, ok := s.D.GetOkExists("dedicated_vm_host_id"); ok {
		tmp := dedicatedVmHostId.(string)
		request.DedicatedVmHostId = &tmp
		//@codegen: Adding wait to ensure that the DVH is available
		time.Sleep(1 * time.Minute)
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

	if rawExtendedMetadata, ok := s.D.GetOkExists("extended_metadata"); ok {
		extendedMetadata, err := mapToExtendedMetadata(rawExtendedMetadata.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.ExtendedMetadata = extendedMetadata
	}

	if faultDomain, ok := s.D.GetOkExists("fault_domain"); ok {
		tmp := faultDomain.(string)
		request.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	if image, ok := s.D.GetOkExists("image"); ok {
		tmp := image.(string)
		request.ImageId = &tmp
	}

	if instanceOptions, ok := s.D.GetOkExists("instance_options"); ok {
		if tmpList := instanceOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "instance_options", 0)
			tmp, err := s.mapToInstanceOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InstanceOptions = &tmp
		}
	}

	if ipxeScript, ok := s.D.GetOkExists("ipxe_script"); ok {
		tmp := ipxeScript.(string)
		request.IpxeScript = &tmp
	}

	if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists("is_pv_encryption_in_transit_enabled"); ok {
		tmp := isPvEncryptionInTransitEnabled.(bool)
		request.IsPvEncryptionInTransitEnabled = &tmp
	}

	if launchOptions, ok := s.D.GetOkExists("launch_options"); ok {
		if tmpList := launchOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "launch_options", 0)
			tmp, err := s.mapToLaunchOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LaunchOptions = &tmp
		}
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = utils.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	if platformConfig, ok := s.D.GetOkExists("platform_config"); ok {
		if tmpList := platformConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "platform_config", 0)
			tmp, err := s.mapToLaunchInstancePlatformConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PlatformConfig = tmp
		}
	}

	if preemptibleInstanceConfig, ok := s.D.GetOkExists("preemptible_instance_config"); ok {
		if tmpList := preemptibleInstanceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "preemptible_instance_config", 0)
			tmp, err := s.mapToPreemptibleInstanceConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PreemptibleInstanceConfig = &tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if shapeConfig, ok := s.D.GetOkExists("shape_config"); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shape_config", 0)
			tmp, err := s.mapToLaunchInstanceShapeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ShapeConfig = &tmp
		}
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			tmp, err := s.mapToInstanceSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceDetails = &tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.LaunchInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workRequestId := response.OpcWorkRequestId

	s.Res = &response.Instance

	workRequestErr := tfresource.ResourceRefreshForHybridPolling(s.WorkRequestClient, workRequestId, "instance", oci_work_requests.WorkRequestResourceActionTypeCreated, s.DisableNotFoundRetries, s.D, s)
	if workRequestErr != nil {
		return workRequestErr
	}

	return nil
}

func (s *CoreInstanceResourceCrud) Get() error {
	request := oci_core.GetInstanceRequest{}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance
	return nil
}

func (s *CoreInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			err := s.mapToUpdateInstanceBootVolumeSizeInGbs(fieldKeyFormat)
			if err != nil {
				return err
			}
		}
	}

	// Update shape, shape config, fault domain and launch options
	err := s.updateOptionsViaWorkRequest()

	if err != nil {
		return err
	}

	request := oci_core.UpdateInstanceRequest{}

	if agentConfig, ok := s.D.GetOkExists("agent_config"); ok {
		if tmpList := agentConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "agent_config", 0)
			tmp, err := s.mapToUpdateInstanceAgentConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AgentConfig = &tmp
		}
	}

	if availabilityConfig, ok := s.D.GetOkExists("availability_config"); ok {
		if tmpList := availabilityConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "availability_config", 0)
			tmp, err := s.mapToUpdateInstanceAvailabilityConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AvailabilityConfig = &tmp
		}
	}

	if capacityReservationId, ok := s.D.GetOkExists("capacity_reservation_id"); ok {
		tmp := capacityReservationId.(string)
		request.CapacityReservationId = &tmp
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

	if rawExtendedMetadata, ok := s.D.GetOkExists("extended_metadata"); ok {
		extendedMetadata, err := mapToExtendedMetadata(rawExtendedMetadata.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.ExtendedMetadata = extendedMetadata
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	if instanceOptions, ok := s.D.GetOkExists("instance_options"); ok {
		if tmpList := instanceOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "instance_options", 0)
			tmp, err := s.mapToInstanceOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InstanceOptions = &tmp
		}
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = utils.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance

	// Check for changes in the create_vnic_details sub resource and separately Update the vnic

	_, ok := s.D.GetOkExists("create_vnic_details")
	if !s.D.HasChange("create_vnic_details") || !ok {
		log.Printf("[DEBUG] No changes to primary VNIC. Instance ID: \"%v\"", s.Res.Id)
		return nil
	}

	vnic, err := s.getPrimaryVnic()
	if err != nil {
		log.Printf("[ERROR] Primary VNIC could not be found during instance Update: %q (Instance ID: \"%v\", State: %q)", err, s.Res.Id, s.Res.LifecycleState)
		return err
	}

	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "create_vnic_details", 0)
	err = s.updateVnicAssignPublicIp(vnic, fieldKeyFormat)
	if err != nil {
		return err
	}

	updateVnicDetails, err := s.mapToUpdateVnicDetailsInstance(fieldKeyFormat)
	if err != nil {
		return err
	}

	vnicOpts := oci_core.UpdateVnicRequest{
		VnicId:            vnic.Id,
		UpdateVnicDetails: updateVnicDetails,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
		},
	}

	_, err = s.VirtualNetworkClient.UpdateVnic(context.Background(), vnicOpts)

	if err != nil {
		log.Printf("[ERROR] Primary VNIC could not be updated during instance Update: %q (Instance ID: \"%v\", State: %q)", err, s.Res.Id, s.Res.LifecycleState)
		return err
	}

	return nil
}

func (s *CoreInstanceResourceCrud) InstanceAction(action oci_core.InstanceActionActionEnum, state oci_core.InstanceLifecycleStateEnum) error {
	request := oci_core.InstanceActionRequest{}
	request.Action = action

	tmp := s.D.Id()
	request.InstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.InstanceAction(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == state }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))

}

func (s *CoreInstanceResourceCrud) Delete() error {
	request := oci_core.TerminateInstanceRequest{}

	tmp := s.D.Id()
	request.InstanceId = &tmp

	if preserveBootVolume, ok := s.D.GetOkExists("preserve_boot_volume"); ok {
		tmp := preserveBootVolume.(bool)
		request.PreserveBootVolume = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.TerminateInstance(context.Background(), request)
	return err
}

func (s *CoreInstanceResourceCrud) SetData() error {
	if s.Res.AgentConfig != nil {
		s.D.Set("agent_config", []interface{}{InstanceAgentConfigToMap(s.Res.AgentConfig)})
	} else {
		s.D.Set("agent_config", nil)
	}

	if s.Res.AvailabilityConfig != nil {
		s.D.Set("availability_config", []interface{}{InstanceAvailabilityConfigToMap(s.Res.AvailabilityConfig)})
	} else {
		s.D.Set("availability_config", nil)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", *s.Res.CapacityReservationId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DedicatedVmHostId != nil {
		s.D.Set("dedicated_vm_host_id", *s.Res.DedicatedVmHostId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExtendedMetadata != nil {
		s.D.Set("extended_metadata", tfresource.GenericMapToJsonMap(s.Res.ExtendedMetadata))
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		s.D.Set("image", *s.Res.ImageId)
	}

	if s.Res.InstanceOptions != nil {
		s.D.Set("instance_options", []interface{}{InstanceOptionsToMap(s.Res.InstanceOptions)})
	} else {
		s.D.Set("instance_options", nil)
	}

	if s.Res.IpxeScript != nil {
		s.D.Set("ipxe_script", *s.Res.IpxeScript)
	}

	s.D.Set("launch_mode", s.Res.LaunchMode)

	if s.Res.LaunchOptions != nil {
		s.D.Set("launch_options", []interface{}{LaunchOptionsToMap(s.Res.LaunchOptions)})
	} else {
		s.D.Set("launch_options", nil)
	}

	if s.Res.Metadata != nil {
		err := s.D.Set("metadata", s.Res.Metadata)
		if err != nil {
			log.Printf("error setting metadata %q", err)
		}
	}

	if s.Res.PlatformConfig != nil {
		platformConfigArray := []interface{}{}
		if platformConfigMap := PlatformConfigToMap(&s.Res.PlatformConfig); platformConfigMap != nil {
			platformConfigArray = append(platformConfigArray, platformConfigMap)
		}
		s.D.Set("platform_config", platformConfigArray)
	} else {
		s.D.Set("platform_config", nil)
	}

	if s.Res.PreemptibleInstanceConfig != nil {
		s.D.Set("preemptible_instance_config", []interface{}{PreemptibleInstanceConfigDetailsToMap(s.Res.PreemptibleInstanceConfig)})
	} else {
		s.D.Set("preemptible_instance_config", nil)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	if s.Res.ShapeConfig != nil {
		s.D.Set("shape_config", []interface{}{InstanceShapeConfigToMap(s.Res.ShapeConfig)})
	} else {
		s.D.Set("shape_config", []interface{}{})
	}

	bootVolume, bootVolumeErr := s.getBootVolume()
	if bootVolumeErr != nil {
		log.Printf("[WARN] Could not get the boot volume: %q", bootVolumeErr)
	}

	if s.Res.SourceDetails != nil {
		var sourceDetailsFromConfig map[string]interface{}
		if details, ok := s.D.GetOkExists("source_details"); ok {
			if tmpList := details.([]interface{}); len(tmpList) > 0 {
				sourceDetailsFromConfig = tmpList[0].(map[string]interface{})
			}
		}
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := InstanceSourceDetailsToMap(&s.Res.SourceDetails, bootVolume, sourceDetailsFromConfig); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		err := s.D.Set("source_details", sourceDetailsArray)
		if err != nil {
			return err
		}
	} else {
		s.D.Set("source_details", nil)
	}

	if bootVolume != nil && bootVolume.Id != nil {
		s.D.Set("boot_volume_id", *bootVolume.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeMaintenanceRebootDue != nil {
		s.D.Set("time_maintenance_reboot_due", s.Res.TimeMaintenanceRebootDue.String())
	} else {
		// If the maintenance time is cleared after reboot, the service will return a nil.
		// We should explicitly zero it out to avoid returning the previously cached reboot time.
		s.D.Set("time_maintenance_reboot_due", "")
	}

	if s.Res.LifecycleState != oci_core.InstanceLifecycleStateTerminated &&
		s.Res.LifecycleState != oci_core.InstanceLifecycleStateProvisioning &&
		s.Res.LifecycleState != oci_core.InstanceLifecycleStateTerminating {
		vnic, vnicError := s.getPrimaryVnic()
		if vnicError != nil || vnic == nil {
			log.Printf("[WARN] Primary VNIC could not be found during instance refresh: %q", vnicError)
		} else {
			s.D.Set("hostname_label", vnic.HostnameLabel)
			s.D.Set("nsg_ids", vnic.NsgIds)
			s.D.Set("public_ip", vnic.PublicIp)
			s.D.Set("private_ip", vnic.PrivateIp)
			s.D.Set("subnet_id", vnic.SubnetId)

			var createVnicDetails map[string]interface{}
			if details, ok := s.D.GetOkExists("create_vnic_details"); ok {
				if tmpList := details.([]interface{}); len(tmpList) > 0 {
					createVnicDetails = tmpList[0].(map[string]interface{})
				}
			}

			err := s.D.Set("create_vnic_details", []interface{}{CreateVnicDetailsToMap(vnic, createVnicDetails, false)})
			if err != nil {
				log.Printf("[WARN] create_vnic_details could not be set: %q", err)
			}
		}
	}

	return nil
}

func (s *CoreInstanceResourceCrud) mapToCreateVnicDetailsInstance(fieldKeyFormat string) (oci_core.CreateVnicDetails, error) {
	result := oci_core.CreateVnicDetails{}

	if assignPrivateDnsRecord, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_private_dns_record")); ok {
		tmp := assignPrivateDnsRecord.(bool)
		result.AssignPrivateDnsRecord = &tmp
		if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "assign_private_dns_record")) {
			result.AssignPrivateDnsRecord = &tmp
		}
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
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

func CreateVnicDetailsToMap(obj *oci_core.Vnic, createVnicDetails map[string]interface{}, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if createVnicDetails != nil {
		result["assign_private_dns_record"] = createVnicDetails["assign_private_dns_record"]
	}
	// "assign_public_ip" isn't part of the VNIC's state & is only useful at creation time (and
	// subsequent force-new creations). So persist the user-defined value in the config & Update it
	// when the user changes that value.
	if createVnicDetails != nil {
		assignPublicIP, _ := tfresource.NormalizeBoolString(createVnicDetails["assign_public_ip"].(string)) // Must be valid.
		result["assign_public_ip"] = assignPublicIP
	} else {
		// We may be importing this value; so let's set it to whether the public IP is set.
		result["assign_public_ip"] = strconv.FormatBool(obj.PublicIp != nil && *obj.PublicIp != "")
	}

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
		result["nsg_ids"] = schema.NewSet(utils.LiteralTypeHashCodeForSets, nsgIds)
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

	return result
}

func (s *CoreInstanceResourceCrud) mapToInstanceAgentPluginConfigDetails(fieldKeyFormat string) (oci_core.InstanceAgentPluginConfigDetails, error) {
	result := oci_core.InstanceAgentPluginConfigDetails{}

	if desiredState, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "desired_state")); ok {
		result.DesiredState = oci_core.InstanceAgentPluginConfigDetailsDesiredStateEnum(desiredState.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func (s *CoreInstanceResourceCrud) mapToInstanceOptions(fieldKeyFormat string) (oci_core.InstanceOptions, error) {
	result := oci_core.InstanceOptions{}

	if areLegacyImdsEndpointsDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_legacy_imds_endpoints_disabled")); ok {
		tmp := areLegacyImdsEndpointsDisabled.(bool)
		result.AreLegacyImdsEndpointsDisabled = &tmp
	}

	return result, nil
}

func (s *CoreInstanceResourceCrud) mapToUpdateVnicDetailsInstance(fieldKeyFormat string) (oci_core.UpdateVnicDetails, error) {
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
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

func InstanceOptionsToMap(obj *oci_core.InstanceOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreLegacyImdsEndpointsDisabled != nil {
		result["are_legacy_imds_endpoints_disabled"] = bool(*obj.AreLegacyImdsEndpointsDisabled)
	}

	return result
}

func (s *CoreInstanceResourceCrud) updateVnicAssignPublicIp(vnic *oci_core.Vnic, fieldKeyFormat string) error {

	if assignPublicIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")) {

		tmp := assignPublicIp.(string)
		assignPublicIpBoolVal, err := strconv.ParseBool(tmp)
		if err != nil {
			return err
		}

		if assignPublicIpBoolVal {

			listPrivateIpsResponse, err := s.VirtualNetworkClient.ListPrivateIps(context.Background(), oci_core.ListPrivateIpsRequest{
				VnicId: vnic.Id,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
				},
			})

			if err != nil {
				return err
			}

			for _, privateIp := range listPrivateIpsResponse.Items {
				if strings.EqualFold(*privateIp.IpAddress, *vnic.PrivateIp) {
					_, err = s.VirtualNetworkClient.CreatePublicIp(context.Background(), oci_core.CreatePublicIpRequest{
						CreatePublicIpDetails: oci_core.CreatePublicIpDetails{
							CompartmentId: vnic.CompartmentId,
							Lifetime:      oci_core.CreatePublicIpDetailsLifetimeEphemeral,
							PrivateIpId:   privateIp.Id,
						},
						RequestMetadata: common.RequestMetadata{
							RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
						},
					})
					return err
				}
			}

			return fmt.Errorf("unable to assign Ephemeral public ip for the vnic private ip: %s", *vnic.PrivateIp)

		} else {
			publicIpByIpAddressResponse, err := s.VirtualNetworkClient.GetPublicIpByIpAddress(context.Background(), oci_core.GetPublicIpByIpAddressRequest{
				GetPublicIpByIpAddressDetails: oci_core.GetPublicIpByIpAddressDetails{
					IpAddress: vnic.PublicIp,
				},
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
				},
			})

			if err == nil {
				_, err = s.VirtualNetworkClient.DeletePublicIp(context.Background(), oci_core.DeletePublicIpRequest{
					PublicIpId: publicIpByIpAddressResponse.Id,
					RequestMetadata: common.RequestMetadata{
						RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
					},
				})

				return err
			}
		}
	}

	return nil
}

func (s *CoreInstanceResourceCrud) mapToInstanceSourceDetails(fieldKeyFormat string) (oci_core.InstanceSourceDetails, error) {
	var baseObject oci_core.InstanceSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("bootVolume"):
		details := oci_core.InstanceSourceViaBootVolumeDetails{}
		if sourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_id")); ok {
			tmp := sourceId.(string)
			details.BootVolumeId = &tmp
		}
		baseObject = details
	case strings.ToLower("image"):
		details := oci_core.InstanceSourceViaImageDetails{}
		if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok {
			tmp := bootVolumeSizeInGBs.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("unable to convert bootVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.BootVolumeSizeInGBs = &tmpInt64
		}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if sourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_id")); ok {
			tmp := sourceId.(string)
			details.ImageId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func InstanceSourceDetailsToMap(obj *oci_core.InstanceSourceDetails, bootVolume *oci_core.BootVolume, sourceDetailsFromConfig map[string]interface{}) map[string]interface{} {
	// We need to use the values provided by the customer to prevent force new in case the service does not return the value
	result := sourceDetailsFromConfig
	if result == nil {
		result = map[string]interface{}{}
	}
	switch v := (*obj).(type) {
	case oci_core.InstanceSourceViaBootVolumeDetails:
		result["source_type"] = "bootVolume"

		if v.BootVolumeId != nil {
			result["source_id"] = string(*v.BootVolumeId)
		}
	case oci_core.InstanceSourceViaImageDetails:
		result["source_type"] = "image"

		if v.BootVolumeSizeInGBs != nil {
			result["boot_volume_size_in_gbs"] = strconv.FormatInt(*v.BootVolumeSizeInGBs, 10)
		} else if bootVolume != nil && bootVolume.SizeInGBs != nil {
			// The service could omit the boot volume size in the InstanceSourceViaImageDetails, so use the boot volume
			// SizeInGBs property if that's the case.
			result["boot_volume_size_in_gbs"] = strconv.FormatInt(*bootVolume.SizeInGBs, 10)
		}

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}

		if v.ImageId != nil {
			result["source_id"] = string(*v.ImageId)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceResourceCrud) mapToLaunchInstanceAgentConfigDetails(fieldKeyFormat string) (oci_core.LaunchInstanceAgentConfigDetails, error) {
	result := oci_core.LaunchInstanceAgentConfigDetails{}

	if areAllPluginsDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_all_plugins_disabled")); ok {
		tmp := areAllPluginsDisabled.(bool)
		result.AreAllPluginsDisabled = &tmp
	}

	if isManagementDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_management_disabled")); ok {
		tmp := isManagementDisabled.(bool)
		result.IsManagementDisabled = &tmp
	}

	if isMonitoringDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monitoring_disabled")); ok {
		tmp := isMonitoringDisabled.(bool)
		result.IsMonitoringDisabled = &tmp
	}

	if pluginsConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plugins_config")); ok {
		interfaces := pluginsConfig.([]interface{})
		tmp := make([]oci_core.InstanceAgentPluginConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "plugins_config"), stateDataIndex)
			converted, err := s.mapToInstanceAgentPluginConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "plugins_config")) {
			result.PluginsConfig = tmp
		}
	}

	return result, nil
}

func (s *CoreInstanceResourceCrud) mapToUpdateInstanceAgentConfigDetails(fieldKeyFormat string) (oci_core.UpdateInstanceAgentConfigDetails, error) {
	result := oci_core.UpdateInstanceAgentConfigDetails{}

	if areAllPluginsDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_all_plugins_disabled")); ok {
		tmp := areAllPluginsDisabled.(bool)
		result.AreAllPluginsDisabled = &tmp
	}

	if isMonitoringDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monitoring_disabled")); ok {
		tmp := isMonitoringDisabled.(bool)
		result.IsMonitoringDisabled = &tmp
	}

	if isManagementDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_management_disabled")); ok {
		tmp := isManagementDisabled.(bool)
		result.IsManagementDisabled = &tmp
	}

	if pluginsConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plugins_config")); ok {
		interfaces := pluginsConfig.([]interface{})
		tmp := make([]oci_core.InstanceAgentPluginConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "plugins_config"), stateDataIndex)
			converted, err := s.mapToInstanceAgentPluginConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "plugins_config")) {
			result.PluginsConfig = tmp
		}
	}

	return result, nil
}

func InstanceAgentConfigToMap(obj *oci_core.InstanceAgentConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreAllPluginsDisabled != nil {
		result["are_all_plugins_disabled"] = bool(*obj.AreAllPluginsDisabled)
	}

	if obj.IsManagementDisabled != nil {
		result["is_management_disabled"] = bool(*obj.IsManagementDisabled)
	}

	if obj.IsMonitoringDisabled != nil {
		result["is_monitoring_disabled"] = bool(*obj.IsMonitoringDisabled)
	}

	pluginsConfig := []interface{}{}
	for _, item := range obj.PluginsConfig {
		pluginsConfig = append(pluginsConfig, InstanceAgentPluginConfigDetailsToMap(item))
	}
	result["plugins_config"] = pluginsConfig

	return result
}

func (s *CoreInstanceResourceCrud) mapToLaunchInstanceAvailabilityConfigDetails(fieldKeyFormat string) (oci_core.LaunchInstanceAvailabilityConfigDetails, error) {
	result := oci_core.LaunchInstanceAvailabilityConfigDetails{}

	if isLiveMigrationPreferred, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_live_migration_preferred")); ok {
		tmp := isLiveMigrationPreferred.(bool)
		result.IsLiveMigrationPreferred = &tmp
	}

	if recoveryAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_action")); ok {
		result.RecoveryAction = oci_core.LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum(recoveryAction.(string))
	}

	return result, nil
}

func InstanceAvailabilityConfigToMap(obj *oci_core.InstanceAvailabilityConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsLiveMigrationPreferred != nil {
		result["is_live_migration_preferred"] = bool(*obj.IsLiveMigrationPreferred)
	}

	result["recovery_action"] = string(obj.RecoveryAction)

	return result
}

func (s *CoreInstanceResourceCrud) mapToLaunchInstancePlatformConfig(fieldKeyFormat string) (oci_core.LaunchInstancePlatformConfig, error) {
	var baseObject oci_core.LaunchInstancePlatformConfig
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("AMD_MILAN_BM"):
		details := oci_core.AmdMilanBmLaunchInstancePlatformConfig{}
		if numaNodesPerSocket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "numa_nodes_per_socket")); ok {
			details.NumaNodesPerSocket = oci_core.AmdMilanBmLaunchInstancePlatformConfigNumaNodesPerSocketEnum(numaNodesPerSocket.(string))
		}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("AMD_ROME_BM"):
		details := oci_core.AmdRomeBmLaunchInstancePlatformConfig{}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("AMD_VM"):
		details := oci_core.AmdVmLaunchInstancePlatformConfig{}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("INTEL_SKYLAKE_BM"):
		details := oci_core.IntelSkylakeBmLaunchInstancePlatformConfig{}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("INTEL_VM"):
		details := oci_core.IntelVmLaunchInstancePlatformConfig{}
		if isMeasuredBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_measured_boot_enabled")); ok {
			tmp := isMeasuredBootEnabled.(bool)
			details.IsMeasuredBootEnabled = &tmp
		}
		if isSecureBootEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_secure_boot_enabled")); ok {
			tmp := isSecureBootEnabled.(bool)
			details.IsSecureBootEnabled = &tmp
		}
		if isTrustedPlatformModuleEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_trusted_platform_module_enabled")); ok {
			tmp := isTrustedPlatformModuleEnabled.(bool)
			details.IsTrustedPlatformModuleEnabled = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func PlatformConfigToMap(obj *oci_core.PlatformConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.AmdMilanBmPlatformConfig:
		result["type"] = "AMD_MILAN_BM"

		result["numa_nodes_per_socket"] = string(v.NumaNodesPerSocket)

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.AmdRomeBmPlatformConfig:
		result["type"] = "AMD_ROME_BM"

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.AmdVmPlatformConfig:
		result["type"] = "AMD_VM"

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.IntelSkylakeBmPlatformConfig:
		result["type"] = "INTEL_SKYLAKE_BM"

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	case oci_core.IntelVmPlatformConfig:
		result["type"] = "INTEL_VM"

		if v.IsMeasuredBootEnabled != nil {
			result["is_measured_boot_enabled"] = bool(*v.IsMeasuredBootEnabled)
		}

		if v.IsSecureBootEnabled != nil {
			result["is_secure_boot_enabled"] = bool(*v.IsSecureBootEnabled)
		}

		if v.IsTrustedPlatformModuleEnabled != nil {
			result["is_trusted_platform_module_enabled"] = bool(*v.IsTrustedPlatformModuleEnabled)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceResourceCrud) mapToLaunchInstanceShapeConfigDetails(fieldKeyFormat string) (oci_core.LaunchInstanceShapeConfigDetails, error) {
	result := oci_core.LaunchInstanceShapeConfigDetails{}

	if baselineOcpuUtilization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "baseline_ocpu_utilization")); ok {
		result.BaselineOcpuUtilization = oci_core.LaunchInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum(baselineOcpuUtilization.(string))
	}

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

func (s *CoreInstanceResourceCrud) mapToUpdateInstanceShapeConfigDetails(fieldKeyFormat string) (oci_core.UpdateInstanceShapeConfigDetails, error) {
	result := oci_core.UpdateInstanceShapeConfigDetails{}

	if baselineOcpuUtilization, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "baseline_ocpu_utilization")); ok {
		result.BaselineOcpuUtilization = oci_core.UpdateInstanceShapeConfigDetailsBaselineOcpuUtilizationEnum(baselineOcpuUtilization.(string))
	}

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

func InstanceShapeConfigToMap(obj *oci_core.InstanceShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["baseline_ocpu_utilization"] = string(obj.BaselineOcpuUtilization)

	if obj.GpuDescription != nil {
		result["gpu_description"] = string(*obj.GpuDescription)
	}

	if obj.Gpus != nil {
		result["gpus"] = int(*obj.Gpus)
	}

	if obj.LocalDiskDescription != nil {
		result["local_disk_description"] = string(*obj.LocalDiskDescription)
	}

	if obj.LocalDisks != nil {
		result["local_disks"] = int(*obj.LocalDisks)
	}

	if obj.LocalDisksTotalSizeInGBs != nil {
		result["local_disks_total_size_in_gbs"] = float32(*obj.LocalDisksTotalSizeInGBs)
	}

	if obj.MaxVnicAttachments != nil {
		result["max_vnic_attachments"] = int(*obj.MaxVnicAttachments)
	}

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

func (s *CoreInstanceResourceCrud) mapToUpdateLaunchOptions(fieldKeyFormat string) (oci_core.UpdateLaunchOptions, error) {
	result := oci_core.UpdateLaunchOptions{}

	if bootVolumeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_type")); ok {
		result.BootVolumeType = oci_core.UpdateLaunchOptionsBootVolumeTypeEnum(bootVolumeType.(string))
	}

	if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
		tmp := isPvEncryptionInTransitEnabled.(bool)
		result.IsPvEncryptionInTransitEnabled = &tmp
	}

	if networkType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_type")); ok {
		result.NetworkType = oci_core.UpdateLaunchOptionsNetworkTypeEnum(networkType.(string))
	}

	return result, nil
}

func (s *CoreInstanceResourceCrud) mapToLaunchOptions(fieldKeyFormat string) (oci_core.LaunchOptions, error) {
	result := oci_core.LaunchOptions{}

	if bootVolumeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_type")); ok {
		result.BootVolumeType = oci_core.LaunchOptionsBootVolumeTypeEnum(bootVolumeType.(string))
	}

	if firmware, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "firmware")); ok {
		result.Firmware = oci_core.LaunchOptionsFirmwareEnum(firmware.(string))
	}

	if isConsistentVolumeNamingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_consistent_volume_naming_enabled")); ok {
		tmp := isConsistentVolumeNamingEnabled.(bool)
		result.IsConsistentVolumeNamingEnabled = &tmp
	}

	if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
		tmp := isPvEncryptionInTransitEnabled.(bool)
		result.IsPvEncryptionInTransitEnabled = &tmp
	}

	if networkType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_type")); ok {
		result.NetworkType = oci_core.LaunchOptionsNetworkTypeEnum(networkType.(string))
	}

	if remoteDataVolumeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remote_data_volume_type")); ok {
		result.RemoteDataVolumeType = oci_core.LaunchOptionsRemoteDataVolumeTypeEnum(remoteDataVolumeType.(string))
	}

	return result, nil
}

func mapToExtendedMetadata(rm map[string]interface{}) (map[string]interface{}, error) {
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

func (s *CoreInstanceResourceCrud) getPrimaryVnic() (*oci_core.Vnic, error) {
	request := oci_core.ListVnicAttachmentsRequest{
		CompartmentId: s.Res.CompartmentId,
		InstanceId:    s.Res.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
		},
	}
	var attachments []oci_core.VnicAttachment

	for {
		result, err := s.Client.ListVnicAttachments(context.Background(), request)
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, result.Items...)
		request.Page = result.OpcNextPage

		if request.Page == nil {
			break
		}
	}

	if len(attachments) < 1 {
		return nil, errors.New("No VNIC attachments found.")
	}

	for _, attachment := range attachments {
		if attachment.LifecycleState == oci_core.VnicAttachmentLifecycleStateAttached {
			request := oci_core.GetVnicRequest{
				VnicId: attachment.VnicId,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: tfresource.GetRetryPolicy(true, "core"),
				},
			}
			response, _ := s.VirtualNetworkClient.GetVnic(context.Background(), request)
			vnic := &response.Vnic

			// Ignore errors on GetVnic, since we might not have permissions to view some secondary VNICs.
			if vnic != nil && vnic.IsPrimary != nil && *vnic.IsPrimary {
				return vnic, nil
			}
		}
	}

	return nil, errors.New("Primary VNIC not found.")
}

func (s *CoreInstanceResourceCrud) getBootVolume() (*oci_core.BootVolume, error) {
	request := oci_core.ListBootVolumeAttachmentsRequest{
		AvailabilityDomain: s.Res.AvailabilityDomain,
		CompartmentId:      s.Res.CompartmentId,
		InstanceId:         s.Res.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
		},
	}

	response, err := s.Client.ListBootVolumeAttachments(context.Background(), request)
	if err != nil {
		return nil, err
	}

	if len(response.Items) < 1 {
		return nil, fmt.Errorf("Could not find any attached boot volumes")
	}

	bootVolumeId := response.Items[0].BootVolumeId
	if bootVolumeId == nil {
		return nil, fmt.Errorf("Found a boot volume attachment with no boot volume ID")
	}

	bootVolumeRequest := oci_core.GetBootVolumeRequest{
		BootVolumeId: bootVolumeId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
		},
	}
	bootVolumeResponse, err := s.BlockStorageClient.GetBootVolume(context.Background(), bootVolumeRequest)
	if err != nil {
		return nil, err
	}

	return &bootVolumeResponse.BootVolume, nil
}

func (s *CoreInstanceResourceCrud) mapToPreemptibleInstanceConfigDetails(fieldKeyFormat string) (oci_core.PreemptibleInstanceConfigDetails, error) {
	result := oci_core.PreemptibleInstanceConfigDetails{}

	if preemptionAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preemption_action")); ok {
		if tmpList := preemptionAction.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "preemption_action"), 0)
			tmp, err := s.mapToPreemptionAction(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert preemption_action, encountered error: %v", err)
			}
			result.PreemptionAction = tmp
		}
	}

	return result, nil
}

func PreemptibleInstanceConfigDetailsToMap(obj *oci_core.PreemptibleInstanceConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PreemptionAction != nil {
		preemptionActionArray := []interface{}{}
		if preemptionActionMap := PreemptionActionToMap(&obj.PreemptionAction); preemptionActionMap != nil {
			preemptionActionArray = append(preemptionActionArray, preemptionActionMap)
		}
		result["preemption_action"] = preemptionActionArray
	}

	return result
}

func (s *CoreInstanceResourceCrud) mapToPreemptionAction(fieldKeyFormat string) (oci_core.PreemptionAction, error) {
	var baseObject oci_core.PreemptionAction
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("TERMINATE"):
		details := oci_core.TerminatePreemptionAction{}
		if preserveBootVolume, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preserve_boot_volume")); ok {
			tmp := preserveBootVolume.(bool)
			details.PreserveBootVolume = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func PreemptionActionToMap(obj *oci_core.PreemptionAction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.TerminatePreemptionAction:
		result["type"] = "TERMINATE"

		if v.PreserveBootVolume != nil {
			result["preserve_boot_volume"] = bool(*v.PreserveBootVolume)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeInstanceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.InstanceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *CoreInstanceResourceCrud) mapToUpdateInstanceBootVolumeSizeInGbs(fieldKeyFormat string) error {
	if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")) {
		tmp := bootVolumeSizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert bootVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		err = s.updateBootVolumeSizeInGbs(tmpInt64)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CoreInstanceResourceCrud) updateBootVolumeSizeInGbs(bootVolumeSizeInGBs interface{}) error {
	changeBootVolumeDetailsRequest := oci_core.UpdateBootVolumeRequest{}

	if bootVolumeId, ok := s.D.GetOkExists("boot_volume_id"); ok {
		tmp := bootVolumeId.(string)
		changeBootVolumeDetailsRequest.BootVolumeId = &tmp
	}

	bootVolumeSizeInGBsTmp := bootVolumeSizeInGBs.(int64)
	changeBootVolumeDetailsRequest.SizeInGBs = &bootVolumeSizeInGBsTmp

	changeBootVolumeDetailsRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.BlockStorageClient.UpdateBootVolume(context.Background(), changeBootVolumeDetailsRequest)
	if err != nil {
		return err
	}

	_, err = waitForBootVolumeIfItIsUpdating(response.Id, s.BlockStorageClient, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return nil
}

func waitForBootVolumeIfItIsUpdating(bootVolumeID *string, client *oci_core.BlockstorageClient, timeout time.Duration) (*oci_core.GetBootVolumeResponse, error) {
	getBootVolumeRequest := oci_core.GetBootVolumeRequest{}

	getBootVolumeRequest.BootVolumeId = bootVolumeID

	bootVolumeUpdating := func(response common.OCIOperationResponse) bool {
		if getBootVolumeResponse, ok := response.Response.(oci_core.GetBootVolumeResponse); ok {
			if getBootVolumeResponse.LifecycleState == oci_core.BootVolumeLifecycleStateProvisioning {
				return true
			}
		}
		return false
	}

	getBootVolumeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicyWithAdditionalRetryCondition(timeout, bootVolumeUpdating, "core")
	getBootVolumeResponse, err := client.GetBootVolume(context.Background(), getBootVolumeRequest)
	return &getBootVolumeResponse, err
}

func (s *CoreInstanceResourceCrud) updateOptionsViaWorkRequest() error {
	request := oci_core.UpdateInstanceRequest{}

	if faultDomain, ok := s.D.GetOkExists("fault_domain"); ok && s.D.HasChange("fault_domain") {
		tmp := faultDomain.(string)
		request.FaultDomain = &tmp
	}

	if launchOptions, ok := s.D.GetOkExists("launch_options"); ok && s.D.HasChange("launch_options") {
		if tmpList := launchOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "launch_options", 0)
			tmp, err := s.mapToUpdateLaunchOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LaunchOptions = &tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok && s.D.HasChange("shape") {
		oldRaw, newRaw := s.D.GetChange("shape")
		if newRaw != "" && oldRaw != "" {
			shapeTmp := shape.(string)
			request.Shape = &shapeTmp
		}
		// the following if block is a temp solution and should be removed once service fixed on ther side
		if shapeConfig, ok := s.D.GetOkExists("shape_config"); strings.Contains(strings.ToLower(shape.(string)), "flex") && ok && !s.D.HasChange("shape_config") {
			if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shape_config", 0)
				tmp, err := s.mapToUpdateInstanceShapeConfigDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.ShapeConfig = &tmp
			}
		}
	}

	if shapeConfig, ok := s.D.GetOkExists("shape_config"); ok && s.D.HasChange("shape_config") {
		shape := s.D.Get("shape")
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 && strings.Contains(strings.ToLower(shape.(string)), "flex") {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shape_config", 0)
			tmp, err := s.mapToUpdateInstanceShapeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ShapeConfig = &tmp
		}
	}

	if request.Shape == nil && request.ShapeConfig == nil && request.LaunchOptions == nil && request.FaultDomain == nil {
		// no-op
		return nil
	}

	idTmp := s.D.Id()
	request.InstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstance(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "instance", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}
	return nil
}

func (s *CoreInstanceResourceCrud) mapToUpdateInstanceAvailabilityConfigDetails(fieldKeyFormat string) (oci_core.UpdateInstanceAvailabilityConfigDetails, error) {
	result := oci_core.UpdateInstanceAvailabilityConfigDetails{}

	if isLiveMigrationPreferred, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_live_migration_preferred")); ok {
		tmp := isLiveMigrationPreferred.(bool)
		result.IsLiveMigrationPreferred = &tmp
	}

	if recoveryAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_action")); ok {
		tmp := oci_core.UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum(recoveryAction.(string))
		result.RecoveryAction = tmp
	}

	return result, nil
}
