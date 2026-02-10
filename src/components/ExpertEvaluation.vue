<template>
  <div class="tomato-query expert-evaluation">
    <h2>专家评估</h2>

    <div class="query-form">
      <div class="form-group">
        <label for="evalTomatoId">番茄ID:</label>
        <input
            v-model="formData.tomato_id"
            id="evalTomatoId"
            type="text"
            required
            placeholder="请输入番茄ID"
        />
      </div>
      <div class="form-group">
        <label for="evalRole">评估类型:</label>
        <select v-model="formData.role" id="evalRole" required>
          <option value="">请选择评估类型</option>
          <option value="sampling">抽样评估</option>
          <option value="harvest">收获评估</option>
        </select>
      </div>
      <div class="form-group">
        <label for="evalUid">专家UID:</label>
        <input v-model="formData.uid" id="evalUid" type="text" placeholder="可选" />
      </div>

      <button @click="performEvaluation" :disabled="loading" class="query-btn">
        {{ loading ? '评估中...' : '执行评估' }}
      </button>
    </div>

    <div v-if="result" class="result-container">
      <h3>评估结果</h3>
      <div class="result-card">
        <div class="result-section">
          <h4>{{ evaluationTitle }}</h4>
          <div class="info-grid">
            <div class="info-item full-width">
              <span class="label">{{ evaluationTitle }}:</span>
              <div class="harvest-evaluation">
                <template v-if="parsedEvaluation?.sections">
                  <div
                      v-for="(section, idx) in parsedEvaluation.sections"
                      :key="idx"
                      class="evaluation-section"
                      :class="{ 'overall-score': section.isOverall }"
                  >
                    <h5>{{ section.title }}</h5>
                    <ul>
                      <li v-for="(value, key) in section.items" :key="key">
                        <strong>{{ key }}：</strong>
                        <span>{{ formatValue(value) }}</span>
                      </li>
                    </ul>
                  </div>
                </template>
                <span v-else class="value pre-wrap">无评估数据</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="error" class="error-message">
      {{ error }}
    </div>
  </div>
</template>

<script>

const _parseReportText = (text) => {
  const lines = String(text).split('\n').filter(l => l.trim());
  const sections = [];
  let currentSection = null;

  for (const line of lines) {
    const titleMatch = line.match(/^\*\*(.+?)\*\*$/) || line.match(/^([^：]+)：\s*$/);
    if (titleMatch) {
      const title = (titleMatch[1] || titleMatch[0]).replace(/\*\*/g, '').trim();
      currentSection = { title, items: {} };
      sections.push(currentSection);
    } else if (currentSection && line.includes('：')) {
      const [key, ...rest] = line.split('：');
      const value = rest.join('：').trim();
      if (key && value !== undefined) {
        currentSection.items[key.trim()] = value;
      }
    }
  }

  const overallItem = sections.flatMap(s => Object.entries(s.items))
      .find(([k]) => k.includes('整体评价') || k.includes('评分') || k.includes('总分'));
  if (overallItem) {
    const [key, value] = overallItem;
    sections.push({
      title: '整体评价',
      items: { '评分': value },
      isOverall: true
    });
    sections.forEach(s => delete s.items[key]);
  }

  return { sections };
};

