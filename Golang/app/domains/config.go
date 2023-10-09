package domains

const (
	MaxProcessCPU      = 2
	MinimumTimeCache   = 5
	DefaultTimeCache   = 10
	MiddleTimeCache    = 30
	MaxTimeCache       = 60
	MaxOpenConnection  = 100
	MaxIddleConnection = 10
	OneMinuteCache     = 1
	OneDayCache        = 1440
)

const DateISOFormat = "2006-01-02"
const DayInSeconds = 86400
const NO_INDEX = 0
const (
	MagazineInt  = 1
	BookInt      = 2
	NewspaperInt = 3
	BonusItemInt = 4
	AudiobookInt = 5
)

const (
	ReceiptTestToProd = 21007
	ReceiptProdToTest = 21008
)

var (
	Release = "release"
	Staging = "test"
)
