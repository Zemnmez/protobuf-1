// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Code generated by generate-types. DO NOT EDIT.

package proto

import (
	"math"

	"github.com/golang/protobuf/v2/internal/encoding/wire"
	"github.com/golang/protobuf/v2/reflect/protoreflect"
)

// unmarshalScalar decodes a value of the given kind.
//
// Message values are decoded into a []byte which aliases the input data.
func (o UnmarshalOptions) unmarshalScalar(b []byte, wtyp wire.Type, num wire.Number, kind protoreflect.Kind) (val protoreflect.Value, n int, err error) {
	switch kind {
	case protoreflect.BoolKind:
		if wtyp != wire.VarintType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(wire.DecodeBool(v)), n, nil
	case protoreflect.EnumKind:
		if wtyp != wire.VarintType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(protoreflect.EnumNumber(v)), n, nil
	case protoreflect.Int32Kind:
		if wtyp != wire.VarintType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(int32(v)), n, nil
	case protoreflect.Sint32Kind:
		if wtyp != wire.VarintType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(int32(wire.DecodeZigZag(v & math.MaxUint32))), n, nil
	case protoreflect.Uint32Kind:
		if wtyp != wire.VarintType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(uint32(v)), n, nil
	case protoreflect.Int64Kind:
		if wtyp != wire.VarintType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(int64(v)), n, nil
	case protoreflect.Sint64Kind:
		if wtyp != wire.VarintType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(wire.DecodeZigZag(v)), n, nil
	case protoreflect.Uint64Kind:
		if wtyp != wire.VarintType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(v), n, nil
	case protoreflect.Sfixed32Kind:
		if wtyp != wire.Fixed32Type {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeFixed32(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(int32(v)), n, nil
	case protoreflect.Fixed32Kind:
		if wtyp != wire.Fixed32Type {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeFixed32(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(uint32(v)), n, nil
	case protoreflect.FloatKind:
		if wtyp != wire.Fixed32Type {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeFixed32(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(math.Float32frombits(uint32(v))), n, nil
	case protoreflect.Sfixed64Kind:
		if wtyp != wire.Fixed64Type {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeFixed64(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(int64(v)), n, nil
	case protoreflect.Fixed64Kind:
		if wtyp != wire.Fixed64Type {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeFixed64(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(v), n, nil
	case protoreflect.DoubleKind:
		if wtyp != wire.Fixed64Type {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeFixed64(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(math.Float64frombits(v)), n, nil
	case protoreflect.StringKind:
		if wtyp != wire.BytesType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeBytes(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(string(v)), n, nil
	case protoreflect.BytesKind:
		if wtyp != wire.BytesType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeBytes(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(append(([]byte)(nil), v...)), n, nil
	case protoreflect.MessageKind:
		if wtyp != wire.BytesType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeBytes(b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(v), n, nil
	case protoreflect.GroupKind:
		if wtyp != wire.StartGroupType {
			return val, 0, errUnknown
		}
		v, n := wire.ConsumeGroup(num, b)
		if n < 0 {
			return val, 0, wire.ParseError(n)
		}
		return protoreflect.ValueOf(v), n, nil
	default:
		return val, 0, errUnknown
	}
}

func (o UnmarshalOptions) unmarshalList(b []byte, wtyp wire.Type, num wire.Number, list protoreflect.List, kind protoreflect.Kind) (n int, err error) {
	switch kind {
	case protoreflect.BoolKind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeVarint(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(wire.DecodeBool(v)))
			}
			return n, nil
		}
		if wtyp != wire.VarintType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(wire.DecodeBool(v)))
		return n, nil
	case protoreflect.EnumKind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeVarint(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(protoreflect.EnumNumber(v)))
			}
			return n, nil
		}
		if wtyp != wire.VarintType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(protoreflect.EnumNumber(v)))
		return n, nil
	case protoreflect.Int32Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeVarint(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(int32(v)))
			}
			return n, nil
		}
		if wtyp != wire.VarintType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(int32(v)))
		return n, nil
	case protoreflect.Sint32Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeVarint(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(int32(wire.DecodeZigZag(v & math.MaxUint32))))
			}
			return n, nil
		}
		if wtyp != wire.VarintType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(int32(wire.DecodeZigZag(v & math.MaxUint32))))
		return n, nil
	case protoreflect.Uint32Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeVarint(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(uint32(v)))
			}
			return n, nil
		}
		if wtyp != wire.VarintType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(uint32(v)))
		return n, nil
	case protoreflect.Int64Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeVarint(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(int64(v)))
			}
			return n, nil
		}
		if wtyp != wire.VarintType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(int64(v)))
		return n, nil
	case protoreflect.Sint64Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeVarint(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(wire.DecodeZigZag(v)))
			}
			return n, nil
		}
		if wtyp != wire.VarintType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(wire.DecodeZigZag(v)))
		return n, nil
	case protoreflect.Uint64Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeVarint(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(v))
			}
			return n, nil
		}
		if wtyp != wire.VarintType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeVarint(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(v))
		return n, nil
	case protoreflect.Sfixed32Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeFixed32(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(int32(v)))
			}
			return n, nil
		}
		if wtyp != wire.Fixed32Type {
			return 0, errUnknown
		}
		v, n := wire.ConsumeFixed32(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(int32(v)))
		return n, nil
	case protoreflect.Fixed32Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeFixed32(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(uint32(v)))
			}
			return n, nil
		}
		if wtyp != wire.Fixed32Type {
			return 0, errUnknown
		}
		v, n := wire.ConsumeFixed32(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(uint32(v)))
		return n, nil
	case protoreflect.FloatKind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeFixed32(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(math.Float32frombits(uint32(v))))
			}
			return n, nil
		}
		if wtyp != wire.Fixed32Type {
			return 0, errUnknown
		}
		v, n := wire.ConsumeFixed32(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(math.Float32frombits(uint32(v))))
		return n, nil
	case protoreflect.Sfixed64Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeFixed64(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(int64(v)))
			}
			return n, nil
		}
		if wtyp != wire.Fixed64Type {
			return 0, errUnknown
		}
		v, n := wire.ConsumeFixed64(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(int64(v)))
		return n, nil
	case protoreflect.Fixed64Kind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeFixed64(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(v))
			}
			return n, nil
		}
		if wtyp != wire.Fixed64Type {
			return 0, errUnknown
		}
		v, n := wire.ConsumeFixed64(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(v))
		return n, nil
	case protoreflect.DoubleKind:
		if wtyp == wire.BytesType {
			buf, n := wire.ConsumeBytes(b)
			if n < 0 {
				return 0, wire.ParseError(n)
			}
			for len(buf) > 0 {
				v, n := wire.ConsumeFixed64(buf)
				if n < 0 {
					return 0, wire.ParseError(n)
				}
				buf = buf[n:]
				list.Append(protoreflect.ValueOf(math.Float64frombits(v)))
			}
			return n, nil
		}
		if wtyp != wire.Fixed64Type {
			return 0, errUnknown
		}
		v, n := wire.ConsumeFixed64(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(math.Float64frombits(v)))
		return n, nil
	case protoreflect.StringKind:
		if wtyp != wire.BytesType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeBytes(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(string(v)))
		return n, nil
	case protoreflect.BytesKind:
		if wtyp != wire.BytesType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeBytes(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		list.Append(protoreflect.ValueOf(append(([]byte)(nil), v...)))
		return n, nil
	case protoreflect.MessageKind:
		if wtyp != wire.BytesType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeBytes(b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		m := list.NewMessage()
		if err := o.unmarshalMessage(v, m); err != nil {
			return 0, err
		}
		list.Append(protoreflect.ValueOf(m))
		return n, nil
	case protoreflect.GroupKind:
		if wtyp != wire.StartGroupType {
			return 0, errUnknown
		}
		v, n := wire.ConsumeGroup(num, b)
		if n < 0 {
			return 0, wire.ParseError(n)
		}
		m := list.NewMessage()
		if err := o.unmarshalMessage(v, m); err != nil {
			return 0, err
		}
		list.Append(protoreflect.ValueOf(m))
		return n, nil
	default:
		return 0, errUnknown
	}
}