// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: entry.go
// DO NOT EDIT!

package oui

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Entry) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(512)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Entry) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{ "manufacturer":`)
	fflib.WriteJsonString(buf, string(mj.Manufacturer))
	buf.WriteString(`,"address":`)
	if mj.Address != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Address {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.WriteJsonString(buf, string(v))
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"prefix":`)

	{
		obj, err = mj.Prefix.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)
	}

	buf.WriteByte(',')
	if len(mj.Country) != 0 {
		buf.WriteString(`"country":`)
		fflib.WriteJsonString(buf, string(mj.Country))
		buf.WriteByte(',')
	}
	if mj.Local != false {
		if mj.Local {
			buf.WriteString(`"local":true`)
		} else {
			buf.WriteString(`"local":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Multicast != false {
		if mj.Multicast {
			buf.WriteString(`"multicast":true`)
		} else {
			buf.WriteString(`"multicast":false`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}
