// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/diagram"
)

func TestDataModelConstructor(t *testing.T) {
	v := diagram.NewDataModel()
	if v == nil {
		t.Errorf("diagram.NewDataModel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.DataModel should validate: %s", err)
	}
}

func TestDataModelMarshalUnmarshal(t *testing.T) {
	v := diagram.NewDataModel()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewDataModel()
	xml.Unmarshal(buf, v2)
}
