<template>
  <div class="tomato-trace-system">
    <div class="hero-section">
      <h1>欢迎使用番茄溯源系统</h1>
      <p>通过区块链技术确保番茄产品全程可追溯</p>
    </div>

    <div class="features-grid">
      <div class="feature-card">
        <div class="feature-icon">🔍</div>
        <h3>信息查询</h3>
        <p>快速查询番茄产品的详细信息和溯源记录</p>
      </div>
      <div class="feature-card">
        <div class="feature-icon">📝</div>
        <h3>角色管理</h3>
        <p>注册和管理不同角色的用户权限</p>
      </div>
      <div class="feature-card">
        <div class="feature-icon">🌱</div>
        <h3>产品管理</h3>
        <p>添加和更新番茄产品的各个阶段信息</p>
      </div>
      <div class="feature-card">
        <div class="feature-icon">🔬</div>
        <h3>专家评估</h3>
        <p>专业质量评估和数据分析服务</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TomatoTraceSystem'
}
</script>

<style scoped lang="scss">
.tomato-trace-system {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;

  .hero-section {
    text-align: center;
    margin-bottom: 2.5rem;
    padding: 2.5rem 2rem;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border-radius: 15px;
    box-shadow: 0 8px 24px rgba(0,0,0,0.15);

    h1 {
      font-size: 2.2rem;
      margin-bottom: 0.8rem;
    }
    p {
      font-size: 1.1rem;
      opacity: 0.95;
    }
  }

  .features-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
    gap: 1.8rem;
  }

  .feature-card {
    background: white;
    padding: 1.8rem 1.5rem;
    border-radius: 10px;
    box-shadow: 0 4px 12px rgba(0,0,0,0.06);
    text-align: center;
    transition: transform 0.3s ease, box-shadow 0.3s ease;

    &:hover {
      transform: translateY(-4px);
      box-shadow: 0 6px 18px rgba(0,0,0,0.1);
    }

    .feature-icon {
      font-size: 2.6rem;
      margin-bottom: 1rem;
    }
    h3 {
      color: #667eea;
      margin-bottom: 0.8rem;
      font-size: 1.3rem;
    }
    p {
      color: #666;
      font-size: 0.95rem;
      line-height: 1.5;
    }
  }
}
</style>
