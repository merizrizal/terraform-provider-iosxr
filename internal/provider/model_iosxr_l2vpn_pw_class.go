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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type L2VPNPWClass struct {
	Device                                                 types.String `tfsdk:"device"`
	Id                                                     types.String `tfsdk:"id"`
	DeleteMode                                             types.String `tfsdk:"delete_mode"`
	Name                                                   types.String `tfsdk:"name"`
	EncapsulationMpls                                      types.Bool   `tfsdk:"encapsulation_mpls"`
	EncapsulationMplsTransportModeEthernet                 types.Bool   `tfsdk:"encapsulation_mpls_transport_mode_ethernet"`
	EncapsulationMplsTransportModeVlan                     types.Bool   `tfsdk:"encapsulation_mpls_transport_mode_vlan"`
	EncapsulationMplsTransportModePassthrough              types.Bool   `tfsdk:"encapsulation_mpls_transport_mode_passthrough"`
	EncapsulationMplsLoadBalancingPwLabel                  types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_pw_label"`
	EncapsulationMplsLoadBalancingFlowLabelTransmit        types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_transmit"`
	EncapsulationMplsLoadBalancingFlowLabelTransmitStatic  types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_transmit_static"`
	EncapsulationMplsLoadBalancingFlowLabelReceive         types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_receive"`
	EncapsulationMplsLoadBalancingFlowLabelReceiveStatic   types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_receive_static"`
	EncapsulationMplsLoadBalancingFlowLabelBoth            types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_both"`
	EncapsulationMplsLoadBalancingFlowLabelBothStatic      types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_both_static"`
	EncapsulationMplsLoadBalancingFlowLabelCodeOne7        types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_code_one7"`
	EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_code_one7_disable"`
}

type L2VPNPWClassData struct {
	Device                                                 types.String `tfsdk:"device"`
	Id                                                     types.String `tfsdk:"id"`
	Name                                                   types.String `tfsdk:"name"`
	EncapsulationMpls                                      types.Bool   `tfsdk:"encapsulation_mpls"`
	EncapsulationMplsTransportModeEthernet                 types.Bool   `tfsdk:"encapsulation_mpls_transport_mode_ethernet"`
	EncapsulationMplsTransportModeVlan                     types.Bool   `tfsdk:"encapsulation_mpls_transport_mode_vlan"`
	EncapsulationMplsTransportModePassthrough              types.Bool   `tfsdk:"encapsulation_mpls_transport_mode_passthrough"`
	EncapsulationMplsLoadBalancingPwLabel                  types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_pw_label"`
	EncapsulationMplsLoadBalancingFlowLabelTransmit        types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_transmit"`
	EncapsulationMplsLoadBalancingFlowLabelTransmitStatic  types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_transmit_static"`
	EncapsulationMplsLoadBalancingFlowLabelReceive         types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_receive"`
	EncapsulationMplsLoadBalancingFlowLabelReceiveStatic   types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_receive_static"`
	EncapsulationMplsLoadBalancingFlowLabelBoth            types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_both"`
	EncapsulationMplsLoadBalancingFlowLabelBothStatic      types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_both_static"`
	EncapsulationMplsLoadBalancingFlowLabelCodeOne7        types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_code_one7"`
	EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable types.Bool   `tfsdk:"encapsulation_mpls_load_balancing_flow_label_code_one7_disable"`
}

func (data L2VPNPWClass) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/pw-classes/pw-class[pw-class-name=%s]", data.Name.ValueString())
}

func (data L2VPNPWClassData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/pw-classes/pw-class[pw-class-name=%s]", data.Name.ValueString())
}

