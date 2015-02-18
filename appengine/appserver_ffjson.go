package appengine

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Response) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(128)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Response) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.Data != nil {
		if true {
			buf.WriteString(`"data":`)

			{
				err = mj.Data.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	if len(mj.Error) != 0 {
		buf.WriteString(`"error":`)
		fflib.WriteJsonString(buf, string(mj.Error))
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}
