<template>
  <v-container class="pa-0">
    <v-carousel
      max-width=600
      height=640
      hide-delimiter-background
    >
      <v-carousel-item
        v-for="entry in diaries"
        :key="entry.id"
      >
        <v-card class="ma-2" elevation="4" height="620" style="position: relative;">
          <v-sheet class="pa-0" style="position: relative;">
            <v-img
              :src="entry.image"
              :alt="entry.title"
              cover
              width="100%"
              height="270"
            />
            <v-sheet class="pa-4" style="position: absolute; top: 0; left: 0;">
              <v-icon size="24">mdi-calendar</v-icon>
              <span class="text-white">{{ entry.date }}</span>
            </v-sheet>
            <v-sheet class="pa-4" style="position: absolute; top: 0; right: 0;">
              <v-icon size="24">mdi-bookmark</v-icon>
            </v-sheet>
          </v-sheet>

          <v-sheet class="pl-6 pr-6 mt-4">
            <h3 class="text-center font-weight-bold">{{ entry.title }}</h3>
            <p class="mt-2 text-body-1">{{ entry.content }}</p>
          </v-sheet>

          <v-sheet class="d-flex justify-center" style="position: relative;">
            <v-btn
              color="primary"
              class="px-4 py-2 mb-2"
              @click="emit('select', entry.id)"
            >
              <span class="text-white font-weight-bold">体験詳細</span>
            </v-btn>
            <v-sheet style="position: absolute; top: 0; right: 0;">
              <v-btn icon>
                <v-icon size="x-large">mdi-share-variant</v-icon>
              </v-btn>
              <v-btn icon class="mr-4">
                <v-icon size="x-large">mdi-download</v-icon>
              </v-btn>
            </v-sheet>
          </v-sheet>
        </v-card>
      </v-carousel-item>
    </v-carousel>
  </v-container>
</template>

<script setup>
defineProps({
  diaries: {
    type: Array,
    required: true
  }
})

const emit = defineEmits(['select'])
</script>

<style scoped>
:deep(.v-carousel__controls__item.v-btn) {
  /* アクティブなドットの色を変更する場合 */
  color: rgb(0, 0, 0, 0.54);
}
</style>
