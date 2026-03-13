// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedDatabasePrivateEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDistributedDatabaseDistributedDatabasePrivateEndpointWithContext,
		ReadContext:   readDistributedDatabaseDistributedDatabasePrivateEndpointWithContext,
		UpdateContext: updateDistributedDatabaseDistributedDatabasePrivateEndpointWithContext,
		DeleteContext: deleteDistributedDatabaseDistributedDatabasePrivateEndpointWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"description": {
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
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"reinstate_proxy_instance_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"globally_distributed_autonomous_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_deployment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"globally_distributed_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_deployment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_compute_instance_id": {
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
			"vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDistributedDatabaseDistributedDatabasePrivateEndpointWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbPrivateEndpointServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The code generator returns a plain `error` from a Context-based CRUD function,
	// but Terraform SDK v2 requires `diag.Diagnostics` as the return type.
	// Returning `error` directly results in a compile-time failure.
	//
	// Convert the error using tfresource.HandleDiagError to satisfy the SDK contract.
	// See JIRA: TOP-9389

	if _, ok := sync.D.GetOkExists("reinstate_proxy_instance_trigger"); ok {
		/*err := sync.ReinstateProxyInstance()
		if err != nil {
			return err
		}*/
		err := sync.ReinstateProxyInstanceWithContext(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	return nil

}

func readDistributedDatabaseDistributedDatabasePrivateEndpointWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbPrivateEndpointServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDistributedDatabaseDistributedDatabasePrivateEndpointWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbPrivateEndpointServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if _, ok := sync.D.GetOkExists("reinstate_proxy_instance_trigger"); ok && sync.D.HasChange("reinstate_proxy_instance_trigger") {
		oldRaw, newRaw := sync.D.GetChange("reinstate_proxy_instance_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ReinstateProxyInstanceWithContext(ctx)

			// WORKAROUND FOR GENERATED CODE ISSUE:
			// Context-based Update functions must return `diag.Diagnostics`.
			// The generator incorrectly propagates a raw `error`, which is incompatible
			// with the Terraform Plugin SDK v2 function signature.
			//
			// Convert the error into Diagnostics using tfresource.HandleDiagError.
			// See JIRA: TOP-9389

			/*if err != nil {
				return err
			}*/
			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}

		} else {
			// WORKAROUND FOR GENERATED CODE ISSUE:
			// Validation errors in Context-based CRUD functions must also return
			// `diag.Diagnostics`. Returning `fmt.Errorf(...)` directly causes
			// a type mismatch at compile time.
			//
			// Wrap the error using tfresource.HandleDiagError to comply with
			// Terraform SDK v2 expectations.
			// See JIRA: TOP-9389
			sync.D.Set("reinstate_proxy_instance_trigger", oldRaw)
			// return fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(
				m,
				fmt.Errorf("new value of trigger should be greater than the old value"),
			)
		}
	}
	// NOTE (TOP-9398):
	// UpdateResourceWithContext requires an explicit context.Context parameter
	// in newer tfresource helpers. Passing the Terraform operation context
	// ensures proper timeout, retry, and cancellation propagation.
	/*if err := tfresource.UpdateResourceWithContext(d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}*/

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil
}

func deleteDistributedDatabaseDistributedDatabasePrivateEndpointWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbPrivateEndpointServiceClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_distributed_database.DistributedDbPrivateEndpointServiceClient
	Res                    *oci_distributed_database.DistributedDatabasePrivateEndpoint
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_distributed_database.DistributedDbWorkRequestServiceClient
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateCreating),
	}
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateActive),
	}
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateDeleting),
	}
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_distributed_database.DistributedDatabasePrivateEndpointLifecycleStateDeleted),
	}
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_distributed_database.CreateDistributedDatabasePrivateEndpointRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.CreateDistributedDatabasePrivateEndpoint(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDistributedDatabasePrivateEndpointFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) getDistributedDatabasePrivateEndpointFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_distributed_database.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	distributedDatabasePrivateEndpointId, err := distributedDatabasePrivateEndpointWaitForWorkRequest(ctx, workId, "distributeddatabaseprivateendpoint",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*distributedDatabasePrivateEndpointId)

	return s.GetWithContext(ctx)
}

func distributedDatabasePrivateEndpointWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "distributed_database", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_distributed_database.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func distributedDatabasePrivateEndpointWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_distributed_database.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_distributed_database.DistributedDbWorkRequestServiceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "distributed_database")
	retryPolicy.ShouldRetryOperation = distributedDatabasePrivateEndpointWorkRequestShouldRetryFunc(timeout)

	response := oci_distributed_database.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_distributed_database.OperationStatusInProgress),
			string(oci_distributed_database.OperationStatusAccepted),
			string(oci_distributed_database.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_distributed_database.OperationStatusSucceeded),
			string(oci_distributed_database.OperationStatusFailed),
			string(oci_distributed_database.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_distributed_database.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_distributed_database.OperationStatusFailed || response.Status == oci_distributed_database.OperationStatusCanceled {
		return nil, getErrorFromDistributedDatabaseDistributedDatabasePrivateEndpointWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDistributedDatabaseDistributedDatabasePrivateEndpointWorkRequest(ctx context.Context, client *oci_distributed_database.DistributedDbWorkRequestServiceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_distributed_database.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_distributed_database.ListWorkRequestErrorsRequest{
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

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedDatabasePrivateEndpointRequest{}

	tmp := s.D.Id()
	request.DistributedDatabasePrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.GetDistributedDatabasePrivateEndpoint(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedDatabasePrivateEndpoint
	return nil
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_distributed_database.UpdateDistributedDatabasePrivateEndpointRequest{}

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

	tmp := s.D.Id()
	request.DistributedDatabasePrivateEndpointId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.UpdateDistributedDatabasePrivateEndpoint(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DistributedDatabasePrivateEndpoint
	return nil
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_distributed_database.DeleteDistributedDatabasePrivateEndpointRequest{}

	tmp := s.D.Id()
	request.DistributedDatabasePrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.DeleteDistributedDatabasePrivateEndpoint(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := distributedDatabasePrivateEndpointWaitForWorkRequest(ctx, workId, "distributeddatabaseprivateendpoint",
		oci_distributed_database.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) SetData() error {
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

	globallyDistributedAutonomousDatabases := []interface{}{}
	for _, item := range s.Res.GloballyDistributedAutonomousDatabases {
		globallyDistributedAutonomousDatabases = append(globallyDistributedAutonomousDatabases, DistributedAutonomousDatabaseAssociatedWithPrivateEndpointToMap(item))
	}
	s.D.Set("globally_distributed_autonomous_databases", globallyDistributedAutonomousDatabases)

	globallyDistributedDatabases := []interface{}{}
	for _, item := range s.Res.GloballyDistributedDatabases {
		globallyDistributedDatabases = append(globallyDistributedDatabases, DistributedDatabaseAssociatedWithPrivateEndpointToMap(item))
	}
	s.D.Set("globally_distributed_databases", globallyDistributedDatabases)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.PrivateIp != nil {
		s.D.Set("private_ip", *s.Res.PrivateIp)
	}

	if s.Res.ProxyComputeInstanceId != nil {
		s.D.Set("proxy_compute_instance_id", *s.Res.ProxyComputeInstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

/*
func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) ReinstateProxyInstance() error {
	request := oci_distributed_database.ReinstateProxyInstanceRequest{}

	idTmp := s.D.Id()
	request.DistributedDatabasePrivateEndpointId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	// NOTE (TOP-9400):
	// ReinstateProxyInstance is an action-style API whose OCI Go SDK response
	// does NOT include a DistributedDatabasePrivateEndpoint payload.
	// The code generator incorrectly assumes the response contains the resource
	// model and emits an assignment like `s.Res = &response.<Resource>`,
	// which does not compile.
	// For now, the response-to-resource assignment is intentionally omitted.
	// The resource state should be refreshed via an explicit GET (or work-request
	// wait + GET) once proper generator support is added.
	//

	/*response, err := s.Client.ReinstateProxyInstance(context.Background(), request)
	if err != nil {
		return err
	}*/
// NOTE (TOP-9398):
// The legacy WaitForUpdatedState helper requires the non-context
// ResourceUpdater interface (Update()), which this CRUD intentionally
// does not implement. Use the context-aware waiter instead to align
// with UpdateWithContext-based CRUD implementations.
/*
	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}
*/
/*
	_, err := s.Client.ReinstateProxyInstance(context.Background(), request)
	if err != nil {
		return err
	}



	if err := tfresource.WaitForUpdatedStateWithContext(context.Background(), s.D, s); err != nil {
		return err
	}

	val := s.D.Get("reinstate_proxy_instance_trigger")
	s.D.Set("reinstate_proxy_instance_trigger", val)
*/

// NOTE (TOP-9400):
// ReinstateProxyInstance is an action-style API whose OCI Go SDK response
// does NOT include a DistributedDatabasePrivateEndpoint payload.
// The code generator incorrectly assumes the response contains the resource
// model and emits an assignment like `s.Res = &response.<Resource>`,
// which does not compile.
// For now, the response-to-resource assignment is intentionally omitted.
// The resource state should be refreshed via an explicit GET (or work-request
// wait + GET) once proper generator support is added.
//

//s.Res = &response.DistributedDatabasePrivateEndpoint
/*
	return nil
}*/

// WORKAROUND FOR CODE GENERATOR ISSUE (TOP-9443):
//
// ReinstateProxyInstance is an *action-style* API executed asynchronously via
// a Work Request. The OCI Go SDK response does NOT include a
// DistributedDatabasePrivateEndpoint payload.
//
// The code generator incorrectly treats this API as a standard Update:
//   - It assumes the response contains the resource model
//   - It chains UpdateResourceWithContext() after invoking the action
//   - This results in the Reinstate API being called twice
//
// The first call transitions the Private Endpoint to UPDATING.
// The second call fails with a lifecycle error (OSD-10165) because the
// resource is already updating.
//
// Correct behavior for this API is:
//   1. Invoke ReinstateProxyInstance
//   2. Wait for the associated Work Request to complete
//   3. Refresh resource state via explicit GET
//   4. Do NOT run generic Update logic afterward
//
// This method implements that behavior explicitly, bypassing the
// generated Update flow and ensuring correct lifecycle handling

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) ReinstateProxyInstanceWithContext(ctx context.Context) error {
	req := oci_distributed_database.ReinstateProxyInstanceRequest{}
	id := s.D.Id()
	req.DistributedDatabasePrivateEndpointId = &id
	req.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	resp, err := s.Client.ReinstateProxyInstance(ctx, req)
	if err != nil {
		return err
	}

	workId := resp.OpcWorkRequestId
	if workId == nil {
		// Defensive: if service ever doesn’t return it, fall back to a refresh.
		return s.GetWithContext(ctx)
	}

	// Wait for WR to finish (don’t rely on actionType matching; reinstate isn’t Created/Updated/Deleted).
	if err := waitForDistributedDbWorkRequestCompletion(ctx, workId, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.WorkRequestClient); err != nil {
		return err
	}

	// Refresh resource
	if err := s.GetWithContext(ctx); err != nil {
		return err
	}
	return s.SetData()
}

// Work-request waiter that only cares about WR status == SUCCEEDED (or FAILED/CANCELED)
func waitForDistributedDbWorkRequestCompletion(
	ctx context.Context,
	wId *string,
	timeout time.Duration,
	disableNotFoundRetries bool,
	client *oci_distributed_database.DistributedDbWorkRequestServiceClient,
) error {
	retryPolicy := tfresource.GetRetryPolicy(disableNotFoundRetries, "distributed_database")
	retryPolicy.ShouldRetryOperation = distributedDatabasePrivateEndpointWorkRequestShouldRetryFunc(timeout)

	var resp oci_distributed_database.GetWorkRequestResponse
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_distributed_database.OperationStatusInProgress),
			string(oci_distributed_database.OperationStatusAccepted),
			string(oci_distributed_database.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_distributed_database.OperationStatusSucceeded),
			string(oci_distributed_database.OperationStatusFailed),
			string(oci_distributed_database.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			resp, err = client.GetWorkRequest(ctx, oci_distributed_database.GetWorkRequestRequest{
				WorkRequestId: wId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
			wr := &resp.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return err
	}

	if resp.Status == oci_distributed_database.OperationStatusFailed || resp.Status == oci_distributed_database.OperationStatusCanceled {
		return getErrorFromDistributedDatabaseDistributedDatabasePrivateEndpointWorkRequest(ctx, client, wId, retryPolicy, "distributeddatabaseprivateendpoint", oci_distributed_database.ActionTypeUpdated)
	}
	return nil
}

func DistributedAutonomousDatabaseAssociatedWithPrivateEndpointToMap(obj oci_distributed_database.DistributedAutonomousDatabaseAssociatedWithPrivateEndpoint) map[string]interface{} {
	result := map[string]interface{}{}

	result["db_deployment_type"] = string(obj.DbDeploymentType)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func DistributedDatabaseAssociatedWithPrivateEndpointToMap(obj oci_distributed_database.DistributedDatabaseAssociatedWithPrivateEndpoint) map[string]interface{} {
	result := map[string]interface{}{}

	result["db_deployment_type"] = string(obj.DbDeploymentType)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func DistributedDatabasePrivateEndpointSummaryToMap(obj oci_distributed_database.DistributedDatabasePrivateEndpointSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

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

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		// NOTE (TOP-9425):
		// The code generator emitted a reference to a non-existent helper
		// `nsgIdsHashCodeForSets`, which causes a compilation failure.
		// For TypeSet attributes like `nsg_ids` (set of strings), reuse the
		// standard hash function already defined in the schema to ensure
		// consistent behavior and avoid undefined-symbol errors.
		//result["nsg_ids"] = schema.NewSet(nsgIdsHashCodeForSets, nsgIds)
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)

	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func (s *DistributedDatabaseDistributedDatabasePrivateEndpointResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_distributed_database.ChangeDistributedDatabasePrivateEndpointCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DistributedDatabasePrivateEndpointId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")

	response, err := s.Client.ChangeDistributedDatabasePrivateEndpointCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDistributedDatabasePrivateEndpointFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database"), oci_distributed_database.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
