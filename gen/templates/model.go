//go:build ignore
// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/sjson"
	"github.com/tidwall/gjson"
)

{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
{{- if and (not .NoDelete) (not .NoDeleteAttributes)}}
	DeleteMode types.String `tfsdk:"delete_mode"`
{{- end}}
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "StringList") (eq .Type "Int64List")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}

type {{camelCase .Name}}Data struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "StringList") (eq .Type "Int64List")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}

{{- range .Attributes}}
{{- $cname := toGoName .TfName}}
{{- if eq .Type "List"}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{$cname}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if or (eq .Type "StringList") (eq .Type "Int64List")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}

{{- range .Attributes}}
{{- $cname := toGoName .TfName}}
{{- if eq .Type "List"}}
{{- range .Attributes}}
{{- if eq .Type "List"}}
type {{$name}}{{$cname}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if or (eq .Type "StringList") (eq .Type "Int64List")}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}

func (data {{camelCase .Name}}) getPath() string {
{{- if hasId .Attributes}}
	return fmt.Sprintf("{{.Path}}"{{range .Attributes}}{{if or .Id .Reference}}, data.{{toGoName .TfName}}.Value{{.Type}}(){{end}}{{end}})
{{- else}}
	return "{{.Path}}"
{{- end}}
}

func (data {{camelCase .Name}}Data) getPath() string {
{{- if hasId .Attributes}}
	return fmt.Sprintf("{{.Path}}"{{range .Attributes}}{{if or .Id .Reference}}, data.{{toGoName .TfName}}.Value{{.Type}}(){{end}}{{end}})
{{- else}}
	return "{{.Path}}"
{{- end}}
}

func (data {{camelCase .Name}}) toBody(ctx context.Context) string {
	body := "{}"
	{{- range .Attributes}}
	{{- if and (not .Reference) (ne .Type "List")}}
	if !data.{{toGoName .TfName}}.IsNull() && !data.{{toGoName .TfName}}.IsUnknown() {
		{{- if eq .Type "Int64"}}
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", strconv.FormatInt(data.{{toGoName .TfName}}.ValueInt64(), 10))
		{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
		if data.{{toGoName .TfName}}.ValueBool() {
			body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", map[string]string{})
		}
		{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", data.{{toGoName .TfName}}.ValueBool())
		{{- else if eq .Type "String"}}
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", data.{{toGoName .TfName}}.ValueString())
		{{- else if eq .Type "StringList"}}
		var values []string
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", values)
		{{- else if eq .Type "Int64List"}}
		var values []int
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", values)
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- range .Attributes}}
	{{- if eq .Type "List"}}
	{{- $list := toJsonPath .YangName .XPath }}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", []interface{}{})
		for index, item := range data.{{toGoName .TfName}} {
			{{- range .Attributes}}
			{{- if (ne .Type "List")}}
			if !item.{{toGoName .TfName}}.IsNull() && !item.{{toGoName .TfName}}.IsUnknown() {
				{{- if eq .Type "Int64"}}
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", strconv.FormatInt(item.{{toGoName .TfName}}.ValueInt64(), 10))
				{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
				if item.{{toGoName .TfName}}.ValueBool() {
					body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", map[string]string{})
				}
				{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", item.{{toGoName .TfName}}.ValueBool())
				{{- else if eq .Type "String"}}
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", item.{{toGoName .TfName}}.ValueString())
				{{- else if eq .Type "StringList"}}
				var values []string
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", values)
				{{- else if eq .Type "Int64List"}}
				var values []int
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", values)
				{{- end}}
			}
			{{- end}}
			{{- end}}
			{{- range .Attributes}}
			{{- if eq .Type "List"}}
			{{- $clist := toJsonPath .YangName .XPath }}
			if len(item.{{toGoName .TfName}}) > 0 {
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", []interface{}{})
				for cindex, citem := range item.{{toGoName .TfName}} {
					{{- range .Attributes}}
					if !citem.{{toGoName .TfName}}.IsNull() && !citem.{{toGoName .TfName}}.IsUnknown() {
						{{- if eq .Type "Int64"}}
						body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{$clist}}"+"."+strconv.Itoa(cindex)+"."+"{{toJsonPath .YangName .XPath}}", strconv.FormatInt(citem.{{toGoName .TfName}}.ValueInt64(), 10))
						{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
						if citem.{{toGoName .TfName}}.ValueBool() {
							body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{$clist}}"+"."+strconv.Itoa(cindex)+"."+"{{toJsonPath .YangName .XPath}}", map[string]string{})
						}
						{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
						body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{$clist}}"+"."+strconv.Itoa(cindex)+"."+"{{toJsonPath .YangName .XPath}}", citem.{{toGoName .TfName}}.ValueBool())
						{{- else if eq .Type "String"}}
						body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{$clist}}"+"."+strconv.Itoa(cindex)+"."+"{{toJsonPath .YangName .XPath}}", citem.{{toGoName .TfName}}.ValueString())
						{{- else if eq .Type "StringList"}}
						var values []string
						citem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{$clist}}"+"."+strconv.Itoa(cindex)+"."+"{{toJsonPath .YangName .XPath}}", values)
						{{- else if eq .Type "Int64List"}}
						var values []int
						citem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{$clist}}"+"."+strconv.Itoa(cindex)+"."+"{{toJsonPath .YangName .XPath}}", values)
						{{- end}}
					}
					{{- end}}
				}
			}
			{{- end}}
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
	return body
}

