package adaptivecard

type DataQuery struct {
	Dataset string `json:"dataset"`
	Count   int    `json:"count"`
	Skip    int    `json:"skip"`
}

func NewDataQuery(dataset string, opts ...func(*DataQuery)) *DataQuery {
	dq := &DataQuery{
		Dataset: dataset,
	}

	for _, o := range opts {
		o(dq)
	}

	return dq
}

func WithDataQueryCount(c int) func(*DataQuery) {
	return func(dq *DataQuery) {
		dq.Count = c
	}
}

func WithDataQuerySkip(s int) func(*DataQuery) {
	return func(dq *DataQuery) {
		dq.Skip = s
	}
}
