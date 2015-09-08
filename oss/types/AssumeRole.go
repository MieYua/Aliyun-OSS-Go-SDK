package types

type AssumeRole struct {
	RoleArn         string `xml:""`
	RoleSessionName string `xml:""`
	Policy          SecurityTokenJSON
	DurationSeconds int
}
