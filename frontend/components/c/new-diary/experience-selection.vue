<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'ExperienceSelection',
  props: {
    video: {
      type: String,
      required: true
    },
    videoTitle: {
      type: String,
      required: true
    },
    videoDescription: {
      type: String,
      required: true
    },
    experiences: {
      type: Array,
      required: true
    }
  },
  methods: {
    selectExperience(id: string) {
      this.$emit('select-experience', id)
    }
  }
})
</script>

<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h3 class="text-h5 mb-4">銚子について</h3>
        <v-card>
          <v-card-text>
            <video :src="video" controls class="w-100"></video>
            <v-card-title class="text-h5">{{ videoTitle }}</v-card-title>
            <v-card-subtitle>{{ videoDescription }}</v-card-subtitle>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <h3 class="text-h5 mb-4">体験</h3>
        <v-row>
          <v-col v-for="experience in experiences" :key="experience.id" cols="12" sm="6" md="6" lg="6">
            <v-card>
              <v-img :src="experience.image" :alt="experience.title" height="200" style="position:relative" cover>
                <v-chip
                  v-if="experience.hasFurusatoNozei" color="accent" outlined
                  style="position:absolute; top: 10px; right: 10px; background: #fafafa !important;">
                  <strong>ふるさと納税対象</strong>
                </v-chip>
              </v-img>
              <v-card-title>{{ experience.title }}</v-card-title>
              <v-card-text>
                <p>{{ experience.description }}</p>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
</style>

