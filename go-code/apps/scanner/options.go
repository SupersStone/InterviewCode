package scanner

type Option func(*Scanner)

func WithStartBlock(startBlock uint64) Option {
	return func(bs *Scanner) {
		bs.startBlock = startBlock
	}
}

func WithDelayBlock(delayBlock uint64) Option {
	return func(bs *Scanner) {
		bs.delayBlock = delayBlock
	}
}

func WithScanInterval(interval int) Option {
	return func(bs *Scanner) {
		bs.scanInterval = interval
	}
}

func WithScanAmount(scanAmount int) Option {
	return func(bs *Scanner) {
		bs.scanAmount = scanAmount
	}
}

func WithChainID(chainID int) Option {
	return func(bs *Scanner) {
		bs.chainID = chainID
	}
}
