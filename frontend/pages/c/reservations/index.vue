<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <h2 class="text-center mb-4">予約確認</h2>
        <v-card v-if="reservations.length === 0" flat>
          <v-card-text class="text-center">予約がありません</v-card-text>
        </v-card>
        <v-card
          v-for="reservation in reservations"
          v-else
          :key="reservation.id"
          class="mb-4"
        >
          <v-row no-gutters>
            <v-col cols="4">
              <v-img
                :src="reservation.image"
                height="100%"
                cover
              />
            </v-col>
            <v-col cols="8">
              <v-card-title>{{ reservation.title }}</v-card-title>
              <v-card-subtitle>
                {{ formatDate(reservation.departureDate) }} - {{ reservation.departureCity }}
              </v-card-subtitle>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                  color="primary"
                  text
                  @click="goToItinerary(reservation.id)"
                >
                  旅程確認・日記更新
                </v-btn>
              </v-card-actions>
            </v-col>
          </v-row>
        </v-card>
      </v-col>
    </v-row>
    <v-snackbar
      v-model="snackbar.show"
      :timeout="snackbar.timeout"
      :color="snackbar.color"
      dense
      :close-text="snackbar.closeText"
      :close-icon="snackbar.closeIcon"
    >
      {{ snackbar.text }}
      <template #action="{ attrs }">
        <v-btn
          text
          v-bind="attrs"
          @click="snackbar.show = false"
        >
          <v-icon>{{ snackbar.closeIcon }}</v-icon>
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script>
import { useUserStore } from '~/store/user'

export default {
  name: 'IndexPage',
  layout: 'mobile',
  data() {
    return {
      userInfo: {},
      reservations: [],
      snackbar: {
        show: false,
        text: '',
        timeout: 3000,
        color: 'success',
        closeText: '閉じる',
        closeIcon: 'mdi-close'
      }
    }
  },
  async mounted() {
    if (this.$route.query.reservation === 'success') {
      this.showSnackbar('予約が完了しました！')
    }
    const userStore = useUserStore()
    userStore.initializeUser()
    this.userInfo = userStore.userInfo
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
      const params = new URLSearchParams({
        user_id: this.userInfo.id,
        filter: 'yet'
      })
      try {
        const response = await fetch(`${process.env.BASE_URL}/reservations?${params.toString()}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })
        console.log(response)
        // API呼び出しのコメントアウト
        // const response = await this.$axios.get('/api/reservations_filter=yet')
        // this.reservations = response.data

        this.reservations = response
        // ダミーデータを返す
        // this.reservations = [
        //   {
        //     id: 1,
        //     title: '東京旅行',
        //     departureDate: '2023-07-01',
        //     departureCity: '大阪',
        //     image: 'https://example.com/tokyo.jpg'
        //   },
        //   {
        //     id: 2,
        //     title: '京都観光',
        //     departureDate: '2023-08-15',
        //     departureCity: '名古屋',
        //     image: 'https://example.com/kyoto.jpg'
        //   },
        //   {
        //     id: 3,
        //     title: '北海道ツアー',
        //     departureDate: '2023-09-20',
        //     departureCity: '東京',
        //     image: 'https://example.com/hokkaido.jpg'
        //   }
        // ]
      } catch
        (error) {
        console.error('予約の取得に失敗しました', error)
      }
    }
    ,
    showSnackbar(text) {
      this.snackbar.text = text
      this.snackbar.show = true
    }
  }
}
</script>
