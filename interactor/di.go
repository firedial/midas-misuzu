package interactor

import (
    "github.com/firedial/midas-go/dao"
)

var attributeRepository repository.AttributeRepository = &dao.MysqlAttributeRepository{}
var balanceRepository repository.BalanceRepository = &dao.MysqlBalanceRepository{}
var sumRepository repository.SumRepository = &dao.MysqlSumRepository{}


