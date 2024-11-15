<script setup>
import { ref, computed } from 'vue'
import { mapState, storeToRefs } from 'pinia'
import { useExperienceStore } from '~/components/b/experiences/_sub/experiences'

const experienceStore = useExperienceStore()
const { dates } = storeToRefs(experienceStore)
const { fetchExperiences } = experienceStore

const search = ref('')
const page = ref(1)
const itemsPerPage = ref(10)
const experiences = ref([])

const headers = computed(() => [
  { text: '体験', value: 'name' },
  { text: '詳細', value: 'description' },
  ...generateDateHeaders(),
])

const filteredExperiences = computed(() =>
  experiences.value.filter(exp =>
    exp.name.toLowerCase().includes(search.value.toLowerCase())
  )
)

function generateDateHeaders() {
  return dates.value.map(date => ({
    text: date,
    value: `status.${date}`,
    align: 'center',
    sortable: false,
  }))
}

async function loadExperiences() {
  // Commented out actual API call
  // await fetchExperiences()

  // Dummy data
  experiences.value = [
    {
      id: 1,
      name: '富士山登山',
      description: '日本一高い山に登る体験',
      status: {
        '2023-07-01': '予約済',
        '2023-07-02': '空きあり',
        '2023-07-03': '仮予約',
      },
    },
    // Add more dummy data as needed
  ]
}

export default {
  name: 'ExperienceIndex',
  data() {
    return {
      search: '',
      page: 1,
      itemsPerPage: 10,
      headers: [
        { text: '体験', value: 'name' },
        { text: '詳細', value: 'description' },
        ...this.generateDateHeaders(),
      ],
      experiences: [],
    }
  },
  computed: {
    ...mapState(['dates']),
    filteredExperiences() {
      return this.experiences.filter(exp =>
        exp.name.toLowerCase().includes(this.search.toLowerCase())
      )
    },
  },
  methods: {
    ...mapActions(['fetchExperiences']),
    generateDateHeaders() {
      return this.dates.map(date => ({
        text: date,
        value: `status.${date}`,
        align: 'center',
        sortable: false,
      }))
    },
    async loadExperiences() {
      // Commented out actual API call
      // await this.fetchExperiences()

      // Dummy data
      this.experiences = [
        {
          id: 1,
          name: '富士山登山',
          description: '日本一高い山に登る体験',
          status: {
            '2023-07-01': '予約済',
            '2023-07-02': '空きあり',
            '2023-07-03': '仮予約',
          },
        },
        {
          id: 2,
          name: '茶道体験',
          description: '伝統的な日本の茶道を学ぶ',
          status: {
            '2023-07-01': '空きあり',
            '2023-07-02': '予約済',
            '2023-07-03': '空きあり',
          },
        },
      ]
    },
  },
  created() {
    this.loadExperiences()
  },
}
</script>

<template>
  <v-container>
    <v-row>
      <v-col>
        <v-btn color="primary" @click="loadExperiences">登録</v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field
          v-model="search"
          label="体験名で検索"
          prepend-icon="mdi-magnify"
        ></v-text-field>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-data-table
          :headers="headers"
          :items="filteredExperiences"
          :search="search"
          :page.sync="page"
          :items-per-page="itemsPerPage"
          hide-default-footer
          class="elevation-1"
        >
          <template v-slot:item="{ item }">
            <tr>
              <td>{{ item.name }}</td>
              <td>{{ item.description }}</td>
              <td v-for="date in dates" :key="date" class="text-center">
                <v-chip
                  :color="getStatusColor(item.status[date])"
                  small
                >
                  {{ item.status[date] }}
                </v-chip>
              </td>
            </tr>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-pagination
          v-model="page"
          :length="Math.ceil(filteredExperiences.length / itemsPerPage)"
        ></v-pagination>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.v-data-table >>> .v-data-table__wrapper {
  overflow-x: auto;
}
</style>
