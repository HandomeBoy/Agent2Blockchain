<template>
  <div class="tomato-query">
    <h2>番茄ID列表</h2>

    <div class="query-form">
      <div class="form-group">
        <label>总数：</label>
        <span class="value">{{ count }}</span>
      </div>
      <button @click="fetchIds" :disabled="loading" class="query-btn">
        {{ loading ? '加载中...' : '刷新列表' }}
      </button>
    </div>

    <div v-if="ids.length > 0" class="result-container">
      <h3>全部番茄ID（{{ count }} 个）</h3>
      <div class="result-card">
        <div class="info-grid">
          <div
              v-for="(id, index) in ids"
              :key="index"
              class="info-item full-width"
          >
            <span class="label">番茄ID {{ index + 1 }}：</span>
            <span class="value pre-wrap">{{ id }}</span>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="loading" class="result-container">
      <h3>加载中...</h3>
      <div class="result-card" style="text-align: center; padding: 2rem;">
        <div class="spinner">●●●</div>
      </div>
    </div>

    <div v-else class="result-container">
      <h3>无数据</h3>
      <div class="result-card" style="text-align: center; padding: 2rem;">
        暂无番茄ID记录
      </div>
    </div>

    <div v-if="error" class="error-message">
      {{ error }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'TomatoIdList',
  data() {
    return {
      ids: [],
      count: 0,
      error: null,
      loading: false
    };
  },
  methods: {
    async fetchIds() {
      this.loading = true;
      this.error = null;
      this.ids = [];
      try {
        const response = await fetch('http://localhost:8081/tomato-ids');
        if (!response.ok) throw new Error(`HTTP ${response.status}`);
        const data = await response.json();
        this.count = data.count || 0;
        this.ids = Array.isArray(data.tomato_ids) ? data.tomato_ids : [];
      } catch (err) {
        this.error = `加载失败: ${err.message}`;
      } finally {
        this.loading = false;
      }
    }
  },
  mounted() {
    this.fetchIds();
  }
};
</script>

<style scoped lang="scss">
.tomato-query {
  max-width: 800px;
  margin: 0 auto;
  background: #f5f3ff;
  border-radius: 12px;
  padding: 1.5rem;
}

.query-form {
  background: #ffffff;
  padding: 2.2rem;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(74, 108, 247, 0.08);
  margin-bottom: 2rem;
  border: 1px solid #e6ebf5;

  .form-group {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1.6rem;
  }

  label {
    font-weight: 600;
    color: #2c3e50;
    font-size: 0.95rem;
  }

  .value {
    background: #f8fbff;
    padding: 0.5rem 1rem;
    border-radius: 6px;
    font-family: monospace;
    font-size: 0.9rem;
    word-break: break-all;
  }

  .query-btn {
    background: linear-gradient(90deg, #5a67d8, #805ad5);
    color: #ffffff;
    border: none;
    padding: 12px 24px;
    border-radius: 12px;
    font-size: 16px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s ease;
    width: auto;
    box-shadow: 0 4px 8px rgba(90, 103, 216, 0.2);

    &:hover:not(:disabled) {
      transform: translateY(-3px);
      box-shadow: 0 6px 16px rgba(90, 103, 216, 0.35);
    }

    &:disabled {
      opacity: 0.6;
      cursor: not-allowed;
      transform: none;
      box-shadow: none;
    }
  }
}

.result-container {
  background: #ffffff;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(74, 108, 247, 0.08);
  overflow: hidden;
  border: 1px solid #e6ebf5;

  h3 {
    padding: 1.6rem 1.8rem;
    background: #f9fbff;
    margin: 0;
    font-size: 1.4rem;
    font-weight: 600;
    color: #2c3e50;
    border-bottom: 1px solid #e6ebf5;
  }

  .result-card {
    padding: 1.8rem;
  }

  .info-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 1.3rem;
  }

  .info-item {
    display: flex;
    flex-direction: column;
    padding: 0.8rem 1rem;
    background: #f8fbff;
    border-radius: 8px;
    border: 1px solid #e6ebf5;

    &.full-width {
      grid-column: 1 / -1;
      padding: 0.9rem 1rem;
    }

    .label {
      font-weight: 600;
      color: #5d6b82;
      margin-bottom: 0.4rem;
      font-size: 0.95rem;
    }

    .value {
      color: #2c3e50;
      font-size: 1rem;
      line-height: 1.5;
      white-space: pre-wrap;
      word-break: break-word;
      font-family: monospace;
      background: #f0f4f8;
      padding: 0.6rem;
      border-radius: 4px;
    }

    .pre-wrap {
      white-space: pre-wrap;
      font-family: 'Segoe UI', system-ui, sans-serif;
    }
  }

  .error-message {
    background: #fff8f8;
    color: #dc3545;
    padding: 1.2rem;
    border-radius: 8px;
    border: 1px solid #f8d7da;
    margin-top: 1.5rem;
    font-size: 0.95rem;
  }

  .spinner {
    font-size: 1.2rem;
    color: #5a67d8;
    animation: blink 1.2s infinite;
  }

  @keyframes blink {
    50% { opacity: 0.3; }
  }
}
</style>
