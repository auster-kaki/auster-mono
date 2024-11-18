<script>
export default {
  name: 'NewDiaryItinerary',
  props: {
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
