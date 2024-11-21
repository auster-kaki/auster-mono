<script>
import NewDiaryItinerary from '~/components/c/new-diary/itinerary.vue'
import ReservationDiary from '~/components/c/diary/ReservationDiary.vue'

export default {
  name: 'Itinerary',
  components: { ReservationDiary, NewDiaryItinerary },
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
      diaryContent: '',
      bring: [],
      itinerary: []
    }
  },
  mounted() {
    this.fetchItinerary()
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
      image: 'http://localhost:3000/auster-mono/destination/choshi.jpg',
      title: 'ダミーデータ！！！',
      content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
    }
    this.diaryContent = '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
  },
  methods: {
    async fetchItinerary() {
      try {
        const response = await fetch(`${process.env.BASE_URL}/reservations/${this.$route.params.id}`)
        const data = await response.json()
        // TODO diaryに入れる
        // TODO diaryContentにも入れる
        // TODO: レスポンスを適切に格納する処理を実装する
        console.log(data)
      } catch (error) {
        console.error('Failed to fetch itinerary:', error)
      }
    },
    async onDiaryClick() {
      try {
        const fileInput = document.createElement('input')
        fileInput.type = 'file'
        fileInput.accept = 'image/*'
        fileInput.onchange = async (event) => {
          const file = event.target.files[0]
          if (file) {
            const formData = new FormData()
            formData.append('photo', file)

            const response = await fetch(`${process.env.BASE_URL}/reservations/${this.$route.params.id}/diary_photo`, {
              method: 'PATCH',
              body: formData
            })

            if (response.ok) {
              // 画面のリロード
              this.$router.go()
              console.log('Image uploaded successfully')
            } else {
              console.error('Failed to upload image')
            }
          }
        }
        await fileInput.click()
      } catch (error) {
        console.error('Error uploading image:', error)
      }
    },
    async updateDiary(updatedDescription) {
      console.log('updateDiary', updatedDescription)
      try {
        const response = await fetch(`${process.env.BASE_URL}/reservations/${this.$route.params.id}/diary_description`, {
          method: 'PATCH',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            Description: updatedDescription
          })
        })

        if (response.ok) {
          console.log('Diary updated successfully')
          this.diary.content = updatedDescription
          this.diaryContent = updatedDescription
        } else {
          console.error('Failed to update diary')
        }
      } catch (error) {
        console.error('Error updating diary:', error)
      }
    }
  }
}
</script>

<template>
  <v-container class="mb-8">
    <h2 class="text-center mb-4">日記</h2>
    <ReservationDiary
      v-model="diaryContent"
      :diary="diary"
      class="mb-8"
      :is-history="true"
      @camera-click="onDiaryClick"
      @update-diary="updateDiary"
    />
    {{ diaryContent }}
    <h2 class="text-center">旅程</h2>
    <NewDiaryItinerary
      :bring="bring"
      :itinerary="itinerary"
    />
  </v-container>
</template>

<style scoped>

</style>
