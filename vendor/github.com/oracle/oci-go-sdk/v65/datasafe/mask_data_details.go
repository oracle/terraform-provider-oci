// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaskDataDetails Details to mask data.
type MaskDataDetails struct {

	// The OCID of the target database to be masked. If it's not provided, the value of the
	// targetId attribute in the MaskingPolicy resource is used. The OCID of the target
	// database to be masked. If it's not provided, the value of the targetId attribute in
	// the MaskingPolicy resource is used.
	TargetId *string `mandatory:"false" json:"targetId"`

	// Indicates if the masking request is to decrypt the data values previously encrypted using Deterministic Encryption. Note that, to
	// correctly decrypt the encrypted data values, it requires the same seed value that was provided to encrypt those data values.
	IsDecrypt *bool `mandatory:"false" json:"isDecrypt"`

	// Indicates if the masking request is to rerun the previously failed masking steps. If a masking request is submitted with the
	// isIgnoreErrorsEnabled attribute set to true, the masking process tracks the failed masking steps. Another masking request can be
	// submitted with the isRun attribute set to true to rerun those failed masking steps. It helps save time by executing only the failed
	// masking steps and not doing the whole masking again.
	IsRerun *bool `mandatory:"false" json:"isRerun"`

	// The tablespace that should be used to create the mapping tables, DMASK objects, and other temporary tables for data masking.
	// If no tablespace is provided, the DEFAULT tablespace is used.
	Tablespace *string `mandatory:"false" json:"tablespace"`

	// Indicates if the masking process should continue on hitting an error. It provides fault tolerance support and is enabled by
	// default. In fault-tolerant mode, the masking process saves the failed step and continues. You can then submit another masking
	// request (with isRerun attribute set to true) to execute only the failed steps.
	IsIgnoreErrorsEnabled *bool `mandatory:"false" json:"isIgnoreErrorsEnabled"`

	// The seed value to be used in case of Deterministic Encryption and Deterministic Substitution masking formats.
	Seed *string `mandatory:"false" json:"seed"`

	// Indicates if the interim DMASK tables should be moved to the user-specified tablespace. As interim tables can be large in size,
	// set it to false if moving them causes performance overhead during masking.
	IsMoveInterimTablesEnabled *bool `mandatory:"false" json:"isMoveInterimTablesEnabled"`

	// Indicates if data masking should be performed using a saved masking script. Setting this attribute to true skips masking script
	// generation and executes the masking script stored in the Data Safe repository. It helps save time if there are no changes in
	// the database tables and their dependencies.
	IsExecuteSavedScriptEnabled *bool `mandatory:"false" json:"isExecuteSavedScriptEnabled"`

	// Indicates if the temporary tables created during a masking operation should be dropped after masking.
	// Set this attribute to false to preserve the temporary tables. Masking creates temporary tables that map the original sensitive
	// data values to mask values. These temporary tables are dropped after masking if this attribute is set as true. But, in some cases, you may want
	// to preserve this information to track how masking changed your data. Note that doing so compromises security. These tables
	// must be dropped before the database is available for unprivileged users.
	// If it's not provided, the value of the isDropTempTablesEnabled attribute in the MaskingPolicy resource is used.
	IsDropTempTablesEnabled *bool `mandatory:"false" json:"isDropTempTablesEnabled"`

	// Indicates if redo logging is enabled during a masking operation. Set this attribute to true to
	// enable redo logging. If set as false, masking disables redo logging and flashback logging to purge any original unmasked
	// data from logs. However, in certain circumstances when you only want to test masking, rollback changes, and retry masking,
	// you could enable logging and use a flashback database to retrieve the original unmasked data after it has been masked.
	// If it's not provided, the value of the isRedoLoggingEnabled attribute in the MaskingPolicy resource is used.
	IsRedoLoggingEnabled *bool `mandatory:"false" json:"isRedoLoggingEnabled"`

	// Indicates if statistics gathering is enabled. Set this attribute to false to disable statistics
	// gathering. The masking process gathers statistics on masked database tables after masking completes.
	// If it's not provided, the value of the isRefreshStatsEnabled attribute in the MaskingPolicy resource is used.
	IsRefreshStatsEnabled *bool `mandatory:"false" json:"isRefreshStatsEnabled"`

	// Specifies options to enable parallel execution when running data masking. Allowed values are 'NONE' (no parallelism),
	// 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree
	// of parallelism. Parallel execution helps effectively use multiple CPUs and improve masking performance. Refer to the
	// Oracle Database parallel execution framework when choosing an explicit degree of parallelism.
	// https://www.oracle.com/pls/topic/lookup?ctx=dblatest&en/database/oracle/oracle-database&id=VLDBG-GUID-3E2AE088-2505-465E-A8B2-AC38813EA355
	// If it's not provided, the value of the parallelDegree attribute in the MaskingPolicy resource is used.
	ParallelDegree *string `mandatory:"false" json:"parallelDegree"`

	// Specifies how to recompile invalid objects post data masking. Allowed values are 'SERIAL' (recompile in serial),
	// 'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree
	// attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after masking completes.
	// If it's not provided, the value of the parallelDegree attribute in the MaskingPolicy resource is used.
	Recompile *string `mandatory:"false" json:"recompile"`
}

func (m MaskDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
