<template>
  <div class="tomato-management">
    <h2>番茄信息管理</h2>

    <div class="management-tabs">
      <button
          v-for="tab in managementTabs"
          :key="tab.id"
          :class="{ active: currentManagementTab === tab.id }"
          @click="currentManagementTab = tab.id"
      >
        {{ tab.name }}
      </button>
    </div>

    <div class="management-content">
      <!-- 添加番茄信息 -->
      <div v-if="currentManagementTab === 'add'" class="operation-panel">
        <h3>添加番茄信息</h3>
        <div class="form-grid">
          <div class="form-group">
            <label for="uid">用户UID:</label>
            <input v-model="addFormData.uid" id="uid" type="text" placeholder="" />
          </div>
          <div class="form-group">
            <label for="productionArea">产地:</label>
            <input v-model="addFormData.production_area" id="productionArea" type="text" required />
          </div>
          <div class="form-group">
            <label for="harvestDate">
              收获日期：<span class="date-format-hint">(YYYY-MM-DD)</span>
            </label>
            <input
                v-model="addFormData.harvest_date"
                id="harvestDate"
                type="text"
                required
                class="date-input"
            />
          </div>
        </div>
        <button @click="addTomato" :disabled="loading" class="action-btn">
          {{ loading ? '添加中...' : '添加番茄信息' }}
        </button>
      </div>

      <!-- 更新运输信息 -->
      <div v-if="currentManagementTab === 'transport'" class="operation-panel">
        <h3>更新运输信息</h3>
        <div class="form-grid">
          <div class="form-group">
            <label for="transportTomatoId">番茄ID:</label>
            <input v-model="transportFormData.tomato_id" id="transportTomatoId" type="text" required />
          </div>
          <div class="form-group">
            <label for="transportUid">用户UID:</label>
            <input v-model="transportFormData.uid" id="transportUid" type="text" placeholder="" />
          </div>
          <div class="form-group">
            <label for="transportInfo">运输信息:</label>
            <textarea v-model="transportFormData.transport_info" id="transportInfo" required></textarea>
          </div>
        </div>
        <button @click="updateTransport" :disabled="loading" class="action-btn">
          {{ loading ? '更新中...' : '更新运输信息' }}
        </button>
      </div>

      <!-- 更新加工信息 -->
      <div v-if="currentManagementTab === 'processing'" class="operation-panel">
        <h3>更新加工信息</h3>
        <div class="form-grid">
          <div class="form-group">
            <label for="processingTomatoId">番茄ID:</label>
            <input v-model="processingFormData.tomato_id" id="processingTomatoId" type="text" required />
          </div>
          <div class="form-group">
            <label for="processingUid">用户UID:</label>
            <input v-model="processingFormData.uid" id="processingUid" type="text" placeholder="" />
          </div>
          <div class="form-group">
            <label for="processingDate">
              加工日期：<span class="date-format-hint">(YYYY-MM-DD)</span>
            </label>
            <input
                v-model="processingFormData.processing_date"
                id="processingDate"
                type="text"
                required
                class="date-input"
            />
          </div>
        </div>
        <button @click="updateProcessing" :disabled="loading" class="action-btn">
          {{ loading ? '更新中...' : '更新加工信息' }}
        </button>
      </div>
    </div>

    <!-- 成功提示区域 -->
    <div v-if="result" class="result-success">
      <h3>{{ result.message }}</h3>
      <div v-if="result.tomato_id" class="detail-item">
        <span class="label">番茄ID:</span>
        <span class="value">{{ result.tomato_id }}</span>
        <button @click="copyToClipboard(result.tomato_id)" class="copy-btn">📋 复制 ID</button>
      </div>
    </div>

    <!-- 错误提示区域 -->
    <div v-if="error" class="error-message">{{ error }}</div>
  </div>
</template>

