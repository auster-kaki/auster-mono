<script>
export default {
  name: 'NewDiaryItinerary',
  props: {
    bring: {
      type: Array,
      required: true
    },
    itinerary: {
      type: Array,
      required: true
    }
  },
  emits: ['goBack', 'goToConfirm'],
  methods: {
    onGoBack() {
      this.$emit('goBack')
    },
    onGoToConfirm() {
      this.$emit('goToConfirm')
    }
  }
}
</script>

<template>
  <v-container>
    <v-timeline align-top dense>
      <v-timeline-item
        color="secondary"
        small
      >
        <v-card>
          <v-card-title class="text-h6">
            持ち物リスト
          </v-card-title>
          <v-card-text>
            <ul>
              <li v-for="(item, index) in bring" :key="index">
                {{ item }}
              </li>
            </ul>
          </v-card-text>
        </v-card>
      </v-timeline-item>
      <v-timeline-item
        v-for="(item, index) in itinerary"
        :key="index"
        :color="item.type === '移動' ? 'grey' : 'primary'"
        small
      >
        <v-card>
          <v-card-title class="text-h6">
            {{ item.type }}
          </v-card-title>
          <v-card-text>
            {{ item.description }}
          </v-card-text>
        </v-card>
      </v-timeline-item>
    </v-timeline>
  </v-container>
</template>

<style scoped>
</style>
