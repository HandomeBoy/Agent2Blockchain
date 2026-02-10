<script>
export default {
  name: 'WelcomePage'
}
</script>

<template>
  <div class="welcome-page">
    <!-- 顶部横幅：动态高度，留出卡片空间 -->
    <div class="hero-section">
      <h1 class="title">欢迎使用番茄溯源系统</h1>
      <p class="subtitle">通过区块链技术确保番茄产品全程可追溯</p>
    </div>

    <!-- 底部卡片区：固定高度 280px，1行4列 -->
    <div class="features-grid">
      <div class="feature-card" @click="$router.push('/main?tab=tomato-query')">
        <div class="feature-icon">🔍</div>
        <h3>信息查询</h3>
        <p>快速查询番茄产品的详细信息和溯源记录</p>
      </div>
      <div class="feature-card" @click="$router.push('/main?tab=role-registration')">
        <div class="feature-icon">📝</div>
        <h3>角色管理</h3>
        <p>注册和管理不同角色的用户权限</p>
      </div>
      <div class="feature-card" @click="$router.push('/main?tab=tomato-management')">
        <div class="feature-icon">🌱</div>
        <h3>产品管理</h3>
        <p>添加和更新番茄产品的各个阶段信息</p>
      </div>
      <div class="feature-card" @click="$router.push('/main?tab=expert-evaluation')">
        <div class="feature-icon">🔬</div>
        <h3>专家评估</h3>
        <p>专业质量评估和数据分析服务</p>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
:root {
  --card-bg: #ffffff;
  --card-hover-shadow: 0 8px 24px rgba(90, 103, 216, 0.12);
  --border-light: rgba(0, 0, 0, 0.06);
}

.welcome-page {
  display: flex;
  flex-direction: column;
  min-height: 100vh; // 保证滚动兼容
  margin: 0;
  padding: 0;
  background: #f8fafc;
  font-family: 'Segoe UI', system-ui, sans-serif;
}

.hero-section {
  flex: 1; // 剩余空间给横幅（≈ 70%~75%）
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  padding: 0 1rem;
  background: linear-gradient(135deg, #5a67d8, #764ba2);
  color: white;
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: -50%;
    left: -50%;
    width: 300%;
    height: 300%;
    background: radial-gradient(circle, rgba(255,255,255,0.1) 0%, transparent 70%);
    z-index: -1;
  }

  .title {
    font-size: 2.6rem;
    font-weight: 700;
    margin-bottom: 0.8rem;
    text-shadow: 0 2px 8px rgba(0,0,0,0.1);
    animation: fadeInUp 0.8s ease-out;
  }

  .subtitle {
    font-size: 1.2rem;
    opacity: 0.92;
    max-width: 600px;
    line-height: 1.6;
    animation: fadeInUp 0.8s ease-out 0.2s both;
  }

  @keyframes fadeInUp {
    from { opacity: 0; transform: translateY(20px); }
    to { opacity: 1; transform: translateY(0); }
  }
}

.features-grid {
  height: 280px; // 固定高度，确保卡片不被压缩
  display: flex;
  align-items: center; // 垂直居中
  justify-content: center;
  padding: 0 1rem;
  gap: 1.2rem; // 略增水平间隙，提升呼吸感

  @media (max-width: 768px) {
    flex-wrap: wrap;
    gap: 0.9rem;
    padding: 0 0.8rem;
  }

  @media (max-width: 576px) {
    gap: 0.7rem;
    padding: 0 0.5rem;
  }
}

.feature-card {
  background: var(--card-bg);
  padding: 1.5rem 1.2rem;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.04);
  text-align: center;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.17, 0.67, 0.88, 1);
  width: calc(25% - 1.2rem); // 动态计算宽度（4列 - 3个 gap）
  max-width: 220px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border: 1px solid var(--border-light);

  &:hover {
    transform: translateY(-5px) scale(1.03);
    box-shadow: var(--card-hover-shadow);
    border-color: rgba(90, 103, 216, 0.3);
  }

  .feature-icon {
    font-size: 2.5rem;
    margin-bottom: 0.8rem;
    color: #5a67d8;
    transition: transform 0.3s;
  }

  &:hover .feature-icon {
    transform: scale(1.15);
  }

  h3 {
    color: #2d3748;
    margin-bottom: 0.7rem;
    font-size: 1.3rem;
    font-weight: 600;
    line-height: 1.4;
  }

  p {
    color: #4a5568;
    font-size: 0.94rem;
    line-height: 1.55;
    margin: 0;
    text-align: center;
    min-height: 2.4em;
  }
}
</style>
