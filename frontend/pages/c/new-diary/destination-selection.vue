<template>
  <v-layout mobile>
    <v-main>
      <v-container>
        <h1 class="text-h4 mb-4">行き先を選択してください</h1>
        <v-row>
          <v-col v-for="destination in destinations" :key="destination.name" cols="12" sm="6">
            <v-card @click="goToExperienceSelection(destination.id)">
              <v-img :src="destination.image" height="200"></v-img>
              <v-card-title>{{ destination.name }}</v-card-title>
              <v-card-text>
                <p>{{ destination.description }}</p>
                <v-btn :to="destination.experiencesLink" text color="primary">体験一覧を見る</v-btn>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
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
        {
          id: 1,
          name: '銚子',
          image: 'https://example.com/choshi.jpg',
          description: '銚子は千葉県の東端に位置し、新鮮な海産物と美しい海岸線で知られています。',
          experiencesLink: '/experiences/choshi'
        },
        {
          id: 2,
          name: '鎌倉',
          image: 'https://example.com/kamakura.jpg',
          description: '鎌倉は歴史的な寺院や神社、そして美しい海岸を持つ古都です。',
          experiencesLink: '/experiences/kamakura'
        },
        {
          id: 3,
          name: '熱海',
          image: 'https://example.com/atami.jpg',
          description: '熱海は温泉と海の両方を楽しめる人気のリゾート地です。',
          experiencesLink: '/experiences/atami'
        },
        {
          id: 4,
          name: '伊豆',
          image: 'https://example.com/izu.jpg',
          description: '伊豆半島は美しい海岸線、温泉、そして豊かな自然で知られています。',
          experiencesLink: '/experiences/izu'
        }
      ],
    }
  },
  methods: {
    goToExperienceSelection(id) {
      this.$router.push({
        path: '/c/new-diary/experience-selection',
        query: { id }
      })
    }
  }
}
</script>

<style scoped>
.selected {
  border: 2px solid #1976D2;
}
</style>
