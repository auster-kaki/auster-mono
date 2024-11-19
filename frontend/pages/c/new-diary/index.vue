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
          :video="experienceForm.video"
          :video-title="experienceForm.videoTitle"
          :video-description="experienceForm.videoDescription"
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
        <NewDiaryItinerary
          :bring="bring"
          :itinerary="itinerary"
        />
        <v-container>
          <v-row class="mt-4 pb-4">
            <v-btn text @click="currentStep -= 1">戻る</v-btn>
            <v-spacer />
            <v-btn color="primary" @click="onGoToConfirm">確認画面へ</v-btn>
          </v-row>
        </v-container>
      </v-stepper-content>
      <v-stepper-content step="6">
        <NewDiaryConfirm
          :booking-info="bookingInfo"
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
        video: '',
        videoTitle: '',
        videoDescription: '',
        experiences: [],
        videoData: {
          video: '',
          videoTitle: '',
          videoDescription: '',
          experiences: []
        }
      },
      createdDiary: {
        id: 1,
        date: '2024/11/23',
        image: 'choshi.jpg',
        title: '大物ヒラマサとの出会い',
        content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
      },
      bring: [],
      itinerary: [],
      bookingInfo: {
        expressTickets: [],
        experience: {}
      },
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
    handleDestinationSelected(id) {
      this.destinationForm.id = id
      this.experienceForm = {
        video: 'https://example.com/video1.mp4',
        videoTitle: '美しい山々を体験しよう',
        videoDescription: '雄大な山々の素晴らしさを発見してください',
        experiences: [
          {
            id: 1,
            image: 'https://placehold.jp/300x200.png',
            title: '山登り',
            description: '山道からの息をのむような景色をお楽しみください',
            hasFurusatoNozei: true
          },
          {
            id: 2,
            image: 'https://example.com/mountain2.jpg',
            title: '高原ハイキング',
            description: '爽やかな高原の空気を感じながら散策しましょう',
            hasFurusatoNozei: false
          },
          {
            id: 3,
            image: 'https://example.com/beach1.jpg',
            title: 'ビーチヨガ',
            description: '波の音を聞きながら心身をリフレッシュ',
            hasFurusatoNozei: true
          },
          {
            id: 4,
            image: 'https://example.com/diving1.jpg',
            title: 'スキューバダイビング',
            description: '色とりどりの魚たちと泳ぐ underwater adventure',
            hasFurusatoNozei: false
          }
        ]
      }
      this.currentStep += 1
    },
    handleSelectExperience(id) {
      // TODO 地域情報からこれらの情報取ってくる
      this.$emit('diary-selected', {
        id: 1,
        date: '2024/11/23',
        image: 'choshi.jpg',
        title: '大物ヒラマサとの出会い',
        content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
      })
      this.currentStep += 1
    },
    handleSelectDiary(id) {
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
      this.currentStep += 1
    },
    onGoToConfirm() {
      this.bookingInfo = {
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
      this.currentStep += 1
    },
    handleConfirm() {
      // 予約処理を実装
      console.log('予約が確認されました')
      this.$router.push({ path: '/c/home', query: { reservation: 'success' } })
    }
  }
}
</script>
