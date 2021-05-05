package util

const (
	USD   = "USD"
	EUR   = "EUR"
	POUND = "POUND"
	GHS   = "GHS"
	CAD   = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, GHS:
		return true
	}
	return false
}
