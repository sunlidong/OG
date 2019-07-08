package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/annchain/OG/common/math"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Tx) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 6 {
		err = msgp.ArrayError{Wanted: 6, Got: zb0001}
		return
	}
	err = z.TxBase.DecodeMsg(dc)
	if err != nil {
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			return
		}
		z.From = nil
	} else {
		if z.From == nil {
			z.From = new(Address)
		}
		err = z.From.DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	err = z.To.DecodeMsg(dc)
	if err != nil {
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			return
		}
		z.Value = nil
	} else {
		if z.Value == nil {
			z.Value = new(math.BigInt)
		}
		err = z.Value.DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	z.TokenId, err = dc.ReadInt32()
	if err != nil {
		return
	}
	z.Data, err = dc.ReadBytes(z.Data)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Tx) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 6
	err = en.Append(0x96)
	if err != nil {
		return
	}
	err = z.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	if z.From == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.From.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = z.To.EncodeMsg(en)
	if err != nil {
		return
	}
	if z.Value == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Value.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = en.WriteInt32(z.TokenId)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Tx) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 6
	o = append(o, 0x96)
	o, err = z.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	if z.From == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.From.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o, err = z.To.MarshalMsg(o)
	if err != nil {
		return
	}
	if z.Value == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Value.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o = msgp.AppendInt32(o, z.TokenId)
	o = msgp.AppendBytes(o, z.Data)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tx) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 6 {
		err = msgp.ArrayError{Wanted: 6, Got: zb0001}
		return
	}
	bts, err = z.TxBase.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.From = nil
	} else {
		if z.From == nil {
			z.From = new(Address)
		}
		bts, err = z.From.UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	bts, err = z.To.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Value = nil
	} else {
		if z.Value == nil {
			z.Value = new(math.BigInt)
		}
		bts, err = z.Value.UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	z.TokenId, bts, err = msgp.ReadInt32Bytes(bts)
	if err != nil {
		return
	}
	z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tx) Msgsize() (s int) {
	s = 1 + z.TxBase.Msgsize()
	if z.From == nil {
		s += msgp.NilSize
	} else {
		s += z.From.Msgsize()
	}
	s += z.To.Msgsize()
	if z.Value == nil {
		s += msgp.NilSize
	} else {
		s += z.Value.Msgsize()
	}
	s += msgp.Int32Size + msgp.BytesPrefixSize + len(z.Data)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Txs) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Txs, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Tx)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Txs) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Txs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Txs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Txs, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Tx)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Txs) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
