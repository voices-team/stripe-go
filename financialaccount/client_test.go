package financialaccount

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v72"
	_ "github.com/stripe/stripe-go/v72/testing"
)

func TestFinancialAccountClose(t *testing.T) {
	account, err := Close("fa_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestFinancialAccountGet(t *testing.T) {
	account, err := Get()
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestFinancialAccountGetByID(t *testing.T) {
	account, err := GetByID("fa_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestFinancialAccountList(t *testing.T) {
	i := List(&stripe.FinancialAccountListParams{})

	// Verify that we can get at least one account
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.FinancialAccount())
	assert.NotNil(t, i.FinancialAccountList())
}

func TestFinancialAccountNew(t *testing.T) {
	account, err := New(&stripe.FinancialAccountParams{
		SupportedCurrencies: []*string{
			stripe.String("usd"),
		},
		Features: &stripe.FinancialAccountFeaturesParams{
			CardIssuing: &stripe.FinancialAccountFeaturesCardIssuingParams{
				Requested: stripe.Bool(true),
			},
			DepositInsurance: &stripe.FinancialAccountFeaturesDepositInsuranceParams{
				Requested: stripe.Bool(true),
			},
			FinancialAddresses: &stripe.FinancialAccountFeaturesFinancialAddressesParams{
				ABA: &stripe.FinancialAccountFeaturesFinancialAddressesABAParams{
					Requested: stripe.Bool(true),
				},
			},
			InboundTransfers: &stripe.FinancialAccountFeaturesInboundTransfersParams{
				ACH: &stripe.FinancialAccountFeaturesInboundTransfersACHParams{
					Requested: stripe.Bool(true),
				},
				Card: &stripe.FinancialAccountFeaturesInboundTransfersCardParams{
					Requested: stripe.Bool(true),
				},
			},
			IntraStripeFlows: &stripe.FinancialAccountFeaturesIntraStripeFlowsParams{
				Requested: stripe.Bool(true),
			},
			OutboundPayments: &stripe.FinancialAccountFeaturesOutboundPaymentsParams{
				ACH: &stripe.FinancialAccountFeaturesInboundTransfersACHParams{
					Requested: stripe.Bool(true),
				},
				UsDomesticWire: &stripe.FinancialAccountFeaturesInboundTransfersUsDomesticWireParams{
					Requested: stripe.Bool(true),
				},
			},
			OutboundTransfers: &stripe.FinancialAccountFeaturesOutboundTransfersParams{
				ACH: &stripe.FinancialAccountFeaturesInboundTransfersACHParams{
					Requested: stripe.Bool(true),
				},
				UsDomesticWire: &stripe.FinancialAccountFeaturesInboundTransfersUsDomesticWireParams{
					Requested: stripe.Bool(true),
				},
			},
		},
		PlatformRestrictions: &stripe.FinancialAccountPlatformRestrictionsParams{
			InboundFlows:  stripe.PlatformRestrictionStatusP(stripe.PlatformRestrictionUnrestricted),
			OutboundFlows: stripe.PlatformRestrictionStatusP(stripe.PlatformRestrictionUnrestricted),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestFinancialAccountUpdate(t *testing.T) {
	account, err := Update("fa_123", &stripe.FinancialAccountParams{
		Features: &stripe.FinancialAccountFeaturesParams{
			CardIssuing: &stripe.FinancialAccountFeaturesCardIssuingParams{
				Requested: stripe.Bool(true),
			},
		},
		PlatformRestrictions: &stripe.FinancialAccountPlatformRestrictionsParams{
			OutboundFlows: stripe.PlatformRestrictionStatusP(stripe.PlatformRestrictionRestricted),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, account)
}
