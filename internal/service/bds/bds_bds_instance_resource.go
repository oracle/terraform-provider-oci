// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func BdsBdsInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.TwelveHours,
			Update: &tfresource.TwelveHours,
			Delete: &tfresource.TwelveHours,
		},
		Create: createBdsBdsInstance,
		Read:   readBdsBdsInstance,
		Update: updateBdsBdsInstance,
		Delete: deleteBdsBdsInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"cluster_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"cluster_public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_high_availability": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"is_secure": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"is_force_remove_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"start_cluster_shape_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_type_shape_configs": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shape": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"master_node": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return ShapeChangeDiffSuppressFunction("master", d)
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"block_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						"number_of_nodes": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},

						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"nvmes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"util_node": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return ShapeChangeDiffSuppressFunction("util", d)
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"block_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						"number_of_nodes": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},

						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"nvmes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"worker_node": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return ShapeChangeDiffSuppressFunction("worker", d)
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"block_volume_size_in_gbs": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						"number_of_nodes": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(3),
						},

						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"nvmes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
					},
				},
			},
			"compute_only_worker_node": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return ShapeChangeDiffSuppressFunction("compute_only_worker", d)
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						"block_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						"number_of_nodes": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},

						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"nvmes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
					},
				},
			},
			"ignore_existing_nodes_shape": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"edge_node": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return ShapeChangeDiffSuppressFunction("edge", d)
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						"block_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						"number_of_nodes": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},

						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"nvmes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
					},
				},
			},
			"kafka_broker_node": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return ShapeChangeDiffSuppressFunction("kafka_broker", d)
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						"block_volume_size_in_gbs": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						"number_of_kafka_nodes": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"nvmes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
					},
				},
			},
			// Optional
			"bootstrap_script_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_sql_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"block_volume_size_in_gbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"nvmes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ocpus": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"shape": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_kerberos_mapped_to_database_users": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"kerberos_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"keytab_file": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"principal_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			// Optional
			"bds_cluster_version_summary": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bds_version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"odh_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"cluster_profile": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kerberos_realm_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"is_nat_gateway_required": {
							Type:     schema.TypeBool,
							Computed: true,
							Optional: true,
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
					string(oci_bds.BdsInstanceLifecycleStateInactive),
					string(oci_bds.BdsInstanceLifecycleStateActive),
				}, true),
			},
			"is_force_stop_jobs": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"remove_node": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_patch_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"cluster_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ambari_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bd_cell_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bda_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bdm_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bds_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"big_data_manager_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloudera_manager_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"csql_cell_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hue_server_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"jupyter_hub_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"odh_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_refreshed": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cloud_sql_configured": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"node_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"odh_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attached_block_volumes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"volume_attachment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"volume_size_in_gbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssh_fingerprint": {
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
						"ocpus": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"nvmes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"memory_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"time_maintenance_reboot_due": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_reboot_required": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"local_disks_total_size_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"os_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_kafka_configured": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"number_of_nodes": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_nodes_requiring_maintenance_reboot": {
				Type:     schema.TypeInt,
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

func createBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_bds.BdsInstanceLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_bds.BdsInstanceLifecycleStateInactive {
			powerOff = true
		}
	}

	cloudSqlRequest := oci_bds.AddCloudSqlRequest{}
	cloudSql := false

	if cloudSqlConfigured, ok := sync.D.GetOkExists("is_cloud_sql_configured"); ok {
		if cloudSqlConfigured.(bool) {
			cloudSql = true
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_sql_details", 0)
			if blockVolumeSizeInGBs, ok := sync.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_size_in_gbs")); ok {
				tmp := blockVolumeSizeInGBs.(string)
				tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return fmt.Errorf("unable to convert blockVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
				}
				cloudSqlRequest.BlockVolumeSizeInGBs = &tmpInt64
			} else {
				return fmt.Errorf("block_volume_size_in_gbs is required in cloud_sql_details")
			}

			if shape, ok := sync.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
				tmp := shape.(string)
				cloudSqlRequest.Shape = &tmp
			} else {
				return fmt.Errorf("shape is required in cloud_sql_details")
			}
		}
	}

	if err := tfresource.CreateResource(d, sync); err != nil {
		return err
	}

	if cloudSql {
		id := sync.D.Id()
		cloudSqlRequest.BdsInstanceId = &id
		if clusterAdminPassword, ok := sync.D.GetOkExists("cluster_admin_password"); ok {
			tmp := clusterAdminPassword.(string)
			cloudSqlRequest.ClusterAdminPassword = &tmp
		}
		if err := sync.addCloudSql(cloudSqlRequest); err != nil {
			return err
		}
		return tfresource.ReadResource(sync)
	}

	if powerOff {
		if err := sync.StopBdsInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_bds.BdsInstanceLifecycleStateInactive)
	}
	return nil
}

func readBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

func updateBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_bds.BdsInstanceLifecycleStateActive == oci_bds.BdsInstanceLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_bds.BdsInstanceLifecycleStateInactive == oci_bds.BdsInstanceLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartBdsInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_bds.BdsInstanceLifecycleStateActive)
	}

	if removeNode, ok := sync.D.GetOkExists("remove_node"); ok {
		if removeNode != "" {
			err := sync.RemoveNode()
			if err != nil {
				return err
			}
		}

	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopBdsInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_bds.BdsInstanceLifecycleStateInactive)
	}

	return nil
}

func deleteBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BdsBdsInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.BdsInstance
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BdsBdsInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateActive),
		string(oci_bds.BdsInstanceLifecycleStateFailed),
	}
}

func (s *BdsBdsInstanceResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateUpdating),
	}
}

