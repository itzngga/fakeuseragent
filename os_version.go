package fakeuseragent

import (
	"fmt"
	"math/rand"
	"time"
)

type Devices interface {
	String() string
}

type DesktopDevice int

const (
	Windows DesktopDevice = iota
	Macintosh
	Linux
)

func (d DesktopDevice) String() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch d {
	case Windows:
		return fmt.Sprintf("(Windows NT %s; %s)", WIN_NT[rng.Intn(len(WIN_NT))], WIN_BIT[rng.Intn(len(WIN_BIT))])
	case Macintosh:
		return fmt.Sprintf("(Macintosh; Intel Mac OS X %s)", MACINTOSH_VER[rng.Intn(len(MACINTOSH_VER))])
	case Linux:
		return fmt.Sprintf("(X11; %s)", LINUX_VER[rng.Intn(len(LINUX_VER))])
	default:
		return ""
	}
}

type MobileDevice int

const (
	Galaxy MobileDevice = iota
	Huawei
	HarmonyOS
	Iphone
	Redmi
	Mi
)

func (m MobileDevice) String() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch m {
	case Galaxy:
		return fmt.Sprintf("(Linux; Android %s; %s)", ANDROID_VER[rng.Intn(len(ANDROID_VER))], GALAXY[rng.Intn(len(GALAXY))])
	case Huawei:
		return fmt.Sprintf("(Linux; Android %s; %s)", ANDROID_VER[rng.Intn(len(ANDROID_VER))], HUAWEI[rng.Intn(len(HUAWEI))])
	case HarmonyOS:
		return fmt.Sprintf("(%s)", HARMONYOS[rng.Intn(len(HARMONYOS))])
	case Iphone:
		return fmt.Sprintf("(iPhone; CPU iPhone OS %s like Mac OS X)", IPHONE[rng.Intn(len(IPHONE))])
	case Mi:
		return fmt.Sprintf("(Linux; Android %s; %s)", ANDROID_VER[rng.Intn(len(ANDROID_VER))], XIAOMI[rng.Intn(len(XIAOMI))])
	case Redmi:
		return fmt.Sprintf("(Linux; Android %s; %s)", ANDROID_VER[rng.Intn(len(ANDROID_VER))], REDMI[rng.Intn(len(REDMI))])
	default:
		return ""
	}
}

type DevicesEnum int

const (
	Desktop DevicesEnum = iota
	Mobile
)

func RandomDevice(rng *rand.Rand) Devices {
	switch rng.Intn(6) {
	case 0, 1, 2:
		return DesktopDevice(rng.Intn(3))
	default:
		return MobileDevice(rng.Intn(6))
	}
}

func MobileDeviceFunc(rng *rand.Rand) Devices {
	return MobileDevice(rng.Intn(6))
}

func PCDeviceFunc(rng *rand.Rand) Devices {
	return DesktopDevice(rng.Intn(3))
}
