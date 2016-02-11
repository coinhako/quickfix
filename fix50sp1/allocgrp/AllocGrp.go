package allocgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/clrinstgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp1/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties"
	"github.com/quickfixgo/quickfix/fix50sp1/settlinstructionsdata"
)

//NoAllocs is a repeating group in AllocGrp
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//MatchStatus is a non-required field for NoAllocs.
	MatchStatus *string `fix:"573"`
	//AllocPrice is a non-required field for NoAllocs.
	AllocPrice *float64 `fix:"366"`
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//ProcessCode is a non-required field for NoAllocs.
	ProcessCode *string `fix:"81"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//NotifyBrokerOfCredit is a non-required field for NoAllocs.
	NotifyBrokerOfCredit *bool `fix:"208"`
	//AllocHandlInst is a non-required field for NoAllocs.
	AllocHandlInst *int `fix:"209"`
	//AllocText is a non-required field for NoAllocs.
	AllocText *string `fix:"161"`
	//EncodedAllocTextLen is a non-required field for NoAllocs.
	EncodedAllocTextLen *int `fix:"360"`
	//EncodedAllocText is a non-required field for NoAllocs.
	EncodedAllocText *string `fix:"361"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//AllocAvgPx is a non-required field for NoAllocs.
	AllocAvgPx *float64 `fix:"153"`
	//AllocNetMoney is a non-required field for NoAllocs.
	AllocNetMoney *float64 `fix:"154"`
	//SettlCurrAmt is a non-required field for NoAllocs.
	SettlCurrAmt *float64 `fix:"119"`
	//AllocSettlCurrAmt is a non-required field for NoAllocs.
	AllocSettlCurrAmt *float64 `fix:"737"`
	//SettlCurrency is a non-required field for NoAllocs.
	SettlCurrency *string `fix:"120"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//SettlCurrFxRate is a non-required field for NoAllocs.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoAllocs.
	SettlCurrFxRateCalc *string `fix:"156"`
	//AllocAccruedInterestAmt is a non-required field for NoAllocs.
	AllocAccruedInterestAmt *float64 `fix:"742"`
	//AllocInterestAtMaturity is a non-required field for NoAllocs.
	AllocInterestAtMaturity *float64 `fix:"741"`
	//MiscFeesGrp Component
	MiscFeesGrp miscfeesgrp.Component
	//ClrInstGrp Component
	ClrInstGrp clrinstgrp.Component
	//AllocSettlInstType is a non-required field for NoAllocs.
	AllocSettlInstType *int `fix:"780"`
	//SettlInstructionsData Component
	SettlInstructionsData settlinstructionsdata.Component
	//SecondaryIndividualAllocID is a non-required field for NoAllocs.
	SecondaryIndividualAllocID *string `fix:"989"`
	//AllocMethod is a non-required field for NoAllocs.
	AllocMethod *int `fix:"1002"`
	//AllocCustomerCapacity is a non-required field for NoAllocs.
	AllocCustomerCapacity *string `fix:"993"`
	//IndividualAllocType is a non-required field for NoAllocs.
	IndividualAllocType *int `fix:"992"`
	//AllocPositionEffect is a non-required field for NoAllocs.
	AllocPositionEffect *string `fix:"1047"`
	//ClearingFeeIndicator is a non-required field for NoAllocs.
	ClearingFeeIndicator *string `fix:"635"`
}

//Component is a fix50sp1 AllocGrp Component
type Component struct {
	//NoAllocs is a non-required field for AllocGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func New() *Component { return new(Component) }
