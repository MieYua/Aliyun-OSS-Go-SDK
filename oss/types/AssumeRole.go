package types

type AssumeRole struct {
	RoleArn         string
	RoleSessionName string
	Policy          SecurityTokenJSON
	DurationSeconds int
}
