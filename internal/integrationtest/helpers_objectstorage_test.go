// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	tf_objectstorage "github.com/oracle/terraform-provider-oci/internal/service/objectstorage"
)

// issue-routing-tag: terraform/default
func TestUnitSafe_splitSizeToOffsetsAndLimits(t *testing.T) {

	offsets, _, _ := tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.DefaultFilePartSize*8 + 1)
	if len(offsets) != 9 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.DefaultFilePartSize*8+1)
		return
	}

	offsets, _, _ = tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.DefaultFilePartSize * 7)
	if len(offsets) != 7 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.DefaultFilePartSize*7)
		return
	}

	offsets, _, _ = tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.DefaultFilePartSize + 1)
	if len(offsets) != 2 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.DefaultFilePartSize+1)
		return
	}

	offsets, _, _ = tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.DefaultFilePartSize)
	if len(offsets) != 1 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.DefaultFilePartSize)
		return
	}

	offsets, _, _ = tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.DefaultFilePartSize / 2)
	if len(offsets) != 1 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.DefaultFilePartSize/2)
		return
	}

	offsets, _, _ = tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.DefaultFilePartSize * tf_objectstorage.MaxCount)
	if len(offsets) != int(tf_objectstorage.MaxCount) {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.DefaultFilePartSize*tf_objectstorage.MaxCount)
		return
	}

	offsets, _, _ = tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.DefaultFilePartSize*tf_objectstorage.MaxCount + 1)
	if len(offsets) != int(tf_objectstorage.MaxCount) {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.DefaultFilePartSize*tf_objectstorage.MaxCount)
		return
	}

	offsets, _, _ = tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.DefaultFilePartSize * tf_objectstorage.MaxCount * 2)
	if len(offsets) != int(tf_objectstorage.MaxCount) {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.DefaultFilePartSize*tf_objectstorage.MaxCount*2)
		return
	}

	offsets, _, _ = tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.MaxPartSize * tf_objectstorage.MaxCount)
	if len(offsets) != int(tf_objectstorage.MaxCount) {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), tf_objectstorage.MaxCount*tf_objectstorage.MaxPartSize)
		return
	}

	_, _, err := tf_objectstorage.SplitSizeToOffsetsAndLimits(tf_objectstorage.MaxPartSize*tf_objectstorage.MaxCount + 1)
	if err == nil {
		t.Errorf("The error should be returned for the too large file size")
		return
	}

	return
}
