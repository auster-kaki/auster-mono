<script>

export default {
  name: 'NewDiaryItinerary',
  layout: 'mobile',
  data() {
    return {
      itinerary: []
    }
  },
  mounted() {
    this.fetchItinerary()
  },
  methods: {
    fetchItinerary() {
      // API通信の代わりにダミーデータを返す
      this.itinerary = [
        { type: '移動', duration: 30, description: '目安: 30分' },
        { type: '観光', duration: 30, description: 'しあわせ三蔵記念撮影(30分)' },
        { type: '移動', duration: 45, description: '目安：45分' },
        { type: 'アクティビティ', duration: 180, description: '14:00~ 一本釣り体験( 3時間 )' },
        { type: '移動', duration: 120, description: '目安: 2時間' }
      ]

      // 実際のAPI通信処理（コメントアウト）
      /*
      try {
        const response = await fetch('/api/itinerary')
        const data = await response.json()
        this.itinerary = data
      } catch (error) {
        console.error('Failed to fetch itinerary:', error)
      }
      */
    }
  }
}
</script>

<template>
  <v-container>
    <v-timeline align-top dense>
      <v-timeline-item
        v-for="(item, index) in itinerary"
        :key="index"
        :color="item.type === '移動' ? 'grey' : 'primary'"
        small
      >
        <v-card>
          <v-card-title class="text-h6">
            {{ item.type }}
          </v-card-title>
          <v-card-text>
            {{ item.description }}
          </v-card-text>
        </v-card>
      </v-timeline-item>
    </v-timeline>
    <v-row class="mt-4 pb-4">
      <v-col cols="6">
        <v-btn block>戻る</v-btn>
      </v-col>
      <v-col cols="6">
        <v-btn block color="primary" to="/c/new-diary/confirm">確認画面へ</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>

</style>
