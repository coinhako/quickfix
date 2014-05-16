//Package news msg type = B.
package news

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a News wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//OrigTime is a non-required field for News.
func (m Message) OrigTime() (*field.OrigTimeField, errors.MessageRejectError) {
	f := &field.OrigTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from News.
func (m Message) GetOrigTime(f *field.OrigTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Urgency is a non-required field for News.
func (m Message) Urgency() (*field.UrgencyField, errors.MessageRejectError) {
	f := &field.UrgencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUrgency reads a Urgency from News.
func (m Message) GetUrgency(f *field.UrgencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RelatdSym is a non-required field for News.
func (m Message) RelatdSym() (*field.RelatdSymField, errors.MessageRejectError) {
	f := &field.RelatdSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRelatdSym reads a RelatdSym from News.
func (m Message) GetRelatdSym(f *field.RelatdSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LinesOfText is a required field for News.
func (m Message) LinesOfText() (*field.LinesOfTextField, errors.MessageRejectError) {
	f := &field.LinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLinesOfText reads a LinesOfText from News.
func (m Message) GetLinesOfText(f *field.LinesOfTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a required field for News.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from News.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for News.
func (m Message) RawDataLength() (*field.RawDataLengthField, errors.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from News.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for News.
func (m Message) RawData() (*field.RawDataField, errors.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from News.
func (m Message) GetRawData(f *field.RawDataField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds News messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for News.
func Builder(
	linesoftext *field.LinesOfTextField,
	text *field.TextField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("B"))
	builder.Body().Set(linesoftext)
	builder.Body().Set(text)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "B", r
}