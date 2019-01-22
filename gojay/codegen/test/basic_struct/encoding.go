// Code generated by GoJayGen. DO NOT EDIT.\n\n
package basic_struct

import (
	"github.com/francoispqt/gojay"
	"time"
)

type Ints []int

//UnmarshalJSONArray decodes JSON array elements into slice
func (a *Ints) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value int
	if err := dec.Int(&value); err != nil {
		return err
	}
	*a = append(*a, value)
	return nil
}

//MarshalJSONArray encodes arrays into JSON
func (a Ints) MarshalJSONArray(enc *gojay.Encoder) {
	for _, item := range a {
		enc.Int(item)
	}
}

//IsNil checks if array is nil
func (a Ints) IsNil() bool {
	return len(a) == 0
}

type Float32sPtr []*float32

//UnmarshalJSONArray decodes JSON array elements into slice
func (a *Float32sPtr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value float32
	if err := dec.Float32(&value); err != nil {
		return err
	}
	*a = append(*a, &value)
	return nil
}

//MarshalJSONArray encodes arrays into JSON
func (a Float32sPtr) MarshalJSONArray(enc *gojay.Encoder) {
	for _, item := range a {
		enc.Float32(*item)
	}
}

//IsNil checks if array is nil
func (a Float32sPtr) IsNil() bool {
	return len(a) == 0
}

type SubMessagesPtr []*SubMessage

func (s *SubMessagesPtr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value = &SubMessage{}
	if err := dec.Object(value); err != nil {
		return err
	}
	*s = append(*s, value)
	return nil
}

func (s SubMessagesPtr) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range s {
		enc.Object(s[i])
	}
}

func (s SubMessagesPtr) IsNil() bool {
	return len(s) == 0
}

type SubMessages []SubMessage

func (s *SubMessages) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value = SubMessage{}
	if err := dec.Object(&value); err != nil {
		return err
	}
	*s = append(*s, value)
	return nil
}

func (s SubMessages) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range s {
		enc.Object(&s[i])
	}
}

func (s SubMessages) IsNil() bool {
	return len(s) == 0
}

//MarshalJSONObject implements MarshalerJSONObject
func (m *SubMessage) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("Id", m.Id)
	enc.StringKey("Description", m.Description)
	enc.TimeKey("StartTime", &m.StartTime, time.RFC3339)
	if m.EndTime != nil {
		enc.TimeKey("EndTime", m.EndTime, time.RFC3339)
	}
}

//IsNil checks if instance is nil
func (m *SubMessage) IsNil() bool {
	return m == nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (m *SubMessage) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "Id":
		return dec.Int(&m.Id)

	case "Description":
		return dec.String(&m.Description)

	case "StartTime":
		var format = time.RFC3339
		var value = time.Time{}
		err := dec.DecodeTime(&value, format)
		return err

	case "EndTime":
		var format = time.RFC3339
		var value = &time.Time{}
		err := dec.DecodeTime(value, format)
		if err == nil {
			m.EndTime = value
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (m *SubMessage) NKeys() int { return 4 }

//MarshalJSONObject implements MarshalerJSONObject
func (m *Message) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("Id", m.Id)
	enc.StringKey("Name", m.Name)
	enc.Float64Key("Price", m.Price)
	var intsSlice = Ints(m.Ints)
	enc.ArrayKey("Ints", intsSlice)
	var floatsSlice = Float32sPtr(m.Floats)
	enc.ArrayKey("Floats", floatsSlice)
	enc.ObjectKey("SubMessageX", m.SubMessageX)
	var messagesXSlice = SubMessagesPtr(m.MessagesX)
	enc.ArrayKey("MessagesX", messagesXSlice)
	enc.ObjectKey("SubMessageY", &m.SubMessageY)
	var messagesYSlice = SubMessages(m.MessagesY)
	enc.ArrayKey("MessagesY", messagesYSlice)
	enc.BoolKey("IsTrue", *m.IsTrue)
	var payloadSlice = gojay.EmbeddedJSON(m.Payload)
	enc.AddEmbeddedJSONKey("Payload", &payloadSlice)
}

//IsNil checks if instance is nil
func (m *Message) IsNil() bool {
	return m == nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (m *Message) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "Id":
		return dec.Int(&m.Id)

	case "Name":
		return dec.String(&m.Name)

	case "Price":
		return dec.Float64(&m.Price)

	case "Ints":
		var aSlice = Ints{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			m.Ints = []int(aSlice)
		}
		return err

	case "Floats":
		var aSlice = Float32sPtr{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			m.Floats = []*float32(aSlice)
		}
		return err

	case "SubMessageX":
		var value = &SubMessage{}
		err := dec.Object(value)
		if err == nil {
			m.SubMessageX = value
		}

		return err

	case "MessagesX":
		var aSlice = SubMessagesPtr{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			m.MessagesX = []*SubMessage(aSlice)
		}
		return err

	case "SubMessageY":
		err := dec.Object(&m.SubMessageY)

		return err

	case "MessagesY":
		var aSlice = SubMessages{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			m.MessagesY = []SubMessage(aSlice)
		}
		return err

	case "IsTrue":
		var value bool
		err := dec.Bool(&value)
		if err == nil {
			m.IsTrue = &value
		}
		return err

	case "Payload":
		var value = gojay.EmbeddedJSON{}
		err := dec.AddEmbeddedJSON(&value)
		if err == nil && len(value) > 0 {
			m.Payload = []byte(value)
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (m *Message) NKeys() int { return 11 }
