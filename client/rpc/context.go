package rpc

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

type Context struct {
	Write io.Writer
}

func (*Context) CmdHead(domain uint8, cmdCode uint16) []byte {
	head := make([]byte, 1400)
	head[6] = 0
	head[7] = byte(domain)
	binary.BigEndian.PutUint16(head[8:10], uint16(cmdCode))
	return head[0:20]
}

func (*Context) MsgHead(domain uint8, cmdCode uint16) []byte {
	head := make([]byte, 1400)
	head[6] = 1
	head[7] = byte(domain)
	binary.BigEndian.PutUint16(head[8:10], uint16(cmdCode))
	return head[0:20]
}

func (c *Context) RenderJson(head []byte, obj interface{}) error {

	// set codec  json
	head[11] = 0
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	// set  content length
	binary.BigEndian.PutUint32(head[12:16], uint32(len(jsonBytes)))
	_, err = c.Write.Write(append(head, jsonBytes...))
	return err
}

func (c *Context) RenderString(head []byte, obj string) error {

	// set codec  string
	head[11] = 2
	bytes := []byte(obj)

	// set  content length
	binary.BigEndian.PutUint32(head[12:16], uint32(len(bytes)))
	_, err := c.Write.Write(append(head, bytes...))
	return err
}

func (c *Context) RenderBin(head, payload []byte) error {

	// set codec  bin
	head[11] = 1

	// set  content length
	binary.BigEndian.PutUint32(head[12:16], uint32(len(payload)))
	_, err := c.Write.Write(append(head, payload...))
	return err
}

/**
json /0
bin/1
string/2

	msg[CmdCodeIndex+1] = 0
	msg[CodecIndex] = 1

	binary.BigEndian.PutUint32(msg[LenIndex:LenIndex+4], uint32(len(tepid)))
*/

//
//import (
//	"encoding/json"
//
//	"io"
//)
//
//type Context struct {
//	io.Writer
//	payload []byte
//
//}
//
//func (c *Context) JSON(msgType int, domain int, cmdCode int, obj interface{}) {
//	// codec / domain codec undomain
//	c.Render(cmdCode, JSON{Data: obj})
//}
//
//type Render interface {
//
//	// Render writes data with custom ContentType.
//	Render(to io.Writer) error
//	// WriteContentType writes custom ContentType.
//	//WriteContentType(w http.ResponseWriter)
//}
//
//// Render writes the response headers and calls render.Render to render data.
//func (c *Context) Render(code int, r Render) {
//
//	if err := r.Render(c.Writer); err != nil {
//		panic(err)
//	}
//}
//
//// Render (JSON) writes data with custom ContentType.
//func (r JSON) Render(w io.Writer) (err error) {
//	if err = r.WriteJSON(w, r.Data); err != nil {
//		panic(err)
//	}
//	return
//}
//
//// JSON contains the given interface object.
//type JSON struct {
//	Data interface{}
//}
//
//// WriteContentType (JSON) writes JSON ContentType.
//func (r JSON) WriteHead(w io.Writer) {
//	//writeContentType(w, jsonContentType)
//}
//
//// WriteJSON marshals the given interface object and writes it with custom ContentType.
//func (r JSON) WriteJSON(w io.Writer, obj interface{}) error {
//	//writeContentType(w, jsonContentType)
//	jsonBytes, err := json.Marshal(obj)
//	if err != nil {
//		return err
//	}
//	_, err = w.Write(jsonBytes)
//	return err
//}
