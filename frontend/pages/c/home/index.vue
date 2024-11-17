<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <h2 class="text-center mb-4">ホーム画面</h2>

        <!-- TODO リスト -->
        <v-card v-if="todos.length > 0" class="mb-4">
          <v-card-title>準備リスト</v-card-title>
          <v-list>
            <v-list-item v-for="(todo, index) in todos" :key="index">
              <v-list-item-action>
                <v-checkbox v-model="todo.completed"></v-checkbox>
              </v-list-item-action>
              <v-list-item-content>
                <v-list-item-title>{{ todo.text }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>

        <!-- 申し込み中の予約 -->
        <v-card v-if="pendingReservations.length > 0" class="mb-4">
          <v-card-title>申し込み中の予約</v-card-title>
          <v-list>
            <v-list-item v-for="reservation in pendingReservations" :key="reservation.id">
              <v-list-item-content>
                <v-list-item-title>{{ reservation.title }}</v-list-item-title>
                <v-list-item-subtitle>{{ reservation.date }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>

        <!-- 新しい絵日記の通知 -->
        <v-alert v-if="newDiary" type="info" class="mb-4">
          新しい絵日記が届いています！
          <v-btn @click="showNewDiary" text color="primary">確認する</v-btn>
        </v-alert>

        <!-- 絵日記の表示 -->
        <v-card v-if="currentDiary" class="mb-4">
          <v-card-title>{{ currentDiary.title }}</v-card-title>
          <v-card-subtitle>{{ currentDiary.date }}</v-card-subtitle>
          <v-img :src="currentDiary.image" height="200" contain></v-img>
          <v-card-text>{{ currentDiary.content }}</v-card-text>
          <v-card-actions>
            <v-btn color="primary" @click="reserveDiary">予約する</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'IndexPage',
  layout: 'mobile',
  data() {
    return {
      todos: [],
      pendingReservations: [],
      newDiary: false,
      currentDiary: null,
    }
  },
  mounted() {
    this.fetchTodos()
    this.fetchPendingReservations()
    this.checkNewDiary()
  },
  methods: {
    fetchTodos() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/todos')
      // this.todos = response.data

      // ダミーデータ
      this.todos = [
        { text: '持ち物の準備', completed: false },
        { text: '乗車券の準備', completed: true },
      ]
    },
    fetchPendingReservations() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/pending-reservations')
      // this.pendingReservations = response.data

      // ダミーデータ
      this.pendingReservations = [
        { id: 1, title: '山登り体験', date: '2023-06-15' },
        { id: 2, title: '料理教室', date: '2023-06-20' },
      ]
    },
    checkNewDiary() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/new-diary')
      // this.newDiary = response.data.hasNewDiary

      // ダミーデータ
      this.newDiary = true
    },
    showNewDiary() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/latest-diary')
      // this.currentDiary = response.data

      // ダミーデータ
      this.currentDiary = {
        date: '2023-06-10',
        title: '素晴らしい山登り体験',
        image: 'https://example.com/mountain.jpg',
        content: '今日は素晴らしい山登り体験をしました。景色が最高でした！'
      }
      this.newDiary = false
    },
    reserveDiary() {
      // 予約処理のロジックをここに実装
      alert('予約が完了しました！')
    }
  }
}
</script>
