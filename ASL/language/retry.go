package language

type RetryItem struct {
    ErrorEquals []string
    // Interval between retry (seconds)
    Interval uint32
    BackoffRates uint32
    MaxAttempts uint32
}

type Retry struct {
    On []RetryItem
}
