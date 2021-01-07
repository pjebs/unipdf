//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package fjson provides support for loading PDF form field data from JSON data/files.
package fjson ;import (_d "encoding/json";_ff "github.com/unidoc/unipdf/v3/core";_cc "github.com/unidoc/unipdf/v3/model";_db "io";_c "os";);

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_bg ,_gaf :=_c .Open (filePath );if _gaf !=nil {return nil ,_gaf ;};defer _bg .Close ();return LoadFromPDF (_bg );};

// FieldValues implements model.FieldValueProvider interface.
func (_ca *FieldData )FieldValues ()(map[string ]_ff .PdfObject ,error ){_ef :=make (map[string ]_ff .PdfObject );for _ ,_fae :=range _ca ._b {if len (_fae .Value )> 0{_ef [_fae .Name ]=_ff .MakeString (_fae .Value );};};return _ef ,nil ;};

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_b []fieldValue };

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _db .ReadSeeker )(*FieldData ,error ){_aa ,_ffa :=_cc .NewPdfReader (rs );if _ffa !=nil {return nil ,_ffa ;};if _aa .AcroForm ==nil {return nil ,nil ;};var _e []fieldValue ;_fb :=_aa .AcroForm .AllFields ();for _ ,_eg :=range _fb {var _bbe []string ;_dcb :=make (map[string ]struct{});_fd ,_gd :=_eg .FullName ();if _gd !=nil {return nil ,_gd ;};if _ad ,_bc :=_eg .V .(*_ff .PdfObjectString );_bc {_e =append (_e ,fieldValue {Name :_fd ,Value :_ad .Decoded ()});continue ;};var _gf string ;for _ ,_fg :=range _eg .Annotations {_dbg ,_fga :=_ff .GetName (_fg .AS );if _fga {_gf =_dbg .String ();};_fa ,_ffad :=_ff .GetDict (_fg .AP );if !_ffad {continue ;};_ga ,_ :=_ff .GetDict (_fa .Get ("\u004e"));for _ ,_gfg :=range _ga .Keys (){_ae :=_gfg .String ();if _ ,_aae :=_dcb [_ae ];!_aae {_bbe =append (_bbe ,_ae );_dcb [_ae ]=struct{}{};};};_bbb ,_ :=_ff .GetDict (_fa .Get ("\u0044"));for _ ,_bd :=range _bbb .Keys (){_df :=_bd .String ();if _ ,_age :=_dcb [_df ];!_age {_bbe =append (_bbe ,_df );_dcb [_df ]=struct{}{};};};};_gb :=fieldValue {Name :_fd ,Value :_gf ,Options :_bbe };_e =append (_e ,_gb );};_da :=FieldData {_b :_e };return &_da ,nil ;};

// JSON returns the field data as a string in JSON format.
func (_ba FieldData )JSON ()(string ,error ){_gba ,_bce :=_d .MarshalIndent (_ba ._b ,"","\u0020\u0020\u0020\u0020");return string (_gba ),_bce ;};

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _db .Reader )(*FieldData ,error ){var _g FieldData ;_dc :=_d .NewDecoder (r ).Decode (&_g ._b );if _dc !=nil {return nil ,_dc ;};return &_g ,nil ;};

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_bb ,_a :=_c .Open (filePath );if _a !=nil {return nil ,_a ;};defer _bb .Close ();return LoadFromJSON (_bb );};type fieldValue struct{Name string `json:"name"`;Value string `json:"value"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};