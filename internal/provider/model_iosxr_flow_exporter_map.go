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
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type FlowExporterMap struct {
	Device                              types.String `tfsdk:"device"`
	Id                                  types.String `tfsdk:"id"`
	Name                                types.String `tfsdk:"name"`
	DestinationIpv4Address              types.String `tfsdk:"destination_ipv4_address"`
	DestinationIpv6Address              types.String `tfsdk:"destination_ipv6_address"`
	DestinationVrf                      types.String `tfsdk:"destination_vrf"`
	Source                              types.String `tfsdk:"source"`
	Dscp                                types.Int64  `tfsdk:"dscp"`
	PacketLength                        types.Int64  `tfsdk:"packet_length"`
	TransportUdp                        types.Int64  `tfsdk:"transport_udp"`
	DfbitSet                            types.Bool   `tfsdk:"dfbit_set"`
	VersionExportFormat                 types.String `tfsdk:"version_export_format"`
	VersionTemplateDataTimeout          types.Int64  `tfsdk:"version_template_data_timeout"`
	VersionTemplateOptionsTimeout       types.Int64  `tfsdk:"version_template_options_timeout"`
	VersionTemplateTimeout              types.Int64  `tfsdk:"version_template_timeout"`
	VersionOptionsInterfaceTableTimeout types.Int64  `tfsdk:"version_options_interface_table_timeout"`
	VersionOptionsSamplerTableTimeout   types.Int64  `tfsdk:"version_options_sampler_table_timeout"`
	VersionOptionsClassTableTimeout     types.Int64  `tfsdk:"version_options_class_table_timeout"`
	VersionOptionsVrfTableTimeout       types.Int64  `tfsdk:"version_options_vrf_table_timeout"`
}

type FlowExporterMapData struct {
	Device                              types.String `tfsdk:"device"`
	Id                                  types.String `tfsdk:"id"`
	Name                                types.String `tfsdk:"name"`
	DestinationIpv4Address              types.String `tfsdk:"destination_ipv4_address"`
	DestinationIpv6Address              types.String `tfsdk:"destination_ipv6_address"`
	DestinationVrf                      types.String `tfsdk:"destination_vrf"`
	Source                              types.String `tfsdk:"source"`
	Dscp                                types.Int64  `tfsdk:"dscp"`
	PacketLength                        types.Int64  `tfsdk:"packet_length"`
	TransportUdp                        types.Int64  `tfsdk:"transport_udp"`
	DfbitSet                            types.Bool   `tfsdk:"dfbit_set"`
	VersionExportFormat                 types.String `tfsdk:"version_export_format"`
	VersionTemplateDataTimeout          types.Int64  `tfsdk:"version_template_data_timeout"`
	VersionTemplateOptionsTimeout       types.Int64  `tfsdk:"version_template_options_timeout"`
	VersionTemplateTimeout              types.Int64  `tfsdk:"version_template_timeout"`
	VersionOptionsInterfaceTableTimeout types.Int64  `tfsdk:"version_options_interface_table_timeout"`
	VersionOptionsSamplerTableTimeout   types.Int64  `tfsdk:"version_options_sampler_table_timeout"`
	VersionOptionsClassTableTimeout     types.Int64  `tfsdk:"version_options_class_table_timeout"`
	VersionOptionsVrfTableTimeout       types.Int64  `tfsdk:"version_options_vrf_table_timeout"`
}

func (data FlowExporterMap) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-flow-cfg:/flow/exporter-maps/exporter-map[exporter-map-name=%s]", data.Name.ValueString())
}

func (data FlowExporterMapData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-flow-cfg:/flow/exporter-maps/exporter-map[exporter-map-name=%s]", data.Name.ValueString())
}

