//Package advertisement msg type = 7.
package advertisement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoLegs is a repeating group in Advertisement
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoUnderlyings is a repeating group in Advertisement
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//Message is a Advertisement FIX Message
type Message struct {
	FIXMsgType string `fix:"7"`
	Header     fix44.Header
	//AdvId is a required field for Advertisement.
	AdvId string `fix:"2"`
	//AdvTransType is a required field for Advertisement.
	AdvTransType string `fix:"5"`
	//AdvRefID is a non-required field for Advertisement.
	AdvRefID *string `fix:"3"`
	//Instrument Component
	Instrument instrument.Component
	//NoLegs is a non-required field for Advertisement.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for Advertisement.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//AdvSide is a required field for Advertisement.
	AdvSide string `fix:"4"`
	//Quantity is a required field for Advertisement.
	Quantity float64 `fix:"53"`
	//QtyType is a non-required field for Advertisement.
	QtyType *int `fix:"854"`
	//Price is a non-required field for Advertisement.
	Price *float64 `fix:"44"`
	//Currency is a non-required field for Advertisement.
	Currency *string `fix:"15"`
	//TradeDate is a non-required field for Advertisement.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for Advertisement.
	TransactTime *time.Time `fix:"60"`
	//Text is a non-required field for Advertisement.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for Advertisement.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Advertisement.
	EncodedText *string `fix:"355"`
	//URLLink is a non-required field for Advertisement.
	URLLink *string `fix:"149"`
	//LastMkt is a non-required field for Advertisement.
	LastMkt *string `fix:"30"`
	//TradingSessionID is a non-required field for Advertisement.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for Advertisement.
	TradingSessionSubID *string `fix:"625"`
	Trailer             fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX44, "7", r
}
