// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]
#![allow(unused_imports)]

use inccounter::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;

mod consts;
mod contract;
mod keys;
mod params;
mod results;
mod state;
mod inccounter;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_CALL_INCREMENT, func_call_increment_thunk);
    exports.add_func(FUNC_CALL_INCREMENT_RECURSE5X, func_call_increment_recurse5x_thunk);
    exports.add_func(FUNC_ENDLESS_LOOP, func_endless_loop_thunk);
    exports.add_func(FUNC_INCREMENT, func_increment_thunk);
    exports.add_func(FUNC_INCREMENT_WITH_DELAY, func_increment_with_delay_thunk);
    exports.add_func(FUNC_INIT, func_init_thunk);
    exports.add_func(FUNC_LOCAL_STATE_INTERNAL_CALL, func_local_state_internal_call_thunk);
    exports.add_func(FUNC_LOCAL_STATE_POST, func_local_state_post_thunk);
    exports.add_func(FUNC_LOCAL_STATE_SANDBOX_CALL, func_local_state_sandbox_call_thunk);
    exports.add_func(FUNC_POST_INCREMENT, func_post_increment_thunk);
    exports.add_func(FUNC_REPEAT_MANY, func_repeat_many_thunk);
    exports.add_func(FUNC_TEST_LEB128, func_test_leb128_thunk);
    exports.add_func(FUNC_WHEN_MUST_INCREMENT, func_when_must_increment_thunk);
    exports.add_view(VIEW_GET_COUNTER, view_get_counter_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct CallIncrementContext {
    state: MutableIncCounterState,
}

fn func_call_increment_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcCallIncrement");
    let f = CallIncrementContext {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_call_increment(ctx, &f);
    ctx.log("inccounter.funcCallIncrement ok");
}

pub struct CallIncrementRecurse5xContext {
    state: MutableIncCounterState,
}

fn func_call_increment_recurse5x_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcCallIncrementRecurse5x");
    let f = CallIncrementRecurse5xContext {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_call_increment_recurse5x(ctx, &f);
    ctx.log("inccounter.funcCallIncrementRecurse5x ok");
}

pub struct EndlessLoopContext {
    state: MutableIncCounterState,
}

fn func_endless_loop_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcEndlessLoop");
    let f = EndlessLoopContext {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_endless_loop(ctx, &f);
    ctx.log("inccounter.funcEndlessLoop ok");
}

pub struct IncrementContext {
    state: MutableIncCounterState,
}

fn func_increment_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcIncrement");
    let f = IncrementContext {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_increment(ctx, &f);
    ctx.log("inccounter.funcIncrement ok");
}

pub struct IncrementWithDelayContext {
    params: ImmutableIncrementWithDelayParams,
    state:  MutableIncCounterState,
}

fn func_increment_with_delay_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcIncrementWithDelay");
    let f = IncrementWithDelayContext {
        params: ImmutableIncrementWithDelayParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.delay().exists(), "missing mandatory delay");
    func_increment_with_delay(ctx, &f);
    ctx.log("inccounter.funcIncrementWithDelay ok");
}

pub struct InitContext {
    params: ImmutableInitParams,
    state:  MutableIncCounterState,
}

fn func_init_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcInit");
    let f = InitContext {
        params: ImmutableInitParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_init(ctx, &f);
    ctx.log("inccounter.funcInit ok");
}

pub struct LocalStateInternalCallContext {
    state: MutableIncCounterState,
}

fn func_local_state_internal_call_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcLocalStateInternalCall");
    let f = LocalStateInternalCallContext {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_local_state_internal_call(ctx, &f);
    ctx.log("inccounter.funcLocalStateInternalCall ok");
}

pub struct LocalStatePostContext {
    state: MutableIncCounterState,
}

fn func_local_state_post_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcLocalStatePost");
    let f = LocalStatePostContext {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_local_state_post(ctx, &f);
    ctx.log("inccounter.funcLocalStatePost ok");
}

pub struct LocalStateSandboxCallContext {
    state: MutableIncCounterState,
}

fn func_local_state_sandbox_call_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcLocalStateSandboxCall");
    let f = LocalStateSandboxCallContext {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_local_state_sandbox_call(ctx, &f);
    ctx.log("inccounter.funcLocalStateSandboxCall ok");
}

pub struct PostIncrementContext {
    state: MutableIncCounterState,
}

fn func_post_increment_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcPostIncrement");
    let f = PostIncrementContext {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_post_increment(ctx, &f);
    ctx.log("inccounter.funcPostIncrement ok");
}

pub struct RepeatManyContext {
    params: ImmutableRepeatManyParams,
    state:  MutableIncCounterState,
}

fn func_repeat_many_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcRepeatMany");
    let f = RepeatManyContext {
        params: ImmutableRepeatManyParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_repeat_many(ctx, &f);
    ctx.log("inccounter.funcRepeatMany ok");
}

pub struct TestLeb128Context {
    state: MutableIncCounterState,
}

fn func_test_leb128_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcTestLeb128");
    let f = TestLeb128Context {
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_test_leb128(ctx, &f);
    ctx.log("inccounter.funcTestLeb128 ok");
}

pub struct WhenMustIncrementContext {
    params: ImmutableWhenMustIncrementParams,
    state:  MutableIncCounterState,
}

fn func_when_must_increment_thunk(ctx: &ScFuncContext) {
    ctx.log("inccounter.funcWhenMustIncrement");
    let f = WhenMustIncrementContext {
        params: ImmutableWhenMustIncrementParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    func_when_must_increment(ctx, &f);
    ctx.log("inccounter.funcWhenMustIncrement ok");
}

pub struct GetCounterContext {
    results: MutableGetCounterResults,
    state:   ImmutableIncCounterState,
}

fn view_get_counter_thunk(ctx: &ScViewContext) {
    ctx.log("inccounter.viewGetCounter");
    let f = GetCounterContext {
        results: MutableGetCounterResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableIncCounterState {
            id: OBJ_ID_STATE,
        },
    };
    view_get_counter(ctx, &f);
    ctx.log("inccounter.viewGetCounter ok");
}

// @formatter:on
