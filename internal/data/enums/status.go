package enums

type StatusType int

const (
	StatustypeEnabled StatusType = iota
	StatustypeDisabled
	StatustypeUndefined
)

var statusType = map[string]StatusType{
	StatustypeEnabled.String():   StatustypeEnabled,
	StatustypeDisabled.String():  StatustypeDisabled,
	StatustypeUndefined.String(): StatustypeUndefined,
}

func (p StatusType) String() string {
	return []string{"status_type_enabled", "status_type_disabled", "status_type_undefined"}[p]
}

func StringToStatusType(s string) StatusType {
	r, ok := statusType[s]
	if ok {
		return r
	}
	return StatustypeUndefined
}
