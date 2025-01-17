package emv

import "github.com/mercadolibre/iso8583/field"

type Data struct {
	AcquirerIdentifier                                          *field.String  `index:"9F01"`
	AdditionalTerminalCapabilities                              *field.String  `index:"9F40"`
	AmountAuthorisedBinary                                      *field.String  `index:"81"`
	AmountAuthorisedNumeric                                     *field.Numeric `index:"9F02"`
	AmountOtherBinary                                           *field.String  `index:"9F04"`
	AmountOtherNumeric                                          *field.Numeric `index:"9F03"`
	AmountReferenceCurrency                                     *field.String  `index:"9F3A"`
	ApplicationCryptogram                                       *field.String  `index:"9F26"`
	ApplicationCurrencyCode                                     *field.String  `index:"9F42"`
	ApplicationCurrencyExponent                                 *field.String  `index:"9F44"`
	ApplicationDiscretionaryData                                *field.String  `index:"9F05"`
	ApplicationEffectiveDate                                    *field.String  `index:"5F25"`
	ApplicationExpirationDate                                   *field.String  `index:"5F24"`
	ApplicationFileLocatorAFL                                   *field.String  `index:"94"`
	ApplicationIdentifierAIDcard                                *field.String  `index:"4F"`
	ApplicationIdentifierAIDterminal                            *field.String  `index:"9F06"`
	ApplicationInterchangeProfile                               *field.String  `index:"82"`
	ApplicationLabel                                            *field.String  `index:"50"`
	ApplicationPreferredName                                    *field.String  `index:"9F12"`
	ApplicationPrimaryAccountNumberPAN                          *field.String  `index:"5A"`
	ApplicationPrimaryAccountNumberPANSequenceNumber            *field.String  `index:"5F34"`
	ApplicationPriorityIndicator                                *field.String  `index:"87"`
	ApplicationReferenceCurrency                                *field.String  `index:"9F3B"`
	ApplicationReferenceCurrencyExponent                        *field.String  `index:"9F43"`
	ApplicationSelectionRegisteredProprietaryData               *field.String  `index:"9F0A"`
	ApplicationTemplate                                         *field.String  `index:"61"`
	ApplicationTransactionCounter                               *field.Numeric `index:"9F36"`
	ApplicationUsageControl                                     *field.String  `index:"9F07"`
	ApplicationVersionNumber                                    *field.String  `index:"9F08"`
	ApplicationVersionNumberTerminal                            *field.String  `index:"9F09"`
	AuthorisationCode                                           *field.String  `index:"89"`
	AuthorisationResponseCode                                   *field.String  `index:"8A"`
	BankIdentifierCodeBIC                                       *field.String  `index:"5F54"`
	CardBITGroupTemplate                                        *field.String  `index:"9F31"`
	CardRiskManagementDataObjectList1CDOL1                      *field.String  `index:"8C"`
	CardRiskManagementDataObjectList2CDOL2                      *field.String  `index:"8D"`
	CardholderName                                              *field.String  `index:"5F20"`
	CardholderNameExtended                                      *field.String  `index:"9F0B"`
	CardholderVerificationMethodCVMList                         *field.String  `index:"8E"`
	CardholderVerificationMethodCVMResults                      *field.String  `index:"9F34"`
	CertificationAuthorityPublicKeyIndex                        *field.String  `index:"8F"`
	CertificationAuthorityPublicKeyIndexTerminal                *field.String  `index:"9F22"`
	CommandTemplate                                             *field.String  `index:"83"`
	CryptogramInformationData                                   *field.String  `index:"9F27"`
	DataAuthenticationCode                                      *field.String  `index:"9F45"`
	DedicatedFileDFName                                         *field.String  `index:"84"`
	DirectoryDefinitionFileDDFName                              *field.String  `index:"9D"`
	DirectoryDiscretionaryTemplate                              *field.String  `index:"73"`
	DynamicDataAuthenticationDataObjectListDDOL                 *field.String  `index:"9F49"`
	EMVProprietaryTemplate                                      *field.String  `index:"70"`
	FacialTryCounter                                            *field.String  `index:"DF50"`
	FileControlInformationFCIIssuerDiscretionaryData            *field.String  `index:"BF0C"`
	FileControlInformationFCIProprietaryTemplate                *field.String  `index:"A5"`
	FileControlInformationFCITemplate                           *field.String  `index:"6F"`
	FingerTryCounter                                            *field.String  `index:"DF51"`
	ICCDynamicNumber                                            *field.String  `index:"9F4C"`
	IntegratedCircuitCardICCPINEnciphermentPublicKeyCertificate *field.String  `index:"9F2D"`
	IntegratedCircuitCardICCPINEnciphermentPublicKeyExponent    *field.String  `index:"9F2E"`
	IntegratedCircuitCardICCPINEnciphermentPublicKeyRemainder   *field.String  `index:"9F2F"`
	IntegratedCircuitCardICCPublicKeyCertificate                *field.String  `index:"9F46"`
	IntegratedCircuitCardICCPublicKeyExponent                   *field.String  `index:"9F47"`
	IntegratedCircuitCardICCPublicKeyRemainder                  *field.String  `index:"9F48"`
	InterfaceDeviceIFDSerialNumber                              *field.String  `index:"9F1E"`
	InternationalBankAccountNumberIBAN                          *field.String  `index:"5F53"`
	IssuerActionCodeDefault                                     *field.String  `index:"9F0D"`
	IssuerActionCodeDenial                                      *field.String  `index:"9F0E"`
	IssuerActionCodeOnline                                      *field.String  `index:"9F0F"`
	IssuerApplicationData                                       *field.String  `index:"9F10"`
	IssuerAuthenticationData                                    *field.String  `index:"91"`
	IssuerCodeTableIndex                                        *field.String  `index:"9F11"`
	IssuerCountryCode                                           *field.String  `index:"5F28"`
	IssuerCountryCodealpha2format                               *field.String  `index:"5F55"`
	IssuerCountryCodealpha3format                               *field.String  `index:"5F56"`
	IssuerIdentificationNumberIIN                               *field.String  `index:"42"`
	IssuerIdentificationNumberExtended                          *field.String  `index:"9F0C"`
	IssuerPublicKeyCertificate                                  *field.String  `index:"90"`
	IssuerPublicKeyExponent                                     *field.String  `index:"9F32"`
	IssuerPublicKeyRemainder                                    *field.String  `index:"92"`
	IssuerScriptCommand                                         *field.String  `index:"86"`
	IssuerScriptIdentifier                                      *field.String  `index:"9F18"`
	IssuerScriptTemplate1                                       *field.String  `index:"71"`
	IssuerScriptTemplate2                                       *field.String  `index:"72"`
	IssuerURL                                                   *field.String  `index:"5F50"`
	LanguagePreference                                          *field.String  `index:"5F2D"`
	LastOnlineApplicationTransactionCounterATCRegister          *field.String  `index:"9F13"`
	LogEntry                                                    *field.String  `index:"9F4D"`
	LogFormat                                                   *field.String  `index:"9F4F"`
	LowerConsecutiveOfflineLimit                                *field.String  `index:"9F14"`
	MerchantCategoryCode                                        *field.String  `index:"9F15"`
	MerchantIdentifier                                          *field.String  `index:"9F16"`
	MerchantNameandLocation                                     *field.String  `index:"9F4E"`
	PaymentAccountReferencePAR                                  *field.String  `index:"9F24"`
	PersonalIdentificationNumberPINTryCounter                   *field.String  `index:"9F17"`
	PointofServicePOSEntryMode                                  *field.String  `index:"9F39"`
	ProcessingOptionsDataObjectListPDOL                         *field.String  `index:"9F38"`
	ResponseMessageTemplateFormat1                              *field.String  `index:"80"`
	ResponseMessageTemplateFormat2                              *field.String  `index:"77"`
	ServiceCode                                                 *field.String  `index:"5F30"`
	ShortFileIdentifierSFI                                      *field.String  `index:"88"`
	SignedDynamicApplicationData                                *field.String  `index:"9F4B"`
	SignedStaticApplicationData                                 *field.String  `index:"93"`
	StaticDataAuthenticationTagList                             *field.String  `index:"9F4A"`
	TerminalCapabilities                                        *field.String  `index:"9F33"`
	TerminalCountryCode                                         *field.String  `index:"9F1A"`
	TerminalFloorLimit                                          *field.String  `index:"9F1B"`
	TerminalIdentification                                      *field.String  `index:"9F1C"`
	TerminalRiskManagementData                                  *field.String  `index:"9F1D"`
	TerminalType                                                *field.String  `index:"9F35"`
	TerminalVerificationResults                                 *field.String  `index:"95"`
	TokenRequestorID                                            *field.String  `index:"9F19"`
	Track1DiscretionaryData                                     *field.String  `index:"9F1F"`
	Track2DiscretionaryData                                     *field.String  `index:"9F20"`
	Track2EquivalentData                                        *field.String  `index:"57"`
	TransactionCertificateTCHashValue                           *field.String  `index:"98"`
	TransactionCertificateDataObjectListTDOL                    *field.String  `index:"97"`
	TransactionCurrencyCode                                     *field.String  `index:"5F2A"`
	TransactionCurrencyExponent                                 *field.String  `index:"5F36"`
	TransactionDate                                             *field.String  `index:"9A"`
	TransactionPersonalIdentificationNumberPINData              *field.String  `index:"99"`
	TransactionReferenceCurrencyCode                            *field.String  `index:"9F3C"`
	TransactionReferenceCurrencyExponent                        *field.String  `index:"9F3D"`
	TransactionSequenceCounter                                  *field.String  `index:"9F41"`
	TransactionStatusInformation                                *field.String  `index:"9B"`
	TransactionTime                                             *field.String  `index:"9F21"`
	TransactionType                                             *field.String  `index:"9C"`
	UnpredictableNumber                                         *field.String  `index:"9F37"`
	UpperConsecutiveOfflineLimit                                *field.String  `index:"9F23"`
}