func (data FlowExporterMap) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "exporter-map-name", data.Name.ValueString())
	}
	if !data.DestinationIpv4Address.IsNull() && !data.DestinationIpv4Address.IsUnknown() {
		body, _ = sjson.Set(body, "destination.ipv4-address", data.DestinationIpv4Address.ValueString())
	}
	if !data.DestinationIpv6Address.IsNull() && !data.DestinationIpv6Address.IsUnknown() {
		body, _ = sjson.Set(body, "destination.ipv6-address", data.DestinationIpv6Address.ValueString())
	}
	if !data.DestinationVrf.IsNull() && !data.DestinationVrf.IsUnknown() {
		body, _ = sjson.Set(body, "destination.vrf", data.DestinationVrf.ValueString())
	}
	if !data.Source.IsNull() && !data.Source.IsUnknown() {
		body, _ = sjson.Set(body, "source", data.Source.ValueString())
	}
	if !data.Dscp.IsNull() && !data.Dscp.IsUnknown() {
		body, _ = sjson.Set(body, "dscp", strconv.FormatInt(data.Dscp.ValueInt64(), 10))
	}
	if !data.PacketLength.IsNull() && !data.PacketLength.IsUnknown() {
		body, _ = sjson.Set(body, "packet-length", strconv.FormatInt(data.PacketLength.ValueInt64(), 10))
	}
	if !data.TransportUdp.IsNull() && !data.TransportUdp.IsUnknown() {
		body, _ = sjson.Set(body, "transport.udp", strconv.FormatInt(data.TransportUdp.ValueInt64(), 10))
	}
	if !data.DfbitSet.IsNull() && !data.DfbitSet.IsUnknown() {
		if data.DfbitSet.ValueBool() {
			body, _ = sjson.Set(body, "dfbit.set", map[string]string{})
		}
	}
	if !data.VersionExportFormat.IsNull() && !data.VersionExportFormat.IsUnknown() {
		body, _ = sjson.Set(body, "version.export-format", data.VersionExportFormat.ValueString())
	}
	if !data.VersionTemplateDataTimeout.IsNull() && !data.VersionTemplateDataTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "version.template.data.timeout", strconv.FormatInt(data.VersionTemplateDataTimeout.ValueInt64(), 10))
	}
	if !data.VersionTemplateOptionsTimeout.IsNull() && !data.VersionTemplateOptionsTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "version.template.options.timeout", strconv.FormatInt(data.VersionTemplateOptionsTimeout.ValueInt64(), 10))
	}
	if !data.VersionTemplateTimeout.IsNull() && !data.VersionTemplateTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "version.template.timeout", strconv.FormatInt(data.VersionTemplateTimeout.ValueInt64(), 10))
	}
	if !data.VersionOptionsInterfaceTableTimeout.IsNull() && !data.VersionOptionsInterfaceTableTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "version.options.interface-table.timeout", strconv.FormatInt(data.VersionOptionsInterfaceTableTimeout.ValueInt64(), 10))
	}
	if !data.VersionOptionsSamplerTableTimeout.IsNull() && !data.VersionOptionsSamplerTableTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "version.options.sampler-table.timeout", strconv.FormatInt(data.VersionOptionsSamplerTableTimeout.ValueInt64(), 10))
	}
	if !data.VersionOptionsClassTableTimeout.IsNull() && !data.VersionOptionsClassTableTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "version.options.class-table.timeout", strconv.FormatInt(data.VersionOptionsClassTableTimeout.ValueInt64(), 10))
	}
	if !data.VersionOptionsVrfTableTimeout.IsNull() && !data.VersionOptionsVrfTableTimeout.IsUnknown() {
		body, _ = sjson.Set(body, "version.options.vrf-table.timeout", strconv.FormatInt(data.VersionOptionsVrfTableTimeout.ValueInt64(), 10))
	}
	return body
}

