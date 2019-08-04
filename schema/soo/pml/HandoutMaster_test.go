// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/pml"
)

func TestHandoutMasterConstructor(t *testing.T) {
	v := pml.NewHandoutMaster()
	if v == nil {
		t.Errorf("pml.NewHandoutMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.HandoutMaster should validate: %s", err)
	}
}

func TestHandoutMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewHandoutMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewHandoutMaster()
	xml.Unmarshal(buf, v2)
}
