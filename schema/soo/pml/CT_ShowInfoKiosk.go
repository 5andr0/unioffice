// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_ShowInfoKiosk struct {
	// Restart Show
	RestartAttr *uint32
}

func NewCT_ShowInfoKiosk() *CT_ShowInfoKiosk {
	ret := &CT_ShowInfoKiosk{}
	return ret
}

func (m *CT_ShowInfoKiosk) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RestartAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "restart"},
			Value: fmt.Sprintf("%v", *m.RestartAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShowInfoKiosk) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "restart" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RestartAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ShowInfoKiosk: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_ShowInfoKiosk and its children
func (m *CT_ShowInfoKiosk) Validate() error {
	return m.ValidateWithPath("CT_ShowInfoKiosk")
}

// ValidateWithPath validates the CT_ShowInfoKiosk and its children, prefixing error messages with path
func (m *CT_ShowInfoKiosk) ValidateWithPath(path string) error {
	return nil
}
