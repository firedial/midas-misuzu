package interactor

import (
    "github.com/firedial/midas-misuzu/dao"
	"github.com/firedial/midas-misuzu/repository"
)

var attributeRepository repository.AttributeRepository = &dao.MysqlAttributeRepository{}
var balanceRepository repository.BalanceRepository = &dao.MysqlBalanceRepository{}
var sumRepository repository.SumRepository = &dao.MysqlSumRepository{}


