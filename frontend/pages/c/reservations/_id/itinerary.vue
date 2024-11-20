<script>

import DiaryCarousel from '~/components/c/diary/DiaryCarousel.vue'
import NewDiaryItinerary from '~/components/c/new-diary/itinerary.vue'

export default {
  name: 'Itinerary',
  components: { NewDiaryItinerary, DiaryCarousel },
  layout: 'mobile',
  data() {
    return {
      diary: {
        id: '',
        date: '',
        image: '',
        title: '',
        content: ''
      },
      bring: [],
      itinerary: []
    }
  },
  mounted() {
    this.fetchItinerary()
  },
  methods: {
    fetchItinerary() {
      // API通信の代わりにダミーデータを返す
      this.bring = [
        '軍手', '手袋', 'ジャケット', 'ハンカチ'
      ]
      this.itinerary = [
        { type: '移動', duration: 30, description: '目安: 30分' },
        { type: '観光', duration: 30, description: 'しあわせ三蔵記念撮影(30分)' },
        { type: '移動', duration: 45, description: '目安：45分' },
        { type: 'アクティビティ', duration: 180, description: '14:00~ 一本釣り体験( 3時間 )' },
        { type: '移動', duration: 120, description: '目安: 2時間' }
      ]
      this.diary = {
        id: 1,
        date: '2024/11/23',
        image: 'http://localhost:3000/auster-mono/_nuxt/static/destination/choshi.jpg',
        title: 'ダミーデータ！！！',
        content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
      }

      // 実際のAPI通信処理（コメントアウト）
      /*
      try {
        const response = await fetch('/api/reservations')
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
  <v-container class="mb-8">
    <h2 class="text-center mb-4">日記</h2>
    <DiaryCarousel class="mb-8" :diary="diary" :is-history="true" />
    <h2 class="text-center">旅程</h2>
    <NewDiaryItinerary
      :bring="bring"
      :itinerary="itinerary"
    />
  </v-container>
</template>

<style scoped>

</style>
