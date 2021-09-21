// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;

#[derive(Clone, Copy)]
pub struct ImmutableCallOnChainResults {
    pub(crate) id: i32,
}

impl ImmutableCallOnChainResults {
    pub fn int_value(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct MutableCallOnChainResults {
    pub(crate) id: i32,
}

impl MutableCallOnChainResults {
    pub fn int_value(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetMintedSupplyResults {
    pub(crate) id: i32,
}

impl ImmutableGetMintedSupplyResults {
    pub fn minted_color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_RESULT_MINTED_COLOR))
    }

    pub fn minted_supply(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_MINTED_SUPPLY))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetMintedSupplyResults {
    pub(crate) id: i32,
}

impl MutableGetMintedSupplyResults {
    pub fn minted_color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_RESULT_MINTED_COLOR))
    }

    pub fn minted_supply(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_MINTED_SUPPLY))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableRunRecursionResults {
    pub(crate) id: i32,
}

impl ImmutableRunRecursionResults {
    pub fn int_value(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct MutableRunRecursionResults {
    pub(crate) id: i32,
}

impl MutableRunRecursionResults {
    pub fn int_value(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableTestChainOwnerIDFullResults {
    pub(crate) id: i32,
}

impl ImmutableTestChainOwnerIDFullResults {
    pub fn chain_owner_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, idx_map(IDX_RESULT_CHAIN_OWNER_ID))
    }
}

#[derive(Clone, Copy)]
pub struct MutableTestChainOwnerIDFullResults {
    pub(crate) id: i32,
}

impl MutableTestChainOwnerIDFullResults {
    pub fn chain_owner_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, idx_map(IDX_RESULT_CHAIN_OWNER_ID))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableFibonacciResults {
    pub(crate) id: i32,
}

impl ImmutableFibonacciResults {
    pub fn int_value(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct MutableFibonacciResults {
    pub(crate) id: i32,
}

impl MutableFibonacciResults {
    pub fn int_value(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_INT_VALUE))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetCounterResults {
    pub(crate) id: i32,
}

impl ImmutableGetCounterResults {
    pub fn counter(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_COUNTER))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetCounterResults {
    pub(crate) id: i32,
}

impl MutableGetCounterResults {
    pub fn counter(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_COUNTER))
    }
}

pub struct MapStringToImmutableInt64 {
    pub(crate) obj_id: i32,
}

impl MapStringToImmutableInt64 {
    pub fn get_int64(&self, key: &str) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.obj_id, key.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetIntResults {
    pub(crate) id: i32,
}

impl ImmutableGetIntResults {
    pub fn values(&self) -> MapStringToImmutableInt64 {
        MapStringToImmutableInt64 { obj_id: self.id }
    }
}

pub struct MapStringToMutableInt64 {
    pub(crate) obj_id: i32,
}

impl MapStringToMutableInt64 {
    pub fn clear(&self) {
        clear(self.obj_id)
    }

    pub fn get_int64(&self, key: &str) -> ScMutableInt64 {
        ScMutableInt64::new(self.obj_id, key.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetIntResults {
    pub(crate) id: i32,
}

impl MutableGetIntResults {
    pub fn values(&self) -> MapStringToMutableInt64 {
        MapStringToMutableInt64 { obj_id: self.id }
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableTestChainOwnerIDViewResults {
    pub(crate) id: i32,
}

impl ImmutableTestChainOwnerIDViewResults {
    pub fn chain_owner_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, idx_map(IDX_RESULT_CHAIN_OWNER_ID))
    }
}

#[derive(Clone, Copy)]
pub struct MutableTestChainOwnerIDViewResults {
    pub(crate) id: i32,
}

impl MutableTestChainOwnerIDViewResults {
    pub fn chain_owner_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, idx_map(IDX_RESULT_CHAIN_OWNER_ID))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableTestSandboxCallResults {
    pub(crate) id: i32,
}

impl ImmutableTestSandboxCallResults {
    pub fn sandbox_call(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_RESULT_SANDBOX_CALL))
    }
}

#[derive(Clone, Copy)]
pub struct MutableTestSandboxCallResults {
    pub(crate) id: i32,
}

impl MutableTestSandboxCallResults {
    pub fn sandbox_call(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_RESULT_SANDBOX_CALL))
    }
}
