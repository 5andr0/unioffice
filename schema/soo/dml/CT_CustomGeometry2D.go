// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_CustomGeometry2D struct {
	AvLst   *CT_GeomGuideList
	GdLst   *CT_GeomGuideList
	AhLst   *CT_AdjustHandleList
	CxnLst  *CT_ConnectionSiteList
	Rect    *CT_GeomRect
	PathLst *CT_Path2DList
}

func NewCT_CustomGeometry2D() *CT_CustomGeometry2D {
	ret := &CT_CustomGeometry2D{}
	ret.PathLst = NewCT_Path2DList()
	return ret
}

func (m *CT_CustomGeometry2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.AvLst != nil {
		seavLst := xml.StartElement{Name: xml.Name{Local: "a:avLst"}}
		e.EncodeElement(m.AvLst, seavLst)
	}
	if m.GdLst != nil {
		segdLst := xml.StartElement{Name: xml.Name{Local: "a:gdLst"}}
		e.EncodeElement(m.GdLst, segdLst)
	}
	if m.AhLst != nil {
		seahLst := xml.StartElement{Name: xml.Name{Local: "a:ahLst"}}
		e.EncodeElement(m.AhLst, seahLst)
	}
	if m.CxnLst != nil {
		secxnLst := xml.StartElement{Name: xml.Name{Local: "a:cxnLst"}}
		e.EncodeElement(m.CxnLst, secxnLst)
	}
	if m.Rect != nil {
		serect := xml.StartElement{Name: xml.Name{Local: "a:rect"}}
		e.EncodeElement(m.Rect, serect)
	}
	sepathLst := xml.StartElement{Name: xml.Name{Local: "a:pathLst"}}
	e.EncodeElement(m.PathLst, sepathLst)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomGeometry2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PathLst = NewCT_Path2DList()
lCT_CustomGeometry2D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "avLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "avLst"}:
				m.AvLst = NewCT_GeomGuideList()
				if err := d.DecodeElement(m.AvLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gdLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "gdLst"}:
				m.GdLst = NewCT_GeomGuideList()
				if err := d.DecodeElement(m.GdLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "ahLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "ahLst"}:
				m.AhLst = NewCT_AdjustHandleList()
				if err := d.DecodeElement(m.AhLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cxnLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cxnLst"}:
				m.CxnLst = NewCT_ConnectionSiteList()
				if err := d.DecodeElement(m.CxnLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "rect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "rect"}:
				m.Rect = NewCT_GeomRect()
				if err := d.DecodeElement(m.Rect, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "pathLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "pathLst"}:
				if err := d.DecodeElement(m.PathLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_CustomGeometry2D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomGeometry2D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomGeometry2D and its children
func (m *CT_CustomGeometry2D) Validate() error {
	return m.ValidateWithPath("CT_CustomGeometry2D")
}

// ValidateWithPath validates the CT_CustomGeometry2D and its children, prefixing error messages with path
func (m *CT_CustomGeometry2D) ValidateWithPath(path string) error {
	if m.AvLst != nil {
		if err := m.AvLst.ValidateWithPath(path + "/AvLst"); err != nil {
			return err
		}
	}
	if m.GdLst != nil {
		if err := m.GdLst.ValidateWithPath(path + "/GdLst"); err != nil {
			return err
		}
	}
	if m.AhLst != nil {
		if err := m.AhLst.ValidateWithPath(path + "/AhLst"); err != nil {
			return err
		}
	}
	if m.CxnLst != nil {
		if err := m.CxnLst.ValidateWithPath(path + "/CxnLst"); err != nil {
			return err
		}
	}
	if m.Rect != nil {
		if err := m.Rect.ValidateWithPath(path + "/Rect"); err != nil {
			return err
		}
	}
	if err := m.PathLst.ValidateWithPath(path + "/PathLst"); err != nil {
		return err
	}
	return nil
}