func (s *BdsBdsInstanceResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.BdsInstanceLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceResourceCrud) Create() error {
	request := oci_bds.CreateBdsInstanceRequest{}

	if _, ok := s.D.GetOkExists("start_cluster_shape_configs"); ok {
		return fmt.Errorf("[ERROR] start_cluster_shape_configs is not permitted during create bds instance")
	}

	if _, ok := s.D.GetOkExists("is_force_remove_enabled"); ok {
		return fmt.Errorf("[ERROR] is_force_remove_enabled is not permitted during create bds instance")
	}

	if _, ok := s.D.GetOkExists("remove_node"); ok {
		return fmt.Errorf("[ERROR] remove_node is not permitted during create bds instance")
	}

	if bdsClusterVersionSummary, ok := s.D.GetOkExists("bds_cluster_version_summary"); ok {
		if tmpList := bdsClusterVersionSummary.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bds_cluster_version_summary", 0)
			tmp, err := s.mapToBdsClusterVersionSummary(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BdsClusterVersionSummary = &tmp
		}
	}

	if bootstrapScriptUrl, ok := s.D.GetOkExists("bootstrap_script_url"); ok {
		tmp := bootstrapScriptUrl.(string)
		request.BootstrapScriptUrl = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if clusterProfile, ok := s.D.GetOkExists("cluster_profile"); ok {
		request.ClusterProfile = oci_bds.BdsInstanceClusterProfileEnum(clusterProfile.(string))
	}

	if clusterPublicKey, ok := s.D.GetOkExists("cluster_public_key"); ok {
		tmp := clusterPublicKey.(string)
		request.ClusterPublicKey = &tmp
	}

	if clusterVersion, ok := s.D.GetOkExists("cluster_version"); ok {
		request.ClusterVersion = oci_bds.BdsInstanceClusterVersionEnum(clusterVersion.(string))
	}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isHighAvailability, ok := s.D.GetOkExists("is_high_availability"); ok {
		tmp := isHighAvailability.(bool)
		request.IsHighAvailability = &tmp
	}

	if isSecure, ok := s.D.GetOkExists("is_secure"); ok {
		tmp := isSecure.(bool)
		request.IsSecure = &tmp
	}

	if kerberosRealmName, ok := s.D.GetOkExists("kerberos_realm_name"); ok {
		tmp := kerberosRealmName.(string)
		request.KerberosRealmName = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if networkConfig, ok := s.D.GetOkExists("network_config"); ok {
		if tmpList := networkConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_config", 0)
			tmp, err := s.mapToNetworkConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfig = &tmp
		}
	}

	numOfNode := 0
	if _, ok := s.D.GetOkExists("master_node"); ok {
		fieldKey := fmt.Sprintf("%s.%d.%s", "master_node", 0, "number_of_nodes")
		if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
			numOfNode = numOfNode + numOfWorkers.(int)
		}
	}
	if _, ok := s.D.GetOkExists("util_node"); ok {
		fieldKey := fmt.Sprintf("%s.%d.%s", "util_node", 0, "number_of_nodes")
		if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
			numOfNode = numOfNode + numOfWorkers.(int)
		}
	}
	if _, ok := s.D.GetOkExists("worker_node"); ok {
		fieldKey := fmt.Sprintf("%s.%d.%s", "worker_node", 0, "number_of_nodes")
		if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
			numOfNode = numOfNode + numOfWorkers.(int)
		}
	}

	if clusterProfile, ok := s.D.GetOkExists("cluster_profile"); ok {
		if clusterProfile == "KAFKA" {
			if _, ok := s.D.GetOkExists("kafka_broker_node"); ok {
				fieldKey := fmt.Sprintf("%s.%d.%s", "kafka_broker_node", 0, "number_of_kafka_nodes")
				if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
					numOfNode = numOfNode + numOfWorkers.(int)
				}
			}
		}
	}

	createNodeDetails := make([]oci_bds.CreateNodeDetails, numOfNode)
	currentPos := 0

	if nodes, ok := s.D.GetOkExists("master_node"); ok {
		interfaces := nodes.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "master_node", stateDataIndex)
			converted, err := s.mapToCreateNodeDetails(fieldKeyFormat, "MASTER")
			if err != nil {
				return err
			}
			fieldKey := fmt.Sprintf(fieldKeyFormat, "number_of_nodes")
			if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
				for idx := 0; idx < numOfWorkers.(int); idx++ {
					createNodeDetails[currentPos] = converted
					currentPos = currentPos + 1
				}

			}
		}
	}
	if nodes, ok := s.D.GetOkExists("util_node"); ok {
		interfaces := nodes.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "util_node", stateDataIndex)
			converted, err := s.mapToCreateNodeDetails(fieldKeyFormat, "UTILITY")
			if err != nil {
				return err
			}

			fieldKey := fmt.Sprintf(fieldKeyFormat, "number_of_nodes")
			if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
				for idx := 0; idx < numOfWorkers.(int); idx++ {
					createNodeDetails[currentPos] = converted
					currentPos = currentPos + 1
				}

			}
		}
	}

	if nodes, ok := s.D.GetOkExists("worker_node"); ok {
		interfaces := nodes.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "worker_node", stateDataIndex)
			converted, err := s.mapToCreateNodeDetails(fieldKeyFormat, "WORKER")
			if err != nil {
				return err
			}

			fieldKey := fmt.Sprintf(fieldKeyFormat, "number_of_nodes")
			if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
				for idx := 0; idx < numOfWorkers.(int); idx++ {
					createNodeDetails[currentPos] = converted
					currentPos = currentPos + 1
				}

			}
		}
	}

	if clusterProfile, ok := s.D.GetOkExists("cluster_profile"); ok {
		if clusterProfile == "KAFKA" {
			if nodes, ok := s.D.GetOkExists("kafka_broker_node"); ok {
				interfaces := nodes.([]interface{})
				for i := range interfaces {
					stateDataIndex := i
					fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "kafka_broker_node", stateDataIndex)
					converted, err := s.mapToCreateNodeDetails(fieldKeyFormat, "KAFKA_BROKER")
					if err != nil {
						return err
					}
					fieldKey := fmt.Sprintf(fieldKeyFormat, "number_of_kafka_nodes")
					if numOfWorkers, ok := s.D.GetOkExists(fieldKey); ok {
						for idx := 0; idx < numOfWorkers.(int); idx++ {
							createNodeDetails[currentPos] = converted
							currentPos = currentPos + 1
						}

					}
				}
			}
		}
	}

	request.Nodes = createNodeDetails

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_bds.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_bds.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "bds") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	createResultError := s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
	if createResultError != nil {
		return createResultError
	}

	if clusterProfile, ok := s.D.GetOkExists("cluster_profile"); ok {
		isAddKafka, ok := s.D.GetOkExists("is_kafka_configured")
		if clusterProfile != "KAFKA" && isAddKafka == true && ok {
			if nodes, ok := s.D.GetOkExists("kafka_broker_node"); ok && nodes != nil {
				interfaces := nodes.([]interface{})
				if len(interfaces) == 0 {
					return fmt.Errorf("kafka broker node definition is missing")
				}
				err := s.AddKafka()
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("kafka broker node definition is missing")
			}
		}
	}

	_, computeWorkerAdditionError := s.updateComputeWorkersIfRequired()
	if computeWorkerAdditionError != nil {
		return computeWorkerAdditionError
	}

	_, edgeAdditionError := s.updateEdgeIfRequired()
	if edgeAdditionError != nil {
		return edgeAdditionError
	}
	return nil
}

func (s *BdsBdsInstanceResourceCrud) getBdsInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceId, err := bdsInstanceWaitForWorkRequest(workId, "bds",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceId)

	return s.Get()
}

func bdsInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func bdsInstanceWaitForWorkRequest(wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_bds.OperationStatusInProgress),
			string(oci_bds.OperationStatusAccepted),
			string(oci_bds.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bds.OperationStatusSucceeded),
			string(oci_bds.OperationStatusFailed),
			string(oci_bds.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_bds.GetWorkRequestRequest{
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
			if res.ActionType == action || res.ActionType == oci_bds.ActionTypesInProgress {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceWorkRequest(client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_bds.ListWorkRequestErrorsRequest{
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

func (s *BdsBdsInstanceResourceCrud) Get() error {
	request := oci_bds.GetBdsInstanceRequest{}

	tmp := s.D.Id()
	request.BdsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsInstance
	return nil
}

func (s *BdsBdsInstanceResourceCrud) Update() error {
	isKafkaBrokerAdded := false
	if cloudSqlConfigured, ok := s.D.GetOkExists("is_cloud_sql_configured"); ok && s.D.HasChange("is_cloud_sql_configured") {
		oldRaw, newRaw := s.D.GetChange("is_cloud_sql_configured")
		if newRaw != "" && oldRaw != "" {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_sql_details", 0)
			if cloudSqlConfigured.(bool) {
				request := oci_bds.AddCloudSqlRequest{}
				id := s.D.Id()
				request.BdsInstanceId = &id
				if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
					tmp := clusterAdminPassword.(string)
					request.ClusterAdminPassword = &tmp
				}
				if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
					tmp := shape.(string)
					request.Shape = &tmp
				} else {
					return fmt.Errorf("shape is required in cloud_sql_details")
				}
				if blockVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_size_in_gbs")); ok {
					tmp := blockVolumeSizeInGBs.(string)
					tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
					if err != nil {
						return fmt.Errorf("unable to convert blockVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
					}
					request.BlockVolumeSizeInGBs = &tmpInt64
				} else {
					return fmt.Errorf("block_volume_size_in_gbs is required in cloud_sql_details")
				}

				err := s.addCloudSql(request)
				if err != nil {
					return err
				}
			} else {
				request := oci_bds.RemoveCloudSqlRequest{}
				id := s.D.Id()
				request.BdsInstanceId = &id
				if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
					tmp := clusterAdminPassword.(string)
					request.ClusterAdminPassword = &tmp

				}
				err := s.deleteCloudSql(request)
				if err != nil {
					return err
				}

			}
		}
	}

	if nodes, ok := s.D.GetOkExists("kafka_broker_node"); ok && nodes != nil {
		interfaces := nodes.([]interface{})
		if len(interfaces) > 0 {
			kafkaConfigured, ok := s.D.GetOkExists("is_kafka_configured")
			if ok && !kafkaConfigured.(bool) {
				return fmt.Errorf("Found a kafka broker node definition without is_kafka_configured set to true")
			}
		}
	}

	// kafka
	if kafkaConfigured, ok := s.D.GetOkExists("is_kafka_configured"); ok && s.D.HasChange("is_kafka_configured") {
		oldRaw, newRaw := s.D.GetChange("is_kafka_configured")
		if newRaw != "" && oldRaw != "" {
			if kafkaConfigured.(bool) {
				if nodes, ok := s.D.GetOkExists("kafka_broker_node"); ok && nodes != nil {
					interfaces := nodes.([]interface{})
					if len(interfaces) == 0 {
						return fmt.Errorf("kafka broker node definition is missing")
					}
					err := s.AddKafka()
					isKafkaBrokerAdded = true
					if err != nil {
						return err
					}
				} else {
					return fmt.Errorf("kafka broker node definition is missing")
				}
			} else {
				if clusterProfile, ok := s.D.GetOkExists("cluster_profile"); ok {
					if clusterProfile == "KAFKA" {
						if nodes, ok := s.D.GetOkExists("kafka_broker_node"); ok && nodes != nil {
							interfaces := nodes.([]interface{})
							if len(interfaces) > 0 {
								return fmt.Errorf("remove_kafka operation not permitted")
							}
						}
					}
				}
				err := s.RemoveKafka()
				isKafkaBrokerAdded = true
				if err != nil {
					return err
				}
			}
		}
	} else {
		isKafkaBrokerAdded1, kafkaBrokerErr := s.updateKafkaBrokerIfRequired()
		if kafkaBrokerErr != nil {
			return kafkaBrokerErr
		}
		isKafkaBrokerAdded = isKafkaBrokerAdded1
	}

	if _, ok := s.D.GetOkExists("bootstrap_script_url"); ok {
		err := s.ExecuteBootstrapScript()
		if err != nil {
			return err
		}
	}

	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	workerNodeFieldKeyFormat := "worker_node.0.%s"
	masterNodeFieldKeyFormat := "master_node.0.%s"
	utilNodeFieldKeyFormat := "util_node.0.%s"
	computeOnlyWorkerNodeFieldKeyFormat := "compute_only_worker_node.0.%s"
	edgeNodeFieldKeyFormat := "edge_node.0.%s"
	cloudSqlNodeFieldKeyFormat := "cloud_sql_details.0.%s"
	kafkaBrokerNodeFieldKeyFormat := "kafka_broker_node.0.%s"

	_, blockVolumeSizeInGbsPresent := s.D.GetOkExists(fmt.Sprintf(workerNodeFieldKeyFormat, "block_volume_size_in_gbs"))
	if blockVolumeSizeInGbsPresent && s.D.HasChange(fmt.Sprintf(workerNodeFieldKeyFormat, "block_volume_size_in_gbs")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(workerNodeFieldKeyFormat, "block_volume_size_in_gbs"))

		tmpOld := oldRaw.(string)
		tmpInt64Old, err := strconv.ParseInt(tmpOld, 10, 64)
		if err != nil {
			return err
		}

		tmpNew := newRaw.(string)
		tmpInt64New, err := strconv.ParseInt(tmpNew, 10, 64)

		if err != nil {
			return err
		}

		if tmpInt64New > tmpInt64Old {
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				dif := tmpInt64New - tmpInt64Old
				err := s.updateWorkerBlockStorage(s.D.Id(), clusterAdminPassword, dif, oci_bds.AddBlockStorageDetailsNodeTypeWorker)
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("the new value should be larger than previous one")
		}
	}

	_, kafkaBlockVolumeSizeInGbsPresent := s.D.GetOkExists(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "block_volume_size_in_gbs"))
	if kafkaBlockVolumeSizeInGbsPresent && s.D.HasChange(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "block_volume_size_in_gbs")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "block_volume_size_in_gbs"))
		if oldRaw != "" {
			tmpOld := oldRaw.(string)
			tmpInt64Old, err := strconv.ParseInt(tmpOld, 10, 64)
			if err != nil {
				return err
			}

			tmpNew := newRaw.(string)
			tmpInt64New, err := strconv.ParseInt(tmpNew, 10, 64)

			if err != nil {
				return err
			}

			if tmpInt64New > tmpInt64Old {
				if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
					dif := tmpInt64New - tmpInt64Old
					err := s.updateWorkerBlockStorage(s.D.Id(), clusterAdminPassword, dif, oci_bds.AddBlockStorageDetailsNodeTypeKafkaBroker)
					if err != nil {
						return err
					}
				}
			} else {
				return fmt.Errorf("the new value should be larger than previous one")
			}
		}
	}
	_, numOfWorkersPresent := s.D.GetOkExists(fmt.Sprintf(workerNodeFieldKeyFormat, "number_of_nodes"))
	if numOfWorkersPresent && s.D.HasChange(fmt.Sprintf(workerNodeFieldKeyFormat, "number_of_nodes")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(workerNodeFieldKeyFormat, "number_of_nodes"))
		tmpOld := oldRaw.(int)
		tmpNew := newRaw.(int)
		if tmpNew > tmpOld {
			var blockVolumeSizeInGBInt64 int64
			var workerNodeShapeStr string
			var err error
			blockVolumeSizeInGBs := s.D.Get("worker_node.0.block_volume_size_in_gbs")
			if blockVolumeSizeInGBs != "" {
				tmp := blockVolumeSizeInGBs.(string)
				blockVolumeSizeInGBInt64, err = strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return err
				}
			}
			if workerNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(workerNodeFieldKeyFormat, "shape")); ok {
				workerNodeShapeStr = workerNodeShape.(string)
				if err != nil {
					return err
				}
			}
			workerShapeConfig, _ := s.mapToShapeConfigDetails("worker_node.0.shape_config.0.%s")
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				if blockVolumeSizeInGBInt64 != 0 {
					err = s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, oci_bds.AddWorkerNodesDetailsNodeTypeWorker, &blockVolumeSizeInGBInt64, &workerNodeShapeStr, &workerShapeConfig)
				} else {
					err = s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, oci_bds.AddWorkerNodesDetailsNodeTypeWorker, nil, &workerNodeShapeStr, &workerShapeConfig)
				}
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("cluster admin password not provided")
			}
		} else {
			return fmt.Errorf("the new value should be larger than previous one")
		}
	}

	//	 ADD MASTER AND UTILITY NODES
	_, numOfMastersPresent := s.D.GetOkExists(fmt.Sprintf(masterNodeFieldKeyFormat, "number_of_nodes"))
	if numOfMastersPresent && s.D.HasChange(fmt.Sprintf(masterNodeFieldKeyFormat, "number_of_nodes")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(masterNodeFieldKeyFormat, "number_of_nodes"))
		tmpOld := oldRaw.(int)
		tmpNew := newRaw.(int)
		if tmpNew > tmpOld {
			var blockVolumeSizeInGBInt64 int64
			var masterNodeShapeStr string
			var err error
			blockVolumeSizeInGBs := s.D.Get("master_node.0.block_volume_size_in_gbs")
			if blockVolumeSizeInGBs != "" {
				tmp := blockVolumeSizeInGBs.(string)
				blockVolumeSizeInGBInt64, err = strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return err
				}
			}
			if masterNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(masterNodeFieldKeyFormat, "shape")); ok {
				masterNodeShapeStr = masterNodeShape.(string)
				if err != nil {
					return err
				}
			}
			masterShapeConfig, _ := s.mapToShapeConfigDetails("master_node.0.shape_config.0.%s")
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				if blockVolumeSizeInGBInt64 != 0 {
					err = s.updateMasterNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, &blockVolumeSizeInGBInt64, &masterNodeShapeStr, &masterShapeConfig)
				} else {
					err = s.updateMasterNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, nil, &masterNodeShapeStr, &masterShapeConfig)
				}
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("cluster admin password not provided")
			}
		} else {
			return fmt.Errorf("the new value should be larger than previous one")
		}
	}
	_, numOfUtilityPresent := s.D.GetOkExists(fmt.Sprintf(utilNodeFieldKeyFormat, "number_of_nodes"))
	if numOfUtilityPresent && s.D.HasChange(fmt.Sprintf(utilNodeFieldKeyFormat, "number_of_nodes")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(utilNodeFieldKeyFormat, "number_of_nodes"))
		tmpOld := oldRaw.(int)
		tmpNew := newRaw.(int)
		if tmpNew > tmpOld {
			var blockVolumeSizeInGBInt64 int64
			var utilNodeShapeStr string
			var err error
			blockVolumeSizeInGBs := s.D.Get("util_node.0.block_volume_size_in_gbs")
			if blockVolumeSizeInGBs != "" {
				tmp := blockVolumeSizeInGBs.(string)
				blockVolumeSizeInGBInt64, err = strconv.ParseInt(tmp, 10, 64)
				if err != nil {
					return err
				}
			}
			if utilNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(utilNodeFieldKeyFormat, "shape")); ok {
				utilNodeShapeStr = utilNodeShape.(string)
				if err != nil {
					return err
				}
			}
			utilShapeConfig, _ := s.mapToShapeConfigDetails("util_node.0.shape_config.0.%s")
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				if blockVolumeSizeInGBInt64 != 0 {
					err = s.updateUtilityNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, &blockVolumeSizeInGBInt64, &utilNodeShapeStr, &utilShapeConfig)
				} else {
					err = s.updateUtilityNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, nil, &utilNodeShapeStr, &utilShapeConfig)
				}
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("cluster admin password not provided")
			}
		} else {
			return fmt.Errorf("the new value should be larger than previous one")
		}
	}

	isComputeWorkerAdded, computeWorkerErr := s.updateComputeWorkersIfRequired()
	if computeWorkerErr != nil {
		return computeWorkerErr
	}

	isEdgeAdded, edgeErr := s.updateEdgeIfRequired()
	if edgeErr != nil {
		return edgeErr
	}

	result := oci_bds.ChangeShapeNodes{}

	changeShapeRequest := oci_bds.ChangeShapeRequest{}
	var ignoreMasterShape, ignoreUtilShape, ignoreWorkerShape, ignoreComputeWorkerShape, ignoreEdgeShape, ignoreKafkaBrokerShape = false, false, false, false, false, false
	if ignoreExistingNodesShape, ok := s.D.GetOkExists("ignore_existing_nodes_shape"); ok {
		interfaces := ignoreExistingNodesShape.([]interface{})
		tmp := make([]string, len(interfaces))
		// Add now node types when they are released
		for i := range interfaces {
			tmp[i] = strings.TrimSpace(strings.ToLower(interfaces[i].(string)))
			if tmp[i] == "master" {
				ignoreMasterShape = true
			}
			if tmp[i] == "utility" {
				ignoreUtilShape = true
			}
			if tmp[i] == "worker" {
				ignoreWorkerShape = true
			}
			if tmp[i] == "compute_only_worker" {
				ignoreComputeWorkerShape = true
			}
			if tmp[i] == "edge" {
				ignoreEdgeShape = true
			}
			if tmp[i] == "kafka_broker" {
				ignoreKafkaBrokerShape = true
			}
		}
		s.D.Set("ignore_existing_nodes_shape", ignoreExistingNodesShape)
	}
	if ignoreWorkerShape == false {
		workerNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(workerNodeFieldKeyFormat, "shape"))
		if ok && (s.D.HasChange(fmt.Sprintf(workerNodeFieldKeyFormat, "shape")) ||
			s.D.HasChange(fmt.Sprintf(workerNodeFieldKeyFormat, "shape_config"))) {
			tmp := workerNodeShape.(string)
			result.Worker = &tmp
			if nodeConfig, ok := s.D.GetOkExists("worker_node.0.shape_config"); ok && len(nodeConfig.([]interface{})) != 0 {
				workerShapeConfig, _ := s.mapToShapeConfigDetails("worker_node.0.shape_config.0.%s")
				result.WorkerShapeConfig = &workerShapeConfig
			}
		}
	}
	if ignoreMasterShape == false {
		masterNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(masterNodeFieldKeyFormat, "shape"))
		if ok && (s.D.HasChange(fmt.Sprintf(masterNodeFieldKeyFormat, "shape")) ||
			s.D.HasChange(fmt.Sprintf(masterNodeFieldKeyFormat, "shape_config"))) {
			tmp := masterNodeShape.(string)
			result.Master = &tmp
			if nodeConfig, ok := s.D.GetOkExists("master_node.0.shape_config"); ok && len(nodeConfig.([]interface{})) != 0 {
				masterShapeConfig, _ := s.mapToShapeConfigDetails("master_node.0.shape_config.0.%s")
				result.MasterShapeConfig = &masterShapeConfig
			}
		}
	}
	if ignoreUtilShape == false {
		utilNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(utilNodeFieldKeyFormat, "shape"))
		if ok && (s.D.HasChange(fmt.Sprintf(utilNodeFieldKeyFormat, "shape")) ||
			s.D.HasChange(fmt.Sprintf(utilNodeFieldKeyFormat, "shape_config"))) {
			tmp := utilNodeShape.(string)
			result.Utility = &tmp
			if nodeConfig, ok := s.D.GetOkExists("util_node.0.shape_config"); ok && len(nodeConfig.([]interface{})) != 0 {
				utilShapeConfig, _ := s.mapToShapeConfigDetails("util_node.0.shape_config.0.%s")
				result.UtilityShapeConfig = &utilShapeConfig
			}
		}
	}
	if ignoreComputeWorkerShape == false {
		computeWorker, ok := s.D.GetOkExists(fmt.Sprintf(computeOnlyWorkerNodeFieldKeyFormat, "shape"))
		if ok && (!isComputeWorkerAdded) && (s.D.HasChange(fmt.Sprintf(computeOnlyWorkerNodeFieldKeyFormat, "shape")) ||
			s.D.HasChange(fmt.Sprintf(computeOnlyWorkerNodeFieldKeyFormat, "shape_config"))) {
			tmp := computeWorker.(string)
			result.ComputeOnlyWorker = &tmp
			if nodeConfig, ok := s.D.GetOkExists("compute_only_worker_node.0.shape_config"); ok && len(nodeConfig.([]interface{})) != 0 {
				computeWorkerShapeConfig, _ := s.mapToShapeConfigDetails("compute_only_worker_node.0.shape_config.0.%s")
				result.ComputeOnlyWorkerShapeConfig = &computeWorkerShapeConfig
			}
		}
	}
	if ignoreEdgeShape == false {
		edge, ok := s.D.GetOkExists(fmt.Sprintf(edgeNodeFieldKeyFormat, "shape"))
		if ok && (!isEdgeAdded) && (s.D.HasChange(fmt.Sprintf(edgeNodeFieldKeyFormat, "shape")) ||
			s.D.HasChange(fmt.Sprintf(edgeNodeFieldKeyFormat, "shape_config"))) {
			tmp := edge.(string)
			result.Edge = &tmp
			if nodeConfig, ok := s.D.GetOkExists("edge_node.0.shape_config"); ok && len(nodeConfig.([]interface{})) != 0 {
				edgeShapeConfig, _ := s.mapToShapeConfigDetails("edge_node.0.shape_config.0.%s")
				result.EdgeShapeConfig = &edgeShapeConfig
			}
		}
	}
	if ignoreKafkaBrokerShape == false {
		kafkaBroker, ok := s.D.GetOkExists(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "shape"))
		if ok && (!isKafkaBrokerAdded) && (s.D.HasChange(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "shape")) ||
			s.D.HasChange(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "shape_config"))) {
			tmp := kafkaBroker.(string)
			result.KafkaBroker = &tmp
			if nodeConfig, ok := s.D.GetOkExists("kafka_broker_node.0.shape_config"); ok && len(nodeConfig.([]interface{})) != 0 {
				kafkaBrokerShapeConfig, _ := s.mapToShapeConfigDetails("kafka_broker_node.0.shape_config.0.%s")
				result.KafkaBrokerShapeConfig = &kafkaBrokerShapeConfig
			}
		}
	}
	if _, ok := s.D.GetOkExists("is_cloud_sql_configured"); ok {
		cloudSqlNodeShape, ok := s.D.GetOkExists(fmt.Sprintf(cloudSqlNodeFieldKeyFormat, "shape"))
		if ok && s.D.HasChange(fmt.Sprintf(cloudSqlNodeFieldKeyFormat, "shape")) {
			tmp := cloudSqlNodeShape.(string)
			result.Cloudsql = &tmp
		}
	}
	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		clusterAdminPasswordTmp := clusterAdminPassword.(string)
		changeShapeRequest.ClusterAdminPassword = &clusterAdminPasswordTmp

		changeShapeRequest.Nodes = &result
		if !reflect.DeepEqual(result, oci_bds.ChangeShapeNodes{}) {
			tmp := s.D.Id()
			changeShapeRequest.BdsInstanceId = &tmp

			response, err := s.Client.ChangeShape(context.Background(), changeShapeRequest)
			if err != nil {
				return err
			}

			workId := response.OpcWorkRequestId
			err = s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return err
			}
		}
	}

	request := oci_bds.UpdateBdsInstanceRequest{}

	tmp := s.D.Id()
	request.BdsInstanceId = &tmp

	if bootstrapScriptUrl, ok := s.D.GetOkExists("bootstrap_script_url"); ok {
		tmp := bootstrapScriptUrl.(string)
		request.BootstrapScriptUrl = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if networkConfig, ok := s.D.GetOkExists("network_config"); ok && s.D.HasChange("network_config") {
		if tmpList := networkConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_config", 0)
			tmp, err := s.mapToNetworkConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	if s.D.HasChange("bootstrap_script_url") || s.D.HasChange("defined_tags") || s.D.HasChange("display_name") || s.D.HasChange("freeform_tags") || s.D.HasChange("kms_key_id") || s.D.HasChange("network_config") {
		response, err := s.Client.UpdateBdsInstance(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
	}
	return nil
}

func (s *BdsBdsInstanceResourceCrud) updateComputeWorkersIfRequired() (bool, error) {
	areWorkersAdded := false
	computeOnlyWorkerNodeFieldKeyFormat := "compute_only_worker_node.0.%s"
	var computeWorkerBlockVolumeSizeGBInt64 int64
	var computeWorkerBlockVolumeConversionError, err error
	computeWorkerBlockVolumeSizeInGBs := s.D.Get("edge_node.0.block_volume_size_in_gbs")
	if computeWorkerBlockVolumeSizeInGBs != "" {
		computeWorkerBlockVolumeSizeGBInt64, computeWorkerBlockVolumeConversionError = strconv.ParseInt(computeWorkerBlockVolumeSizeInGBs.(string), 10, 64)
		if computeWorkerBlockVolumeConversionError != nil {
			return false, computeWorkerBlockVolumeConversionError
		}
	}
	compute_worker_shape, _ := s.D.GetOkExists("compute_only_worker_node.0.shape")
	compute_worker_shape_string := compute_worker_shape.(string)
	compute_worker_shape_config, _ := s.mapToShapeConfigDetails("compute_only_worker_node.0.shape_config.0.%s")
	_, numOfComputeOnlyWorkersPresent := s.D.GetOkExists(fmt.Sprintf(computeOnlyWorkerNodeFieldKeyFormat, "number_of_nodes"))
	if numOfComputeOnlyWorkersPresent && s.D.HasChange(fmt.Sprintf(computeOnlyWorkerNodeFieldKeyFormat, "number_of_nodes")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(computeOnlyWorkerNodeFieldKeyFormat, "number_of_nodes"))
		var tmpOld = 0
		if oldRaw != nil {
			tmpOld = oldRaw.(int)
		}
		tmpNew := newRaw.(int)
		if tmpNew > tmpOld {
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				if computeWorkerBlockVolumeSizeGBInt64 != 0 {
					err = s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, oci_bds.AddWorkerNodesDetailsNodeTypeComputeOnlyWorker, &computeWorkerBlockVolumeSizeGBInt64, &compute_worker_shape_string, &compute_worker_shape_config)
				} else {
					err = s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, oci_bds.AddWorkerNodesDetailsNodeTypeComputeOnlyWorker, nil, &compute_worker_shape_string, &compute_worker_shape_config)
				}
				if err != nil {
					return false, err
				}
				areWorkersAdded = true
			} else {
				return false, fmt.Errorf("cluster admin password not provided")
			}
		} else {
			return false, fmt.Errorf("the new number of compute only worker node should be larger than previous one")
		}
	}
	return areWorkersAdded, nil
}

