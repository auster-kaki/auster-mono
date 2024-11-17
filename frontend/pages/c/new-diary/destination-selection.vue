<template>
  <v-layout mobile>
    <v-main>
      <v-container>
        <h1 class="text-h4 mb-4">行き先を選択してください</h1>
        <v-row>
          <v-col v-for="destination in destinations" :key="destination.name" cols="12" sm="6">
            <v-card
              :class="{ 'selected': selectedDestination === destination }" @click="selectDestination(destination)">
              <v-card-title>{{ destination.name }}</v-card-title>
            </v-card>
          </v-col>
        </v-row>
        <v-dialog v-model="selectedDestination" max-width="500">
          <v-card v-if="selectedDestination">
            <v-card-title>{{ selectedDestination.name }}の動画</v-card-title>
            <v-card-text>
              <video controls :src="selectedDestination.video" width="100%"></video>
            </v-card-text>
            <v-card-actions>
              <v-btn color="primary" @click="selectedDestination = null">閉じる</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-container>
    </v-main>
  </v-layout>
</template>

<script>
export default {
  layout: 'mobile',
  data() {
    return {
      destinations: [
        { name: '銚子', video: 'https://example.com/choshi.mp4' },
        { name: '鎌倉', video: 'https://example.com/kamakura.mp4' },
        { name: '熱海', video: 'https://example.com/atami.mp4' },
        { name: '伊豆', video: 'https://example.com/izu.mp4' }
      ],
      selectedDestination: null
    }
  },
  methods: {
    selectDestination(destination) {
      this.selectedDestination = destination
    }
  }
}
</script>

<style scoped>
.selected {
  border: 2px solid #1976D2;
}
</style>
