<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <h1 class="text-center mb-6">設定一覧</h1>

        <v-select
          v-model="selectedUser"
          :items="users"
          item-text="name"
          item-value="id"
          label="ユーザー切り替え"
          outlined
          @change="loadUserSettings"
        ></v-select>

        <v-form ref="form" v-model="valid" @submit.prevent="updateSettings">
          <v-text-field
            v-model="settings.name"
            label="ユーザー名"
            required
          ></v-text-field>

          <v-radio-group v-model="settings.gender" row>
            <v-radio label="男性" value="male"></v-radio>
            <v-radio label="女性" value="female"></v-radio>
            <v-radio label="その他" value="other"></v-radio>
          </v-radio-group>

          <v-text-field
            v-model.number="settings.age"
            label="年齢"
            type="number"
            required
          ></v-text-field>

          <v-combobox
            v-model="settings.hobbies"
            :items="hobbyOptions"
            label="趣味"
            multiple
            chips
            small-chips
            deletable-chips
          ></v-combobox>

          <v-file-input
            v-model="settings.photo"
            accept="image/*"
            label="顔写真"
            prepend-icon="mdi-camera"
          ></v-file-input>

          <v-btn
            color="primary"
            class="mt-4"
            block
            type="submit"
            :disabled="!valid"
          >
            設定を更新
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'IndexPage',
  layout: 'mobile',
  data() {
    return {
      valid: false,
      selectedUser: 1,
      users: [
        { id: 1, name: 'ユーザー1' },
        { id: 2, name: 'ユーザー2' },
        { id: 3, name: 'ユーザー3' },
      ],
      settings: {
        name: '',
        gender: '',
        age: null,
        hobbies: [],
        photo: null,
      },
      hobbyOptions: ['読書', '映画鑑賞', 'スポーツ', '料理', '旅行', '音楽'],
    }
  },
  mounted() {
    this.loadUserSettings()
  },
  methods: {
    loadUserSettings() {
      this.fetchUserSettings(this.selectedUser).then((response) => {
        this.settings = response.data
      })
    },
    updateSettings() {
      if (this.$refs.form.validate()) {
        this.saveUserSettings(this.selectedUser, this.settings).then(() => {
          alert('設定が更新されました')
        })
      }
    },
    fetchUserSettings(userId) {
      return fetch(`http://localhost:8080/users/${userId}`, {
        headers: {
          "Content-Type": "application/json"
        }
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          return { data };
        })
        .catch(error => {
          console.error('There was a problem with the fetch operation:', error);
          throw error;
        });
    },
    saveUserSettings(userId, settings) {
      return new Promise((resolve) => {
        setTimeout(() => {
          console.log(`User ${userId} settings updated:`, settings)
          resolve()
        }, 500)
      })
    },
  },
}
</script>
