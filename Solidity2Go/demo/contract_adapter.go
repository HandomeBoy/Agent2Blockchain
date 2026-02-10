package main

import (
	"demo/interfaces"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"

	"demo/contract"
)

// ContractAdapter 适配器，将生成的合约实例转换为接口
type ContractAdapter struct {
	Contract *contract.TomatoTraceable
}

// NewContractAdapter 创建新的合约适配器
func NewContractAdapter(contract *contract.TomatoTraceable) *ContractAdapter {
	return &ContractAdapter{Contract: contract}
}

// AddTomatoInfo 添加番茄信息
func (adapter *ContractAdapter) AddTomatoInfo(opts *bind.TransactOpts, productionArea string, harvestDateStr string) (string, *types.Transaction, *types.Receipt, error) {
	return adapter.Contract.AddTomatoInfo(opts, productionArea, harvestDateStr)
}

// UpdateHarvestInfo 更新收获信息
func (adapter *ContractAdapter) UpdateHarvestInfo(opts *bind.TransactOpts, tomatoId string, productionArea string, harvestDateStr string) (*types.Transaction, *types.Receipt, error) {
	return adapter.Contract.UpdateHarvestInfo(opts, tomatoId, productionArea, harvestDateStr)
}

// UpdateTransportInfo 更新运输信息
func (adapter *ContractAdapter) UpdateTransportInfo(opts *bind.TransactOpts, tomatoId string, transportInfo string) (*types.Transaction, *types.Receipt, error) {
	return adapter.Contract.UpdateTransportInfo(opts, tomatoId, transportInfo)
}

// UpdateProcessingInfo 更新加工信息
func (adapter *ContractAdapter) UpdateProcessingInfo(opts *bind.TransactOpts, tomatoId string, processingDateStr string) (*types.Transaction, *types.Receipt, error) {
	return adapter.Contract.UpdateProcessingInfo(opts, tomatoId, processingDateStr)
}

// UpdateQualityInfo 更新质量信息
func (adapter *ContractAdapter) UpdateQualityInfo(opts *bind.TransactOpts, tomatoId string, harvestEvaluation string, samplingEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return adapter.Contract.UpdateQualityInfo(opts, tomatoId, harvestEvaluation, samplingEvaluation)
}

// QueryTomatoInfo 查询番茄信息
func (adapter *ContractAdapter) QueryTomatoInfo(opts *bind.CallOpts, user common.Address, tomatoId string) (interfaces.TomatoTraceableTomatoBaseInfo, interfaces.TomatoTraceableProcessingInfo, interfaces.TomatoTraceableTransportInfo, error) {
	log.Printf("DEBUG: ContractAdapter.QueryTomatoInfo - Now querying database directly for user=%s, tomatoId=%s", user.Hex(), tomatoId)

	return interfaces.TomatoTraceableTomatoBaseInfo{},
		interfaces.TomatoTraceableProcessingInfo{},
		interfaces.TomatoTraceableTransportInfo{},
		fmt.Errorf("method deprecated - use service layer to query database directly")
}

// GetUserRole 获取用户角色
func (adapter *ContractAdapter) GetUserRole(opts *bind.CallOpts, user common.Address) (uint8, error) {
	role, err := adapter.Contract.GetUserRole(opts, user)
	if err != nil {
		return 0, err
	}
	return uint8(role), nil
}

// RegisterRole 注册角色
func (adapter *ContractAdapter) RegisterRole(opts *bind.TransactOpts, role uint8) (*types.Transaction, *types.Receipt, error) {
	// 使用defer-recover机制来处理可能的panic
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered in RegisterRole: %v", r)
		}
	}()

	return adapter.Contract.RegisterRole(opts, role)
}

// GetAllTomatoIds 获取所有番茄ID
func (adapter *ContractAdapter) GetAllTomatoIds(opts *bind.CallOpts) ([]string, error) {
	// 直接调用合约方法
	tomatoIDs, err := adapter.Contract.GetAllTomatoIds(opts)
	if err != nil {
		log.Printf("Error in GetAllTomatoIds: %v", err)
		// 检查是否是ABI解组错误
		if strings.Contains(err.Error(), "unmarshall an empty string") || strings.Contains(err.Error(), "abi:") {
			log.Printf("ABI error detected in GetAllTomatoIds, returning empty array")
			// 如果是ABI错误，返回空数组而不是错误
			return []string{}, nil
		}
		return nil, err
	}

	return tomatoIDs, nil
}

// TomatoBaseInfos 获取番茄基本信息
func (adapter *ContractAdapter) TomatoBaseInfos(opts *bind.CallOpts, arg0 string) (interfaces.TomatoTraceableTomatoBaseInfo, error) {
	info, err := adapter.Contract.TomatoBaseInfos(opts, arg0)
	if err != nil {
		log.Printf("Error in TomatoBaseInfos: %v", err)
		// 检查是否是ABI解组错误
		if strings.Contains(err.Error(), "unmarshall an empty string") || strings.Contains(err.Error(), "abi:") {
			log.Printf("ABI error detected in TomatoBaseInfos, returning default struct")
			// 如果是ABI错误，返回默认结构体
			return interfaces.TomatoTraceableTomatoBaseInfo{
				TomatoId:          arg0,
				ProductionArea:    "",
				HarvestDate:       big.NewInt(0),
				HarvestEvaluation: "",
				Grower:            common.Address{},
			}, nil
		}
		return interfaces.TomatoTraceableTomatoBaseInfo{}, err
	}

	result := interfaces.TomatoTraceableTomatoBaseInfo{
		TomatoId:          info.TomatoId,
		ProductionArea:    info.ProductionArea,
		HarvestDate:       info.HarvestDate,
		HarvestEvaluation: info.HarvestEvaluation,
		Grower:            info.Grower,
	}

	return result, nil
}

// TransportInfos 获取运输信息
func (adapter *ContractAdapter) TransportInfos(opts *bind.CallOpts, arg0 string) (interfaces.TomatoTraceableTransportInfo, error) {
	info, err := adapter.Contract.TransportInfos(opts, arg0)
	if err != nil {
		log.Printf("Error in TransportInfos: %v", err)
		// 检查是否是ABI解组错误
		if strings.Contains(err.Error(), "unmarshall an empty string") || strings.Contains(err.Error(), "abi:") {
			log.Printf("ABI error detected in TransportInfos, returning default struct")
			// 如果是ABI错误，返回默认结构体
			return interfaces.TomatoTraceableTransportInfo{
				TransportInfo: "",
				Logistic:      common.Address{},
			}, nil
		}
		return interfaces.TomatoTraceableTransportInfo{}, err
	}

	result := interfaces.TomatoTraceableTransportInfo{
		TransportInfo: info.TransportInfo,
		Logistic:      info.Logistic,
	}

	return result, nil
}

// UpdateHarvestEvaluation 更新收获评估
func (adapter *ContractAdapter) UpdateHarvestEvaluation(opts *bind.TransactOpts, tomatoId string, harvestEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return adapter.Contract.UpdateHarvestEvaluation(opts, tomatoId, harvestEvaluation)
}

// UpdateSamplingEvaluation 更新采样评估
func (adapter *ContractAdapter) UpdateSamplingEvaluation(opts *bind.TransactOpts, tomatoId string, samplingEvaluation string) (*types.Transaction, *types.Receipt, error) {
	return adapter.Contract.UpdateSamplingEvaluation(opts, tomatoId, samplingEvaluation)
}
