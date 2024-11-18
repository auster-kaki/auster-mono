<template>
  <v-container class="pa-0">
    <v-carousel
      max-width=600
      height=640
      hide-delimiter-background
    >
      <v-carousel-item
        v-for="(entry, index) in entries"
        :key="index"
      >
        <v-card class="ma-2" elevation="4" height="620" style="position: relative;">
        <v-sheet class="pa-0" style="position: relative;">
          <v-img 
            :src="entry.image" 
            :alt="entry.title" 
            cover 
            width="100%" 
            height="270"
          />
          <v-sheet class="pa-4" style="position: absolute; top: 0; left: 0;">
            <v-icon size="24">mdi-calendar</v-icon>
            <span class="text-white">{{ entry.date }}</span>
          </v-sheet>
          <v-sheet class="pa-4" style="position: absolute; top: 0; right: 0;">
            <v-icon size="24">mdi-bookmark</v-icon>
          </v-sheet>
        </v-sheet>

        <v-sheet class="pl-6 pr-6 mt-4">
          <h3 class="text-center font-weight-bold">{{ entry.title }}</h3>
          <p class="mt-2 text-body-1">{{ entry.content }}</p>
        </v-sheet>

        <v-sheet class="d-flex justify-center" style="position: relative;">
          <v-btn
            color="accent"
            class="px-4 py-2 mb-2"
            @click="goToEntry(index)"
          >
            <span class="text-white font-weight-bold">体験詳細</span>
          </v-btn>
          <v-sheet style="position: absolute; top: 0; right: 0;">
            <v-btn icon>
              <v-icon size="x-large">mdi-share-variant</v-icon>
            </v-btn>
            <v-btn icon class="mr-4">
              <v-icon size="x-large">mdi-download</v-icon>
            </v-btn>
          </v-sheet>
        </v-sheet>
        </v-card>
      </v-carousel-item>
    </v-carousel>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'

const currentIndex = ref(0)

const entries = [
  {
    date: '2024/11/23',
    image: '/placeholder.svg?height=400&width=600',
    title: '大物ヒラマサとの出会い',
    content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
  },
  {
    date: '2024/11/24',
    image: '/placeholder.svg?height=400&width=600',
    title: '穏やかな朝の散歩',
    content: '今朝は日の出とともに目覚め、近所の公園を散歩しました。紅葉が見頃で、朝日に照らされた葉が美しく輝いていました。静かな朝の時間を過ごすことで、一日を穏やかな気持ちで始められそうです。'
  },
  {
    date: '2024/11/25',
    image: '/placeholder.svg?height=400&width=600',
    title: '新しいカフェでの発見',
    content: '街中にオープンしたという評判のカフェに行ってきました。インテリアがとてもおしゃれで、窓から差し込む光が心地よい空間でした。注文したラテアートが素晴らしく、バリスタの技術に感動。また訪れたい場所が増えました。'
  }
]

const goToEntry = (index) => {
  currentIndex.value = index
}
</script>

<style scoped>
:deep(.v-carousel__controls__item.v-btn) {
  /* アクティブなドットの色を変更する場合 */
  color: rgb(0, 0, 0,0.54);
}
</style>