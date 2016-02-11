//Package securitystatus msg type = f.
package securitystatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//Message is a SecurityStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"f"`
	Header     fix42.Header
	//SecurityStatusReqID is a non-required field for SecurityStatus.
	SecurityStatusReqID *string `fix:"324"`
	//Symbol is a required field for SecurityStatus.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for SecurityStatus.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for SecurityStatus.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for SecurityStatus.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for SecurityStatus.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for SecurityStatus.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for SecurityStatus.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for SecurityStatus.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for SecurityStatus.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for SecurityStatus.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for SecurityStatus.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for SecurityStatus.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for SecurityStatus.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for SecurityStatus.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for SecurityStatus.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for SecurityStatus.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for SecurityStatus.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for SecurityStatus.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for SecurityStatus.
	EncodedSecurityDesc *string `fix:"351"`
	//Currency is a non-required field for SecurityStatus.
	Currency *string `fix:"15"`
	//TradingSessionID is a non-required field for SecurityStatus.
	TradingSessionID *string `fix:"336"`
	//UnsolicitedIndicator is a non-required field for SecurityStatus.
	UnsolicitedIndicator *bool `fix:"325"`
	//SecurityTradingStatus is a non-required field for SecurityStatus.
	SecurityTradingStatus *int `fix:"326"`
	//FinancialStatus is a non-required field for SecurityStatus.
	FinancialStatus *string `fix:"291"`
	//CorporateAction is a non-required field for SecurityStatus.
	CorporateAction *string `fix:"292"`
	//HaltReasonChar is a non-required field for SecurityStatus.
	HaltReasonChar *string `fix:"327"`
	//InViewOfCommon is a non-required field for SecurityStatus.
	InViewOfCommon *bool `fix:"328"`
	//DueToRelated is a non-required field for SecurityStatus.
	DueToRelated *bool `fix:"329"`
	//BuyVolume is a non-required field for SecurityStatus.
	BuyVolume *float64 `fix:"330"`
	//SellVolume is a non-required field for SecurityStatus.
	SellVolume *float64 `fix:"331"`
	//HighPx is a non-required field for SecurityStatus.
	HighPx *float64 `fix:"332"`
	//LowPx is a non-required field for SecurityStatus.
	LowPx *float64 `fix:"333"`
	//LastPx is a non-required field for SecurityStatus.
	LastPx *float64 `fix:"31"`
	//TransactTime is a non-required field for SecurityStatus.
	TransactTime *time.Time `fix:"60"`
	//Adjustment is a non-required field for SecurityStatus.
	Adjustment *int `fix:"334"`
	Trailer    fix42.Trailer
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
	return enum.BeginStringFIX42, "f", r
}