export default {
  name: 'ExpertEvaluation',
  data() {
    return {
      formData: {
        tomato_id: '',
        role: '',
        uid: ''
      },
      result: null,
      error: null,
      loading: false
    };
  },
  computed: {
    evaluationTitle() {
      return this.formData.role === 'sampling' ? '抽样评估' : '收获评估';
    },
    parsedEvaluation() {
      if (!this.result?.data) return null;
      try {
        return _parseReportText(this.result.data);
      } catch (e) {
        console.warn('解析失败', e);
        return { sections: [{ title: this.evaluationTitle, items: { '原始数据': this.result.data } }] };
      }
    }
  },
  methods: {
    formatValue(val) {
      if (typeof val !== 'string') return val;
      return val.replace(/^["']|["']$/g, '').trim();
    },
    async performEvaluation() {
      if (!this.formData.tomato_id || !this.formData.role) {
        this.error = '请填写必填字段';
        return;
      }

      this.loading = true;
      this.error = null;
      this.result = null;

      try {
        const response = await fetch('http://localhost:8081/tomato/expert-evaluate', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(this.formData)
        });

        if (!response.ok) throw new Error(`HTTP ${response.status}`);
        const json = await response.json();

        if (json.code !== '0') throw new Error(json.message || '评估失败');

        let raw = json.data;

        // 尝试最多 2 层 JSON.parse
        for (let i = 0; i < 2; i++) {
          try {
            const parsed = JSON.parse(raw);
            if (typeof parsed === 'string') {
              raw = parsed;
              break;
            } else if (parsed.output && typeof parsed.output === 'string') {
              raw = parsed.output;
              break;
            } else if (typeof parsed === 'object') {
              raw = JSON.stringify(parsed);
            }
          } catch (e) {
            break;
          }
        }

        // 清理常见转义：\" → ", \\n → \n
        raw = raw
            .replace(/\\"/g, '"')
            .replace(/\\\\/g, '\\')
            .replace(/\\n/g, '\n')
            .replace(/\\r/g, '\r')
            .replace(/\\t/g, '\t');

        // 确保是字符串
        if (typeof raw !== 'string') {
          raw = String(raw);
        }

        this.result = { data: raw };
      } catch (err) {
        this.error = `评估失败: ${err.message}`;
      } finally {
        this.loading = false;
      }
    }

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
    margin-bottom: 1.6rem;
  }

  label {
    display: block;
    margin-bottom: 0.6rem;
    font-weight: 600;
    color: #2c3e50;
    font-size: 0.95rem;
  }

  input,
  select {
    width: 100%;
    padding: 0.85rem 1rem;
    border: 1px solid #d0d9eb;
    border-radius: 8px;
    font-size: 1rem;
    color: #2c3e50;
    background: #f8fbff;
    transition: border-color 0.3s ease;
  }

  input:focus,
  select:focus {
    outline: none;
    border-color: #4a6cf7;
    background: #ffffff;
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
    width: 100%;
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

  .result-section {
    margin-bottom: 2rem;
    &:last-child { margin-bottom: 0; }

    h4 {
      color: #4a6cf7;
      margin-bottom: 1.2rem;
      padding-bottom: 0.6rem;
      border-bottom: 2px solid #4a6cf7;
      font-weight: 600;
      font-size: 1.15rem;
    }
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
    }

    .pre-wrap {
      white-space: pre-wrap;
      font-family: 'Segoe UI', system-ui, sans-serif;
    }
  }

  .harvest-evaluation {
    padding-top: 0.4rem;

    .evaluation-section {
      margin-bottom: 1.2rem;

      h5 {
        color: #2c3e50;
        font-weight: 600;
        margin-bottom: 0.6rem;
        font-size: 1rem;
        border-left: 3px solid #5a67d8;
        padding-left: 0.8rem;
      }

      ul {
        list-style: none;
        padding-left: 0;
        margin: 0;

        li {
          margin-bottom: 0.5rem;
          font-size: 0.95rem;
          line-height: 1.5;

          strong {
            color: #5d6b82;
            font-weight: 600;
          }

          span {
            color: #2c3e50;
          }
        }
      }

      &.overall-score {
        background: #f0f9ff;
        border-radius: 8px;
        padding: 1rem;
        border-left: 4px solid #10b981;

        h5 {
          color: #10b981;
          font-size: 1.1rem;
        }

        li {
          display: flex;
          justify-content: space-between;
          font-weight: 600;
          color: #155724;

          strong {
            color: inherit;
          }
        }
      }
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
}
</style>
