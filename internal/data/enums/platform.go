package enums

type PlatformType int

const (
	PlatformType_Undefined PlatformType = iota
	PlatformType_WholesaleSeller
	PlatformType_WholesaleBuyer
)

var platformType = map[string]PlatformType{
	PlatformType_Undefined.String():       PlatformType_Undefined,
	PlatformType_WholesaleSeller.String(): PlatformType_WholesaleSeller,
	PlatformType_WholesaleBuyer.String():  PlatformType_WholesaleBuyer,
}

func (p PlatformType) String() string {
	return []string{"platform_type_undefined", "platform_type_google", "platform_type_facebook"}[p]
}

//func StringToPlatformType(s string) PlatformType {
//	r, ok := platformType[s]
//	if ok {
//		return r
//	}
//	return PlatformType_Undefined
//}
