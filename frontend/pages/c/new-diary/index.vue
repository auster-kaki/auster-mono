<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <h1 class="text-center mb-6">日記作成画面</h1>
        <v-form ref="form" :model-value="valid" @submit.prevent="onSubmit">
          <v-text-field
            v-model="departure"
            label="出発地"
            :rules="[v => !!v || '出発地は必須です']"
            required
          ></v-text-field>
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
      ]
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
  methods: {
    validateForm() {
      this.valid = this.departure !== '' && this.interests.length > 0
    },
    onSubmit() {
      if (this.valid) {
        // 行き先選択画面へ遷移
        this.$router.push('/c/new-diary/destination-selection')
      }
    }
  }
}
</script>
