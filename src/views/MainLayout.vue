<script setup>
import { ref, computed, onMounted } from 'vue'
import TomatoQuery from '@/components/TomatoQuery.vue'
import RoleRegistration from '@/components/RoleRegistration.vue'
import TomatoManagement from '@/components/TomatoManagement.vue'
import ExpertEvaluation from '@/components/ExpertEvaluation.vue'
import TomatoIdList from '@/components/TomatoIdList.vue'

const currentTab = ref('tomato-query')

const tabs = [
  { id: 'tomato-ids', name: '番茄ID列表' },
  { id: 'tomato-query', name: '番茄查询' },
  { id: 'role-registration', name: '角色注册' },
  { id: 'tomato-management', name: '番茄管理' },
  { id: 'expert-evaluation', name: '专家评估' }
]

const currentTabComponent = computed(() => {
  switch (currentTab.value) {
    case 'tomato-query': return TomatoQuery
    case 'role-registration': return RoleRegistration
    case 'tomato-management': return TomatoManagement
    case 'expert-evaluation': return ExpertEvaluation
    case 'tomato-ids': return TomatoIdList
    default: return TomatoQuery
  }
})

onMounted(() => {
  const tab = new URLSearchParams(window.location.search).get('tab')
  if (tab && tabs.some(t => t.id === tab)) {
    currentTab.value = tab
  }
})
</script>

<template>
  <div class="main-layout">

    <div class="bottom-section">
      <aside class="sidebar">
        <div class="sidebar-header">
          <h2>番茄溯源</h2>
          <h2>系统</h2>
        </div>
        <ul class="menu">
          <li
              v-for="tab in tabs"
              :key="tab.id"
              :class="{ active: currentTab === tab.id }"
          >
            <button @click="currentTab = tab.id">
              {{ tab.name }}
            </button>
          </li>
        </ul>
      </aside>

      <main class="content-area">
        <component :is="currentTabComponent"></component>
      </main>
    </div>
  </div>
</template>

<style scoped lang="scss">
.main-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.bottom-section {
  flex: 1;
  display: flex;
}

.sidebar {
  width: 240px;
  background: #f5f3ff;
  border-right: 1px solid #e6ebf5;
  display: flex;
  flex-direction: column;

  .sidebar-header {
    padding: 1.5rem 1rem;
    text-align: center;
    background: linear-gradient(135deg, #5a67d8, #805ad5);
    border-radius: 12px 12px 0 0;
    color: white;
    position: relative;

    h2 {
      margin: 0;
      font-weight: 700;
      line-height: 1.2;

      &:first-child { font-size: 28px; margin-bottom: 0.2em; }
      &:last-child { font-size: 24px; }
    }
  }

  .menu {
    flex-grow: 1;
    list-style: none;
    margin: 0;
    padding: 0;

    li {
      margin: 0;
      &:not(:last-child) { border-bottom: 1px solid #e6ebf5; }
    }

    button {
      width: 100%;
      height: 60px;
      display: flex;
      align-items: center;
      justify-content: center;
      border: none;
      // ✅ 菜单项背景与 sidebar 一致（视觉融合）
      background: #f5f3ff;
      color: #555;
      font-size: 14px;
      transition: all 0.3s ease;
      padding: 0.25rem 0;
      cursor: pointer;

      &:hover {
        background: #ede9fe;
        color: #2c3e50;
      }

      &.active {
        background: white;
        color: #5a67d8;
        font-weight: 600;
        box-shadow: inset 4px 0 0 0 #5a67d8;
      }
    }
  }
}

.content-area {
  flex: 1;
  padding: 2rem 2rem;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
  background: #f5f3ff;
  transition: background 0.3s ease;

  > * {
    margin: 0 auto;
    max-width: 100%;
  }
}

@media (max-width: 768px) {
  .bottom-section { flex-direction: column; }
  .sidebar { width: 100%; border-right: none; border-bottom: 1px solid #e6ebf5; }
  .content-area { padding: 1.5rem 1rem; }
}
</style>
