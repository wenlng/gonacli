const goaddon = require('bindings')('goaddon');
// JS call API
module.exports = { 
  intSum32 : goaddon.int_sum32,
  intSum64 : goaddon.int_sum64,
  uintSum32 : goaddon.uint_sum32,
  compareInt : goaddon.compare_int,
  floatSum : goaddon.float_sum,
  doubleSum : goaddon.double_sum,
  formatStr : goaddon.format_str,
  emptyString : goaddon.empty_string,
  filterMap : goaddon.filter_map,
  countMap : goaddon.count_map,
  isMapType : goaddon.is_map_type,
  filterSlice : goaddon.filter_slice,
  countSlice : goaddon.count_slice,
  isSliceType : goaddon.is_slice_type,
  asyncCallbackSleep : goaddon.async_callback_sleep,
  asyncCallbackReStr : goaddon.async_callback_re_str,
  asyncCallbackReUintSum32 : goaddon.async_callback_re_uint_sum32,
  asyncCallbackReArr : goaddon.async_callback_re_arr,
  asyncCallbackReObject : goaddon.async_callback_re_object,
  asyncCallbackReCount : goaddon.async_callback_re_count,
  asyncCallbackReBool : goaddon.async_callback_re_bool,
  asyncCallbackMArg : goaddon.async_callback_m_arg
};
