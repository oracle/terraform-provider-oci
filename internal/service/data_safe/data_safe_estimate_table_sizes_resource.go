// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License 2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// DataSafeEstimateTableSizesResource starts an asynchronous table-size
// estimation for a subsetting policy. The operation is intentionally modeled
// as a create-only resource because the estimate is produced by the action and
// has no independent lifecycle.
func DataSafeEstimateTableSizesResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			policyID := d.Get("subsetting_policy_id").(string)
			credentials := d.Get("target_credentials").([]interface{})
			if len(credentials) == 0 {
				return diag.Errorf("target_credentials must be specified")
			}
			values := credentials[0].(map[string]interface{})
			userName := values["user_name"].(string)
			password := values["password"].(string)

			request := oci_data_safe.EstimateTableSizesRequest{
				SubsettingPolicyId: &policyID,
				EstimateTableSizesDetails: oci_data_safe.EstimateTableSizesDetails{
					TargetCredentials: &oci_data_safe.Credentials{
						UserName: &userName,
						Password: &password,
					},
				},
			}
			if targetID, ok := d.GetOk("target_id"); ok {
				targetIDValue := targetID.(string)
				request.EstimateTableSizesDetails.TargetId = &targetIDValue
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")
			dataSafeClient := m.(*client.OracleClients).DataSafeClient()
			response, err := dataSafeClient.EstimateTableSizes(ctx, request)
			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
			if response.OpcWorkRequestId == nil {
				return diag.Errorf("estimate table sizes operation did not return a work request ID")
			}
			if _, err = subsettingPolicyWaitForWorkRequest(ctx, response.OpcWorkRequestId, "subsettingpolicy", oci_data_safe.WorkRequestResourceActionTypeUpdated, d.Timeout(schema.TimeoutCreate), false, dataSafeClient); err != nil {
				return tfresource.HandleDiagError(m, err)
			}

			d.SetId(policyID)
			return nil
		},
		ReadContext:   func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics { return nil },
		DeleteContext: func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics { return nil },
		Schema: map[string]*schema.Schema{
			"subsetting_policy_id": {Type: schema.TypeString, Required: true, ForceNew: true},
			"target_id":            {Type: schema.TypeString, Optional: true, ForceNew: true},
			"target_credentials": {
				Type: schema.TypeList, Required: true, ForceNew: true, MinItems: 1, MaxItems: 1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"user_name": {Type: schema.TypeString, Required: true},
					"password":  {Type: schema.TypeString, Required: true, Sensitive: true},
				}},
			},
		},
	}
}
