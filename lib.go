package fakeuseragent

import (
	"fmt"
	"math/rand"
	"time"
)

type Product int

const (
	MozillaProduct Product = iota
	OperaProduct
)

func (p Product) String() string {
	switch p {
	case MozillaProduct:
		return "Mozilla/5.0"
	case OperaProduct:
		return "Opera/9.80"
	default:
		return ""
	}
}

type UserAgent struct {
	product Product
	osVer   Devices
	browser Browser
}

func (ua UserAgent) String() string {
	return fmt.Sprintf("%s %s %s", ua.product, ua.osVer, ua.browser)
}

func RandomUserAgent() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return UserAgent{
		product: MozillaProduct,
		osVer:   RandomDevice(rng),
		browser: RandomBrowser(rng),
	}.String()
}

func MobileUserAgent() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return UserAgent{
		product: MozillaProduct,
		osVer:   MobileDeviceFunc(rng),
		browser: RandomBrowser(rng),
	}.String()
}

func DesktopUserAgent() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return UserAgent{
		product: MozillaProduct,
		osVer:   PCDeviceFunc(rng),
		browser: RandomBrowser(rng),
	}.String()
}

func CustomUserAgent(osVer Devices, browser Browser) string {
	return UserAgent{
		product: MozillaProduct,
		osVer:   osVer,
		browser: browser,
	}.String()
}
