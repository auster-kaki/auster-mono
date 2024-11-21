<template>
    <div class="container">
      <h1 class="title">旅行提案</h1>
      <v-btn small class="add-experience-button" color="info" block to="/b/home">事業者向け画面へ</v-btn>
      <div class="table-container">
        <table class="custom-table">
          <thead>
            <tr class="custom-header">
              <th class="custom-th">旅行者名</th>
              <th class="custom-th">日記URL</th>
              <th class="custom-th">Instagram</th>
              <th class="custom-th">現在のステップ</th>
              <th class="custom-th">提案状況</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="traveler in travelers" :key="traveler.id" class="custom-tr">
              <td class="custom-td">{{ traveler.name }}</td>
              <td class="custom-td">
                <a :href="traveler.diaryUrl" target="_blank" class="diary-link">
                  {{ traveler.diaryUrl }}
                </a>
              </td>
              <td class="custom-td">
                <div class="instagram-links">
                  <a v-for="(link, index) in traveler.instagramLinks" :key="index" :href="link" target="_blank">
                    <v-img src="/instagram-icon.png" alt="Instagram" height="20" width="20" class="center-image"/>
                  </a>
                </div>
              </td>
              <td class="custom-td">
                <span :class="getStepClass(traveler.currentStep)">
                  {{ traveler.currentStep }}
                </span>
              </td>
              <td class="custom-td">
                <button
                  @click="toggleProposalStatus(traveler)"
                  :class="getProposalClass(traveler.proposalSent)"
                >
                  {{ traveler.proposalSent ? '提案済み' : '未提案' }}
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <!-- スナックバーコンポーネントの追加 -->
      <v-snackbar v-model="snackbar.show" :timeout="3000" top center  class="custom-snackbar">
        {{ snackbar.message }}
        <template v-slot:action="{ attrs }">
          <v-btn color="white" text @click="snackbar.show = false">閉じる</v-btn>
        </template>
      </v-snackbar>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const travelers = ref([
    {
      id: 1,
      name: '山田太郎',
      diaryUrl: 'https://example.com/diary/yamada',
      instagramLinks: ['https://instagram.com/yamada1'],
      currentStep: 4,
      proposalSent: true
    },
    {
      id: 2,
      name: '佐藤花子',
      diaryUrl: 'https://example.com/diary/sato',
      instagramLinks: ['https://instagram.com/sato'],
      currentStep: 2,
      proposalSent: false
    },
    {
      id: 3,
      name: '鈴木一郎',
      diaryUrl: 'https://example.com/diary/suzuki',
      instagramLinks: ['https://instagram.com/suzuki1'],
      currentStep: 3,
      proposalSent: true
    }
  ])
  
  const snackbar = ref({
    show: false,
    message: ''
  })
  
  function getStepClass(step) {
    const baseClass = 'step-badge'
    switch (step) {
      case 4:
        return `${baseClass} step-4`
      case 3:
        return `${baseClass} step-3`
      case 2:
        return `${baseClass} step-2`
      default:
        return `${baseClass} step-1`
    }
  }
  
  function getProposalClass(sent) {
    return sent ? 'proposal-sent' : 'proposal-not-sent'
  }
  
  function toggleProposalStatus(traveler) {
    traveler.proposalSent = !traveler.proposalSent
    if (traveler.proposalSent) {
      snackbar.value = {
        show: true,
        message: '提案メールを送信しました。'
      }
      // ここに実際の提案メール送信のロジックを追加できます
    }
  }
  </script>
  
  <style scoped>
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 1rem;
  }
  .title {
    font-size: 1.875rem;
    font-weight: bold;
    margin-bottom: 1.5rem;
    text-align: center;
  }
  .table-container {
    overflow-x: auto;
  }
  .custom-table {
    border-collapse: collapse;
    width: 100%;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    border: 1px solid #d1d5db;
  }
  .custom-header {
    background-color: #f3f4f6;
  }
  .custom-th, .custom-td {
    padding: 0.75rem;
    text-align: center;
    border: 1px solid #d1d5db;
  }
  .custom-tr:nth-child(even) {
    background-color: #f9fafb;
  }
  .custom-tr:nth-child(odd) {
    background-color: #ffffff;
  }
  .diary-link {
    color: #2563eb;
    text-decoration: none;
  }
  .diary-link:hover {
    text-decoration: underline;
  }
  .instagram-links {
    display: flex;
    justify-content: center;
    gap: 0.5rem;
  }
  button {
    padding: 0.5rem 1rem;
    border-radius: 9999px;
    font-weight: 600;
    color: white;
    cursor: pointer;
    border: none;
  }
  .step-badge {
    display: inline-block;
    border-radius: 9999px;
    padding: 0.25rem 0.75rem;
    font-size: 0.875rem;
    font-weight: 600;
    color: white;
  }
  .step-1 { background-color: #6b7280; }
  .step-2 { background-color: #3b82f6; }
  .step-3 { background-color: #f59e0b; }
  .step-4 { background-color: #ef4444; }
  .proposal-sent {
    background-color: #10b981;
  }
  .proposal-sent:hover {
    background-color: #059669;
  }
  .proposal-not-sent {
    background-color: #ef4444;
  }
  .proposal-not-sent:hover {
    background-color: #dc2626;
  }
  .add-experience-button {
    padding: 0.25rem 0.5rem;
    background-color: #2563eb; /* 青色 */
    color: rgb(0, 0, 0);
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin-bottom: 1rem; /* ボタンとテーブルの間にスペース */
    font-size: 0.875rem;
  }
  /* 画像を中央に配置するためのクラス */
  .center-image {
    display: block;
    margin: 0 auto;
  }
  </style>