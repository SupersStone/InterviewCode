package scanner

import (
	"dao-exchange/apps/scanner/repo"
	"dao-exchange/pkg/myerr"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// ContractsCache map
type ContractsCache struct {
	cache   sync.Map
	chain   string
	chainID int
	dao     *repo.Dao
}

// NewContractsCache new cache
func NewContractsCache(chain string, chainID int, dao *repo.Dao) *ContractsCache {
	return &ContractsCache{
		cache:   sync.Map{},
		chain:   chain,
		chainID: chainID,
		dao:     dao,
	}

}

// StartCacheContracts cache contracts loop
func (c *ContractsCache) StartCacheContracts(stop <-chan struct{}, scanInterval time.Duration) {
	myerr.Recover(nil)
	ticker := time.NewTicker(scanInterval)
	for {
		<-ticker.C

		select {
		case <-stop:
			logrus.Info("cache support contracts task get stop signal, exit")
			return
		default:
			if err := c.CacheContracts(); err != nil {
				if errors.Is(err, myerr.ErrRecordNotFound) {
					logrus.Warnln(err)
					continue
				}
				logrus.Error(err)
			}
		}
	}
}

// Contains include something
func (c *ContractsCache) Contains(address common.Address) bool {
	v, ok := c.cache.Load(address)
	if !ok {
		return false
	}

	return v.(bool)
}

// CacheContracts isStart 启动时为true 不执行合约初始化扫描
func (c *ContractsCache) CacheContracts() error {
	contracts, err := c.dao.QueryContracts(c.chain, c.chainID)
	if err != nil {
		return errors.Wrapf(err, "query supported contracts failed, chain: %s, chain_id %d", c.chain, c.chainID)
	}

	addCount := 0
	for _, contract := range contracts {
		if !common.IsHexAddress(contract.Address) {
			logrus.Errorf("address: %s is in valid address, chain: %s", contract.Address, contract.Chain)
			continue
		}

		_, loaded := c.cache.LoadOrStore(common.HexToAddress(contract.Address), true)
		if !loaded {
			addCount++
			logrus.Infof("add new supported contract address to cache, address: %s", contract.Address)
		}
	}

	logrus.Infof("cache supported contract address sucess, contract amount: %d", addCount)

	return nil
}
