// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Read:   tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CreateContext: createDistributedDatabaseDistributedAutonomousDatabasesConfigureShardingWithContext,
		ReadContext:   readDistributedDatabaseDistributedAutonomousDatabasesConfigureShardingWithContext,
		DeleteContext: deleteDistributedDatabaseDistributedAutonomousDatabasesConfigureShardingWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"distributed_autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"ca_bundle_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_rebalance_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"work_request_id": {Type: schema.TypeString, Computed: true},
			"action_pending":  {Type: schema.TypeBool, Computed: true, Description: "An accepted action still requires completion polling. Retries resume its saved work request."},
		},
	}
}

func createDistributedDatabaseDistributedAutonomousDatabasesConfigureShardingWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if err := sync.CreateWithContext(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	return nil
}

func readDistributedDatabaseDistributedAutonomousDatabasesConfigureShardingWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.Get("action_pending").(bool) {
		return nil
	}
	sync := &DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()
	if err := sync.resumeAction(ctx, schema.TimeoutRead); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	return nil
}

func deleteDistributedDatabaseDistributedAutonomousDatabasesConfigureShardingWithContext(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

type DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_distributed_database.DistributedAutonomousDbServiceClient
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_distributed_database.DistributedDbWorkRequestServiceClient
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceCrud) createAction(ctx context.Context) error {
	if s.D.Get("action_pending").(bool) {
		return s.resumeAction(ctx, schema.TimeoutCreate)
	}
	request := oci_distributed_database.ConfigureDistributedAutonomousDatabaseShardingRequest{}
	distributedAutonomousDatabaseID := s.D.Get("distributed_autonomous_database_id").(string)
	if err := waitForDistributedAutonomousDatabaseActionReady(ctx, s.D, s.Client, distributedAutonomousDatabaseID); err != nil {
		return err
	}

	if caBundleID, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleID.(string)
		request.CaBundleId = &tmp
	}
	if certificateID, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateID.(string)
		request.CertificateId = &tmp
	}
	request.DistributedAutonomousDatabaseId = &distributedAutonomousDatabaseID
	if isRebalanceRequired, ok := s.D.GetOkExists("is_rebalance_required"); ok {
		tmp := isRebalanceRequired.(bool)
		request.IsRebalanceRequired = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")
	response, err := s.Client.ConfigureDistributedAutonomousDatabaseSharding(ctx, request)
	if err != nil {
		return err
	}

	workID := response.OpcWorkRequestId
	if workID == nil || strings.TrimSpace(*workID) == "" {
		return fmt.Errorf("missing opc-work-request-id for ConfigureDistributedAutonomousDatabaseSharding")
	}
	s.D.SetId(distributedDatabaseActionID("DistributedAutonomousDatabaseConfigureShardingAction-", s.D, "distributed_autonomous_database_id", "ca_bundle_id", "certificate_id", "is_rebalance_required"))
	if err := s.D.Set("work_request_id", *workID); err != nil {
		return err
	}
	if err := s.D.Set("action_pending", true); err != nil {
		return err
	}
	return s.resumeAction(ctx, schema.TimeoutCreate)
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceCrud) CreateWithContext(ctx context.Context) error {
	return s.createAction(ctx)
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceCrud) resumeAction(ctx context.Context, timeoutKey string) error {
	work := strings.TrimSpace(s.D.Get("work_request_id").(string))
	if work == "" {
		return fmt.Errorf("pending configure-sharding action has no work_request_id; refusing to resubmit")
	}
	ctx, cancel := context.WithTimeout(ctx, s.D.Timeout(timeoutKey))
	defer cancel()
	if _, err := distributedAutonomousDatabaseWaitForWorkRequest(ctx, &work, "", oci_distributed_database.ActionTypeUpdated, s.D.Timeout(timeoutKey), s.DisableNotFoundRetries, s.WorkRequestClient); err != nil {
		return err
	}
	id := s.D.Get("distributed_autonomous_database_id").(string)
	if err := waitForDistributedAutonomousDatabaseActionReady(ctx, s.D, s.Client, id); err != nil {
		return err
	}
	return s.D.Set("action_pending", false)
}
