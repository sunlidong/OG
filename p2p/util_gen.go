package p2p

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *ArrByte) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 []byte
		zb0001, err = dc.ReadBytes([]byte((*z)))
		if err != nil {
			return
		}
		(*z) = ArrByte(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ArrByte) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes([]byte(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ArrByte) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, []byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ArrByte) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 []byte
		zb0001, bts, err = msgp.ReadBytesBytes(bts, []byte((*z)))
		if err != nil {
			return
		}
		(*z) = ArrByte(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ArrByte) Msgsize() (s int) {
	s = msgp.BytesPrefixSize + len([]byte(z))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ArrString) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(ArrString, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], err = dc.ReadString()
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ArrString) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		err = en.WriteString(z[zb0003])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ArrString) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		o = msgp.AppendString(o, z[zb0003])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ArrString) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(ArrString, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ArrString) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		s += msgp.StringPrefixSize + len(z[zb0003])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ArrUint) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(ArrUint, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], err = dc.ReadUint()
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ArrUint) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		err = en.WriteUint(z[zb0003])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ArrUint) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		o = msgp.AppendUint(o, z[zb0003])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ArrUint) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(ArrUint, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], bts, err = msgp.ReadUintBytes(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ArrUint) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (len(z) * (msgp.UintSize))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Mint64) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int64
		zb0001, err = dc.ReadInt64()
		if err != nil {
			return
		}
		(*z) = Mint64(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Mint64) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt64(int64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Mint64) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt64(o, int64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Mint64) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int64
		zb0001, bts, err = msgp.ReadInt64Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Mint64(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Mint64) Msgsize() (s int) {
	s = msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Muint) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint
		zb0001, err = dc.ReadUint()
		if err != nil {
			return
		}
		(*z) = Muint(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Muint) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint(uint(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Muint) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint(o, uint(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Muint) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint
		zb0001, bts, err = msgp.ReadUintBytes(bts)
		if err != nil {
			return
		}
		(*z) = Muint(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Muint) Msgsize() (s int) {
	s = msgp.UintSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Muint64) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint64
		zb0001, err = dc.ReadUint64()
		if err != nil {
			return
		}
		(*z) = Muint64(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Muint64) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint64(uint64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Muint64) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint64(o, uint64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Muint64) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint64
		zb0001, bts, err = msgp.ReadUint64Bytes(bts)
		if err != nil {
			return
		}
		(*z) = Muint64(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Muint64) Msgsize() (s int) {
	s = msgp.Uint64Size
	return
}
