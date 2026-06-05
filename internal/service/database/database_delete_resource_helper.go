// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func deleteDatabaseResourceWithReadBeforeDelete(ctx context.Context, d *schema.ResourceData, m interface{}, sync tfresource.ResourceDeleterWithContext) diag.Diagnostics {
	if reader, ok := sync.(tfresource.ResourceReaderWithContext); ok {
		readResponse := tfresource.ReadResourceWithContext(ctx, reader)
		if readResponse == nil && d.Id() == "" {
			return nil
		}
		return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync, readResponse))
	}
	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}
