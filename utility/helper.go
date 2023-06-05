package utility

type ResultStorage struct {
	MetricGroup string

	Output interface{}

	Err error
}
