<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Confirm',
  data() {
    return {
      bookingInfo: {
        expressTickets: [
          {
            date: '2024/11/09',
            time: '11:00~',
            trainName: 'しおさい３号',
            price: 1500
          },
          {
            date: '2024/11/10',
            time: '11:00~',
            trainName: 'しおさい14号',
            price: 1500
          }
        ],
        experience: {
          name: '一本釣り体験',
          price: 3000
        }
      }
    }
  },
  methods: {
    confirmBooking() {
      // ここに予約処理を実装
      console.log('予約が確認されました')
      this.$router.push(`/c/home`)
    }
  }
})
</script>

<template>
  <v-container>
    <v-card>
      <v-card-title>以下で申し込みますか</v-card-title>
      <v-card-text>
        <h3>特急券</h3>
        <v-list>
          <v-list-item v-for="(ticket, index) in bookingInfo.expressTickets" :key="index">
            <v-list-item-content>
              <v-list-item-title>{{ ticket.date }} {{ ticket.time }} {{ ticket.trainName }}</v-list-item-title>
              <v-list-item-subtitle>{{ ticket.price }}円</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>

        <h3>{{ bookingInfo.experience.name }}</h3>
        <v-list>
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title>{{ bookingInfo.experience.price }}円</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-card-text>
      <v-card-actions>
        <v-btn to="/c/new-diary/itinerary">戻る</v-btn>
        <v-spacer />
        <v-btn color="primary" @click="confirmBooking">予約する</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<style scoped>
/* スタイルが必要な場合はここに追加 */
</style>
