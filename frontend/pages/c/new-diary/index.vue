<template>
  <v-stepper v-model="currentStep" class="mb-16" flat max-width="1200px" style="margin: 0 auto;">
    <v-stepper-header style="box-shadow: unset !important;">
      <v-stepper-step :complete="currentStep > 1" step="1" />
      <v-divider />
      <v-stepper-step :complete="currentStep > 2" step="2" />
      <v-divider />
      <v-stepper-step :complete="currentStep > 3" step="3" />
      <v-divider />
      <v-stepper-step :complete="currentStep > 4" step="4" />
      <v-divider />
      <v-stepper-step :complete="currentStep > 5" step="5" />
      <v-divider />
      <v-stepper-step :complete="currentStep > 6" step="6" />
    </v-stepper-header>

    <v-stepper-items>
      <v-stepper-content step="1">
        <departure-selection
          :departure-place="departureForm.departurePlace"
          :departure-date="departureForm.departureDate"
          :return-date="departureForm.returnDate"
          :departure-time="departureForm.departureTime"
          :return-time="departureForm.returnTime"
          :interests="departureForm.interests"
          @submit="handleDepatureSelectionSubmit"
        />
      </v-stepper-content>

      <v-stepper-content step="2">
        <destination-selection
          @destination-selected="handleDestinationSelected"
        />
        <v-btn text @click="currentStep -= 1">
          戻る
        </v-btn>
      </v-stepper-content>

      <v-stepper-content step="3">
        <experience-selection
          :experiences="experienceForm.experiences"
          @click="handleSelectExperience"
        />
        <v-container>
          <v-row class="mt-4 pb-4">
            <v-btn text @click="currentStep -= 1">戻る</v-btn>
          </v-row>
        </v-container>
      </v-stepper-content>
      <v-stepper-content style="padding: 8px" step="4">
        <diary-carousel :diary="createdDiary" @select="handleSelectDiary" />
        <v-btn text @click="currentStep -= 1">
          戻る
        </v-btn>
      </v-stepper-content>
      <v-stepper-content step="5">
        <h2 class="text-center mb-4">旅程</h2>
        <NewDiaryItinerary
          :bring="bring"
          :itinerary="itinerary"
        />
        <v-container>
          <v-row class="mt-4 pb-4">
            <v-btn text @click="currentStep -= 1">戻る</v-btn>
            <v-spacer />
            <v-btn color="primary" @click="onGoToConfirm">申し込み確認画面へ</v-btn>
          </v-row>
        </v-container>
      </v-stepper-content>
      <v-stepper-content step="6">
        <NewDiaryConfirm
          :itinerary="itinerary"
          @confirm="handleConfirm"
        />
        <v-container>
          <v-row class="mt-4 pb-4">
            <v-btn text @click="currentStep -= 1">戻る</v-btn>
            <v-spacer />
            <v-btn color="primary" @click="handleConfirm">予約する</v-btn>
          </v-row>
        </v-container>
      </v-stepper-content>
    </v-stepper-items>
  </v-stepper>
</template>

<script>
import DepartureSelection from '~/components/c/new-diary/departure-selection.vue'
import DestinationSelection from '~/components/c/new-diary/destination-selection.vue'
import ExperienceSelection from '~/components/c/new-diary/experience-selection.vue'
import NewDiaryItinerary from '~/components/c/new-diary/itinerary.vue'
import NewDiaryConfirm from '~/components/c/new-diary/confirm.vue'
import { useUserStore } from '~/store/user'
import DiaryCarousel from '~/components/c/diary/DiaryCarousel.vue'

