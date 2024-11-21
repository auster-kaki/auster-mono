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
  methods: {
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
        v-for="item in itinerary"
        :key="item.id"
        :color="item.kind === 'move' ? 'grey' : 'primary'"
        small
      >
        <v-card>
          <v-card-title class="text-h6">
            {{ item.kind === 'move' ? '移動' : '体験' }} <p>{{ item.description }}</p>
          </v-card-title>
          <v-card-text>
            <p>所要時間: {{ item.takeTime }}</p>
            <p v-if="item.kind !== 'move'">費用: {{ item.price }}円</p>
            <p>順番: {{ item.Order }}</p>
          </v-card-text>
        </v-card>
      </v-timeline-item>
    </v-timeline>
  </v-container>
</template>

<style scoped>
</style>