<script>
export default {
  name: 'TomatoManagement',
  data() {
    return {
      currentManagementTab: 'add',
      managementTabs: [
        { id: 'add', name: '添加信息' },
        { id: 'transport', name: '运输信息' },
        { id: 'processing', name: '加工信息' }
      ],
      addFormData: {
        uid: '',
        production_area: '',
        harvest_date: ''
      },
      transportFormData: {
        tomato_id: '',
        transport_info: '',
        uid: ''
      },
      processingFormData: {
        tomato_id: '',
        processing_date: '',
        uid: ''
      },
      result: null,
      error: null,
      loading: false
    };
  },
  watch: {
    // 👇 切换 Tab 时清空提示
    currentManagementTab() {
      this.result = null;
      this.error = null;
    }
  },
  methods: {
    async addTomato() {
      const { production_area, harvest_date } = this.addFormData;
      const dateRegex = /^\d{4}-\d{2}-\d{2}$/;

      if (!production_area) {
        this.error = '产地为必填项';
        return;
      }
      if (!harvest_date) {
        this.error = '收获日期为必填项';
        return;
      }
      if (!dateRegex.test(harvest_date)) {
        this.error = '收获日期格式错误，请使用 YYYY-MM-DD 格式（例如：2026-01-01）';
        return;
      }

      this.loading = true;
      this.error = null;
      this.result = null;

      try {
        const response = await fetch('http://localhost:8081/add-tomato', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            ...(this.addFormData.uid && { 'X-User-UID': this.addFormData.uid })
          },
          body: JSON.stringify(this.addFormData)
        });

        if (!response.ok) throw new Error(`HTTP ${response.status}`);
        const resultData = await response.json();
        this.result = {
          ...resultData,
          message: '番茄信息添加成功'
        };
        this.addFormData = { uid: '', production_area: '', harvest_date: '' };
      } catch (err) {
        this.error = `添加失败：${err.message}`;
      } finally {
        this.loading = false;
      }
    },

    async updateTransport() {
      const { tomato_id, transport_info } = this.transportFormData;

      if (!tomato_id) {
        this.error = '番茄ID为必填项';
        return;
      }
      if (!transport_info) {
        this.error = '运输信息为必填项';
        return;
      }

      this.loading = true;
      this.error = null;
      this.result = null;

      try {
        const response = await fetch('http://localhost:8081/tomato/transport', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            ...(this.transportFormData.uid && { 'X-User-UID': this.transportFormData.uid })
          },
          body: JSON.stringify(this.transportFormData)
        });

        if (!response.ok) throw new Error(`HTTP ${response.status}`);
        const resultData = await response.json();
        this.result = {
          ...resultData,
          message: '运输信息更新成功'
        };
      } catch (err) {
        this.error = `更新失败：${err.message}`;
      } finally {
        this.loading = false;
      }
    },

    async updateProcessing() {
      const { tomato_id, processing_date } = this.processingFormData;
      const dateRegex = /^\d{4}-\d{2}-\d{2}$/;

      if (!tomato_id) {
        this.error = '番茄ID为必填项';
        return;
      }
      if (!processing_date) {
        this.error = '加工日期为必填项';
        return;
      }
      if (!dateRegex.test(processing_date)) {
        this.error = '加工日期格式错误，请使用 YYYY-MM-DD 格式（例如：2024-06-15）';
        return;
      }

      this.loading = true;
      this.error = null;
      this.result = null;

      try {
        const response = await fetch('http://localhost:8081/tomato/processing', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            ...(this.processingFormData.uid && { 'X-User-UID': this.processingFormData.uid })
          },
          body: JSON.stringify(this.processingFormData)
        });

        if (!response.ok) throw new Error(`HTTP ${response.status}`);
        const resultData = await response.json();
        this.result = {
          ...resultData,
          message: '加工信息更新成功'
        };
      } catch (err) {
        this.error = `更新失败：${err.message}`;
      } finally {
        this.loading = false;
      }
    },

    copyToClipboard(text) {
      navigator.clipboard.writeText(text).then(() => {
        // 可选：保留原 alert 或改用 Toast；此处暂不处理，避免干扰主提示
        // alert('📋 番茄 ID 已复制到剪贴板！');
      }).catch(err => {
        console.error('复制失败:', err);
        this.error = '❌ 复制失败，请手动复制';
      });
    }
  }
};
</script>

<style scoped lang="scss">
.tomato-management {
  max-width: 1000px;
  margin: 0 auto;
}

.management-tabs {
  display: flex;
  background: #f8f9fa;
  border-radius: 10px 10px 0 0;
  overflow: hidden;
  margin-bottom: 0;
}

.management-tabs button {
  flex: 1;
  padding: 1rem;
  border: none;
  background: transparent;
  cursor: pointer;
  transition: all 0.3s ease;
}

.management-tabs button:hover {
  background: #e9ecef;
}

.management-tabs button.active {
  background: #667eea;
  color: white;
}

.management-content {
  background: white;
  border-radius: 0 0 10px 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  padding: 2rem;
}

.operation-panel h3 {
  color: #667eea;
  margin-bottom: 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 2px solid #667eea;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: 0.5rem;
  font-weight: bold;
  color: #333;
}

.date-format-hint {
  font-size: 0.85rem;
  color: #6c757d;
  font-weight: normal;
  margin-left: 0.25rem;
}

.form-group input,
.form-group textarea {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 1rem;

  &.date-input {
    appearance: none;
    background: white;
    width: 100%;

    &::-webkit-calendar-picker-indicator {
      display: none;
    }
  }
}

.form-group textarea {
  min-height: 100px;
  resize: vertical;
}

.action-btn {
  background: linear-gradient(90deg, #5a67d8, #805ad5);
  color: #ffffff;
  border: none;
  padding: 12px 24px;
  border-radius: 12px;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.3s ease;
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
  margin-top: 2rem;
}

.result-success h3 {
  color: #155724;
  margin-top: 0;
}

.detail-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  background: white;
  border-radius: 5px;
  margin-top: 1rem;
}

.label {
  font-weight: bold;
  color: #155724;
}

.value {
  color: #333;
  word-break: break-all;
}

.copy-btn {
  background: #5a67d8;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  cursor: pointer;
  transition: all 0.3s ease;

  &:hover {
    background: #4c51bf;
  }
}

.error-message {
  background: #f8d7da;
  color: #721c24;
  padding: 1rem;
  border-radius: 5px;
  border: 1px solid #f5c6cb;
  margin-top: 1rem;
}
</style>