export default {
  components: {
    DiaryCarousel,
    NewDiaryConfirm,
    NewDiaryItinerary,
    ExperienceSelection,
    DestinationSelection,
    DepartureSelection
  },
  layout: 'mobile',
  data() {
    return {
      currentStep: 1,
      departureForm: {
        departurePlace: '',
        departureDate: null,
        returnDate: null,
        departureTime: null,
        returnTime: null,
        interests: []
      },
      destinationForm: {
        id: ''
      },
      experienceForm: {
        experiences: []
      },
      selectedTravelSpotId: '',
      createdDiary: {
        id: 1,
        date: '2024/11/23',
        image: 'http://localhost:3000/auster-mono/_nuxt/static/destination/choshi.jpg',
        title: 'ダミーデータ！！！',
        content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
      },
      bring: [],
      itinerary: [],
      userInfo: {
        id: ''
      }
    }
  },
  mounted() {
    const userStore = useUserStore()
    userStore.initializeUser()
    this.userInfo = userStore.userInfo
  },
  methods: {
    handleDepatureSelectionSubmit(model) {
      this.departureForm = model
      this.currentStep += 1
    },
    async handleDestinationSelected(id) {
      this.destinationForm.id = id
      const params = new URLSearchParams({
        user_id: this.userInfo.id,
        hobby_id: this.departureForm.interests
      })

      try {
        const response = await fetch(`${process.env.BASE_URL}/travel_spots?${params.toString()}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })

        if (!response.ok) {
          throw new Error('Network response was not ok')
        }
        const data = await response.json()

        let experiences = []
        if (data.length >= 5) {
          // ランダムに4つ選んでexperiencesに値を入れる
          const shuffled = data.sort(() => 0.5 - Math.random())
          experiences = shuffled.slice(0, 4).map((spot, _i) => ({
            id: spot.ID,
            image: spot.image || 'https://placehold.jp/300x200.png',
            title: spot.Name,
            description: spot.Description,
            hasFurusatoNozei: Math.random() < 0.5 // ランダムにtrueかfalseを設定
          }))
        } else {
          // そのまま入れる
          experiences = data.map((spot, _i) => ({
            id: spot.ID,
            image: spot.image || 'https://placehold.jp/300x200.png',
            title: spot.Name,
            description: spot.Description,
            hasFurusatoNozei: Math.random() < 0.5 // ランダムにtrueかfalseを設定
          }))
        }
        console.log(this.experienceForm.experiences)

        // TODO マスターデータ差し替える
        this.experienceForm = {
          video: 'https://example.com/video1.mp4',
          videoTitle: '美しい山々を体験しよう',
          videoDescription: '雄大な山々の素晴らしさを発見してください',
          experiences
        }
        this.currentStep += 1
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
      }
    },
    async handleSelectExperience(id) {
      try {
        const response = await fetch(`${process.env.BASE_URL}/travel_spots/${id}/diary`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            user_id: this.userInfo.id,
            hobby_id: this.departureForm.interests,
            date: this.departureForm.departureDate
          })
        })
        this.selectedTravelSpotId = id
        if (!response.ok) {
          throw new Error('Network response was not ok')
        }
        this.currentStep += 1

        // ゴニョゴニョして、responseを本文と画像に分けてそれぞれ格納する
        const contentType = response.headers.get('Content-Type')
        const boundary = contentType.match(/boundary=(.*)/)[1]
        const rawData = await response.text()
        const parts = rawData.split(`--${boundary}`)
        let responseJson = {}
        let imageUrl = ''
        for (const part of parts) {
          if (part.includes('Content-Type: application/json')) {
            // JSON部分を解析
            const jsonBody = part.split('\r\n\r\n')[1].trim()
            responseJson = JSON.parse(jsonBody)
          } else if (part.includes('Content-Type: image/jpeg')) {
            // 画像データ部分を解析
            const binaryData = part.split('\r\n\r\n')[1].trim()
            const binaryArray = new Uint8Array(binaryData.length)

            // バイナリデータを生成
            for (let i = 0; i < binaryData.length; i++) {
              binaryArray[i] = binaryData.charCodeAt(i)
            }

            // Blob URLを生成して設定
            const blob = new Blob([binaryArray], { type: 'image/jpeg' })
            imageUrl = URL.createObjectURL(blob)
          }
        }

        // 解析した内容を元に画像等を格納する
        this.createdDiary = {
          id: responseJson.ID,
          date: this.departureForm.departureDate,
          image: imageUrl && imageUrl.startsWith('http')
            ? imageUrl
            : 'http://localhost:3000/auster-mono/_nuxt/static/destination/choshi.jpg',
          title: responseJson.Title,
          content: responseJson.Body
        }

      } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
      }
    },
    async handleSelectDiary(_id) {
      const params = new URLSearchParams({
        user_id: this.userInfo.id
      })
      try {
        const response = await fetch(`${process.env.BASE_URL}/travel_spots/${this.selectedTravelSpotId}/itineraries?${params.toString()}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })
        const data = await response.json()
        this.bring = data.Items.map((item) => item.Name)
        this.itinerary = data.TravelSpotItineraries.map((itinerary) => ({
          id: itinerary.ID,
          kind: itinerary.Kind,
          takeTime: itinerary.TakeTime,
          price: itinerary.Price,
          order: itinerary.Order
        }))
        this.currentStep += 1
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
      }
    },
    onGoToConfirm() {
      this.currentStep += 1
    },
    async handleConfirm() {
      await fetch(`${process.env.BASE_URL}/reservations`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          user_id: this.userInfo.id,
          travel_spot_id: this.selectedTravelSpotId,
          travel_spot_diary_id: this.createdDiary.id.toString(),
          from_date: this.departureForm.departureDate,
          to_date: this.departureForm.returnDate,
        })
      })
      this.$router.push({ path: '/c/reservations', query: { reservation: 'success' } })
    }
  }
}
</script>
