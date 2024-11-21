<template>
  <div class="container mx-auto p-4">
    <h1 class="text-3xl font-bold mb-6 text-center">体験一覧</h1>
    <v-btn small  class="add-experience-button" color="info" nuxt to="/b/experiences">体験を追加する</v-btn>
    <v-btn small  class="add-experience-button" color="info" nuxt to="/b/suggest">提案を行う</v-btn>
    <div class="overflow-x-auto">
      <table class="custom-table">
        <thead>
          <tr class="custom-header">
            <th class="custom-th">体験名</th>
            <th class="custom-th">体験情報</th>
            <th class="custom-th">準備物</th>
            <th class="custom-th">金額</th>
            <th class="custom-th">予約状況（一週間）</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="experience in experiences" :key="experience.id" class="custom-tr">
            <td class="custom-td">{{ experience.name }}</td>
            <td class="custom-td">{{ experience.description }}</td>
            <td class="custom-td">
              <ul class="list-disc list-inside text-left">
                <li v-for="item in experience.preparations" :key="item">{{ item }}</li>
              </ul>
            </td>
            <td class="custom-td">{{ experience.price.toLocaleString() }}円</td>
            <td class="custom-td">
              <div class="schedule-row">
                <div v-for="(day, index) in experience.weekSchedule" :key="index" class="schedule-item">
                  <div class="status-date">{{ formatDate(day.date) }}</div>
                  <div :class="getStatusClass(day.status)" class="status-label">
                    {{ day.status }}
                  </div>
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const experiences = ref([
  {
    id: 1,
    name: '山岳トレッキング体験',
    description: '美しい山々を巡る素晴らしいトレッキング体験です。',
    capacity: 10,
    preparations: ['帽子', '飲み物'],
    isFurusatoNozei: true,
    price: 5000,
    ageRange: '15歳以上',
    weekSchedule: generateWeekSchedule()
  },
  {
    id: 2,
    name: '伝統工芸体験',
    description: '地元の職人から学ぶ伝統工芸体験です。',
    capacity: 5,
    preparations: ['エプロン', 'タオル'],
    isFurusatoNozei: false,
    price: 3000,
    ageRange: '全年齢',
    weekSchedule: generateWeekSchedule()
  },
  {
    id: 3,
    name: '農業収穫体験',
    description: '季節の野菜や果物の収穫を体験できます。',
    capacity: 15,
    preparations: ['帽子', '軍手', '長靴', '日焼け止め'],
    isFurusatoNozei: true,
    price: 2000,
    ageRange: '6歳以上',
    weekSchedule: generateWeekSchedule()
  }
])

function generateWeekSchedule() {
  const weekSchedule = []
  const startDate = new Date()
  for (let i = 0; i < 7; i++) {
    const date = new Date(startDate)
    date.setDate(date.getDate() + i)
    weekSchedule.push({
      date,
      status: getRandomStatus()
    })
  }
  return weekSchedule
}

function getRandomStatus() {
  const statuses = ['予約済み', '仮予約', '空き']
  return statuses[Math.floor(Math.random() * statuses.length)]
}

function getStatusClass(status) {
  switch (status) {
    case '予約済み':
      return 'status-reserved'
    case '仮予約':
      return 'status-provisional'
    case '空き':
      return 'status-available'
    default:
      return ''
}
}

function formatDate(date) {
  return new Intl.DateTimeFormat('ja-JP', { month: 'short', day: 'numeric' }).format(date)
}

function addExperience() {
  router.push('b/experiences')
}
</script>

<style scoped>
.custom-table {
  border-collapse: collapse;
  width: 100%;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border: 1px solid #d1d5db; /* 罫線の色 */
}

.custom-header {
  background-color: #f3f4f6; /* ヘッダーの背景色 */
}

.custom-th, .custom-td {
  padding: 0.75rem;
  text-align: center;
  border: 1px solid #d1d5db; /* セルの罫線色 */
}

.custom-tr:nth-child(even) {
  background-color: #f9fafb; /* 偶数行の背景色 */
}

.custom-tr:nth-child(odd) {
  background-color: #ffffff; /* 奇数行の背景色 */
}

/* 追加のカスタムスタイルが必要な場合はここに記述 */

.status-reserved {
  background-color: #f8d7da; /* 赤系背景 */
  color: #721c24; /* 赤系文字 */
  padding: 0.25rem;
  border-radius: 0.25rem;
  font-size: 0.75rem;
  margin-top: 0.25rem;
}

.status-provisional {
  background-color: #fff3cd; /* 黄色系背景 */
  color: #856404; /* 黄色系文字 */
  padding: 0.25rem;
  border-radius: 0.25rem;
  font-size: 0.75rem;
  margin-top: 0.25rem;
}

.status-available {
  background-color: #d4edda; /* 緑系背景 */
  color: #155724; /* 緑系文字 */
  padding: 0.25rem;
  border-radius: 0.25rem;
  font-size: 0.75rem;
  margin-top: 0.25rem;
}
.wide-column {
    width: 40%; /* 例えば、50% に変更 */
}

/* スケジュール行のスタイル */
.schedule-row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  gap: 8px;
}

/* スケジュール項目のスタイル */
.schedule-item {
  flex: 1;
  text-align: center;
}

/* ステータス日付のスタイル */
.status-date {
  font-size: 0.75rem;
  margin-bottom: 4px;
}

/* ステータスラベルの共通スタイル */
.status-label {
  font-size: 0.75rem;
  padding: 4px 8px;
  border-radius: 4px;
  display: inline-block;
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

.add-experience-button:hover {
  background-color: #1e40af; /* より濃い青色 */
}
</style>

