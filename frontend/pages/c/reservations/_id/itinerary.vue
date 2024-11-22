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
  async mounted() {
    await this.fetchItinerary()
    // API通信の代わりにダミーデータを返す
    this.bring = [
      'ダミー軍手', 'ダミー手袋', 'ダミージャケット', 'ダミーハンカチ'
    ]
    // this.itinerary = [
    //   { type: '移動', duration: 30, description: '目安: 30分' },
    //   { type: '観光', duration: 30, description: 'しあわせ三蔵記念撮影(30分)' },
    //   { type: '移動', duration: 45, description: '目安：45分' },
    //   { type: 'アクティビティ', duration: 180, description: '14:00~ 一本釣り体験( 3時間 )' },
    //   { type: '移動', duration: 120, description: '目安: 2時間' }
    // ]
    // this.diary = {
    //   id: 1,
    //   date: '2024/11/23',
    //   image: 'http://localhost:3000/auster-mono/destination/choshi.jpg',
    //   title: 'ダミーデータ！！！',
    //   content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
    // }
    // this.diaryContent = '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
  },
  methods: {
    async fetchItinerary() {
      try {
        const response = await fetch(`${process.env.BASE_URL}/reservations/${this.$route.params.id}`)
        const data = await response.json()
        this.diary = {
          id: data.id,
          date: data.from_date.split('T')[0],
          image: `${process.env.BASE_URL}/images/${data.diary_photo_path}`,
          title: data.diary_title,
          content: data.diary_description
        }
        this.diaryContent = data.diary_description
        this.itinerary = data.travelSpotItineraries.map(item => ({
          type: item.type,
          takeTime: item.take_time,
          price: item.price,
          title: item.title,
          description: item.description
        }))
        // TODO this.bring
        // {
        //   "diary_description": "bbbbbbb",
        //   "diary_photo_path": "/assets/images/travel_spots/6/IMG_5132.jpg",
        //   "diary_title": "タイトル1",
        //   "from_date": "2024-11-22T00:00:00Z",
        //   "id": "csvtnecpng7s73cntutg",
        //   "is_offer": false,
        //   "to_date": "2024-11-23T00:00:00Z",
        //   "travelSpotItineraries": [
        //   {
        //     "description": "",
        //     "id": "iti10",
        //     "kind": "move",
        //     "order": 5,
        //     "price": 1500,
        //     "take_time": 60,
        //     "title": "特急 しおさい14号"
        //   },
        //   {
        //     "description": "標高90mの屋上展望スペースからは360度の大パノラマが楽しめます。水平線の両端が丸みを帯びて見えることから、その名が付けられました。(千葉県銚子市天王台1421-13)",
        //     "id": "iti7",
        //     "kind": "spot",
        //     "order": 2,
        //     "price": 500,
        //     "take_time": 60,
        //     "title": "地球の丸く見える丘展望館"
        //   },
        // ],
        //   "travel_spot_description": "船釣りでクロカジキを釣る体験ができます",
        //   "travel_spot_title": "釣り体験 シイラ"
        // }
        // TODO diaryに入れる
        // TODO diaryContentにも入れる
        // TODO: レスポンスを適切に格納する処理を実装する
        console.log(data)
      } catch (error) {
        console.error('Failed to fetch itinerary:', error)
      }
    },
    async onDiaryClick() {
      // this.$router.go() TODO
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
    <h2 class="text-center">旅程</h2>
    <NewDiaryItinerary
      :bring="bring"
      :itinerary="itinerary"
    />
  </v-container>
</template>

<style scoped>

</style>
