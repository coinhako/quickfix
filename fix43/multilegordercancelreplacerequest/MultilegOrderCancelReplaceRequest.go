//Package multilegordercancelreplacerequest msg type = AC.
package multilegordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/instrumentleg"
	"github.com/quickfixgo/quickfix/fix43/nestedparties"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"time"
)

//NoAllocs is a repeating group in MultilegOrderCancelReplaceRequest
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//NoTradingSessions is a repeating group in MultilegOrderCancelReplaceRequest
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//NoLegs is a repeating group in MultilegOrderCancelReplaceRequest
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegPrice is a non-required field for NoLegs.
	LegPrice *float64 `fix:"566"`
	//LegSettlmntTyp is a non-required field for NoLegs.
	LegSettlmntTyp *string `fix:"587"`
	//LegFutSettDate is a non-required field for NoLegs.
	LegFutSettDate *string `fix:"588"`
}

//Message is a MultilegOrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AC"`
	Header     fix43.Header
	//OrderID is a non-required field for MultilegOrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//OrigClOrdID is a required field for MultilegOrderCancelReplaceRequest.
	OrigClOrdID string `fix:"41"`
	//ClOrdID is a required field for MultilegOrderCancelReplaceRequest.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplaceRequest.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for MultilegOrderCancelReplaceRequest.
	ClOrdLinkID *string `fix:"583"`
	//OrigOrdModTime is a non-required field for MultilegOrderCancelReplaceRequest.
	OrigOrdModTime *time.Time `fix:"586"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for MultilegOrderCancelReplaceRequest.
	Account *string `fix:"1"`
	//AccountType is a non-required field for MultilegOrderCancelReplaceRequest.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for MultilegOrderCancelReplaceRequest.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for MultilegOrderCancelReplaceRequest.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for MultilegOrderCancelReplaceRequest.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlmntTyp is a non-required field for MultilegOrderCancelReplaceRequest.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for MultilegOrderCancelReplaceRequest.
	FutSettDate *string `fix:"64"`
	//CashMargin is a non-required field for MultilegOrderCancelReplaceRequest.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplaceRequest.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a required field for MultilegOrderCancelReplaceRequest.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for MultilegOrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for MultilegOrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for MultilegOrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for MultilegOrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for MultilegOrderCancelReplaceRequest.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for MultilegOrderCancelReplaceRequest.
	ProcessCode *string `fix:"81"`
	//Side is a required field for MultilegOrderCancelReplaceRequest.
	Side string `fix:"54"`
	//Instrument Component
	Instrument instrument.Component
	//PrevClosePx is a non-required field for MultilegOrderCancelReplaceRequest.
	PrevClosePx *float64 `fix:"140"`
	//NoLegs is a required field for MultilegOrderCancelReplaceRequest.
	NoLegs []NoLegs `fix:"555"`
	//LocateReqd is a non-required field for MultilegOrderCancelReplaceRequest.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for MultilegOrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//QuantityType is a non-required field for MultilegOrderCancelReplaceRequest.
	QuantityType *int `fix:"465"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a required field for MultilegOrderCancelReplaceRequest.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for MultilegOrderCancelReplaceRequest.
	PriceType *int `fix:"423"`
	//Price is a non-required field for MultilegOrderCancelReplaceRequest.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for MultilegOrderCancelReplaceRequest.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for MultilegOrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for MultilegOrderCancelReplaceRequest.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for MultilegOrderCancelReplaceRequest.
	SolicitedFlag *bool `fix:"377"`
	//IOIid is a non-required field for MultilegOrderCancelReplaceRequest.
	IOIid *string `fix:"23"`
	//QuoteID is a non-required field for MultilegOrderCancelReplaceRequest.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for MultilegOrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for MultilegOrderCancelReplaceRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for MultilegOrderCancelReplaceRequest.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for MultilegOrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
	GTBookingInst *int `fix:"427"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//OrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for MultilegOrderCancelReplaceRequest.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for MultilegOrderCancelReplaceRequest.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for MultilegOrderCancelReplaceRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for MultilegOrderCancelReplaceRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for MultilegOrderCancelReplaceRequest.
	EncodedText *string `fix:"355"`
	//PositionEffect is a non-required field for MultilegOrderCancelReplaceRequest.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplaceRequest.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for MultilegOrderCancelReplaceRequest.
	MaxShow *float64 `fix:"210"`
	//PegDifference is a non-required field for MultilegOrderCancelReplaceRequest.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for MultilegOrderCancelReplaceRequest.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for MultilegOrderCancelReplaceRequest.
	DiscretionOffset *float64 `fix:"389"`
	//CancellationRights is a non-required field for MultilegOrderCancelReplaceRequest.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplaceRequest.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for MultilegOrderCancelReplaceRequest.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for MultilegOrderCancelReplaceRequest.
	Designation *string `fix:"494"`
	//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplaceRequest.
	MultiLegRptTypeReq *int `fix:"563"`
	//NetMoney is a non-required field for MultilegOrderCancelReplaceRequest.
	NetMoney *float64 `fix:"118"`
	Trailer  fix43.Trailer
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
	return enum.BeginStringFIX43, "AC", r
}
