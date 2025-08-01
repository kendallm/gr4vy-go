// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RedirectPaymentMethodCreateMethod - The method to use, this can be any of the methods that support redirect requests.
type RedirectPaymentMethodCreateMethod string

const (
	RedirectPaymentMethodCreateMethodAbitab          RedirectPaymentMethodCreateMethod = "abitab"
	RedirectPaymentMethodCreateMethodAffirm          RedirectPaymentMethodCreateMethod = "affirm"
	RedirectPaymentMethodCreateMethodAfterpay        RedirectPaymentMethodCreateMethod = "afterpay"
	RedirectPaymentMethodCreateMethodAlipay          RedirectPaymentMethodCreateMethod = "alipay"
	RedirectPaymentMethodCreateMethodAlipayhk        RedirectPaymentMethodCreateMethod = "alipayhk"
	RedirectPaymentMethodCreateMethodArcuspaynetwork RedirectPaymentMethodCreateMethod = "arcuspaynetwork"
	RedirectPaymentMethodCreateMethodBacs            RedirectPaymentMethodCreateMethod = "bacs"
	RedirectPaymentMethodCreateMethodBancontact      RedirectPaymentMethodCreateMethod = "bancontact"
	RedirectPaymentMethodCreateMethodBanked          RedirectPaymentMethodCreateMethod = "banked"
	RedirectPaymentMethodCreateMethodBcp             RedirectPaymentMethodCreateMethod = "bcp"
	RedirectPaymentMethodCreateMethodBecs            RedirectPaymentMethodCreateMethod = "becs"
	RedirectPaymentMethodCreateMethodBitpay          RedirectPaymentMethodCreateMethod = "bitpay"
	RedirectPaymentMethodCreateMethodBlik            RedirectPaymentMethodCreateMethod = "blik"
	RedirectPaymentMethodCreateMethodBoleto          RedirectPaymentMethodCreateMethod = "boleto"
	RedirectPaymentMethodCreateMethodBoost           RedirectPaymentMethodCreateMethod = "boost"
	RedirectPaymentMethodCreateMethodCashapp         RedirectPaymentMethodCreateMethod = "cashapp"
	RedirectPaymentMethodCreateMethodClearpay        RedirectPaymentMethodCreateMethod = "clearpay"
	RedirectPaymentMethodCreateMethodDana            RedirectPaymentMethodCreateMethod = "dana"
	RedirectPaymentMethodCreateMethodDcb             RedirectPaymentMethodCreateMethod = "dcb"
	RedirectPaymentMethodCreateMethodDlocal          RedirectPaymentMethodCreateMethod = "dlocal"
	RedirectPaymentMethodCreateMethodEbanx           RedirectPaymentMethodCreateMethod = "ebanx"
	RedirectPaymentMethodCreateMethodEfecty          RedirectPaymentMethodCreateMethod = "efecty"
	RedirectPaymentMethodCreateMethodEps             RedirectPaymentMethodCreateMethod = "eps"
	RedirectPaymentMethodCreateMethodEverydaypay     RedirectPaymentMethodCreateMethod = "everydaypay"
	RedirectPaymentMethodCreateMethodGcash           RedirectPaymentMethodCreateMethod = "gcash"
	RedirectPaymentMethodCreateMethodGem             RedirectPaymentMethodCreateMethod = "gem"
	RedirectPaymentMethodCreateMethodGemds           RedirectPaymentMethodCreateMethod = "gemds"
	RedirectPaymentMethodCreateMethodGiropay         RedirectPaymentMethodCreateMethod = "giropay"
	RedirectPaymentMethodCreateMethodGivingblock     RedirectPaymentMethodCreateMethod = "givingblock"
	RedirectPaymentMethodCreateMethodGocardless      RedirectPaymentMethodCreateMethod = "gocardless"
	RedirectPaymentMethodCreateMethodGopay           RedirectPaymentMethodCreateMethod = "gopay"
	RedirectPaymentMethodCreateMethodGrabpay         RedirectPaymentMethodCreateMethod = "grabpay"
	RedirectPaymentMethodCreateMethodIdeal           RedirectPaymentMethodCreateMethod = "ideal"
	RedirectPaymentMethodCreateMethodKakaopay        RedirectPaymentMethodCreateMethod = "kakaopay"
	RedirectPaymentMethodCreateMethodKcp             RedirectPaymentMethodCreateMethod = "kcp"
	RedirectPaymentMethodCreateMethodKhipu           RedirectPaymentMethodCreateMethod = "khipu"
	RedirectPaymentMethodCreateMethodKlarna          RedirectPaymentMethodCreateMethod = "klarna"
	RedirectPaymentMethodCreateMethodLatitude        RedirectPaymentMethodCreateMethod = "latitude"
	RedirectPaymentMethodCreateMethodLatitudeds      RedirectPaymentMethodCreateMethod = "latitudeds"
	RedirectPaymentMethodCreateMethodLaybuy          RedirectPaymentMethodCreateMethod = "laybuy"
	RedirectPaymentMethodCreateMethodLinepay         RedirectPaymentMethodCreateMethod = "linepay"
	RedirectPaymentMethodCreateMethodLinkaja         RedirectPaymentMethodCreateMethod = "linkaja"
	RedirectPaymentMethodCreateMethodMaybankqrpay    RedirectPaymentMethodCreateMethod = "maybankqrpay"
	RedirectPaymentMethodCreateMethodMercadopago     RedirectPaymentMethodCreateMethod = "mercadopago"
	RedirectPaymentMethodCreateMethodMultibanco      RedirectPaymentMethodCreateMethod = "multibanco"
	RedirectPaymentMethodCreateMethodMultipago       RedirectPaymentMethodCreateMethod = "multipago"
	RedirectPaymentMethodCreateMethodNetbanking      RedirectPaymentMethodCreateMethod = "netbanking"
	RedirectPaymentMethodCreateMethodNupay           RedirectPaymentMethodCreateMethod = "nupay"
	RedirectPaymentMethodCreateMethodNequi           RedirectPaymentMethodCreateMethod = "nequi"
	RedirectPaymentMethodCreateMethodOney10x         RedirectPaymentMethodCreateMethod = "oney_10x"
	RedirectPaymentMethodCreateMethodOney12x         RedirectPaymentMethodCreateMethod = "oney_12x"
	RedirectPaymentMethodCreateMethodOney3x          RedirectPaymentMethodCreateMethod = "oney_3x"
	RedirectPaymentMethodCreateMethodOney4x          RedirectPaymentMethodCreateMethod = "oney_4x"
	RedirectPaymentMethodCreateMethodOney6x          RedirectPaymentMethodCreateMethod = "oney_6x"
	RedirectPaymentMethodCreateMethodOvo             RedirectPaymentMethodCreateMethod = "ovo"
	RedirectPaymentMethodCreateMethodOxxo            RedirectPaymentMethodCreateMethod = "oxxo"
	RedirectPaymentMethodCreateMethodPagoefectivo    RedirectPaymentMethodCreateMethod = "pagoefectivo"
	RedirectPaymentMethodCreateMethodPayid           RedirectPaymentMethodCreateMethod = "payid"
	RedirectPaymentMethodCreateMethodPaymaya         RedirectPaymentMethodCreateMethod = "paymaya"
	RedirectPaymentMethodCreateMethodPaypal          RedirectPaymentMethodCreateMethod = "paypal"
	RedirectPaymentMethodCreateMethodPaypalpaylater  RedirectPaymentMethodCreateMethod = "paypalpaylater"
	RedirectPaymentMethodCreateMethodPayto           RedirectPaymentMethodCreateMethod = "payto"
	RedirectPaymentMethodCreateMethodPayvalida       RedirectPaymentMethodCreateMethod = "payvalida"
	RedirectPaymentMethodCreateMethodPicpay          RedirectPaymentMethodCreateMethod = "picpay"
	RedirectPaymentMethodCreateMethodPix             RedirectPaymentMethodCreateMethod = "pix"
	RedirectPaymentMethodCreateMethodPse             RedirectPaymentMethodCreateMethod = "pse"
	RedirectPaymentMethodCreateMethodRabbitlinepay   RedirectPaymentMethodCreateMethod = "rabbitlinepay"
	RedirectPaymentMethodCreateMethodRapipago        RedirectPaymentMethodCreateMethod = "rapipago"
	RedirectPaymentMethodCreateMethodRazorpay        RedirectPaymentMethodCreateMethod = "razorpay"
	RedirectPaymentMethodCreateMethodRedpagos        RedirectPaymentMethodCreateMethod = "redpagos"
	RedirectPaymentMethodCreateMethodScalapay        RedirectPaymentMethodCreateMethod = "scalapay"
	RedirectPaymentMethodCreateMethodSepa            RedirectPaymentMethodCreateMethod = "sepa"
	RedirectPaymentMethodCreateMethodServipag        RedirectPaymentMethodCreateMethod = "servipag"
	RedirectPaymentMethodCreateMethodShopeepay       RedirectPaymentMethodCreateMethod = "shopeepay"
	RedirectPaymentMethodCreateMethodSingteldash     RedirectPaymentMethodCreateMethod = "singteldash"
	RedirectPaymentMethodCreateMethodSmartpay        RedirectPaymentMethodCreateMethod = "smartpay"
	RedirectPaymentMethodCreateMethodSofort          RedirectPaymentMethodCreateMethod = "sofort"
	RedirectPaymentMethodCreateMethodSpei            RedirectPaymentMethodCreateMethod = "spei"
	RedirectPaymentMethodCreateMethodStitch          RedirectPaymentMethodCreateMethod = "stitch"
	RedirectPaymentMethodCreateMethodStripedd        RedirectPaymentMethodCreateMethod = "stripedd"
	RedirectPaymentMethodCreateMethodStripetoken     RedirectPaymentMethodCreateMethod = "stripetoken"
	RedirectPaymentMethodCreateMethodTapi            RedirectPaymentMethodCreateMethod = "tapi"
	RedirectPaymentMethodCreateMethodTapifintechs    RedirectPaymentMethodCreateMethod = "tapifintechs"
	RedirectPaymentMethodCreateMethodThaiqr          RedirectPaymentMethodCreateMethod = "thaiqr"
	RedirectPaymentMethodCreateMethodTouchngo        RedirectPaymentMethodCreateMethod = "touchngo"
	RedirectPaymentMethodCreateMethodTruemoney       RedirectPaymentMethodCreateMethod = "truemoney"
	RedirectPaymentMethodCreateMethodTrustly         RedirectPaymentMethodCreateMethod = "trustly"
	RedirectPaymentMethodCreateMethodTrustlyeurope   RedirectPaymentMethodCreateMethod = "trustlyeurope"
	RedirectPaymentMethodCreateMethodUpi             RedirectPaymentMethodCreateMethod = "upi"
	RedirectPaymentMethodCreateMethodVenmo           RedirectPaymentMethodCreateMethod = "venmo"
	RedirectPaymentMethodCreateMethodVipps           RedirectPaymentMethodCreateMethod = "vipps"
	RedirectPaymentMethodCreateMethodWaave           RedirectPaymentMethodCreateMethod = "waave"
	RedirectPaymentMethodCreateMethodWebpay          RedirectPaymentMethodCreateMethod = "webpay"
	RedirectPaymentMethodCreateMethodWechat          RedirectPaymentMethodCreateMethod = "wechat"
	RedirectPaymentMethodCreateMethodYape            RedirectPaymentMethodCreateMethod = "yape"
	RedirectPaymentMethodCreateMethodZippay          RedirectPaymentMethodCreateMethod = "zippay"
)