func (s *BdsBdsInstanceResourceCrud) updateEdgeIfRequired() (bool, error) {
	areEdgeAdded := false
	edgeNodeFieldKeyFormat := "edge_node.0.%s"
	var edgeBlockVolumeSizeInGBInt64 int64
	var edgeBlockVolumeConversionError, err error
	edgeBlockVolumeSizeInGBs := s.D.Get("edge_node.0.block_volume_size_in_gbs")
	if edgeBlockVolumeSizeInGBs != "" {
		edgeBlockVolumeSizeInGBInt64, edgeBlockVolumeConversionError = strconv.ParseInt(edgeBlockVolumeSizeInGBs.(string), 10, 64)
		if edgeBlockVolumeConversionError != nil {
			return false, edgeBlockVolumeConversionError
		}
	}
	edge_shape, _ := s.D.GetOkExists("edge_node.0.shape")
	edge_shape_string := edge_shape.(string)
	edge_shape_config, _ := s.mapToShapeConfigDetails("edge_node.0.shape_config.0.%s")
	_, numOfEdgePresent := s.D.GetOkExists(fmt.Sprintf(edgeNodeFieldKeyFormat, "number_of_nodes"))
	if numOfEdgePresent && s.D.HasChange(fmt.Sprintf(edgeNodeFieldKeyFormat, "number_of_nodes")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(edgeNodeFieldKeyFormat, "number_of_nodes"))
		var tmpOld = 0
		if oldRaw != nil {
			tmpOld = oldRaw.(int)
		}
		tmpNew := newRaw.(int)
		if tmpNew > tmpOld {
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				if edgeBlockVolumeSizeInGBInt64 != 0 {
					err = s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, oci_bds.AddWorkerNodesDetailsNodeTypeEdge, &edgeBlockVolumeSizeInGBInt64, &edge_shape_string, &edge_shape_config)
				} else {
					err = s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, oci_bds.AddWorkerNodesDetailsNodeTypeEdge, nil, &edge_shape_string, &edge_shape_config)
				}
				if err != nil {
					return false, err
				}

				areEdgeAdded = true
			} else {
				return false, fmt.Errorf("cluster admin password not provided")
			}
		} else {
			return false, fmt.Errorf("the new number of edge node should be larger than previous one")
		}
	}
	return areEdgeAdded, nil
}