func (data *FlowExporterMap) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "destination.ipv4-address"); value.Exists() && !data.DestinationIpv4Address.IsNull() {
		data.DestinationIpv4Address = types.StringValue(value.String())
	} else {
		data.DestinationIpv4Address = types.StringNull()
	}
	if value := gjson.GetBytes(res, "destination.ipv6-address"); value.Exists() && !data.DestinationIpv6Address.IsNull() {
		data.DestinationIpv6Address = types.StringValue(value.String())
	} else {
		data.DestinationIpv6Address = types.StringNull()
	}
	if value := gjson.GetBytes(res, "destination.vrf"); value.Exists() && !data.DestinationVrf.IsNull() {
		data.DestinationVrf = types.StringValue(value.String())
	} else {
		data.DestinationVrf = types.StringNull()
	}
	if value := gjson.GetBytes(res, "source"); value.Exists() && !data.Source.IsNull() {
		data.Source = types.StringValue(value.String())
	} else {
		data.Source = types.StringNull()
	}
	if value := gjson.GetBytes(res, "dscp"); value.Exists() && !data.Dscp.IsNull() {
		data.Dscp = types.Int64Value(value.Int())
	} else {
		data.Dscp = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "packet-length"); value.Exists() && !data.PacketLength.IsNull() {
		data.PacketLength = types.Int64Value(value.Int())
	} else {
		data.PacketLength = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "transport.udp"); value.Exists() && !data.TransportUdp.IsNull() {
		data.TransportUdp = types.Int64Value(value.Int())
	} else {
		data.TransportUdp = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "dfbit.set"); !data.DfbitSet.IsNull() {
		if value.Exists() {
			data.DfbitSet = types.BoolValue(true)
		} else {
			data.DfbitSet = types.BoolValue(false)
		}
	} else {
		data.DfbitSet = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "version.export-format"); value.Exists() && !data.VersionExportFormat.IsNull() {
		data.VersionExportFormat = types.StringValue(value.String())
	} else {
		data.VersionExportFormat = types.StringNull()
	}
	if value := gjson.GetBytes(res, "version.template.data.timeout"); value.Exists() && !data.VersionTemplateDataTimeout.IsNull() {
		data.VersionTemplateDataTimeout = types.Int64Value(value.Int())
	} else {
		data.VersionTemplateDataTimeout = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "version.template.options.timeout"); value.Exists() && !data.VersionTemplateOptionsTimeout.IsNull() {
		data.VersionTemplateOptionsTimeout = types.Int64Value(value.Int())
	} else {
		data.VersionTemplateOptionsTimeout = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "version.template.timeout"); value.Exists() && !data.VersionTemplateTimeout.IsNull() {
		data.VersionTemplateTimeout = types.Int64Value(value.Int())
	} else {
		data.VersionTemplateTimeout = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "version.options.interface-table.timeout"); value.Exists() && !data.VersionOptionsInterfaceTableTimeout.IsNull() {
		data.VersionOptionsInterfaceTableTimeout = types.Int64Value(value.Int())
	} else {
		data.VersionOptionsInterfaceTableTimeout = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "version.options.sampler-table.timeout"); value.Exists() && !data.VersionOptionsSamplerTableTimeout.IsNull() {
		data.VersionOptionsSamplerTableTimeout = types.Int64Value(value.Int())
	} else {
		data.VersionOptionsSamplerTableTimeout = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "version.options.class-table.timeout"); value.Exists() && !data.VersionOptionsClassTableTimeout.IsNull() {
		data.VersionOptionsClassTableTimeout = types.Int64Value(value.Int())
	} else {
		data.VersionOptionsClassTableTimeout = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "version.options.vrf-table.timeout"); value.Exists() && !data.VersionOptionsVrfTableTimeout.IsNull() {
		data.VersionOptionsVrfTableTimeout = types.Int64Value(value.Int())
	} else {
		data.VersionOptionsVrfTableTimeout = types.Int64Null()
	}
}

func (data *FlowExporterMap) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "destination.ipv4-address"); value.Exists() {
		data.DestinationIpv4Address = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "destination.ipv6-address"); value.Exists() {
		data.DestinationIpv6Address = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "destination.vrf"); value.Exists() {
		data.DestinationVrf = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "source"); value.Exists() {
		data.Source = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "dscp"); value.Exists() {
		data.Dscp = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "packet-length"); value.Exists() {
		data.PacketLength = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "transport.udp"); value.Exists() {
		data.TransportUdp = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "dfbit.set"); value.Exists() {
		data.DfbitSet = types.BoolValue(true)
	} else {
		data.DfbitSet = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "version.export-format"); value.Exists() {
		data.VersionExportFormat = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "version.template.data.timeout"); value.Exists() {
		data.VersionTemplateDataTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.template.options.timeout"); value.Exists() {
		data.VersionTemplateOptionsTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.template.timeout"); value.Exists() {
		data.VersionTemplateTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.options.interface-table.timeout"); value.Exists() {
		data.VersionOptionsInterfaceTableTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.options.sampler-table.timeout"); value.Exists() {
		data.VersionOptionsSamplerTableTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.options.class-table.timeout"); value.Exists() {
		data.VersionOptionsClassTableTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.options.vrf-table.timeout"); value.Exists() {
		data.VersionOptionsVrfTableTimeout = types.Int64Value(value.Int())
	}
}

func (data *FlowExporterMapData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "destination.ipv4-address"); value.Exists() {
		data.DestinationIpv4Address = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "destination.ipv6-address"); value.Exists() {
		data.DestinationIpv6Address = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "destination.vrf"); value.Exists() {
		data.DestinationVrf = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "source"); value.Exists() {
		data.Source = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "dscp"); value.Exists() {
		data.Dscp = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "packet-length"); value.Exists() {
		data.PacketLength = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "transport.udp"); value.Exists() {
		data.TransportUdp = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "dfbit.set"); value.Exists() {
		data.DfbitSet = types.BoolValue(true)
	} else {
		data.DfbitSet = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "version.export-format"); value.Exists() {
		data.VersionExportFormat = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "version.template.data.timeout"); value.Exists() {
		data.VersionTemplateDataTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.template.options.timeout"); value.Exists() {
		data.VersionTemplateOptionsTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.template.timeout"); value.Exists() {
		data.VersionTemplateTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.options.interface-table.timeout"); value.Exists() {
		data.VersionOptionsInterfaceTableTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.options.sampler-table.timeout"); value.Exists() {
		data.VersionOptionsSamplerTableTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.options.class-table.timeout"); value.Exists() {
		data.VersionOptionsClassTableTimeout = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "version.options.vrf-table.timeout"); value.Exists() {
		data.VersionOptionsVrfTableTimeout = types.Int64Value(value.Int())
	}
}