func (e RedirectPaymentMethodCreateMethod) ToPointer() *RedirectPaymentMethodCreateMethod {
	return &e
}

// RedirectPaymentMethodCreate - Create a transaction for an APM/LPM that requires a redirect.
type RedirectPaymentMethodCreate struct {
	// The method to use, this can be any of the methods that support redirect requests.
	Method RedirectPaymentMethodCreateMethod `json:"method"`
	// The `id` of a stored buyer to use Use this instead of the `buyer_external_identifier`.
	BuyerID *string `json:"buyer_id,omitempty"`
	// The `external_identifier` of a stored buyer to use. Use this instead of the `buyer_id`.
	BuyerExternalIdentifier *string `json:"buyer_external_identifier,omitempty"`
	// The 2-letter ISO code of the country to use this payment method for. This is used to select the payment service to use.
	Country string `json:"country"`
	// The ISO-4217 currency code to use this payment method for. This is used to select the payment service to use.
	Currency string `json:"currency"`
	// The redirect URL to redirect a buyer to after they have authorized the payment method.
	RedirectURL string `json:"redirect_url"`
	// The merchant identifier for this payment method.
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
}

func (o *RedirectPaymentMethodCreate) GetMethod() RedirectPaymentMethodCreateMethod {
	if o == nil {
		return RedirectPaymentMethodCreateMethod("")
	}
	return o.Method
}

func (o *RedirectPaymentMethodCreate) GetBuyerID() *string {
	if o == nil {
		return nil
	}
	return o.BuyerID
}

func (o *RedirectPaymentMethodCreate) GetBuyerExternalIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.BuyerExternalIdentifier
}

func (o *RedirectPaymentMethodCreate) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *RedirectPaymentMethodCreate) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *RedirectPaymentMethodCreate) GetRedirectURL() string {
	if o == nil {
		return ""
	}
	return o.RedirectURL
}

func (o *RedirectPaymentMethodCreate) GetExternalIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ExternalIdentifier
}