func (s *BdsBdsInstanceResourceCrud) updateKafkaBrokerIfRequired() (bool, error) {
	areKafkaBrokerAdded := false
	kafkaBrokerNodeFieldKeyFormat := "kafka_broker_node.0.%s"
	var kafkaBrokerBlockVolumeSizeGBInt64 int64
	var kafkaBrokerBlockVolumeConversionError, err error
	kafkaBrokerBlockVolumeSizeInGBs := s.D.Get("kafka_broker_node.0.block_volume_size_in_gbs")
	if kafkaBrokerBlockVolumeSizeInGBs != "" {
		kafkaBrokerBlockVolumeSizeGBInt64, kafkaBrokerBlockVolumeConversionError = strconv.ParseInt(kafkaBrokerBlockVolumeSizeInGBs.(string), 10, 64)
		if kafkaBrokerBlockVolumeConversionError != nil {
			return false, kafkaBrokerBlockVolumeConversionError
		}
	}
	kafka_broker_shape, _ := s.D.GetOkExists("kafka_broker_node.0.shape")
	kafka_broker_shape_string := kafka_broker_shape.(string)
	kafka_broker_shape_config, _ := s.mapToShapeConfigDetails("kafka_broker_node.0.shape_config.0.%s")
	_, numOfKafkaBrokersPresent := s.D.GetOkExists(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "number_of_kafka_nodes"))
	if numOfKafkaBrokersPresent && s.D.HasChange(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "number_of_kafka_nodes")) {
		oldRaw, newRaw := s.D.GetChange(fmt.Sprintf(kafkaBrokerNodeFieldKeyFormat, "number_of_kafka_nodes"))
		var tmpOld = 0
		if oldRaw != nil {
			tmpOld = oldRaw.(int)
		}
		tmpNew := newRaw.(int)
		if tmpNew > tmpOld {
			if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
				if kafkaBrokerBlockVolumeSizeGBInt64 != 0 {
					err = s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, oci_bds.AddWorkerNodesDetailsNodeTypeKafkaBroker, &kafkaBrokerBlockVolumeSizeGBInt64, &kafka_broker_shape_string, &kafka_broker_shape_config)
				} else {
					err = s.updateWorkerNode(s.D.Id(), clusterAdminPassword, tmpNew-tmpOld, oci_bds.AddWorkerNodesDetailsNodeTypeKafkaBroker, nil, &kafka_broker_shape_string, &kafka_broker_shape_config)
				}
				if err != nil {
					return false, err
				}
				areKafkaBrokerAdded = true
			} else {
				return false, fmt.Errorf("cluster admin password not provided")
			}
		} else {
			return false, fmt.Errorf("the new number of kafka broker node should be larger than previous one")
		}
	}
	return areKafkaBrokerAdded, nil
}

