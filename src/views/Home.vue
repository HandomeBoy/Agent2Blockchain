<template>
  <div class="home">
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
  name: 'Home'
}
</script>

<style scoped>
.home {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.hero-section {
  text-align: center;
  margin-bottom: 3rem;
  padding: 3rem 2rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 15px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.2);
}

.hero-section h1 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

.hero-section p {
  font-size: 1.2rem;
  opacity: 0.9;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
}

.feature-card {
  background: white;
  padding: 2rem;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0,0,0,0.1);
  text-align: center;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0,0,0,0.15);
}

.feature-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.feature-card h3 {
  color: #667eea;
  margin-bottom: 1rem;
}

.feature-card p {
  color: #666;
  line-height: 1.6;
}
</style>
