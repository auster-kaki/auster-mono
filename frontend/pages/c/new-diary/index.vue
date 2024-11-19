<template>
  <v-stepper v-model="currentStep" class="mb-16">
    <v-stepper-header>
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
          @submit="currentStep += 1"
        />
      </v-stepper-content>

      <v-stepper-content step="2">
        <destination-selection
          @destination-selected="handleDestinationSelected"
        />
        <v-btn text @click="currentStep -= 1">
          Back
        </v-btn>
      </v-stepper-content>

      <v-stepper-content step="3">
        <experience-selection
          :video="experienceForm.video"
          :video-title="experienceForm.videoTitle"
          :video-description="experienceForm.videoDescription"
          :experiences="experienceForm.experiences"
          @select-experience="handleSelectExperience"
        />
        <v-btn text @click="currentStep -= 1">
          Back
        </v-btn>
      </v-stepper-content>
      <v-stepper-content step="4">
        <v-btn text @click="currentStep += 1">体験詳細</v-btn>
        <v-btn text @click="currentStep -= 1">
          Back
        </v-btn>
      </v-stepper-content>
      <v-stepper-content step="5">
        <NewDiaryItinerary
          :itinerary="itinerary"
        />
        <v-container>
          <v-row class="mt-4 pb-4">
            <v-col cols="6">
              <v-btn block @click="currentStep += 1">戻る</v-btn>
            </v-col>
            <v-col cols="6">
              <v-btn block color="primary" @click="onGoToConfirm">確認画面へ</v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-stepper-content>
      <v-stepper-content step="6">
        <NewDiaryConfirm
          :booking-info="bookingInfo"
          @confirm="handleConfirm"
        />
        <v-btn text @click="currentStep -= 1">
          Back
        </v-btn>
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

export default {
  components: { NewDiaryConfirm, NewDiaryItinerary, ExperienceSelection, DestinationSelection, DepartureSelection },
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
    submitForm() {
      // Handle form submission
      console.log('Form submitted')
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
      this.$router.push('/c/home')
      alert('予約が確認されました')
    }
  }
}
</script>