func (s *BdsBdsInstanceResourceCrud) Delete() error {
	request := oci_bds.DeleteBdsInstanceRequest{}

	tmp := s.D.Id()
	request.BdsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeleteBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bdsInstanceWaitForWorkRequest(workId, "bds",
		oci_bds.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BdsBdsInstanceResourceCrud) SetData() error {
	if s.Res.BdsClusterVersionSummary != nil {
		s.D.Set("bds_cluster_version_summary", []interface{}{BdsClusterVersionSummaryToMap(s.Res.BdsClusterVersionSummary)})
	} else {
		s.D.Set("bds_cluster_version_summary", nil)
	}

	if s.Res.BootstrapScriptUrl != nil {
		s.D.Set("bootstrap_script_url", *s.Res.BootstrapScriptUrl)
	}

	if s.Res.IsCloudSqlConfigured != nil {
		s.D.Set("is_cloud_sql_configured", *s.Res.IsCloudSqlConfigured)
	}

	if s.Res.CloudSqlDetails != nil {
		s.D.Set("cloud_sql_details", []interface{}{CloudSqlDetailsToMap(s.Res.CloudSqlDetails)})
	} else {
		s.D.Set("cloud_sql_details", []interface{}{})
		s.D.Set("is_cloud_sql_configured", false)
	}

	if s.Res.ClusterDetails != nil {
		s.D.Set("cluster_details", []interface{}{ClusterDetailsToMap(s.Res.ClusterDetails)})
	} else {
		s.D.Set("cluster_details", nil)
	}

	s.D.Set("cluster_profile", s.Res.ClusterProfile)

	s.D.Set("cluster_version", s.Res.ClusterVersion)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsHighAvailability != nil {
		s.D.Set("is_high_availability", *s.Res.IsHighAvailability)
	}

	if s.Res.IsKafkaConfigured != nil {
		s.D.Set("is_kafka_configured", *s.Res.IsKafkaConfigured)
	}

	if s.Res.IsSecure != nil {
		s.D.Set("is_secure", *s.Res.IsSecure)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.NetworkConfig != nil {
		s.D.Set("network_config", []interface{}{NetworkConfigToMap(s.Res.NetworkConfig)})
	} else {
		s.D.Set("network_config", nil)
	}

	nodes := []interface{}{}
	nodeMap := make(map[string]map[string]interface{})
	for _, item := range s.Res.Nodes {
		node := BdsNodeToMap(item)
		nodes = append(nodes, node)
		PopulateNodeTemplate(item, nodeMap)
	}
	masterNodeConfig := nodeMap["MASTER"]
	utilNodeConfig := nodeMap["UTILITY"]
	workerNodeConfig := nodeMap["WORKER"]
	computeOnlyWorkerNodeConfig := nodeMap["COMPUTE_ONLY_WORKER"]
	edgeNodeConfig := nodeMap["EDGE"]
	kafkaBrokerNodeConfig := nodeMap["KAFKA_BROKER"]
	s.deleteShapeConfigIfMissingInInput("master_node", masterNodeConfig)
	s.deleteShapeConfigIfMissingInInput("util_node", utilNodeConfig)
	s.deleteShapeConfigIfMissingInInput("worker_node", workerNodeConfig)
	s.deleteShapeConfigIfMissingInInput("compute_only_worker_node", computeOnlyWorkerNodeConfig)
	s.deleteShapeConfigIfMissingInInput("edge", edgeNodeConfig)
	s.deleteShapeConfigIfMissingInInput("kafka_broker_node", kafkaBrokerNodeConfig)
	s.D.Set("nodes", nodes)
	s.D.Set("master_node", []interface{}{masterNodeConfig})
	s.D.Set("util_node", []interface{}{utilNodeConfig})
	s.D.Set("worker_node", []interface{}{workerNodeConfig})

	if _, ok := nodeMap["COMPUTE_ONLY_WORKER"]; ok {
		s.D.Set("compute_only_worker_node", []interface{}{computeOnlyWorkerNodeConfig})
	}

	if _, ok := nodeMap["EDGE"]; ok {
		s.D.Set("edge_node", []interface{}{edgeNodeConfig})
	}

	if _, ok := nodeMap["KAFKA_BROKER"]; ok {
		s.D.Set("kafka_broker_node", []interface{}{kafkaBrokerNodeConfig})
	}

	if s.Res.NumberOfNodes != nil {
		s.D.Set("number_of_nodes", *s.Res.NumberOfNodes)
	}

	if s.Res.NumberOfNodesRequiringMaintenanceReboot != nil {
		s.D.Set("number_of_nodes_requiring_maintenance_reboot", *s.Res.NumberOfNodesRequiringMaintenanceReboot)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *BdsBdsInstanceResourceCrud) StartBdsInstance() error {
	request := oci_bds.StartBdsInstanceRequest{}

	idTmp := s.D.Id()
	request.BdsInstanceId = &idTmp

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if startClusterShapeConfigs, ok := s.D.GetOkExists("start_cluster_shape_configs"); ok {
		if tmpList := startClusterShapeConfigs.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "start_cluster_shape_configs", 0)
			tmp, err := s.mapToStartClusterShapeConfigs(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.StartClusterShapeConfigs = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	_, err := s.Client.StartBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_bds.BdsInstanceLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) StopBdsInstance() error {
	request := oci_bds.StopBdsInstanceRequest{}

	idTmp := s.D.Id()
	request.BdsInstanceId = &idTmp

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if isForceStopJobs, ok := s.D.GetOkExists("is_force_stop_jobs"); ok {
		tmp := isForceStopJobs.(bool)
		request.IsForceStopJobs = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	_, err := s.Client.StopBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_bds.BdsInstanceLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) AddKafka() error {

	request := oci_bds.AddKafkaRequest{}

	idTmp := s.D.Id()
	request.BdsInstanceId = &idTmp

	if nodes, ok := s.D.GetOkExists("kafka_broker_node"); ok {
		interfaces := nodes.([]interface{})
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "kafka_broker_node", stateDataIndex)

			if blockVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_size_in_gbs")); ok {
				if blockVolumeSizeInGBs != "" {
					tmp := blockVolumeSizeInGBs.(string)
					tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
					if err != nil {
						return err
					}
					request.BlockVolumeSizeInGBs = &tmpInt64
				}
			}

			if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
				tmp := shape.(string)
				request.Shape = &tmp
			}

			if shapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config")); ok {
				if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
					fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config"), 0)
					tmp, err := s.mapToShapeConfigDetails(fieldKeyFormatNextLevel)
					if err != nil {
						return err
					}
					request.ShapeConfig = &tmp
				}
			}

			if numberOfKafkaNodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "number_of_kafka_nodes")); ok {
				tmp := numberOfKafkaNodes.(int)
				request.NumberOfKafkaNodes = &tmp
			}
		}
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")
	response, err := s.Client.AddKafka(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) ExecuteBootstrapScript() error {
	request := oci_bds.ExecuteBootstrapScriptRequest{}

	idTmp := s.D.Id()
	request.BdsInstanceId = &idTmp

	if bootstrapScriptUrl, ok := s.D.GetOkExists("bootstrap_script_url"); ok {
		tmp := bootstrapScriptUrl.(string)
		request.BootstrapScriptUrl = &tmp
	}

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.ExecuteBootstrapScript(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) deleteShapeConfigIfMissingInInput(node_type string, node_map map[string]interface{}) {
	if _, ok := s.D.GetOkExists(node_type); ok {
		fieldKey := fmt.Sprintf("%s.%d.%s", node_type, 0, "shape_config")
		if nodeConfig, ok := s.D.GetOkExists(fieldKey); !ok || len(nodeConfig.([]interface{})) == 0 {
			delete(node_map, "shape_config")
		}
	}
}
func (s *BdsBdsInstanceResourceCrud) RemoveKafka() error {
	request := oci_bds.RemoveKafkaRequest{}

	idTmp := s.D.Id()
	request.BdsInstanceId = &idTmp

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.RemoveKafka(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) RemoveNode() error {
	request := oci_bds.RemoveNodeRequest{}

	idTmp := s.D.Id()
	request.BdsInstanceId = &idTmp

	if clusterAdminPassword, ok := s.D.GetOkExists("cluster_admin_password"); ok {
		tmp := clusterAdminPassword.(string)
		request.ClusterAdminPassword = &tmp
	}

	if isForceRemoveEnabled, ok := s.D.GetOkExists("is_force_remove_enabled"); ok {
		tmp := isForceRemoveEnabled.(bool)
		request.IsForceRemoveEnabled = &tmp
	}

	if nodeId, ok := s.D.GetOkExists("remove_node"); ok {
		tmp := nodeId.(string)
		request.NodeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.RemoveNode(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) mapToBdsClusterVersionSummary(fieldKeyFormat string) (oci_bds.BdsClusterVersionSummary, error) {
	result := oci_bds.BdsClusterVersionSummary{}

	if bdsVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bds_version")); ok {
		tmp := bdsVersion.(string)
		result.BdsVersion = &tmp
	}

	if odhVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "odh_version")); ok {
		tmp := odhVersion.(string)
		result.OdhVersion = &tmp
	}

	return result, nil
}

func (s *BdsBdsInstanceResourceCrud) mapToStartClusterShapeConfigs(fieldKeyFormat string) (oci_bds.StartClusterShapeConfigs, error) {
	request := oci_bds.StartClusterShapeConfigs{}

	if nodeTypeShapeConfigs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_type_shape_configs")); ok {
		interfaces := nodeTypeShapeConfigs.([]interface{})
		tmp := make([]oci_bds.NodeTypeShapeConfig, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%s.%d.%%s", "start_cluster_shape_configs.0", "node_type_shape_configs", stateDataIndex)
			converted, err := s.mapToNodeTypeShapeConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return request, err
			}
			tmp[i] = converted

			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "node_type_shape_configs")) {
				request.NodeTypeShapeConfigs = tmp
			}
		}
	}
	return request, nil
}
func (s *BdsBdsInstanceResourceCrud) mapToNodeTypeShapeConfig(fieldKeyFormat string) (oci_bds.NodeTypeShapeConfig, error) {
	request := oci_bds.NodeTypeShapeConfig{}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}
	if nodeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "node_type")); ok {
		tmp := oci_bds.NodeNodeTypeEnum(nodeType.(string))
		request.NodeType = tmp
	}

	return request, nil
}

