// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;

pub struct Erc721Events {
}

impl Erc721Events {

	pub fn approval(&self, approved: &ScAgentID, owner: &ScAgentID, token_id: &ScHash) {
		let mut encoder = EventEncoder::new("erc721.approval");
		encoder.agent_id(&approved);
		encoder.agent_id(&owner);
		encoder.hash(&token_id);
		encoder.emit();
	}

	pub fn approval_for_all(&self, approval: bool, operator: &ScAgentID, owner: &ScAgentID) {
		let mut encoder = EventEncoder::new("erc721.approvalForAll");
		encoder.bool(approval);
		encoder.agent_id(&operator);
		encoder.agent_id(&owner);
		encoder.emit();
	}

	pub fn init(&self, name: &str, symbol: &str) {
		let mut encoder = EventEncoder::new("erc721.init");
		encoder.string(&name);
		encoder.string(&symbol);
		encoder.emit();
	}

	pub fn mint(&self, balance: u64, owner: &ScAgentID, token_id: &ScHash) {
		let mut encoder = EventEncoder::new("erc721.mint");
		encoder.uint64(balance);
		encoder.agent_id(&owner);
		encoder.hash(&token_id);
		encoder.emit();
	}

	pub fn transfer(&self, from: &ScAgentID, to: &ScAgentID, token_id: &ScHash) {
		let mut encoder = EventEncoder::new("erc721.transfer");
		encoder.agent_id(&from);
		encoder.agent_id(&to);
		encoder.hash(&token_id);
		encoder.emit();
	}
}
