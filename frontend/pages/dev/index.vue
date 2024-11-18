<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import DiaryComponent from '@/components/c/diary/DiaryComponent.vue'
import DiaryCarousel from '@/components/c/diary/DiaryCarousel.vue'

const baseUrl = ref('https://api.example.com')
const endpoint = ref('/users')
const method = ref('GET')
const pathParams = ref('')
const queryParams = ref('')
const requestBody = ref('')
const response = ref('')

const diaryData = ref({
  id: 1,
  title: '美しい富士山',
  date: '2023-05-01',
  location: '山梨県',
  content:
    '富士山の麓でキャンプを楽しみました。澄んだ空気と雄大な景色に感動しました。',
  image: 'https://example.com/fuji.jpg'
})

const sendRequest = async () => {
  try {
    const url = `${baseUrl.value}${endpoint.value}${pathParams.value}`
    const config = {
      method: method.value,
      url,
      params: queryParams.value ? JSON.parse(queryParams.value) : {},
      data: requestBody.value ? JSON.parse(requestBody.value) : {}
    }
    const result = await axios(config)
    response.value = JSON.stringify(result.data, null, 2)
  } catch (error: any) {
    response.value = `Error: ${error.message}`
  }
}
</script>

<template>
  <div>
    <h1>Diaryテスト画面</h1>
    <diary-component :diary="diaryData" />
    <diary-carousel />
    <div>
      <h1>DiaryCarouselテスト画面</h1>
  </div>
      <h1>API テスト画面</h1>
    <div>
      <label>
        ベースURL:
        <input v-model="baseUrl" type="text" />
      </label>
    </div>
    <div>
      <label>
        エンドポイント:
        <input v-model="endpoint" type="text" />
      </label>
    </div>
    <div>
      <label>
        メソッド:
        <select v-model="method">
          <option>GET</option>
          <option>POST</option>
          <option>PUT</option>
          <option>DELETE</option>
        </select>
      </label>
    </div>
    <div>
      <label>
        パスパラメータ:
        <input v-model="pathParams" type="text" placeholder="/1" />
      </label>
    </div>
    <div>
      <label>
        クエリパラメータ (JSON):
        <input
          v-model="queryParams"
          type="text"
          placeholder='{"key": "value"}'
        />
      </label>
    </div>
    <div>
      <label>
        リクエストボディ (JSON):
        <textarea
          v-model="requestBody"
          placeholder='{"key": "value"}'
        ></textarea>
      </label>
    </div>
    <button @click="sendRequest">送信</button>
    <div v-if="response">
      <h2>レスポンス:</h2>
      <pre>{{ response }}</pre>
    </div>
  </div>
</template>

<style scoped>
div {
  margin-bottom: 10px;
}

label {
  display: block;
}

input,
select,
textarea {
  width: 100%;
  max-width: 400px;
}

textarea {
  height: 100px;
}

pre {
  background-color: #f0f0f0;
  padding: 10px;
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>