func BdsClusterVersionSummaryToMap(obj *oci_bds.BdsClusterVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BdsVersion != nil {
		result["bds_version"] = string(*obj.BdsVersion)
	}

	if obj.OdhVersion != nil {
		result["odh_version"] = string(*obj.OdhVersion)
	}

	return result
}

func CloudSqlDetailsToMap(obj *oci_bds.CloudSqlDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BlockVolumeSizeInGBs != nil {
		result["block_volume_size_in_gbs"] = strconv.FormatInt(*obj.BlockVolumeSizeInGBs, 10)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.IsKerberosMappedToDatabaseUsers != nil {
		result["is_kerberos_mapped_to_database_users"] = bool(*obj.IsKerberosMappedToDatabaseUsers)
	}

	kerberosDetails := []interface{}{}
	for _, item := range obj.KerberosDetails {
		kerberosDetails = append(kerberosDetails, KerberosDetailsToMap(item))
	}
	result["kerberos_details"] = kerberosDetails

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	return result
}

func ClusterDetailsToMap(obj *oci_bds.ClusterDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AmbariUrl != nil {
		result["ambari_url"] = string(*obj.AmbariUrl)
	}

	if obj.BdCellVersion != nil {
		result["bd_cell_version"] = string(*obj.BdCellVersion)
	}

	if obj.BdaVersion != nil {
		result["bda_version"] = string(*obj.BdaVersion)
	}

	if obj.BdmVersion != nil {
		result["bdm_version"] = string(*obj.BdmVersion)
	}

	if obj.BdsVersion != nil {
		result["bds_version"] = string(*obj.BdsVersion)
	}

	if obj.BigDataManagerUrl != nil {
		result["big_data_manager_url"] = string(*obj.BigDataManagerUrl)
	}

	if obj.ClouderaManagerUrl != nil {
		result["cloudera_manager_url"] = string(*obj.ClouderaManagerUrl)
	}

	if obj.CsqlCellVersion != nil {
		result["csql_cell_version"] = string(*obj.CsqlCellVersion)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.HueServerUrl != nil {
		result["hue_server_url"] = string(*obj.HueServerUrl)
	}

	if obj.JupyterHubUrl != nil {
		result["jupyter_hub_url"] = string(*obj.JupyterHubUrl)
	}

	if obj.OdhVersion != nil {
		result["odh_version"] = string(*obj.OdhVersion)
	}

	if obj.OsVersion != nil {
		result["os_version"] = string(*obj.OsVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeRefreshed != nil {
		result["time_refreshed"] = obj.TimeRefreshed.String()
	}

	return result
}

func (s *BdsBdsInstanceResourceCrud) mapToCreateNodeDetails(fieldKeyFormat, nodeType string) (oci_bds.CreateNodeDetails, error) {
	result := oci_bds.CreateNodeDetails{}

	if blockVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volume_size_in_gbs")); ok {
		tmp := blockVolumeSizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert blockVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.BlockVolumeSizeInGBs = &tmpInt64
	}

	result.NodeType = oci_bds.NodeNodeTypeEnum(nodeType)

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if shapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config")); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config"), 0)
			tmp, err := s.mapToShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert shape_config, encountered error: %v", err)
			}
			result.ShapeConfig = &tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func BdsNodeToMap(obj oci_bds.Node) map[string]interface{} {
	result := map[string]interface{}{}

	attachedBlockVolumes := []interface{}{}
	for _, item := range obj.AttachedBlockVolumes {
		attachedBlockVolumes = append(attachedBlockVolumes, VolumeAttachmentDetailToMap(item))
	}

	result["attached_block_volumes"] = attachedBlockVolumes

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.ImageId != nil {
		result["image_id"] = string(*obj.ImageId)
	}

	if obj.InstanceId != nil {
		result["instance_id"] = string(*obj.InstanceId)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.IsRebootRequired != nil {
		result["is_reboot_required"] = bool(*obj.IsRebootRequired)
	}

	if obj.LocalDisksTotalSizeInGBs != nil {
		result["local_disks_total_size_in_gbs"] = float64(*obj.LocalDisksTotalSizeInGBs)
	}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = int(*obj.MemoryInGBs)
	}

	result["node_type"] = string(obj.NodeType)

	if obj.Nvmes != nil {
		result["nvmes"] = int(*obj.Nvmes)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = int(*obj.Ocpus)
	}

	if obj.OdhVersion != nil {
		result["odh_version"] = string(*obj.OdhVersion)
	}

	if obj.OsVersion != nil {
		result["os_version"] = string(*obj.OsVersion)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SshFingerprint != nil {
		result["ssh_fingerprint"] = string(*obj.SshFingerprint)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeMaintenanceRebootDue != nil {
		result["time_maintenance_reboot_due"] = obj.TimeMaintenanceRebootDue.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func KerberosDetailsToMap(obj oci_bds.KerberosDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeytabFile != nil {
		result["keytab_file"] = string(*obj.KeytabFile)
	}

	if obj.PrincipalName != nil {
		result["principal_name"] = string(*obj.PrincipalName)
	}

	return result
}

func (s *BdsBdsInstanceResourceCrud) mapToNetworkConfig(fieldKeyFormat string) (oci_bds.NetworkConfig, error) {
	result := oci_bds.NetworkConfig{}

	if cidrBlock, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cidr_block")); ok {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	if isNatGatewayRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_nat_gateway_required")); ok {
		tmp := isNatGatewayRequired.(bool)
		result.IsNatGatewayRequired = &tmp
	}

	return result, nil
}

func NetworkConfigToMap(obj *oci_bds.NetworkConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CidrBlock != nil {
		result["cidr_block"] = string(*obj.CidrBlock)
	}

	if obj.IsNatGatewayRequired != nil {
		result["is_nat_gateway_required"] = bool(*obj.IsNatGatewayRequired)
	}

	return result
}

func (s *BdsBdsInstanceResourceCrud) mapToShapeConfigDetails(fieldKeyFormat string) (oci_bds.ShapeConfigDetails, error) {
	result := oci_bds.ShapeConfigDetails{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := memoryInGBs.(int)
		result.MemoryInGBs = &tmp
	}

	if nvmes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nvmes")); ok {
		tmp := nvmes.(int)
		if tmp != 0 {
			result.Nvmes = &tmp
		}
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := ocpus.(int)
		result.Ocpus = &tmp
	}

	return result, nil
}

func VolumeAttachmentDetailToMap(obj oci_bds.VolumeAttachmentDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.VolumeAttachmentId != nil {
		result["volume_attachment_id"] = string(*obj.VolumeAttachmentId)
	}

	if obj.VolumeSizeInGBs != nil {
		result["volume_size_in_gbs"] = strconv.FormatInt(*obj.VolumeSizeInGBs, 10)
	}

	return result
}

func (s *BdsBdsInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_bds.ChangeBdsInstanceCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BdsInstanceId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.ChangeBdsInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) updateWorkerBlockStorage(id string, clusterAdminPassword interface{}, blockVolumeSizeInGBs int64, nodeType oci_bds.AddBlockStorageDetailsNodeTypeEnum) error {
	addBlockStorageRequest := oci_bds.AddBlockStorageRequest{}

	addBlockStorageRequest.BdsInstanceId = &id

	addBlockStorageRequest.NodeType = nodeType

	tmpClusterAdminPassword := clusterAdminPassword.(string)
	addBlockStorageRequest.ClusterAdminPassword = &tmpClusterAdminPassword

	addBlockStorageRequest.BlockVolumeSizeInGBs = &blockVolumeSizeInGBs

	addBlockStorageRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.AddBlockStorage(context.Background(), addBlockStorageRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) updateWorkerNode(id string, clusterAdminPassword interface{}, numberOfWorker int, nodeType oci_bds.AddWorkerNodesDetailsNodeTypeEnum, blockVolumeSizeInGBs *int64, shape *string, shapeConfig *oci_bds.ShapeConfigDetails) error {
	addWorkerNodesRequest := oci_bds.AddWorkerNodesRequest{}
	addWorkerNodesRequest.BdsInstanceId = &id

	addWorkerNodesRequest.NodeType = nodeType

	if shape != nil {
		addWorkerNodesRequest.Shape = shape
	}

	if shapeConfig != nil {
		addWorkerNodesRequest.ShapeConfig = shapeConfig
	}

	if blockVolumeSizeInGBs != nil {
		addWorkerNodesRequest.BlockVolumeSizeInGBs = blockVolumeSizeInGBs
	}

	clusterAdminPasswordTmp := clusterAdminPassword.(string)
	addWorkerNodesRequest.ClusterAdminPassword = &clusterAdminPasswordTmp

	addWorkerNodesRequest.NumberOfWorkerNodes = &numberOfWorker

	addWorkerNodesRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.AddWorkerNodes(context.Background(), addWorkerNodesRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) updateMasterNode(id string, clusterAdminPassword interface{}, numberOfMaster int, blockVolumeSizeInGBs *int64, shape *string, shapeConfig *oci_bds.ShapeConfigDetails) error {
	addMasterNodesRequest := oci_bds.AddMasterNodesRequest{}
	addMasterNodesRequest.BdsInstanceId = &id

	if shape != nil {
		addMasterNodesRequest.Shape = shape
	}

	if shapeConfig != nil {
		addMasterNodesRequest.ShapeConfig = shapeConfig
	}

	if blockVolumeSizeInGBs != nil {
		addMasterNodesRequest.BlockVolumeSizeInGBs = blockVolumeSizeInGBs
	}

	clusterAdminPasswordTmp := clusterAdminPassword.(string)
	addMasterNodesRequest.ClusterAdminPassword = &clusterAdminPasswordTmp

	addMasterNodesRequest.NumberOfMasterNodes = &numberOfMaster

	addMasterNodesRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.AddMasterNodes(context.Background(), addMasterNodesRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) updateUtilityNode(id string, clusterAdminPassword interface{}, numberOfUtility int, blockVolumeSizeInGBs *int64, shape *string, shapeConfig *oci_bds.ShapeConfigDetails) error {
	addUtilityNodesRequest := oci_bds.AddUtilityNodesRequest{}
	addUtilityNodesRequest.BdsInstanceId = &id

	if shape != nil {
		addUtilityNodesRequest.Shape = shape
	}

	if shapeConfig != nil {
		addUtilityNodesRequest.ShapeConfig = shapeConfig
	}

	if blockVolumeSizeInGBs != nil {
		addUtilityNodesRequest.BlockVolumeSizeInGBs = blockVolumeSizeInGBs
	}

	clusterAdminPasswordTmp := clusterAdminPassword.(string)
	addUtilityNodesRequest.ClusterAdminPassword = &clusterAdminPasswordTmp

	addUtilityNodesRequest.NumberOfUtilityNodes = &numberOfUtility

	addUtilityNodesRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.AddUtilityNodes(context.Background(), addUtilityNodesRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) addCloudSql(request oci_bds.AddCloudSqlRequest) error {
	response, err := s.Client.AddCloudSql(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BdsBdsInstanceResourceCrud) deleteCloudSql(request oci_bds.RemoveCloudSqlRequest) error {
	response, err := s.Client.RemoveCloudSql(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	return s.getBdsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func PopulateNodeTemplate(obj oci_bds.Node, nodeMap map[string]map[string]interface{}) {
	switch nodeType := string(obj.NodeType); nodeType {
	case "MASTER":
		if node, ok := nodeMap["MASTER"]; ok {
			node["number_of_nodes"] = node["number_of_nodes"].(int) + 1
		} else {
			nodeMap["MASTER"] = BdsNodeToTemplateMap(obj)
		}
	case "UTILITY":
		if node, ok := nodeMap["UTILITY"]; ok {
			node["number_of_nodes"] = node["number_of_nodes"].(int) + 1
		} else {
			nodeMap["UTILITY"] = BdsNodeToTemplateMap(obj)
		}
	case "WORKER":
		if node, ok := nodeMap["WORKER"]; ok {
			node["number_of_nodes"] = node["number_of_nodes"].(int) + 1
		} else {
			nodeMap["WORKER"] = BdsNodeToTemplateMap(obj)
		}
	case "COMPUTE_ONLY_WORKER":
		if node, ok := nodeMap["COMPUTE_ONLY_WORKER"]; ok {
			node["number_of_nodes"] = node["number_of_nodes"].(int) + 1
		} else {
			nodeMap["COMPUTE_ONLY_WORKER"] = BdsNodeToTemplateMap(obj)
		}
	case "EDGE":
		if node, ok := nodeMap["EDGE"]; ok {
			node["number_of_nodes"] = node["number_of_nodes"].(int) + 1
		} else {
			nodeMap["EDGE"] = BdsNodeToTemplateMap(obj)
		}
	}
}

func BdsNodeToTemplateMap(obj oci_bds.Node) map[string]interface{} {
	result := map[string]interface{}{}

	totalSize := int64(0)
	for _, item := range obj.AttachedBlockVolumes {
		if item.VolumeSizeInGBs != nil {
			totalSize += *item.VolumeSizeInGBs
		}
	}
	if totalSize != 0 {
		result["block_volume_size_in_gbs"] = strconv.FormatInt(totalSize, 10)
	}
	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	result["number_of_nodes"] = 1

	if obj.Ocpus != nil && obj.MemoryInGBs != nil {
		shapeConfigMap := map[string]interface{}{}
		shapeConfigMap["ocpus"] = int(*obj.Ocpus)
		shapeConfigMap["memory_in_gbs"] = int(*obj.MemoryInGBs)
		if result["shape"] == "VM.DenseIO.E4.Flex" || result["shape"] == "VM.DenseIO.E5.Flex" || result["shape"] == "VM.DenseIO.Generic" {
			shapeConfigMap["nvmes"] = int(*obj.Nvmes)
		}

		result["shape_config"] = []interface{}{shapeConfigMap}
	}
	return result
}

func ShapeChangeDiffSuppressFunction(nodeType string, d *schema.ResourceData) bool {
	var ignoreMasterShape, ignoreUtilShape, ignoreWorkerShape, ignoreComputeWorkerShape, ignoreEdgeShape, ignoreKafkaBrokerShape = false, false, false, false, false, false
	var addNode bool
	if changeExistingNodesTrigger, ok := d.GetOkExists("ignore_existing_nodes_shape"); ok {
		interfaces := changeExistingNodesTrigger.([]interface{})
		tmp := make([]string, len(interfaces))
		// Add now node types when they are released
		for i := range interfaces {
			tmp[i] = strings.TrimSpace(strings.ToLower(interfaces[i].(string)))
			if tmp[i] == "master" {
				ignoreMasterShape = true
			}
			if tmp[i] == "utility" {
				ignoreUtilShape = true
			}
			if tmp[i] == "worker" {
				ignoreWorkerShape = true
			}
			if tmp[i] == "compute_only_worker" {
				ignoreComputeWorkerShape = true
			}
			if tmp[i] == "edge" {
				ignoreEdgeShape = true
			}
			if tmp[i] == "kafka_broker" {
				ignoreKafkaBrokerShape = true
			}
		}
	} else {
		return false
	}
	// Add now node types when they are released
	if nodeType == "master" && ignoreMasterShape == true {
		addNode = d.HasChange("master_node.0.number_of_nodes")
	} else if nodeType == "util" && ignoreUtilShape == true {
		addNode = d.HasChange("util_node.0.number_of_nodes")
	} else if nodeType == "worker" && ignoreWorkerShape == true {
		addNode = d.HasChange("worker_node.0.number_of_nodes")
	} else if nodeType == "compute_only_worker" && ignoreComputeWorkerShape == true {
		addNode = d.HasChange("compute_only_worker_node.0.number_of_nodes")
	} else if nodeType == "edge" && ignoreEdgeShape == true {
		addNode = d.HasChange("edge_node.0.number_of_nodes")
	} else if nodeType == "kafka_broker" && ignoreKafkaBrokerShape == true {
		addNode = d.HasChange("kafka_broker_node.0.number_of_kafka_nodes")
	} else {
		addNode = true
	}
	if !addNode {
		return true
	}
	return false
}
