<template>
  <div>
    <v-container class="form-container">
      <v-row>
        <v-col cols="12">
          <v-autocomplete
            v-model="departurePlace"
            :items="departurePlaceOptions"
            label="発着地"
            outlined
            dense
          ></v-autocomplete>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="departureDate"
            label="行きの日付"
            outlined
            dense
            type="date"
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="returnDate"
            label="帰りの日付"
            outlined
            dense
            type="date"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="departureTime"
            label="出発時間"
            outlined
            dense
            type="time"
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="returnTime"
            label="到着時間"
            outlined
            dense
            type="time"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12">
          <v-select
            v-model="interests"
            :items="interestOptions"
            label="興味関心"
            outlined
            dense
          ></v-select>
        </v-col>
      </v-row>
      <v-row justify="end">
        <v-btn color="primary" @click="onSubmit">検索する</v-btn>
      </v-row>
    </v-container>
    <div class="bottom-space"></div>
  </div>
</template>

<script>
export default {
  name: 'DepartureSelection',
  props: {
    initialDeparturePlace: {
      type: String,
      default: null
    },
    initialDepartureDate: {
      type: String,
      default: null
    },
    initialReturnDate: {
      type: String,
      default: null
    },
    initialDepartureTime: {
      type: String,
      default: null
    },
    initialReturnTime: {
      type: String,
      default: null
    },
    initialInterests: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      departurePlace: this.initialDeparturePlace,
      departureDate: this.initialDepartureDate,
      departureTime: this.initialDepartureTime,
      returnDate: this.initialReturnDate,
      returnTime: this.initialReturnTime,
      departurePlaceOptions: ['東京', '大阪', '名古屋', '福岡', '札幌', '仙台', '広島', '那覇'],
      interests: this.initialInterests,
      interestOptions: [
        '旅行', '料理', 'スポーツ', '音楽', '映画', '読書', 'アート', 'テクノロジー'
      ]
    }
  },
  methods: {
    onSubmit() {
      this.$emit('submit', {
        departurePlace: this.departurePlace,
        departureDate: this.departureDate,
        departureTime: this.departureTime,
        returnDate: this.returnDate,
        returnTime: this.returnTime,
        interests: this.interests
      })
    }
  }
}
</script>

<style scoped>
/* 上下左右中央寄せのテキスト */
.text-overlay {
  position: absolute; /* 親要素に対して絶対配置 */
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%); /* 正確に中央寄せ */
}

/* 画面下部のスペース */
.bottom-space {
  height: 50px; /* スペースの高さを調整 */
}
</style>
