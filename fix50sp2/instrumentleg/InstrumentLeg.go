package instrumentleg

//NoLegSecurityAltID is a repeating group in InstrumentLeg
type NoLegSecurityAltID struct {
	//LegSecurityAltID is a non-required field for NoLegSecurityAltID.
	LegSecurityAltID *string `fix:"605"`
	//LegSecurityAltIDSource is a non-required field for NoLegSecurityAltID.
	LegSecurityAltIDSource *string `fix:"606"`
}

//Component is a fix50sp2 InstrumentLeg Component
type Component struct {
	//LegSymbol is a non-required field for InstrumentLeg.
	LegSymbol *string `fix:"600"`
	//LegSymbolSfx is a non-required field for InstrumentLeg.
	LegSymbolSfx *string `fix:"601"`
	//LegSecurityID is a non-required field for InstrumentLeg.
	LegSecurityID *string `fix:"602"`
	//LegSecurityIDSource is a non-required field for InstrumentLeg.
	LegSecurityIDSource *string `fix:"603"`
	//NoLegSecurityAltID is a non-required field for InstrumentLeg.
	NoLegSecurityAltID []NoLegSecurityAltID `fix:"604,omitempty"`
	//LegProduct is a non-required field for InstrumentLeg.
	LegProduct *int `fix:"607"`
	//LegCFICode is a non-required field for InstrumentLeg.
	LegCFICode *string `fix:"608"`
	//LegSecurityType is a non-required field for InstrumentLeg.
	LegSecurityType *string `fix:"609"`
	//LegSecuritySubType is a non-required field for InstrumentLeg.
	LegSecuritySubType *string `fix:"764"`
	//LegMaturityMonthYear is a non-required field for InstrumentLeg.
	LegMaturityMonthYear *string `fix:"610"`
	//LegMaturityDate is a non-required field for InstrumentLeg.
	LegMaturityDate *string `fix:"611"`
	//LegCouponPaymentDate is a non-required field for InstrumentLeg.
	LegCouponPaymentDate *string `fix:"248"`
	//LegIssueDate is a non-required field for InstrumentLeg.
	LegIssueDate *string `fix:"249"`
	//LegRepoCollateralSecurityType is a non-required field for InstrumentLeg.
	LegRepoCollateralSecurityType *int `fix:"250"`
	//LegRepurchaseTerm is a non-required field for InstrumentLeg.
	LegRepurchaseTerm *int `fix:"251"`
	//LegRepurchaseRate is a non-required field for InstrumentLeg.
	LegRepurchaseRate *float64 `fix:"252"`
	//LegFactor is a non-required field for InstrumentLeg.
	LegFactor *float64 `fix:"253"`
	//LegCreditRating is a non-required field for InstrumentLeg.
	LegCreditRating *string `fix:"257"`
	//LegInstrRegistry is a non-required field for InstrumentLeg.
	LegInstrRegistry *string `fix:"599"`
	//LegCountryOfIssue is a non-required field for InstrumentLeg.
	LegCountryOfIssue *string `fix:"596"`
	//LegStateOrProvinceOfIssue is a non-required field for InstrumentLeg.
	LegStateOrProvinceOfIssue *string `fix:"597"`
	//LegLocaleOfIssue is a non-required field for InstrumentLeg.
	LegLocaleOfIssue *string `fix:"598"`
	//LegRedemptionDate is a non-required field for InstrumentLeg.
	LegRedemptionDate *string `fix:"254"`
	//LegStrikePrice is a non-required field for InstrumentLeg.
	LegStrikePrice *float64 `fix:"612"`
	//LegStrikeCurrency is a non-required field for InstrumentLeg.
	LegStrikeCurrency *string `fix:"942"`
	//LegOptAttribute is a non-required field for InstrumentLeg.
	LegOptAttribute *string `fix:"613"`
	//LegContractMultiplier is a non-required field for InstrumentLeg.
	LegContractMultiplier *float64 `fix:"614"`
	//LegCouponRate is a non-required field for InstrumentLeg.
	LegCouponRate *float64 `fix:"615"`
	//LegSecurityExchange is a non-required field for InstrumentLeg.
	LegSecurityExchange *string `fix:"616"`
	//LegIssuer is a non-required field for InstrumentLeg.
	LegIssuer *string `fix:"617"`
	//EncodedLegIssuerLen is a non-required field for InstrumentLeg.
	EncodedLegIssuerLen *int `fix:"618"`
	//EncodedLegIssuer is a non-required field for InstrumentLeg.
	EncodedLegIssuer *string `fix:"619"`
	//LegSecurityDesc is a non-required field for InstrumentLeg.
	LegSecurityDesc *string `fix:"620"`
	//EncodedLegSecurityDescLen is a non-required field for InstrumentLeg.
	EncodedLegSecurityDescLen *int `fix:"621"`
	//EncodedLegSecurityDesc is a non-required field for InstrumentLeg.
	EncodedLegSecurityDesc *string `fix:"622"`
	//LegRatioQty is a non-required field for InstrumentLeg.
	LegRatioQty *float64 `fix:"623"`
	//LegSide is a non-required field for InstrumentLeg.
	LegSide *string `fix:"624"`
	//LegCurrency is a non-required field for InstrumentLeg.
	LegCurrency *string `fix:"556"`
	//LegPool is a non-required field for InstrumentLeg.
	LegPool *string `fix:"740"`
	//LegDatedDate is a non-required field for InstrumentLeg.
	LegDatedDate *string `fix:"739"`
	//LegContractSettlMonth is a non-required field for InstrumentLeg.
	LegContractSettlMonth *string `fix:"955"`
	//LegInterestAccrualDate is a non-required field for InstrumentLeg.
	LegInterestAccrualDate *string `fix:"956"`
	//LegUnitOfMeasure is a non-required field for InstrumentLeg.
	LegUnitOfMeasure *string `fix:"999"`
	//LegTimeUnit is a non-required field for InstrumentLeg.
	LegTimeUnit *string `fix:"1001"`
	//LegOptionRatio is a non-required field for InstrumentLeg.
	LegOptionRatio *float64 `fix:"1017"`
	//LegPrice is a non-required field for InstrumentLeg.
	LegPrice *float64 `fix:"566"`
	//LegMaturityTime is a non-required field for InstrumentLeg.
	LegMaturityTime *string `fix:"1212"`
	//LegPutOrCall is a non-required field for InstrumentLeg.
	LegPutOrCall *int `fix:"1358"`
	//LegExerciseStyle is a non-required field for InstrumentLeg.
	LegExerciseStyle *int `fix:"1420"`
	//LegUnitOfMeasureQty is a non-required field for InstrumentLeg.
	LegUnitOfMeasureQty *float64 `fix:"1224"`
	//LegPriceUnitOfMeasure is a non-required field for InstrumentLeg.
	LegPriceUnitOfMeasure *string `fix:"1421"`
	//LegPriceUnitOfMeasureQty is a non-required field for InstrumentLeg.
	LegPriceUnitOfMeasureQty *float64 `fix:"1422"`
	//LegContractMultiplierUnit is a non-required field for InstrumentLeg.
	LegContractMultiplierUnit *int `fix:"1436"`
	//LegFlowScheduleType is a non-required field for InstrumentLeg.
	LegFlowScheduleType *int `fix:"1440"`
}

func New() *Component { return new(Component) }
