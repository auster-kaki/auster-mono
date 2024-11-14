<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <h2 class="text-center mb-4">予約一覧</h2>
        <v-card v-if="reservations.length === 0" flat>
          <v-card-text class="text-center">予約がありません</v-card-text>
        </v-card>
        <v-card
          v-for="reservation in reservations"
          v-else
          :key="reservation.id"
          class="mb-4"
          @click="goToItinerary(reservation.id)"
        >
          <v-img
            :src="reservation.image"
            height="200"
            cover
          ></v-img>
          <v-card-title>{{ reservation.title }}</v-card-title>
          <v-card-subtitle>
            {{ formatDate(reservation.departureDate) }} - {{ reservation.departureCity }}
          </v-card-subtitle>
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
      reservations: []
    }
  },
  async created() {
    await this.fetchReservations()
  },
  methods: {
    formatDate(date) {
      return new Date(date).toLocaleDateString('ja-JP')
    },
    goToItinerary(id) {
      this.$router.push(`/c/reservations/${id}/itinerary`)
    },
    async fetchReservations() {
      try {
        // API呼び出しのコメントアウト
        // const response = await this.$axios.get('/api/reservations')
        // this.reservations = response.data

        // ダミーデータを返す
        this.reservations = [
          {
            id: 1,
            title: '東京旅行',
            departureDate: '2023-07-01',
            departureCity: '大阪',
            image: 'https://example.com/tokyo.jpg'
          },
          {
            id: 2,
            title: '京都観光',
            departureDate: '2023-08-15',
            departureCity: '名古屋',
            image: 'https://example.com/kyoto.jpg'
          },
          {
            id: 3,
            title: '北海道ツアー',
            departureDate: '2023-09-20',
            departureCity: '東京',
            image: 'https://example.com/hokkaido.jpg'
          }
        ]
      } catch (error) {
        console.error('予約の取得に失敗しました', error)
      }
    }
  }
}
</script>