func (data L2VPNPWClass) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "pw-class-name", data.Name.ValueString())
	}
	if !data.EncapsulationMpls.IsNull() && !data.EncapsulationMpls.IsUnknown() {
		if data.EncapsulationMpls.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls", map[string]string{})
		}
	}
	if !data.EncapsulationMplsTransportModeEthernet.IsNull() && !data.EncapsulationMplsTransportModeEthernet.IsUnknown() {
		if data.EncapsulationMplsTransportModeEthernet.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.transport-mode.ethernet", map[string]string{})
		}
	}
	if !data.EncapsulationMplsTransportModeVlan.IsNull() && !data.EncapsulationMplsTransportModeVlan.IsUnknown() {
		if data.EncapsulationMplsTransportModeVlan.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.transport-mode.vlan", map[string]string{})
		}
	}
	if !data.EncapsulationMplsTransportModePassthrough.IsNull() && !data.EncapsulationMplsTransportModePassthrough.IsUnknown() {
		if data.EncapsulationMplsTransportModePassthrough.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.transport-mode.passthrough", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingPwLabel.IsNull() && !data.EncapsulationMplsLoadBalancingPwLabel.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingPwLabel.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.pw-label", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelTransmit.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelTransmit.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingFlowLabelTransmit.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.flow-label.transmit", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.flow-label.transmit.static", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelReceive.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelReceive.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingFlowLabelReceive.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.flow-label.receive", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.flow-label.receive.static", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelBoth.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelBoth.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingFlowLabelBoth.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.flow-label.both", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelBothStatic.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelBothStatic.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingFlowLabelBothStatic.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.flow-label.both.static", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.flow-label.code.one7", map[string]string{})
		}
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.IsUnknown() {
		if data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation.mpls.load-balancing.flow-label.code.one7.disable", map[string]string{})
		}
	}
	return body
}

func (data *L2VPNPWClass) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "encapsulation.mpls"); !data.EncapsulationMpls.IsNull() {
		if value.Exists() {
			data.EncapsulationMpls = types.BoolValue(true)
		} else {
			data.EncapsulationMpls = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMpls = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.transport-mode.ethernet"); !data.EncapsulationMplsTransportModeEthernet.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsTransportModeEthernet = types.BoolValue(true)
		} else {
			data.EncapsulationMplsTransportModeEthernet = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsTransportModeEthernet = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.transport-mode.vlan"); !data.EncapsulationMplsTransportModeVlan.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsTransportModeVlan = types.BoolValue(true)
		} else {
			data.EncapsulationMplsTransportModeVlan = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsTransportModeVlan = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.transport-mode.passthrough"); !data.EncapsulationMplsTransportModePassthrough.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsTransportModePassthrough = types.BoolValue(true)
		} else {
			data.EncapsulationMplsTransportModePassthrough = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsTransportModePassthrough = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.pw-label"); !data.EncapsulationMplsLoadBalancingPwLabel.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingPwLabel = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingPwLabel = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingPwLabel = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.transmit"); !data.EncapsulationMplsLoadBalancingFlowLabelTransmit.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingFlowLabelTransmit = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingFlowLabelTransmit = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelTransmit = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.transmit.static"); !data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.receive"); !data.EncapsulationMplsLoadBalancingFlowLabelReceive.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingFlowLabelReceive = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingFlowLabelReceive = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelReceive = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.receive.static"); !data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.both"); !data.EncapsulationMplsLoadBalancingFlowLabelBoth.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingFlowLabelBoth = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingFlowLabelBoth = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelBoth = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.both.static"); !data.EncapsulationMplsLoadBalancingFlowLabelBothStatic.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingFlowLabelBothStatic = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingFlowLabelBothStatic = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelBothStatic = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.code.one7"); !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7 = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7 = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7 = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.code.one7.disable"); !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.IsNull() {
		if value.Exists() {
			data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable = types.BoolValue(true)
		} else {
			data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable = types.BoolValue(false)
		}
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable = types.BoolNull()
	}
}

func (data *L2VPNPWClassData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "encapsulation.mpls"); value.Exists() {
		data.EncapsulationMpls = types.BoolValue(true)
	} else {
		data.EncapsulationMpls = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.transport-mode.ethernet"); value.Exists() {
		data.EncapsulationMplsTransportModeEthernet = types.BoolValue(true)
	} else {
		data.EncapsulationMplsTransportModeEthernet = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.transport-mode.vlan"); value.Exists() {
		data.EncapsulationMplsTransportModeVlan = types.BoolValue(true)
	} else {
		data.EncapsulationMplsTransportModeVlan = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.transport-mode.passthrough"); value.Exists() {
		data.EncapsulationMplsTransportModePassthrough = types.BoolValue(true)
	} else {
		data.EncapsulationMplsTransportModePassthrough = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.pw-label"); value.Exists() {
		data.EncapsulationMplsLoadBalancingPwLabel = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingPwLabel = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.transmit"); value.Exists() {
		data.EncapsulationMplsLoadBalancingFlowLabelTransmit = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelTransmit = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.transmit.static"); value.Exists() {
		data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.receive"); value.Exists() {
		data.EncapsulationMplsLoadBalancingFlowLabelReceive = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelReceive = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.receive.static"); value.Exists() {
		data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.both"); value.Exists() {
		data.EncapsulationMplsLoadBalancingFlowLabelBoth = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelBoth = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.both.static"); value.Exists() {
		data.EncapsulationMplsLoadBalancingFlowLabelBothStatic = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelBothStatic = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.code.one7"); value.Exists() {
		data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7 = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7 = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation.mpls.load-balancing.flow-label.code.one7.disable"); value.Exists() {
		data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable = types.BoolValue(true)
	} else {
		data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable = types.BoolValue(false)
	}
}

