package constants

type Env int64

const (
	Undefined Env = iota
	Development
	Test
	Staging
	Production
)

func (e Env) String() string {
	switch e {
	case Development:
		return "development"
	case Test:
		return "test"
	case Staging:
		return "staging"
	case Production:
		return "production"
	default:
		return "undefined"
	}
}

func IsDevelopment(env string) bool {
	return env == Development.String()
}

func IsTest(env string) bool {
	return env == Test.String()
}

func IsStaging(env string) bool {
	return env == Staging.String()
}

func IsProduction(env string) bool {
	return env == Production.String()
}
