// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_SSupPr struct {
	CtrlPr *CT_CtrlPr
}

func NewCT_SSupPr() *CT_SSupPr {
	ret := &CT_SSupPr{}
	return ret
}

func (m *CT_SSupPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SSupPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SSupPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SSupPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SSupPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SSupPr and its children
func (m *CT_SSupPr) Validate() error {
	return m.ValidateWithPath("CT_SSupPr")
}

// ValidateWithPath validates the CT_SSupPr and its children, prefixing error messages with path
func (m *CT_SSupPr) ValidateWithPath(path string) error {
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