func (data *L2VPNPWClass) getDeletedItems(ctx context.Context, state L2VPNPWClass) []string {
	deletedItems := make([]string, 0)
	if !state.EncapsulationMpls.IsNull() && data.EncapsulationMpls.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls", state.getPath()))
	}
	if !state.EncapsulationMplsTransportModeEthernet.IsNull() && data.EncapsulationMplsTransportModeEthernet.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/ethernet", state.getPath()))
	}
	if !state.EncapsulationMplsTransportModeVlan.IsNull() && data.EncapsulationMplsTransportModeVlan.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/vlan", state.getPath()))
	}
	if !state.EncapsulationMplsTransportModePassthrough.IsNull() && data.EncapsulationMplsTransportModePassthrough.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/passthrough", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingPwLabel.IsNull() && data.EncapsulationMplsLoadBalancingPwLabel.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/pw-label", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingFlowLabelTransmit.IsNull() && data.EncapsulationMplsLoadBalancingFlowLabelTransmit.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/transmit", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.IsNull() && data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/transmit/static", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingFlowLabelReceive.IsNull() && data.EncapsulationMplsLoadBalancingFlowLabelReceive.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/receive", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.IsNull() && data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/receive/static", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingFlowLabelBoth.IsNull() && data.EncapsulationMplsLoadBalancingFlowLabelBoth.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/both", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingFlowLabelBothStatic.IsNull() && data.EncapsulationMplsLoadBalancingFlowLabelBothStatic.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/both/static", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.IsNull() && data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/code/one7", state.getPath()))
	}
	if !state.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.IsNull() && data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/code/one7/disable", state.getPath()))
	}
	return deletedItems
}

func (data *L2VPNPWClass) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.EncapsulationMpls.IsNull() && !data.EncapsulationMpls.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls", data.getPath()))
	}
	if !data.EncapsulationMplsTransportModeEthernet.IsNull() && !data.EncapsulationMplsTransportModeEthernet.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/ethernet", data.getPath()))
	}
	if !data.EncapsulationMplsTransportModeVlan.IsNull() && !data.EncapsulationMplsTransportModeVlan.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/vlan", data.getPath()))
	}
	if !data.EncapsulationMplsTransportModePassthrough.IsNull() && !data.EncapsulationMplsTransportModePassthrough.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/passthrough", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingPwLabel.IsNull() && !data.EncapsulationMplsLoadBalancingPwLabel.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/pw-label", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelTransmit.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelTransmit.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/transmit", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/transmit/static", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelReceive.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelReceive.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/receive", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/receive/static", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelBoth.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelBoth.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/both", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelBothStatic.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelBothStatic.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/both/static", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/code/one7", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.IsNull() && !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/code/one7/disable", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *L2VPNPWClass) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.EncapsulationMpls.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls", data.getPath()))
	}
	if !data.EncapsulationMplsTransportModeEthernet.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/ethernet", data.getPath()))
	}
	if !data.EncapsulationMplsTransportModeVlan.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/vlan", data.getPath()))
	}
	if !data.EncapsulationMplsTransportModePassthrough.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/transport-mode/passthrough", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingPwLabel.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/pw-label", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelTransmit.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/transmit", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelTransmitStatic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/transmit/static", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelReceive.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/receive", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelReceiveStatic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/receive/static", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelBoth.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/both", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelBothStatic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/both/static", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/code/one7", data.getPath()))
	}
	if !data.EncapsulationMplsLoadBalancingFlowLabelCodeOne7Disable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation/mpls/load-balancing/flow-label/code/one7/disable", data.getPath()))
	}
	return deletePaths
}
