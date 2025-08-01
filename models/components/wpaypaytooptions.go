// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type WpayPaytoOptions struct {
	// Options to pass to the `instrument` resource in the Wpay PayTo API.
	Instrument *WpayPaytoResourceOptions `json:"instrument,omitempty"`
	// Options to pass to the `payment` resource in the Wpay PayTo API.
	Payment *WpayPaytoResourceOptions `json:"payment,omitempty"`
	// Options to pass to the `refund` resource in the Wpay PayTo API.
	Refund *WpayPaytoResourceOptions `json:"refund,omitempty"`
}

func (o *WpayPaytoOptions) GetInstrument() *WpayPaytoResourceOptions {
	if o == nil {
		return nil
	}
	return o.Instrument
}

func (o *WpayPaytoOptions) GetPayment() *WpayPaytoResourceOptions {
	if o == nil {
		return nil
	}
	return o.Payment
}

func (o *WpayPaytoOptions) GetRefund() *WpayPaytoResourceOptions {
	if o == nil {
		return nil
	}
	return o.Refund
}
