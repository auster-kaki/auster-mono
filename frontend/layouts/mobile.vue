<template>
  <v-app>
    <v-main>
      <v-container>
        <Nuxt/>
      </v-container>
    </v-main>
    <v-bottom-navigation v-model="value" fixed>
      <v-btn value="home" to="/c/home">
        <span>ホーム</span>
        <v-icon>mdi-home</v-icon>
      </v-btn>
      <v-btn value="create" to="/c/new-diary">
        <span>日記作成</span>
        <v-icon>mdi-pencil</v-icon>
      </v-btn>
      <v-btn value="reservations" to="/c/reservations">
        <span>予約確認</span>
        <v-icon>mdi-calendar-check</v-icon>
      </v-btn>
      <v-btn value="history" to="/c/histories">
        <span>過去日記</span>
        <v-icon>mdi-history</v-icon>
      </v-btn>
      <v-btn value="settings" to="/c/settings">
        <span>設定</span>
        <v-icon>mdi-cog</v-icon>
      </v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<script>
export default {
  name: 'DefaultLayout',
  data() {
    return {
      value: 'home',
    }
  },
}
</script>
