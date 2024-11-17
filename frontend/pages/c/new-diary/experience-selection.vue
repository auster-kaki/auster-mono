<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'ExperienceSelection',
  layout: 'mobile',
  data() {
    return {
      video: '',
      videoTitle: '',
      videoDescription: '',
      experiences: [] as any[]
    }
  },
  mounted() {
    const id = this.$route.query.id as string
    if (id) {
      this.fetchData(id)
    }
  },
  methods: {
    fetchData(id: string) {
      const dummyData = [
        {
          id: '1',
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
        },
        {
          id: '2',
          video: 'https://example.com/video2.mp4',
          videoTitle: '海の魅力を体験しよう',
          videoDescription: '美しい海岸線と豊かな海の生態系を探索してください',
          experiences: [
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
            },
            {
              id: 5,
              image: 'https://example.com/oldtown1.jpg',
              title: '着物レンタル＆散策',
              description: '伝統的な着物を着て、タイムスリップしたような気分を味わおう',
              hasFurusatoNozei: true
            },
            {
              id: 6,
              image: 'https://example.com/teaceremony1.jpg',
              title: '茶道体験',
              description: '日本の伝統文化である茶道を体験してみましょう',
              hasFurusatoNozei: false
            }
          ]
        }
      ]

      const selectedData = dummyData.find(data => data.id === id)
      if (selectedData) {
        this.video = selectedData.video
        this.videoTitle = selectedData.videoTitle
        this.videoDescription = selectedData.videoDescription
        this.experiences = selectedData.experiences
      }
    },
    selectExperience(id: string) {
      this.$router.push({
        path: '/c/new-diary/diary-selection',
        query: { id }
      })
    }
  }
})
</script>
<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h3 class="text-h5 mb-4">銚子について</h3>
        <v-card>
          <v-card-text>
            <video :src="video" controls class="w-100"></video>
            <v-card-title class="text-h5">{{ videoTitle }}</v-card-title>
            <v-card-subtitle>{{ videoDescription }}</v-card-subtitle>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <h3 class="text-h5 mb-4">体験</h3>
        <v-row>
          <v-col v-for="experience in experiences" :key="experience.id" cols="12" sm="6" md="6" lg="6">
            <v-card>
              <v-img :src="experience.image" :alt="experience.title" height="200" style="position:relative" cover>
                <v-chip
                  v-if="experience.hasFurusatoNozei" color="primary"
                  style="position:absolute; top: 10px; right: 10px;">ふるさと納税対象
                </v-chip>
              </v-img>
              <v-card-title>{{ experience.title }}</v-card-title>
              <v-card-text>
                <p>{{ experience.description }}</p>
              </v-card-text>
              <v-card-actions>
                <v-spacer />
                <v-btn color="primary" @click="selectExperience(experience.id)">体験する</v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
</style>
