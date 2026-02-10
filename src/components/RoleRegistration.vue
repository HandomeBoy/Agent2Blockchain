<template>
  <div class="role-registration">
    <h2>角色注册</h2>

    <div class="registration-form">
      <div class="form-group">
        <label for="roleType">角色类型:</label>
        <select v-model="formData.role_type" id="roleType">
          <option value="">请选择角色类型</option>
          <option value="Consumer">消费者</option>
          <option value="Grower">种植户</option>
          <option value="Logistic">物流公司</option>
          <option value="Factory">工厂</option>
          <option value="QualityExpert">质量专家</option>
        </select>
      </div>

      <button @click="registerRole" :disabled="loading" class="register-btn">
        {{ loading ? '注册中...' : '注册角色' }}
      </button>
    </div>

    <div v-if="result" class="result-success">
      <h3>注册成功!</h3>
      <div class="success-details">
        <div class="detail-item">
          <span class="label">用户ID:</span>
          <span class="value">{{ result.uid }}</span>
        </div>
        <div class="detail-item">
          <span class="label">角色类型:</span>
          <span class="value">{{ result.role_type }}</span>
        </div>
        <div class="detail-item">
          <span class="label">区块链地址:</span>
          <span class="value address">{{ result.address }}</span>
        </div>
        <div class="detail-item">
          <span class="label">消息:</span>
          <span class="value">{{ result.message }}</span>
        </div>
      </div>
    </div>

    <div v-if="error" class="error-message">
      {{ error }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'RoleRegistration',
  data() {
    return {
      formData: {
        role_type: ''
      },
      result: null,
      error: null,
      loading: false
    }
  },
  methods: {
    async registerRole() {
      if (!this.formData.role_type) {
        this.error = '请选择角色类型'
        return
      }

      this.loading = true
      this.error = null
      this.result = null

      try {
        const response = await fetch('http://localhost:8081/role/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(this.formData)
        })

        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || `HTTP error! status: ${response.status}`)
        }

        this.result = await response.json()
      } catch (err) {
        this.error = `注册失败: ${err.message}`
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped lang="scss">
.role-registration {
  max-width: 600px;
  margin: 0 auto;
}

.registration-form {
  background: white;
  padding: 2rem;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
  color: #333;
}

.form-group select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 1rem;
}

.register-btn {
  background: linear-gradient(90deg, #5a67d8, #805ad5);
  color: #ffffff;
  border: none;
  padding: 12px 24px;
  border-radius: 12px;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.3s ease;
  width: 100%;
  box-shadow: 0 4px 8px rgba(90, 103, 216, 0.2);

  &:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 6px 12px rgba(90, 103, 216, 0.3);
  }

  &:disabled {
    opacity: 0.6;
    cursor: not-allowed;
    transform: none;
    box-shadow: none;
  }
}

.result-success {
  background: #d4edda;
  border: 1px solid #c3e6cb;
  border-radius: 10px;
  padding: 2rem;
  margin-bottom: 2rem;
}

.result-success h3 {
  color: #155724;
  margin-top: 0;
  text-align: center;
}

.success-details {
  display: grid;
  gap: 1rem;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  padding: 0.75rem;
  background: white;
  border-radius: 5px;
}

.label {
  font-weight: bold;
  color: #155724;
}

.value {
  color: #333;
  word-break: break-all;
}

.address {
  font-family: monospace;
  font-size: 0.9rem;
}

.error-message {
  background: #f8d7da;
  color: #721c24;
  padding: 1rem;
  border-radius: 5px;
  border: 1px solid #f5c6cb;
}
</style>
