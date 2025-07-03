package kms

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"
)

func KmsVaultReplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createKmsVaultReplica,
		Read:     readKmsVaultReplica,
		Delete:   deleteKmsVaultReplica,
		Schema: map[string]*schema.Schema{
			// Required
			"vault_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"replica_region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"replica_vault_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"idcs_account_name_url": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"vault_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"private_endpoint_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vault_replica_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKmsVaultReplica(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.CreateResource(d, sync)
}

func readKmsVaultReplica(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.ReadResource(sync)
}

func deleteKmsVaultReplica(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.DeleteResource(d, sync)
}

type KmsVaultReplicaResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.KmsVaultClient
	Res                    *oci_kms.VaultReplicaSummary
	DisableNotFoundRetries bool
}

func (s *KmsVaultReplicaResourceCrud) ID() string {
	log.Printf("[INFO] ID()")

	vaultIdStr := "vault_id"
	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		vaultIdStr = tmp
	}
	replicaRegionStr := *s.Res.Region

	log.Printf("[INFO] ID() Setting ID: %s", fmt.Sprintf("%s:%s", vaultIdStr, replicaRegionStr))
	return fmt.Sprintf("%s:%s", vaultIdStr, replicaRegionStr)
}

func (s *KmsVaultReplicaResourceCrud) Create() error {
	log.Printf("[INFO] Create()")
	request := oci_kms.CreateVaultReplicaRequest{}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	if replicaRegion, ok := s.D.GetOkExists("replica_region"); ok {
		tmp := replicaRegion.(string)
		request.ReplicaRegion = &tmp
	}

	if replicaVaultMetadata, ok := s.D.GetOkExists("replica_vault_metadata"); ok {
		if tmpList := replicaVaultMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replica_vault_metadata", 0)
			replicaExternalVaultMetadata := oci_kms.ReplicaExternalVaultMetadata{}

			if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
				tmp := privateEndpointId.(string)
				replicaExternalVaultMetadata.PrivateEndpointId = &tmp
			}

			if idcsAccountNameUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idcs_account_name_url")); ok {
				tmp := idcsAccountNameUrl.(string)
				replicaExternalVaultMetadata.IdcsAccountNameUrl = &tmp
			}

			request.ReplicaVaultMetadata = &replicaExternalVaultMetadata
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	log.Printf("[INFO] Create() Sending Create Replica Request")
	_, err := s.Client.CreateVaultReplica(context.Background(), request)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Create() Waiting for Replica Creation")
	retentionPolicyFunc := func() bool { return s.Res.Status == oci_kms.VaultReplicaSummaryStatusCreated }
	waitErr := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutCreate))
	if waitErr != nil {
		return waitErr
	}
	log.Printf("[INFO] Create() Replica Creation Completed")

	return nil
}

func (s *KmsVaultReplicaResourceCrud) Get() error {
	log.Printf("[INFO] Get()")

	request := oci_kms.ListVaultReplicasRequest{}

	/*
	 *The Vault ID is only present in the config and not in the Response so we need to check
	 * if Vault ID is available in the state/config then use that else extract it from the ID
	 * We can't always use the ID as it is only set after resource creation or provided during import
	 */
	vaultIdStr := "vault_id"
	replicaRegionStr := "replica_region"

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		vaultIdStr = tmp
		request.VaultId = &tmp
	}
	if replicaRegion, ok := s.D.GetOkExists("replica_region"); ok {
		tmp := replicaRegion.(string)
		replicaRegionStr = tmp
	}
	// If the Vault ID or Replica Region String didn't get updated from state/config, use ID
	if vaultIdStr == "vault_id" || replicaRegionStr == "replica_region" {
		log.Printf("[INFO] Get() Vault ID or Replica Region String didn't get updated from state/config, using Resource ID")
		parts := strings.Split(s.D.Id(), ":")
		if len(parts) > 1 {
			vaultId := parts[0]
			request.VaultId = &vaultId
			replicaRegionStr = parts[1]
		} else {
			log.Fatalf("[ERROR] Get() unable to parse current ID: %s. The expected format of the ID is \"{vault_id}:{replica_region}\"", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	log.Printf("[INFO] Get() Calling List Vault Replicas")
	response, err := s.Client.ListVaultReplicas(context.Background(), request)
	if err != nil {
		return err
	}

	vaultReplicaSummaryRes := oci_kms.VaultReplicaSummary{Status: oci_kms.VaultReplicaSummaryStatusDeleted}
	for _, vaultReplicaSummary := range response.Items {
		vaultReplicaRegion := *(vaultReplicaSummary.Region)
		if vaultReplicaRegion == replicaRegionStr {
			vaultReplicaSummaryRes = vaultReplicaSummary
		}
	}
	s.Res = &vaultReplicaSummaryRes
	// Explicitly Call VoidState to Remove Deleted Vault Replica from State
	if s.Res.Status == oci_kms.VaultReplicaSummaryStatusDeleted {
		log.Printf("[INFO] VoidState() as no resource is found, removing from state")
		s.VoidState()
	}

	return nil
}

func (s *KmsVaultReplicaResourceCrud) Delete() error {
	log.Printf("[INFO] Delete()")
	request := oci_kms.DeleteVaultReplicaRequest{}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	if replicaRegion, ok := s.D.GetOkExists("replica_region"); ok {
		tmp := replicaRegion.(string)
		request.ReplicaRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.DeleteVaultReplica(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res == nil || s.Res.Status == oci_kms.VaultReplicaSummaryStatusDeleted }
	waitErr := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutDelete))
	if waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *KmsVaultReplicaResourceCrud) SetData() error {
	log.Printf("[INFO] SetData()")

	parts := strings.Split(s.D.Id(), ":")
	if len(parts) > 1 {
		vaultId := parts[0]
		s.D.Set("vault_id", &vaultId)
	}

	if s.Res.CryptoEndpoint != nil {
		s.D.Set("crypto_endpoint", *s.Res.CryptoEndpoint)
	}

	if s.Res.ManagementEndpoint != nil {
		s.D.Set("management_endpoint", *s.Res.ManagementEndpoint)
	}

	if s.Res.Region != nil {
		s.D.Set("replica_region", *s.Res.Region)
	}

	s.D.Set("vault_replica_status", s.Res.Status)

	return nil
}

// State Necessary to have. Otherwise, cause "Could not set resource state, sync did not have a valid .Res.State, .Resource.State, or .WorkRequest.State" error from setState() in crud_helper
func (s *KmsVaultReplicaResourceCrud) State() oci_kms.VaultReplicaSummaryStatusEnum {
	log.Printf("[INFO] State()")
	return s.Res.Status
}

func (s *KmsVaultReplicaResourceCrud) CreatedPending() []string {
	log.Printf("[INFO] CreatedPending()")
	return []string{
		string(oci_kms.VaultReplicaSummaryStatusCreating),
	}
}

func (s *KmsVaultReplicaResourceCrud) CreatedTarget() []string {
	log.Printf("[INFO] CreatedTarget()")
	return []string{
		string(oci_kms.VaultReplicaSummaryStatusCreated),
	}
}

func (s *KmsVaultReplicaResourceCrud) DeletedPending() []string {
	log.Printf("[INFO] DeletedPending()")
	return []string{
		string(oci_kms.VaultReplicaSummaryStatusDeleting),
	}
}

func (s *KmsVaultReplicaResourceCrud) DeletedTarget() []string {
	log.Printf("[INFO] DeletedTarget()")
	return []string{
		string(oci_kms.VaultReplicaSummaryStatusDeleted),
	}
}