func (data *{{camelCase .Name}}) updateFromBody(ctx context.Context, res []byte) {
	{{- range .Attributes}}
	{{- if and (not .Reference) (not .Id) (not .WriteOnly)}}
	{{- if eq .Type "Int64"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = types.Int64Value(value.Int())
	} else {
		data.{{toGoName .TfName}} = types.Int64Null()
	}
	{{- else if eq .Type "Bool"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); !data.{{toGoName .TfName}}.IsNull() {
		{{- if eq .TypeYangBool "boolean"}}
		if value.Exists() {
			data.{{toGoName .TfName}} = types.BoolValue(value.Bool())
		}
		{{- else}}
		if value.Exists() {
			data.{{toGoName .TfName}} = types.BoolValue(true)
		} else {
			data.{{toGoName .TfName}} = types.BoolValue(false)
		}
		{{- end}}
	} else {
		data.{{toGoName .TfName}} = types.BoolNull()
	}
	{{- else if eq .Type "String"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}value.Raw{{else}}value.String(){{end}})
	} else {
		data.{{toGoName .TfName}} = types.StringNull()
	}
	{{- else if eq .Type "StringList"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = helpers.GetStringList(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.StringType)
	}
	{{- else if eq .Type "Int64List"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = helpers.GetInt64List(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.Int64Type)
	}
	{{- else if eq .Type "List"}}
	{{- $list := (toGoName .TfName)}}
	{{- $listPath := (toJsonPath .YangName .XPath)}}
	for i := range data.{{$list}} {
		keys := [...]string{ {{range .Attributes}}{{if .Id}}"{{.YangName}}", {{end}}{{end}} }
		keyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

		var r gjson.Result
		gjson.GetBytes(res, "{{$listPath}}").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)

		{{- range .Attributes}}
		{{- if not .WriteOnly}}
		{{- if eq .Type "Int64"}}
		if value := r.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = types.Int64Value(value.Int())
		} else {
			data.{{$list}}[i].{{toGoName .TfName}} = types.Int64Null()
		}
		{{- else if eq .Type "Bool"}}
		if value := r.Get("{{toJsonPath .YangName .XPath}}"); !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			{{- if eq .TypeYangBool "boolean"}}
			if value.Exists() {
				data.{{$list}}[i].{{toGoName .TfName}} = types.BoolValue(value.Bool())
			}
			{{- else}}
			if value.Exists() {
				data.{{$list}}[i].{{toGoName .TfName}} = types.BoolValue(true)
			} else {
				data.{{$list}}[i].{{toGoName .TfName}} = types.BoolValue(false)
			}
			{{- end}}
		} else {
			data.{{$list}}[i].{{toGoName .TfName}} = types.BoolNull()
		}
		{{- else if eq .Type "String"}}
		if value := r.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}value.Raw{{else}}value.String(){{end}})
		} else {
			data.{{$list}}[i].{{toGoName .TfName}} = types.StringNull()
		}
		{{- else if eq .Type "StringList"}}
		if value := r.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = helpers.GetStringList(value.Array())
		} else {
			data.{{$list}}[i].{{toGoName .TfName}} = types.ListNull(types.StringType)
		}
		{{- else if eq .Type "Int64List"}}
		if value := r.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = helpers.GetInt64List(value.Array())
		} else {
			data.{{$list}}[i].{{toGoName .TfName}} = types.ListNull(types.Int64Type)
		}
		{{- else if eq .Type "List"}}
		{{- $clist := (toGoName .TfName)}}
		{{- $clistPath := (toJsonPath .YangName .XPath)}}
		for ci := range data.{{$list}}[i].{{$clist}} {
			keys := [...]string{ {{range .Attributes}}{{if .Id}}"{{.YangName}}", {{end}}{{end}} }
			keyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueBool()), {{else}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

			var cr gjson.Result
			r.Get("{{$clistPath}}").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)

			{{- range .Attributes}}
			{{- if not .WriteOnly}}
			{{- if eq .Type "Int64"}}
			if value := cr.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.Int64Value(value.Int())
			} else {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.Int64Null()
			}
			{{- else if eq .Type "Bool"}}
			if value := cr.Get("{{toJsonPath .YangName .XPath}}"); !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				{{- if eq .TypeYangBool "boolean"}}
				if value.Exists() {
					data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.BoolValue(value.Bool())
				}
				{{- else}}
				if value.Exists() {
					data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.BoolValue(true)
				} else {
					data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.BoolValue(false)
				}
				{{- end}}
			} else {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.BoolNull()
			}
			{{- else if eq .Type "String"}}
			if value := cr.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}value.Raw{{else}}value.String(){{end}})
			} else {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.StringNull()
			}
			{{- else if eq .Type "StringList"}}
			if value := cr.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = helpers.GetStringList(value.Array())
			} else {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.ListNull(types.StringType)
			}
			{{- else if eq .Type "Int64List"}}
			if value := cr.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = helpers.GetInt64List(value.Array())
			} else {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.ListNull(types.Int64Type)
			}
			{{- end}}
			{{- end}}
			{{- end}}
		}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res []byte) {
	{{- range .Attributes}}
	{{- $cname := toGoName .TfName}}
	{{- if and (not .Reference) (not .Id) (not .WriteOnly)}}
	{{- if eq .Type "Int64"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.Int64Value(value.Int())
	}
	{{- else if eq .Type "Bool"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		{{- if eq .TypeYangBool "boolean"}}
		data.{{toGoName .TfName}} = types.BoolValue(value.Bool())
		{{- else}}
		data.{{toGoName .TfName}} = types.BoolValue(true)
		{{- end}}
	} else {
		data.{{toGoName .TfName}} = types.BoolValue(false)
	}
	{{- else if eq .Type "String"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}value.Raw{{else}}value.String(){{end}})
	}
	{{- else if eq .Type "StringList"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = helpers.GetStringList(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.StringType)
	}
	{{- else if eq .Type "Int64List"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = helpers.GetInt64List(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.Int64Type)
	}
	{{- else if eq .Type "List"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- if not .WriteOnly}}
			{{- if eq .Type "Int64"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.Int64Value(cValue.Int())
			}
			{{- else if eq .Type "Bool"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				{{- if eq .TypeYangBool "boolean"}}
				item.{{toGoName .TfName}} = types.BoolValue(cValue.Bool())
				{{- else}}
				item.{{toGoName .TfName}} = types.BoolValue(true)
				{{- end}}
			} else {
				item.{{toGoName .TfName}} = types.BoolValue(false)
			}
			{{- else if eq .Type "String"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}cValue.Raw{{else}}cValue.String(){{end}})
			}
			{{- else if eq .Type "StringList"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = helpers.GetStringList(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.ListNull(types.StringType)
			}
			{{- else if eq .Type "Int64List"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = helpers.GetInt64List(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.ListNull(types.Int64Type)
			}
			{{- else if eq .Type "List"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if not .WriteOnly}}
					{{- if eq .Type "Int64"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.Int64Value(ccValue.Int())
					}
					{{- else if eq .Type "Bool"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						{{- if eq .TypeYangBool "boolean"}}
						cItem.{{toGoName .TfName}} = types.BoolValue(ccValue.Bool())
						{{- else}}
						cItem.{{toGoName .TfName}} = types.BoolValue(true)
						{{- end}}
					} else {
						cItem.{{toGoName .TfName}} = types.BoolValue(false)
					}
					{{- else if eq .Type "String"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}ccValue.Raw{{else}}ccValue.String(){{end}})
					}
					{{- else if eq .Type "StringList"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = helpers.GetStringList(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
					}
					{{- else if eq .Type "Int64List"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = helpers.GetInt64List(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.ListNull(types.Int64Type)
					}
					{{- end}}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}Data) fromBody(ctx context.Context, res []byte) {
	{{- range .Attributes}}
	{{- $cname := toGoName .TfName}}
	{{- if and (not .Reference) (not .Id) (not .WriteOnly)}}
	{{- if eq .Type "Int64"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.Int64Value(value.Int())
	}
	{{- else if eq .Type "Bool"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		{{- if eq .TypeYangBool "boolean"}}
		data.{{toGoName .TfName}} = types.BoolValue(value.Bool())
		{{- else}}
		data.{{toGoName .TfName}} = types.BoolValue(true)
		{{- end}}
	} else {
		data.{{toGoName .TfName}} = types.BoolValue(false)
	}
	{{- else if eq .Type "String"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}value.Raw{{else}}value.String(){{end}})
	}
	{{- else if eq .Type "StringList"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = helpers.GetStringList(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.StringType)
	}
	{{- else if eq .Type "Int64List"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = helpers.GetInt64List(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.Int64Type)
	}
	{{- else if eq .Type "List"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- if not .WriteOnly}}
			{{- if eq .Type "Int64"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.Int64Value(cValue.Int())
			}
			{{- else if eq .Type "Bool"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				{{- if eq .TypeYangBool "boolean"}}
				item.{{toGoName .TfName}} = types.BoolValue(cValue.Bool())
				{{- else}}
				item.{{toGoName .TfName}} = types.BoolValue(true)
				{{- end}}
			} else {
				item.{{toGoName .TfName}} = types.BoolValue(false)
			}
			{{- else if eq .Type "String"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}cValue.Raw{{else}}cValue.String(){{end}})
			}
			{{- else if eq .Type "StringList"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = helpers.GetStringList(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.ListNull(types.StringType)
			}
			{{- else if eq .Type "Int64List"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = helpers.GetInt64List(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.ListNull(types.Int64Type)
			}
			{{- else if eq .Type "List"}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if not .WriteOnly}}
					{{- if eq .Type "Int64"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.Int64Value(ccValue.Int())
					}
					{{- else if eq .Type "Bool"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						{{- if eq .TypeYangBool "boolean"}}
						cItem.{{toGoName .TfName}} = types.BoolValue(ccValue.Bool())
						{{- else}}
						cItem.{{toGoName .TfName}} = types.BoolValue(true)
						{{- end}}
					} else {
						cItem.{{toGoName .TfName}} = types.BoolValue(false)
					}
					{{- else if eq .Type "String"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.StringValue({{if .ReadRaw}}ccValue.Raw{{else}}ccValue.String(){{end}})
					}
					{{- else if eq .Type "StringList"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = helpers.GetStringList(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
					}
					{{- else if eq .Type "Int64List"}}
					if ccValue := cv.Get("{{toJsonPath .YangName .XPath}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = helpers.GetInt64List(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.ListNull(types.Int64Type)
					}
					{{- end}}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) getDeletedItems(ctx context.Context, state {{camelCase .Name}}) []string {
	deletedItems := make([]string, 0)
	{{- range .Attributes}}
	{{- if and (not .Reference) (not .Id) (ne .Type "List") (not .NoDelete)}}
	if !state.{{toGoName .TfName}}.IsNull() && data.{{toGoName .TfName}}.IsNull() {
		{{- if .DeleteParent}}
		deletedItems = append(deletedItems, fmt.Sprintf("%v/{{removeLastPathElement .YangName}}", state.getPath()))
		{{- else}}
		deletedItems = append(deletedItems, fmt.Sprintf("%v/{{.YangName}}", state.getPath()))
		{{- end}}
	}
	{{- else if eq .Type "List"}}
	{{- $yangName := .YangName}}
	for i := range state.{{toGoName .TfName}} {
		{{- $list := (toGoName .TfName)}}
		keys := [...]string{ {{range .Attributes}}{{if .Id}}"{{.YangName}}", {{end}}{{end}} }
		stateKeyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(state.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(state.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else}}state.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
		keyString := ""
		for ki := range keys {
			keyString += "["+keys[ki]+"="+stateKeyValues[ki]+"]"
		}
		
		emptyKeys := true
		{{- range .Attributes}}
		{{- if .Id}}
		if !reflect.ValueOf(state.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}()).IsZero() {
			emptyKeys = false
		}
		{{- end}}
		{{- end}}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.{{toGoName .TfName}} {
			found = true
			{{- range .Attributes}}
			{{- if .Id}}
			if state.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}() != data.{{$list}}[j].{{toGoName .TfName}}.Value{{.Type}}() {
				found = false
			} 
			{{- end}}
			{{- end}}
			if found {
				{{- range .Attributes}}
				{{- if and (not .Reference) (not .Id) (ne .Type "List") (not .NoDelete)}}
				if !state.{{$list}}[i].{{toGoName .TfName}}.IsNull() && data.{{$list}}[j].{{toGoName .TfName}}.IsNull() {
					{{- if .DeleteParent}}
					deletedItems = append(deletedItems, fmt.Sprintf("%v/{{$yangName}}%v/{{removeLastPathElement .YangName}}", state.getPath(), keyString))
					{{- else}}
					deletedItems = append(deletedItems, fmt.Sprintf("%v/{{$yangName}}%v/{{.YangName}}", state.getPath(), keyString))
					{{- end}}
				}
				{{- else if eq .Type "List"}}
				{{- $cYangName := .YangName}}
				for ci := range state.{{$list}}[i].{{toGoName .TfName}} {
					{{- $clist := (toGoName .TfName)}}
					ckeys := [...]string{ {{range .Attributes}}{{if .Id}}"{{.YangName}}", {{end}}{{end}} }
					cstateKeyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(state.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(state.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueBool()), {{else}}state.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
					ckeyString := ""
					for cki := range ckeys {
						ckeyString += "["+ckeys[cki]+"="+cstateKeyValues[cki]+"]"
					}
					
					cemptyKeys := true
					{{- range .Attributes}}
					{{- if .Id}}
					if !reflect.ValueOf(state.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}()).IsZero() {
						cemptyKeys = false
					}
					{{- end}}
					{{- end}}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.{{$list}}[j].{{toGoName .TfName}} {
						found = true
						{{- range .Attributes}}
						{{- if .Id}}
						if state.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}() != data.{{$list}}[j].{{$clist}}[cj].{{toGoName .TfName}}.Value{{.Type}}() {
							found = false
						} 
						{{- end}}
						{{- end}}
						if found {
							{{- range .Attributes}}
							{{- if and (not .Reference) (not .Id) (ne .Type "List") (not .NoDelete)}}
							if !state.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() && data.{{$list}}[j].{{$clist}}[cj].{{toGoName .TfName}}.IsNull() {
								{{- if .DeleteParent}}
								deletedItems = append(deletedItems, fmt.Sprintf("%v/{{$yangName}}%v/{{$cYangName}}%v/{{removeLastPathElement .YangName}}", state.getPath(), keyString, ckeyString))
								{{- else}}
								deletedItems = append(deletedItems, fmt.Sprintf("%v/{{$yangName}}%v/{{$cYangName}}%v/{{.YangName}}", state.getPath(), keyString, ckeyString))
								{{- end}}
							}
							{{- end}}
							{{- end}}
							break
						}
					}
					if !found {
						deletedItems = append(deletedItems, fmt.Sprintf("%v/{{$yangName}}%v/{{.YangName}}%v", state.getPath(), keyString, ckeyString))
					}
				}
				{{- end}}
				{{- end}}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/{{.YangName}}%v", state.getPath(), keyString))
		}
	}
	{{- end}}
	{{- end}}
	return deletedItems
}

func (data *{{camelCase .Name}}) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	{{- range .Attributes}}
	{{- if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
	if !data.{{toGoName .TfName}}.IsNull() && !data.{{toGoName .TfName}}.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/{{.YangName}}", data.getPath()))
	}
	{{- end}}
	{{- if eq .Type "List"}}
	{{- $yangName := .YangName}}
	for i := range data.{{toGoName .TfName}} {
		{{- $list := (toGoName .TfName)}}
		keys := [...]string{ {{range .Attributes}}{{if .Id}}"{{.YangName}}", {{end}}{{end}} }
		keyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
		keyString := ""
		for ki := range keys {
			keyString += "["+keys[ki]+"="+keyValues[ki]+"]"
		}
		{{- range .Attributes}}
		{{- if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
		if !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() && !data.{{$list}}[i].{{toGoName .TfName}}.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/{{$yangName}}%v/{{.YangName}}", data.getPath(), keyString))
		}
		{{- end}}
		{{- if eq .Type "List"}}
		{{- $cyangName := .YangName}}
		for ci := range data.{{$list}}[i].{{toGoName .TfName}} {
			{{- $clist := (toGoName .TfName)}}
			ckeys := [...]string{ {{range .Attributes}}{{if .Id}}"{{.YangName}}", {{end}}{{end}} }
			ckeyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueBool()), {{else}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
			ckeyString := ""
			for cki := range ckeys {
				ckeyString += "["+ckeys[cki]+"="+ckeyValues[cki]+"]"
			}
			{{- range .Attributes}}
			{{- if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
			if !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueBool() {
				emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/{{$yangName}}%v/{{.YangName}}/{{$cyangName}}%v/{{.YangName}}", data.getPath(), keyString, ckeyString))
			}
			{{- end}}
			{{- end}}
		}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	return emptyLeafsDelete
}

func (data *{{camelCase .Name}}) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	{{- range .Attributes}}
	{{- if and (not .Reference) (not .Id) (ne .Type "List") (not .NoDelete)}}
	if !data.{{toGoName .TfName}}.IsNull() {
		{{- if .DeleteParent}}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/{{removeLastPathElement .YangName}}", data.getPath()))
		{{- else}}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/{{.YangName}}", data.getPath()))
		{{- end}}
	}
	{{- else if and (eq .Type "List") (not .NoDelete)}}
	for i := range data.{{toGoName .TfName}} {
		{{- $list := (toGoName .TfName)}}
		keys := [...]string{ {{range .Attributes}}{{if .Id}}"{{.YangName}}", {{end}}{{end}} }
		keyValues := [...]string{ {{range .Attributes}}{{if .Id}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

		keyString := ""
		for ki := range keys {
			keyString += "["+keys[ki]+"="+keyValues[ki]+"]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/{{.YangName}}%v", data.getPath(), keyString))
	}
	{{- end}}
	{{- end}}
	return deletePaths
}
