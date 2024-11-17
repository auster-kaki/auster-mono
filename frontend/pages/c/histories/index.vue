<template>
  <v-container>
    <v-tabs v-model="activeTab">
      <v-tab>過去の旅</v-tab>
      <v-tab>出会った人々</v-tab>
    </v-tabs>

    <v-tabs-items v-model="activeTab">
      <v-tab-item>
        <v-container>
          <v-row>
            <v-col
              v-for="(diary, index) in diaries"
              :key="index"
              cols="12"
              sm="6"
              md="4"
            >
              <diary-card
                :diary="diary"
                @click:diary="openDiaryModal(diary)"
              ></diary-card>
            </v-col>
          </v-row>
        </v-container>
      </v-tab-item>

      <v-tab-item>
        <v-list>
          <v-list-item
            v-for="(person, index) in people"
            :key="index"
            @click="openPersonModal(person)"
          >
            <v-list-item-avatar>
              <v-img :src="person.avatar"></v-img>
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title>{{ person.name }}</v-list-item-title>
              <v-list-item-subtitle
                >{{ person.location }}
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-tab-item>
    </v-tabs-items>
    <v-dialog v-model="diaryModalOpen" max-width="600">
      <diary-component :diary="selectedDiary" />
    </v-dialog>
    <v-dialog v-model="personModalOpen" max-width="600">
      <v-card v-if="selectedPerson">
        <v-card-title>{{ selectedPerson.name }}</v-card-title>
        <v-card-subtitle>{{ selectedPerson.location }}</v-card-subtitle>
        <v-card-text>
          <p>{{ selectedPerson.experience }}</p>
          <v-btn @click="openChat(selectedPerson)">チャットを開く</v-btn>
        </v-card-text>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import DiaryCard from '@/components/c/diary/DiaryCard.vue'
import DiaryComponent from '@/components/c/diary/DiaryComponent.vue'

export default {
  components: {
    DiaryCard,
    DiaryComponent,
  },
  layout: 'mobile',
  data() {
    return {
      activeTab: null,
      diaries: [],
      people: [],
      diaryModalOpen: false,
      personModalOpen: false,
      selectedDiary: null,
      selectedPerson: null,
    }
  },
  mounted() {
    this.fetchDiaries()
    this.fetchPeople()
  },
  methods: {
    fetchDiaries() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/trips')
      // this.trips = response.data

      // ダミーデータ
      this.diaries = [
        {
          id: 1,
          title: '美しい富士山',
          date: '2023-05-01',
          location: '山梨県',
          content:
            '富士山の麓でキャンプを楽しみました。澄んだ空気と雄大な景色に感動しました。',
          image: 'https://example.com/fuji.jpg',
        },
        {
          id: 2,
          title: '古都京都の旅',
          date: '2023-06-15',
          location: '京都府',
          content:
            '金閣寺や清水寺など、歴史ある寺社を巡りました。日本の伝統文化に触れる素晴らしい経験でした。',
          image: 'https://example.com/kyoto.jpg',
        },
      ]
    },
    fetchPeople() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/people')
      // this.people = response.data

      // ダミーデータ
      this.people = [
        {
          id: 1,
          name: '山田太郎',
          location: '北海道',
          avatar: 'https://example.com/yamada.jpg',
          experience:
            '北海道の大自然の中でのサバイバル体験を教えてくれました。',
        },
        {
          id: 2,
          name: '佐藤花子',
          location: '沖縄県',
          avatar: 'https://example.com/sato.jpg',
          experience: '沖縄の伝統的な織物技術を教えてくれました。',
        },
      ]
    },
    openDiaryModal(diary) {
      this.selectedDiary = diary
      this.diaryModalOpen = true
    },
    openPersonModal(person) {
      this.selectedPerson = person
      this.personModalOpen = true
    },
    openChat(person) {
      // チャットを開く処理
    },
  },
}
</script>
