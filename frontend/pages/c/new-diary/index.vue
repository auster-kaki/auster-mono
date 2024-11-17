<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <h1 class="text-center mb-6">日記作成画面</h1>
        <v-form ref="form" :model-value="valid" @submit.prevent="onSubmit">
          <v-select
            v-model="departure"
            :items="departureOptions"
            label="出発地"
            :rules="[v => !!v || '出発地は必須です']"
            required
          ></v-select>
          <v-select
            v-model="interests"
            :items="interestOptions"
            label="興味趣味"
            multiple
            chips
            :rules="[v => v.length > 0 || '興味趣味は必須です']"
            required
          ></v-select>
          <v-btn
            color="primary"
            block
            large
            type="submit"
            :disabled="!valid"
          >
            探す
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  layout: 'mobile',
  data() {
    return {
      form: null,
      valid: false,
      departure: '',
      interests: [],
      interestOptions: [
        '旅行', '料理', 'スポーツ', '音楽', '映画', '読書', 'アート', 'テクノロジー'
      ],
      departureOptions: []
    }
  },
  watch: {
    departure() {
      this.validateForm()
    },
    interests: {
      handler() {
        this.validateForm()
      },
      deep: true
    }
  },
  mounted() {
    this.fetchDepartureOptions()
  },
  methods: {
    validateForm() {
      this.valid = this.departure !== '' && this.interests.length > 0
    },
    onSubmit() {
      if (this.valid) {
        // 行き先選択画面へ遷移
        this.$router.push('/c/new-diary/destination-selection')
      }
    },
    async fetchDepartureOptions() {
      try {
        // API通信のダミーコード
        const response = await this.getDummyDepartureOptions()
        this.departureOptions = response.data
      } catch (error) {
        console.error('出発地の取得に失敗しました:', error)
      }
    },
    getDummyDepartureOptions() {
      // ダミーデータを返す関数
      return new Promise((resolve) => {
        setTimeout(() => {
          resolve({
            data: ['東京', '大阪', '名古屋', '福岡', '札幌', '仙台', '広島', '那覇']
          })
        }, 500)
      })
    }
  }
}
</script>
