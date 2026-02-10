package main

import (
	"database/sql"
	"demo/interfaces"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DatabaseAdapter 适配器，将GORM转换为接口
type DatabaseAdapter struct {
	DB *gorm.DB
}

// NewDatabaseAdapter 创建新的数据库适配器
func NewDatabaseAdapter(db *gorm.DB) *DatabaseAdapter {
	return &DatabaseAdapter{DB: db}
}

// InsertTomatoInfo 插入番茄信息
func (adapter *DatabaseAdapter) InsertTomatoInfo(info interfaces.TomatoInfo) error {
	query := `
		INSERT INTO tomato_info (
			tomato_id, production_area, harvest_date, transport_info, 
			harvest_image, sampling_image, created_at, updated_at,
			harvest_evaluation, sampling_evaluation, processing_date
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	var harvestImage, samplingImage []byte
	if info.HarvestImage != nil {
		harvestImage = info.HarvestImage
	}
	if info.SamplingImage != nil {
		samplingImage = info.SamplingImage
	}

	// 处理可能为零值的时间字段
	processingDate := sql.NullTime{
		Valid: !info.ProcessingDate.IsZero(),
		Time:  info.ProcessingDate,
	}
	createdAt := sql.NullTime{
		Valid: !info.CreatedAt.IsZero(),
		Time:  info.CreatedAt,
	}
	updatedAt := sql.NullTime{
		Valid: !info.UpdatedAt.IsZero(),
		Time:  info.UpdatedAt,
	}

	// 对于 HarvestDate，如果为零值则使用当前时间
	harvestDate := info.HarvestDate
	if harvestDate.IsZero() {
		harvestDate = time.Now()
	}

	_, err := adapter.DB.DB().Exec(query,
		info.TomatoID,
		info.ProductionArea,
		harvestDate,
		info.TransportInfo,
		harvestImage,
		samplingImage,
		createdAt,
		updatedAt,
		info.HarvestEvaluation,
		info.SamplingEvaluation,
		processingDate,
	)
	return err
}

// Exec 执行SQL语句
func (adapter *DatabaseAdapter) Exec(query string, args ...interface{}) (sql.Result, error) {
	result := adapter.DB.Exec(query, args...)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.CommonDB().Exec(query, args...) // 返回标准库sql.Result
}

// Query 执行查询
func (adapter *DatabaseAdapter) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return adapter.DB.DB().Query(query, args...)
}

// GetUserRoleByUID 根据UID获取用户角色
func (adapter *DatabaseAdapter) GetUserRoleByUID(uid string) (string, error) {
	var role string
	query := "SELECT role FROM role WHERE Uid = ?"
	row := adapter.DB.DB().QueryRow(query, uid)
	err := row.Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}

// GetUserRoleAndAddressByUID 根据UID获取用户角色和地址
func (adapter *DatabaseAdapter) GetUserRoleAndAddressByUID(uid string) (string, string, error) {
	log.Printf("DEBUG: GetUserRoleAndAddressByUID called with uid='%s'", uid)
	var role, address string
	query := "SELECT role, address FROM role WHERE Uid = ?"
	row := adapter.DB.DB().QueryRow(query, uid)
	log.Printf("DEBUG: Executing query: %s with param='%s'", query, uid)
	err := row.Scan(&role, &address)
	if err != nil {
		log.Printf("ERROR: Database query failed for uid='%s': %v", uid, err)
		return "", "", err
	}
	log.Printf("DEBUG: Query success - role='%s', address='%s'", role, address)
	return role, address, nil
}

// GetAllRoleAddresses 获取所有角色及其对应地址的映射
func (adapter *DatabaseAdapter) GetAllRoleAddresses() (map[string]string, error) {
	log.Printf("DEBUG: GetAllRoleAddresses called")
	query := "SELECT role, address FROM role"
	rows, err := adapter.DB.DB().Query(query)
	if err != nil {
		log.Printf("ERROR: Failed to query all role addresses: %v", err)
		return nil, err
	}
	defer rows.Close()

	roleAddresses := make(map[string]string)
	for rows.Next() {
		var role, address string
		err := rows.Scan(&role, &address)
		if err != nil {
			log.Printf("ERROR: Failed to scan row in GetAllRoleAddresses: %v", err)
			return nil, err
		}
		roleAddresses[role] = address
		log.Printf("DEBUG: Found role '%s' with address '%s'", role, address)
	}

	if err = rows.Err(); err != nil {
		log.Printf("ERROR: Rows error in GetAllRoleAddresses: %v", err)
		return nil, err
	}

	log.Printf("DEBUG: Retrieved %d role-address mappings", len(roleAddresses))
	return roleAddresses, nil
}

// QueryRow 执行查询并返回一行
func (adapter *DatabaseAdapter) QueryRow(query string, args ...interface{}) *sql.Row {
	return adapter.DB.DB().QueryRow(query, args...)
}

// GetTomatoInfoByID 根据ID获取番茄信息
func (adapter *DatabaseAdapter) GetTomatoInfoByID(id string) (*interfaces.TomatoInfo, error) {
	query := `
		SELECT 
			tomato_id, production_area, harvest_date, transport_info,
			harvest_image, sampling_image, created_at, updated_at,
			harvest_evaluation, sampling_evaluation, processing_date
		FROM tomato_info 
		WHERE tomato_id = ?
	`
	row := adapter.DB.DB().QueryRow(query, id)

	var info interfaces.TomatoInfo
	var harvestImage, samplingImage []byte
	var createdAt, updatedAt, processingDate sql.NullTime

	err := row.Scan(
		&info.TomatoID,
		&info.ProductionArea,
		&info.HarvestDate,
		&info.TransportInfo,
		&harvestImage,
		&samplingImage,
		&createdAt,
		&updatedAt,
		&info.HarvestEvaluation,
		&info.SamplingEvaluation,
		&processingDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tomato with id %s not found", id)
		}
		return nil, err
	}

	info.HarvestImage = harvestImage
	info.SamplingImage = samplingImage
	if createdAt.Valid {
		info.CreatedAt = createdAt.Time
	} else {
		info.CreatedAt = time.Now()
	}
	if updatedAt.Valid {
		info.UpdatedAt = updatedAt.Time
	} else {
		info.UpdatedAt = time.Now()
	}
	if processingDate.Valid {
		info.ProcessingDate = processingDate.Time
	} else {
		info.ProcessingDate = time.Time{}
	}

	return &info, nil
}

// GetAllTomatoIDs 获取所有番茄ID列表
func (adapter *DatabaseAdapter) GetAllTomatoIDs() ([]string, error) {
	query := "SELECT tomato_id FROM tomato_info"
	rows, err := adapter.DB.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tomatoIDs []string
	for rows.Next() {
		var tomatoID string
		err := rows.Scan(&tomatoID)
		if err != nil {
			return nil, err
		}
		tomatoIDs = append(tomatoIDs, tomatoID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tomatoIDs, nil
}

// Close 关闭数据库连接
func (adapter *DatabaseAdapter) Close() error {
	return adapter.DB.Close()
}

// UpdateHarvestImageOnly 仅更新收获图片字段
func (adapter *DatabaseAdapter) UpdateHarvestImageOnly(tomatoID string, harvestImage []byte) error {
	query := `UPDATE tomato_info SET harvest_image = ?, updated_at = ? WHERE tomato_id = ?`
	_, err := adapter.DB.DB().Exec(query, harvestImage, time.Now(), tomatoID)
	if err != nil {
		return fmt.Errorf("failed to update harvest image: %v", err)
	}
	return nil
}

// UpdateSamplingImageOnly 仅更新采样图片字段
func (adapter *DatabaseAdapter) UpdateSamplingImageOnly(tomatoID string, samplingImage []byte) error {
	query := `UPDATE tomato_info SET sampling_image = ?, updated_at = ? WHERE tomato_id = ?`
	_, err := adapter.DB.DB().Exec(query, samplingImage, time.Now(), tomatoID)
	if err != nil {
		return fmt.Errorf("failed to update sampling image: %v", err)
	}
	return nil
}

// UpdateHarvestEvaluationOnly 仅更新收获评估字段
func (adapter *DatabaseAdapter) UpdateHarvestEvaluationOnly(tomatoID string, harvestEvaluation string) error {
	query := `UPDATE tomato_info SET harvest_evaluation = ?, updated_at = ? WHERE tomato_id = ?`
	_, err := adapter.DB.DB().Exec(query, harvestEvaluation, time.Now(), tomatoID)
	if err != nil {
		return fmt.Errorf("failed to update harvest evaluation: %v", err)
	}
	return nil
}

// UpdateSamplingEvaluationOnly 仅更新采样评估字段
func (adapter *DatabaseAdapter) UpdateSamplingEvaluationOnly(tomatoID string, samplingEvaluation string) error {
	query := `UPDATE tomato_info SET sampling_evaluation = ?, updated_at = ? WHERE tomato_id = ?`
	_, err := adapter.DB.DB().Exec(query, samplingEvaluation, time.Now(), tomatoID)
	if err != nil {
		return fmt.Errorf("failed to update sampling evaluation: %v", err)
	}
	return nil
}

// UpdateTransportInfoOnly 仅更新运输信息字段
func (adapter *DatabaseAdapter) UpdateTransportInfoOnly(tomatoID string, transportInfo string) error {
	query := `UPDATE tomato_info SET transport_info = ?, updated_at = ? WHERE tomato_id = ?`
	_, err := adapter.DB.DB().Exec(query, transportInfo, time.Now(), tomatoID)
	if err != nil {
		return fmt.Errorf("failed to update transport info: %v", err)
	}
	return nil
}

// UpdateProcessingInfoOnly 仅更新加工信息字段
func (adapter *DatabaseAdapter) UpdateProcessingInfoOnly(tomatoID string, processingDateStr string) error {
	query := `UPDATE tomato_info SET processing_date = ?, updated_at = ? WHERE tomato_id = ?`
	processingDate, err := time.Parse("2006-01-02", processingDateStr)
	if err != nil {
		// 如果日期格式不正确，使用当前时间
		log.Printf("Warning: Invalid processing date format '%s': %v", processingDateStr, err)
		processingDate = time.Now()
	}
	_, err = adapter.DB.DB().Exec(query, processingDate, time.Now(), tomatoID)
	if err != nil {
		return fmt.Errorf("failed to update processing info: %v", err)
	}
	return nil
}
