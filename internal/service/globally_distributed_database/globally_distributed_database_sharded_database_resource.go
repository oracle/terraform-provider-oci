// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0 frf

package globally_distributed_database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GloballyDistributedDatabaseShardedDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createGloballyDistributedDatabaseShardedDatabase,
		Read:   readGloballyDistributedDatabaseShardedDatabase,
		Update: updateGloballyDistributedDatabaseShardedDatabase,
		Delete: deleteGloballyDistributedDatabaseShardedDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"catalog_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_password": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"cloud_autonomous_vm_cluster_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"compute_count": {
							Type:     schema.TypeFloat,
							Required: true,
							ForceNew: true,
						},
						"data_storage_size_in_gbs": {
							Type:     schema.TypeFloat,
							Required: true,
							ForceNew: true,
						},
						"is_auto_scaling_enabled": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"encryption_key_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"kms_key_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"vault_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"kms_key_version_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"peer_cloud_autonomous_vm_cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"container_database_parent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shard_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supporting_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_ssl_certificate_expires": {
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
			"character_set": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_deployment_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"DEDICATED",
				}, true),
			},
			"db_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"listener_port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"listener_port_tls": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ncharacter_set": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ons_port_local": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ons_port_remote": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shard_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_password": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"cloud_autonomous_vm_cluster_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"compute_count": {
							Type:     schema.TypeFloat,
							Required: true,
							ForceNew: true,
						},
						"data_storage_size_in_gbs": {
							Type:     schema.TypeFloat,
							Required: true,
							ForceNew: true,
						},
						"is_auto_scaling_enabled": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"encryption_key_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"kms_key_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"vault_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"kms_key_version_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"peer_cloud_autonomous_vm_cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"shard_space": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"container_database_parent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shard_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supporting_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_ssl_certificate_expires": {
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
			"sharded_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sharding_method": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"chunks": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cluster_certificate_common_name": {
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
								"INSERT",
								"MERGE",
								"REMOVE",
							}, true),
						},
						"selection": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							//Type:             schema.TypeMap,
							Type:     schema.TypeString,
							Required: true,
							//DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
			},
			"configure_gsms_trigger": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"configure_sharding_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"download_gsm_certificate_signing_request_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"generate_gsm_certificate_signing_request_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"generate_wallet_trigger": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"get_connection_string_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			/*"fetched_all_connection_strings": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},*/
			"start_database_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stop_database_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upload_signed_certificate_and_generate_wallet_trigger": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"validate_network_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"all_connection_strings": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
			"gsms": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compute_count": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"data_storage_size_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supporting_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_ssl_certificate_expires": {
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
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_state": {
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
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createGloballyDistributedDatabaseShardedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabaseShardedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("configure_gsms_trigger"); ok {
		err := sync.ConfigureShardedDatabaseGsms()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("configure_sharding_trigger"); ok {
		err := sync.ConfigureSharding()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("download_gsm_certificate_signing_request_trigger"); ok {
		err := sync.DownloadGsmCertificateSigningRequest()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("generate_gsm_certificate_signing_request_trigger"); ok {
		err := sync.GenerateGsmCertificateSigningRequest()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("generate_wallet_trigger"); ok {
		err := sync.GenerateWallet()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("get_connection_string_trigger"); ok {
		err := sync.FetchConnectionString()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("start_database_trigger"); ok {
		err := sync.StartShardedDatabase()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("stop_database_trigger"); ok {
		err := sync.StopShardedDatabase()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("upload_signed_certificate_and_generate_wallet_trigger"); ok {
		err := sync.UploadSignedCertificateAndGenerateWallet()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("validate_network_trigger"); ok {
		err := sync.ValidateNetwork()
		if err != nil {
			return err
		}
	}
	return nil

}

func readGloballyDistributedDatabaseShardedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabaseShardedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	return tfresource.ReadResource(sync)
}

func updateGloballyDistributedDatabaseShardedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabaseShardedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()

	if _, ok := sync.D.GetOkExists("configure_gsms_trigger"); ok && sync.D.HasChange("configure_gsms_trigger") {
		oldRaw, newRaw := sync.D.GetChange("configure_gsms_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConfigureShardedDatabaseGsms()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("configure_gsms_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("configure_sharding_trigger"); ok && sync.D.HasChange("configure_sharding_trigger") {
		oldRaw, newRaw := sync.D.GetChange("configure_sharding_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ConfigureSharding()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("configure_sharding_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("download_gsm_certificate_signing_request_trigger"); ok && sync.D.HasChange("download_gsm_certificate_signing_request_trigger") {
		oldRaw, newRaw := sync.D.GetChange("download_gsm_certificate_signing_request_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.DownloadGsmCertificateSigningRequest()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("download_gsm_certificate_signing_request_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("generate_gsm_certificate_signing_request_trigger"); ok && sync.D.HasChange("generate_gsm_certificate_signing_request_trigger") {
		oldRaw, newRaw := sync.D.GetChange("generate_gsm_certificate_signing_request_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.GenerateGsmCertificateSigningRequest()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("generate_gsm_certificate_signing_request_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("generate_wallet_trigger"); ok && sync.D.HasChange("generate_wallet_trigger") {
		oldRaw, newRaw := sync.D.GetChange("generate_wallet_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.GenerateWallet()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("generate_wallet_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("get_connection_string_trigger"); ok && sync.D.HasChange("get_connection_string_trigger") {
		oldRaw, newRaw := sync.D.GetChange("get_connection_string_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.FetchConnectionString()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("get_connection_string_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("start_database_trigger"); ok && sync.D.HasChange("start_database_trigger") {
		oldRaw, newRaw := sync.D.GetChange("start_database_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.StartShardedDatabase()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("start_database_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("stop_database_trigger"); ok && sync.D.HasChange("stop_database_trigger") {
		oldRaw, newRaw := sync.D.GetChange("stop_database_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.StopShardedDatabase()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("stop_database_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("upload_signed_certificate_and_generate_wallet_trigger"); ok && sync.D.HasChange("upload_signed_certificate_and_generate_wallet_trigger") {
		oldRaw, newRaw := sync.D.GetChange("upload_signed_certificate_and_generate_wallet_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.UploadSignedCertificateAndGenerateWallet()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("upload_signed_certificate_and_generate_wallet_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("validate_network_trigger"); ok && sync.D.HasChange("validate_network_trigger") {
		oldRaw, newRaw := sync.D.GetChange("validate_network_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ValidateNetwork()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("validate_network_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteGloballyDistributedDatabaseShardedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &GloballyDistributedDatabaseShardedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ShardedDatabaseServiceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GloballyDistributedDatabaseShardedDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_globally_distributed_database.ShardedDatabaseServiceClient
	Res                    *oci_globally_distributed_database.ShardedDatabase
	PatchResponse          *oci_globally_distributed_database.ShardedDatabase
	DisableNotFoundRetries bool
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) ID() string {
	shardedDatabase := *s.Res
	return *shardedDatabase.GetId()
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateCreating),
	}
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateActive),
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateNeedsAttention),
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateInactive),
	}
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateDeleting),
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateUnavailable),
	}
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateDeleted),
	}
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) Create() error {
	request := oci_globally_distributed_database.CreateShardedDatabaseRequest{}
	err := s.populateTopLevelPolymorphicCreateShardedDatabaseRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.CreateShardedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	//identifier = response.id
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	err = s.getShardedDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database"), oci_globally_distributed_database.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	err = s.Patch()
	if err != nil {
		log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
		return err
	}
	return nil
}
func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) Patch() error {
	request := oci_globally_distributed_database.PatchShardedDatabaseRequest{}

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
		interfaces := patchOperations.([]interface{})
		tmp := make([]oci_globally_distributed_database.PatchInstruction, len(interfaces))
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
		tmpId := s.D.Id()
		if tmpId != "" {
			request.ShardedDatabaseId = &tmpId
		}

		if shardedDatabaseId, ok := s.D.GetOkExists("id"); ok {
			tmp := shardedDatabaseId.(string)
			request.ShardedDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")
		response, err := s.Client.PatchShardedDatabase(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		return s.getShardedDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database"), oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
	}

	return nil

}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) getShardedDatabaseFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_globally_distributed_database.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	shardedDatabaseId, err := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*shardedDatabaseId)

	return s.Get()
}

func shardedDatabaseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "globally_distributed_database", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_globally_distributed_database.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func shardedDatabaseWaitForWorkRequest(wId *string, entityType string, action oci_globally_distributed_database.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_globally_distributed_database.ShardedDatabaseServiceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "globally_distributed_database")
	retryPolicy.ShouldRetryOperation = shardedDatabaseWorkRequestShouldRetryFunc(timeout)

	response := oci_globally_distributed_database.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_globally_distributed_database.OperationStatusInProgress),
			string(oci_globally_distributed_database.OperationStatusAccepted),
			string(oci_globally_distributed_database.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_globally_distributed_database.OperationStatusSucceeded),
			string(oci_globally_distributed_database.OperationStatusFailed),
			string(oci_globally_distributed_database.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_globally_distributed_database.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_globally_distributed_database.OperationStatusFailed || response.Status == oci_globally_distributed_database.OperationStatusCanceled {
		return nil, getErrorFromGloballyDistributedDatabaseShardedDatabaseWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGloballyDistributedDatabaseShardedDatabaseWorkRequest(client *oci_globally_distributed_database.ShardedDatabaseServiceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_globally_distributed_database.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_globally_distributed_database.ListWorkRequestErrorsRequest{
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

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) Get() error {
	request := oci_globally_distributed_database.GetShardedDatabaseRequest{}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp := metadata.(string)
		request.Metadata = &tmp
	}

	tmp := s.D.Id()
	request.ShardedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.GetShardedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ShardedDatabase
	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_globally_distributed_database.UpdateShardedDatabaseRequest{}

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

	tmp := s.D.Id()
	request.ShardedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.UpdateShardedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ShardedDatabase

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {
		if tmpList := patchOperations.([]interface{}); len(tmpList) > 0 {
			err = s.Patch()
			if err != nil {
				log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
				return err
			}
		}
	}
	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) Delete() error {
	request := oci_globally_distributed_database.DeleteShardedDatabaseRequest{}

	tmp := s.D.Id()
	request.ShardedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.DeleteShardedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		oci_globally_distributed_database.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_globally_distributed_database.DedicatedShardedDatabase:
		s.D.Set("db_deployment_type", "DEDICATED")

		catalogDetails := []interface{}{}
		shardDetails := []interface{}{}
		tmpId := s.D.Id()

		if tmpId != "" {
			tmpReq := oci_globally_distributed_database.GetShardedDatabaseRequest{
				ShardedDatabaseId: &tmpId,
			}
			tmpResp, _ := s.Client.GetShardedDatabase(context.Background(), tmpReq)
			tmpShardedDb := tmpResp.ShardedDatabase
			dedicateShardedDb := tmpShardedDb.(oci_globally_distributed_database.DedicatedShardedDatabase)
			cDAdminPassword := ""

			for _, item := range dedicateShardedDb.CatalogDetails {
				if createCatalogDetails, ok := s.D.GetOkExists("catalog_details"); ok {
					interfaces := createCatalogDetails.([]interface{})
					tmp := make([]oci_globally_distributed_database.CreateDedicatedCatalogDetail, len(interfaces))
					for i := range interfaces {
						stateDataIndex := i
						fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "catalog_details", stateDataIndex)
						converted, err := s.mapToCreateDedicatedCatalogDetail(fieldKeyFormat)
						if err != nil {
							return err
						}
						tmp[i] = converted
						if *tmp[i].CloudAutonomousVmClusterId == *item.CloudAutonomousVmClusterId {
							cDAdminPassword = *tmp[i].AdminPassword
						}
					}
				}
				catalogDetails = append(catalogDetails, DedicatedCatalogDetailsToMapForResource(item, cDAdminPassword))
			}

			sDAdminPassword := ""
			for _, item := range dedicateShardedDb.ShardDetails {
				if createShardDetails, ok := s.D.GetOkExists("shard_details"); ok {
					interfaces := createShardDetails.([]interface{})
					tmp := make([]oci_globally_distributed_database.CreateDedicatedShardDetail, len(interfaces))
					for i := range interfaces {
						stateDataIndex := i
						fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shard_details", stateDataIndex)
						converted, err := s.mapToCreateDedicatedShardDetail(fieldKeyFormat)
						if err != nil {
							return err
						}
						tmp[i] = converted
						if *tmp[i].CloudAutonomousVmClusterId == *item.CloudAutonomousVmClusterId {
							sDAdminPassword = *tmp[i].AdminPassword
						}
					}
				}
				shardDetails = append(shardDetails, DedicatedShardDetailsToMapForResource(item, sDAdminPassword))
			}
		}

		s.D.Set("catalog_details", catalogDetails)

		s.D.Set("shard_details", shardDetails)

		if v.CharacterSet != nil {
			s.D.Set("character_set", *v.CharacterSet)
		}

		if v.Chunks != nil {
			s.D.Set("chunks", *v.Chunks)
		}

		if v.ClusterCertificateCommonName != nil {
			s.D.Set("cluster_certificate_common_name", *v.ClusterCertificateCommonName)
		}

		if v.ConnectionStrings != nil {
			s.D.Set("connection_strings", []interface{}{ConnectionStringToMap(v.ConnectionStrings)})
		} else {
			s.D.Set("connection_strings", nil)
		}

		if v.DbVersion != nil {
			s.D.Set("db_version", *v.DbVersion)
		}

		s.D.Set("db_workload", v.DbWorkload)

		gsms := []interface{}{}
		for _, item := range v.Gsms {
			gsms = append(gsms, GsmDetailsToMap(item))
		}
		s.D.Set("gsms", gsms)

		if v.ListenerPort != nil {
			s.D.Set("listener_port", *v.ListenerPort)
		}

		if v.ListenerPortTls != nil {
			s.D.Set("listener_port_tls", *v.ListenerPortTls)
		}

		if v.NcharacterSet != nil {
			s.D.Set("ncharacter_set", *v.NcharacterSet)
		}

		if v.OnsPortLocal != nil {
			s.D.Set("ons_port_local", *v.OnsPortLocal)
		}

		if v.OnsPortRemote != nil {
			s.D.Set("ons_port_remote", *v.OnsPortRemote)
		}

		if v.Prefix != nil {
			s.D.Set("prefix", *v.Prefix)
		}

		if v.PrivateEndpoint != nil {
			s.D.Set("private_endpoint", *v.PrivateEndpoint)
		}

		s.D.Set("sharding_method", v.ShardingMethod)

		if v.TimeZone != nil {
			s.D.Set("time_zone", *v.TimeZone)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		/*if v.Id != nil {
			s.D.Set("id", *v.Id)
		}*/

		if v.LifecycleStateDetails != nil {
			s.D.Set("lifecycle_state_details", *v.LifecycleStateDetails)
		}

		s.D.Set("lifecycle_state", v.LifecycleState)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'db_deployment_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) ConfigureShardedDatabaseGsms() error {
	request := oci_globally_distributed_database.ConfigureShardedDatabaseGsmsRequest{}

	if isLatestGsmImage, ok := s.D.GetOkExists("configure_gsms_trigger_is_latest_gsm_image"); ok {
		tmp := isLatestGsmImage.(bool)
		request.IsLatestGsmImage = &tmp
	}

	if oldGsmNames, ok := s.D.GetOkExists("configure_gsms_trigger_old_gsm_names"); ok {
		interfaces := oldGsmNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("configure_gsms_trigger_old_gsm_names") {
			request.OldGsmNames = tmp
		}
	}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.ConfigureShardedDatabaseGsms(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("configure_gsms_trigger")
	s.D.Set("configure_gsms_trigger", val)

	workId := response.OpcWorkRequestId

	_, configureShardedDatabaseGsmsWorkRequestErr := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	//s.Res = &response.ShardedDatabase
	return configureShardedDatabaseGsmsWorkRequestErr
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) ConfigureSharding() error {
	request := oci_globally_distributed_database.ConfigureShardingRequest{}

	if isRebalanceRequired, ok := s.D.GetOkExists("is_rebalance_required"); ok {
		tmp := isRebalanceRequired.(bool)
		request.IsRebalanceRequired = &tmp
	}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.ConfigureSharding(context.Background(), request)

	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("configure_sharding_trigger")
	s.D.Set("configure_sharding_trigger", val)

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, configureShardingWorkRequestErr := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	//s.Res = &response.ShardedDatabase
	return configureShardingWorkRequestErr
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) DownloadGsmCertificateSigningRequest() error {
	request := oci_globally_distributed_database.DownloadGsmCertificateSigningRequestRequest{}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	_, err := s.Client.DownloadGsmCertificateSigningRequest(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("download_gsm_certificate_signing_request_trigger")
	s.D.Set("download_gsm_certificate_signing_request_trigger", val)

	//s.Res = &response.ShardedDatabase
	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) GenerateGsmCertificateSigningRequest() error {
	request := oci_globally_distributed_database.GenerateGsmCertificateSigningRequestRequest{}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.GenerateGsmCertificateSigningRequest(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("generate_gsm_certificate_signing_request_trigger")
	s.D.Set("generate_gsm_certificate_signing_request_trigger", val)

	workId := response.OpcWorkRequestId

	_, generateGsmCertificateSigningRequesWorkRequestErr := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	//s.Res = &response.ShardedDatabase
	return generateGsmCertificateSigningRequesWorkRequestErr
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) GenerateWallet() error {
	request := oci_globally_distributed_database.GenerateWalletRequest{}

	/*if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}*/

	if password, ok := s.D.GetOkExists("generate_wallet_password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	_, err := s.Client.GenerateWallet(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("generate_wallet_trigger")
	s.D.Set("generate_wallet_trigger", val)

	//s.Res = &response.ShardedDatabase
	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) FetchConnectionString() error {
	request := oci_globally_distributed_database.FetchConnectionStringRequest{}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.FetchConnectionString(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("get_connection_string_trigger")
	s.D.Set("get_connection_string_trigger", val)

	s.D.Set("connection_strings", response.AllConnectionStrings)
	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) StartShardedDatabase() error {
	request := oci_globally_distributed_database.StartShardedDatabaseRequest{}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.StartShardedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("start_database_trigger")
	s.D.Set("start_database_trigger", val)

	workId := response.OpcWorkRequestId

	_, startShardedDatabaseWorkRequestErr := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	//s.Res = &response.ShardedDatabase
	return startShardedDatabaseWorkRequestErr
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) StopShardedDatabase() error {
	request := oci_globally_distributed_database.StopShardedDatabaseRequest{}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.StopShardedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("stop_database_trigger")
	s.D.Set("stop_database_trigger", val)

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, stopShardedDatabaseWorkRequestErr := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	//s.Res = &response.ShardedDatabase
	return stopShardedDatabaseWorkRequestErr
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) UploadSignedCertificateAndGenerateWallet() error {
	request := oci_globally_distributed_database.UploadSignedCertificateAndGenerateWalletRequest{}

	if caSignedCertificate, ok := s.D.GetOkExists("ca_signed_certificate"); ok {
		tmp := caSignedCertificate.(string)
		request.CaSignedCertificate = &tmp
	}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.UploadSignedCertificateAndGenerateWallet(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("upload_signed_certificate_and_generate_wallet_trigger")
	s.D.Set("upload_signed_certificate_and_generate_wallet_trigger", val)

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, uploadSignedCertificateAndGenerateWalletWorkRequestErr := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	//s.Res = &response.ShardedDatabase
	return uploadSignedCertificateAndGenerateWalletWorkRequestErr
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) ValidateNetwork() error {
	request := oci_globally_distributed_database.ValidateNetworkRequest{}

	if isPrimary, ok := s.D.GetOkExists("is_primary"); ok {
		tmp := isPrimary.(bool)
		request.IsPrimary = &tmp
	}

	if isSurrogate, ok := s.D.GetOkExists("is_surrogate"); ok {
		tmp := isSurrogate.(bool)
		request.IsSurrogate = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	idTmp := s.D.Id()
	request.ShardedDatabaseId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.ValidateNetwork(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("validate_network_trigger")
	s.D.Set("validate_network_trigger", val)

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, validateNetworkWorkRequestErr := shardedDatabaseWaitForWorkRequest(workId, "shardeddatabase",
		oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)

	//s.Res = &response.ShardedDatabase
	return validateNetworkWorkRequestErr

}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToConnectionString(fieldKeyFormat string) (oci_globally_distributed_database.ConnectionString, error) {
	result := oci_globally_distributed_database.ConnectionString{}

	if allConnectionStrings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "all_connection_strings")); ok {
		result.AllConnectionStrings = tfresource.ObjectMapToStringMap(allConnectionStrings.(map[string]interface{}))
	}

	return result, nil
}

func ConnectionStringToMap(obj *oci_globally_distributed_database.ConnectionString) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	return result
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToCreateDedicatedCatalogDetail(fieldKeyFormat string) (oci_globally_distributed_database.CreateDedicatedCatalogDetail, error) {
	result := oci_globally_distributed_database.CreateDedicatedCatalogDetail{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_autonomous_vm_cluster_id")); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		result.CloudAutonomousVmClusterId = &tmp
	}

	if computeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_count")); ok {
		//tmp := computeCount.(float32)
		tmp := float32(computeCount.(float64))
		result.ComputeCount = &tmp
	}

	if dataStorageSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gbs")); ok {
		tmp := dataStorageSizeInGbs.(float64)
		result.DataStorageSizeInGbs = &tmp
	}

	if encryptionKeyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_key_details")); ok {
		if tmpList := encryptionKeyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "encryption_key_details"), 0)
			tmp, err := s.mapToDedicatedShardOrCatalogEncryptionKeyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert encryption_key_details, encountered error: %v", err)
			}
			result.EncryptionKeyDetails = &tmp
		}
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_scaling_enabled")); ok {
		tmp := isAutoScalingEnabled.(bool)
		result.IsAutoScalingEnabled = &tmp
	}

	if peerCloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_cloud_autonomous_vm_cluster_id")); ok {
		tmp := peerCloudAutonomousVmClusterId.(string)
		result.PeerCloudAutonomousVmClusterId = &tmp
	}

	return result, nil
}

func CreateDedicatedCatalogDetailToMap(obj oci_globally_distributed_database.CreateDedicatedCatalogDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminPassword != nil {
		result["admin_password"] = string(*obj.AdminPassword)
	}

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = string(*obj.CloudAutonomousVmClusterId)
	}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	if obj.EncryptionKeyDetails != nil {
		result["encryption_key_details"] = []interface{}{DedicatedShardOrCatalogEncryptionKeyDetailsToMap(obj.EncryptionKeyDetails)}
	}

	if obj.IsAutoScalingEnabled != nil {
		result["is_auto_scaling_enabled"] = bool(*obj.IsAutoScalingEnabled)
	}

	if obj.PeerCloudAutonomousVmClusterId != nil {
		result["peer_cloud_autonomous_vm_cluster_id"] = string(*obj.PeerCloudAutonomousVmClusterId)
	}

	return result
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToCreateDedicatedShardDetail(fieldKeyFormat string) (oci_globally_distributed_database.CreateDedicatedShardDetail, error) {
	result := oci_globally_distributed_database.CreateDedicatedShardDetail{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_autonomous_vm_cluster_id")); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		result.CloudAutonomousVmClusterId = &tmp
	}

	if computeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_count")); ok {
		//tmp := computeCount.(float32)
		tmp := float32(computeCount.(float64))
		result.ComputeCount = &tmp
	}

	if dataStorageSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gbs")); ok {
		tmp := dataStorageSizeInGbs.(float64)
		result.DataStorageSizeInGbs = &tmp
	}

	if encryptionKeyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_key_details")); ok {
		if tmpList := encryptionKeyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "encryption_key_details"), 0)
			tmp, err := s.mapToDedicatedShardOrCatalogEncryptionKeyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert encryption_key_details, encountered error: %v", err)
			}
			result.EncryptionKeyDetails = &tmp
		}
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_scaling_enabled")); ok {
		tmp := isAutoScalingEnabled.(bool)
		result.IsAutoScalingEnabled = &tmp
	}

	if peerCloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_cloud_autonomous_vm_cluster_id")); ok {
		tmp := peerCloudAutonomousVmClusterId.(string)
		result.PeerCloudAutonomousVmClusterId = &tmp
	}

	if shardSpace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_space")); ok {
		tmp := shardSpace.(string)
		result.ShardSpace = &tmp
	}

	return result, nil
}

func CreateDedicatedShardDetailToMap(obj oci_globally_distributed_database.CreateDedicatedShardDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminPassword != nil {
		result["admin_password"] = string(*obj.AdminPassword)
	}

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = string(*obj.CloudAutonomousVmClusterId)
	}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	if obj.EncryptionKeyDetails != nil {
		result["encryption_key_details"] = []interface{}{DedicatedShardOrCatalogEncryptionKeyDetailsToMap(obj.EncryptionKeyDetails)}
	}

	if obj.IsAutoScalingEnabled != nil {
		result["is_auto_scaling_enabled"] = bool(*obj.IsAutoScalingEnabled)
	}

	if obj.PeerCloudAutonomousVmClusterId != nil {
		result["peer_cloud_autonomous_vm_cluster_id"] = string(*obj.PeerCloudAutonomousVmClusterId)
	}

	if obj.ShardSpace != nil {
		result["shard_space"] = string(*obj.ShardSpace)
	}

	return result
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToDedicatedCatalogDetails(fieldKeyFormat string) (oci_globally_distributed_database.DedicatedCatalogDetails, error) {
	result := oci_globally_distributed_database.DedicatedCatalogDetails{}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_autonomous_vm_cluster_id")); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		result.CloudAutonomousVmClusterId = &tmp
	}

	if computeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_count")); ok {
		//tmp := computeCount.(float32)
		tmp := float32(computeCount.(float64))
		result.ComputeCount = &tmp
	}

	if containerDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_database_id")); ok {
		tmp := containerDatabaseId.(string)
		result.ContainerDatabaseId = &tmp
	}

	if containerDatabaseParentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_database_parent_id")); ok {
		tmp := containerDatabaseParentId.(string)
		result.ContainerDatabaseParentId = &tmp
	}

	if dataStorageSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gbs")); ok {
		tmp := dataStorageSizeInGbs.(float64)
		result.DataStorageSizeInGbs = &tmp
	}

	if encryptionKeyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_key_details")); ok {
		if tmpList := encryptionKeyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "encryption_key_details"), 0)
			tmp, err := s.mapToDedicatedShardOrCatalogEncryptionKeyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert encryption_key_details, encountered error: %v", err)
			}
			result.EncryptionKeyDetails = &tmp
		}
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_scaling_enabled")); ok {
		tmp := isAutoScalingEnabled.(bool)
		result.IsAutoScalingEnabled = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		result.Metadata = metadata.(map[string]interface{})
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if peerCloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_cloud_autonomous_vm_cluster_id")); ok {
		tmp := peerCloudAutonomousVmClusterId.(string)
		result.PeerCloudAutonomousVmClusterId = &tmp
	}

	if shardGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_group")); ok {
		tmp := shardGroup.(string)
		result.ShardGroup = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_globally_distributed_database.DedicatedCatalogDetailsStatusEnum(status.(string))
	}

	if supportingResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "supporting_resource_id")); ok {
		tmp := supportingResourceId.(string)
		result.SupportingResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeSslCertificateExpires, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_ssl_certificate_expires")); ok {
		tmp, err := time.Parse(time.RFC3339, timeSslCertificateExpires.(string))
		if err != nil {
			return result, err
		}
		result.TimeSslCertificateExpires = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return result, err
		}
		result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func DedicatedCatalogDetailsToMap(obj oci_globally_distributed_database.DedicatedCatalogDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = string(*obj.CloudAutonomousVmClusterId)
	}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.ContainerDatabaseId != nil {
		result["container_database_id"] = string(*obj.ContainerDatabaseId)
	}

	if obj.ContainerDatabaseParentId != nil {
		result["container_database_parent_id"] = string(*obj.ContainerDatabaseParentId)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	if obj.EncryptionKeyDetails != nil {
		result["encryption_key_details"] = []interface{}{DedicatedShardOrCatalogEncryptionKeyDetailsToMap(obj.EncryptionKeyDetails)}
	}

	if obj.IsAutoScalingEnabled != nil {
		result["is_auto_scaling_enabled"] = bool(*obj.IsAutoScalingEnabled)
	}

	result["metadata"] = obj.Metadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PeerCloudAutonomousVmClusterId != nil {
		result["peer_cloud_autonomous_vm_cluster_id"] = string(*obj.PeerCloudAutonomousVmClusterId)
	}

	if obj.ShardGroup != nil {
		result["shard_group"] = string(*obj.ShardGroup)
	}

	result["status"] = string(obj.Status)

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeSslCertificateExpires != nil {
		result["time_ssl_certificate_expires"] = obj.TimeSslCertificateExpires.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	return result
}

func DedicatedCatalogDetailsToMapForResource(obj oci_globally_distributed_database.DedicatedCatalogDetails, adminPassword string) map[string]interface{} {
	result := map[string]interface{}{}

	result["admin_password"] = adminPassword

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = string(*obj.CloudAutonomousVmClusterId)
	}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.ContainerDatabaseId != nil {
		result["container_database_id"] = string(*obj.ContainerDatabaseId)
	}

	if obj.ContainerDatabaseParentId != nil {
		result["container_database_parent_id"] = string(*obj.ContainerDatabaseParentId)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	if obj.EncryptionKeyDetails != nil {
		result["encryption_key_details"] = []interface{}{DedicatedShardOrCatalogEncryptionKeyDetailsToMap(obj.EncryptionKeyDetails)}
	}

	if obj.IsAutoScalingEnabled != nil {
		result["is_auto_scaling_enabled"] = bool(*obj.IsAutoScalingEnabled)
	}

	result["metadata"] = obj.Metadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PeerCloudAutonomousVmClusterId != nil {
		result["peer_cloud_autonomous_vm_cluster_id"] = string(*obj.PeerCloudAutonomousVmClusterId)
	}

	if obj.ShardGroup != nil {
		result["shard_group"] = string(*obj.ShardGroup)
	}

	result["status"] = string(obj.Status)

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeSslCertificateExpires != nil {
		result["time_ssl_certificate_expires"] = obj.TimeSslCertificateExpires.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	return result
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToDedicatedShardDetails(fieldKeyFormat string) (oci_globally_distributed_database.DedicatedShardDetails, error) {
	result := oci_globally_distributed_database.DedicatedShardDetails{}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_autonomous_vm_cluster_id")); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		result.CloudAutonomousVmClusterId = &tmp
	}

	if computeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_count")); ok {
		//tmp := computeCount.(float32)
		tmp := float32(computeCount.(float64))
		result.ComputeCount = &tmp
	}

	if containerDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_database_id")); ok {
		tmp := containerDatabaseId.(string)
		result.ContainerDatabaseId = &tmp
	}

	if containerDatabaseParentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_database_parent_id")); ok {
		tmp := containerDatabaseParentId.(string)
		result.ContainerDatabaseParentId = &tmp
	}

	if dataStorageSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gbs")); ok {
		tmp := dataStorageSizeInGbs.(float64)
		result.DataStorageSizeInGbs = &tmp
	}

	if encryptionKeyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_key_details")); ok {
		if tmpList := encryptionKeyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "encryption_key_details"), 0)
			tmp, err := s.mapToDedicatedShardOrCatalogEncryptionKeyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert encryption_key_details, encountered error: %v", err)
			}
			result.EncryptionKeyDetails = &tmp
		}
	}

	if isAutoScalingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_scaling_enabled")); ok {
		tmp := isAutoScalingEnabled.(bool)
		result.IsAutoScalingEnabled = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		result.Metadata = metadata.(map[string]interface{})
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if peerCloudAutonomousVmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "peer_cloud_autonomous_vm_cluster_id")); ok {
		tmp := peerCloudAutonomousVmClusterId.(string)
		result.PeerCloudAutonomousVmClusterId = &tmp
	}

	if shardGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_group")); ok {
		tmp := shardGroup.(string)
		result.ShardGroup = &tmp
	}

	if shardSpace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shard_space")); ok {
		tmp := shardSpace.(string)
		result.ShardSpace = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_globally_distributed_database.DedicatedShardDetailsStatusEnum(status.(string))
	}

	if supportingResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "supporting_resource_id")); ok {
		tmp := supportingResourceId.(string)
		result.SupportingResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeSslCertificateExpires, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_ssl_certificate_expires")); ok {
		tmp, err := time.Parse(time.RFC3339, timeSslCertificateExpires.(string))
		if err != nil {
			return result, err
		}
		result.TimeSslCertificateExpires = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return result, err
		}
		result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func DedicatedShardDetailsToMap(obj oci_globally_distributed_database.DedicatedShardDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = string(*obj.CloudAutonomousVmClusterId)
	}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.ContainerDatabaseId != nil {
		result["container_database_id"] = string(*obj.ContainerDatabaseId)
	}

	if obj.ContainerDatabaseParentId != nil {
		result["container_database_parent_id"] = string(*obj.ContainerDatabaseParentId)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	if obj.EncryptionKeyDetails != nil {
		result["encryption_key_details"] = []interface{}{DedicatedShardOrCatalogEncryptionKeyDetailsToMap(obj.EncryptionKeyDetails)}
	}

	if obj.IsAutoScalingEnabled != nil {
		result["is_auto_scaling_enabled"] = bool(*obj.IsAutoScalingEnabled)
	}

	result["metadata"] = obj.Metadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PeerCloudAutonomousVmClusterId != nil {
		result["peer_cloud_autonomous_vm_cluster_id"] = string(*obj.PeerCloudAutonomousVmClusterId)
	}

	if obj.ShardGroup != nil {
		result["shard_group"] = string(*obj.ShardGroup)
	}

	if obj.ShardSpace != nil {
		result["shard_space"] = string(*obj.ShardSpace)
	}

	result["status"] = string(obj.Status)

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeSslCertificateExpires != nil {
		result["time_ssl_certificate_expires"] = obj.TimeSslCertificateExpires.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	return result
}

func DedicatedShardDetailsToMapForResource(obj oci_globally_distributed_database.DedicatedShardDetails, adminPassword string) map[string]interface{} {
	result := map[string]interface{}{}

	result["admin_password"] = adminPassword

	if obj.CloudAutonomousVmClusterId != nil {
		result["cloud_autonomous_vm_cluster_id"] = string(*obj.CloudAutonomousVmClusterId)
	}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.ContainerDatabaseId != nil {
		result["container_database_id"] = string(*obj.ContainerDatabaseId)
	}

	if obj.ContainerDatabaseParentId != nil {
		result["container_database_parent_id"] = string(*obj.ContainerDatabaseParentId)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	if obj.EncryptionKeyDetails != nil {
		result["encryption_key_details"] = []interface{}{DedicatedShardOrCatalogEncryptionKeyDetailsToMap(obj.EncryptionKeyDetails)}
	}

	if obj.IsAutoScalingEnabled != nil {
		result["is_auto_scaling_enabled"] = bool(*obj.IsAutoScalingEnabled)
	}

	result["metadata"] = obj.Metadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PeerCloudAutonomousVmClusterId != nil {
		result["peer_cloud_autonomous_vm_cluster_id"] = string(*obj.PeerCloudAutonomousVmClusterId)
	}

	if obj.ShardGroup != nil {
		result["shard_group"] = string(*obj.ShardGroup)
	}

	if obj.ShardSpace != nil {
		result["shard_space"] = string(*obj.ShardSpace)
	}

	result["status"] = string(obj.Status)

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeSslCertificateExpires != nil {
		result["time_ssl_certificate_expires"] = obj.TimeSslCertificateExpires.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	return result
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToDedicatedShardOrCatalogEncryptionKeyDetails(fieldKeyFormat string) (oci_globally_distributed_database.DedicatedShardOrCatalogEncryptionKeyDetails, error) {
	result := oci_globally_distributed_database.DedicatedShardOrCatalogEncryptionKeyDetails{}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}

	if kmsKeyVersionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_version_id")); ok {
		tmp := kmsKeyVersionId.(string)
		result.KmsKeyVersionId = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}

func DedicatedShardOrCatalogEncryptionKeyDetailsToMap(obj *oci_globally_distributed_database.DedicatedShardOrCatalogEncryptionKeyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	if obj.KmsKeyVersionId != nil {
		result["kms_key_version_id"] = string(*obj.KmsKeyVersionId)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToGsmDetails(fieldKeyFormat string) (oci_globally_distributed_database.GsmDetails, error) {
	result := oci_globally_distributed_database.GsmDetails{}

	if computeCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_count")); ok {
		//tmp := computeCount.(float32)
		tmp := float32(computeCount.(float64))
		result.ComputeCount = &tmp
	}

	if dataStorageSizeInGbs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gbs")); ok {
		tmp := dataStorageSizeInGbs.(float64)
		result.DataStorageSizeInGbs = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		result.Metadata = metadata.(map[string]interface{})
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_globally_distributed_database.GsmDetailsStatusEnum(status.(string))
	}

	if supportingResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "supporting_resource_id")); ok {
		tmp := supportingResourceId.(string)
		result.SupportingResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeSslCertificateExpires, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_ssl_certificate_expires")); ok {
		tmp, err := time.Parse(time.RFC3339, timeSslCertificateExpires.(string))
		if err != nil {
			return result, err
		}
		result.TimeSslCertificateExpires = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return result, err
		}
		result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func GsmDetailsToMap(obj oci_globally_distributed_database.GsmDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ComputeCount != nil {
		result["compute_count"] = float32(*obj.ComputeCount)
	}

	if obj.DataStorageSizeInGbs != nil {
		result["data_storage_size_in_gbs"] = float64(*obj.DataStorageSizeInGbs)
	}

	result["metadata"] = obj.Metadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["status"] = string(obj.Status)

	if obj.SupportingResourceId != nil {
		result["supporting_resource_id"] = string(*obj.SupportingResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeSslCertificateExpires != nil {
		result["time_ssl_certificate_expires"] = obj.TimeSslCertificateExpires.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	return result
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_globally_distributed_database.PatchInstruction, error) {
	var baseObject oci_globally_distributed_database.PatchInstruction
	//discriminator
	operationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation"))
	var operation string
	if ok {
		operation = operationRaw.(string)
	} else {
		operation = "" // default value
	}
	switch strings.ToLower(operation) {
	case strings.ToLower("INSERT"):
		details := oci_globally_distributed_database.PatchInsertInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("MERGE"):
		details := oci_globally_distributed_database.PatchMergeInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOVE"):
		details := oci_globally_distributed_database.PatchRemoveInstruction{}
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

func PatchInstructionToMap(obj oci_globally_distributed_database.PatchInstruction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_globally_distributed_database.PatchInsertInstruction:
		result["operation"] = "INSERT"

		if v.Value != nil {
			//result["value"] = []interface{}{objectToMap(v.Value)}
			result["value"] = []interface{}{v.Value}
		}
	case oci_globally_distributed_database.PatchMergeInstruction:
		result["operation"] = "MERGE"

		if v.Value != nil {
			result["value"] = []interface{}{v.Value}
		}
	case oci_globally_distributed_database.PatchRemoveInstruction:
		result["operation"] = "REMOVE"
	default:
		log.Printf("[WARN] Received 'operation' of unknown type %v", obj)
		return nil
	}

	return result
}

func ShardedDatabaseSummaryToMap(obj oci_globally_distributed_database.ShardedDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_globally_distributed_database.DedicatedShardedDatabaseSummary:
		result["db_deployment_type"] = "DEDICATED"

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.FreeformTags != nil {
			result["freeform_tags"] = v.FreeformTags
		}

		if v.Id != nil {
			result["id"] = v.Id
		}

		result["state"] = v.LifecycleState

		result["lifecycle_state"] = v.LifecycleState

		result["lifecycle_state_details"] = v.LifecycleStateDetails

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_created"] = v.TimeUpdated.String()
		}

		if v.CharacterSet != nil {
			result["character_set"] = string(*v.CharacterSet)
		}

		if v.Chunks != nil {
			result["chunks"] = int(*v.Chunks)
		}

		if v.ClusterCertificateCommonName != nil {
			result["cluster_certificate_common_name"] = string(*v.ClusterCertificateCommonName)
		}

		if v.DbVersion != nil {
			result["db_version"] = string(*v.DbVersion)
		}

		result["db_workload"] = string(v.DbWorkload)

		if v.ListenerPort != nil {
			result["listener_port"] = int(*v.ListenerPort)
		}

		if v.ListenerPortTls != nil {
			result["listener_port_tls"] = int(*v.ListenerPortTls)
		}

		if v.NcharacterSet != nil {
			result["ncharacter_set"] = string(*v.NcharacterSet)
		}

		if v.OnsPortLocal != nil {
			result["ons_port_local"] = int(*v.OnsPortLocal)
		}

		if v.OnsPortRemote != nil {
			result["ons_port_remote"] = int(*v.OnsPortRemote)
		}

		if v.Prefix != nil {
			result["prefix"] = string(*v.Prefix)
		}

		result["sharding_method"] = string(v.ShardingMethod)

		if v.TotalCpuCount != nil {
			result["total_cpu_count"] = int(*v.TotalCpuCount)
		}

		if v.TotalDataStorageSizeInGbs != nil {
			result["total_data_storage_size_in_gbs"] = float64(*v.TotalDataStorageSizeInGbs)
		}
	default:
		log.Printf("[WARN] Received 'db_deployment_type' of unknown type %v", obj)
		return nil
	}

	return result
}

/*
func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) mapToobject(fieldKeyFormat string) (oci_globally_distributed_database.Object, error) {
	result := oci_globally_distributed_database.Object{}

	return result, nil
}

func objectToMap(obj *oci_globally_distributed_database.Object) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}
*/

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) populateTopLevelPolymorphicCreateShardedDatabaseRequest(request *oci_globally_distributed_database.CreateShardedDatabaseRequest) error {
	//discriminator
	dbDeploymentTypeRaw, ok := s.D.GetOkExists("db_deployment_type")
	var dbDeploymentType string
	if ok {
		dbDeploymentType = dbDeploymentTypeRaw.(string)
	} else {
		dbDeploymentType = "" // default value
	}
	switch strings.ToLower(dbDeploymentType) {
	case strings.ToLower("DEDICATED"):
		details := oci_globally_distributed_database.CreateDedicatedShardedDatabase{}
		if catalogDetails, ok := s.D.GetOkExists("catalog_details"); ok {
			interfaces := catalogDetails.([]interface{})
			tmp := make([]oci_globally_distributed_database.CreateDedicatedCatalogDetail, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "catalog_details", stateDataIndex)
				converted, err := s.mapToCreateDedicatedCatalogDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("catalog_details") {
				details.CatalogDetails = tmp
			}
		}
		if characterSet, ok := s.D.GetOkExists("character_set"); ok {
			tmp := characterSet.(string)
			details.CharacterSet = &tmp
		}
		if chunks, ok := s.D.GetOkExists("chunks"); ok {
			tmp := chunks.(int)
			details.Chunks = &tmp
		}
		if clusterCertificateCommonName, ok := s.D.GetOkExists("cluster_certificate_common_name"); ok {
			tmp := clusterCertificateCommonName.(string)
			details.ClusterCertificateCommonName = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
			details.DbWorkload = oci_globally_distributed_database.CreateDedicatedShardedDatabaseDbWorkloadEnum(dbWorkload.(string))
		}
		if listenerPort, ok := s.D.GetOkExists("listener_port"); ok {
			tmp := listenerPort.(int)
			details.ListenerPort = &tmp
		}
		if listenerPortTls, ok := s.D.GetOkExists("listener_port_tls"); ok {
			tmp := listenerPortTls.(int)
			details.ListenerPortTls = &tmp
		}
		if ncharacterSet, ok := s.D.GetOkExists("ncharacter_set"); ok {
			tmp := ncharacterSet.(string)
			details.NcharacterSet = &tmp
		}
		if onsPortLocal, ok := s.D.GetOkExists("ons_port_local"); ok {
			tmp := onsPortLocal.(int)
			details.OnsPortLocal = &tmp
		}
		if onsPortRemote, ok := s.D.GetOkExists("ons_port_remote"); ok {
			tmp := onsPortRemote.(int)
			details.OnsPortRemote = &tmp
		}
		if prefix, ok := s.D.GetOkExists("prefix"); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		if shardDetails, ok := s.D.GetOkExists("shard_details"); ok {
			interfaces := shardDetails.([]interface{})
			tmp := make([]oci_globally_distributed_database.CreateDedicatedShardDetail, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "shard_details", stateDataIndex)
				converted, err := s.mapToCreateDedicatedShardDetail(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("shard_details") {
				details.ShardDetails = tmp
			}
		}
		if shardingMethod, ok := s.D.GetOkExists("sharding_method"); ok {
			details.ShardingMethod = oci_globally_distributed_database.CreateDedicatedShardedDatabaseShardingMethodEnum(shardingMethod.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateShardedDatabaseDetails = details
	default:
		return fmt.Errorf("unknown db_deployment_type '%v' was specified", dbDeploymentType)
	}
	return nil
}

func (s *GloballyDistributedDatabaseShardedDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_globally_distributed_database.ChangeShardedDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ShardedDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database")

	response, err := s.Client.ChangeShardedDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getShardedDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "globally_distributed_database"), oci_globally_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
