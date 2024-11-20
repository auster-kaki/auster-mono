<script lang="ts">
import { defineComponent, PropType } from 'vue'

interface ExpressTicket {
  date: string
  time: string
  trainName: string
  price: number
}

interface Experience {
  name: string
  price: number
}

interface BookingInfo {
  expressTickets: ExpressTicket[]
  experience: Experience
}

export default defineComponent({
  name: 'NewDiaryConfirm',
  props: {
    bookingInfo: {
      type: Object as PropType<BookingInfo>,
      required: true
    }
  },
  emits: ['confirm'],
  methods: {
    confirmBooking() {
      this.$emit('confirm')
    }
  }
})
</script>

<template>
  <v-container>
    <v-card elevation="0">
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
    </v-card>
  </v-container>
</template>

<style scoped>
/* スタイルが必要な場合はここに追加 */
</style>
