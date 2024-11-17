<script setup lang="ts">
import { defineProps } from 'vue'

interface Diary {
  id: number
  image: string
  title: string
  date: string
  location: string
  content: string
}

const props = defineProps<{
  diary: Diary
}>()

const openDiaryModal = (diary: Diary) => {
  // ここにモーダルを開く処理を実装
  console.log('Opening modal for Diary:', diary)
}
</script>

<template>
  <v-card @click="openDiaryModal(diary)">
    <v-img :src="diary.image" height="200"></v-img>
    <v-card-title>{{ diary.title }}</v-card-title>
    <v-card-subtitle>{{ diary.date }} - {{ diary.location }}</v-card-subtitle>
    <v-card-text>{{ diary.content.substring(0, 100) }}...</v-card-text>
  </v-card>
</template>

<style scoped></style>
