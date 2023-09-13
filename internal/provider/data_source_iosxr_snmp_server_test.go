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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrSNMPServer(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "location", "My location"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "contact", "My contact"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "rf", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "bfd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "ntp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "ethernet_oam_events", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "copy_complete", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "traps_snmp_linkup", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "traps_snmp_linkdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "power", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "config", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "entity", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "system", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "bridgemib", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "entity_state_operstatus", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "entity_redundancy_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "trap_source_both", "Loopback10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "l2vpn_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "l2vpn_vc_up", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "l2vpn_vc_down", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "sensor", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "fru_ctrl", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "isis_authentication_failure", "enable"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "bgp_cbgp2_updown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "bgp_bgp4_mib_updown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "users.0.user_name", "USER1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "users.0.group_name", "GROUP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "users.0.v3_auth_md5_encryption_aes", "073C05626E2A4841141D"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "groups.0.group_name", "GROUP12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "groups.0.v3_priv", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "groups.0.v3_read", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "groups.0.v3_write", "VIEW2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "groups.0.v3_context", "CONTEXT1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "groups.0.v3_notify", "VIEW3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "groups.0.v3_ipv4", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "groups.0.v3_ipv6", "ACL2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "communities.0.community", "COMMUNITY1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "communities.0.view", "VIEW1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "communities.0.ro", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "communities.0.rw", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "communities.0.sdrowner", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "communities.0.systemowner", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "communities.0.ipv4", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_snmp_server.test", "communities.0.ipv6", "ACL2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrSNMPServerConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrSNMPServerConfig() string {
	config := `resource "iosxr_snmp_server" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	location = "My location"` + "\n"
	config += `	contact = "My contact"` + "\n"
	config += `	rf = true` + "\n"
	config += `	bfd = true` + "\n"
	config += `	ntp = true` + "\n"
	config += `	ethernet_oam_events = true` + "\n"
	config += `	copy_complete = true` + "\n"
	config += `	traps_snmp_linkup = true` + "\n"
	config += `	traps_snmp_linkdown = true` + "\n"
	config += `	power = true` + "\n"
	config += `	config = true` + "\n"
	config += `	entity = true` + "\n"
	config += `	system = true` + "\n"
	config += `	bridgemib = true` + "\n"
	config += `	entity_state_operstatus = true` + "\n"
	config += `	entity_redundancy_all = true` + "\n"
	config += `	trap_source_both = "Loopback10"` + "\n"
	config += `	l2vpn_all = true` + "\n"
	config += `	l2vpn_vc_up = true` + "\n"
	config += `	l2vpn_vc_down = true` + "\n"
	config += `	sensor = true` + "\n"
	config += `	fru_ctrl = true` + "\n"
	config += `	isis_authentication_failure = "enable"` + "\n"
	config += `	bgp_cbgp2_updown = true` + "\n"
	config += `	bgp_bgp4_mib_updown = true` + "\n"
	config += `	users = [{` + "\n"
	config += `		user_name = "USER1"` + "\n"
	config += `		group_name = "GROUP1"` + "\n"
	config += `		v3_auth_md5_encryption_aes = "073C05626E2A4841141D"` + "\n"
	config += `	}]` + "\n"
	config += `	groups = [{` + "\n"
	config += `		group_name = "GROUP12"` + "\n"
	config += `		v3_priv = true` + "\n"
	config += `		v3_read = "VIEW1"` + "\n"
	config += `		v3_write = "VIEW2"` + "\n"
	config += `		v3_context = "CONTEXT1"` + "\n"
	config += `		v3_notify = "VIEW3"` + "\n"
	config += `		v3_ipv4 = "ACL1"` + "\n"
	config += `		v3_ipv6 = "ACL2"` + "\n"
	config += `	}]` + "\n"
	config += `	communities = [{` + "\n"
	config += `		community = "COMMUNITY1"` + "\n"
	config += `		view = "VIEW1"` + "\n"
	config += `		ro = true` + "\n"
	config += `		rw = false` + "\n"
	config += `		sdrowner = false` + "\n"
	config += `		systemowner = true` + "\n"
	config += `		ipv4 = "ACL1"` + "\n"
	config += `		ipv6 = "ACL2"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_snmp_server" "test" {
			depends_on = [iosxr_snmp_server.test]
		}
	`
	return config
}