func (data *FlowExporterMap) getDeletedItems(ctx context.Context, state FlowExporterMap) []string {
	deletedItems := make([]string, 0)
	if !state.DestinationIpv4Address.IsNull() && data.DestinationIpv4Address.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/destination/ipv4-address", state.getPath()))
	}
	if !state.DestinationIpv6Address.IsNull() && data.DestinationIpv6Address.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/destination/ipv6-address", state.getPath()))
	}
	if !state.DestinationVrf.IsNull() && data.DestinationVrf.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/destination/vrf", state.getPath()))
	}
	if !state.Source.IsNull() && data.Source.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/source", state.getPath()))
	}
	if !state.Dscp.IsNull() && data.Dscp.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/dscp", state.getPath()))
	}
	if !state.PacketLength.IsNull() && data.PacketLength.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/packet-length", state.getPath()))
	}
	if !state.TransportUdp.IsNull() && data.TransportUdp.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/transport/udp", state.getPath()))
	}
	if !state.DfbitSet.IsNull() && data.DfbitSet.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/dfbit/set", state.getPath()))
	}
	if !state.VersionExportFormat.IsNull() && data.VersionExportFormat.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/version/export-format", state.getPath()))
	}
	if !state.VersionTemplateDataTimeout.IsNull() && data.VersionTemplateDataTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/version/template/data/timeout", state.getPath()))
	}
	if !state.VersionTemplateOptionsTimeout.IsNull() && data.VersionTemplateOptionsTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/version/template/options/timeout", state.getPath()))
	}
	if !state.VersionTemplateTimeout.IsNull() && data.VersionTemplateTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/version/template/timeout", state.getPath()))
	}
	if !state.VersionOptionsInterfaceTableTimeout.IsNull() && data.VersionOptionsInterfaceTableTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/version/options/interface-table/timeout", state.getPath()))
	}
	if !state.VersionOptionsSamplerTableTimeout.IsNull() && data.VersionOptionsSamplerTableTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/version/options/sampler-table/timeout", state.getPath()))
	}
	if !state.VersionOptionsClassTableTimeout.IsNull() && data.VersionOptionsClassTableTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/version/options/class-table/timeout", state.getPath()))
	}
	if !state.VersionOptionsVrfTableTimeout.IsNull() && data.VersionOptionsVrfTableTimeout.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/version/options/vrf-table/timeout", state.getPath()))
	}
	return deletedItems
}

func (data *FlowExporterMap) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.DfbitSet.IsNull() && !data.DfbitSet.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/dfbit/set", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *FlowExporterMap) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.DestinationIpv4Address.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/destination/ipv4-address", data.getPath()))
	}
	if !data.DestinationIpv6Address.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/destination/ipv6-address", data.getPath()))
	}
	if !data.DestinationVrf.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/destination/vrf", data.getPath()))
	}
	if !data.Source.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/source", data.getPath()))
	}
	if !data.Dscp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dscp", data.getPath()))
	}
	if !data.PacketLength.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/packet-length", data.getPath()))
	}
	if !data.TransportUdp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/transport/udp", data.getPath()))
	}
	if !data.DfbitSet.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/dfbit/set", data.getPath()))
	}
	if !data.VersionExportFormat.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version/export-format", data.getPath()))
	}
	if !data.VersionTemplateDataTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version/template/data/timeout", data.getPath()))
	}
	if !data.VersionTemplateOptionsTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version/template/options/timeout", data.getPath()))
	}
	if !data.VersionTemplateTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version/template/timeout", data.getPath()))
	}
	if !data.VersionOptionsInterfaceTableTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version/options/interface-table/timeout", data.getPath()))
	}
	if !data.VersionOptionsSamplerTableTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version/options/sampler-table/timeout", data.getPath()))
	}
	if !data.VersionOptionsClassTableTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version/options/class-table/timeout", data.getPath()))
	}
	if !data.VersionOptionsVrfTableTimeout.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version/options/vrf-table/timeout", data.getPath()))
	}
	return deletePaths
}
