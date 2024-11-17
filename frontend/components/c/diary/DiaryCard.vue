<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

interface Diary {
  id: number
  image: string
  title: string
  date: string
  location: string
  content: string
}

const props = defineProps<{
  value: Diary
}>()

const emit = defineEmits(['click'])

const openDiaryModal = () => {
  emit('click', props.value)
}
</script>

<template>
  <v-card @click="openDiaryModal">
    <v-img :src="value.image" height="200"></v-img>
    <v-card-title>{{ value.title }}</v-card-title>
    <v-card-subtitle>{{ value.date }} - {{ value.location }}</v-card-subtitle>
    <v-card-text>{{ value.content.substring(0, 100) }}...</v-card-text>
  </v-card>
</template>

<style scoped></style>
